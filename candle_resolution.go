package finnhub

// CandleResolution the type for candle resolutions
type CandleResolution int

const (
	// CandleResolutionSecond resolution 1 second
	CandleResolutionSecond CandleResolution = iota

	// CandleResolution5Second resolution 5 seconds
	CandleResolution5Second

	// CandleResolution15Second resolution 15 seconds
	CandleResolution15Second

	// CandleResolution30Second resolution 30 seconds
	CandleResolution30Second

	// CandleResolutionMinute resolution 1 minute
	CandleResolutionMinute

	// CandleResolutionDay resolution 1 day
	CandleResolutionDay

	// CandleResolutionWeek resolution 1 week
	CandleResolutionWeek

	// CandleResolutionMonth resolution 1 month
	CandleResolutionMonth
)

func (d CandleResolution) String() string {
	return [...]string{"1", "5", "15", "30", "60", "D", "W", "M"}[d]
}
