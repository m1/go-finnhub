package finnhub

import (
	"encoding/json"
	"time"
)

const (
	// CandleStatusOK candle status ok
	CandleStatusOK     = "ok"

	// CandleStatusNoData candle has no data
	CandleStatusNoData = "no_data"

	// CandleDefaultCount default count for returned candles
	CandleDefaultCount = 200
)

// Candle is candlestick data for stocks
type Candle struct {
	Close  []float64   `json:"c"`
	High   []float64   `json:"h"`
	Low    []float64   `json:"l"`
	Open   []float64   `json:"o"`
	Status string      `json:"s"`
	Times  []time.Time `json:"t"`
	Volume []float64   `json:"v"`
}

// UnmarshalJSON decodes json
func (c *Candle) UnmarshalJSON(data []byte) error {
	type Alias Candle
	candle := &struct {
		Times []int64 `json:"t"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	var err error
	if err = json.Unmarshal(data, &candle); err != nil {
		return err
	}

	for _, t := range candle.Times {
		c.Times = append(c.Times, time.Unix(t, 0))
	}

	return nil
}
