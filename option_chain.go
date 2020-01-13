package finnhub

// OptionChain the data structure for option chains
type OptionChain struct {
	Code     string            `json:"code"`
	Exchange string            `json:"exchange"`
	Data     []OptionChainData `json:"data"`
}
