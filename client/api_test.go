package client

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/m1/go-finnhub"
)

type MockResponse struct {
	Path string
	Test string
}

var (
	mockParams = map[string]string{}
)

func NewTestAPI() *API {
	testSrv := http.NewServeMux()
	testSrv.HandleFunc("/v1/TestAPI_Call_valid", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, `{"path":"TestAPI_Call","test":"valid"}`)
	})
	testSrv.HandleFunc("/v1/TestAPI_Call_500", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	testSrv.HandleFunc("/v1/TestAPI_Call_eof", func(w http.ResponseWriter, r *http.Request) {
		hj, ok := w.(http.Hijacker)
		if !ok {
			return
		}
		conn, bufrw, err := hj.Hijack()
		if err != nil {
			return
		}
		bufrw.Flush()
		conn.Close()
	})
	testSrv.HandleFunc("/v1/TestAPI_Call_429", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	})
	testSrv.HandleFunc("/v1/TestAPI_Call_ticker-not-found", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, APIErrTickerNotFound)
	})
	testSrv.HandleFunc("/v1/TestAPI_Call_invalid-json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, `{"err"}`)
	})
	testSrv.HandleFunc("/v1/TestAPI_Call_empty_response", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, `{}`)
	})
	srv := httptest.NewServer(testSrv)
	return &API{endpoint: srv.URL, Key: "token"}
}

func TestAPI_Call(t *testing.T) {
	validParams := mockParams
	type args struct {
		method   string
		path     string
		params   finnhub.URLParams
		response interface{}
	}
	tests := []struct {
		name            string
		args            args
		wantErr         error
		wantErrContains string
		want            interface{}
	}{
		{
			name: "valid",
			args: args{method: "GET", path: "TestAPI_Call_valid", params: validParams, response: &MockResponse{}},
			want: &MockResponse{Path: "TestAPI_Call", Test: "valid"},
		},
		{
			name:    "500",
			args:    args{method: "GET", path: "TestAPI_Call_500", params: validParams, response: &MockResponse{}},
			wantErr: ErrServer,
		},
		{
			name:            "invalid method",
			args:            args{method: ",", path: "TestAPI_Call_invalid", params: validParams, response: &MockResponse{}},
			wantErrContains: "invalid method",
		},
		{
			name:            "eof",
			args:            args{method: "GET", path: "TestAPI_Call_eof", params: validParams, response: &MockResponse{}},
			wantErrContains: "EOF",
		},
		{
			name:    "429 - too many requests",
			args:    args{method: "GET", path: "TestAPI_Call_429", params: validParams, response: &MockResponse{}},
			wantErr: ErrTooManyRequests,
		},
		{
			name:    "ticker not found",
			args:    args{method: "GET", path: "TestAPI_Call_ticker-not-found", params: validParams, response: &MockResponse{}},
			wantErr: ErrTickerNotFound,
		},
		{
			name:    "invalid json",
			args:    args{method: "GET", path: "TestAPI_Call_invalid-json", params: validParams, response: &MockResponse{}},
			wantErr: errors.New("invalid character '}' after object key"),
		},
		{
			name:    "empty response",
			args:    args{method: "GET", path: "TestAPI_Call_empty_response", params: validParams, response: &MockResponse{}},
			wantErr: ErrEmptyResponse,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewTestAPI()
			err := a.Call(tt.args.method, tt.args.path, tt.args.params, tt.args.response)
			if (err != nil && tt.wantErrContains != "") && strings.Contains(err.Error(), tt.wantErrContains) {
				return
			}
			if (err != nil && tt.wantErr != nil) && err.Error() == tt.wantErr.Error() {
				return
			}
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("Call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.args.response, tt.want) && tt.wantErr == nil {
				t.Errorf("Call() got = %v, want %v", tt.args.response, tt.want)
			}
		})
	}
}

func TestAPI_Get(t *testing.T) {
	validParams := mockParams
	type args struct {
		path     string
		params   finnhub.URLParams
		response interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
		want    interface{}
	}{
		{
			name: "valid",
			args: args{path: "TestAPI_Call_valid", params: validParams, response: &MockResponse{}},
			want: &MockResponse{Path: "TestAPI_Call", Test: "valid"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewTestAPI()
			err := a.Get(tt.args.path, tt.args.params, tt.args.response)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.response, tt.want) && tt.wantErr == nil {
				t.Errorf("Get() got = %v, want %v", tt.args.response, tt.want)
			}
		})
	}
}

func TestNewAPI(t *testing.T) {
	type args struct {
		key           string
		clientVersion string
	}
	tests := []struct {
		name string
		args args
		want *API
	}{
		{
			name: "valid",
			args: args{key: "key", clientVersion: "v"},
			want: &API{
				Key:           "key",
				ClientVersion: "v",
				UserAgent:     fmt.Sprintf(UserAgentFmt, "v"),
				Client: http.Client{
					Timeout: time.Second * 30,
				},
				endpoint: APIEndpoint,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAPI(tt.args.key, tt.args.clientVersion); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEmptyResponse(t *testing.T) {
	type args struct {
		p interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "valid",
			args: args{p: MockResponse{}},
			want: MockResponse{},
		},
		{
			name: "valid ptr",
			args: args{p: &MockResponse{}},
			want: &MockResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEmptyResponse(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEmptyResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEmptyResponse(t *testing.T) {
	type args struct {
		response interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid not empty",
			args: args{MockResponse{}},
			want: true,
		},
		{
			name: "valid empty",
			args: args{MockResponse{Path: "test"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEmptyResponse(tt.args.response); got != tt.want {
				t.Errorf("isEmptyResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
