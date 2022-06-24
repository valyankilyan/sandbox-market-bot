// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: api/myinvestapi/account.proto

package myinvestapi

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

type TinkoffToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *TinkoffToken) Reset() {
	*x = TinkoffToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_myinvestapi_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TinkoffToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TinkoffToken) ProtoMessage() {}

func (x *TinkoffToken) ProtoReflect() protoreflect.Message {
	mi := &file_api_myinvestapi_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TinkoffToken.ProtoReflect.Descriptor instead.
func (*TinkoffToken) Descriptor() ([]byte, []int) {
	return file_api_myinvestapi_account_proto_rawDescGZIP(), []int{0}
}

func (x *TinkoffToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type PayinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    *TinkoffToken `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Quantity *Quantity     `protobuf:"bytes,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
}

func (x *PayinRequest) Reset() {
	*x = PayinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_myinvestapi_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayinRequest) ProtoMessage() {}

func (x *PayinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_myinvestapi_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayinRequest.ProtoReflect.Descriptor instead.
func (*PayinRequest) Descriptor() ([]byte, []int) {
	return file_api_myinvestapi_account_proto_rawDescGZIP(), []int{1}
}

func (x *PayinRequest) GetToken() *TinkoffToken {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *PayinRequest) GetQuantity() *Quantity {
	if x != nil {
		return x.Quantity
	}
	return nil
}

var File_api_myinvestapi_account_proto protoreflect.FileDescriptor

var file_api_myinvestapi_account_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0c, 0x54, 0x69, 0x6e,
	0x6b, 0x6f, 0x66, 0x66, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x72, 0x0a, 0x0c, 0x50, 0x61, 0x79, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2f, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x69, 0x6e,
	0x6b, 0x6f, 0x66, 0x66, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x31, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69,
	0x2e, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x32, 0xde, 0x01, 0x0a, 0x0f, 0x4d, 0x79, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x50, 0x61, 0x79, 0x49, 0x6e,
	0x12, 0x19, 0x2e, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x50,
	0x61, 0x79, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x45, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x6d, 0x79, 0x69, 0x6e,
	0x76, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x76, 0x61, 0x6c, 0x79, 0x61, 0x6e, 0x6b, 0x69, 0x6c, 0x79, 0x61, 0x6e, 0x2f,
	0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2d, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x62,
	0x6f, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_myinvestapi_account_proto_rawDescOnce sync.Once
	file_api_myinvestapi_account_proto_rawDescData = file_api_myinvestapi_account_proto_rawDesc
)

func file_api_myinvestapi_account_proto_rawDescGZIP() []byte {
	file_api_myinvestapi_account_proto_rawDescOnce.Do(func() {
		file_api_myinvestapi_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_myinvestapi_account_proto_rawDescData)
	})
	return file_api_myinvestapi_account_proto_rawDescData
}

var file_api_myinvestapi_account_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_myinvestapi_account_proto_goTypes = []interface{}{
	(*TinkoffToken)(nil),    // 0: myinvestapi.TinkoffToken
	(*PayinRequest)(nil),    // 1: myinvestapi.PayinRequest
	(*Quantity)(nil),        // 2: myinvestapi.Quantity
	(*CurrencyRequest)(nil), // 3: myinvestapi.CurrencyRequest
	(*empty.Empty)(nil),     // 4: google.protobuf.Empty
	(*CurrencyList)(nil),    // 5: myinvestapi.CurrencyList
}
var file_api_myinvestapi_account_proto_depIdxs = []int32{
	0, // 0: myinvestapi.PayinRequest.Token:type_name -> myinvestapi.TinkoffToken
	2, // 1: myinvestapi.PayinRequest.Quantity:type_name -> myinvestapi.Quantity
	1, // 2: myinvestapi.MyInvestService.PayIn:input_type -> myinvestapi.PayinRequest
	3, // 3: myinvestapi.MyInvestService.GetCurrencies:input_type -> myinvestapi.CurrencyRequest
	4, // 4: myinvestapi.MyInvestService.GetAllCurrencies:input_type -> google.protobuf.Empty
	4, // 5: myinvestapi.MyInvestService.PayIn:output_type -> google.protobuf.Empty
	5, // 6: myinvestapi.MyInvestService.GetCurrencies:output_type -> myinvestapi.CurrencyList
	5, // 7: myinvestapi.MyInvestService.GetAllCurrencies:output_type -> myinvestapi.CurrencyList
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_myinvestapi_account_proto_init() }
func file_api_myinvestapi_account_proto_init() {
	if File_api_myinvestapi_account_proto != nil {
		return
	}
	file_api_myinvestapi_currency_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_myinvestapi_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TinkoffToken); i {
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
		file_api_myinvestapi_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayinRequest); i {
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
			RawDescriptor: file_api_myinvestapi_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_myinvestapi_account_proto_goTypes,
		DependencyIndexes: file_api_myinvestapi_account_proto_depIdxs,
		MessageInfos:      file_api_myinvestapi_account_proto_msgTypes,
	}.Build()
	File_api_myinvestapi_account_proto = out.File
	file_api_myinvestapi_account_proto_rawDesc = nil
	file_api_myinvestapi_account_proto_goTypes = nil
	file_api_myinvestapi_account_proto_depIdxs = nil
}
