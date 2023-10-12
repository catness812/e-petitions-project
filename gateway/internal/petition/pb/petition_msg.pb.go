// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
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

	Id          uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Category    string  `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Image       string  `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Status      *Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	UserId      uint32  `protobuf:"varint,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VoteGoal    uint32  `protobuf:"varint,8,opt,name=vote_goal,json=voteGoal,proto3" json:"vote_goal,omitempty"`
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

func (x *Petition) GetId() uint32 {
	if x != nil {
		return x.Id
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

func (x *Petition) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Petition) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Petition) GetVoteGoal() uint32 {
	if x != nil {
		return x.VoteGoal
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Status) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type CreateVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PetitionId uint32 `protobuf:"varint,2,opt,name=petition_id,json=petitionId,proto3" json:"petition_id,omitempty"`
}

func (x *CreateVoteRequest) Reset() {
	*x = CreateVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVoteRequest) ProtoMessage() {}

func (x *CreateVoteRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateVoteRequest.ProtoReflect.Descriptor instead.
func (*CreateVoteRequest) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{2}
}

func (x *CreateVoteRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateVoteRequest) GetPetitionId() uint32 {
	if x != nil {
		return x.PetitionId
	}
	return 0
}

type CreatePetitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Image       string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	UserId      uint32 `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Category    string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	VoteGoal    uint32 `protobuf:"varint,8,opt,name=vote_goal,json=voteGoal,proto3" json:"vote_goal,omitempty"`
}

func (x *CreatePetitionRequest) Reset() {
	*x = CreatePetitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePetitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetitionRequest) ProtoMessage() {}

func (x *CreatePetitionRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreatePetitionRequest.ProtoReflect.Descriptor instead.
func (*CreatePetitionRequest) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePetitionRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePetitionRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreatePetitionRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreatePetitionRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreatePetitionRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreatePetitionRequest) GetVoteGoal() uint32 {
	if x != nil {
		return x.VoteGoal
	}
	return 0
}

type PetitionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PetitionId) Reset() {
	*x = PetitionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PetitionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PetitionId) ProtoMessage() {}

func (x *PetitionId) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PetitionId.ProtoReflect.Descriptor instead.
func (*PetitionId) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{4}
}

func (x *PetitionId) GetId() uint32 {
	if x != nil {
		return x.Id
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
		mi := &file_petition_msg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPetitionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPetitionsRequest) ProtoMessage() {}

func (x *GetPetitionsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPetitionsRequest.ProtoReflect.Descriptor instead.
func (*GetPetitionsRequest) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{5}
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
		mi := &file_petition_msg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPetitionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPetitionsResponse) ProtoMessage() {}

func (x *GetPetitionsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPetitionsResponse.ProtoReflect.Descriptor instead.
func (*GetPetitionsResponse) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{6}
}

func (x *GetPetitionsResponse) GetPetitions() []*Petition {
	if x != nil {
		return x.Petitions
	}
	return nil
}

type UpdatePetitionStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdatePetitionStatusRequest) Reset() {
	*x = UpdatePetitionStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePetitionStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePetitionStatusRequest) ProtoMessage() {}

func (x *UpdatePetitionStatusRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdatePetitionStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdatePetitionStatusRequest) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePetitionStatusRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePetitionStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetUserPetitionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page   uint32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit  uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetUserPetitionsRequest) Reset() {
	*x = GetUserPetitionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPetitionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPetitionsRequest) ProtoMessage() {}

func (x *GetUserPetitionsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetUserPetitionsRequest.ProtoReflect.Descriptor instead.
func (*GetUserPetitionsRequest) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserPetitionsRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetUserPetitionsRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetUserPetitionsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetUserPetitionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Petitions []*Petition `protobuf:"bytes,1,rep,name=petitions,proto3" json:"petitions,omitempty"`
}

func (x *GetUserPetitionsResponse) Reset() {
	*x = GetUserPetitionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPetitionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPetitionsResponse) ProtoMessage() {}

func (x *GetUserPetitionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petition_msg_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPetitionsResponse.ProtoReflect.Descriptor instead.
func (*GetUserPetitionsResponse) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserPetitionsResponse) GetPetitions() []*Petition {
	if x != nil {
		return x.Petitions
	}
	return nil
}

type GetUserVotedPetitionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page   uint32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit  uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetUserVotedPetitionsRequest) Reset() {
	*x = GetUserVotedPetitionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserVotedPetitionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserVotedPetitionsRequest) ProtoMessage() {}

func (x *GetUserVotedPetitionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petition_msg_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserVotedPetitionsRequest.ProtoReflect.Descriptor instead.
func (*GetUserVotedPetitionsRequest) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserVotedPetitionsRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetUserVotedPetitionsRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetUserVotedPetitionsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetUserVotedPetitionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Petitions []*Petition `protobuf:"bytes,1,rep,name=petitions,proto3" json:"petitions,omitempty"`
}

func (x *GetUserVotedPetitionsResponse) Reset() {
	*x = GetUserVotedPetitionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petition_msg_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserVotedPetitionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserVotedPetitionsResponse) ProtoMessage() {}

func (x *GetUserVotedPetitionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petition_msg_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserVotedPetitionsResponse.ProtoReflect.Descriptor instead.
func (*GetUserVotedPetitionsResponse) Descriptor() ([]byte, []int) {
	return file_petition_msg_proto_rawDescGZIP(), []int{11}
}

func (x *GetUserVotedPetitionsResponse) GetPetitions() []*Petition {
	if x != nil {
		return x.Petitions
	}
	return nil
}

var File_petition_msg_proto protoreflect.FileDescriptor

var file_petition_msg_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x08,
	0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x22,
	0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22,
	0x4d, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xb7,
	0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x76,
	0x6f, 0x74, 0x65, 0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x76, 0x6f, 0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x22, 0x1c, 0x0a, 0x0a, 0x50, 0x65, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x45, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x09, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x45,
	0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x49, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x09, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x61,
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x74, 0x65, 0x64, 0x50, 0x65,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x4e, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x74, 0x65,
	0x64, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_petition_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_petition_msg_proto_goTypes = []interface{}{
	(*Petition)(nil),                      // 0: proto.Petition
	(*Status)(nil),                        // 1: proto.Status
	(*CreateVoteRequest)(nil),             // 2: proto.CreateVoteRequest
	(*CreatePetitionRequest)(nil),         // 3: proto.CreatePetitionRequest
	(*PetitionId)(nil),                    // 4: proto.PetitionId
	(*GetPetitionsRequest)(nil),           // 5: proto.GetPetitionsRequest
	(*GetPetitionsResponse)(nil),          // 6: proto.GetPetitionsResponse
	(*UpdatePetitionStatusRequest)(nil),   // 7: proto.UpdatePetitionStatusRequest
	(*GetUserPetitionsRequest)(nil),       // 8: proto.GetUserPetitionsRequest
	(*GetUserPetitionsResponse)(nil),      // 9: proto.GetUserPetitionsResponse
	(*GetUserVotedPetitionsRequest)(nil),  // 10: proto.GetUserVotedPetitionsRequest
	(*GetUserVotedPetitionsResponse)(nil), // 11: proto.GetUserVotedPetitionsResponse
}
var file_petition_msg_proto_depIdxs = []int32{
	1, // 0: proto.Petition.status:type_name -> proto.Status
	0, // 1: proto.GetPetitionsResponse.petitions:type_name -> proto.Petition
	0, // 2: proto.GetUserPetitionsResponse.petitions:type_name -> proto.Petition
	0, // 3: proto.GetUserVotedPetitionsResponse.petitions:type_name -> proto.Petition
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
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
			switch v := v.(*Status); i {
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
			switch v := v.(*CreateVoteRequest); i {
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
		file_petition_msg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PetitionId); i {
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
		file_petition_msg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_petition_msg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePetitionStatusRequest); i {
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
			switch v := v.(*GetUserPetitionsRequest); i {
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
		file_petition_msg_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPetitionsResponse); i {
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
		file_petition_msg_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserVotedPetitionsRequest); i {
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
		file_petition_msg_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserVotedPetitionsResponse); i {
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
			NumMessages:   12,
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
