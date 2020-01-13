package finnhub

import (
	"encoding/json"
	"strings"
	"time"
)

// Option the data structure for options
type Option struct {
	ContractName      string     `json:"contractName"`
	ContractSize      string     `json:"contractSize"`
	Currency          string     `json:"currency"`
	Type              string     `json:"type"`
	InTheMoney        bool       `json:"inTheMoney"`
	LastTradeDateTime *time.Time `json:"lastTradeDateTime"`
	ExpirationDate    time.Time  `json:"expirationDate"`
	Strike            float64    `json:"strike,string"`
	LastPrice         float64    `json:"lastPrice,string"`
	Bid               float64    `json:"bid,string"`
	Ask               float64    `json:"ask,string"`
	Change            float64    `json:"change,string"`
	ChangePercent     float64    `json:"changePercent,string"`
	Volume            int        `json:"volume"`
	OpenInterest      int        `json:"openInterest"`
	ImpliedVolatility float64    `json:"impliedVolatility,string"`
	Delta             float64    `json:"delta,string"`
	Gamma             float64    `json:"gamma,string"`
	Theta             float64    `json:"theta,string"`
	Vega              float64    `json:"vega,string"`
	Rho               float64    `json:"rho,string"`
	Theoretical       float64    `json:"theoretical,string"`
	IntrinsicValue    float64    `json:"intrinsicValue,string"`
	TimeValue         float64    `json:"timeValue,string"`
	UpdatedAt         time.Time  `json:"updatedAt"`
}

// UnmarshalJSON decodes the json data
func (o *Option) UnmarshalJSON(data []byte) error {
	type Alias Option
	opt := &struct {
		InTheMoney        string `json:"inTheMoney"`
		LastTradeDateTime string `json:"lastTradeDateTime"`
		ExpirationDate    string `json:"expirationDate"`
		UpdatedAt         string `json:"updatedAt"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	var err error
	if err = json.Unmarshal(data, &opt); err != nil {
		return err
	}

	o.InTheMoney = false
	if strings.ToLower(opt.InTheMoney) == "true" {
		o.InTheMoney = true
	}

	lastTrade, _ := time.Parse(DateLayoutDateTime, opt.LastTradeDateTime)
	o.LastTradeDateTime = &lastTrade

	o.ExpirationDate, err = time.Parse(DateLayoutDate, opt.ExpirationDate)
	if err != nil {
		return err
	}

	o.UpdatedAt, err = time.Parse(DateLayoutDateTime, opt.UpdatedAt)
	return err
}
