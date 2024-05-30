// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: rpcDatabase.proto

package proto

import (
	context "context"
	proto "github.com/je4/genericproto/v2/pkg/generic/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Database_Ping_FullMethodName            = "/mediaserverproto.Database/Ping"
	Database_GetItem_FullMethodName         = "/mediaserverproto.Database/GetItem"
	Database_GetItemMetadata_FullMethodName = "/mediaserverproto.Database/GetItemMetadata"
	Database_GetStorage_FullMethodName      = "/mediaserverproto.Database/GetStorage"
	Database_GetCache_FullMethodName        = "/mediaserverproto.Database/GetCache"
	Database_GetCaches_FullMethodName       = "/mediaserverproto.Database/GetCaches"
	Database_DeleteCache_FullMethodName     = "/mediaserverproto.Database/DeleteCache"
	Database_GetCollection_FullMethodName   = "/mediaserverproto.Database/GetCollection"
	Database_GetCollections_FullMethodName  = "/mediaserverproto.Database/GetCollections"
	Database_CreateItem_FullMethodName      = "/mediaserverproto.Database/CreateItem"
	Database_DeleteItem_FullMethodName      = "/mediaserverproto.Database/DeleteItem"
	Database_GetIngestItem_FullMethodName   = "/mediaserverproto.Database/GetIngestItem"
	Database_SetIngestItem_FullMethodName   = "/mediaserverproto.Database/SetIngestItem"
	Database_ExistsItem_FullMethodName      = "/mediaserverproto.Database/ExistsItem"
	Database_InsertCache_FullMethodName     = "/mediaserverproto.Database/InsertCache"
)

// DatabaseClient is the client API for Database service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*proto.DefaultResponse, error)
	GetItem(ctx context.Context, in *ItemIdentifier, opts ...grpc.CallOption) (*Item, error)
	GetItemMetadata(ctx context.Context, in *ItemIdentifier, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	GetStorage(ctx context.Context, in *StorageIdentifier, opts ...grpc.CallOption) (*Storage, error)
	GetCache(ctx context.Context, in *CacheRequest, opts ...grpc.CallOption) (*Cache, error)
	GetCaches(ctx context.Context, in *CachesRequest, opts ...grpc.CallOption) (*CachesResult, error)
	DeleteCache(ctx context.Context, in *CacheRequest, opts ...grpc.CallOption) (*proto.DefaultResponse, error)
	GetCollection(ctx context.Context, in *CollectionIdentifier, opts ...grpc.CallOption) (*Collection, error)
	GetCollections(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Database_GetCollectionsClient, error)
	CreateItem(ctx context.Context, in *NewItem, opts ...grpc.CallOption) (*proto.DefaultResponse, error)
	DeleteItem(ctx context.Context, in *ItemIdentifier, opts ...grpc.CallOption) (*proto.DefaultResponse, error)
	GetIngestItem(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IngestItem, error)
	SetIngestItem(ctx context.Context, in *IngestMetadata, opts ...grpc.CallOption) (*proto.DefaultResponse, error)
	ExistsItem(ctx context.Context, in *ItemIdentifier, opts ...grpc.CallOption) (*proto.DefaultResponse, error)
	InsertCache(ctx context.Context, in *Cache, opts ...grpc.CallOption) (*proto.DefaultResponse, error)
}

type databaseClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseClient(cc grpc.ClientConnInterface) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*proto.DefaultResponse, error) {
	out := new(proto.DefaultResponse)
	err := c.cc.Invoke(ctx, Database_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) GetItem(ctx context.Context, in *ItemIdentifier, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, Database_GetItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) GetItemMetadata(ctx context.Context, in *ItemIdentifier, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, Database_GetItemMetadata_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) GetStorage(ctx context.Context, in *StorageIdentifier, opts ...grpc.CallOption) (*Storage, error) {
	out := new(Storage)
	err := c.cc.Invoke(ctx, Database_GetStorage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) GetCache(ctx context.Context, in *CacheRequest, opts ...grpc.CallOption) (*Cache, error) {
	out := new(Cache)
	err := c.cc.Invoke(ctx, Database_GetCache_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) GetCaches(ctx context.Context, in *CachesRequest, opts ...grpc.CallOption) (*CachesResult, error) {
	out := new(CachesResult)
	err := c.cc.Invoke(ctx, Database_GetCaches_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) DeleteCache(ctx context.Context, in *CacheRequest, opts ...grpc.CallOption) (*proto.DefaultResponse, error) {
	out := new(proto.DefaultResponse)
	err := c.cc.Invoke(ctx, Database_DeleteCache_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) GetCollection(ctx context.Context, in *CollectionIdentifier, opts ...grpc.CallOption) (*Collection, error) {
	out := new(Collection)
	err := c.cc.Invoke(ctx, Database_GetCollection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) GetCollections(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Database_GetCollectionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Database_ServiceDesc.Streams[0], Database_GetCollections_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &databaseGetCollectionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Database_GetCollectionsClient interface {
	Recv() (*Collection, error)
	grpc.ClientStream
}

type databaseGetCollectionsClient struct {
	grpc.ClientStream
}

func (x *databaseGetCollectionsClient) Recv() (*Collection, error) {
	m := new(Collection)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *databaseClient) CreateItem(ctx context.Context, in *NewItem, opts ...grpc.CallOption) (*proto.DefaultResponse, error) {
	out := new(proto.DefaultResponse)
	err := c.cc.Invoke(ctx, Database_CreateItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) DeleteItem(ctx context.Context, in *ItemIdentifier, opts ...grpc.CallOption) (*proto.DefaultResponse, error) {
	out := new(proto.DefaultResponse)
	err := c.cc.Invoke(ctx, Database_DeleteItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) GetIngestItem(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IngestItem, error) {
	out := new(IngestItem)
	err := c.cc.Invoke(ctx, Database_GetIngestItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) SetIngestItem(ctx context.Context, in *IngestMetadata, opts ...grpc.CallOption) (*proto.DefaultResponse, error) {
	out := new(proto.DefaultResponse)
	err := c.cc.Invoke(ctx, Database_SetIngestItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) ExistsItem(ctx context.Context, in *ItemIdentifier, opts ...grpc.CallOption) (*proto.DefaultResponse, error) {
	out := new(proto.DefaultResponse)
	err := c.cc.Invoke(ctx, Database_ExistsItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) InsertCache(ctx context.Context, in *Cache, opts ...grpc.CallOption) (*proto.DefaultResponse, error) {
	out := new(proto.DefaultResponse)
	err := c.cc.Invoke(ctx, Database_InsertCache_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServer is the server API for Database service.
// All implementations must embed UnimplementedDatabaseServer
// for forward compatibility
type DatabaseServer interface {
	Ping(context.Context, *emptypb.Empty) (*proto.DefaultResponse, error)
	GetItem(context.Context, *ItemIdentifier) (*Item, error)
	GetItemMetadata(context.Context, *ItemIdentifier) (*wrapperspb.StringValue, error)
	GetStorage(context.Context, *StorageIdentifier) (*Storage, error)
	GetCache(context.Context, *CacheRequest) (*Cache, error)
	GetCaches(context.Context, *CachesRequest) (*CachesResult, error)
	DeleteCache(context.Context, *CacheRequest) (*proto.DefaultResponse, error)
	GetCollection(context.Context, *CollectionIdentifier) (*Collection, error)
	GetCollections(*emptypb.Empty, Database_GetCollectionsServer) error
	CreateItem(context.Context, *NewItem) (*proto.DefaultResponse, error)
	DeleteItem(context.Context, *ItemIdentifier) (*proto.DefaultResponse, error)
	GetIngestItem(context.Context, *emptypb.Empty) (*IngestItem, error)
	SetIngestItem(context.Context, *IngestMetadata) (*proto.DefaultResponse, error)
	ExistsItem(context.Context, *ItemIdentifier) (*proto.DefaultResponse, error)
	InsertCache(context.Context, *Cache) (*proto.DefaultResponse, error)
	mustEmbedUnimplementedDatabaseServer()
}

// UnimplementedDatabaseServer must be embedded to have forward compatible implementations.
type UnimplementedDatabaseServer struct {
}

func (UnimplementedDatabaseServer) Ping(context.Context, *emptypb.Empty) (*proto.DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedDatabaseServer) GetItem(context.Context, *ItemIdentifier) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedDatabaseServer) GetItemMetadata(context.Context, *ItemIdentifier) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemMetadata not implemented")
}
func (UnimplementedDatabaseServer) GetStorage(context.Context, *StorageIdentifier) (*Storage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorage not implemented")
}
func (UnimplementedDatabaseServer) GetCache(context.Context, *CacheRequest) (*Cache, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCache not implemented")
}
func (UnimplementedDatabaseServer) GetCaches(context.Context, *CachesRequest) (*CachesResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaches not implemented")
}
func (UnimplementedDatabaseServer) DeleteCache(context.Context, *CacheRequest) (*proto.DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCache not implemented")
}
func (UnimplementedDatabaseServer) GetCollection(context.Context, *CollectionIdentifier) (*Collection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollection not implemented")
}
func (UnimplementedDatabaseServer) GetCollections(*emptypb.Empty, Database_GetCollectionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCollections not implemented")
}
func (UnimplementedDatabaseServer) CreateItem(context.Context, *NewItem) (*proto.DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedDatabaseServer) DeleteItem(context.Context, *ItemIdentifier) (*proto.DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}
func (UnimplementedDatabaseServer) GetIngestItem(context.Context, *emptypb.Empty) (*IngestItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIngestItem not implemented")
}
func (UnimplementedDatabaseServer) SetIngestItem(context.Context, *IngestMetadata) (*proto.DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIngestItem not implemented")
}
func (UnimplementedDatabaseServer) ExistsItem(context.Context, *ItemIdentifier) (*proto.DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistsItem not implemented")
}
func (UnimplementedDatabaseServer) InsertCache(context.Context, *Cache) (*proto.DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertCache not implemented")
}
func (UnimplementedDatabaseServer) mustEmbedUnimplementedDatabaseServer() {}

// UnsafeDatabaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseServer will
// result in compilation errors.
type UnsafeDatabaseServer interface {
	mustEmbedUnimplementedDatabaseServer()
}

func RegisterDatabaseServer(s grpc.ServiceRegistrar, srv DatabaseServer) {
	s.RegisterService(&Database_ServiceDesc, srv)
}

func _Database_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_GetItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).GetItem(ctx, req.(*ItemIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_GetItemMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).GetItemMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_GetItemMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).GetItemMetadata(ctx, req.(*ItemIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_GetStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).GetStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_GetStorage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).GetStorage(ctx, req.(*StorageIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_GetCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).GetCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_GetCache_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).GetCache(ctx, req.(*CacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_GetCaches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CachesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).GetCaches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_GetCaches_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).GetCaches(ctx, req.(*CachesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_DeleteCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).DeleteCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_DeleteCache_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).DeleteCache(ctx, req.(*CacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_GetCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).GetCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_GetCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).GetCollection(ctx, req.(*CollectionIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_GetCollections_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DatabaseServer).GetCollections(m, &databaseGetCollectionsServer{stream})
}

type Database_GetCollectionsServer interface {
	Send(*Collection) error
	grpc.ServerStream
}

type databaseGetCollectionsServer struct {
	grpc.ServerStream
}

func (x *databaseGetCollectionsServer) Send(m *Collection) error {
	return x.ServerStream.SendMsg(m)
}

func _Database_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_CreateItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).CreateItem(ctx, req.(*NewItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_DeleteItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).DeleteItem(ctx, req.(*ItemIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_GetIngestItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).GetIngestItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_GetIngestItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).GetIngestItem(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_SetIngestItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngestMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).SetIngestItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_SetIngestItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).SetIngestItem(ctx, req.(*IngestMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_ExistsItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).ExistsItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_ExistsItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).ExistsItem(ctx, req.(*ItemIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_InsertCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cache)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).InsertCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_InsertCache_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).InsertCache(ctx, req.(*Cache))
	}
	return interceptor(ctx, in, info, handler)
}

// Database_ServiceDesc is the grpc.ServiceDesc for Database service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Database_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mediaserverproto.Database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Database_Ping_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _Database_GetItem_Handler,
		},
		{
			MethodName: "GetItemMetadata",
			Handler:    _Database_GetItemMetadata_Handler,
		},
		{
			MethodName: "GetStorage",
			Handler:    _Database_GetStorage_Handler,
		},
		{
			MethodName: "GetCache",
			Handler:    _Database_GetCache_Handler,
		},
		{
			MethodName: "GetCaches",
			Handler:    _Database_GetCaches_Handler,
		},
		{
			MethodName: "DeleteCache",
			Handler:    _Database_DeleteCache_Handler,
		},
		{
			MethodName: "GetCollection",
			Handler:    _Database_GetCollection_Handler,
		},
		{
			MethodName: "CreateItem",
			Handler:    _Database_CreateItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _Database_DeleteItem_Handler,
		},
		{
			MethodName: "GetIngestItem",
			Handler:    _Database_GetIngestItem_Handler,
		},
		{
			MethodName: "SetIngestItem",
			Handler:    _Database_SetIngestItem_Handler,
		},
		{
			MethodName: "ExistsItem",
			Handler:    _Database_ExistsItem_Handler,
		},
		{
			MethodName: "InsertCache",
			Handler:    _Database_InsertCache_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCollections",
			Handler:       _Database_GetCollections_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpcDatabase.proto",
}
