// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package memoryserver

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

// MemoryServerClient is the client API for MemoryServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemoryServerClient interface {
	// Sends a greeting
	CheckWithMemoryServer(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*DoIRemember, error)
}

type memoryServerClient struct {
	cc grpc.ClientConnInterface
}

func NewMemoryServerClient(cc grpc.ClientConnInterface) MemoryServerClient {
	return &memoryServerClient{cc}
}

func (c *memoryServerClient) CheckWithMemoryServer(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*DoIRemember, error) {
	out := new(DoIRemember)
	err := c.cc.Invoke(ctx, "/memoryserver.MemoryServer/checkWithMemoryServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemoryServerServer is the server API for MemoryServer service.
// All implementations must embed UnimplementedMemoryServerServer
// for forward compatibility
type MemoryServerServer interface {
	// Sends a greeting
	CheckWithMemoryServer(context.Context, *NumberRequest) (*DoIRemember, error)
	mustEmbedUnimplementedMemoryServerServer()
}

// UnimplementedMemoryServerServer must be embedded to have forward compatible implementations.
type UnimplementedMemoryServerServer struct {
}

func (UnimplementedMemoryServerServer) CheckWithMemoryServer(context.Context, *NumberRequest) (*DoIRemember, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckWithMemoryServer not implemented")
}
func (UnimplementedMemoryServerServer) mustEmbedUnimplementedMemoryServerServer() {}

// UnsafeMemoryServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemoryServerServer will
// result in compilation errors.
type UnsafeMemoryServerServer interface {
	mustEmbedUnimplementedMemoryServerServer()
}

func RegisterMemoryServerServer(s grpc.ServiceRegistrar, srv MemoryServerServer) {
	s.RegisterService(&MemoryServer_ServiceDesc, srv)
}

func _MemoryServer_CheckWithMemoryServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServerServer).CheckWithMemoryServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/memoryserver.MemoryServer/checkWithMemoryServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServerServer).CheckWithMemoryServer(ctx, req.(*NumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MemoryServer_ServiceDesc is the grpc.ServiceDesc for MemoryServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemoryServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "memoryserver.MemoryServer",
	HandlerType: (*MemoryServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "checkWithMemoryServer",
			Handler:    _MemoryServer_CheckWithMemoryServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ms.proto",
}
