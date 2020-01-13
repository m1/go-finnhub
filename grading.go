package finnhub

import (
	"encoding/json"
	"time"
)

// Grading the data structure for upgrade/downgrade
type Grading struct {
	Symbol    string    `json:"symbol"`
	Company   string    `json:"company"`
	FromGrade string    `json:"fromGrade"`
	ToGrade   string    `json:"toGrade"`
	Action    string    `json:"action"`
	Time      time.Time `json:"gradeTime"`
}

// UnmarshalJSON decodes the json data
func (g *Grading) UnmarshalJSON(data []byte) error {
	type Alias Grading
	price := &struct {
		Time int64 `json:"gradeTime"`
		*Alias
	}{
		Alias: (*Alias)(g),
	}
	var err error
	if err = json.Unmarshal(data, &price); err != nil {
		return err
	}
	g.Time = time.Unix(price.Time, 0)
	return nil
}
