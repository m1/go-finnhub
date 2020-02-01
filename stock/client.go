package stock

import (
	"errors"
	"strconv"

	"github.com/m1/go-finnhub"
)

const (
	// URLProfile profile endpoint url
	URLProfile = "stock/profile"

	// URLCEOCompensation ceo compensation endpoint url
	URLCEOCompensation = "stock/ceo-compensation"

	// URLRecommendation recommnedation endpoint url
	URLRecommendation = "stock/recommendation"

	// URLPriceTarget price target endpoint url
	URLPriceTarget = "stock/price-target"

	// URLOptionChain option chain endpoint url
	URLOptionChain = "stock/option-chain"

	// URLPeers peers endpoint url
	URLPeers = "stock/peers"

	// URLEarnings earnings endpoint url
	URLEarnings = "stock/earnings"

	// URLCandle candle endpoint url
	URLCandle = "stock/candle"

	// URLExchange exchange endpoint url
	URLExchange = "stock/exchange"

	// URLSymbol symbol endpoint url
	URLSymbol = "stock/symbol"

	// URLQuote quote endpoint url
	URLQuote = "quote"

	// URLGradings gradings endpoint url
	URLGradings = "stock/upgrade-downgrade"
)

var (
	// ErrCandlesWrongParams wrong candle params
	ErrCandlesWrongParams = errors.New("wrong candle params - count or from/to must be set")

	// ErrCandleNoData no data for supplied candle
	ErrCandleNoData = errors.New("candle with selected params returned no data - try changing the resolution")
)

// Client returns a new stock client
type Client struct {
	API finnhub.Backend
}

// GetProfile gets general information of a company
func (c *Client) GetProfile(symbol string) (*finnhub.Company, error) {
	var company finnhub.Company
	err := c.API.Get(URLProfile, finnhub.URLParams{finnhub.ParamSymbol: symbol}, &company)
	return &company, err
}

// GetCEO gets latest company's CEO compensation. This endpoint only available
// for US companies
func (c *Client) GetCEO(symbol string) (*finnhub.CEO, error) {
	var ceo finnhub.CEO
	err := c.API.Get(URLCEOCompensation, finnhub.URLParams{finnhub.ParamSymbol: symbol}, &ceo)
	return &ceo, err
}

// GetRecommendations gets latest analyst recommendation trends for a company
func (c *Client) GetRecommendations(symbol string) ([]finnhub.Recommendation, error) {
	var recommendation []finnhub.Recommendation
	err := c.API.Get(URLRecommendation, finnhub.URLParams{finnhub.ParamSymbol: symbol}, &recommendation)
	return recommendation, err
}

// GetPriceTarget gets latest price target consensus
func (c *Client) GetPriceTarget(symbol string) (*finnhub.PriceTarget, error) {
	var target finnhub.PriceTarget
	err := c.API.Get(URLPriceTarget, finnhub.URLParams{finnhub.ParamSymbol: symbol}, &target)
	return &target, err
}

// GetOptionChain gets company option chain. This endpoint only available for
// US companies
func (c *Client) GetOptionChain(symbol string) (*finnhub.OptionChain, error) {
	var options finnhub.OptionChain
	err := c.API.Get(URLOptionChain, finnhub.URLParams{finnhub.ParamSymbol: symbol}, &options)
	return &options, err
}

// GetPeers gets company peers. Return a list of peers in the same country and
// GICS sub-industry
func (c *Client) GetPeers(symbol string) ([]string, error) {
	var peers []string
	err := c.API.Get(URLPeers, finnhub.URLParams{finnhub.ParamSymbol: symbol}, &peers)
	return peers, err
}

// GetEarnings gets company quarterly earnings
func (c *Client) GetEarnings(symbol string) ([]finnhub.Earning, error) {
	var earnings []finnhub.Earning
	err := c.API.Get(URLEarnings, finnhub.URLParams{finnhub.ParamSymbol: symbol}, &earnings)
	return earnings, err
}

// GetCandle gets candlestick data for stocks
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

// GetExchanges lists supported stock exchanges
func (c *Client) GetExchanges() ([]finnhub.Exchange, error) {
	var exchanges []finnhub.Exchange
	err := c.API.Get(URLExchange, finnhub.URLParams{}, &exchanges)
	return exchanges, err
}

// GetSymbols lists supported stocks by exchange
func (c *Client) GetSymbols(exchange string) ([]finnhub.Symbol, error) {
	var symbols []finnhub.Symbol
	err := c.API.Get(URLSymbol, finnhub.URLParams{finnhub.ParamExchange: exchange}, &symbols)
	return symbols, err
}

// GetQuote gets quote data. Constant polling is not recommended. Use websocket
// if you need real-time update
func (c *Client) GetQuote(symbol string) (*finnhub.Quote, error) {
	var quote finnhub.Quote
	err := c.API.Get(URLQuote, finnhub.URLParams{finnhub.ParamSymbol: symbol}, &quote)
	return &quote, err
}

// GetGradings gets latest stock upgrade and downgrade
func (c *Client) GetGradings(args *finnhub.GradingParams) ([]finnhub.Grading, error) {
	var gradings []finnhub.Grading
	params := finnhub.URLParams{}
	if args != nil {
		params[finnhub.ParamSymbol] = args.Symbol
	}
	err := c.API.Get(URLGradings, params, &gradings)
	return gradings, err
}
