// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: auth.proto

package authv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	AuthWithPassword(ctx context.Context, in *PasswordAuthRequest, opts ...grpc.CallOption) (*AuthToken, error)
	SetPasswordAuth(ctx context.Context, in *SetPasswordAuthMethod, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeletePasswordAuth(ctx context.Context, in *DeletePasswordAuthMethod, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) AuthWithPassword(ctx context.Context, in *PasswordAuthRequest, opts ...grpc.CallOption) (*AuthToken, error) {
	out := new(AuthToken)
	err := c.cc.Invoke(ctx, "/panels.auth.v1.AuthService/AuthWithPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SetPasswordAuth(ctx context.Context, in *SetPasswordAuthMethod, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/panels.auth.v1.AuthService/SetPasswordAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeletePasswordAuth(ctx context.Context, in *DeletePasswordAuthMethod, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/panels.auth.v1.AuthService/DeletePasswordAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	AuthWithPassword(context.Context, *PasswordAuthRequest) (*AuthToken, error)
	SetPasswordAuth(context.Context, *SetPasswordAuthMethod) (*emptypb.Empty, error)
	DeletePasswordAuth(context.Context, *DeletePasswordAuthMethod) (*emptypb.Empty, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) AuthWithPassword(context.Context, *PasswordAuthRequest) (*AuthToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthWithPassword not implemented")
}
func (UnimplementedAuthServiceServer) SetPasswordAuth(context.Context, *SetPasswordAuthMethod) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPasswordAuth not implemented")
}
func (UnimplementedAuthServiceServer) DeletePasswordAuth(context.Context, *DeletePasswordAuthMethod) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePasswordAuth not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_AuthWithPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthWithPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/panels.auth.v1.AuthService/AuthWithPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthWithPassword(ctx, req.(*PasswordAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SetPasswordAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPasswordAuthMethod)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SetPasswordAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/panels.auth.v1.AuthService/SetPasswordAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SetPasswordAuth(ctx, req.(*SetPasswordAuthMethod))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeletePasswordAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePasswordAuthMethod)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeletePasswordAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/panels.auth.v1.AuthService/DeletePasswordAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeletePasswordAuth(ctx, req.(*DeletePasswordAuthMethod))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "panels.auth.v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthWithPassword",
			Handler:    _AuthService_AuthWithPassword_Handler,
		},
		{
			MethodName: "SetPasswordAuth",
			Handler:    _AuthService_SetPasswordAuth_Handler,
		},
		{
			MethodName: "DeletePasswordAuth",
			Handler:    _AuthService_DeletePasswordAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
