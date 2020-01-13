package finnhub

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestPriceTarget_UnmarshalJSON(t *testing.T) {
	priceTargetData := map[string]interface{}{
		"lastUpdated":  "2006-01-02 15:04:05",
		"symbol":       "symbol",
		"targetHigh":   1,
		"targetLow":    1,
		"targetMean":   1,
		"targetMedian": 1,
	}

	bytes, err := json.Marshal(&priceTargetData)
	assert.NoError(t, err)

	var earning PriceTarget
	err = json.Unmarshal(bytes, &earning)
	assert.NoError(t, err)
}
