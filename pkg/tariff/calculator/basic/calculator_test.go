package basic

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestCalculator_Calculate(t *testing.T) {
	calculator := Calculator{
		BaseCost:        decimal.NewFromInt(5),
		ConsumptionCost: decimal.NewFromFloat(0.22),
	}
	testData := []struct {
		input    decimal.Decimal
		expected decimal.Decimal
	}{
		{decimal.NewFromInt(3500), decimal.NewFromInt(830)},
		{decimal.NewFromInt(4500), decimal.NewFromInt(1050)},
		{decimal.NewFromInt(6000), decimal.NewFromInt(1380)},
	}
	for _, data := range testData {
		actual := calculator.Calculate(data.input)
		if !actual.Equal(data.expected) {
			t.Errorf("assertion error Calculate(%s) != %s, actual is %s", data.input, data.expected, actual)
		}
	}
}

func BenchmarkCalculator_Calculate__5__0_22__6000(b *testing.B) {
	benchmarkCalculatorCalculate(
		decimal.NewFromInt(5),
		decimal.NewFromFloat(0.22),
		decimal.NewFromInt(6000),
		b,
	)
}

func benchmarkCalculatorCalculate(baseCost, consCost, annualCons decimal.Decimal, b *testing.B) {
	calculator := Calculator{
		BaseCost:        baseCost,
		ConsumptionCost: consCost,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculator.Calculate(annualCons)
	}
}
