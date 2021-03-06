// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: CRUDService.proto

package services

import (
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

type GetPostsRequest_IdType int32

const (
	GetPostsRequest_POST_ID GetPostsRequest_IdType = 0
	GetPostsRequest_USER_ID GetPostsRequest_IdType = 1
	GetPostsRequest_ALL     GetPostsRequest_IdType = 2
)

// Enum value maps for GetPostsRequest_IdType.
var (
	GetPostsRequest_IdType_name = map[int32]string{
		0: "POST_ID",
		1: "USER_ID",
		2: "ALL",
	}
	GetPostsRequest_IdType_value = map[string]int32{
		"POST_ID": 0,
		"USER_ID": 1,
		"ALL":     2,
	}
)

func (x GetPostsRequest_IdType) Enum() *GetPostsRequest_IdType {
	p := new(GetPostsRequest_IdType)
	*p = x
	return p
}

func (x GetPostsRequest_IdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetPostsRequest_IdType) Descriptor() protoreflect.EnumDescriptor {
	return file_CRUDService_proto_enumTypes[0].Descriptor()
}

func (GetPostsRequest_IdType) Type() protoreflect.EnumType {
	return &file_CRUDService_proto_enumTypes[0]
}

func (x GetPostsRequest_IdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetPostsRequest_IdType.Descriptor instead.
func (GetPostsRequest_IdType) EnumDescriptor() ([]byte, []int) {
	return file_CRUDService_proto_rawDescGZIP(), []int{0, 0}
}

type GetPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Needed GetPostsRequest_IdType `protobuf:"varint,2,opt,name=needed,proto3,enum=services.GetPostsRequest_IdType" json:"needed,omitempty"`
}

func (x *GetPostsRequest) Reset() {
	*x = GetPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CRUDService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsRequest) ProtoMessage() {}

func (x *GetPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_CRUDService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsRequest.ProtoReflect.Descriptor instead.
func (*GetPostsRequest) Descriptor() ([]byte, []int) {
	return file_CRUDService_proto_rawDescGZIP(), []int{0}
}

func (x *GetPostsRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetPostsRequest) GetNeeded() GetPostsRequest_IdType {
	if x != nil {
		return x.Needed
	}
	return GetPostsRequest_POST_ID
}

type ContentDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body  string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ContentDTO) Reset() {
	*x = ContentDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CRUDService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentDTO) ProtoMessage() {}

func (x *ContentDTO) ProtoReflect() protoreflect.Message {
	mi := &file_CRUDService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentDTO.ProtoReflect.Descriptor instead.
func (*ContentDTO) Descriptor() ([]byte, []int) {
	return file_CRUDService_proto_rawDescGZIP(), []int{1}
}

func (x *ContentDTO) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ContentDTO) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type PostDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId  int32       `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content *ContentDTO `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *PostDTO) Reset() {
	*x = PostDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CRUDService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostDTO) ProtoMessage() {}

func (x *PostDTO) ProtoReflect() protoreflect.Message {
	mi := &file_CRUDService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostDTO.ProtoReflect.Descriptor instead.
func (*PostDTO) Descriptor() ([]byte, []int) {
	return file_CRUDService_proto_rawDescGZIP(), []int{2}
}

func (x *PostDTO) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PostDTO) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PostDTO) GetContent() *ContentDTO {
	if x != nil {
		return x.Content
	}
	return nil
}

type SavePostDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32       `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content *ContentDTO `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SavePostDTO) Reset() {
	*x = SavePostDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CRUDService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePostDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePostDTO) ProtoMessage() {}

func (x *SavePostDTO) ProtoReflect() protoreflect.Message {
	mi := &file_CRUDService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePostDTO.ProtoReflect.Descriptor instead.
func (*SavePostDTO) Descriptor() ([]byte, []int) {
	return file_CRUDService_proto_rawDescGZIP(), []int{3}
}

func (x *SavePostDTO) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SavePostDTO) GetContent() *ContentDTO {
	if x != nil {
		return x.Content
	}
	return nil
}

type UpdatePostDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content *ContentDTO `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UpdatePostDTO) Reset() {
	*x = UpdatePostDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CRUDService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostDTO) ProtoMessage() {}

func (x *UpdatePostDTO) ProtoReflect() protoreflect.Message {
	mi := &file_CRUDService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostDTO.ProtoReflect.Descriptor instead.
func (*UpdatePostDTO) Descriptor() ([]byte, []int) {
	return file_CRUDService_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePostDTO) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePostDTO) GetContent() *ContentDTO {
	if x != nil {
		return x.Content
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CRUDService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_CRUDService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_CRUDService_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CRUDService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_CRUDService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_CRUDService_proto_rawDescGZIP(), []int{6}
}

func (x *ErrorResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_CRUDService_proto protoreflect.FileDescriptor

var file_CRUDService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x52, 0x55, 0x44, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x88, 0x01,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x38, 0x0a, 0x06, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x06, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x22, 0x2b, 0x0a, 0x06, 0x49,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x22, 0x36, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x44, 0x54, 0x4f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x22, 0x62, 0x0a, 0x07, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x54, 0x4f, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0x56, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x44, 0x54, 0x4f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x44, 0x54, 0x4f, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x4f, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x44, 0x54, 0x4f, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x1f, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25,
	0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xff, 0x01, 0x0a, 0x0b, 0x43, 0x52, 0x55, 0x44, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x61, 0x76,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x54, 0x4f, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x54, 0x4f, 0x12, 0x3a, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x44, 0x54, 0x4f, 0x30, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x54, 0x4f, 0x1a, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6c, 0x65, 0x6d, 0x31, 0x30, 0x39, 0x32, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CRUDService_proto_rawDescOnce sync.Once
	file_CRUDService_proto_rawDescData = file_CRUDService_proto_rawDesc
)

func file_CRUDService_proto_rawDescGZIP() []byte {
	file_CRUDService_proto_rawDescOnce.Do(func() {
		file_CRUDService_proto_rawDescData = protoimpl.X.CompressGZIP(file_CRUDService_proto_rawDescData)
	})
	return file_CRUDService_proto_rawDescData
}

var file_CRUDService_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CRUDService_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_CRUDService_proto_goTypes = []interface{}{
	(GetPostsRequest_IdType)(0), // 0: services.GetPostsRequest.IdType
	(*GetPostsRequest)(nil),     // 1: services.GetPostsRequest
	(*ContentDTO)(nil),          // 2: services.ContentDTO
	(*PostDTO)(nil),             // 3: services.PostDTO
	(*SavePostDTO)(nil),         // 4: services.SavePostDTO
	(*UpdatePostDTO)(nil),       // 5: services.UpdatePostDTO
	(*DeleteRequest)(nil),       // 6: services.DeleteRequest
	(*ErrorResponse)(nil),       // 7: services.ErrorResponse
}
var file_CRUDService_proto_depIdxs = []int32{
	0, // 0: services.GetPostsRequest.needed:type_name -> services.GetPostsRequest.IdType
	2, // 1: services.PostDTO.content:type_name -> services.ContentDTO
	2, // 2: services.SavePostDTO.content:type_name -> services.ContentDTO
	2, // 3: services.UpdatePostDTO.content:type_name -> services.ContentDTO
	4, // 4: services.CRUDService.SavePost:input_type -> services.SavePostDTO
	1, // 5: services.CRUDService.GetPosts:input_type -> services.GetPostsRequest
	6, // 6: services.CRUDService.DeletePost:input_type -> services.DeleteRequest
	5, // 7: services.CRUDService.UpdatePost:input_type -> services.UpdatePostDTO
	3, // 8: services.CRUDService.SavePost:output_type -> services.PostDTO
	3, // 9: services.CRUDService.GetPosts:output_type -> services.PostDTO
	7, // 10: services.CRUDService.DeletePost:output_type -> services.ErrorResponse
	7, // 11: services.CRUDService.UpdatePost:output_type -> services.ErrorResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_CRUDService_proto_init() }
func file_CRUDService_proto_init() {
	if File_CRUDService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CRUDService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostsRequest); i {
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
		file_CRUDService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentDTO); i {
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
		file_CRUDService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostDTO); i {
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
		file_CRUDService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavePostDTO); i {
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
		file_CRUDService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostDTO); i {
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
		file_CRUDService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_CRUDService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
			RawDescriptor: file_CRUDService_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_CRUDService_proto_goTypes,
		DependencyIndexes: file_CRUDService_proto_depIdxs,
		EnumInfos:         file_CRUDService_proto_enumTypes,
		MessageInfos:      file_CRUDService_proto_msgTypes,
	}.Build()
	File_CRUDService_proto = out.File
	file_CRUDService_proto_rawDesc = nil
	file_CRUDService_proto_goTypes = nil
	file_CRUDService_proto_depIdxs = nil
}
