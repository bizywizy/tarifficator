package tariff

import (
	"github.com/shopspring/decimal"
	"sort"
)

type repository interface {
	List() ([]Tariff, error)
}

type Comparison struct {
	Tariff     Tariff
	AnnualCost decimal.Decimal
}

func NewComparator(repository repository) *comparator {
	return &comparator{repository: repository}
}

type comparator struct {
	repository repository
}

func (c *comparator) Compare(annualConsumption decimal.Decimal) ([]Comparison, error) {
	tariffs, err := c.repository.List()
	if err != nil {
		return nil, err
	}
	comparisons := make([]Comparison, 0, len(tariffs))
	for _, t := range tariffs {
		comparisons = append(comparisons, Comparison{
			Tariff:     t,
			AnnualCost: t.Calculate(annualConsumption),
		})
	}
	sort.Slice(comparisons, func(i, j int) bool {
		return comparisons[i].AnnualCost.LessThan(comparisons[j].AnnualCost)
	})
	return comparisons, nil
}
