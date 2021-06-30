package tariff

import (
	"github.com/shopspring/decimal"
	"testing"
)

type testTariffCalculatorMock struct {
	called      bool
	returnValue decimal.Decimal
}

func (m *testTariffCalculatorMock) Calculate(_ decimal.Decimal) decimal.Decimal {
	m.called = true
	return m.returnValue
}

func TestTariff_Calculate(t *testing.T) {
	mock := testTariffCalculatorMock{
		returnValue: decimal.NewFromInt(322),
	}
	tt := Tariff{
		Name:       "test",
		Calculator: &mock,
	}

	actual := tt.Calculate(decimal.NewFromInt(228))
	if !mock.called {
		t.Error("mocked Calculate method was not called")
	}
	if !actual.Equal(mock.returnValue) {
		t.Errorf("mocked Calculate method was called but return value is wrong, actual: %s", actual)
	}
}
