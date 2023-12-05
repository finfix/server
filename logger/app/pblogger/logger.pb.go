// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: app/pblogger/logger.proto

package pblogger

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

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Уровень лога
	Level string `protobuf:"bytes,1,opt,name=level,proto3" jsonapi:"level,omitempty"`
	// Путь до строки, на которой был вызван лог
	Path string `protobuf:"bytes,2,opt,name=path,proto3" jsonapi:"path,omitempty"`
	// Сообщение лога
	Message string `protobuf:"bytes,3,opt,name=message,proto3" jsonapi:"message,omitempty"`
	// Время лога
	Time string `protobuf:"bytes,4,opt,name=time,proto3" jsonapi:"time,omitempty"`
	// Сервис, в котором был вызван лог
	Service string `protobuf:"bytes,5,opt,name=service,proto3" jsonapi:"service,omitempty"`
	// Контекст ошибки
	Context *string `protobuf:"bytes,6,opt,name=context,proto3,oneof" jsonapi:"context,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_pblogger_logger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_app_pblogger_logger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_app_pblogger_logger_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Log) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Log) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Log) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Log) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Log) GetContext() string {
	if x != nil && x.Context != nil {
		return *x.Context
	}
	return ""
}

var File_app_pblogger_logger_proto protoreflect.FileDescriptor

var file_app_pblogger_logger_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x62, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x32, 0x39, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x12, 0x2f, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x67, 0x12, 0x0d, 0x2e, 0x70, 0x62,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x79, 0x43, 0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70,
	0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_pblogger_logger_proto_rawDescOnce sync.Once
	file_app_pblogger_logger_proto_rawDescData = file_app_pblogger_logger_proto_rawDesc
)

func file_app_pblogger_logger_proto_rawDescGZIP() []byte {
	file_app_pblogger_logger_proto_rawDescOnce.Do(func() {
		file_app_pblogger_logger_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_pblogger_logger_proto_rawDescData)
	})
	return file_app_pblogger_logger_proto_rawDescData
}

var file_app_pblogger_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_app_pblogger_logger_proto_goTypes = []interface{}{
	(*Log)(nil),           // 0: pblogger.Log
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_app_pblogger_logger_proto_depIdxs = []int32{
	0, // 0: pblogger.logger.AddLog:input_type -> pblogger.Log
	1, // 1: pblogger.logger.AddLog:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_pblogger_logger_proto_init() }
func file_app_pblogger_logger_proto_init() {
	if File_app_pblogger_logger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_pblogger_logger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
	file_app_pblogger_logger_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_pblogger_logger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_pblogger_logger_proto_goTypes,
		DependencyIndexes: file_app_pblogger_logger_proto_depIdxs,
		MessageInfos:      file_app_pblogger_logger_proto_msgTypes,
	}.Build()
	File_app_pblogger_logger_proto = out.File
	file_app_pblogger_logger_proto_rawDesc = nil
	file_app_pblogger_logger_proto_goTypes = nil
	file_app_pblogger_logger_proto_depIdxs = nil
}
