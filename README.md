# go-finnhub

[![GoDoc](https://godoc.org/github.com/m1/go-finnhub?status.svg)](https://godoc.org/github.com/m1/go-finnhub)
[![Build Status](https://travis-ci.org/m1/go-finnhub.svg?branch=master)](https://travis-ci.org/m1/go-finnhub)
[![Go Report Card](https://goreportcard.com/badge/github.com/m1/go-finnhub)](https://goreportcard.com/report/github.com/m1/go-finnhub)
[![Release](https://img.shields.io/github/release/m1/go-finnhub.svg)](https://github.com/m1/go-finnhub/releases/latest)
[![codecov](https://codecov.io/gh/m1/go-finnhub/branch/master/graph/badge.svg)](https://codecov.io/gh/m1/go-finnhub)

__Simple and easy to use client for stock, forex and crpyto data from [finnhub.io](https://finnhub.io/) written in Go. Access real-time market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges__

## Installation

`go get github.com/m1/go-finnhub`

## Usage

First sign up for your api token here [finnhub.io](https://finnhub.io/)

Follow this basic example, for more in-depth documentation see the [docs](https://godoc.org/github.com/m1/go-finnhub):
```go
c := client.New("your_token_here")

// Stocks
company, err := c.Stock.GetProfile("AAPL")
ceo, err := c.Stock.GetCEO("AAPL")
recommendation, err := c.Stock.GetRecommendations("AAPL")
target, err := c.Stock.GetPriceTarget("AAPL")
options, err := c.Stock.GetOptionChain("DBD")
peers, err := c.Stock.GetPeers("AAPL")
earnings, err := c.Stock.GetEarnings("AAPL")
candle, err := c.Stock.GetCandle("AAPL", finnhub.CandleResolutionDay, nil)
exchanges, err := c.Stock.GetExchanges()
symbols, err := c.Stock.GetSymbols("US")
gradings, err := c.Stock.GetGradings(&finnhub.GradingParams{Symbol: "AAPL"})

// Crypto
exchanges, err := c.Crypto.GetExchanges()
symbols, err := c.Crypto.GetSymbols("Binance")
candle, err := c.Crypto.GetCandle("BINANCE:BEAMUSDT", finnhub.CandleResolutionMonth, nil)

// Forex
exchanges, err := c.Forex.GetExchanges()
symbols, err := c.Forex.GetSymbols("oanda")
candle, err := c.Forex.GetCandle("OANDA:XAU_GBP", finnhub.CandleResolutionMonth, nil)

// News
news, err := c.News.Get(nil)
news, err = c.News.Get(&finnhub.NewsParams{Category: finnhub.NewsCategoryCrypto})
news, err = c.News.GetCompany("AAPL")
sentiment, err := c.News.GetSentiment("AAPL")
```