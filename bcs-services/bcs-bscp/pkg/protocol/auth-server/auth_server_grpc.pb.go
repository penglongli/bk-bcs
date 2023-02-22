// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: auth_server.proto

package pbas

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

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	// init auth center's auth model.
	InitAuthCenter(ctx context.Context, in *InitAuthCenterReq, opts ...grpc.CallOption) (*InitAuthCenterResp, error)
	// 获取用户鉴权信息
	GetUserInfo(ctx context.Context, in *UserCredentialReq, opts ...grpc.CallOption) (*UserInfoResp, error)
	// iam pull resource callback.
	PullResource(ctx context.Context, in *PullResourceReq, opts ...grpc.CallOption) (*PullResourceResp, error)
	// authorize resource batch.
	AuthorizeBatch(ctx context.Context, in *AuthorizeBatchReq, opts ...grpc.CallOption) (*AuthorizeBatchResp, error)
	// get iam permission to apply.
	GetPermissionToApply(ctx context.Context, in *GetPermissionToApplyReq, opts ...grpc.CallOption) (*GetPermissionToApplyResp, error)
	ListSpace(ctx context.Context, in *ListSpaceReq, opts ...grpc.CallOption) (*ListSpaceResp, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) InitAuthCenter(ctx context.Context, in *InitAuthCenterReq, opts ...grpc.CallOption) (*InitAuthCenterResp, error) {
	out := new(InitAuthCenterResp)
	err := c.cc.Invoke(ctx, "/pbas.Auth/InitAuthCenter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetUserInfo(ctx context.Context, in *UserCredentialReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	out := new(UserInfoResp)
	err := c.cc.Invoke(ctx, "/pbas.Auth/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) PullResource(ctx context.Context, in *PullResourceReq, opts ...grpc.CallOption) (*PullResourceResp, error) {
	out := new(PullResourceResp)
	err := c.cc.Invoke(ctx, "/pbas.Auth/PullResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AuthorizeBatch(ctx context.Context, in *AuthorizeBatchReq, opts ...grpc.CallOption) (*AuthorizeBatchResp, error) {
	out := new(AuthorizeBatchResp)
	err := c.cc.Invoke(ctx, "/pbas.Auth/AuthorizeBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetPermissionToApply(ctx context.Context, in *GetPermissionToApplyReq, opts ...grpc.CallOption) (*GetPermissionToApplyResp, error) {
	out := new(GetPermissionToApplyResp)
	err := c.cc.Invoke(ctx, "/pbas.Auth/GetPermissionToApply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ListSpace(ctx context.Context, in *ListSpaceReq, opts ...grpc.CallOption) (*ListSpaceResp, error) {
	out := new(ListSpaceResp)
	err := c.cc.Invoke(ctx, "/pbas.Auth/ListSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations should embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	// init auth center's auth model.
	InitAuthCenter(context.Context, *InitAuthCenterReq) (*InitAuthCenterResp, error)
	// 获取用户鉴权信息
	GetUserInfo(context.Context, *UserCredentialReq) (*UserInfoResp, error)
	// iam pull resource callback.
	PullResource(context.Context, *PullResourceReq) (*PullResourceResp, error)
	// authorize resource batch.
	AuthorizeBatch(context.Context, *AuthorizeBatchReq) (*AuthorizeBatchResp, error)
	// get iam permission to apply.
	GetPermissionToApply(context.Context, *GetPermissionToApplyReq) (*GetPermissionToApplyResp, error)
	ListSpace(context.Context, *ListSpaceReq) (*ListSpaceResp, error)
}

// UnimplementedAuthServer should be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) InitAuthCenter(context.Context, *InitAuthCenterReq) (*InitAuthCenterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitAuthCenter not implemented")
}
func (UnimplementedAuthServer) GetUserInfo(context.Context, *UserCredentialReq) (*UserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedAuthServer) PullResource(context.Context, *PullResourceReq) (*PullResourceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullResource not implemented")
}
func (UnimplementedAuthServer) AuthorizeBatch(context.Context, *AuthorizeBatchReq) (*AuthorizeBatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeBatch not implemented")
}
func (UnimplementedAuthServer) GetPermissionToApply(context.Context, *GetPermissionToApplyReq) (*GetPermissionToApplyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissionToApply not implemented")
}
func (UnimplementedAuthServer) ListSpace(context.Context, *ListSpaceReq) (*ListSpaceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSpace not implemented")
}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_InitAuthCenter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitAuthCenterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).InitAuthCenter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbas.Auth/InitAuthCenter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).InitAuthCenter(ctx, req.(*InitAuthCenterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCredentialReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbas.Auth/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUserInfo(ctx, req.(*UserCredentialReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_PullResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullResourceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).PullResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbas.Auth/PullResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).PullResource(ctx, req.(*PullResourceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AuthorizeBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeBatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthorizeBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbas.Auth/AuthorizeBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthorizeBatch(ctx, req.(*AuthorizeBatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetPermissionToApply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionToApplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetPermissionToApply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbas.Auth/GetPermissionToApply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetPermissionToApply(ctx, req.(*GetPermissionToApplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ListSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSpaceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ListSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbas.Auth/ListSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ListSpace(ctx, req.(*ListSpaceReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbas.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitAuthCenter",
			Handler:    _Auth_InitAuthCenter_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _Auth_GetUserInfo_Handler,
		},
		{
			MethodName: "PullResource",
			Handler:    _Auth_PullResource_Handler,
		},
		{
			MethodName: "AuthorizeBatch",
			Handler:    _Auth_AuthorizeBatch_Handler,
		},
		{
			MethodName: "GetPermissionToApply",
			Handler:    _Auth_GetPermissionToApply_Handler,
		},
		{
			MethodName: "ListSpace",
			Handler:    _Auth_ListSpace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_server.proto",
}
