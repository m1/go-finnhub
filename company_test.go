package finnhub

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCompany_UnmarshalJSON(t *testing.T) {
	companyData := map[string]interface{}{
		"address":     "address",
		"city":        "city",
		"country":     "country",
		"currency":    "currency",
		"cusip":       "cusip",
		"description": "description",
		"exchange":    "exchange",
		"ggroup":      "ggroup",
		"gind":        "gind",
		"gsector":     "gsector",
		"gsubind":     "gsubind",
		"isin":        "isin",
		"naics":       "naics",
		"name":        "name",
		"phone":       "phone",
		"state":       "state",
		"ticker":      "ticker",
		"weburl":      "http://example.com",
	}

	bytes, err := json.Marshal(&companyData)
	assert.NoError(t, err)

	var company Company
	err = json.Unmarshal(bytes, &company)
	assert.NoError(t, err)
}
