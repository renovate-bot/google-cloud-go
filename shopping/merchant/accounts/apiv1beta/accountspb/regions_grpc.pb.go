// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.7
// source: google/shopping/merchant/accounts/v1beta/regions.proto

package accountspb

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
	RegionsService_GetRegion_FullMethodName    = "/google.shopping.merchant.accounts.v1beta.RegionsService/GetRegion"
	RegionsService_CreateRegion_FullMethodName = "/google.shopping.merchant.accounts.v1beta.RegionsService/CreateRegion"
	RegionsService_UpdateRegion_FullMethodName = "/google.shopping.merchant.accounts.v1beta.RegionsService/UpdateRegion"
	RegionsService_DeleteRegion_FullMethodName = "/google.shopping.merchant.accounts.v1beta.RegionsService/DeleteRegion"
	RegionsService_ListRegions_FullMethodName  = "/google.shopping.merchant.accounts.v1beta.RegionsService/ListRegions"
)

// RegionsServiceClient is the client API for RegionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegionsServiceClient interface {
	// Retrieves a region defined in your Merchant Center account.
	GetRegion(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*Region, error)
	// Creates a region definition in your Merchant Center account. Executing this
	// method requires admin access.
	CreateRegion(ctx context.Context, in *CreateRegionRequest, opts ...grpc.CallOption) (*Region, error)
	// Updates a region definition in your Merchant Center account. Executing this
	// method requires admin access.
	UpdateRegion(ctx context.Context, in *UpdateRegionRequest, opts ...grpc.CallOption) (*Region, error)
	// Deletes a region definition from your Merchant Center account. Executing
	// this method requires admin access.
	DeleteRegion(ctx context.Context, in *DeleteRegionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Lists the regions in your Merchant Center account.
	ListRegions(ctx context.Context, in *ListRegionsRequest, opts ...grpc.CallOption) (*ListRegionsResponse, error)
}

type regionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegionsServiceClient(cc grpc.ClientConnInterface) RegionsServiceClient {
	return &regionsServiceClient{cc}
}

func (c *regionsServiceClient) GetRegion(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*Region, error) {
	out := new(Region)
	err := c.cc.Invoke(ctx, RegionsService_GetRegion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionsServiceClient) CreateRegion(ctx context.Context, in *CreateRegionRequest, opts ...grpc.CallOption) (*Region, error) {
	out := new(Region)
	err := c.cc.Invoke(ctx, RegionsService_CreateRegion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionsServiceClient) UpdateRegion(ctx context.Context, in *UpdateRegionRequest, opts ...grpc.CallOption) (*Region, error) {
	out := new(Region)
	err := c.cc.Invoke(ctx, RegionsService_UpdateRegion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionsServiceClient) DeleteRegion(ctx context.Context, in *DeleteRegionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RegionsService_DeleteRegion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionsServiceClient) ListRegions(ctx context.Context, in *ListRegionsRequest, opts ...grpc.CallOption) (*ListRegionsResponse, error) {
	out := new(ListRegionsResponse)
	err := c.cc.Invoke(ctx, RegionsService_ListRegions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegionsServiceServer is the server API for RegionsService service.
// All implementations should embed UnimplementedRegionsServiceServer
// for forward compatibility
type RegionsServiceServer interface {
	// Retrieves a region defined in your Merchant Center account.
	GetRegion(context.Context, *GetRegionRequest) (*Region, error)
	// Creates a region definition in your Merchant Center account. Executing this
	// method requires admin access.
	CreateRegion(context.Context, *CreateRegionRequest) (*Region, error)
	// Updates a region definition in your Merchant Center account. Executing this
	// method requires admin access.
	UpdateRegion(context.Context, *UpdateRegionRequest) (*Region, error)
	// Deletes a region definition from your Merchant Center account. Executing
	// this method requires admin access.
	DeleteRegion(context.Context, *DeleteRegionRequest) (*emptypb.Empty, error)
	// Lists the regions in your Merchant Center account.
	ListRegions(context.Context, *ListRegionsRequest) (*ListRegionsResponse, error)
}

// UnimplementedRegionsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRegionsServiceServer struct {
}

func (UnimplementedRegionsServiceServer) GetRegion(context.Context, *GetRegionRequest) (*Region, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegion not implemented")
}
func (UnimplementedRegionsServiceServer) CreateRegion(context.Context, *CreateRegionRequest) (*Region, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRegion not implemented")
}
func (UnimplementedRegionsServiceServer) UpdateRegion(context.Context, *UpdateRegionRequest) (*Region, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRegion not implemented")
}
func (UnimplementedRegionsServiceServer) DeleteRegion(context.Context, *DeleteRegionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRegion not implemented")
}
func (UnimplementedRegionsServiceServer) ListRegions(context.Context, *ListRegionsRequest) (*ListRegionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRegions not implemented")
}

// UnsafeRegionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegionsServiceServer will
// result in compilation errors.
type UnsafeRegionsServiceServer interface {
	mustEmbedUnimplementedRegionsServiceServer()
}

func RegisterRegionsServiceServer(s grpc.ServiceRegistrar, srv RegionsServiceServer) {
	s.RegisterService(&RegionsService_ServiceDesc, srv)
}

func _RegionsService_GetRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionsServiceServer).GetRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionsService_GetRegion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionsServiceServer).GetRegion(ctx, req.(*GetRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionsService_CreateRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionsServiceServer).CreateRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionsService_CreateRegion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionsServiceServer).CreateRegion(ctx, req.(*CreateRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionsService_UpdateRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionsServiceServer).UpdateRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionsService_UpdateRegion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionsServiceServer).UpdateRegion(ctx, req.(*UpdateRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionsService_DeleteRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionsServiceServer).DeleteRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionsService_DeleteRegion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionsServiceServer).DeleteRegion(ctx, req.(*DeleteRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionsService_ListRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRegionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionsServiceServer).ListRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionsService_ListRegions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionsServiceServer).ListRegions(ctx, req.(*ListRegionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegionsService_ServiceDesc is the grpc.ServiceDesc for RegionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.shopping.merchant.accounts.v1beta.RegionsService",
	HandlerType: (*RegionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRegion",
			Handler:    _RegionsService_GetRegion_Handler,
		},
		{
			MethodName: "CreateRegion",
			Handler:    _RegionsService_CreateRegion_Handler,
		},
		{
			MethodName: "UpdateRegion",
			Handler:    _RegionsService_UpdateRegion_Handler,
		},
		{
			MethodName: "DeleteRegion",
			Handler:    _RegionsService_DeleteRegion_Handler,
		},
		{
			MethodName: "ListRegions",
			Handler:    _RegionsService_ListRegions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/shopping/merchant/accounts/v1beta/regions.proto",
}
