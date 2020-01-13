package stock

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/m1/go-finnhub"
)

var (
	mockCompany123      = finnhub.Company{Name: "company-123"}
	mockCompanyAbc      = finnhub.Company{Name: "company-abc"}
	mockCeo             = finnhub.CEO{Name: "ceo-123"}
	mockRecommendations = []finnhub.Recommendation{
		{Symbol: "abc"},
		{Symbol: "123"},
	}
	mockPriceTarget = finnhub.PriceTarget{Symbol: "company-123"}
	mockOptionChain = finnhub.OptionChain{Code: "option-chain-123"}
	mockPeers       = []string{
		"peer-123",
		"peer-abc",
	}
	mockEarnings = []finnhub.Earning{
		{Symbol: "abc"},
		{Symbol: "123"},
	}
	mockCandle1               = finnhub.Candle{Status: "status-1"}
	mockCandle2               = finnhub.Candle{Status: "status-2"}
	mockCandle3               = finnhub.Candle{Status: "status-3"}
	mockCandle4               = finnhub.Candle{Status: "status-4"}
	mockCandleNoStatusCompany = finnhub.Company{Name: finnhub.CandleStatusNoData}
	mockExchange1             = finnhub.Exchange{Name: "exchange-1"}
	mockExchange2             = finnhub.Exchange{Name: "exchange-2"}
	mockExchanges             = []finnhub.Exchange{mockExchange1, mockExchange2}
	mockSymbolsExchange1      = []finnhub.Symbol{{Symbol: "company-123"}, {Symbol: "company-abc"}}
	mockSymbolsExchange2      = []finnhub.Symbol{{Symbol: "company-321"}, {Symbol: "company-cba"}}
	mockQuoteCompany123       = finnhub.Quote{Open: 1}
	mockQuoteCompanyAbc       = finnhub.Quote{Open: 2}
	mockGradingsBlank         = []finnhub.Grading{{Symbol: "blank"}}
	mockGradingsCompany123    = []finnhub.Grading{{Symbol: "company-123"}, {Symbol: "company-123"}}
	errMock                   = errors.New("mock")
)

type BackendMock struct{}

func NewBackendMock() *BackendMock {
	return &BackendMock{}
}

