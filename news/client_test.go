package news

import (
	"errors"
	"reflect"
	"testing"

	"github.com/m1/go-finnhub"
)

var (
	mockNewsNoCat          = []finnhub.News{{ID: 1}}
	mockNewsExample1Cat    = []finnhub.News{{ID: 2}}
	mockNewsExample2Cat    = []finnhub.News{{ID: 3}}
	mockNewsExample1Symbol = []finnhub.News{{ID: 4}}
	mockNewsExample2Symbol = []finnhub.News{{ID: 5}}
	mockSentimentExample1  = finnhub.NewsSentiment{Symbol: "example1"}
	mockSentimentExample2  = finnhub.NewsSentiment{Symbol: "example2"}
	errMock                = errors.New("mock")
)

type BackendMock struct{}

func NewBackendMock() *BackendMock {
	return &BackendMock{}
}

func (b BackendMock) Get(path string, params finnhub.URLParams, response interface{}) error {
	switch path {
	case URLNews:
		news := response.(*[]finnhub.News)
		*news = mockNewsNoCat

		cat, ok := params[finnhub.ParamCategory]
		if ok {
			switch cat {
			case "example1":
				*news = mockNewsExample1Cat
			case "example2":
				*news = mockNewsExample2Cat
			default:
				return errMock
			}
		}

		symbol, ok := params[finnhub.ParamSymbol]
		if ok {
			switch symbol {
			case "example1":
				*news = mockNewsExample1Symbol
			case "example2":
				*news = mockNewsExample2Symbol
			default:
				return errMock
			}
		}

		return nil
	case URLSentiment:
		sentiment := response.(*finnhub.NewsSentiment)
		*sentiment = mockSentimentExample1
		switch params[finnhub.ParamSymbol] {
		case "example1":
			*sentiment = mockSentimentExample1
		case "example2":
			*sentiment = mockSentimentExample2
		default:
			return errMock
		}
	}
	return nil
}

func (b BackendMock) Call(method string, path string, params finnhub.URLParams, response interface{}) error {
	panic("stub")
}

func TestClient_Get(t *testing.T) {
	type fields struct {
		API finnhub.Backend
	}
	type args struct {
		args *finnhub.NewsParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []finnhub.News
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			want:   mockNewsNoCat,
		},
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{args: &finnhub.NewsParams{Category: "example1"}},
			want:   mockNewsExample1Cat,
		},
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{args: &finnhub.NewsParams{Category: "example2"}},
			want:   mockNewsExample2Cat,
		},
		{
			name:    "invalid",
			fields:  fields{API: NewBackendMock()},
			args:    args{args: &finnhub.NewsParams{Category: "example3"}},
			wantErr: errMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API: tt.fields.API,
			}
			got, err := c.Get(tt.args.args)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && tt.wantErr == nil {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetCompany(t *testing.T) {
	type fields struct {
		API finnhub.Backend
	}
	type args struct {
		symbol string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []finnhub.News
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{symbol: "example1"},
			want:   mockNewsExample1Symbol,
		},
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{symbol: "example2"},
			want:   mockNewsExample2Symbol,
		},
		{
			name:    "invalid",
			fields:  fields{API: NewBackendMock()},
			args:    args{symbol: "example3"},
			wantErr: errMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API: tt.fields.API,
			}
			got, err := c.GetCompany(tt.args.symbol)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GetCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && tt.wantErr == nil {
				t.Errorf("GetCompany() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetSentiment(t *testing.T) {
	type fields struct {
		API finnhub.Backend
	}
	type args struct {
		symbol string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    finnhub.NewsSentiment
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{symbol: "example1"},
			want:   mockSentimentExample1,
		},
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{symbol: "example2"},
			want:   mockSentimentExample2,
		},
		{
			name:    "invalid",
			fields:  fields{API: NewBackendMock()},
			args:    args{symbol: "example3"},
			wantErr: errMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API: tt.fields.API,
			}
			got, err := c.GetSentiment(tt.args.symbol)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GetSentiment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && tt.wantErr == nil {
				t.Errorf("GetSentiment() got = %v, want %v", got, tt.want)
			}
		})
	}
}
