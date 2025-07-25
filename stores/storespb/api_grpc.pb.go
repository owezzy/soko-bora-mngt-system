// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: storespb/api.proto

package storespb

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

// StoresServiceClient is the client API for StoresService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoresServiceClient interface {
	CreateStore(ctx context.Context, in *CreateStoreRequest, opts ...grpc.CallOption) (*CreateStoreResponse, error)
	EnableParticipation(ctx context.Context, in *EnableParticipationRequest, opts ...grpc.CallOption) (*EnableParticipationResponse, error)
	DisableParticipation(ctx context.Context, in *DisableParticipationRequest, opts ...grpc.CallOption) (*DisableParticipationResponse, error)
	RebrandStore(ctx context.Context, in *RebrandStoreRequest, opts ...grpc.CallOption) (*RebrandStoreResponse, error)
	GetStore(ctx context.Context, in *GetStoreRequest, opts ...grpc.CallOption) (*GetStoreResponse, error)
	GetStores(ctx context.Context, in *GetStoresRequest, opts ...grpc.CallOption) (*GetStoresResponse, error)
	GetParticipatingStores(ctx context.Context, in *GetParticipatingStoresRequest, opts ...grpc.CallOption) (*GetParticipatingStoresResponse, error)
	AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error)
	RebrandProduct(ctx context.Context, in *RebrandProductRequest, opts ...grpc.CallOption) (*RebrandProductResponse, error)
	IncreaseProductPrice(ctx context.Context, in *IncreaseProductPriceRequest, opts ...grpc.CallOption) (*IncreaseProductPriceResponse, error)
	DecreaseProductPrice(ctx context.Context, in *DecreaseProductPriceRequest, opts ...grpc.CallOption) (*DecreaseProductPriceResponse, error)
	RemoveProduct(ctx context.Context, in *RemoveProductRequest, opts ...grpc.CallOption) (*RemoveProductResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
	GetCatalog(ctx context.Context, in *GetCatalogRequest, opts ...grpc.CallOption) (*GetCatalogResponse, error)
}

type storesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoresServiceClient(cc grpc.ClientConnInterface) StoresServiceClient {
	return &storesServiceClient{cc}
}

