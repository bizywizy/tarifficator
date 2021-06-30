package tariff

import "github.com/shopspring/decimal"

type calculator interface {
	Calculate(decimal decimal.Decimal) decimal.Decimal
}

type Tariff struct {
	Name       string
	Calculator calculator
}

func (t *Tariff) Calculate(annualConsumption decimal.Decimal) decimal.Decimal {
	return t.Calculator.Calculate(annualConsumption)
}
