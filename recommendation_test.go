package finnhub

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestRecommendation_UnmarshalJSON(t *testing.T) {
	recData := map[string]interface{}{
		"buy":        1,
		"hold":       1,
		"period":     "2006-01-02",
		"sell":       1,
		"strongBuy":  1,
		"strongSell": 1,
		"symbol":     "symbol",
	}

	bytes, err := json.Marshal(&recData)
	assert.NoError(t, err)

	var earning Recommendation
	err = json.Unmarshal(bytes, &earning)
	assert.NoError(t, err)
}
