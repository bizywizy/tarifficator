package memory

import "tarifficator/pkg/tariff"

func NewRepository() *repository {
	return &repository{}
}

type repository struct {
	storage []tariff.Tariff
}

func (r *repository) Insert(t tariff.Tariff) error {
	r.storage = append(r.storage, t)
	return nil
}

func (r *repository) List() ([]tariff.Tariff, error) {
	return r.storage, nil
}
