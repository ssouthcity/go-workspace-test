package main

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ssouthcity/go-workspace-test/lib"
	pb "github.com/ssouthcity/go-workspace-test/server/gen"
)

type CarHandler struct {
	pb.CarServiceServer

	repo lib.CarRepository
}

func (h *CarHandler) GetCars(_ *empty.Empty, stream pb.CarService_GetCarsServer) error {
	cars, err := h.repo.GetCars()
	if err != nil {
		return err
	}

	for _, c := range cars {
		if err := stream.Send(&pb.Car{
			Brand: c.Brand,
			Model: c.Model,
		}); err != nil {
			return err
		}
	}

	return nil
}
