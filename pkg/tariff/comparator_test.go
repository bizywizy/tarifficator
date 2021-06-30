package tariff

import (
	"github.com/shopspring/decimal"
	"testing"
)

type testComparatorRepositoryMock struct {
	returnValue []Tariff
	err         error
}

func (r testComparatorRepositoryMock) List() ([]Tariff, error) {
	if r.err != nil {
		return nil, r.err
	}
	return r.returnValue, nil
}

type testComparatorCalculatorMock struct {
	returnValue decimal.Decimal
}

func (c testComparatorCalculatorMock) Calculate(_ decimal.Decimal) decimal.Decimal {
	return c.returnValue
}

func TestComparator_Compare(t *testing.T) {
	expensiveTariff := Tariff{
		Name:       "expensive",
		Calculator: testComparatorCalculatorMock{decimal.NewFromInt(3)},
	}
	cheapTariff := Tariff{
		Name:       "cheap",
		Calculator: testComparatorCalculatorMock{decimal.NewFromInt(1)},
	}
	middleTariff := Tariff{
		Name:       "middle",
		Calculator: testComparatorCalculatorMock{decimal.NewFromInt(2)},
	}
	repository := testComparatorRepositoryMock{
		returnValue: []Tariff{expensiveTariff, cheapTariff, middleTariff},
	}
	comparator := comparator{repository: repository}

	comparisons, err := comparator.Compare(decimal.Decimal{})
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if (comparisons[0].Tariff != cheapTariff) || (comparisons[1].Tariff != middleTariff) || (comparisons[2].Tariff != expensiveTariff) {
		t.Errorf("wrong order: {%s, %s, %s}", comparisons[0].Tariff.Name, comparisons[1].Tariff.Name, comparisons[2].Tariff.Name)
	}
}
