package crypto

import (
	"errors"
	"strconv"

	"github.com/m1/go-finnhub"
)

const (
	// URLExchange the url for the exchange endpoint
	URLExchange = "crypto/exchange"

	// URLSymbol the url for the symbol endpoint
	URLSymbol = "crypto/symbol"

	// URLCandle the url for the candle endpoint
	URLCandle = "crypto/candle"
)

var (
	// ErrCandlesWrongParams wrong candle params
	ErrCandlesWrongParams = errors.New("wrong candle params - count or from/to must be set")

	// ErrCandleNoData no data for supplied candle
	ErrCandleNoData = errors.New("candle with selected params returned no data - try changing the resolution")
)

// Client returns a new crypto client
type Client struct {
	API finnhub.Backend
}

// GetExchanges lists supported crypto exchanges
func (c *Client) GetExchanges() ([]string, error) {
	var exchanges []string
	err := c.API.Get(URLExchange, finnhub.URLParams{}, &exchanges)
	return exchanges, err
}

// GetSymbols lists supported crypto symbols by exchange
func (c *Client) GetSymbols(exchange string) ([]finnhub.Symbol, error) {
	var symbols []finnhub.Symbol
	err := c.API.Get(URLSymbol, finnhub.URLParams{finnhub.ParamExchange: exchange}, &symbols)
	return symbols, err
}

// GetCandle gets candlestick data for crypto symbols
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
