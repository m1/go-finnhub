package finnhub

// Quote the data structure for quotes
type Quote struct {
	Open          float64 `json:"o"`
	High          float64 `json:"h"`
	Low           float64 `json:"l"`
	Current       float64 `json:"c"`
	PreviousClose float64 `json:"pc"`
}
