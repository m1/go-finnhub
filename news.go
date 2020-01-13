package finnhub

import (
	"encoding/json"
	"net/url"
	"time"
)

// News the data structure for news
type News struct {
	Category string    `json:"category"`
	DateTime time.Time `json:"dateTime"`
	Headline string    `json:"headline"`
	ID       int       `json:"id"`
	Image    *url.URL  `json:"image"`
	Related  string    `json:"related"`
	Source   string    `json:"source"`
	Summary  string    `json:"summary"`
	URL      *url.URL  `json:"url"`
}

// UnmarshalJSON decodes the json data
func (n *News) UnmarshalJSON(data []byte) error {
	type Alias News
	news := &struct {
		DateTime int64  `json:"datetime"`
		Image    string `json:"image"`
		URL      string `json:"url"`
		*Alias
	}{
		Alias: (*Alias)(n),
	}
	var err error
	if err = json.Unmarshal(data, &news); err != nil {
		return err
	}
	n.DateTime = time.Unix(news.DateTime, 0)
	if news.Image != "" {
		n.Image, err = url.Parse(news.Image)

		if err != nil {
			return err
		}
	}
	if news.URL != "" {
		n.URL, err = url.Parse(news.URL)
		if err != nil {
			return err
		}
	}
	return nil
}
