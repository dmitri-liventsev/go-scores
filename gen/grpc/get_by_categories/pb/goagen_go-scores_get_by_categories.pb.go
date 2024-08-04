// Code generated with goa v3.18.2, DO NOT EDIT.
//
// get by categories protocol buffer definition
//
// Command:
// $ goa gen go-scores/design

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: goagen_go-scores_get_by_categories.proto

package get_by_categoriespb

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

type GetAggregatedScoresRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Start date (YYYY-MM-DD)
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// End date (YYYY-MM-DD)
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *GetAggregatedScoresRequest) Reset() {
	*x = GetAggregatedScoresRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_go_scores_get_by_categories_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAggregatedScoresRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAggregatedScoresRequest) ProtoMessage() {}

func (x *GetAggregatedScoresRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_go_scores_get_by_categories_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAggregatedScoresRequest.ProtoReflect.Descriptor instead.
func (*GetAggregatedScoresRequest) Descriptor() ([]byte, []int) {
	return file_goagen_go_scores_get_by_categories_proto_rawDescGZIP(), []int{0}
}

func (x *GetAggregatedScoresRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *GetAggregatedScoresRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type GetAggregatedScoresResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *CategoryMeta    `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Data []*CategoryScore `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAggregatedScoresResponse) Reset() {
	*x = GetAggregatedScoresResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_go_scores_get_by_categories_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAggregatedScoresResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAggregatedScoresResponse) ProtoMessage() {}

func (x *GetAggregatedScoresResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_go_scores_get_by_categories_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAggregatedScoresResponse.ProtoReflect.Descriptor instead.
func (*GetAggregatedScoresResponse) Descriptor() ([]byte, []int) {
	return file_goagen_go_scores_get_by_categories_proto_rawDescGZIP(), []int{1}
}

func (x *GetAggregatedScoresResponse) GetMeta() *CategoryMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *GetAggregatedScoresResponse) GetData() []*CategoryScore {
	if x != nil {
		return x.Data
	}
	return nil
}

type CategoryMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Periods    []*Period   `protobuf:"bytes,1,rep,name=periods,proto3" json:"periods,omitempty"`
	Categories []*Category `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *CategoryMeta) Reset() {
	*x = CategoryMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_go_scores_get_by_categories_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryMeta) ProtoMessage() {}

func (x *CategoryMeta) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_go_scores_get_by_categories_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryMeta.ProtoReflect.Descriptor instead.
func (*CategoryMeta) Descriptor() ([]byte, []int) {
	return file_goagen_go_scores_get_by_categories_proto_rawDescGZIP(), []int{2}
}

func (x *CategoryMeta) GetPeriods() []*Period {
	if x != nil {
		return x.Periods
	}
	return nil
}

func (x *CategoryMeta) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

type Period struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    *int32 `protobuf:"zigzag32,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Start *int64 `protobuf:"zigzag64,3,opt,name=start,proto3,oneof" json:"start,omitempty"`
	End   *int64 `protobuf:"zigzag64,4,opt,name=end,proto3,oneof" json:"end,omitempty"`
}

func (x *Period) Reset() {
	*x = Period{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_go_scores_get_by_categories_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Period) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Period) ProtoMessage() {}

func (x *Period) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_go_scores_get_by_categories_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Period.ProtoReflect.Descriptor instead.
func (*Period) Descriptor() ([]byte, []int) {
	return file_goagen_go_scores_get_by_categories_proto_rawDescGZIP(), []int{3}
}

func (x *Period) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Period) GetStart() int64 {
	if x != nil && x.Start != nil {
		return *x.Start
	}
	return 0
}

