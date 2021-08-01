// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protos/public/news/news.proto

package gen

import (
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

type GetNewsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags          []string               `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	LastTimestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_timestamp,json=lastTimestamp,proto3" json:"last_timestamp,omitempty"`
}

func (x *GetNewsReq) Reset() {
	*x = GetNewsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_public_news_news_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsReq) ProtoMessage() {}

func (x *GetNewsReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_public_news_news_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsReq.ProtoReflect.Descriptor instead.
func (*GetNewsReq) Descriptor() ([]byte, []int) {
	return file_protos_public_news_news_proto_rawDescGZIP(), []int{0}
}

func (x *GetNewsReq) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *GetNewsReq) GetLastTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.LastTimestamp
	}
	return nil
}

type GetNewsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles []*NewsArticle `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
}

func (x *GetNewsResp) Reset() {
	*x = GetNewsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_public_news_news_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsResp) ProtoMessage() {}

func (x *GetNewsResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_public_news_news_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsResp.ProtoReflect.Descriptor instead.
func (*GetNewsResp) Descriptor() ([]byte, []int) {
	return file_protos_public_news_news_proto_rawDescGZIP(), []int{1}
}

func (x *GetNewsResp) GetArticles() []*NewsArticle {
	if x != nil {
		return x.Articles
	}
	return nil
}

type CreateNewsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *NewsArticle `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *CreateNewsReq) Reset() {
	*x = CreateNewsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_public_news_news_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewsReq) ProtoMessage() {}

func (x *CreateNewsReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_public_news_news_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewsReq.ProtoReflect.Descriptor instead.
func (*CreateNewsReq) Descriptor() ([]byte, []int) {
	return file_protos_public_news_news_proto_rawDescGZIP(), []int{2}
}

func (x *CreateNewsReq) GetArticle() *NewsArticle {
	if x != nil {
		return x.Article
	}
	return nil
}

type CreateNewsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateNewsResp) Reset() {
	*x = CreateNewsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_public_news_news_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewsResp) ProtoMessage() {}

func (x *CreateNewsResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_public_news_news_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewsResp.ProtoReflect.Descriptor instead.
func (*CreateNewsResp) Descriptor() ([]byte, []int) {
	return file_protos_public_news_news_proto_rawDescGZIP(), []int{3}
}

func (x *CreateNewsResp) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NewsArticle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Tags      []string               `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *NewsArticle) Reset() {
	*x = NewsArticle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_public_news_news_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsArticle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsArticle) ProtoMessage() {}

func (x *NewsArticle) ProtoReflect() protoreflect.Message {
	mi := &file_protos_public_news_news_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsArticle.ProtoReflect.Descriptor instead.
func (*NewsArticle) Descriptor() ([]byte, []int) {
	return file_protos_public_news_news_proto_rawDescGZIP(), []int{4}
}

func (x *NewsArticle) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NewsArticle) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewsArticle) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *NewsArticle) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_protos_public_news_news_proto protoreflect.FileDescriptor

var file_protos_public_news_news_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x6e, 0x65, 0x77, 0x73, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6e, 0x65, 0x77,
	0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x63, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x48, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x6f, 0x67, 0x67, 0x6c,
	0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x73,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x22, 0x48, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x37, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x81, 0x01,
	0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x32, 0xb4, 0x01, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x1d, 0x2e, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x6e, 0x65, 0x77, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x56, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2f, 0x61, 0x61, 0x72, 0x6f, 0x6e, 0x62, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x6e,
	0x69, 0x63, 0x61, 0x6c, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_public_news_news_proto_rawDescOnce sync.Once
	file_protos_public_news_news_proto_rawDescData = file_protos_public_news_news_proto_rawDesc
)

func file_protos_public_news_news_proto_rawDescGZIP() []byte {
	file_protos_public_news_news_proto_rawDescOnce.Do(func() {
		file_protos_public_news_news_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_public_news_news_proto_rawDescData)
	})
	return file_protos_public_news_news_proto_rawDescData
}

var file_protos_public_news_news_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_public_news_news_proto_goTypes = []interface{}{
	(*GetNewsReq)(nil),            // 0: toggle.test.news.GetNewsReq
	(*GetNewsResp)(nil),           // 1: toggle.test.news.GetNewsResp
	(*CreateNewsReq)(nil),         // 2: toggle.test.news.CreateNewsReq
	(*CreateNewsResp)(nil),        // 3: toggle.test.news.CreateNewsResp
	(*NewsArticle)(nil),           // 4: toggle.test.news.NewsArticle
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_protos_public_news_news_proto_depIdxs = []int32{
	5, // 0: toggle.test.news.GetNewsReq.last_timestamp:type_name -> google.protobuf.Timestamp
	4, // 1: toggle.test.news.GetNewsResp.articles:type_name -> toggle.test.news.NewsArticle
	4, // 2: toggle.test.news.CreateNewsReq.article:type_name -> toggle.test.news.NewsArticle
	5, // 3: toggle.test.news.NewsArticle.timestamp:type_name -> google.protobuf.Timestamp
	0, // 4: toggle.test.news.NewsService.GetNewsArticle:input_type -> toggle.test.news.GetNewsReq
	2, // 5: toggle.test.news.NewsService.CreateNewsArticle:input_type -> toggle.test.news.CreateNewsReq
	1, // 6: toggle.test.news.NewsService.GetNewsArticle:output_type -> toggle.test.news.GetNewsResp
	3, // 7: toggle.test.news.NewsService.CreateNewsArticle:output_type -> toggle.test.news.CreateNewsResp
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_public_news_news_proto_init() }
func file_protos_public_news_news_proto_init() {
	if File_protos_public_news_news_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_public_news_news_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewsReq); i {
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
		file_protos_public_news_news_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewsResp); i {
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
		file_protos_public_news_news_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewsReq); i {
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
		file_protos_public_news_news_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewsResp); i {
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
		file_protos_public_news_news_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsArticle); i {
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
			RawDescriptor: file_protos_public_news_news_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_public_news_news_proto_goTypes,
		DependencyIndexes: file_protos_public_news_news_proto_depIdxs,
		MessageInfos:      file_protos_public_news_news_proto_msgTypes,
	}.Build()
	File_protos_public_news_news_proto = out.File
	file_protos_public_news_news_proto_rawDesc = nil
	file_protos_public_news_news_proto_goTypes = nil
	file_protos_public_news_news_proto_depIdxs = nil
}
