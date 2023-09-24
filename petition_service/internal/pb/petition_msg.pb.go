// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: petition_msg.proto

package pb

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

type Petition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetitionId  uint32 `protobuf:"varint,1,opt,name=petition_id,json=petitionId,proto3" json:"petition_id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Category    string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Image       string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Status      uint32 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	UserId      uint32 `protobuf:"varint,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Petition) Reset() {
	*x = Petition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Petition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Petition) ProtoMessage() {}

func (x *Petition) ProtoReflect() protoreflect.Message {
	mi := &file_petition_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Petition.ProtoReflect.Descriptor instead.
func (*Petition) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Petition) GetPetitionId() uint32 {
	if x != nil {
		return x.PetitionId
	}
	return 0
}

func (x *Petition) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Petition) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Petition) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Petition) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Petition) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Petition) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreatePetitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Petition *Petition `protobuf:"bytes,1,opt,name=petition,proto3" json:"petition,omitempty"`
}

func (x *CreatePetitionRequest) Reset() {
	*x = CreatePetitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePetitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetitionRequest) ProtoMessage() {}

func (x *CreatePetitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petition_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetitionRequest.ProtoReflect.Descriptor instead.
func (*CreatePetitionRequest) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePetitionRequest) GetPetition() *Petition {
	if x != nil {
		return x.Petition
	}
	return nil
}

type CreatePetitionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetitionId uint32 `protobuf:"varint,1,opt,name=petition_id,json=petitionId,proto3" json:"petition_id,omitempty"`
}

func (x *CreatePetitionResponse) Reset() {
	*x = CreatePetitionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePetitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetitionResponse) ProtoMessage() {}

func (x *CreatePetitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petition_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetitionResponse.ProtoReflect.Descriptor instead.
func (*CreatePetitionResponse) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePetitionResponse) GetPetitionId() uint32 {
	if x != nil {
		return x.PetitionId
	}
	return 0
}

type GetPetitionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetPetitionsRequest) Reset() {
	*x = GetPetitionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPetitionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPetitionsRequest) ProtoMessage() {}

func (x *GetPetitionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petition_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPetitionsRequest.ProtoReflect.Descriptor instead.
func (*GetPetitionsRequest) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{3}
}

func (x *GetPetitionsRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetPetitionsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetPetitionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Petitions []*Petition `protobuf:"bytes,1,rep,name=petitions,proto3" json:"petitions,omitempty"`
}

func (x *GetPetitionsResponse) Reset() {
	*x = GetPetitionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPetitionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPetitionsResponse) ProtoMessage() {}

func (x *GetPetitionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petition_msg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPetitionsResponse.ProtoReflect.Descriptor instead.
func (*GetPetitionsResponse) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{4}
}

func (x *GetPetitionsResponse) GetPetitions() []*Petition {
	if x != nil {
		return x.Petitions
	}
	return nil
}

type UpdatePetitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdatePetitionRequest) Reset() {
	*x = UpdatePetitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePetitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePetitionRequest) ProtoMessage() {}

func (x *UpdatePetitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petition_msg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePetitionRequest.ProtoReflect.Descriptor instead.
func (*UpdatePetitionRequest) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePetitionRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePetitionRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdatePetitionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdatePetitionResponse) Reset() {
	*x = UpdatePetitionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePetitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePetitionResponse) ProtoMessage() {}

func (x *UpdatePetitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petition_msg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePetitionResponse.ProtoReflect.Descriptor instead.
func (*UpdatePetitionResponse) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePetitionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeletePetitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePetitionRequest) Reset() {
	*x = DeletePetitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePetitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePetitionRequest) ProtoMessage() {}

func (x *DeletePetitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petition_msg_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePetitionRequest.ProtoReflect.Descriptor instead.
func (*DeletePetitionRequest) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePetitionRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeletePetitionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeletePetitionResponse) Reset() {
	*x = DeletePetitionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePetitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePetitionResponse) ProtoMessage() {}

func (x *DeletePetitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petition_msg_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePetitionResponse.ProtoReflect.Descriptor instead.
func (*DeletePetitionResponse) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePetitionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_petition_msg_proto protoreflect.FileDescriptor

var file_petition_msg_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x08,
	0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x65, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70,
	0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x08, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x39, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x45, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d,
	0x0a, 0x09, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3f, 0x0a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x32,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x16, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_petition_msg_proto_rawDescOnce sync.Once
	file_petition_msg_proto_rawDescData = file_petition_msg_proto_rawDesc
)

func file_petition_msg_proto_rawDescGZIP() []byte {
	file_petition_msg_proto_rawDescOnce.Do(func() {
		file_petition_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_petition_msg_proto_rawDescData)
	})
	return file_petition_msg_proto_rawDescData
}

var file_petition_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_petition_msg_proto_goTypes = []interface{}{
	(*Petition)(nil),               // 0: proto.Petition
	(*CreatePetitionRequest)(nil),  // 1: proto.CreatePetitionRequest
	(*CreatePetitionResponse)(nil), // 2: proto.CreatePetitionResponse
	(*GetPetitionsRequest)(nil),    // 3: proto.GetPetitionsRequest
	(*GetPetitionsResponse)(nil),   // 4: proto.GetPetitionsResponse
	(*UpdatePetitionRequest)(nil),  // 5: proto.UpdatePetitionRequest
	(*UpdatePetitionResponse)(nil), // 6: proto.UpdatePetitionResponse
	(*DeletePetitionRequest)(nil),  // 7: proto.DeletePetitionRequest
	(*DeletePetitionResponse)(nil), // 8: proto.DeletePetitionResponse
}
var file_petition_msg_proto_depIdxs = []int32{
	0, // 0: proto.CreatePetitionRequest.petition:type_name -> proto.Petition
	0, // 1: proto.GetPetitionsResponse.petitions:type_name -> proto.Petition
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_petition_msg_proto_init() }
func file_petition_msg_proto_init() {
	if File_petition_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_petition_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Petition); i {
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
		file_petition_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePetitionRequest); i {
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
		file_petition_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePetitionResponse); i {
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
		file_petition_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPetitionsRequest); i {
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
		file_petition_msg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPetitionsResponse); i {
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
		file_petition_msg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePetitionRequest); i {
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
		file_petition_msg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePetitionResponse); i {
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
		file_petition_msg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePetitionRequest); i {
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
		file_petition_msg_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePetitionResponse); i {
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
			RawDescriptor: file_petition_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_petition_msg_proto_goTypes,
		DependencyIndexes: file_petition_msg_proto_depIdxs,
		MessageInfos:      file_petition_msg_proto_msgTypes,
	}.Build()
	File_petition_msg_proto = out.File
	file_petition_msg_proto_rawDesc = nil
	file_petition_msg_proto_goTypes = nil
	file_petition_msg_proto_depIdxs = nil
}
