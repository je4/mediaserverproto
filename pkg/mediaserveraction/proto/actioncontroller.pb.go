// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: actioncontroller.proto

package proto

import (
	proto "github.com/je4/genericproto/v2/pkg/generic/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_actioncontroller_proto protoreflect.FileDescriptor

var file_actioncontroller_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xed, 0x01, 0x0a, 0x10, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x18, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x06, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0xb3, 0x01, 0x0a, 0x29, 0x63, 0x68, 0x2e, 0x75, 0x6e, 0x69, 0x62,
	0x61, 0x73, 0x2e, 0x75, 0x62, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x42, 0x15, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x65, 0x34, 0x2f, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x55, 0x42,
	0x42, 0xaa, 0x02, 0x26, 0x55, 0x6e, 0x69, 0x62, 0x61, 0x73, 0x2e, 0x55, 0x42, 0x2e, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_actioncontroller_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil),         // 0: google.protobuf.Empty
	(*ActionParam)(nil),           // 1: mediaserveractionproto.ActionParam
	(*proto.DefaultResponse)(nil), // 2: genericproto.DefaultResponse
	(*proto.StringList)(nil),      // 3: genericproto.StringList
	(*ActionResponse)(nil),        // 4: mediaserveractionproto.ActionResponse
}
var file_actioncontroller_proto_depIdxs = []int32{
	0, // 0: mediaserveractionproto.ActionController.Ping:input_type -> google.protobuf.Empty
	0, // 1: mediaserveractionproto.ActionController.GetParams:input_type -> google.protobuf.Empty
	1, // 2: mediaserveractionproto.ActionController.Action:input_type -> mediaserveractionproto.ActionParam
	2, // 3: mediaserveractionproto.ActionController.Ping:output_type -> genericproto.DefaultResponse
	3, // 4: mediaserveractionproto.ActionController.GetParams:output_type -> genericproto.StringList
	4, // 5: mediaserveractionproto.ActionController.Action:output_type -> mediaserveractionproto.ActionResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_actioncontroller_proto_init() }
func file_actioncontroller_proto_init() {
	if File_actioncontroller_proto != nil {
		return
	}
	file_action_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_actioncontroller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_actioncontroller_proto_goTypes,
		DependencyIndexes: file_actioncontroller_proto_depIdxs,
	}.Build()
	File_actioncontroller_proto = out.File
	file_actioncontroller_proto_rawDesc = nil
	file_actioncontroller_proto_goTypes = nil
	file_actioncontroller_proto_depIdxs = nil
}
