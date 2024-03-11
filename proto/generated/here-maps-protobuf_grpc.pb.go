// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: proto/here-maps-protobuf.proto

package generated

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	HereMapsService_GetCoordinates_FullMethodName = "/heremaps.HereMapsService/GetCoordinates"
)

// HereMapsServiceClient is the client API for HereMapsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HereMapsServiceClient interface {
	GetCoordinates(ctx context.Context, in *HereMapsRequest, opts ...grpc.CallOption) (*CoordinatesResponse, error)
}

type hereMapsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHereMapsServiceClient(cc grpc.ClientConnInterface) HereMapsServiceClient {
	return &hereMapsServiceClient{cc}
}

func (c *hereMapsServiceClient) GetCoordinates(ctx context.Context, in *HereMapsRequest, opts ...grpc.CallOption) (*CoordinatesResponse, error) {
	out := new(CoordinatesResponse)
	err := c.cc.Invoke(ctx, HereMapsService_GetCoordinates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HereMapsServiceServer is the server API for HereMapsService service.
// All implementations must embed UnimplementedHereMapsServiceServer
// for forward compatibility
type HereMapsServiceServer interface {
	GetCoordinates(context.Context, *HereMapsRequest) (*CoordinatesResponse, error)
	mustEmbedUnimplementedHereMapsServiceServer()
}

// UnimplementedHereMapsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHereMapsServiceServer struct {
}

func (UnimplementedHereMapsServiceServer) GetCoordinates(context.Context, *HereMapsRequest) (*CoordinatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoordinates not implemented")
}
func (UnimplementedHereMapsServiceServer) mustEmbedUnimplementedHereMapsServiceServer() {}

// UnsafeHereMapsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HereMapsServiceServer will
// result in compilation errors.
type UnsafeHereMapsServiceServer interface {
	mustEmbedUnimplementedHereMapsServiceServer()
}

func RegisterHereMapsServiceServer(s grpc.ServiceRegistrar, srv HereMapsServiceServer) {
	s.RegisterService(&HereMapsService_ServiceDesc, srv)
}

func _HereMapsService_GetCoordinates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HereMapsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HereMapsServiceServer).GetCoordinates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HereMapsService_GetCoordinates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HereMapsServiceServer).GetCoordinates(ctx, req.(*HereMapsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HereMapsService_ServiceDesc is the grpc.ServiceDesc for HereMapsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HereMapsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "heremaps.HereMapsService",
	HandlerType: (*HereMapsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCoordinates",
			Handler:    _HereMapsService_GetCoordinates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/here-maps-protobuf.proto",
}