package forex

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/m1/go-finnhub"
)

var (
	mockExchange1             = finnhub.Exchange{Name: "exchange-1"}
	mockExchange2             = finnhub.Exchange{Name: "exchange-2"}
	mockExchanges             = []string{mockExchange1.Name, mockExchange2.Name}
	mockExchange1Symbols      = []finnhub.Symbol{{Symbol: "symbol-1"}, {Symbol: "symbol-2"}}
	mockExchange2Symbols      = []finnhub.Symbol{{Symbol: "symbol-3"}, {Symbol: "symbol-4"}}
	mockCandle1               = finnhub.Candle{Status: "status-1"}
	mockCandle2               = finnhub.Candle{Status: "status-2"}
	mockCandle3               = finnhub.Candle{Status: "status-3"}
	mockCandle4               = finnhub.Candle{Status: "status-4"}
	mockCandleNoStatusCompany = finnhub.Company{Name: finnhub.CandleStatusNoData}
	errMock                   = errors.New("mock")
)

type BackendMock struct{}

func NewBackendMock() *BackendMock {
	return &BackendMock{}
}

func (b BackendMock) Get(path string, params finnhub.URLParams, response interface{}) error {
	switch path {
	case URLExchange:
		exchanges := response.(*[]string)
		*exchanges = mockExchanges
	case URLSymbol:
		symbols := response.(*[]finnhub.Symbol)
		*symbols = mockExchange1Symbols
		switch params[finnhub.ParamExchange] {
		case mockExchange1.Name:
			*symbols = mockExchange1Symbols
		case mockExchange2.Name:
			*symbols = mockExchange2Symbols
		default:
			return errMock
		}
	case URLCandle:
		candles := response.(*finnhub.Candle)
		*candles = mockCandle1
		if params[finnhub.ParamCount] == "20" {
			*candles = mockCandle2
		} else if params[finnhub.ParamFrom] == "1576521447" && params[finnhub.ParamTo] == "1576521447" {
			*candles = mockCandle3
		} else if params[finnhub.ParamCount] == strconv.Itoa(finnhub.CandleDefaultCount) {
			*candles = mockCandle4
		} else if params[finnhub.ParamSymbol] == mockCandleNoStatusCompany.Name {
			candles.Status = finnhub.CandleStatusNoData
		}
	}
	return nil
}

func (b BackendMock) Call(method string, path string, params finnhub.URLParams, response interface{}) error {
	panic("stub")
}

func TestClient_GetExchanges(t *testing.T) {
	type fields struct {
		API finnhub.Backend
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			want:   mockExchanges,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API: tt.fields.API,
			}
			got, err := c.GetExchanges()
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GetExchanges() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && tt.wantErr == nil {
				t.Errorf("GetExchanges() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetSymbols(t *testing.T) {
	type fields struct {
		API finnhub.Backend
	}
	type args struct {
		exchange string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []finnhub.Symbol
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{exchange: mockExchange1.Name},
			want:   mockExchange1Symbols,
		},
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{exchange: mockExchange2.Name},
			want:   mockExchange2Symbols,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API: tt.fields.API,
			}
			got, err := c.GetSymbols(tt.args.exchange)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GetSymbols() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && tt.wantErr == nil {
				t.Errorf("GetSymbols() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetCandle(t *testing.T) {
	type fields struct {
		API finnhub.Backend
	}
	type args struct {
		symbol     string
		resolution finnhub.CandleResolution
		args       *finnhub.CandleParams
	}
	count := 20
	count500 := 500
	from := time.Unix(1576521447, 0)
	to := time.Unix(1576521447, 0)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *finnhub.Candle
		wantErr error
	}{
		{
			name:   "valid with count",
			fields: fields{API: NewBackendMock()},
			args: args{args: &finnhub.CandleParams{
				Count: &count,
			}},
			want: &mockCandle2,
		},
		{
			name:   "valid with from/to",
			fields: fields{API: NewBackendMock()},
			args: args{args: &finnhub.CandleParams{
				From: &from,
				To:   &to,
			}},
			want: &mockCandle3,
		},
		{
			name:    "invalid no args",
			fields:  fields{API: NewBackendMock()},
			args:    args{args: &finnhub.CandleParams{}},
			wantErr: ErrCandlesWrongParams,
		},
		{
			name:   "valid with nil",
			fields: fields{API: NewBackendMock()},
			args:   args{args: nil},
			want:   &mockCandle4,
		},
		{
			name:    "candle no status",
			fields:  fields{API: NewBackendMock()},
			args:    args{symbol: mockCandleNoStatusCompany.Name, args: &finnhub.CandleParams{Count: &count500}},
			wantErr: ErrCandleNoData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API: tt.fields.API,
			}
			got, err := c.GetCandle(tt.args.symbol, tt.args.resolution, tt.args.args)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GetCandle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && tt.wantErr == nil {
				t.Errorf("GetCandle() got = %v, want %v", got, tt.want)
			}
		})
	}
}
