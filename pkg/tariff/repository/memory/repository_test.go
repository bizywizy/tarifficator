package memory

import (
	"tarifficator/pkg/tariff"
	"testing"
)

func TestRepository(t *testing.T) {
	repo := NewRepository()
	expected := []tariff.Tariff{{}, {}}
	for _, tt := range expected {
		err := repo.Insert(tt)
		if err != nil {
			t.Errorf("unexpected err on insert: %v", err)
		}
	}
	actual, err := repo.List()
	if err != nil {
		t.Errorf("unexpected err on list: %v", err)
	}

	for _, expectedTariff := range expected {
		found := false
		for _, actualTariff := range actual {
			if expectedTariff == actualTariff {
				found = true
				break
			}
			if !found {
				t.Errorf("not found expected %s in actual list %v", expectedTariff, actual)
			}
		}
	}
}
