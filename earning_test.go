package finnhub

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestEarning_UnmarshalJSON(t *testing.T) {
	earningData := map[string]interface{}{
		"actual":   1,
		"estimate": 1,
		"period":   "2006-01-02",
		"symbol":   "test",
	}

	bytes, err := json.Marshal(&earningData)
	assert.NoError(t, err)

	var earning Earning
	err = json.Unmarshal(bytes, &earning)
	assert.NoError(t, err)
}
