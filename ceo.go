package finnhub

// CEO is the company CEO
type CEO struct {
	Symbol              string `json:"symbol"`
	Name                string `json:"name"`
	CompanyName         string `json:"companyName"`
	Location            string `json:"location"`
	Salary              int    `json:"salary"`
	Bonus               int    `json:"bonus"`
	StockAwards         int    `json:"stockAwards"`
	OptionAwards        int    `json:"optionAwards"`
	NonEquityIncentives int    `json:"nonEquityIncentives"`
	PensionAndDeferred  int    `json:"pensionAndDeferred"`
	OtherComp           int    `json:"otherComp"`
	Total               int    `json:"total"`
	Year                string `json:"year"`
}
