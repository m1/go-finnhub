package finnhub

import (
	"encoding/json"
	"time"
)

// PriceTarget the data structure for price targets
type PriceTarget struct {
	LastUpdated  time.Time `json:"lastUpdated"`
	Symbol       string    `json:"symbol"`
	TargetHigh   float64   `json:"targetHigh"`
	TargetLow    float64   `json:"targetLow"`
	TargetMean   float64   `json:"targetMean"`
	TargetMedian float64   `json:"targetMedian"`
}

// UnmarshalJSON decodes the json data
func (p *PriceTarget) UnmarshalJSON(data []byte) error {
	type Alias PriceTarget
	price := &struct {
		LastUpdated string `json:"lastUpdated"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	var err error
	if err = json.Unmarshal(data, &price); err != nil {
		return err
	}
	p.LastUpdated, err = time.Parse(DateLayoutDateTime, price.LastUpdated)
	return err
}
