// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.2
// source: user_svc.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserControllerClient is the client API for UserController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserControllerClient interface {
	CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	CreateUserOTP(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	UpdateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...grpc.CallOption) (*GetUserByEmailResponse, error)
	GetUserEmailById(ctx context.Context, in *GetUserEmailByIdRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	AddAdmin(ctx context.Context, in *AddAdminRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
}

type userControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserControllerClient(cc grpc.ClientConnInterface) UserControllerClient {
	return &userControllerClient{cc}
}

func (c *userControllerClient) CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/rpctransport.UserController/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) CreateUserOTP(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/rpctransport.UserController/CreateUserOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) UpdateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/rpctransport.UserController/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/rpctransport.UserController/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...grpc.CallOption) (*GetUserByEmailResponse, error) {
	out := new(GetUserByEmailResponse)
	err := c.cc.Invoke(ctx, "/rpctransport.UserController/GetUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) GetUserEmailById(ctx context.Context, in *GetUserEmailByIdRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/rpctransport.UserController/GetUserEmailById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) AddAdmin(ctx context.Context, in *AddAdminRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/rpctransport.UserController/AddAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserControllerServer is the server API for UserController service.
// All implementations must embed UnimplementedUserControllerServer
// for forward compatibility
type UserControllerServer interface {
	CreateUser(context.Context, *UserRequest) (*wrapperspb.StringValue, error)
	CreateUserOTP(context.Context, *UserRequest) (*wrapperspb.StringValue, error)
	UpdateUser(context.Context, *UserRequest) (*wrapperspb.StringValue, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*wrapperspb.StringValue, error)
	GetUserByEmail(context.Context, *GetUserByEmailRequest) (*GetUserByEmailResponse, error)
	GetUserEmailById(context.Context, *GetUserEmailByIdRequest) (*wrapperspb.StringValue, error)
	AddAdmin(context.Context, *AddAdminRequest) (*wrapperspb.StringValue, error)
	// mustEmbedUnimplementedUserControllerServer()
}

// UnimplementedUserControllerServer must be embedded to have forward compatible implementations.
type UnimplementedUserControllerServer struct {
}

func (UnimplementedUserControllerServer) CreateUser(context.Context, *UserRequest) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserControllerServer) CreateUserOTP(context.Context, *UserRequest) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserOTP not implemented")
}
func (UnimplementedUserControllerServer) UpdateUser(context.Context, *UserRequest) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserControllerServer) DeleteUser(context.Context, *DeleteUserRequest) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserControllerServer) GetUserByEmail(context.Context, *GetUserByEmailRequest) (*GetUserByEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}
func (UnimplementedUserControllerServer) GetUserEmailById(context.Context, *GetUserEmailByIdRequest) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserEmailById not implemented")
}
func (UnimplementedUserControllerServer) AddAdmin(context.Context, *AddAdminRequest) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAdmin not implemented")
}
func (UnimplementedUserControllerServer) mustEmbedUnimplementedUserControllerServer() {}

// UnsafeUserControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserControllerServer will
// result in compilation errors.
type UnsafeUserControllerServer interface {
	mustEmbedUnimplementedUserControllerServer()
}

func RegisterUserControllerServer(s grpc.ServiceRegistrar, srv UserControllerServer) {
	s.RegisterService(&UserController_ServiceDesc, srv)
}

func _UserController_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpctransport.UserController/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).CreateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_CreateUserOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).CreateUserOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpctransport.UserController/CreateUserOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).CreateUserOTP(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpctransport.UserController/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).UpdateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpctransport.UserController/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpctransport.UserController/GetUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).GetUserByEmail(ctx, req.(*GetUserByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_GetUserEmailById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserEmailByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).GetUserEmailById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpctransport.UserController/GetUserEmailById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).GetUserEmailById(ctx, req.(*GetUserEmailByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_AddAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).AddAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpctransport.UserController/AddAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).AddAdmin(ctx, req.(*AddAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserController_ServiceDesc is the grpc.ServiceDesc for UserController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpctransport.UserController",
	HandlerType: (*UserControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserController_CreateUser_Handler,
		},
		{
			MethodName: "CreateUserOTP",
			Handler:    _UserController_CreateUserOTP_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserController_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserController_DeleteUser_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _UserController_GetUserByEmail_Handler,
		},
		{
			MethodName: "GetUserEmailById",
			Handler:    _UserController_GetUserEmailById_Handler,
		},
		{
			MethodName: "AddAdmin",
			Handler:    _UserController_AddAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_svc.proto",
}
