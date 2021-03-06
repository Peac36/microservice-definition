// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

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

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	Create(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	Update(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostResponse, error)
	Delete(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
	Get(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error)
	Index(ctx context.Context, in *IndexPostRequest, opts ...grpc.CallOption) (*IndexPostResponse, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) Create(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, "/servicesDefinitions.PostService/create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) Update(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostResponse, error) {
	out := new(UpdatePostResponse)
	err := c.cc.Invoke(ctx, "/servicesDefinitions.PostService/update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) Delete(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, "/servicesDefinitions.PostService/delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) Get(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error) {
	out := new(GetPostResponse)
	err := c.cc.Invoke(ctx, "/servicesDefinitions.PostService/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) Index(ctx context.Context, in *IndexPostRequest, opts ...grpc.CallOption) (*IndexPostResponse, error) {
	out := new(IndexPostResponse)
	err := c.cc.Invoke(ctx, "/servicesDefinitions.PostService/index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	Create(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	Update(context.Context, *UpdatePostRequest) (*UpdatePostResponse, error)
	Delete(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	Get(context.Context, *GetPostRequest) (*GetPostResponse, error)
	Index(context.Context, *IndexPostRequest) (*IndexPostResponse, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) Create(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPostServiceServer) Update(context.Context, *UpdatePostRequest) (*UpdatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPostServiceServer) Delete(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPostServiceServer) Get(context.Context, *GetPostRequest) (*GetPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPostServiceServer) Index(context.Context, *IndexPostRequest) (*IndexPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicesDefinitions.PostService/create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Create(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicesDefinitions.PostService/update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Update(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicesDefinitions.PostService/delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Delete(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicesDefinitions.PostService/get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Get(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicesDefinitions.PostService/index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Index(ctx, req.(*IndexPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "servicesDefinitions.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _PostService_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _PostService_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _PostService_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _PostService_Get_Handler,
		},
		{
			MethodName: "index",
			Handler:    _PostService_Index_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "definition.proto",
}

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryServiceClient interface {
	Create(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error)
	FindMany(ctx context.Context, in *FindManyCategoryRequest, opts ...grpc.CallOption) (*FindManyCategoryResponse, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) Create(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error) {
	out := new(CreateCategoryResponse)
	err := c.cc.Invoke(ctx, "/servicesDefinitions.CategoryService/create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) FindMany(ctx context.Context, in *FindManyCategoryRequest, opts ...grpc.CallOption) (*FindManyCategoryResponse, error) {
	out := new(FindManyCategoryResponse)
	err := c.cc.Invoke(ctx, "/servicesDefinitions.CategoryService/findMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations must embed UnimplementedCategoryServiceServer
// for forward compatibility
type CategoryServiceServer interface {
	Create(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error)
	FindMany(context.Context, *FindManyCategoryRequest) (*FindManyCategoryResponse, error)
	mustEmbedUnimplementedCategoryServiceServer()
}

// UnimplementedCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryServiceServer struct {
}

func (UnimplementedCategoryServiceServer) Create(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCategoryServiceServer) FindMany(context.Context, *FindManyCategoryRequest) (*FindManyCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMany not implemented")
}
func (UnimplementedCategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s grpc.ServiceRegistrar, srv CategoryServiceServer) {
	s.RegisterService(&CategoryService_ServiceDesc, srv)
}

func _CategoryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicesDefinitions.CategoryService/create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Create(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_FindMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindManyCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).FindMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicesDefinitions.CategoryService/findMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).FindMany(ctx, req.(*FindManyCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryService_ServiceDesc is the grpc.ServiceDesc for CategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "servicesDefinitions.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _CategoryService_Create_Handler,
		},
		{
			MethodName: "findMany",
			Handler:    _CategoryService_FindMany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "definition.proto",
}
