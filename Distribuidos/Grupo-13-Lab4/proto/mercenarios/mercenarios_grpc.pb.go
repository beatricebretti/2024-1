// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: proto/mercenarios.proto

package mercenarios

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
	Mercenario_ReportarPreparacion_FullMethodName = "/mercenarios.Mercenario/ReportarPreparacion"
)

// MercenarioClient is the client API for Mercenario service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MercenarioClient interface {
	// Método para reportar la preparación del mercenario
	ReportarPreparacion(ctx context.Context, in *PreparacionRequest, opts ...grpc.CallOption) (*PreparacionResponse, error)
}

type mercenarioClient struct {
	cc grpc.ClientConnInterface
}

func NewMercenarioClient(cc grpc.ClientConnInterface) MercenarioClient {
	return &mercenarioClient{cc}
}

func (c *mercenarioClient) ReportarPreparacion(ctx context.Context, in *PreparacionRequest, opts ...grpc.CallOption) (*PreparacionResponse, error) {
	out := new(PreparacionResponse)
	err := c.cc.Invoke(ctx, Mercenario_ReportarPreparacion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MercenarioServer is the server API for Mercenario service.
// All implementations must embed UnimplementedMercenarioServer
// for forward compatibility
type MercenarioServer interface {
	// Método para reportar la preparación del mercenario
	ReportarPreparacion(context.Context, *PreparacionRequest) (*PreparacionResponse, error)
	mustEmbedUnimplementedMercenarioServer()
}

// UnimplementedMercenarioServer must be embedded to have forward compatible implementations.
type UnimplementedMercenarioServer struct {
}

func (UnimplementedMercenarioServer) ReportarPreparacion(context.Context, *PreparacionRequest) (*PreparacionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportarPreparacion not implemented")
}
func (UnimplementedMercenarioServer) mustEmbedUnimplementedMercenarioServer() {}

// UnsafeMercenarioServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MercenarioServer will
// result in compilation errors.
type UnsafeMercenarioServer interface {
	mustEmbedUnimplementedMercenarioServer()
}

func RegisterMercenarioServer(s grpc.ServiceRegistrar, srv MercenarioServer) {
	s.RegisterService(&Mercenario_ServiceDesc, srv)
}

func _Mercenario_ReportarPreparacion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreparacionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MercenarioServer).ReportarPreparacion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mercenario_ReportarPreparacion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MercenarioServer).ReportarPreparacion(ctx, req.(*PreparacionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mercenario_ServiceDesc is the grpc.ServiceDesc for Mercenario service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mercenario_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mercenarios.Mercenario",
	HandlerType: (*MercenarioServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportarPreparacion",
			Handler:    _Mercenario_ReportarPreparacion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mercenarios.proto",
}
