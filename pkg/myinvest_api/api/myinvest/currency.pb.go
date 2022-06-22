// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: api/myinvest/currency.proto

package myinvest

import (
	_ "github.com/golang/protobuf/ptypes/empty"
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

type CurrencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortName []string `protobuf:"bytes,1,rep,name=ShortName,proto3" json:"ShortName,omitempty"`
}

func (x *CurrencyRequest) Reset() {
	*x = CurrencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_myinvest_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyRequest) ProtoMessage() {}

func (x *CurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_myinvest_currency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyRequest.ProtoReflect.Descriptor instead.
func (*CurrencyRequest) Descriptor() ([]byte, []int) {
	return file_api_myinvest_currency_proto_rawDescGZIP(), []int{0}
}

func (x *CurrencyRequest) GetShortName() []string {
	if x != nil {
		return x.ShortName
	}
	return nil
}

type Currency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortName string      `protobuf:"bytes,1,opt,name=ShortName,proto3" json:"ShortName,omitempty"`
	Name      string      `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Price     *MoneyValue `protobuf:"bytes,3,opt,name=Price,proto3" json:"Price,omitempty"`
	Lot       int32       `protobuf:"varint,4,opt,name=Lot,proto3" json:"Lot,omitempty"`
}

func (x *Currency) Reset() {
	*x = Currency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_myinvest_currency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Currency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Currency) ProtoMessage() {}

func (x *Currency) ProtoReflect() protoreflect.Message {
	mi := &file_api_myinvest_currency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Currency.ProtoReflect.Descriptor instead.
func (*Currency) Descriptor() ([]byte, []int) {
	return file_api_myinvest_currency_proto_rawDescGZIP(), []int{1}
}

func (x *Currency) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *Currency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Currency) GetPrice() *MoneyValue {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *Currency) GetLot() int32 {
	if x != nil {
		return x.Lot
	}
	return 0
}

type CurrencyList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currencies []*Currency `protobuf:"bytes,1,rep,name=Currencies,proto3" json:"Currencies,omitempty"`
}

func (x *CurrencyList) Reset() {
	*x = CurrencyList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_myinvest_currency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyList) ProtoMessage() {}

func (x *CurrencyList) ProtoReflect() protoreflect.Message {
	mi := &file_api_myinvest_currency_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyList.ProtoReflect.Descriptor instead.
func (*CurrencyList) Descriptor() ([]byte, []int) {
	return file_api_myinvest_currency_proto_rawDescGZIP(), []int{2}
}

func (x *CurrencyList) GetCurrencies() []*Currency {
	if x != nil {
		return x.Currencies
	}
	return nil
}

var File_api_myinvest_currency_proto protoreflect.FileDescriptor

var file_api_myinvest_currency_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d,
	0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x79,
	0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7e, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x4c, 0x6f, 0x74, 0x22, 0x46, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x79, 0x69, 0x6e,
	0x76, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x52, 0x0a, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x32, 0x5f, 0x0a,
	0x11, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x39,
	0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x61, 0x6c,
	0x79, 0x61, 0x6e, 0x6b, 0x69, 0x6c, 0x79, 0x61, 0x6e, 0x2f, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x78, 0x2d, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x62, 0x6f, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_myinvest_currency_proto_rawDescOnce sync.Once
	file_api_myinvest_currency_proto_rawDescData = file_api_myinvest_currency_proto_rawDesc
)

func file_api_myinvest_currency_proto_rawDescGZIP() []byte {
	file_api_myinvest_currency_proto_rawDescOnce.Do(func() {
		file_api_myinvest_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_myinvest_currency_proto_rawDescData)
	})
	return file_api_myinvest_currency_proto_rawDescData
}

var file_api_myinvest_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_myinvest_currency_proto_goTypes = []interface{}{
	(*CurrencyRequest)(nil), // 0: myinvest_api.CurrencyRequest
	(*Currency)(nil),        // 1: myinvest_api.Currency
	(*CurrencyList)(nil),    // 2: myinvest_api.CurrencyList
	(*MoneyValue)(nil),      // 3: myinvest_api.MoneyValue
}
var file_api_myinvest_currency_proto_depIdxs = []int32{
	3, // 0: myinvest_api.Currency.Price:type_name -> myinvest_api.MoneyValue
	1, // 1: myinvest_api.CurrencyList.Currencies:type_name -> myinvest_api.Currency
	0, // 2: myinvest_api.CurrenciesService.GetCurrencies:input_type -> myinvest_api.CurrencyRequest
	2, // 3: myinvest_api.CurrenciesService.GetCurrencies:output_type -> myinvest_api.CurrencyList
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_myinvest_currency_proto_init() }
func file_api_myinvest_currency_proto_init() {
	if File_api_myinvest_currency_proto != nil {
		return
	}
	file_api_myinvest_account_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_myinvest_currency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyRequest); i {
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
		file_api_myinvest_currency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Currency); i {
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
		file_api_myinvest_currency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyList); i {
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
			RawDescriptor: file_api_myinvest_currency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_myinvest_currency_proto_goTypes,
		DependencyIndexes: file_api_myinvest_currency_proto_depIdxs,
		MessageInfos:      file_api_myinvest_currency_proto_msgTypes,
	}.Build()
	File_api_myinvest_currency_proto = out.File
	file_api_myinvest_currency_proto_rawDesc = nil
	file_api_myinvest_currency_proto_goTypes = nil
	file_api_myinvest_currency_proto_depIdxs = nil
}
