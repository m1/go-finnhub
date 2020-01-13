package finnhub

const (
	// ParamToken the param for tokens
	ParamToken = "token"

	// ParamSymbol the param for symbols
	ParamSymbol = "symbol"

	// ParamCount the param for count
	ParamCount = "count"

	// ParamFrom the param for from
	ParamFrom = "from"

	// ParamTo the param for to
	ParamTo = "to"

	// ParamResolution the param for resolution
	ParamResolution = "resolution"

	// ParamExchange the param for exchange
	ParamExchange = "exchange"

	// ParamCategory the param for category
	ParamCategory = "category"
)

// URLParams the data structure to hold the URL parameters
type URLParams map[string]string

// Backend the interface for the API
type Backend interface {
	Get(path string, params URLParams, response interface{}) error
	Call(method string, path string, params URLParams, response interface{}) error
}
