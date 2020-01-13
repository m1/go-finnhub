package finnhub

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCandle_UnmarshalJSON(t *testing.T) {
	candleData := map[string]interface{}{
		"c": []float64{1},
		"h": []float64{1},
		"l": []float64{1},
		"o": []float64{1},
		"s": "test",
		"t": []float64{1577836800},
		"v": []float64{1},
	}

	bytes, err := json.Marshal(&candleData)
	assert.NoError(t, err)

	var candle Candle
	err = json.Unmarshal(bytes, &candle)
	assert.NoError(t, err)
}
