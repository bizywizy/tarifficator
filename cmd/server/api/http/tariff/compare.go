package tariff

import (
	"github.com/shopspring/decimal"
	"net/http"
	"tarifficator/cmd/server/api/http/util"
	"tarifficator/pkg/tariff"
)

const (
	consumptionKey = "consumption"

	consumptionKeyNotFoundErrMsg = "You have to provide `consumption` key in query string."
	invalidConsumptionErrMsg     = "`consumption` must be a valid decimal, like \"1234.56\"."
	unexpectedErrorErrMsg        = "Unexpected error occurred. Try again later."

	notFoundErrKey        = "not_found"
	invalidErrKey         = "invalid"
	notValidDecimalErrKey = "not_valid_decimal"
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
	rawConsumption, ok := r.URL.Query()[consumptionKey]
	if !ok {
		util.ResponseError(
			w,
			http.StatusBadRequest,
			consumptionKeyNotFoundErrMsg,
			util.ErrorDetail{consumptionKey: []string{notFoundErrKey}},
		)
		return
	}
	if len(rawConsumption) != 1 {
		util.ResponseError(
			w,
			http.StatusBadRequest,
			invalidConsumptionErrMsg,
			util.ErrorDetail{consumptionKey: []string{invalidErrKey}},
		)
		return
	}

	consumption, err := decimal.NewFromString(rawConsumption[0])
	if err != nil {
		util.ResponseError(
			w,
			http.StatusBadRequest,
			invalidConsumptionErrMsg,
			util.ErrorDetail{consumptionKey: []string{notValidDecimalErrKey}},
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
