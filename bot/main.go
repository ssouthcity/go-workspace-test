package main

import (
	"log"

	"github.com/ssouthcity/go-workspace-test/server/client"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("server:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial grpc server: %v", err)
	}

	carClient := client.NewCarClient(conn)

	cars, err := carClient.GetCars()
	if err != nil {
		log.Fatalf("failed to get cars: %v", err)
	}

	for _, c := range cars {
		log.Println(c.String())
	}
}
