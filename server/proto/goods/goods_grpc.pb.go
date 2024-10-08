//指定当前proto的语法版本，有2和3之分

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.9
// source: goods.proto

// 指定等会文件生成出来的package

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GoodsService_ListGoods_FullMethodName = "/server.proto.goods.GoodsService/ListGoods"
)

// GoodsServiceClient is the client API for GoodsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// service 服务名称{}
type GoodsServiceClient interface {
	// rpc 服务的函数名 （传入参数）返回（返回参数）
	ListGoods(ctx context.Context, in *ListGoodsReq, opts ...grpc.CallOption) (*ListGoodsRsp, error)
}

type goodsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsServiceClient(cc grpc.ClientConnInterface) GoodsServiceClient {
	return &goodsServiceClient{cc}
}

func (c *goodsServiceClient) ListGoods(ctx context.Context, in *ListGoodsReq, opts ...grpc.CallOption) (*ListGoodsRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListGoodsRsp)
	err := c.cc.Invoke(ctx, GoodsService_ListGoods_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServiceServer is the server API for GoodsService service.
// All implementations should embed UnimplementedGoodsServiceServer
// for forward compatibility.
//
// service 服务名称{}
type GoodsServiceServer interface {
	// rpc 服务的函数名 （传入参数）返回（返回参数）
	ListGoods(context.Context, *ListGoodsReq) (*ListGoodsRsp, error)
}

// UnimplementedGoodsServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGoodsServiceServer struct{}

func (UnimplementedGoodsServiceServer) ListGoods(context.Context, *ListGoodsReq) (*ListGoodsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGoods not implemented")
}
func (UnimplementedGoodsServiceServer) testEmbeddedByValue() {}

// UnsafeGoodsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServiceServer will
// result in compilation errors.
type UnsafeGoodsServiceServer interface {
	mustEmbedUnimplementedGoodsServiceServer()
}

func RegisterGoodsServiceServer(s grpc.ServiceRegistrar, srv GoodsServiceServer) {
	// If the following call pancis, it indicates UnimplementedGoodsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GoodsService_ServiceDesc, srv)
}

func _GoodsService_ListGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGoodsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).ListGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoodsService_ListGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).ListGoods(ctx, req.(*ListGoodsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GoodsService_ServiceDesc is the grpc.ServiceDesc for GoodsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoodsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.proto.goods.GoodsService",
	HandlerType: (*GoodsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListGoods",
			Handler:    _GoodsService_ListGoods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goods.proto",
}
