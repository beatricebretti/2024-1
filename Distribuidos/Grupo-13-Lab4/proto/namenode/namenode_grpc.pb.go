// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: proto/namenode.proto

package namenode

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
	NameNode_DistributeDecisions_FullMethodName = "/namenode.NameNode/DistributeDecisions"
)

// NameNodeClient is the client API for NameNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NameNodeClient interface {
	DistributeDecisions(ctx context.Context, in *DecisionRequest, opts ...grpc.CallOption) (*DecisionResponse, error)
}

type nameNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNameNodeClient(cc grpc.ClientConnInterface) NameNodeClient {
	return &nameNodeClient{cc}
}

func (c *nameNodeClient) DistributeDecisions(ctx context.Context, in *DecisionRequest, opts ...grpc.CallOption) (*DecisionResponse, error) {
	out := new(DecisionResponse)
	err := c.cc.Invoke(ctx, NameNode_DistributeDecisions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NameNodeServer is the server API for NameNode service.
// All implementations must embed UnimplementedNameNodeServer
// for forward compatibility
type NameNodeServer interface {
	DistributeDecisions(context.Context, *DecisionRequest) (*DecisionResponse, error)
	mustEmbedUnimplementedNameNodeServer()
}

// UnimplementedNameNodeServer must be embedded to have forward compatible implementations.
type UnimplementedNameNodeServer struct {
}

func (UnimplementedNameNodeServer) DistributeDecisions(context.Context, *DecisionRequest) (*DecisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DistributeDecisions not implemented")
}
func (UnimplementedNameNodeServer) mustEmbedUnimplementedNameNodeServer() {}

// UnsafeNameNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NameNodeServer will
// result in compilation errors.
type UnsafeNameNodeServer interface {
	mustEmbedUnimplementedNameNodeServer()
}

func RegisterNameNodeServer(s grpc.ServiceRegistrar, srv NameNodeServer) {
	s.RegisterService(&NameNode_ServiceDesc, srv)
}

func _NameNode_DistributeDecisions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServer).DistributeDecisions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NameNode_DistributeDecisions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServer).DistributeDecisions(ctx, req.(*DecisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NameNode_ServiceDesc is the grpc.ServiceDesc for NameNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NameNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "namenode.NameNode",
	HandlerType: (*NameNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DistributeDecisions",
			Handler:    _NameNode_DistributeDecisions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/namenode.proto",
}
