// Code generated with goa v3.18.2, DO NOT EDIT.
//
// get overall protocol buffer definition
//
// Command:
// $ goa gen go-scores/design

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: goagen_go-scores_get_overall.proto

package get_overallpb

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

type GetOverallScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Start date in YYYY-MM-DD format
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// End date in YYYY-MM-DD format
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *GetOverallScoreRequest) Reset() {
	*x = GetOverallScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_go_scores_get_overall_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOverallScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOverallScoreRequest) ProtoMessage() {}

func (x *GetOverallScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_go_scores_get_overall_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOverallScoreRequest.ProtoReflect.Descriptor instead.
func (*GetOverallScoreRequest) Descriptor() ([]byte, []int) {
	return file_goagen_go_scores_get_overall_proto_rawDescGZIP(), []int{0}
}

func (x *GetOverallScoreRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *GetOverallScoreRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type GetOverallScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field float32 `protobuf:"fixed32,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *GetOverallScoreResponse) Reset() {
	*x = GetOverallScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_go_scores_get_overall_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOverallScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOverallScoreResponse) ProtoMessage() {}

func (x *GetOverallScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_go_scores_get_overall_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOverallScoreResponse.ProtoReflect.Descriptor instead.
func (*GetOverallScoreResponse) Descriptor() ([]byte, []int) {
	return file_goagen_go_scores_get_overall_proto_rawDescGZIP(), []int{1}
}

func (x *GetOverallScoreResponse) GetField() float32 {
	if x != nil {
		return x.Field
	}
	return 0
}

var File_goagen_go_scores_get_overall_proto protoreflect.FileDescriptor

var file_goagen_go_scores_get_overall_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x6f, 0x61, 0x67, 0x65, 0x6e, 0x5f, 0x67, 0x6f, 0x2d, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c,
	0x6c, 0x22, 0x3c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x22,
	0x2f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x6a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x12, 0x5c,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x23, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x76, 0x65,
	0x72, 0x61, 0x6c, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e,
	0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goagen_go_scores_get_overall_proto_rawDescOnce sync.Once
	file_goagen_go_scores_get_overall_proto_rawDescData = file_goagen_go_scores_get_overall_proto_rawDesc
)

func file_goagen_go_scores_get_overall_proto_rawDescGZIP() []byte {
	file_goagen_go_scores_get_overall_proto_rawDescOnce.Do(func() {
		file_goagen_go_scores_get_overall_proto_rawDescData = protoimpl.X.CompressGZIP(file_goagen_go_scores_get_overall_proto_rawDescData)
	})
	return file_goagen_go_scores_get_overall_proto_rawDescData
}

var file_goagen_go_scores_get_overall_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_goagen_go_scores_get_overall_proto_goTypes = []any{
	(*GetOverallScoreRequest)(nil),  // 0: get_overall.GetOverallScoreRequest
	(*GetOverallScoreResponse)(nil), // 1: get_overall.GetOverallScoreResponse
}
var file_goagen_go_scores_get_overall_proto_depIdxs = []int32{
	0, // 0: get_overall.GetOverall.GetOverallScore:input_type -> get_overall.GetOverallScoreRequest
	1, // 1: get_overall.GetOverall.GetOverallScore:output_type -> get_overall.GetOverallScoreResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goagen_go_scores_get_overall_proto_init() }
func file_goagen_go_scores_get_overall_proto_init() {
	if File_goagen_go_scores_get_overall_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goagen_go_scores_get_overall_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetOverallScoreRequest); i {
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
		file_goagen_go_scores_get_overall_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetOverallScoreResponse); i {
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
			RawDescriptor: file_goagen_go_scores_get_overall_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goagen_go_scores_get_overall_proto_goTypes,
		DependencyIndexes: file_goagen_go_scores_get_overall_proto_depIdxs,
		MessageInfos:      file_goagen_go_scores_get_overall_proto_msgTypes,
	}.Build()
	File_goagen_go_scores_get_overall_proto = out.File
	file_goagen_go_scores_get_overall_proto_rawDesc = nil
	file_goagen_go_scores_get_overall_proto_goTypes = nil
	file_goagen_go_scores_get_overall_proto_depIdxs = nil
}
