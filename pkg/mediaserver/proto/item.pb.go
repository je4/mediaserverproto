// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: item.proto

package proto

import (
	proto "github.com/je4/genericproto/v2/pkg/generic/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IngestType int32

const (
	IngestType_KEEP IngestType = 0
	IngestType_COPY IngestType = 1
	IngestType_MOVE IngestType = 2
)

// Enum value maps for IngestType.
var (
	IngestType_name = map[int32]string{
		0: "KEEP",
		1: "COPY",
		2: "MOVE",
	}
	IngestType_value = map[string]int32{
		"KEEP": 0,
		"COPY": 1,
		"MOVE": 2,
	}
)

func (x IngestType) Enum() *IngestType {
	p := new(IngestType)
	*p = x
	return p
}

func (x IngestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IngestType) Descriptor() protoreflect.EnumDescriptor {
	return file_item_proto_enumTypes[0].Descriptor()
}

func (IngestType) Type() protoreflect.EnumType {
	return &file_item_proto_enumTypes[0]
}

func (x IngestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IngestType.Descriptor instead.
func (IngestType) EnumDescriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{0}
}

type ItemIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection string `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	Signature  string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ItemIdentifier) Reset() {
	*x = ItemIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemIdentifier) ProtoMessage() {}

func (x *ItemIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemIdentifier.ProtoReflect.Descriptor instead.
func (*ItemIdentifier) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{0}
}

func (x *ItemIdentifier) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *ItemIdentifier) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type NewItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier    *ItemIdentifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Urn           string          `protobuf:"bytes,2,opt,name=urn,proto3" json:"urn,omitempty"`
	Public        *bool           `protobuf:"varint,3,opt,name=public,proto3,oneof" json:"public,omitempty"`
	Parent        *ItemIdentifier `protobuf:"bytes,4,opt,name=parent,proto3,oneof" json:"parent,omitempty"`
	PublicActions []byte          `protobuf:"bytes,5,opt,name=publicActions,proto3,oneof" json:"publicActions,omitempty"`
	IngestType    *IngestType     `protobuf:"varint,6,opt,name=ingestType,proto3,enum=mediaserverproto.IngestType,oneof" json:"ingestType,omitempty"`
}

func (x *NewItem) Reset() {
	*x = NewItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewItem) ProtoMessage() {}

func (x *NewItem) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewItem.ProtoReflect.Descriptor instead.
func (*NewItem) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{1}
}

func (x *NewItem) GetIdentifier() *ItemIdentifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *NewItem) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *NewItem) GetPublic() bool {
	if x != nil && x.Public != nil {
		return *x.Public
	}
	return false
}

func (x *NewItem) GetParent() *ItemIdentifier {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *NewItem) GetPublicActions() []byte {
	if x != nil {
		return x.PublicActions
	}
	return nil
}

func (x *NewItem) GetIngestType() IngestType {
	if x != nil && x.IngestType != nil {
		return *x.IngestType
	}
	return IngestType_KEEP
}

type IngestItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier *ItemIdentifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Urn        string          `protobuf:"bytes,2,opt,name=urn,proto3" json:"urn,omitempty"`
	IngestType IngestType      `protobuf:"varint,4,opt,name=ingestType,proto3,enum=mediaserverproto.IngestType" json:"ingestType,omitempty"`
	Collection *Collection     `protobuf:"bytes,5,opt,name=collection,proto3" json:"collection,omitempty"`
}

func (x *IngestItem) Reset() {
	*x = IngestItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestItem) ProtoMessage() {}

func (x *IngestItem) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestItem.ProtoReflect.Descriptor instead.
func (*IngestItem) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{2}
}

func (x *IngestItem) GetIdentifier() *ItemIdentifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *IngestItem) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *IngestItem) GetIngestType() IngestType {
	if x != nil {
		return x.IngestType
	}
	return IngestType_KEEP
}

func (x *IngestItem) GetCollection() *Collection {
	if x != nil {
		return x.Collection
	}
	return nil
}

type ItemMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       *string `protobuf:"bytes,3,opt,name=type,proto3,oneof" json:"type,omitempty"`
	Subtype    *string `protobuf:"bytes,4,opt,name=subtype,proto3,oneof" json:"subtype,omitempty"`
	Objecttype *string `protobuf:"bytes,5,opt,name=objecttype,proto3,oneof" json:"objecttype,omitempty"`
	Mimetype   *string `protobuf:"bytes,6,opt,name=mimetype,proto3,oneof" json:"mimetype,omitempty"`
	Sha512     *string `protobuf:"bytes,8,opt,name=sha512,proto3,oneof" json:"sha512,omitempty"`
}

func (x *ItemMetadata) Reset() {
	*x = ItemMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemMetadata) ProtoMessage() {}

func (x *ItemMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemMetadata.ProtoReflect.Descriptor instead.
func (*ItemMetadata) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{3}
}

func (x *ItemMetadata) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *ItemMetadata) GetSubtype() string {
	if x != nil && x.Subtype != nil {
		return *x.Subtype
	}
	return ""
}

func (x *ItemMetadata) GetObjecttype() string {
	if x != nil && x.Objecttype != nil {
		return *x.Objecttype
	}
	return ""
}

func (x *ItemMetadata) GetMimetype() string {
	if x != nil && x.Mimetype != nil {
		return *x.Mimetype
	}
	return ""
}

func (x *ItemMetadata) GetSha512() string {
	if x != nil && x.Sha512 != nil {
		return *x.Sha512
	}
	return ""
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier    *ItemIdentifier        `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Urn           string                 `protobuf:"bytes,2,opt,name=urn,proto3" json:"urn,omitempty"`
	Error         *string                `protobuf:"bytes,3,opt,name=error,proto3,oneof" json:"error,omitempty"`
	Created       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Disabled      bool                   `protobuf:"varint,6,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Public        bool                   `protobuf:"varint,7,opt,name=public,proto3" json:"public,omitempty"`
	PublicActions []string               `protobuf:"bytes,8,rep,name=publicActions,proto3" json:"publicActions,omitempty"`
	Status        string                 `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	Metadata      *ItemMetadata          `protobuf:"bytes,10,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Parent        *ItemIdentifier        `protobuf:"bytes,11,opt,name=parent,proto3,oneof" json:"parent,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{4}
}

