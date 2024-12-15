// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: admin/service/v1/i_router.proto

package servicev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RouterService_ListRoute_FullMethodName          = "/admin.service.v1.RouterService/ListRoute"
	RouterService_ListPermissionCode_FullMethodName = "/admin.service.v1.RouterService/ListPermissionCode"
)

// RouterServiceClient is the client API for RouterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 网站后台动态路由服务
type RouterServiceClient interface {
	// 查询路由列表
	ListRoute(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRouteResponse, error)
	// 查询权限码列表
	ListPermissionCode(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListPermissionCodeResponse, error)
}

type routerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRouterServiceClient(cc grpc.ClientConnInterface) RouterServiceClient {
	return &routerServiceClient{cc}
}

func (c *routerServiceClient) ListRoute(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRouteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRouteResponse)
	err := c.cc.Invoke(ctx, RouterService_ListRoute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerServiceClient) ListPermissionCode(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListPermissionCodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPermissionCodeResponse)
	err := c.cc.Invoke(ctx, RouterService_ListPermissionCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServiceServer is the server API for RouterService service.
// All implementations must embed UnimplementedRouterServiceServer
// for forward compatibility.
//
// 网站后台动态路由服务
type RouterServiceServer interface {
	// 查询路由列表
	ListRoute(context.Context, *emptypb.Empty) (*ListRouteResponse, error)
	// 查询权限码列表
	ListPermissionCode(context.Context, *emptypb.Empty) (*ListPermissionCodeResponse, error)
	mustEmbedUnimplementedRouterServiceServer()
}

// UnimplementedRouterServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRouterServiceServer struct{}

func (UnimplementedRouterServiceServer) ListRoute(context.Context, *emptypb.Empty) (*ListRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoute not implemented")
}
func (UnimplementedRouterServiceServer) ListPermissionCode(context.Context, *emptypb.Empty) (*ListPermissionCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPermissionCode not implemented")
}
func (UnimplementedRouterServiceServer) mustEmbedUnimplementedRouterServiceServer() {}
func (UnimplementedRouterServiceServer) testEmbeddedByValue()                       {}

// UnsafeRouterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouterServiceServer will
// result in compilation errors.
type UnsafeRouterServiceServer interface {
	mustEmbedUnimplementedRouterServiceServer()
}

func RegisterRouterServiceServer(s grpc.ServiceRegistrar, srv RouterServiceServer) {
	// If the following call pancis, it indicates UnimplementedRouterServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RouterService_ServiceDesc, srv)
}

func _RouterService_ListRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServiceServer).ListRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouterService_ListRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServiceServer).ListRoute(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouterService_ListPermissionCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServiceServer).ListPermissionCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouterService_ListPermissionCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServiceServer).ListPermissionCode(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RouterService_ServiceDesc is the grpc.ServiceDesc for RouterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.service.v1.RouterService",
	HandlerType: (*RouterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRoute",
			Handler:    _RouterService_ListRoute_Handler,
		},
		{
			MethodName: "ListPermissionCode",
			Handler:    _RouterService_ListPermissionCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/service/v1/i_router.proto",
}