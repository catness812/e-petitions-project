// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: petition_svc.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_petition_svc_proto protoreflect.FileDescriptor

var file_petition_svc_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x65, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xeb, 0x07, 0x0a,
	0x0f, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x35, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x46, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x3b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x12,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3e, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x53, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x62, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x74,
	0x65, 0x64, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x74, 0x65, 0x64,
	0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x56, 0x6f, 0x74, 0x65, 0x64, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x66, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5d, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x50, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x16, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x54, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_petition_svc_proto_goTypes = []interface{}{
	(*PetitionId)(nil),                    // 0: proto.PetitionId
	(*CreatePetitionRequest)(nil),         // 1: proto.CreatePetitionRequest
	(*GetPetitionsRequest)(nil),           // 2: proto.GetPetitionsRequest
	(*UpdatePetitionStatusRequest)(nil),   // 3: proto.UpdatePetitionStatusRequest
	(*UpdatePetitionRequest)(nil),         // 4: proto.UpdatePetitionRequest
	(*CreateVoteRequest)(nil),             // 5: proto.CreateVoteRequest
	(*GetUserPetitionsRequest)(nil),       // 6: proto.GetUserPetitionsRequest
	(*GetUserVotedPetitionsRequest)(nil),  // 7: proto.GetUserVotedPetitionsRequest
	(*Petition)(nil),                      // 8: proto.Petition
	(*PetitionSuggestionRequest)(nil),     // 9: proto.PetitionSuggestionRequest
	(*SearchPetitionsByTitRequest)(nil),   // 10: proto.SearchPetitionsByTitRequest
	(*GetPetitionsResponse)(nil),          // 11: proto.GetPetitionsResponse
	(*emptypb.Empty)(nil),                 // 12: google.protobuf.Empty
	(*GetUserPetitionsResponse)(nil),      // 13: proto.GetUserPetitionsResponse
	(*GetUserVotedPetitionsResponse)(nil), // 14: proto.GetUserVotedPetitionsResponse
	(*PetitionSuggestionResponse)(nil),    // 15: proto.PetitionSuggestionResponse
}
var file_petition_svc_proto_depIdxs = []int32{
	0,  // 0: proto.PetitionService.GetPetitionById:input_type -> proto.PetitionId
	1,  // 1: proto.PetitionService.CreatePetition:input_type -> proto.CreatePetitionRequest
	2,  // 2: proto.PetitionService.GetPetitions:input_type -> proto.GetPetitionsRequest
	3,  // 3: proto.PetitionService.UpdatePetitionStatus:input_type -> proto.UpdatePetitionStatusRequest
	4,  // 4: proto.PetitionService.UpdatePetition:input_type -> proto.UpdatePetitionRequest
	0,  // 5: proto.PetitionService.DeletePetition:input_type -> proto.PetitionId
	0,  // 6: proto.PetitionService.ValidatePetitionId:input_type -> proto.PetitionId
	5,  // 7: proto.PetitionService.CreateVote:input_type -> proto.CreateVoteRequest
	6,  // 8: proto.PetitionService.GetUserPetitions:input_type -> proto.GetUserPetitionsRequest
	7,  // 9: proto.PetitionService.GetUserVotedPetitions:input_type -> proto.GetUserVotedPetitionsRequest
	8,  // 10: proto.PetitionService.CheckIfPetitionsExpired:input_type -> proto.Petition
	9,  // 11: proto.PetitionService.GetAllSimilarPetitions:input_type -> proto.PetitionSuggestionRequest
	10, // 12: proto.PetitionService.SearchPetitionsByTitle:input_type -> proto.SearchPetitionsByTitRequest
	8,  // 13: proto.PetitionService.GetPetitionById:output_type -> proto.Petition
	0,  // 14: proto.PetitionService.CreatePetition:output_type -> proto.PetitionId
	11, // 15: proto.PetitionService.GetPetitions:output_type -> proto.GetPetitionsResponse
	12, // 16: proto.PetitionService.UpdatePetitionStatus:output_type -> google.protobuf.Empty
	12, // 17: proto.PetitionService.UpdatePetition:output_type -> google.protobuf.Empty
	12, // 18: proto.PetitionService.DeletePetition:output_type -> google.protobuf.Empty
	12, // 19: proto.PetitionService.ValidatePetitionId:output_type -> google.protobuf.Empty
	12, // 20: proto.PetitionService.CreateVote:output_type -> google.protobuf.Empty
	13, // 21: proto.PetitionService.GetUserPetitions:output_type -> proto.GetUserPetitionsResponse
	14, // 22: proto.PetitionService.GetUserVotedPetitions:output_type -> proto.GetUserVotedPetitionsResponse
	12, // 23: proto.PetitionService.CheckIfPetitionsExpired:output_type -> google.protobuf.Empty
	15, // 24: proto.PetitionService.GetAllSimilarPetitions:output_type -> proto.PetitionSuggestionResponse
	15, // 25: proto.PetitionService.SearchPetitionsByTitle:output_type -> proto.PetitionSuggestionResponse
	13, // [13:26] is the sub-list for method output_type
	0,  // [0:13] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_petition_svc_proto_init() }
func file_petition_svc_proto_init() {
	if File_petition_svc_proto != nil {
		return
	}
	file_petition_msg_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_petition_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_petition_svc_proto_goTypes,
		DependencyIndexes: file_petition_svc_proto_depIdxs,
	}.Build()
	File_petition_svc_proto = out.File
	file_petition_svc_proto_rawDesc = nil
	file_petition_svc_proto_goTypes = nil
	file_petition_svc_proto_depIdxs = nil
}