func (b BackendMock) Get(path string, params finnhub.URLParams, response interface{}) error {
	switch path {
	case URLProfile:
		company := response.(*finnhub.Company)
		switch params[finnhub.ParamSymbol] {
		case mockCompany123.Name:
			*company = mockCompany123
		case mockCompanyAbc.Name:
			*company = mockCompanyAbc
		default:
			return errMock
		}
	case URLCEOCompensation:
		ceo := response.(*finnhub.CEO)
		*ceo = mockCeo
		switch params[finnhub.ParamSymbol] {
		case mockCeo.Name:
			*ceo = mockCeo
		default:
			return errMock
		}
	case URLRecommendation:
		recommendations := response.(*[]finnhub.Recommendation)
		*recommendations = mockRecommendations
		switch params[finnhub.ParamSymbol] {
		case mockCompany123.Name:
			*recommendations = mockRecommendations
		default:
			return errMock
		}
	case URLPriceTarget:
		target := response.(*finnhub.PriceTarget)
		*target = mockPriceTarget
		switch params[finnhub.ParamSymbol] {
		case mockCompany123.Name:
			*target = mockPriceTarget
		default:
			return errMock
		}
	case URLOptionChain:
		chain := response.(*finnhub.OptionChain)
		*chain = mockOptionChain
		switch params[finnhub.ParamSymbol] {
		case mockCompany123.Name:
			*chain = mockOptionChain
		default:
			return errMock
		}
	case URLPeers:
		peers := response.(*[]string)
		*peers = mockPeers
		switch params[finnhub.ParamSymbol] {
		case mockCompany123.Name:
			*peers = mockPeers
		default:
			return errMock
		}
	case URLEarnings:
		earnings := response.(*[]finnhub.Earning)
		*earnings = mockEarnings
		switch params[finnhub.ParamSymbol] {
		case mockCompany123.Name:
			*earnings = mockEarnings
		default:
			return errMock
		}
	case URLExchange:
		exchanges := response.(*[]finnhub.Exchange)
		*exchanges = mockExchanges
	case URLSymbol:
		symbols := response.(*[]finnhub.Symbol)
		*symbols = mockSymbolsExchange1
		switch params[finnhub.ParamExchange] {
		case mockExchange1.Name:
			*symbols = mockSymbolsExchange1
		case mockExchange2.Name:
			*symbols = mockSymbolsExchange2
		default:
			return errMock
		}
	case URLQuote:
		quote := response.(*finnhub.Quote)
		*quote = mockQuoteCompany123
		switch params[finnhub.ParamSymbol] {
		case mockCompany123.Name:
			*quote = mockQuoteCompany123
		case mockCompanyAbc.Name:
			*quote = mockQuoteCompanyAbc
		default:
			return errMock
		}
	case URLGradings:
		gradings := response.(*[]finnhub.Grading)
		*gradings = mockGradingsBlank
		switch params[finnhub.ParamSymbol] {
		case mockCompany123.Name:
			*gradings = mockGradingsCompany123
		case "":
			*gradings = mockGradingsBlank
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

func TestClient_GetProfile(t *testing.T) {
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
		want    *finnhub.Company
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{symbol: mockCompany123.Name},
			want:   &mockCompany123,
		},
		{
			name:   "valid other symbol",
			fields: fields{API: NewBackendMock()},
			args:   args{symbol: mockCompanyAbc.Name},
			want:   &mockCompanyAbc,
		},
		{
			name:    "error",
			fields:  fields{API: NewBackendMock()},
			args:    args{symbol: "company-err"},
			wantErr: errMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API: tt.fields.API,
			}
			got, err := c.GetProfile(tt.args.symbol)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GetProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && tt.wantErr == nil {
				t.Errorf("GetProfile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetCEO(t *testing.T) {
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
		want    *finnhub.CEO
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{symbol: mockCeo.Name},
			want:   &mockCeo,
		},
		{
			name:    "error",
			fields:  fields{API: NewBackendMock()},
			args:    args{symbol: "company-err"},
			wantErr: errMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API: tt.fields.API,
			}
			got, err := c.GetCEO(tt.args.symbol)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GetCEO() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && tt.wantErr == nil {
				t.Errorf("GetCEO() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetRecommendations(t *testing.T) {
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
		want    []finnhub.Recommendation
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{symbol: mockCompany123.Name},
			want:   mockRecommendations,
		},
		{
			name:    "error",
			fields:  fields{API: NewBackendMock()},
			args:    args{symbol: "company-err"},
			wantErr: errMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API: tt.fields.API,
			}
			got, err := c.GetRecommendations(tt.args.symbol)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GetRecommendations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && tt.wantErr == nil {
				t.Errorf("GetRecommendations() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetPriceTarget(t *testing.T) {
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
		want    *finnhub.PriceTarget
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{symbol: mockCompany123.Name},
			want:   &mockPriceTarget,
		},
		{
			name:    "error",
			fields:  fields{API: NewBackendMock()},
			args:    args{symbol: "company-err"},
			wantErr: errMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API: tt.fields.API,
			}
			got, err := c.GetPriceTarget(tt.args.symbol)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GetPriceTarget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && tt.wantErr == nil {
				t.Errorf("GetPriceTarget() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetOptionChain(t *testing.T) {
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
		want    *finnhub.OptionChain
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{symbol: mockCompany123.Name},
			want:   &mockOptionChain,
		},
		{
			name:    "error",
			fields:  fields{API: NewBackendMock()},
			args:    args{symbol: "company-err"},
			wantErr: errMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API: tt.fields.API,
			}
			got, err := c.GetOptionChain(tt.args.symbol)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GetOptionChain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && tt.wantErr == nil {
				t.Errorf("GetOptionChain() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetPeers(t *testing.T) {
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
		want    []string
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{symbol: mockCompany123.Name},
			want:   mockPeers,
		},
		{
			name:    "error",
			fields:  fields{API: NewBackendMock()},
			args:    args{symbol: "company-err"},
			wantErr: errMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API: tt.fields.API,
			}
			got, err := c.GetPeers(tt.args.symbol)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GetPeers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && tt.wantErr == nil {
				t.Errorf("GetPeers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetEarnings(t *testing.T) {
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
		want    []finnhub.Earning
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{symbol: mockCompany123.Name},
			want:   mockEarnings,
		},
		{
			name:    "error",
			fields:  fields{API: NewBackendMock()},
			args:    args{symbol: "company-err"},
			wantErr: errMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API: tt.fields.API,
			}
			got, err := c.GetEarnings(tt.args.symbol)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GetEarnings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && tt.wantErr == nil {
				t.Errorf("GetEarnings() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetExchanges(t *testing.T) {
	type fields struct {
		API finnhub.Backend
	}
	tests := []struct {
		name    string
		fields  fields
		want    []finnhub.Exchange
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
			want:   mockSymbolsExchange1,
		},
		{
			name:    "err",
			fields:  fields{API: NewBackendMock()},
			args:    args{exchange: ""},
			wantErr: errMock,
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

func TestClient_GetQuote(t *testing.T) {
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
		want    *finnhub.Quote
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{symbol: mockCompany123.Name},
			want:   &mockQuoteCompany123,
		},
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{symbol: mockCompanyAbc.Name},
			want:   &mockQuoteCompanyAbc,
		},
		{
			name:    "err",
			fields:  fields{API: NewBackendMock()},
			args:    args{symbol: ""},
			wantErr: errMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API: tt.fields.API,
			}
			got, err := c.GetQuote(tt.args.symbol)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GetQuote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && tt.wantErr == nil {
				t.Errorf("GetQuote() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetGradings(t *testing.T) {
	type fields struct {
		API finnhub.Backend
	}
	type args struct {
		args *finnhub.GradingParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []finnhub.Grading
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{},
			want:   mockGradingsBlank,
		},
		{
			name:   "valid",
			fields: fields{API: NewBackendMock()},
			args:   args{args: &finnhub.GradingParams{Symbol: mockCompany123.Name}},
			want:   mockGradingsCompany123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API: tt.fields.API,
			}
			got, err := c.GetGradings(tt.args.args)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GetGradings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && tt.wantErr == nil {
				t.Errorf("GetGradings() got = %v, want %v", got, tt.want)
			}
		})
	}
}
