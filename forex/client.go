package forex

import (
	"errors"
	"strconv"

	"github.com/m1/go-finnhub"
)

const (
	// URLExchange exchange endpoint url
	URLExchange = "forex/exchange"

	// URLSymbol symbol endpoint url
	URLSymbol = "forex/symbol"

	// URLCandle candle endpoint url
	URLCandle = "forex/candle"
)

var (
	// ErrCandlesWrongParams wrong candle params
	ErrCandlesWrongParams = errors.New("wrong candle params - count or from/to must be set")

	// ErrCandleNoData no data for supplied candle
	ErrCandleNoData = errors.New("candle with selected params returned no data - try changing the resolution")
)

// Client returns a new forex client
type Client struct {
	API finnhub.Backend
}

// GetExchanges lists supported forex exchanges
func (c *Client) GetExchanges() ([]string, error) {
	var exchanges []string
	err := c.API.Get(URLExchange, finnhub.URLParams{}, &exchanges)
	return exchanges, err
}

// GetSymbols lists supported forex symbols by exchange
func (c *Client) GetSymbols(exchange string) ([]finnhub.Symbol, error) {
	var symbols []finnhub.Symbol
	err := c.API.Get(URLSymbol, finnhub.URLParams{finnhub.ParamExchange: exchange}, &symbols)
	return symbols, err
}

// GetCandle gets candlestick data for forex symbols
func (c *Client) GetCandle(symbol string, resolution finnhub.CandleResolution, args *finnhub.CandleParams) (*finnhub.Candle, error) {
	params := finnhub.URLParams{finnhub.ParamSymbol: symbol, finnhub.ParamResolution: resolution.String()}
	if args != nil {
		if args.Count != nil {
			params[finnhub.ParamCount] = strconv.Itoa(*args.Count)
		} else if args.From == nil || args.To == nil {
			return nil, ErrCandlesWrongParams
		} else {
			params[finnhub.ParamFrom] = strconv.FormatInt(args.From.Unix(), 10)
			params[finnhub.ParamTo] = strconv.FormatInt(args.To.Unix(), 10)
		}
	} else {
		params[finnhub.ParamCount] = strconv.Itoa(finnhub.CandleDefaultCount)
	}

	var candle finnhub.Candle
	err := c.API.Get(URLCandle, params, &candle)
	if err != nil {
		return nil, err
	}

	if candle.Status == finnhub.CandleStatusNoData {
		return nil, ErrCandleNoData
	}
	return &candle, nil
}
