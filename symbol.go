package finnhub

// Symbol the data structure for symbols
type Symbol struct {
	Description   string `json:"description"`
	DisplaySymbol string `json:"displaySymbol"`
	Symbol        string `json:"symbol"`
}
