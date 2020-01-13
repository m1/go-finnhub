package finnhub

import (
	"encoding/json"
	"time"
)

// Earning are the earnings for a company
type Earning struct {
	Actual   float64   `json:"actual"`
	Estimate float64   `json:"estimate"`
	Period   time.Time `json:"period"`
	Symbol   string    `json:"symbol"`
}

// UnmarshalJSON decodes the json data
func (e *Earning) UnmarshalJSON(data []byte) error {
	type Alias Earning
	earning := &struct {
		Period string `json:"period"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	var err error
	if err = json.Unmarshal(data, &earning); err != nil {
		return err
	}

	e.Period, err = time.Parse(DateLayoutDate, earning.Period)
	return err
}