func (x *Item) GetIdentifier() *ItemIdentifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *Item) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *Item) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

func (x *Item) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Item) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *Item) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Item) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *Item) GetPublicActions() []string {
	if x != nil {
		return x.PublicActions
	}
	return nil
}

func (x *Item) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Item) GetMetadata() *ItemMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Item) GetParent() *ItemIdentifier {
	if x != nil {
		return x.Parent
	}
	return nil
}

type ItemsResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items        []*Item             `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	PageResponse *proto.PageResponse `protobuf:"bytes,2,opt,name=pageResponse,proto3" json:"pageResponse,omitempty"`
}

func (x *ItemsResult) Reset() {
	*x = ItemsResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemsResult) ProtoMessage() {}

func (x *ItemsResult) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemsResult.ProtoReflect.Descriptor instead.
func (*ItemsResult) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{5}
}

func (x *ItemsResult) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ItemsResult) GetPageResponse() *proto.PageResponse {
	if x != nil {
		return x.PageResponse
	}
	return nil
}

type ItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier  *ItemIdentifier    `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	PageRequest *proto.PageRequest `protobuf:"bytes,2,opt,name=pageRequest,proto3" json:"pageRequest,omitempty"`
}

func (x *ItemsRequest) Reset() {
	*x = ItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemsRequest) ProtoMessage() {}

func (x *ItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemsRequest.ProtoReflect.Descriptor instead.
func (*ItemsRequest) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{6}
}

func (x *ItemsRequest) GetIdentifier() *ItemIdentifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *ItemsRequest) GetPageRequest() *proto.PageRequest {
	if x != nil {
		return x.PageRequest
	}
	return nil
}

var File_item_proto protoreflect.FileDescriptor

var file_item_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a,
	0x0e, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xde, 0x02,
	0x0a, 0x07, 0x4e, 0x65, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x40, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x12, 0x1b, 0x0a,
	0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52,
	0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x48, 0x01, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x02, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x0a, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x48, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xdc,
	0x01, 0x0a, 0x0a, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x40, 0x0a,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6e, 0x12, 0x3c, 0x0a, 0x0a, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x3c, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe5, 0x01,
	0x0a, 0x0c, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x17,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x73, 0x75, 0x62, 0x74,
	0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6d,
	0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x73, 0x68, 0x61, 0x35, 0x31, 0x32, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x06,
	0x73, 0x68, 0x61, 0x35, 0x31, 0x32, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x73, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73,
	0x68, 0x61, 0x35, 0x31, 0x32, 0x22, 0xe3, 0x03, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x40,
	0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6e, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x24, 0x0a,
	0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x48, 0x01, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x7b, 0x0a, 0x0b, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3e, 0x0a, 0x0c, 0x70, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x70, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x0c, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x70,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x70, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2a, 0x2a, 0x0a, 0x0a, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4b, 0x45, 0x45, 0x50, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x43, 0x4f, 0x50, 0x59, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x4f,
	0x56, 0x45, 0x10, 0x02, 0x42, 0x8a, 0x01, 0x0a, 0x18, 0x63, 0x68, 0x2e, 0x75, 0x6e, 0x69, 0x62,
	0x61, 0x73, 0x2e, 0x75, 0x62, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x42, 0x14, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x65, 0x34, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x55, 0x42, 0x42, 0xaa, 0x02, 0x15, 0x55, 0x6e, 0x69, 0x62,
	0x61, 0x73, 0x2e, 0x55, 0x42, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_item_proto_rawDescOnce sync.Once
	file_item_proto_rawDescData = file_item_proto_rawDesc
)

