package finnhub

// NewsSentiment the data structure for companies news sentiment and statistics
type NewsSentiment struct {
	Buzz struct {
		ArticlesInLastWeek int     `json:"articlesInLastWeek"`
		Buzz               float64 `json:"buzz"`
		WeeklyAverage      float64 `json:"weeklyAverage"`
	} `json:"buzz"`
	CompanyNewsScore            float64 `json:"companyNewsScore"`
	SectorAverageBullishPercent float64 `json:"sectorAverageBullishPercent"`
	SectorAverageNewsScore      float64 `json:"sectorAverageNewsScore"`
	Sentiment                   struct {
		BearishPercent float64 `json:"bearishPercent"`
		BullishPercent float64 `json:"bullishPercent"`
	} `json:"sentiment"`
	Symbol string `json:"symbol"`
}
