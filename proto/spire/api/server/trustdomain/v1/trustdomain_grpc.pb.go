// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package trustdomain

import (
	context "context"
	types "github.com/biareichert/spire-api-sdk/proto/spire/api/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TrustDomainClient is the client API for TrustDomain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrustDomainClient interface {
	// Lists federation relationships with foreign trust domains.
	//
	// The caller must be local or present an admin X509-SVID.
	ListFederationRelationships(ctx context.Context, in *ListFederationRelationshipsRequest, opts ...grpc.CallOption) (*ListFederationRelationshipsResponse, error)
	// Gets a federation relationship with a foreign trust domain.
	// If there is no federation relationship with the specified
	// trust domain, NOT_FOUND is returned.
	//
	// The caller must be local or present an admin X509-SVID.
	GetFederationRelationship(ctx context.Context, in *GetFederationRelationshipRequest, opts ...grpc.CallOption) (*types.FederationRelationship, error)
	// Batch creates one or more federation relationships with
	// foreign trust domains.
	//
	// The caller must be local or present an admin X509-SVID.
	BatchCreateFederationRelationship(ctx context.Context, in *BatchCreateFederationRelationshipRequest, opts ...grpc.CallOption) (*BatchCreateFederationRelationshipResponse, error)
	// Batch updates one or more federation relationships with
	// foreign trust domains.
	//
	// The caller must be local or present an admin X509-SVID.
	BatchUpdateFederationRelationship(ctx context.Context, in *BatchUpdateFederationRelationshipRequest, opts ...grpc.CallOption) (*BatchUpdateFederationRelationshipResponse, error)
	// Batch deletes federation relationships with foreign trust domains.
	//
	// The caller must be local or present an admin X509-SVID.
	BatchDeleteFederationRelationship(ctx context.Context, in *BatchDeleteFederationRelationshipRequest, opts ...grpc.CallOption) (*BatchDeleteFederationRelationshipResponse, error)
	// Refreshes the bundle from the specified federated trust domain.
	// If there is not a federation relationship configured with the
	// specified trust domain, NOT_FOUND is returned.
	//
	// The caller must be local or present an admin X509-SVID.
	RefreshBundle(ctx context.Context, in *RefreshBundleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type trustDomainClient struct {
	cc grpc.ClientConnInterface
}

func NewTrustDomainClient(cc grpc.ClientConnInterface) TrustDomainClient {
	return &trustDomainClient{cc}
}

func (c *trustDomainClient) ListFederationRelationships(ctx context.Context, in *ListFederationRelationshipsRequest, opts ...grpc.CallOption) (*ListFederationRelationshipsResponse, error) {
	out := new(ListFederationRelationshipsResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.trustdomain.v1.TrustDomain/ListFederationRelationships", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustDomainClient) GetFederationRelationship(ctx context.Context, in *GetFederationRelationshipRequest, opts ...grpc.CallOption) (*types.FederationRelationship, error) {
	out := new(types.FederationRelationship)
	err := c.cc.Invoke(ctx, "/spire.api.server.trustdomain.v1.TrustDomain/GetFederationRelationship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustDomainClient) BatchCreateFederationRelationship(ctx context.Context, in *BatchCreateFederationRelationshipRequest, opts ...grpc.CallOption) (*BatchCreateFederationRelationshipResponse, error) {
	out := new(BatchCreateFederationRelationshipResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.trustdomain.v1.TrustDomain/BatchCreateFederationRelationship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustDomainClient) BatchUpdateFederationRelationship(ctx context.Context, in *BatchUpdateFederationRelationshipRequest, opts ...grpc.CallOption) (*BatchUpdateFederationRelationshipResponse, error) {
	out := new(BatchUpdateFederationRelationshipResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.trustdomain.v1.TrustDomain/BatchUpdateFederationRelationship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustDomainClient) BatchDeleteFederationRelationship(ctx context.Context, in *BatchDeleteFederationRelationshipRequest, opts ...grpc.CallOption) (*BatchDeleteFederationRelationshipResponse, error) {
	out := new(BatchDeleteFederationRelationshipResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.trustdomain.v1.TrustDomain/BatchDeleteFederationRelationship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustDomainClient) RefreshBundle(ctx context.Context, in *RefreshBundleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/spire.api.server.trustdomain.v1.TrustDomain/RefreshBundle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrustDomainServer is the server API for TrustDomain service.
// All implementations must embed UnimplementedTrustDomainServer
// for forward compatibility
type TrustDomainServer interface {
	// Lists federation relationships with foreign trust domains.
	//
	// The caller must be local or present an admin X509-SVID.
	ListFederationRelationships(context.Context, *ListFederationRelationshipsRequest) (*ListFederationRelationshipsResponse, error)
	// Gets a federation relationship with a foreign trust domain.
	// If there is no federation relationship with the specified
	// trust domain, NOT_FOUND is returned.
	//
	// The caller must be local or present an admin X509-SVID.
	GetFederationRelationship(context.Context, *GetFederationRelationshipRequest) (*types.FederationRelationship, error)
	// Batch creates one or more federation relationships with
	// foreign trust domains.
	//
	// The caller must be local or present an admin X509-SVID.
	BatchCreateFederationRelationship(context.Context, *BatchCreateFederationRelationshipRequest) (*BatchCreateFederationRelationshipResponse, error)
	// Batch updates one or more federation relationships with
	// foreign trust domains.
	//
	// The caller must be local or present an admin X509-SVID.
	BatchUpdateFederationRelationship(context.Context, *BatchUpdateFederationRelationshipRequest) (*BatchUpdateFederationRelationshipResponse, error)
	// Batch deletes federation relationships with foreign trust domains.
	//
	// The caller must be local or present an admin X509-SVID.
	BatchDeleteFederationRelationship(context.Context, *BatchDeleteFederationRelationshipRequest) (*BatchDeleteFederationRelationshipResponse, error)
	// Refreshes the bundle from the specified federated trust domain.
	// If there is not a federation relationship configured with the
	// specified trust domain, NOT_FOUND is returned.
	//
	// The caller must be local or present an admin X509-SVID.
	RefreshBundle(context.Context, *RefreshBundleRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTrustDomainServer()
}

// UnimplementedTrustDomainServer must be embedded to have forward compatible implementations.
type UnimplementedTrustDomainServer struct {
}

func (UnimplementedTrustDomainServer) ListFederationRelationships(context.Context, *ListFederationRelationshipsRequest) (*ListFederationRelationshipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFederationRelationships not implemented")
}
func (UnimplementedTrustDomainServer) GetFederationRelationship(context.Context, *GetFederationRelationshipRequest) (*types.FederationRelationship, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFederationRelationship not implemented")
}
func (UnimplementedTrustDomainServer) BatchCreateFederationRelationship(context.Context, *BatchCreateFederationRelationshipRequest) (*BatchCreateFederationRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreateFederationRelationship not implemented")
}
func (UnimplementedTrustDomainServer) BatchUpdateFederationRelationship(context.Context, *BatchUpdateFederationRelationshipRequest) (*BatchUpdateFederationRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchUpdateFederationRelationship not implemented")
}
func (UnimplementedTrustDomainServer) BatchDeleteFederationRelationship(context.Context, *BatchDeleteFederationRelationshipRequest) (*BatchDeleteFederationRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteFederationRelationship not implemented")
}
func (UnimplementedTrustDomainServer) RefreshBundle(context.Context, *RefreshBundleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshBundle not implemented")
}
func (UnimplementedTrustDomainServer) mustEmbedUnimplementedTrustDomainServer() {}

// UnsafeTrustDomainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrustDomainServer will
// result in compilation errors.
type UnsafeTrustDomainServer interface {
	mustEmbedUnimplementedTrustDomainServer()
}

func RegisterTrustDomainServer(s grpc.ServiceRegistrar, srv TrustDomainServer) {
	s.RegisterService(&_TrustDomain_serviceDesc, srv)
}

func _TrustDomain_ListFederationRelationships_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFederationRelationshipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustDomainServer).ListFederationRelationships(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.trustdomain.v1.TrustDomain/ListFederationRelationships",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustDomainServer).ListFederationRelationships(ctx, req.(*ListFederationRelationshipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustDomain_GetFederationRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFederationRelationshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustDomainServer).GetFederationRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.trustdomain.v1.TrustDomain/GetFederationRelationship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustDomainServer).GetFederationRelationship(ctx, req.(*GetFederationRelationshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustDomain_BatchCreateFederationRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateFederationRelationshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustDomainServer).BatchCreateFederationRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.trustdomain.v1.TrustDomain/BatchCreateFederationRelationship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustDomainServer).BatchCreateFederationRelationship(ctx, req.(*BatchCreateFederationRelationshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustDomain_BatchUpdateFederationRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchUpdateFederationRelationshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustDomainServer).BatchUpdateFederationRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.trustdomain.v1.TrustDomain/BatchUpdateFederationRelationship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustDomainServer).BatchUpdateFederationRelationship(ctx, req.(*BatchUpdateFederationRelationshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustDomain_BatchDeleteFederationRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchDeleteFederationRelationshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustDomainServer).BatchDeleteFederationRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.trustdomain.v1.TrustDomain/BatchDeleteFederationRelationship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustDomainServer).BatchDeleteFederationRelationship(ctx, req.(*BatchDeleteFederationRelationshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustDomain_RefreshBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustDomainServer).RefreshBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.trustdomain.v1.TrustDomain/RefreshBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustDomainServer).RefreshBundle(ctx, req.(*RefreshBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TrustDomain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.api.server.trustdomain.v1.TrustDomain",
	HandlerType: (*TrustDomainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFederationRelationships",
			Handler:    _TrustDomain_ListFederationRelationships_Handler,
		},
		{
			MethodName: "GetFederationRelationship",
			Handler:    _TrustDomain_GetFederationRelationship_Handler,
		},
		{
			MethodName: "BatchCreateFederationRelationship",
			Handler:    _TrustDomain_BatchCreateFederationRelationship_Handler,
		},
		{
			MethodName: "BatchUpdateFederationRelationship",
			Handler:    _TrustDomain_BatchUpdateFederationRelationship_Handler,
		},
		{
			MethodName: "BatchDeleteFederationRelationship",
			Handler:    _TrustDomain_BatchDeleteFederationRelationship_Handler,
		},
		{
			MethodName: "RefreshBundle",
			Handler:    _TrustDomain_RefreshBundle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spire/api/server/trustdomain/v1/trustdomain.proto",
}
