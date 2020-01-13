package finnhub

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGrading_UnmarshalJSON(t *testing.T) {
	gradingData := map[string]interface{}{
		"symbol":    "test",
		"company":   "test",
		"fromGrade": "from",
		"toGrade":   "to",
		"action":    "up",
		"gradeTime": 1577836800,
	}

	bytes, err := json.Marshal(&gradingData)
	assert.NoError(t, err)

	var earning Grading
	err = json.Unmarshal(bytes, &earning)
	assert.NoError(t, err)
}
