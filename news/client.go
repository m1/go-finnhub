package news

import (
	"github.com/m1/go-finnhub"
)

const (
	// URLNews news endpoint url
	URLNews      = "news"

	// URLSentiment sentiment endpoint url
	URLSentiment = "news-sentiment"
)

// Client returns a new news client
type Client struct {
	API finnhub.Backend
}

// Get lists latest market news
func (c *Client) Get(args *finnhub.NewsParams) ([]finnhub.News, error) {
	var news []finnhub.News
	params := finnhub.URLParams{}
	if args != nil {
		params[finnhub.ParamCategory] = args.Category
	}
	err := c.API.Get(URLNews, params, &news)
	return news, err
}

// GetCompany lists latest company news by symbol. This endpoint is only
// available for US companies
func (c *Client) GetCompany(symbol string) ([]finnhub.News, error) {
	var news []finnhub.News
	err := c.API.Get(URLNews, finnhub.URLParams{finnhub.ParamSymbol: symbol}, &news)
	return news, err
}

// GetSentiment lists company's news sentiment and statistics. This endpoint is
// only available for US companies.
func (c *Client) GetSentiment(symbol string) (finnhub.NewsSentiment, error) {
	var sentiment finnhub.NewsSentiment
	err := c.API.Get(URLSentiment, finnhub.URLParams{finnhub.ParamSymbol: symbol}, &sentiment)
	return sentiment, err
}
