// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: api/myinvestapi/currency.proto

package myinvestapi

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

type CurrencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortName []string `protobuf:"bytes,1,rep,name=ShortName,proto3" json:"ShortName,omitempty"`
}

func (x *CurrencyRequest) Reset() {
	*x = CurrencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_myinvestapi_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyRequest) ProtoMessage() {}

func (x *CurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_myinvestapi_currency_proto_msgTypes[0]
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
	return file_api_myinvestapi_currency_proto_rawDescGZIP(), []int{0}
}

func (x *CurrencyRequest) GetShortName() []string {
	if x != nil {
		return x.ShortName
	}
	return nil
}

type Quotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Units int64 `protobuf:"varint,1,opt,name=Units,proto3" json:"Units,omitempty"`
	Nano  int32 `protobuf:"varint,2,opt,name=Nano,proto3" json:"Nano,omitempty"`
}

func (x *Quotation) Reset() {
	*x = Quotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_myinvestapi_currency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quotation) ProtoMessage() {}

func (x *Quotation) ProtoReflect() protoreflect.Message {
	mi := &file_api_myinvestapi_currency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quotation.ProtoReflect.Descriptor instead.
func (*Quotation) Descriptor() ([]byte, []int) {
	return file_api_myinvestapi_currency_proto_rawDescGZIP(), []int{1}
}

func (x *Quotation) GetUnits() int64 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *Quotation) GetNano() int32 {
	if x != nil {
		return x.Nano
	}
	return 0
}

type Currency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortName string     `protobuf:"bytes,1,opt,name=ShortName,proto3" json:"ShortName,omitempty"`
	Name      string     `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Price     *Quotation `protobuf:"bytes,3,opt,name=Price,proto3" json:"Price,omitempty"`
	Lot       int32      `protobuf:"varint,4,opt,name=Lot,proto3" json:"Lot,omitempty"`
}

func (x *Currency) Reset() {
	*x = Currency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_myinvestapi_currency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Currency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Currency) ProtoMessage() {}

func (x *Currency) ProtoReflect() protoreflect.Message {
	mi := &file_api_myinvestapi_currency_proto_msgTypes[2]
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
	return file_api_myinvestapi_currency_proto_rawDescGZIP(), []int{2}
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

func (x *Currency) GetPrice() *Quotation {
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
		mi := &file_api_myinvestapi_currency_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyList) ProtoMessage() {}

func (x *CurrencyList) ProtoReflect() protoreflect.Message {
	mi := &file_api_myinvestapi_currency_proto_msgTypes[3]
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
	return file_api_myinvestapi_currency_proto_rawDescGZIP(), []int{3}
}

func (x *CurrencyList) GetCurrencies() []*Currency {
	if x != nil {
		return x.Currencies
	}
	return nil
}

var File_api_myinvestapi_currency_proto protoreflect.FileDescriptor

var file_api_myinvestapi_currency_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x22, 0x2f, 0x0a,
	0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x35,
	0x0a, 0x09, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x55,
	0x6e, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x55, 0x6e, 0x69, 0x74,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x4e, 0x61, 0x6e, 0x6f, 0x22, 0x7c, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69,
	0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x4c, 0x6f, 0x74, 0x22, 0x45, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x79, 0x69, 0x6e, 0x76, 0x65,
	0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0a,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x61, 0x6c, 0x79, 0x61, 0x6e, 0x6b,
	0x69, 0x6c, 0x79, 0x61, 0x6e, 0x2f, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2d, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x2d, 0x62, 0x6f, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x79, 0x69,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_myinvestapi_currency_proto_rawDescOnce sync.Once
	file_api_myinvestapi_currency_proto_rawDescData = file_api_myinvestapi_currency_proto_rawDesc
)

func file_api_myinvestapi_currency_proto_rawDescGZIP() []byte {
	file_api_myinvestapi_currency_proto_rawDescOnce.Do(func() {
		file_api_myinvestapi_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_myinvestapi_currency_proto_rawDescData)
	})
	return file_api_myinvestapi_currency_proto_rawDescData
}

var file_api_myinvestapi_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_myinvestapi_currency_proto_goTypes = []interface{}{
	(*CurrencyRequest)(nil), // 0: myinvestapi.CurrencyRequest
	(*Quotation)(nil),       // 1: myinvestapi.Quotation
	(*Currency)(nil),        // 2: myinvestapi.Currency
	(*CurrencyList)(nil),    // 3: myinvestapi.CurrencyList
}
var file_api_myinvestapi_currency_proto_depIdxs = []int32{
	1, // 0: myinvestapi.Currency.Price:type_name -> myinvestapi.Quotation
	2, // 1: myinvestapi.CurrencyList.Currencies:type_name -> myinvestapi.Currency
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_myinvestapi_currency_proto_init() }
func file_api_myinvestapi_currency_proto_init() {
	if File_api_myinvestapi_currency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_myinvestapi_currency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_myinvestapi_currency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quotation); i {
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
		file_api_myinvestapi_currency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_myinvestapi_currency_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_api_myinvestapi_currency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_myinvestapi_currency_proto_goTypes,
		DependencyIndexes: file_api_myinvestapi_currency_proto_depIdxs,
		MessageInfos:      file_api_myinvestapi_currency_proto_msgTypes,
	}.Build()
	File_api_myinvestapi_currency_proto = out.File
	file_api_myinvestapi_currency_proto_rawDesc = nil
	file_api_myinvestapi_currency_proto_goTypes = nil
	file_api_myinvestapi_currency_proto_depIdxs = nil
}
