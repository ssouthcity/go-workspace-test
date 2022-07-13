package client

import (
	"context"
	"io"

	"github.com/ssouthcity/go-workspace-test/lib"
	pb "github.com/ssouthcity/go-workspace-test/server/gen"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CarClient struct {
	client pb.CarServiceClient
}

func NewCarClient(conn grpc.ClientConnInterface) *CarClient {
	return &CarClient{pb.NewCarServiceClient(conn)}
}

func (c *CarClient) GetCars() ([]*lib.Car, error) {
	stream, err := c.client.GetCars(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	cars := []*lib.Car{}

	for {
		c, err := stream.Recv()
		if err == io.EOF {
			break
		}

		cars = append(cars, &lib.Car{
			Brand: c.Brand,
			Model: c.Model,
		})
	}

	return cars, nil
}
