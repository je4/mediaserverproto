// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: rpcAction.proto

package proto

import (
	proto "github.com/je4/genericproto/v2/pkg/generic/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DispatcherDefaultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response     *proto.DefaultResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	NextCallWait int64                  `protobuf:"varint,4,opt,name=nextCallWait,proto3" json:"nextCallWait,omitempty"`
}

func (x *DispatcherDefaultResponse) Reset() {
	*x = DispatcherDefaultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcAction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatcherDefaultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatcherDefaultResponse) ProtoMessage() {}

func (x *DispatcherDefaultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcAction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatcherDefaultResponse.ProtoReflect.Descriptor instead.
func (*DispatcherDefaultResponse) Descriptor() ([]byte, []int) {
	return file_rpcAction_proto_rawDescGZIP(), []int{0}
}

func (x *DispatcherDefaultResponse) GetResponse() *proto.DefaultResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *DispatcherDefaultResponse) GetNextCallWait() int64 {
	if x != nil {
		return x.NextCallWait
	}
	return 0
}

var File_rpcAction_proto protoreflect.FileDescriptor

var file_rpcAction_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x70, 0x63, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x19, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x39, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x6e, 0x65, 0x78, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x57, 0x61, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x57, 0x61, 0x69, 0x74,
	0x32, 0xd5, 0x01, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x69, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x18, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x17, 0x2e,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x22, 0x00, 0x32, 0x9a, 0x02, 0x0a, 0x10, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x3f, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67,
	0x0a, 0x0d, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12,
	0x27, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x2b, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1d, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x8a, 0x01, 0x0a, 0x18, 0x63, 0x68, 0x2e, 0x75, 0x6e, 0x69,
	0x62, 0x61, 0x73, 0x2e, 0x75, 0x62, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x42, 0x14, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x65, 0x34, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x55, 0x42, 0x42, 0xaa, 0x02, 0x15, 0x55, 0x6e, 0x69,
	0x62, 0x61, 0x73, 0x2e, 0x55, 0x42, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpcAction_proto_rawDescOnce sync.Once
	file_rpcAction_proto_rawDescData = file_rpcAction_proto_rawDesc
)

func file_rpcAction_proto_rawDescGZIP() []byte {
	file_rpcAction_proto_rawDescOnce.Do(func() {
		file_rpcAction_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpcAction_proto_rawDescData)
	})
	return file_rpcAction_proto_rawDescData
}

var file_rpcAction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rpcAction_proto_goTypes = []interface{}{
	(*DispatcherDefaultResponse)(nil), // 0: mediaserverproto.DispatcherDefaultResponse
	(*proto.DefaultResponse)(nil),     // 1: genericproto.DefaultResponse
	(*emptypb.Empty)(nil),             // 2: google.protobuf.Empty
	(*ParamsParam)(nil),               // 3: mediaserverproto.ParamsParam
	(*ActionParam)(nil),               // 4: mediaserverproto.ActionParam
	(*ActionDispatcherParam)(nil),     // 5: mediaserverproto.ActionDispatcherParam
	(*proto.StringList)(nil),          // 6: genericproto.StringList
	(*Cache)(nil),                     // 7: mediaserverproto.Cache
}
var file_rpcAction_proto_depIdxs = []int32{
	1, // 0: mediaserverproto.DispatcherDefaultResponse.response:type_name -> genericproto.DefaultResponse
	2, // 1: mediaserverproto.Action.Ping:input_type -> google.protobuf.Empty
	3, // 2: mediaserverproto.Action.GetParams:input_type -> mediaserverproto.ParamsParam
	4, // 3: mediaserverproto.Action.Action:input_type -> mediaserverproto.ActionParam
	2, // 4: mediaserverproto.ActionDispatcher.Ping:input_type -> google.protobuf.Empty
	5, // 5: mediaserverproto.ActionDispatcher.AddController:input_type -> mediaserverproto.ActionDispatcherParam
	5, // 6: mediaserverproto.ActionDispatcher.RemoveController:input_type -> mediaserverproto.ActionDispatcherParam
	1, // 7: mediaserverproto.Action.Ping:output_type -> genericproto.DefaultResponse
	6, // 8: mediaserverproto.Action.GetParams:output_type -> genericproto.StringList
	7, // 9: mediaserverproto.Action.Action:output_type -> mediaserverproto.Cache
	1, // 10: mediaserverproto.ActionDispatcher.Ping:output_type -> genericproto.DefaultResponse
	0, // 11: mediaserverproto.ActionDispatcher.AddController:output_type -> mediaserverproto.DispatcherDefaultResponse
	1, // 12: mediaserverproto.ActionDispatcher.RemoveController:output_type -> genericproto.DefaultResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpcAction_proto_init() }
func file_rpcAction_proto_init() {
	if File_rpcAction_proto != nil {
		return
	}
	file_action_proto_init()
	file_cache_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpcAction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatcherDefaultResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpcAction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_rpcAction_proto_goTypes,
		DependencyIndexes: file_rpcAction_proto_depIdxs,
		MessageInfos:      file_rpcAction_proto_msgTypes,
	}.Build()
	File_rpcAction_proto = out.File
	file_rpcAction_proto_rawDesc = nil
	file_rpcAction_proto_goTypes = nil
	file_rpcAction_proto_depIdxs = nil
}
