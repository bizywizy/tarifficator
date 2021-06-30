package basic

import "github.com/shopspring/decimal"

var months = decimal.NewFromInt(12)

type Calculator struct {
	BaseCost        decimal.Decimal
	ConsumptionCost decimal.Decimal
}

func (c Calculator) Calculate(annualConsumption decimal.Decimal) decimal.Decimal {
	return c.BaseCost.Mul(months).Add(c.ConsumptionCost.Mul(annualConsumption))
}
