package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/ssouthcity/go-workspace-test/server/gen"
	"github.com/ssouthcity/go-workspace-test/server/inmem"
)

func main() {
	lst, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("unable to listen over port: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterCarServiceServer(s, &CarHandler{
		repo: inmem.NewCarRepository(),
	})

	reflection.Register(s)

	if err := s.Serve(lst); err != nil {
		log.Fatalf("unable to serve traffic: %v", err)
	}
}
