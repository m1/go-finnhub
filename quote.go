package finnhub

// Quote the data structure for quotes
type Quote struct {
	Open          int `json:"o"`
	High          int `json:"h"`
	Low           int `json:"l"`
	Current       int `json:"c"`
	PreviousClose int `json:"pc"`
}
