package finnhub

import (
	"encoding/json"
	"time"
)

// Recommendation the data structure for recommendations
type Recommendation struct {
	Buy        int       `json:"buy"`
	Hold       int       `json:"hold"`
	Period     time.Time `json:"period"`
	Sell       int       `json:"sell"`
	StrongBuy  int       `json:"strongBuy"`
	StrongSell int       `json:"strongSell"`
	Symbol     string    `json:"symbol"`
}

// UnmarshalJSON decodes the json data
func (r *Recommendation) UnmarshalJSON(data []byte) error {
	type Alias Recommendation
	rec := &struct {
		Period string `json:"period"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	var err error
	if err = json.Unmarshal(data, &rec); err != nil {
		return err
	}
	r.Period, err = time.Parse(DateLayoutDate, rec.Period)
	return err
}
