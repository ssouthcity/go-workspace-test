package inmem

import "github.com/ssouthcity/go-workspace-test/lib"

type CarRepository struct{}

func NewCarRepository() *CarRepository {
	return &CarRepository{}
}

func (r *CarRepository) GetCars() ([]*lib.Car, error) {
	return []*lib.Car{
		{Brand: "BMW", Model: "i3"},
	}, nil
}
