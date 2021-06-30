package packaged

import "github.com/shopspring/decimal"

type Calculator struct {
	Package         decimal.Decimal
	PackageCost     decimal.Decimal
	ConsumptionCost decimal.Decimal
}

func (c Calculator) Calculate(annualConsumption decimal.Decimal) decimal.Decimal {
	annualCost := c.PackageCost
	rest := annualConsumption.Sub(c.Package)
	if !(rest.IsNegative() || rest.IsZero()) {
		annualCost = annualCost.Add(c.ConsumptionCost.Mul(rest))
	}
	return annualCost
}
