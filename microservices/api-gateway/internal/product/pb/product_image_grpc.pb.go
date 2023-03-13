// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: internal/product/pb/product_image.proto

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

// ProductImageServiceClient is the client API for ProductImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductImageServiceClient interface {
	GetAllByProductId(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	UploadImageToProduct(ctx context.Context, in *UploadImageRequest, opts ...grpc.CallOption) (*UploadImageResponse, error)
	DeleteByNameAndProductId(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*DeleteImageResponse, error)
}

type productImageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductImageServiceClient(cc grpc.ClientConnInterface) ProductImageServiceClient {
	return &productImageServiceClient{cc}
}

func (c *productImageServiceClient) GetAllByProductId(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/product.ProductImageService/GetAllByProductId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productImageServiceClient) UploadImageToProduct(ctx context.Context, in *UploadImageRequest, opts ...grpc.CallOption) (*UploadImageResponse, error) {
	out := new(UploadImageResponse)
	err := c.cc.Invoke(ctx, "/product.ProductImageService/UploadImageToProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productImageServiceClient) DeleteByNameAndProductId(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*DeleteImageResponse, error) {
	out := new(DeleteImageResponse)
	err := c.cc.Invoke(ctx, "/product.ProductImageService/DeleteByNameAndProductId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductImageServiceServer is the server API for ProductImageService service.
// All implementations should embed UnimplementedProductImageServiceServer
// for forward compatibility
type ProductImageServiceServer interface {
	GetAllByProductId(context.Context, *GetAllRequest) (*GetAllResponse, error)
	UploadImageToProduct(context.Context, *UploadImageRequest) (*UploadImageResponse, error)
	DeleteByNameAndProductId(context.Context, *DeleteImageRequest) (*DeleteImageResponse, error)
}

// UnimplementedProductImageServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProductImageServiceServer struct {
}

func (UnimplementedProductImageServiceServer) GetAllByProductId(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByProductId not implemented")
}
func (UnimplementedProductImageServiceServer) UploadImageToProduct(context.Context, *UploadImageRequest) (*UploadImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadImageToProduct not implemented")
}
func (UnimplementedProductImageServiceServer) DeleteByNameAndProductId(context.Context, *DeleteImageRequest) (*DeleteImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByNameAndProductId not implemented")
}

// UnsafeProductImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductImageServiceServer will
// result in compilation errors.
type UnsafeProductImageServiceServer interface {
	mustEmbedUnimplementedProductImageServiceServer()
}

func RegisterProductImageServiceServer(s grpc.ServiceRegistrar, srv ProductImageServiceServer) {
	s.RegisterService(&ProductImageService_ServiceDesc, srv)
}

func _ProductImageService_GetAllByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductImageServiceServer).GetAllByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductImageService/GetAllByProductId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductImageServiceServer).GetAllByProductId(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductImageService_UploadImageToProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductImageServiceServer).UploadImageToProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductImageService/UploadImageToProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductImageServiceServer).UploadImageToProduct(ctx, req.(*UploadImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductImageService_DeleteByNameAndProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductImageServiceServer).DeleteByNameAndProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductImageService/DeleteByNameAndProductId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductImageServiceServer).DeleteByNameAndProductId(ctx, req.(*DeleteImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductImageService_ServiceDesc is the grpc.ServiceDesc for ProductImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductImageService",
	HandlerType: (*ProductImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllByProductId",
			Handler:    _ProductImageService_GetAllByProductId_Handler,
		},
		{
			MethodName: "UploadImageToProduct",
			Handler:    _ProductImageService_UploadImageToProduct_Handler,
		},
		{
			MethodName: "DeleteByNameAndProductId",
			Handler:    _ProductImageService_DeleteByNameAndProductId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/product/pb/product_image.proto",
}
