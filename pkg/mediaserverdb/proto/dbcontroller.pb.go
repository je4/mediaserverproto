// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: dbcontroller.proto

package proto

import (
	proto "github.com/je4/mediaserverproto/v2/pkg/mediaservergeneric/proto"
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

type IngestMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item          *ItemIdentifier `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Status        string          `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	ItemMetadata  *ItemMetadata   `protobuf:"bytes,3,opt,name=itemMetadata,proto3" json:"itemMetadata,omitempty"`
	CacheMetadata *CacheMetadata  `protobuf:"bytes,4,opt,name=cacheMetadata,proto3" json:"cacheMetadata,omitempty"`
}

func (x *IngestMetadata) Reset() {
	*x = IngestMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbcontroller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestMetadata) ProtoMessage() {}

func (x *IngestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_dbcontroller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestMetadata.ProtoReflect.Descriptor instead.
func (*IngestMetadata) Descriptor() ([]byte, []int) {
	return file_dbcontroller_proto_rawDescGZIP(), []int{0}
}

func (x *IngestMetadata) GetItem() *ItemIdentifier {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *IngestMetadata) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *IngestMetadata) GetItemMetadata() *ItemMetadata {
	if x != nil {
		return x.ItemMetadata
	}
	return nil
}

func (x *IngestMetadata) GetCacheMetadata() *CacheMetadata {
	if x != nil {
		return x.CacheMetadata
	}
	return nil
}

var File_dbcontroller_proto protoreflect.FileDescriptor

var file_dbcontroller_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x44, 0x0a, 0x0c, 0x69, 0x74,
	0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x62,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x47, 0x0a, 0x0d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0xe9, 0x06, 0x0a, 0x0c, 0x44, 0x42,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x22, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64,
	0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x22,
	0x00, 0x12, 0x52, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x25, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x62, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x1b, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x62,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x52, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x1a, 0x1f, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x49, 0x74, 0x65,
	0x6d, 0x1a, 0x28, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x2e, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a,
	0x28, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x67,
	0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x67,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x28, 0x2e, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0a, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x28, 0x2e, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x8f, 0x01, 0x0a, 0x1b, 0x63, 0x68, 0x2e, 0x75, 0x6e, 0x69,
	0x62, 0x61, 0x73, 0x2e, 0x75, 0x62, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x64, 0x62, 0x42, 0x11, 0x44, 0x42, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x65, 0x34, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x64, 0x62,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x55, 0x42, 0x42, 0xaa, 0x02, 0x18, 0x55,
	0x6e, 0x69, 0x62, 0x61, 0x73, 0x2e, 0x55, 0x42, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x42, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dbcontroller_proto_rawDescOnce sync.Once
	file_dbcontroller_proto_rawDescData = file_dbcontroller_proto_rawDesc
)

func file_dbcontroller_proto_rawDescGZIP() []byte {
	file_dbcontroller_proto_rawDescOnce.Do(func() {
		file_dbcontroller_proto_rawDescData = protoimpl.X.CompressGZIP(file_dbcontroller_proto_rawDescData)
	})
	return file_dbcontroller_proto_rawDescData
}

