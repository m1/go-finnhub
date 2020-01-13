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

Follow this basic example, for more in-depth documentation see the [docs](https://godoc.org/github.com/m1/go-finnhub):
```go
finnhub := client.New("bn2iiinrh5rdsulinog0")

// Stocks
company, err := finnhub.Stock.GetProfile("AAPL")
ceo, err := finnhub.Stock.GetCEO("AAPL")
recommendation, err := finnhub.Stock.GetRecommendations("AAPL")
target, err := finnhub.Stock.GetPriceTarget("AAPL")
options, err := finnhub.Stock.GetOptionChain("DBD")
peers, err := finnhub.Stock.GetPeers("AAPL")
earnings, err := finnhub.Stock.GetEarnings("AAPL")
candle, err := finnhub.Stock.GetCandle("AAPL", finnhub2.CandleResolutionDay, nil)
exchanges, err := finnhub.Stock.GetExchanges()
symbols, err := finnhub.Stock.GetSymbols("dddadsad")
gradings, err := finnhub.Stock.GetGradings(&finnhub2.GradingParams{Symbol: "hello"})

// Crypto
exchanges, err := finnhub.Crypto.GetExchanges()
symbols, err := finnhub.Crypto.GetSymbols("Binance")
candle, err := finnhub.Crypto.GetCandle("BINANCE:BEAMUSDT", finnhub2.CandleResolutionMonth, nil)

// Forex
exchanges, err := finnhub.Forex.GetExchanges()
symbols, err := finnhub.Forex.GetSymbols("oanda")
candle, err := finnhub.Forex.GetCandle("OANDA:XAU_GBP", finnhub2.CandleResolutionMonth, nil)

// News
news, err := finnhub.News.Get(nil)
news, err = finnhub.News.Get(&finnhub2.NewsParams{Category: finnhub2.NewsCategoryCrypto})
news, err = finnhub.News.GetCompany("APPL")
sentiment, err := finnhub.News.GetSentiment("AAPL ")
```