func (x *Period) GetEnd() int64 {
	if x != nil && x.End != nil {
		return *x.End
	}
	return 0
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *int32  `protobuf:"zigzag32,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_go_scores_get_by_categories_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_go_scores_get_by_categories_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_goagen_go_scores_get_by_categories_proto_rawDescGZIP(), []int{4}
}

func (x *Category) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Category) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type CategoryScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId   *int32         `protobuf:"zigzag32,1,opt,name=category_id,json=categoryId,proto3,oneof" json:"category_id,omitempty"`
	NumOfReviews *int32         `protobuf:"zigzag32,2,opt,name=num_of_reviews,json=numOfReviews,proto3,oneof" json:"num_of_reviews,omitempty"`
	Periods      []*ScorePeriod `protobuf:"bytes,3,rep,name=periods,proto3" json:"periods,omitempty"`
	TotalScore   *float32       `protobuf:"fixed32,4,opt,name=total_score,json=totalScore,proto3,oneof" json:"total_score,omitempty"`
}

func (x *CategoryScore) Reset() {
	*x = CategoryScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_go_scores_get_by_categories_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryScore) ProtoMessage() {}

func (x *CategoryScore) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_go_scores_get_by_categories_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryScore.ProtoReflect.Descriptor instead.
func (*CategoryScore) Descriptor() ([]byte, []int) {
	return file_goagen_go_scores_get_by_categories_proto_rawDescGZIP(), []int{5}
}

func (x *CategoryScore) GetCategoryId() int32 {
	if x != nil && x.CategoryId != nil {
		return *x.CategoryId
	}
	return 0
}

func (x *CategoryScore) GetNumOfReviews() int32 {
	if x != nil && x.NumOfReviews != nil {
		return *x.NumOfReviews
	}
	return 0
}

func (x *CategoryScore) GetPeriods() []*ScorePeriod {
	if x != nil {
		return x.Periods
	}
	return nil
}

func (x *CategoryScore) GetTotalScore() float32 {
	if x != nil && x.TotalScore != nil {
		return *x.TotalScore
	}
	return 0
}

type ScorePeriod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    *int32   `protobuf:"zigzag32,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Score *float32 `protobuf:"fixed32,2,opt,name=score,proto3,oneof" json:"score,omitempty"`
}

func (x *ScorePeriod) Reset() {
	*x = ScorePeriod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goagen_go_scores_get_by_categories_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScorePeriod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScorePeriod) ProtoMessage() {}

func (x *ScorePeriod) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_go_scores_get_by_categories_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScorePeriod.ProtoReflect.Descriptor instead.
func (*ScorePeriod) Descriptor() ([]byte, []int) {
	return file_goagen_go_scores_get_by_categories_proto_rawDescGZIP(), []int{6}
}

func (x *ScorePeriod) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *ScorePeriod) GetScore() float32 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

var File_goagen_go_scores_get_by_categories_proto protoreflect.FileDescriptor

var file_goagen_go_scores_get_by_categories_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x61, 0x67, 0x65, 0x6e, 0x5f, 0x67, 0x6f, 0x2d, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x65, 0x74, 0x5f,
	0x62, 0x79, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x40, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x22,
	0x88, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x33, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x80, 0x01, 0x0a, 0x0c, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x07, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x2e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x07, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73,
	0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x68, 0x0a,
	0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x11, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x12, 0x48, 0x01, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x12, 0x48, 0x02, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x65, 0x6e, 0x64, 0x22, 0x48, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0xf3, 0x01, 0x0a, 0x0d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x6e, 0x75, 0x6d,
	0x5f, 0x6f, 0x66, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x11, 0x48, 0x01, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x4f, 0x66, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x07, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x12, 0x24,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x48, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x6f, 0x66, 0x5f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x4e, 0x0a, 0x0b, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x11, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x48, 0x01, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x32, 0x87, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x74, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x12, 0x2d, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x16, 0x5a, 0x14, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_goagen_go_scores_get_by_categories_proto_rawDescOnce sync.Once
	file_goagen_go_scores_get_by_categories_proto_rawDescData = file_goagen_go_scores_get_by_categories_proto_rawDesc
)

