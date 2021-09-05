// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ControlServiceClient is the client API for ControlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControlServiceClient interface {
	// Create a service, and wait for requests to come through via the stream.
	// The caller will then receive a response message for each client that wants
	// to connect to the service.
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (ControlService_CreateServiceClient, error)
	// Forward to remote service.
	ForwardToService(ctx context.Context, in *ForwardToServiceRequest, opts ...grpc.CallOption) (*ForwardToServiceResponse, error)
	// List services.
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
}

type controlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewControlServiceClient(cc grpc.ClientConnInterface) ControlServiceClient {
	return &controlServiceClient{cc}
}

func (c *controlServiceClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (ControlService_CreateServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &ControlService_ServiceDesc.Streams[0], "/mindmeld.ControlService/CreateService", opts...)
	if err != nil {
		return nil, err
	}
	x := &controlServiceCreateServiceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ControlService_CreateServiceClient interface {
	Recv() (*CreateServiceResponse, error)
	grpc.ClientStream
}

type controlServiceCreateServiceClient struct {
	grpc.ClientStream
}

func (x *controlServiceCreateServiceClient) Recv() (*CreateServiceResponse, error) {
	m := new(CreateServiceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controlServiceClient) ForwardToService(ctx context.Context, in *ForwardToServiceRequest, opts ...grpc.CallOption) (*ForwardToServiceResponse, error) {
	out := new(ForwardToServiceResponse)
	err := c.cc.Invoke(ctx, "/mindmeld.ControlService/ForwardToService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, "/mindmeld.ControlService/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlServiceServer is the server API for ControlService service.
// All implementations must embed UnimplementedControlServiceServer
// for forward compatibility
type ControlServiceServer interface {
	// Create a service, and wait for requests to come through via the stream.
	// The caller will then receive a response message for each client that wants
	// to connect to the service.
	CreateService(*CreateServiceRequest, ControlService_CreateServiceServer) error
	// Forward to remote service.
	ForwardToService(context.Context, *ForwardToServiceRequest) (*ForwardToServiceResponse, error)
	// List services.
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	mustEmbedUnimplementedControlServiceServer()
}

// UnimplementedControlServiceServer must be embedded to have forward compatible implementations.
type UnimplementedControlServiceServer struct {
}

func (UnimplementedControlServiceServer) CreateService(*CreateServiceRequest, ControlService_CreateServiceServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedControlServiceServer) ForwardToService(context.Context, *ForwardToServiceRequest) (*ForwardToServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForwardToService not implemented")
}
func (UnimplementedControlServiceServer) ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedControlServiceServer) mustEmbedUnimplementedControlServiceServer() {}

// UnsafeControlServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControlServiceServer will
// result in compilation errors.
type UnsafeControlServiceServer interface {
	mustEmbedUnimplementedControlServiceServer()
}

func RegisterControlServiceServer(s grpc.ServiceRegistrar, srv ControlServiceServer) {
	s.RegisterService(&ControlService_ServiceDesc, srv)
}

func _ControlService_CreateService_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateServiceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControlServiceServer).CreateService(m, &controlServiceCreateServiceServer{stream})
}

type ControlService_CreateServiceServer interface {
	Send(*CreateServiceResponse) error
	grpc.ServerStream
}

type controlServiceCreateServiceServer struct {
	grpc.ServerStream
}

func (x *controlServiceCreateServiceServer) Send(m *CreateServiceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ControlService_ForwardToService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForwardToServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).ForwardToService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindmeld.ControlService/ForwardToService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).ForwardToService(ctx, req.(*ForwardToServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindmeld.ControlService/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ControlService_ServiceDesc is the grpc.ServiceDesc for ControlService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ControlService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mindmeld.ControlService",
	HandlerType: (*ControlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ForwardToService",
			Handler:    _ControlService_ForwardToService_Handler,
		},
		{
			MethodName: "ListServices",
			Handler:    _ControlService_ListServices_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateService",
			Handler:       _ControlService_CreateService_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mindmeld.proto",
}

// ProxyServiceClient is the client API for ProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyServiceClient interface {
	ProxyConnection(ctx context.Context, opts ...grpc.CallOption) (ProxyService_ProxyConnectionClient, error)
}

type proxyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyServiceClient(cc grpc.ClientConnInterface) ProxyServiceClient {
	return &proxyServiceClient{cc}
}

func (c *proxyServiceClient) ProxyConnection(ctx context.Context, opts ...grpc.CallOption) (ProxyService_ProxyConnectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProxyService_ServiceDesc.Streams[0], "/mindmeld.ProxyService/ProxyConnection", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyServiceProxyConnectionClient{stream}
	return x, nil
}

type ProxyService_ProxyConnectionClient interface {
	Send(*Payload) error
	Recv() (*Payload, error)
	grpc.ClientStream
}

type proxyServiceProxyConnectionClient struct {
	grpc.ClientStream
}

func (x *proxyServiceProxyConnectionClient) Send(m *Payload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *proxyServiceProxyConnectionClient) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProxyServiceServer is the server API for ProxyService service.
// All implementations must embed UnimplementedProxyServiceServer
// for forward compatibility
type ProxyServiceServer interface {
	ProxyConnection(ProxyService_ProxyConnectionServer) error
	mustEmbedUnimplementedProxyServiceServer()
}

// UnimplementedProxyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProxyServiceServer struct {
}

func (UnimplementedProxyServiceServer) ProxyConnection(ProxyService_ProxyConnectionServer) error {
	return status.Errorf(codes.Unimplemented, "method ProxyConnection not implemented")
}
func (UnimplementedProxyServiceServer) mustEmbedUnimplementedProxyServiceServer() {}

// UnsafeProxyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyServiceServer will
// result in compilation errors.
type UnsafeProxyServiceServer interface {
	mustEmbedUnimplementedProxyServiceServer()
}

func RegisterProxyServiceServer(s grpc.ServiceRegistrar, srv ProxyServiceServer) {
	s.RegisterService(&ProxyService_ServiceDesc, srv)
}

func _ProxyService_ProxyConnection_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProxyServiceServer).ProxyConnection(&proxyServiceProxyConnectionServer{stream})
}

type ProxyService_ProxyConnectionServer interface {
	Send(*Payload) error
	Recv() (*Payload, error)
	grpc.ServerStream
}

type proxyServiceProxyConnectionServer struct {
	grpc.ServerStream
}

func (x *proxyServiceProxyConnectionServer) Send(m *Payload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *proxyServiceProxyConnectionServer) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProxyService_ServiceDesc is the grpc.ServiceDesc for ProxyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProxyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mindmeld.ProxyService",
	HandlerType: (*ProxyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProxyConnection",
			Handler:       _ProxyService_ProxyConnection_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "mindmeld.proto",
}