func (c *storesServiceClient) CreateStore(ctx context.Context, in *CreateStoreRequest, opts ...grpc.CallOption) (*CreateStoreResponse, error) {
	out := new(CreateStoreResponse)
	err := c.cc.Invoke(ctx, "/storespb.StoresService/CreateStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesServiceClient) EnableParticipation(ctx context.Context, in *EnableParticipationRequest, opts ...grpc.CallOption) (*EnableParticipationResponse, error) {
	out := new(EnableParticipationResponse)
	err := c.cc.Invoke(ctx, "/storespb.StoresService/EnableParticipation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesServiceClient) DisableParticipation(ctx context.Context, in *DisableParticipationRequest, opts ...grpc.CallOption) (*DisableParticipationResponse, error) {
	out := new(DisableParticipationResponse)
	err := c.cc.Invoke(ctx, "/storespb.StoresService/DisableParticipation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesServiceClient) RebrandStore(ctx context.Context, in *RebrandStoreRequest, opts ...grpc.CallOption) (*RebrandStoreResponse, error) {
	out := new(RebrandStoreResponse)
	err := c.cc.Invoke(ctx, "/storespb.StoresService/RebrandStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesServiceClient) GetStore(ctx context.Context, in *GetStoreRequest, opts ...grpc.CallOption) (*GetStoreResponse, error) {
	out := new(GetStoreResponse)
	err := c.cc.Invoke(ctx, "/storespb.StoresService/GetStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesServiceClient) GetStores(ctx context.Context, in *GetStoresRequest, opts ...grpc.CallOption) (*GetStoresResponse, error) {
	out := new(GetStoresResponse)
	err := c.cc.Invoke(ctx, "/storespb.StoresService/GetStores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesServiceClient) GetParticipatingStores(ctx context.Context, in *GetParticipatingStoresRequest, opts ...grpc.CallOption) (*GetParticipatingStoresResponse, error) {
	out := new(GetParticipatingStoresResponse)
	err := c.cc.Invoke(ctx, "/storespb.StoresService/GetParticipatingStores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesServiceClient) AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error) {
	out := new(AddProductResponse)
	err := c.cc.Invoke(ctx, "/storespb.StoresService/AddProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesServiceClient) RebrandProduct(ctx context.Context, in *RebrandProductRequest, opts ...grpc.CallOption) (*RebrandProductResponse, error) {
	out := new(RebrandProductResponse)
	err := c.cc.Invoke(ctx, "/storespb.StoresService/RebrandProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesServiceClient) IncreaseProductPrice(ctx context.Context, in *IncreaseProductPriceRequest, opts ...grpc.CallOption) (*IncreaseProductPriceResponse, error) {
	out := new(IncreaseProductPriceResponse)
	err := c.cc.Invoke(ctx, "/storespb.StoresService/IncreaseProductPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesServiceClient) DecreaseProductPrice(ctx context.Context, in *DecreaseProductPriceRequest, opts ...grpc.CallOption) (*DecreaseProductPriceResponse, error) {
	out := new(DecreaseProductPriceResponse)
	err := c.cc.Invoke(ctx, "/storespb.StoresService/DecreaseProductPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesServiceClient) RemoveProduct(ctx context.Context, in *RemoveProductRequest, opts ...grpc.CallOption) (*RemoveProductResponse, error) {
	out := new(RemoveProductResponse)
	err := c.cc.Invoke(ctx, "/storespb.StoresService/RemoveProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, "/storespb.StoresService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesServiceClient) GetCatalog(ctx context.Context, in *GetCatalogRequest, opts ...grpc.CallOption) (*GetCatalogResponse, error) {
	out := new(GetCatalogResponse)
	err := c.cc.Invoke(ctx, "/storespb.StoresService/GetCatalog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoresServiceServer is the server API for StoresService service.
// All implementations must embed UnimplementedStoresServiceServer
// for forward compatibility
type StoresServiceServer interface {
	CreateStore(context.Context, *CreateStoreRequest) (*CreateStoreResponse, error)
	EnableParticipation(context.Context, *EnableParticipationRequest) (*EnableParticipationResponse, error)
	DisableParticipation(context.Context, *DisableParticipationRequest) (*DisableParticipationResponse, error)
	RebrandStore(context.Context, *RebrandStoreRequest) (*RebrandStoreResponse, error)
	GetStore(context.Context, *GetStoreRequest) (*GetStoreResponse, error)
	GetStores(context.Context, *GetStoresRequest) (*GetStoresResponse, error)
	GetParticipatingStores(context.Context, *GetParticipatingStoresRequest) (*GetParticipatingStoresResponse, error)
	AddProduct(context.Context, *AddProductRequest) (*AddProductResponse, error)
	RebrandProduct(context.Context, *RebrandProductRequest) (*RebrandProductResponse, error)
	IncreaseProductPrice(context.Context, *IncreaseProductPriceRequest) (*IncreaseProductPriceResponse, error)
	DecreaseProductPrice(context.Context, *DecreaseProductPriceRequest) (*DecreaseProductPriceResponse, error)
	RemoveProduct(context.Context, *RemoveProductRequest) (*RemoveProductResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
	GetCatalog(context.Context, *GetCatalogRequest) (*GetCatalogResponse, error)
	mustEmbedUnimplementedStoresServiceServer()
}

// UnimplementedStoresServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStoresServiceServer struct {
}

func (UnimplementedStoresServiceServer) CreateStore(context.Context, *CreateStoreRequest) (*CreateStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStore not implemented")
}
func (UnimplementedStoresServiceServer) EnableParticipation(context.Context, *EnableParticipationRequest) (*EnableParticipationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableParticipation not implemented")
}
func (UnimplementedStoresServiceServer) DisableParticipation(context.Context, *DisableParticipationRequest) (*DisableParticipationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableParticipation not implemented")
}
func (UnimplementedStoresServiceServer) RebrandStore(context.Context, *RebrandStoreRequest) (*RebrandStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RebrandStore not implemented")
}
func (UnimplementedStoresServiceServer) GetStore(context.Context, *GetStoreRequest) (*GetStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStore not implemented")
}
func (UnimplementedStoresServiceServer) GetStores(context.Context, *GetStoresRequest) (*GetStoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStores not implemented")
}
func (UnimplementedStoresServiceServer) GetParticipatingStores(context.Context, *GetParticipatingStoresRequest) (*GetParticipatingStoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParticipatingStores not implemented")
}
func (UnimplementedStoresServiceServer) AddProduct(context.Context, *AddProductRequest) (*AddProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedStoresServiceServer) RebrandProduct(context.Context, *RebrandProductRequest) (*RebrandProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RebrandProduct not implemented")
}
func (UnimplementedStoresServiceServer) IncreaseProductPrice(context.Context, *IncreaseProductPriceRequest) (*IncreaseProductPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncreaseProductPrice not implemented")
}
func (UnimplementedStoresServiceServer) DecreaseProductPrice(context.Context, *DecreaseProductPriceRequest) (*DecreaseProductPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecreaseProductPrice not implemented")
}
func (UnimplementedStoresServiceServer) RemoveProduct(context.Context, *RemoveProductRequest) (*RemoveProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProduct not implemented")
}
func (UnimplementedStoresServiceServer) GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedStoresServiceServer) GetCatalog(context.Context, *GetCatalogRequest) (*GetCatalogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCatalog not implemented")
}
func (UnimplementedStoresServiceServer) mustEmbedUnimplementedStoresServiceServer() {}

// UnsafeStoresServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoresServiceServer will
// result in compilation errors.
type UnsafeStoresServiceServer interface {
	mustEmbedUnimplementedStoresServiceServer()
}

func RegisterStoresServiceServer(s grpc.ServiceRegistrar, srv StoresServiceServer) {
	s.RegisterService(&StoresService_ServiceDesc, srv)
}

func _StoresService_CreateStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServiceServer).CreateStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storespb.StoresService/CreateStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServiceServer).CreateStore(ctx, req.(*CreateStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoresService_EnableParticipation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableParticipationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServiceServer).EnableParticipation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storespb.StoresService/EnableParticipation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServiceServer).EnableParticipation(ctx, req.(*EnableParticipationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoresService_DisableParticipation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableParticipationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServiceServer).DisableParticipation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storespb.StoresService/DisableParticipation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServiceServer).DisableParticipation(ctx, req.(*DisableParticipationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoresService_RebrandStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RebrandStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServiceServer).RebrandStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storespb.StoresService/RebrandStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServiceServer).RebrandStore(ctx, req.(*RebrandStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoresService_GetStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServiceServer).GetStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storespb.StoresService/GetStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServiceServer).GetStore(ctx, req.(*GetStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoresService_GetStores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServiceServer).GetStores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storespb.StoresService/GetStores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServiceServer).GetStores(ctx, req.(*GetStoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoresService_GetParticipatingStores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParticipatingStoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServiceServer).GetParticipatingStores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storespb.StoresService/GetParticipatingStores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServiceServer).GetParticipatingStores(ctx, req.(*GetParticipatingStoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoresService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storespb.StoresService/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServiceServer).AddProduct(ctx, req.(*AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoresService_RebrandProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RebrandProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServiceServer).RebrandProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storespb.StoresService/RebrandProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServiceServer).RebrandProduct(ctx, req.(*RebrandProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoresService_IncreaseProductPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncreaseProductPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServiceServer).IncreaseProductPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storespb.StoresService/IncreaseProductPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServiceServer).IncreaseProductPrice(ctx, req.(*IncreaseProductPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoresService_DecreaseProductPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecreaseProductPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServiceServer).DecreaseProductPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storespb.StoresService/DecreaseProductPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServiceServer).DecreaseProductPrice(ctx, req.(*DecreaseProductPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoresService_RemoveProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServiceServer).RemoveProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storespb.StoresService/RemoveProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServiceServer).RemoveProduct(ctx, req.(*RemoveProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoresService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storespb.StoresService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoresService_GetCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServiceServer).GetCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storespb.StoresService/GetCatalog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServiceServer).GetCatalog(ctx, req.(*GetCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StoresService_ServiceDesc is the grpc.ServiceDesc for StoresService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoresService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storespb.StoresService",
	HandlerType: (*StoresServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStore",
			Handler:    _StoresService_CreateStore_Handler,
		},
		{
			MethodName: "EnableParticipation",
			Handler:    _StoresService_EnableParticipation_Handler,
		},
		{
			MethodName: "DisableParticipation",
			Handler:    _StoresService_DisableParticipation_Handler,
		},
		{
			MethodName: "RebrandStore",
			Handler:    _StoresService_RebrandStore_Handler,
		},
		{
			MethodName: "GetStore",
			Handler:    _StoresService_GetStore_Handler,
		},
		{
			MethodName: "GetStores",
			Handler:    _StoresService_GetStores_Handler,
		},
		{
			MethodName: "GetParticipatingStores",
			Handler:    _StoresService_GetParticipatingStores_Handler,
		},
		{
			MethodName: "AddProduct",
			Handler:    _StoresService_AddProduct_Handler,
		},
		{
			MethodName: "RebrandProduct",
			Handler:    _StoresService_RebrandProduct_Handler,
		},
		{
			MethodName: "IncreaseProductPrice",
			Handler:    _StoresService_IncreaseProductPrice_Handler,
		},
		{
			MethodName: "DecreaseProductPrice",
			Handler:    _StoresService_DecreaseProductPrice_Handler,
		},
		{
			MethodName: "RemoveProduct",
			Handler:    _StoresService_RemoveProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _StoresService_GetProduct_Handler,
		},
		{
			MethodName: "GetCatalog",
			Handler:    _StoresService_GetCatalog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storespb/api.proto",
}