var file_dbcontroller_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_dbcontroller_proto_goTypes = []interface{}{
	(*IngestMetadata)(nil),        // 0: mediaserverdbproto.IngestMetadata
	(*ItemIdentifier)(nil),        // 1: mediaserverdbproto.ItemIdentifier
	(*ItemMetadata)(nil),          // 2: mediaserverdbproto.ItemMetadata
	(*CacheMetadata)(nil),         // 3: mediaserverdbproto.CacheMetadata
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
	(*StorageIdentifier)(nil),     // 5: mediaserverdbproto.StorageIdentifier
	(*CollectionIdentifier)(nil),  // 6: mediaserverdbproto.CollectionIdentifier
	(*PageToken)(nil),             // 7: mediaserverdbproto.PageToken
	(*NewItem)(nil),               // 8: mediaserverdbproto.NewItem
	(*proto.DefaultResponse)(nil), // 9: mediaservergenericproto.DefaultResponse
	(*Item)(nil),                  // 10: mediaserverdbproto.Item
	(*Storage)(nil),               // 11: mediaserverdbproto.Storage
	(*Collection)(nil),            // 12: mediaserverdbproto.Collection
	(*Collections)(nil),           // 13: mediaserverdbproto.Collections
	(*IngestItem)(nil),            // 14: mediaserverdbproto.IngestItem
}
var file_dbcontroller_proto_depIdxs = []int32{
	1,  // 0: mediaserverdbproto.IngestMetadata.item:type_name -> mediaserverdbproto.ItemIdentifier
	2,  // 1: mediaserverdbproto.IngestMetadata.itemMetadata:type_name -> mediaserverdbproto.ItemMetadata
	3,  // 2: mediaserverdbproto.IngestMetadata.cacheMetadata:type_name -> mediaserverdbproto.CacheMetadata
	4,  // 3: mediaserverdbproto.DBController.Ping:input_type -> google.protobuf.Empty
	1,  // 4: mediaserverdbproto.DBController.GetItem:input_type -> mediaserverdbproto.ItemIdentifier
	5,  // 5: mediaserverdbproto.DBController.GetStorage:input_type -> mediaserverdbproto.StorageIdentifier
	6,  // 6: mediaserverdbproto.DBController.GetCollection:input_type -> mediaserverdbproto.CollectionIdentifier
	7,  // 7: mediaserverdbproto.DBController.GetCollections:input_type -> mediaserverdbproto.PageToken
	8,  // 8: mediaserverdbproto.DBController.CreateItem:input_type -> mediaserverdbproto.NewItem
	1,  // 9: mediaserverdbproto.DBController.DeleteItem:input_type -> mediaserverdbproto.ItemIdentifier
	4,  // 10: mediaserverdbproto.DBController.GetIngestItem:input_type -> google.protobuf.Empty
	0,  // 11: mediaserverdbproto.DBController.SetIngestItem:input_type -> mediaserverdbproto.IngestMetadata
	1,  // 12: mediaserverdbproto.DBController.ExistsItem:input_type -> mediaserverdbproto.ItemIdentifier
	9,  // 13: mediaserverdbproto.DBController.Ping:output_type -> mediaservergenericproto.DefaultResponse
	10, // 14: mediaserverdbproto.DBController.GetItem:output_type -> mediaserverdbproto.Item
	11, // 15: mediaserverdbproto.DBController.GetStorage:output_type -> mediaserverdbproto.Storage
	12, // 16: mediaserverdbproto.DBController.GetCollection:output_type -> mediaserverdbproto.Collection
	13, // 17: mediaserverdbproto.DBController.GetCollections:output_type -> mediaserverdbproto.Collections
	9,  // 18: mediaserverdbproto.DBController.CreateItem:output_type -> mediaservergenericproto.DefaultResponse
	9,  // 19: mediaserverdbproto.DBController.DeleteItem:output_type -> mediaservergenericproto.DefaultResponse
	14, // 20: mediaserverdbproto.DBController.GetIngestItem:output_type -> mediaserverdbproto.IngestItem
	9,  // 21: mediaserverdbproto.DBController.SetIngestItem:output_type -> mediaservergenericproto.DefaultResponse
	9,  // 22: mediaserverdbproto.DBController.ExistsItem:output_type -> mediaservergenericproto.DefaultResponse
	13, // [13:23] is the sub-list for method output_type
	3,  // [3:13] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_dbcontroller_proto_init() }
func file_dbcontroller_proto_init() {
	if File_dbcontroller_proto != nil {
		return
	}
	file_storage_proto_init()
	file_collection_proto_init()
	file_item_proto_init()
	file_cache_proto_init()
	file_pageToken_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dbcontroller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngestMetadata); i {
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
			RawDescriptor: file_dbcontroller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dbcontroller_proto_goTypes,
		DependencyIndexes: file_dbcontroller_proto_depIdxs,
		MessageInfos:      file_dbcontroller_proto_msgTypes,
	}.Build()
	File_dbcontroller_proto = out.File
	file_dbcontroller_proto_rawDesc = nil
	file_dbcontroller_proto_goTypes = nil
	file_dbcontroller_proto_depIdxs = nil
}
