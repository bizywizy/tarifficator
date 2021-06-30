package packaged

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestCalculator_Calculate(t *testing.T) {
	calculator := Calculator{
		Package:         decimal.NewFromInt(4000),
		PackageCost:     decimal.NewFromInt(800),
		ConsumptionCost: decimal.NewFromFloat(0.3),
	}
	testData := []struct {
		input    decimal.Decimal
		expected decimal.Decimal
	}{
		{decimal.NewFromInt(3500), decimal.NewFromInt(800)},
		{decimal.NewFromInt(4500), decimal.NewFromInt(950)},
		{decimal.NewFromInt(6000), decimal.NewFromInt(1400)},
	}
	for _, data := range testData {
		actual := calculator.Calculate(data.input)
		if !actual.Equal(data.expected) {
			t.Errorf("assertion error Calculate(%s) != %s, actual is %s", data.input, data.expected, actual)
		}
	}
}

func BenchmarkCalculator_Calculate__4000__800__0_3__6000(b *testing.B) {
	benchmarkCalculatorCalculate(
		decimal.NewFromInt(4000),
		decimal.NewFromInt(800),
		decimal.NewFromFloat(0.3),
		decimal.NewFromInt(6000),
		b,
	)
}

func benchmarkCalculatorCalculate(pack, packCost, consCost, annualCons decimal.Decimal, b *testing.B) {
	calculator := Calculator{
		Package:         pack,
		PackageCost:     packCost,
		ConsumptionCost: consCost,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculator.Calculate(annualCons)
	}
}
