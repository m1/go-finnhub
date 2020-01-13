package finnhub

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestOption_UnmarshalJSON(t *testing.T) {
	optionData := map[string]interface{}{
		"contractName":      "contract",
		"contractSize":      "size",
		"currency":          "currency",
		"type":              "type",
		"inTheMoney":        "true",
		"lastTradeDateTime": "2006-01-02 15:04:05",
		"expirationDate":    "2006-01-02",
		"strike":            "1",
		"lastPrice":         "1",
		"bid":               "1",
		"ask":               "1",
		"change":            "1",
		"changePercent":     "1",
		"volume":            1,
		"openInterest":      1,
		"impliedVolatility": "1",
		"delta":             "1",
		"gamma":             "1",
		"theta":             "1",
		"vega":              "1",
		"rho":               "1",
		"theoretical":       "1",
		"intrinsicValue":    "1",
		"timeValue":         "1",
		"updatedAt":         "2006-01-02 15:04:05",
	}

	bytes, err := json.Marshal(&optionData)
	assert.NoError(t, err)

	var earning Option
	err = json.Unmarshal(bytes, &earning)
	assert.NoError(t, err)
}
