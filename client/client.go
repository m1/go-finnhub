package client

import (
	"github.com/m1/go-finnhub/crypto"
	"github.com/m1/go-finnhub/forex"
	"github.com/m1/go-finnhub/news"
	"github.com/m1/go-finnhub/stock"
)

const (
	// Version the version of the client
	Version = "0.1.0"
)

// Client holds the individual endpoint clients for the api
type Client struct {
	Stock  stock.Client
	Forex  forex.Client
	Crypto crypto.Client
	News   news.Client
}

// New returns a new client
func New(key string) *Client {
	a := NewAPI(key, Version)
	client := &Client{}
	client.Stock = stock.Client{API: a}
	client.Forex = forex.Client{API: a}
	client.Crypto = crypto.Client{API: a}
	client.News = news.Client{API: a}
	return client
}
