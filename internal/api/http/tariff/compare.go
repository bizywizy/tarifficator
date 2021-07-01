package tariff

import (
	"github.com/shopspring/decimal"
	"net/http"
	"tarifficator/internal/api/http/util"
	"tarifficator/pkg/tariff"
)

const (
	consumptionKey = "consumption"

	invalidConsumptionErrMsg = "You have to provide a valid decimal into `consumption` key in query string. like \"1234.56\"."
	unexpectedErrorErrMsg    = "Unexpected error occurred. Try again later."

	invalidErrKey = "invalid"
)

type comparator interface {
	Compare(decimal.Decimal) ([]tariff.Comparison, error)
}

func NewCompareService(comparator comparator) *CompareService {
	return &CompareService{comparator: comparator}
}

type CompareService struct {
	comparator comparator
}

func (s *CompareService) CompareTariffsHandler(w http.ResponseWriter, r *http.Request) {
	consumption, err := decimal.NewFromString(r.URL.Query().Get(consumptionKey))
	if err != nil {
		util.ResponseError(
			w,
			http.StatusBadRequest,
			invalidConsumptionErrMsg,
			util.ErrorDetail{consumptionKey: []string{invalidErrKey}},
		)
		return
	}
	comparisons, err := s.comparator.Compare(consumption)
	if err != nil {
		util.ResponseError(
			w,
			http.StatusInternalServerError,
			unexpectedErrorErrMsg,
			util.ErrorDetail{},
		)
		return
	}
	dtos := make([]comparisonViewModel, len(comparisons))
	for i, cmp := range comparisons {
		dtos[i] = comparisonViewModel{
			Name:       cmp.Tariff.Name,
			AnnualCost: cmp.AnnualCost.String(),
		}
	}

	responseBody := compareTariffsViewModel{
		Total:   len(dtos),
		Objects: dtos,
	}
	util.ResponseJson(w, http.StatusOK, responseBody)
}
