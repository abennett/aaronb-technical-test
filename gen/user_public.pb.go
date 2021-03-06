// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protos/public/user/user_public.proto

package gen

import (
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

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_public_user_user_public_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_public_user_user_public_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_protos_public_user_user_public_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_public_user_user_public_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_public_user_user_public_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_protos_public_user_user_public_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CreateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserReq) Reset() {
	*x = CreateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_public_user_user_public_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserReq) ProtoMessage() {}

func (x *CreateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_public_user_user_public_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserReq.ProtoReflect.Descriptor instead.
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return file_protos_public_user_user_public_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateUserResp) Reset() {
	*x = CreateUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_public_user_user_public_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResp) ProtoMessage() {}

func (x *CreateUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_public_user_user_public_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResp.ProtoReflect.Descriptor instead.
func (*CreateUserResp) Descriptor() ([]byte, []int) {
	return file_protos_public_user_user_public_proto_rawDescGZIP(), []int{3}
}

func (x *CreateUserResp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type AddUserTagReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *AddUserTagReq) Reset() {
	*x = AddUserTagReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_public_user_user_public_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserTagReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserTagReq) ProtoMessage() {}

func (x *AddUserTagReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_public_user_user_public_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserTagReq.ProtoReflect.Descriptor instead.
func (*AddUserTagReq) Descriptor() ([]byte, []int) {
	return file_protos_public_user_user_public_proto_rawDescGZIP(), []int{4}
}

func (x *AddUserTagReq) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type RemoveUserTagReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *RemoveUserTagReq) Reset() {
	*x = RemoveUserTagReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_public_user_user_public_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserTagReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserTagReq) ProtoMessage() {}

func (x *RemoveUserTagReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_public_user_user_public_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserTagReq.ProtoReflect.Descriptor instead.
func (*RemoveUserTagReq) Descriptor() ([]byte, []int) {
	return file_protos_public_user_user_public_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveUserTagReq) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

var File_protos_public_user_user_public_proto protoreflect.FileDescriptor

var file_protos_public_user_user_public_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x21, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x47, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x29, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x24, 0x0a, 0x10, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67,
	0x32, 0xe4, 0x02, 0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x21, 0x2e, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x22, 0x2e, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x5d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x74,
	0x6f, 0x67, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4c, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x61, 0x67, 0x12, 0x26, 0x2e, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x52, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x61, 0x67, 0x12, 0x29, 0x2e, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2f, 0x61, 0x61, 0x72, 0x6f, 0x6e, 0x62, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69,
	0x63, 0x61, 0x6c, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_public_user_user_public_proto_rawDescOnce sync.Once
	file_protos_public_user_user_public_proto_rawDescData = file_protos_public_user_user_public_proto_rawDesc
)

func file_protos_public_user_user_public_proto_rawDescGZIP() []byte {
	file_protos_public_user_user_public_proto_rawDescOnce.Do(func() {
		file_protos_public_user_user_public_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_public_user_user_public_proto_rawDescData)
	})
	return file_protos_public_user_user_public_proto_rawDescData
}

var file_protos_public_user_user_public_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_public_user_user_public_proto_goTypes = []interface{}{
	(*LoginReq)(nil),         // 0: toggle.test.public.user.LoginReq
	(*LoginResp)(nil),        // 1: toggle.test.public.user.LoginResp
	(*CreateUserReq)(nil),    // 2: toggle.test.public.user.CreateUserReq
	(*CreateUserResp)(nil),   // 3: toggle.test.public.user.CreateUserResp
	(*AddUserTagReq)(nil),    // 4: toggle.test.public.user.AddUserTagReq
	(*RemoveUserTagReq)(nil), // 5: toggle.test.public.user.RemoveUserTagReq
	(*emptypb.Empty)(nil),    // 6: google.protobuf.Empty
}
var file_protos_public_user_user_public_proto_depIdxs = []int32{
	0, // 0: toggle.test.public.user.PublicUserService.Login:input_type -> toggle.test.public.user.LoginReq
	2, // 1: toggle.test.public.user.PublicUserService.CreateUser:input_type -> toggle.test.public.user.CreateUserReq
	4, // 2: toggle.test.public.user.PublicUserService.AddUserTag:input_type -> toggle.test.public.user.AddUserTagReq
	5, // 3: toggle.test.public.user.PublicUserService.RemoveUserTag:input_type -> toggle.test.public.user.RemoveUserTagReq
	1, // 4: toggle.test.public.user.PublicUserService.Login:output_type -> toggle.test.public.user.LoginResp
	3, // 5: toggle.test.public.user.PublicUserService.CreateUser:output_type -> toggle.test.public.user.CreateUserResp
	6, // 6: toggle.test.public.user.PublicUserService.AddUserTag:output_type -> google.protobuf.Empty
	6, // 7: toggle.test.public.user.PublicUserService.RemoveUserTag:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_public_user_user_public_proto_init() }
func file_protos_public_user_user_public_proto_init() {
	if File_protos_public_user_user_public_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_public_user_user_public_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_protos_public_user_user_public_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_protos_public_user_user_public_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserReq); i {
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
		file_protos_public_user_user_public_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResp); i {
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
		file_protos_public_user_user_public_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserTagReq); i {
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
		file_protos_public_user_user_public_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserTagReq); i {
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
			RawDescriptor: file_protos_public_user_user_public_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_public_user_user_public_proto_goTypes,
		DependencyIndexes: file_protos_public_user_user_public_proto_depIdxs,
		MessageInfos:      file_protos_public_user_user_public_proto_msgTypes,
	}.Build()
	File_protos_public_user_user_public_proto = out.File
	file_protos_public_user_user_public_proto_rawDesc = nil
	file_protos_public_user_user_public_proto_goTypes = nil
	file_protos_public_user_user_public_proto_depIdxs = nil
}
