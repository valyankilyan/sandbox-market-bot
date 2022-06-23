// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: api/server/server.proto

package server_api

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_server_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_api_server_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_api_server_server_proto_rawDescGZIP(), []int{0}
}

func (x *Id) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TgId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TgId int64 `protobuf:"varint,1,opt,name=TgId,proto3" json:"TgId,omitempty"`
}

func (x *TgId) Reset() {
	*x = TgId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_server_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TgId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TgId) ProtoMessage() {}

func (x *TgId) ProtoReflect() protoreflect.Message {
	mi := &file_api_server_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TgId.ProtoReflect.Descriptor instead.
func (*TgId) Descriptor() ([]byte, []int) {
	return file_api_server_server_proto_rawDescGZIP(), []int{1}
}

func (x *TgId) GetTgId() int64 {
	if x != nil {
		return x.TgId
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TgUserName   string `protobuf:"bytes,3,opt,name=tgUserName,proto3" json:"tgUserName,omitempty"`
	TgId         int64  `protobuf:"varint,4,opt,name=tgId,proto3" json:"tgId,omitempty"`
	TinkoffToken string `protobuf:"bytes,5,opt,name=tinkoffToken,proto3" json:"tinkoffToken,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_server_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_server_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_server_server_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetTgUserName() string {
	if x != nil {
		return x.TgUserName
	}
	return ""
}

func (x *User) GetTgId() int64 {
	if x != nil {
		return x.TgId
	}
	return 0
}

func (x *User) GetTinkoffToken() string {
	if x != nil {
		return x.TinkoffToken
	}
	return ""
}

var File_api_server_server_proto protoreflect.FileDescriptor

var file_api_server_server_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x04, 0x54, 0x67, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x54, 0x67, 0x49, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x67, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x67, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x67, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x74, 0x67, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69, 0x6e,
	0x6b, 0x6f, 0x66, 0x66, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xe5, 0x01, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x64, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x08, 0x52,
	0x65, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x67, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x54, 0x67, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x76, 0x61, 0x6c, 0x79, 0x61, 0x6e, 0x6b, 0x69, 0x6c, 0x79, 0x61, 0x6e, 0x2f, 0x73, 0x61, 0x6e,
	0x64, 0x62, 0x6f, 0x78, 0x2d, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x62, 0x6f, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_server_server_proto_rawDescOnce sync.Once
	file_api_server_server_proto_rawDescData = file_api_server_server_proto_rawDesc
)

func file_api_server_server_proto_rawDescGZIP() []byte {
	file_api_server_server_proto_rawDescOnce.Do(func() {
		file_api_server_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_server_server_proto_rawDescData)
	})
	return file_api_server_server_proto_rawDescData
}

var file_api_server_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_server_server_proto_goTypes = []interface{}{
	(*Id)(nil),          // 0: server_api.Id
	(*TgId)(nil),        // 1: server_api.TgId
	(*User)(nil),        // 2: server_api.User
	(*empty.Empty)(nil), // 3: google.protobuf.Empty
}
var file_api_server_server_proto_depIdxs = []int32{
	2, // 0: server_api.UserService.CreateUser:input_type -> server_api.User
	1, // 1: server_api.UserService.ReadUser:input_type -> server_api.TgId
	2, // 2: server_api.UserService.UpdateUser:input_type -> server_api.User
	1, // 3: server_api.UserService.DeleteUser:input_type -> server_api.TgId
	0, // 4: server_api.UserService.CreateUser:output_type -> server_api.Id
	2, // 5: server_api.UserService.ReadUser:output_type -> server_api.User
	3, // 6: server_api.UserService.UpdateUser:output_type -> google.protobuf.Empty
	3, // 7: server_api.UserService.DeleteUser:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_server_server_proto_init() }
func file_api_server_server_proto_init() {
	if File_api_server_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_server_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_api_server_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TgId); i {
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
		file_api_server_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_api_server_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_server_server_proto_goTypes,
		DependencyIndexes: file_api_server_server_proto_depIdxs,
		MessageInfos:      file_api_server_server_proto_msgTypes,
	}.Build()
	File_api_server_server_proto = out.File
	file_api_server_server_proto_rawDesc = nil
	file_api_server_server_proto_goTypes = nil
	file_api_server_server_proto_depIdxs = nil
}
