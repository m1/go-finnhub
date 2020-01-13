package finnhub

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestOptionChainData_UnmarshalJSON(t *testing.T) {
	optionChainData := map[string]interface{}{
		"expirationDate": "2006-01-02",
		"options": map[string]interface{}{
			"CALL": []map[string]interface{}{},
			"PUT":  []map[string]interface{}{},
		},
	}

	bytes, err := json.Marshal(&optionChainData)
	assert.NoError(t, err)

	var earning OptionChainData
	err = json.Unmarshal(bytes, &earning)
	assert.NoError(t, err)
}
