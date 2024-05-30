// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: action.proto

package proto

import (
	proto "github.com/je4/genericproto/v2/pkg/generic/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ActionParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item    *Item             `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Action  string            `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Params  map[string]string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Storage *Storage          `protobuf:"bytes,4,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (x *ActionParam) Reset() {
	*x = ActionParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionParam) ProtoMessage() {}

func (x *ActionParam) ProtoReflect() protoreflect.Message {
	mi := &file_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionParam.ProtoReflect.Descriptor instead.
func (*ActionParam) Descriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{0}
}

func (x *ActionParam) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *ActionParam) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *ActionParam) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *ActionParam) GetStorage() *Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

type ParamsParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Action string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *ParamsParam) Reset() {
	*x = ParamsParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamsParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamsParam) ProtoMessage() {}

func (x *ParamsParam) ProtoReflect() protoreflect.Message {
	mi := &file_action_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamsParam.ProtoReflect.Descriptor instead.
func (*ParamsParam) Descriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{1}
}

func (x *ParamsParam) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ParamsParam) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type ActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ActionResponse:
	//
	//	*ActionResponse_Response
	//	*ActionResponse_Cache
	ActionResponse isActionResponse_ActionResponse `protobuf_oneof:"actionResponse"`
}

func (x *ActionResponse) Reset() {
	*x = ActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionResponse) ProtoMessage() {}

func (x *ActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_action_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionResponse.ProtoReflect.Descriptor instead.
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{2}
}

func (m *ActionResponse) GetActionResponse() isActionResponse_ActionResponse {
	if m != nil {
		return m.ActionResponse
	}
	return nil
}

func (x *ActionResponse) GetResponse() *proto.DefaultResponse {
	if x, ok := x.GetActionResponse().(*ActionResponse_Response); ok {
		return x.Response
	}
	return nil
}

func (x *ActionResponse) GetCache() *Cache {
	if x, ok := x.GetActionResponse().(*ActionResponse_Cache); ok {
		return x.Cache
	}
	return nil
}

type isActionResponse_ActionResponse interface {
	isActionResponse_ActionResponse()
}

type ActionResponse_Response struct {
	Response *proto.DefaultResponse `protobuf:"bytes,2,opt,name=response,proto3,oneof"`
}

type ActionResponse_Cache struct {
	Cache *Cache `protobuf:"bytes,3,opt,name=cache,proto3,oneof"`
}

func (*ActionResponse_Response) isActionResponse_ActionResponse() {}

func (*ActionResponse_Cache) isActionResponse_ActionResponse() {}

type ActionDispatcherParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Action      []string `protobuf:"bytes,2,rep,name=action,proto3" json:"action,omitempty"`
	Host        *string  `protobuf:"bytes,3,opt,name=host,proto3,oneof" json:"host,omitempty"`
	Port        uint32   `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Concurrency uint32   `protobuf:"varint,5,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	Param       []string `protobuf:"bytes,6,rep,name=param,proto3" json:"param,omitempty"`
}

func (x *ActionDispatcherParam) Reset() {
	*x = ActionDispatcherParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionDispatcherParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionDispatcherParam) ProtoMessage() {}

func (x *ActionDispatcherParam) ProtoReflect() protoreflect.Message {
	mi := &file_action_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionDispatcherParam.ProtoReflect.Descriptor instead.
func (*ActionDispatcherParam) Descriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{3}
}

func (x *ActionDispatcherParam) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ActionDispatcherParam) GetAction() []string {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *ActionDispatcherParam) GetHost() string {
	if x != nil && x.Host != nil {
		return *x.Host
	}
	return ""
}

func (x *ActionDispatcherParam) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ActionDispatcherParam) GetConcurrency() uint32 {
	if x != nil {
		return x.Concurrency
	}
	return 0
}

func (x *ActionDispatcherParam) GetParam() []string {
	if x != nil {
		return x.Param
	}
	return nil
}

var File_action_proto protoreflect.FileDescriptor

var file_action_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x84, 0x02, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12,
	0x2a, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x48, 0x00, 0x52, 0x05, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x15, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x42, 0x8a, 0x01, 0x0a, 0x18, 0x63, 0x68, 0x2e,
	0x75, 0x6e, 0x69, 0x62, 0x61, 0x73, 0x2e, 0x75, 0x62, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x14, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x65, 0x34, 0x2f, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x55, 0x42, 0x42, 0xaa, 0x02, 0x15,
	0x55, 0x6e, 0x69, 0x62, 0x61, 0x73, 0x2e, 0x55, 0x42, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_action_proto_rawDescOnce sync.Once
	file_action_proto_rawDescData = file_action_proto_rawDesc
)

func file_action_proto_rawDescGZIP() []byte {
	file_action_proto_rawDescOnce.Do(func() {
		file_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_action_proto_rawDescData)
	})
	return file_action_proto_rawDescData
}

var file_action_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_action_proto_goTypes = []interface{}{
	(*ActionParam)(nil),           // 0: mediaserverproto.ActionParam
	(*ParamsParam)(nil),           // 1: mediaserverproto.ParamsParam
	(*ActionResponse)(nil),        // 2: mediaserverproto.ActionResponse
	(*ActionDispatcherParam)(nil), // 3: mediaserverproto.ActionDispatcherParam
	nil,                           // 4: mediaserverproto.ActionParam.ParamsEntry
	(*Item)(nil),                  // 5: mediaserverproto.Item
	(*Storage)(nil),               // 6: mediaserverproto.Storage
	(*proto.DefaultResponse)(nil), // 7: genericproto.DefaultResponse
	(*Cache)(nil),                 // 8: mediaserverproto.Cache
}
var file_action_proto_depIdxs = []int32{
	5, // 0: mediaserverproto.ActionParam.item:type_name -> mediaserverproto.Item
	4, // 1: mediaserverproto.ActionParam.params:type_name -> mediaserverproto.ActionParam.ParamsEntry
	6, // 2: mediaserverproto.ActionParam.storage:type_name -> mediaserverproto.Storage
	7, // 3: mediaserverproto.ActionResponse.response:type_name -> genericproto.DefaultResponse
	8, // 4: mediaserverproto.ActionResponse.cache:type_name -> mediaserverproto.Cache
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_action_proto_init() }
func file_action_proto_init() {
	if File_action_proto != nil {
		return
	}
	file_item_proto_init()
	file_cache_proto_init()
	file_storage_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionParam); i {
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
		file_action_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamsParam); i {
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
		file_action_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionResponse); i {
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
		file_action_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionDispatcherParam); i {
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
	file_action_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ActionResponse_Response)(nil),
		(*ActionResponse_Cache)(nil),
	}
	file_action_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_action_proto_goTypes,
		DependencyIndexes: file_action_proto_depIdxs,
		MessageInfos:      file_action_proto_msgTypes,
	}.Build()
	File_action_proto = out.File
	file_action_proto_rawDesc = nil
	file_action_proto_goTypes = nil
	file_action_proto_depIdxs = nil
}
