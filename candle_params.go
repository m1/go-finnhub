package finnhub

import (
	"time"
)

// CandleParams the params for the candle endpoint
type CandleParams struct {
	Count *int
	From  *time.Time
	To    *time.Time
}
