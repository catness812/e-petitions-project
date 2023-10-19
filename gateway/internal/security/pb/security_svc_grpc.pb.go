// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: security_svc.proto

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

// SecurityServiceClient is the client API for SecurityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecurityServiceClient interface {
	Login(ctx context.Context, in *UserCredentials, opts ...grpc.CallOption) (*Tokens, error)
	RefreshSession(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
	ValidateToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
	SendOTP(ctx context.Context, in *OTPInfo, opts ...grpc.CallOption) (*OTPInfo, error)
	ValidateOTP(ctx context.Context, in *OTPInfo, opts ...grpc.CallOption) (*IsOTPValidated, error)
}

type securityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecurityServiceClient(cc grpc.ClientConnInterface) SecurityServiceClient {
	return &securityServiceClient{cc}
}

func (c *securityServiceClient) Login(ctx context.Context, in *UserCredentials, opts ...grpc.CallOption) (*Tokens, error) {
	out := new(Tokens)
	err := c.cc.Invoke(ctx, "/proto.SecurityService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) RefreshSession(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, "/proto.SecurityService/RefreshSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	out := new(ValidateTokenResponse)
	err := c.cc.Invoke(ctx, "/proto.SecurityService/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) SendOTP(ctx context.Context, in *OTPInfo, opts ...grpc.CallOption) (*OTPInfo, error) {
	out := new(OTPInfo)
	err := c.cc.Invoke(ctx, "/proto.SecurityService/SendOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) ValidateOTP(ctx context.Context, in *OTPInfo, opts ...grpc.CallOption) (*IsOTPValidated, error) {
	out := new(IsOTPValidated)
	err := c.cc.Invoke(ctx, "/proto.SecurityService/ValidateOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityServiceServer is the server API for SecurityService service.
// All implementations must embed UnimplementedSecurityServiceServer
// for forward compatibility
type SecurityServiceServer interface {
	Login(context.Context, *UserCredentials) (*Tokens, error)
	RefreshSession(context.Context, *RefreshRequest) (*RefreshResponse, error)
	ValidateToken(context.Context, *Token) (*ValidateTokenResponse, error)
	SendOTP(context.Context, *OTPInfo) (*OTPInfo, error)
	ValidateOTP(context.Context, *OTPInfo) (*IsOTPValidated, error)
	mustEmbedUnimplementedSecurityServiceServer()
}

// UnimplementedSecurityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecurityServiceServer struct {
}

func (UnimplementedSecurityServiceServer) Login(context.Context, *UserCredentials) (*Tokens, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSecurityServiceServer) RefreshSession(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshSession not implemented")
}
func (UnimplementedSecurityServiceServer) ValidateToken(context.Context, *Token) (*ValidateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedSecurityServiceServer) SendOTP(context.Context, *OTPInfo) (*OTPInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOTP not implemented")
}
func (UnimplementedSecurityServiceServer) ValidateOTP(context.Context, *OTPInfo) (*IsOTPValidated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateOTP not implemented")
}
func (UnimplementedSecurityServiceServer) mustEmbedUnimplementedSecurityServiceServer() {}

// UnsafeSecurityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecurityServiceServer will
// result in compilation errors.
type UnsafeSecurityServiceServer interface {
	mustEmbedUnimplementedSecurityServiceServer()
}

func RegisterSecurityServiceServer(s grpc.ServiceRegistrar, srv SecurityServiceServer) {
	s.RegisterService(&SecurityService_ServiceDesc, srv)
}

func _SecurityService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecurityService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).Login(ctx, req.(*UserCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_RefreshSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).RefreshSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecurityService/RefreshSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).RefreshSession(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecurityService/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).ValidateToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_SendOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OTPInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).SendOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecurityService/SendOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).SendOTP(ctx, req.(*OTPInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_ValidateOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OTPInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).ValidateOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecurityService/ValidateOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).ValidateOTP(ctx, req.(*OTPInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// SecurityService_ServiceDesc is the grpc.ServiceDesc for SecurityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecurityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SecurityService",
	HandlerType: (*SecurityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _SecurityService_Login_Handler,
		},
		{
			MethodName: "RefreshSession",
			Handler:    _SecurityService_RefreshSession_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _SecurityService_ValidateToken_Handler,
		},
		{
			MethodName: "SendOTP",
			Handler:    _SecurityService_SendOTP_Handler,
		},
		{
			MethodName: "ValidateOTP",
			Handler:    _SecurityService_ValidateOTP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "security_svc.proto",
}
