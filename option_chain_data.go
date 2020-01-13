package finnhub

import (
	"encoding/json"
	"time"
)

// OptionChainData the option chain data structure
type OptionChainData struct {
	ExpirationDate time.Time `json:"expirationDate"`
	Options        struct {
		Call []Option `json:"CALL"`
		Put  []Option `json:"PUT"`
	} `json:"options"`
}

// UnmarshalJSON decodes the json data
func (o *OptionChainData) UnmarshalJSON(data []byte) error {
	type Alias OptionChainData
	opt := &struct {
		ExpirationDate string `json:"expirationDate"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	var err error
	if err = json.Unmarshal(data, &opt); err != nil {
		return err
	}

	o.ExpirationDate, err = time.Parse(DateLayoutDate, opt.ExpirationDate)
	if err != nil {
		return err
	}

	return nil
}
