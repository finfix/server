// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: enum.proto

package pbEnum

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

type AccountType int32

const (
	AccountType_Regular  AccountType = 0
	AccountType_Expense  AccountType = 1
	AccountType_Debt     AccountType = 2
	AccountType_Earnings AccountType = 3
)

// Enum value maps for AccountType.
var (
	AccountType_name = map[int32]string{
		0: "Regular",
		1: "Expense",
		2: "Debt",
		3: "Earnings",
	}
	AccountType_value = map[string]int32{
		"Regular":  0,
		"Expense":  1,
		"Debt":     2,
		"Earnings": 3,
	}
)

func (x AccountType) Enum() *AccountType {
	p := new(AccountType)
	*p = x
	return p
}

func (x AccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[0].Descriptor()
}

func (AccountType) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[0]
}

func (x AccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountType.Descriptor instead.
func (AccountType) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{0}
}

type TransactionType int32

const (
	TransactionType_Transfer    TransactionType = 0
	TransactionType_Consumption TransactionType = 1
	TransactionType_Balancing   TransactionType = 2
	TransactionType_Income      TransactionType = 3
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "Transfer",
		1: "Consumption",
		2: "Balancing",
		3: "Income",
	}
	TransactionType_value = map[string]int32{
		"Transfer":    0,
		"Consumption": 1,
		"Balancing":   2,
		"Income":      3,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[1].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[1]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{1}
}

var File_enum_proto protoreflect.FileDescriptor

var file_enum_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2a, 0x3f, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x44, 0x65, 0x62, 0x74, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x73, 0x10, 0x03, 0x2a, 0x4b, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x10,
	0x03, 0x42, 0x17, 0x5a, 0x15, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x45, 0x6e, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_enum_proto_rawDescOnce sync.Once
	file_enum_proto_rawDescData = file_enum_proto_rawDesc
)

func file_enum_proto_rawDescGZIP() []byte {
	file_enum_proto_rawDescOnce.Do(func() {
		file_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_enum_proto_rawDescData)
	})
	return file_enum_proto_rawDescData
}

var file_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_enum_proto_goTypes = []interface{}{
	(AccountType)(0),     // 0: enums.AccountType
	(TransactionType)(0), // 1: enums.TransactionType
}
var file_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enum_proto_init() }
func file_enum_proto_init() {
	if File_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enum_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enum_proto_goTypes,
		DependencyIndexes: file_enum_proto_depIdxs,
		EnumInfos:         file_enum_proto_enumTypes,
	}.Build()
	File_enum_proto = out.File
	file_enum_proto_rawDesc = nil
	file_enum_proto_goTypes = nil
	file_enum_proto_depIdxs = nil
}