func file_goagen_go_scores_get_by_categories_proto_rawDescGZIP() []byte {
	file_goagen_go_scores_get_by_categories_proto_rawDescOnce.Do(func() {
		file_goagen_go_scores_get_by_categories_proto_rawDescData = protoimpl.X.CompressGZIP(file_goagen_go_scores_get_by_categories_proto_rawDescData)
	})
	return file_goagen_go_scores_get_by_categories_proto_rawDescData
}

var file_goagen_go_scores_get_by_categories_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_goagen_go_scores_get_by_categories_proto_goTypes = []any{
	(*GetAggregatedScoresRequest)(nil),  // 0: get_by_categories.GetAggregatedScoresRequest
	(*GetAggregatedScoresResponse)(nil), // 1: get_by_categories.GetAggregatedScoresResponse
	(*CategoryMeta)(nil),                // 2: get_by_categories.CategoryMeta
	(*Period)(nil),                      // 3: get_by_categories.Period
	(*Category)(nil),                    // 4: get_by_categories.Category
	(*CategoryScore)(nil),               // 5: get_by_categories.CategoryScore
	(*ScorePeriod)(nil),                 // 6: get_by_categories.ScorePeriod
}
var file_goagen_go_scores_get_by_categories_proto_depIdxs = []int32{
	2, // 0: get_by_categories.GetAggregatedScoresResponse.meta:type_name -> get_by_categories.CategoryMeta
	5, // 1: get_by_categories.GetAggregatedScoresResponse.data:type_name -> get_by_categories.CategoryScore
	3, // 2: get_by_categories.CategoryMeta.periods:type_name -> get_by_categories.Period
	4, // 3: get_by_categories.CategoryMeta.categories:type_name -> get_by_categories.Category
	6, // 4: get_by_categories.CategoryScore.periods:type_name -> get_by_categories.ScorePeriod
	0, // 5: get_by_categories.GetByCategories.GetAggregatedScores:input_type -> get_by_categories.GetAggregatedScoresRequest
	1, // 6: get_by_categories.GetByCategories.GetAggregatedScores:output_type -> get_by_categories.GetAggregatedScoresResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_goagen_go_scores_get_by_categories_proto_init() }
func file_goagen_go_scores_get_by_categories_proto_init() {
	if File_goagen_go_scores_get_by_categories_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goagen_go_scores_get_by_categories_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetAggregatedScoresRequest); i {
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
		file_goagen_go_scores_get_by_categories_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAggregatedScoresResponse); i {
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
		file_goagen_go_scores_get_by_categories_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CategoryMeta); i {
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
		file_goagen_go_scores_get_by_categories_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Period); i {
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
		file_goagen_go_scores_get_by_categories_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Category); i {
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
		file_goagen_go_scores_get_by_categories_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CategoryScore); i {
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
		file_goagen_go_scores_get_by_categories_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ScorePeriod); i {
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
	file_goagen_go_scores_get_by_categories_proto_msgTypes[3].OneofWrappers = []any{}
	file_goagen_go_scores_get_by_categories_proto_msgTypes[4].OneofWrappers = []any{}
	file_goagen_go_scores_get_by_categories_proto_msgTypes[5].OneofWrappers = []any{}
	file_goagen_go_scores_get_by_categories_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goagen_go_scores_get_by_categories_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goagen_go_scores_get_by_categories_proto_goTypes,
		DependencyIndexes: file_goagen_go_scores_get_by_categories_proto_depIdxs,
		MessageInfos:      file_goagen_go_scores_get_by_categories_proto_msgTypes,
	}.Build()
	File_goagen_go_scores_get_by_categories_proto = out.File
	file_goagen_go_scores_get_by_categories_proto_rawDesc = nil
	file_goagen_go_scores_get_by_categories_proto_goTypes = nil
	file_goagen_go_scores_get_by_categories_proto_depIdxs = nil
}