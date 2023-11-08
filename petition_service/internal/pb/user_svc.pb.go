// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: user_svc.proto

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

type GetUserEmailByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserEmailByIdRequest) Reset() {
	*x = GetUserEmailByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserEmailByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserEmailByIdRequest) ProtoMessage() {}

func (x *GetUserEmailByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserEmailByIdRequest.ProtoReflect.Descriptor instead.
func (*GetUserEmailByIdRequest) Descriptor() ([]byte, []int) {
	return file_user_svc_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserEmailByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseMessage) Reset() {
	*x = ResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMessage) ProtoMessage() {}

func (x *ResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_user_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMessage.ProtoReflect.Descriptor instead.
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return file_user_svc_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CheckUserExistenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CheckUserExistenceRequest) Reset() {
	*x = CheckUserExistenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserExistenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserExistenceRequest) ProtoMessage() {}

func (x *CheckUserExistenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserExistenceRequest.ProtoReflect.Descriptor instead.
func (*CheckUserExistenceRequest) Descriptor() ([]byte, []int) {
	return file_user_svc_proto_rawDescGZIP(), []int{2}
}

func (x *CheckUserExistenceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CheckUserExistenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message bool `protobuf:"varint,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CheckUserExistenceResponse) Reset() {
	*x = CheckUserExistenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_svc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserExistenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserExistenceResponse) ProtoMessage() {}

func (x *CheckUserExistenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_svc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserExistenceResponse.ProtoReflect.Descriptor instead.
func (*CheckUserExistenceResponse) Descriptor() ([]byte, []int) {
	return file_user_svc_proto_rawDescGZIP(), []int{3}
}

func (x *CheckUserExistenceResponse) GetMessage() bool {
	if x != nil {
		return x.Message
	}
	return false
}

type GetAdminEmailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminEmails []string `protobuf:"bytes,1,rep,name=admin_emails,json=adminEmails,proto3" json:"admin_emails,omitempty"`
}

func (x *GetAdminEmailsResponse) Reset() {
	*x = GetAdminEmailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_svc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminEmailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminEmailsResponse) ProtoMessage() {}

func (x *GetAdminEmailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_svc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminEmailsResponse.ProtoReflect.Descriptor instead.
func (*GetAdminEmailsResponse) Descriptor() ([]byte, []int) {
	return file_user_svc_proto_rawDescGZIP(), []int{4}
}

func (x *GetAdminEmailsResponse) GetAdminEmails() []string {
	if x != nil {
		return x.AdminEmails
	}
	return nil
}

type GetAdminEmailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAdminEmailsRequest) Reset() {
	*x = GetAdminEmailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_svc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminEmailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminEmailsRequest) ProtoMessage() {}

func (x *GetAdminEmailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_svc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminEmailsRequest.ProtoReflect.Descriptor instead.
func (*GetAdminEmailsRequest) Descriptor() ([]byte, []int) {
	return file_user_svc_proto_rawDescGZIP(), []int{5}
}

var File_user_svc_proto protoreflect.FileDescriptor

var file_user_svc_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x2b, 0x0a, 0x19, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x1a,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x3b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x73, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x83, 0x02, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x59, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_svc_proto_rawDescOnce sync.Once
	file_user_svc_proto_rawDescData = file_user_svc_proto_rawDesc
)

func file_user_svc_proto_rawDescGZIP() []byte {
	file_user_svc_proto_rawDescOnce.Do(func() {
		file_user_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_svc_proto_rawDescData)
	})
	return file_user_svc_proto_rawDescData
}

var file_user_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_svc_proto_goTypes = []interface{}{
	(*GetUserEmailByIdRequest)(nil),    // 0: proto.GetUserEmailByIdRequest
	(*ResponseMessage)(nil),            // 1: proto.ResponseMessage
	(*CheckUserExistenceRequest)(nil),  // 2: proto.CheckUserExistenceRequest
	(*CheckUserExistenceResponse)(nil), // 3: proto.CheckUserExistenceResponse
	(*GetAdminEmailsResponse)(nil),     // 4: proto.GetAdminEmailsResponse
	(*GetAdminEmailsRequest)(nil),      // 5: proto.GetAdminEmailsRequest
}
var file_user_svc_proto_depIdxs = []int32{
	0, // 0: proto.UserService.GetUserEmailById:input_type -> proto.GetUserEmailByIdRequest
	2, // 1: proto.UserService.CheckUserExistence:input_type -> proto.CheckUserExistenceRequest
	5, // 2: proto.UserService.GetAdminEmails:input_type -> proto.GetAdminEmailsRequest
	1, // 3: proto.UserService.GetUserEmailById:output_type -> proto.ResponseMessage
	3, // 4: proto.UserService.CheckUserExistence:output_type -> proto.CheckUserExistenceResponse
	4, // 5: proto.UserService.GetAdminEmails:output_type -> proto.GetAdminEmailsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_svc_proto_init() }
func file_user_svc_proto_init() {
	if File_user_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserEmailByIdRequest); i {
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
		file_user_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMessage); i {
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
		file_user_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserExistenceRequest); i {
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
		file_user_svc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserExistenceResponse); i {
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
		file_user_svc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminEmailsResponse); i {
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
		file_user_svc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminEmailsRequest); i {
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
			RawDescriptor: file_user_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_svc_proto_goTypes,
		DependencyIndexes: file_user_svc_proto_depIdxs,
		MessageInfos:      file_user_svc_proto_msgTypes,
	}.Build()
	File_user_svc_proto = out.File
	file_user_svc_proto_rawDesc = nil
	file_user_svc_proto_goTypes = nil
	file_user_svc_proto_depIdxs = nil
}