func file_item_proto_rawDescGZIP() []byte {
	file_item_proto_rawDescOnce.Do(func() {
		file_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_item_proto_rawDescData)
	})
	return file_item_proto_rawDescData
}

var file_item_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_item_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_item_proto_goTypes = []interface{}{
	(IngestType)(0),               // 0: mediaserverproto.IngestType
	(*ItemIdentifier)(nil),        // 1: mediaserverproto.ItemIdentifier
	(*NewItem)(nil),               // 2: mediaserverproto.NewItem
	(*IngestItem)(nil),            // 3: mediaserverproto.IngestItem
	(*ItemMetadata)(nil),          // 4: mediaserverproto.ItemMetadata
	(*Item)(nil),                  // 5: mediaserverproto.Item
	(*ItemsResult)(nil),           // 6: mediaserverproto.ItemsResult
	(*ItemsRequest)(nil),          // 7: mediaserverproto.ItemsRequest
	(*Collection)(nil),            // 8: mediaserverproto.Collection
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*proto.PageResponse)(nil),    // 10: genericproto.PageResponse
	(*proto.PageRequest)(nil),     // 11: genericproto.PageRequest
}
var file_item_proto_depIdxs = []int32{
	1,  // 0: mediaserverproto.NewItem.identifier:type_name -> mediaserverproto.ItemIdentifier
	1,  // 1: mediaserverproto.NewItem.parent:type_name -> mediaserverproto.ItemIdentifier
	0,  // 2: mediaserverproto.NewItem.ingestType:type_name -> mediaserverproto.IngestType
	1,  // 3: mediaserverproto.IngestItem.identifier:type_name -> mediaserverproto.ItemIdentifier
	0,  // 4: mediaserverproto.IngestItem.ingestType:type_name -> mediaserverproto.IngestType
	8,  // 5: mediaserverproto.IngestItem.collection:type_name -> mediaserverproto.Collection
	1,  // 6: mediaserverproto.Item.identifier:type_name -> mediaserverproto.ItemIdentifier
	9,  // 7: mediaserverproto.Item.created:type_name -> google.protobuf.Timestamp
	9,  // 8: mediaserverproto.Item.updated:type_name -> google.protobuf.Timestamp
	4,  // 9: mediaserverproto.Item.metadata:type_name -> mediaserverproto.ItemMetadata
	1,  // 10: mediaserverproto.Item.parent:type_name -> mediaserverproto.ItemIdentifier
	5,  // 11: mediaserverproto.ItemsResult.items:type_name -> mediaserverproto.Item
	10, // 12: mediaserverproto.ItemsResult.pageResponse:type_name -> genericproto.PageResponse
	1,  // 13: mediaserverproto.ItemsRequest.identifier:type_name -> mediaserverproto.ItemIdentifier
	11, // 14: mediaserverproto.ItemsRequest.pageRequest:type_name -> genericproto.PageRequest
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_item_proto_init() }
func file_item_proto_init() {
	if File_item_proto != nil {
		return
	}
	file_collection_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemIdentifier); i {
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
		file_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewItem); i {
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
		file_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngestItem); i {
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
		file_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemMetadata); i {
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
		file_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_item_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemsResult); i {
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
		file_item_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemsRequest); i {
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
	file_item_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_item_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_item_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_item_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_item_proto_goTypes,
		DependencyIndexes: file_item_proto_depIdxs,
		EnumInfos:         file_item_proto_enumTypes,
		MessageInfos:      file_item_proto_msgTypes,
	}.Build()
	File_item_proto = out.File
	file_item_proto_rawDesc = nil
	file_item_proto_goTypes = nil
	file_item_proto_depIdxs = nil
}
