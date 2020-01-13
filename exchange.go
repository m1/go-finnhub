package finnhub

// Exchange is the data structure for forex/crypto exchanges
type Exchange struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
