// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: car.proto

package __

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CarServiceClient is the client API for CarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarServiceClient interface {
	GetCars(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (CarService_GetCarsClient, error)
}

type carServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarServiceClient(cc grpc.ClientConnInterface) CarServiceClient {
	return &carServiceClient{cc}
}

func (c *carServiceClient) GetCars(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (CarService_GetCarsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CarService_ServiceDesc.Streams[0], "/api.CarService/GetCars", opts...)
	if err != nil {
		return nil, err
	}
	x := &carServiceGetCarsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CarService_GetCarsClient interface {
	Recv() (*Car, error)
	grpc.ClientStream
}

type carServiceGetCarsClient struct {
	grpc.ClientStream
}

func (x *carServiceGetCarsClient) Recv() (*Car, error) {
	m := new(Car)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CarServiceServer is the server API for CarService service.
// All implementations must embed UnimplementedCarServiceServer
// for forward compatibility
type CarServiceServer interface {
	GetCars(*empty.Empty, CarService_GetCarsServer) error
	mustEmbedUnimplementedCarServiceServer()
}

// UnimplementedCarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCarServiceServer struct {
}

func (UnimplementedCarServiceServer) GetCars(*empty.Empty, CarService_GetCarsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCars not implemented")
}
func (UnimplementedCarServiceServer) mustEmbedUnimplementedCarServiceServer() {}

// UnsafeCarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarServiceServer will
// result in compilation errors.
type UnsafeCarServiceServer interface {
	mustEmbedUnimplementedCarServiceServer()
}

func RegisterCarServiceServer(s grpc.ServiceRegistrar, srv CarServiceServer) {
	s.RegisterService(&CarService_ServiceDesc, srv)
}

func _CarService_GetCars_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CarServiceServer).GetCars(m, &carServiceGetCarsServer{stream})
}

type CarService_GetCarsServer interface {
	Send(*Car) error
	grpc.ServerStream
}

type carServiceGetCarsServer struct {
	grpc.ServerStream
}

func (x *carServiceGetCarsServer) Send(m *Car) error {
	return x.ServerStream.SendMsg(m)
}

// CarService_ServiceDesc is the grpc.ServiceDesc for CarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.CarService",
	HandlerType: (*CarServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCars",
			Handler:       _CarService_GetCars_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "car.proto",
}