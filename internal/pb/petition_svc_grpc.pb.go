// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: petition_svc.proto

package pb

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
	PetitionService_CreatePetition_FullMethodName = "/proto.PetitionService/CreatePetition"
)

// PetitionServiceClient is the client API for PetitionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetitionServiceClient interface {
	CreatePetition(ctx context.Context, in *CreatePetitionRequest, opts ...grpc.CallOption) (*CreatePetitionResponse, error)
}

type petitionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPetitionServiceClient(cc grpc.ClientConnInterface) PetitionServiceClient {
	return &petitionServiceClient{cc}
}

func (c *petitionServiceClient) CreatePetition(ctx context.Context, in *CreatePetitionRequest, opts ...grpc.CallOption) (*CreatePetitionResponse, error) {
	out := new(CreatePetitionResponse)
	err := c.cc.Invoke(ctx, PetitionService_CreatePetition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetitionServiceServer is the server API for PetitionService service.
// All implementations should embed UnimplementedPetitionServiceServer
// for forward compatibility
type PetitionServiceServer interface {
	CreatePetition(context.Context, *CreatePetitionRequest) (*CreatePetitionResponse, error)
}

// UnimplementedPetitionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPetitionServiceServer struct {
}

func (UnimplementedPetitionServiceServer) CreatePetition(context.Context, *CreatePetitionRequest) (*CreatePetitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePetition not implemented")
}

// UnsafePetitionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetitionServiceServer will
// result in compilation errors.
type UnsafePetitionServiceServer interface {
	mustEmbedUnimplementedPetitionServiceServer()
}

func RegisterPetitionServiceServer(s grpc.ServiceRegistrar, srv PetitionServiceServer) {
	s.RegisterService(&PetitionService_ServiceDesc, srv)
}

func _PetitionService_CreatePetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePetitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetitionServiceServer).CreatePetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetitionService_CreatePetition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetitionServiceServer).CreatePetition(ctx, req.(*CreatePetitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PetitionService_ServiceDesc is the grpc.ServiceDesc for PetitionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PetitionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PetitionService",
	HandlerType: (*PetitionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePetition",
			Handler:    _PetitionService_CreatePetition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "petition_svc.proto",
}
