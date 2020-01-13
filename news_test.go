package finnhub

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestNews_UnmarshalJSON(t *testing.T) {
	newsData := map[string]interface{}{
		"category": "technology",
		"datetime": 1567054115,
		"headline": "Facebook acknowledges flaw in Messenger Kids app",
		"id":       25040,
		"image":    "https://example.com",
		"related":  "",
		"source":   "Reuters",
		"summary":  "Facebook Inc  acknowledged a flaw in its Messenger Kids app, weeks after two U.S. senators raised privacy concerns about the application, and said that it spoke to the U.S. Federal Trade Commission about the matter.",
		"url":      "https://example.com",
	}

	bytes, err := json.Marshal(&newsData)
	assert.NoError(t, err)

	var news News
	err = json.Unmarshal(bytes, &news)
	assert.NoError(t, err)

}
