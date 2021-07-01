package main

import (
	"github.com/shopspring/decimal"
	"log"
	"net/http"
	httpApi "tarifficator/internal/api/http"
	"tarifficator/pkg/tariff"
	"tarifficator/pkg/tariff/calculator/basic"
	"tarifficator/pkg/tariff/calculator/packaged"
	"tarifficator/pkg/tariff/repository/memory"
)

func main() {
	tariffRepo := memory.NewRepository()
	_ = tariffRepo.Insert(tariff.Tariff{
		Name: "basic electricity tariff",
		Calculator: basic.Calculator{
			BaseCost:        decimal.NewFromInt(5),
			ConsumptionCost: decimal.NewFromFloat(0.22),
		},
	})
	_ = tariffRepo.Insert(tariff.Tariff{
		Name: "Packaged tariff",
		Calculator: packaged.Calculator{
			Package:         decimal.NewFromInt(4000),
			PackageCost:     decimal.NewFromInt(800),
			ConsumptionCost: decimal.NewFromFloat(0.3),
		},
	})
	tariffComparator := tariff.NewComparator(tariffRepo)

	mux := httpApi.NewMux(tariffComparator)

	log.Fatal(http.ListenAndServe(":8000", mux))
}
