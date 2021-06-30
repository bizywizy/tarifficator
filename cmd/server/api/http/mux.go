package http

import (
	"github.com/shopspring/decimal"
	"net/http"
	tariffApi "tarifficator/cmd/server/api/http/tariff"
	"tarifficator/pkg/tariff"
)

type comparator interface {
	Compare(decimal.Decimal) ([]tariff.Comparison, error)
}

func NewMux(comparator comparator) *http.ServeMux {
	compareService := tariffApi.NewCompareService(comparator)
	mux := http.NewServeMux()
	mux.HandleFunc("/tariffs/compare", compareService.CompareTariffsHandler)
	return mux
}
