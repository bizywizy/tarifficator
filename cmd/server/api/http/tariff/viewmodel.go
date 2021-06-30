package tariff

type comparisonViewModel struct {
	Name       string `json:"name"`
	AnnualCost string `json:"annual_cost"`
}

type compareTariffsViewModel struct {
	Total   int                   `json:"total"`
	Objects []comparisonViewModel `json:"objects"`
}
