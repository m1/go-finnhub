package finnhub

import (
	"encoding/json"
	"net/url"
)

// Company data structure for companies
type Company struct {
	Address           string   `json:"address"`
	City              string   `json:"city"`
	Country           string   `json:"country"`
	Currency          string   `json:"currency"`
	CUSIP             string   `json:"cusip"`
	Description       string   `json:"description"`
	Exchange          string   `json:"exchange"`
	GICSIndustryGroup string   `json:"ggroup"`
	GICSIndustry      string   `json:"gind"`
	GICSSector        string   `json:"gsector"`
	GICSSubIndustry   string   `json:"gsubind"`
	ISIN              string   `json:"isin"`
	NAICS             string   `json:"naics"`
	Name              string   `json:"name"`
	Phone             string   `json:"phone"`
	State             string   `json:"state"`
	Ticker            string   `json:"ticker"`
	WebURL            *url.URL `json:"weburl"`
}

// UnmarshalJSON decodes the json data
func (c *Company) UnmarshalJSON(data []byte) error {
	type Alias Company
	company := &struct {
		WebURL string `json:"weburl"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	var err error
	if err = json.Unmarshal(data, &company); err != nil {
		return err
	}
	c.WebURL, err = url.Parse(company.WebURL)
	if err != nil {
		return err
	}
	return err
}
