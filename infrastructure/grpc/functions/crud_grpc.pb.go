// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: crud.proto

package functions

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CRUDService_Create_FullMethodName      = "/crud.CRUDService/Create"
	CRUDService_Read_FullMethodName        = "/crud.CRUDService/Read"
	CRUDService_BatchCreate_FullMethodName = "/crud.CRUDService/BatchCreate"
	CRUDService_Update_FullMethodName      = "/crud.CRUDService/Update"
	CRUDService_Delete_FullMethodName      = "/crud.CRUDService/Delete"
)

// CRUDServiceClient is the client API for CRUDService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CRUDServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*IdResponse, error)
	Read(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	BatchCreate(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type cRUDServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCRUDServiceClient(cc grpc.ClientConnInterface) CRUDServiceClient {
	return &cRUDServiceClient{cc}
}

func (c *cRUDServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, CRUDService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDServiceClient) Read(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, CRUDService_Read_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDServiceClient) BatchCreate(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CRUDService_BatchCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, CRUDService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDServiceClient) Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CRUDService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CRUDServiceServer is the server API for CRUDService service.
// All implementations must embed UnimplementedCRUDServiceServer
// for forward compatibility
type CRUDServiceServer interface {
	Create(context.Context, *CreateRequest) (*IdResponse, error)
	Read(context.Context, *IdRequest) (*TaskResponse, error)
	BatchCreate(context.Context, *BatchRequest) (*emptypb.Empty, error)
	Update(context.Context, *UpdateRequest) (*TaskResponse, error)
	Delete(context.Context, *IdRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCRUDServiceServer()
}

// UnimplementedCRUDServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCRUDServiceServer struct {
}

func (UnimplementedCRUDServiceServer) Create(context.Context, *CreateRequest) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCRUDServiceServer) Read(context.Context, *IdRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedCRUDServiceServer) BatchCreate(context.Context, *BatchRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreate not implemented")
}
func (UnimplementedCRUDServiceServer) Update(context.Context, *UpdateRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCRUDServiceServer) Delete(context.Context, *IdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCRUDServiceServer) mustEmbedUnimplementedCRUDServiceServer() {}

// UnsafeCRUDServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CRUDServiceServer will
// result in compilation errors.
type UnsafeCRUDServiceServer interface {
	mustEmbedUnimplementedCRUDServiceServer()
}

func RegisterCRUDServiceServer(s grpc.ServiceRegistrar, srv CRUDServiceServer) {
	s.RegisterService(&CRUDService_ServiceDesc, srv)
}

func _CRUDService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRUDService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUDService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRUDService_Read_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServiceServer).Read(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUDService_BatchCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServiceServer).BatchCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRUDService_BatchCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServiceServer).BatchCreate(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUDService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRUDService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUDService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRUDService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServiceServer).Delete(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CRUDService_ServiceDesc is the grpc.ServiceDesc for CRUDService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CRUDService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crud.CRUDService",
	HandlerType: (*CRUDServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CRUDService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _CRUDService_Read_Handler,
		},
		{
			MethodName: "BatchCreate",
			Handler:    _CRUDService_BatchCreate_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CRUDService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CRUDService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crud.proto",
}
