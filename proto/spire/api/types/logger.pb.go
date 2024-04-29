// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: spire/api/types/logger.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The logger log levels.
type LogLevel int32

const (
	LogLevel_UNSPECIFIED LogLevel = 0
	LogLevel_PANIC       LogLevel = 1
	LogLevel_FATAL       LogLevel = 2
	LogLevel_ERROR       LogLevel = 3
	LogLevel_WARN        LogLevel = 4
	LogLevel_INFO        LogLevel = 5
	LogLevel_DEBUG       LogLevel = 6
	LogLevel_TRACE       LogLevel = 7
)

// Enum value maps for LogLevel.
var (
	LogLevel_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "PANIC",
		2: "FATAL",
		3: "ERROR",
		4: "WARN",
		5: "INFO",
		6: "DEBUG",
		7: "TRACE",
	}
	LogLevel_value = map[string]int32{
		"UNSPECIFIED": 0,
		"PANIC":       1,
		"FATAL":       2,
		"ERROR":       3,
		"WARN":        4,
		"INFO":        5,
		"DEBUG":       6,
		"TRACE":       7,
	}
)

func (x LogLevel) Enum() *LogLevel {
	p := new(LogLevel)
	*p = x
	return p
}

func (x LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_spire_api_types_logger_proto_enumTypes[0].Descriptor()
}

func (LogLevel) Type() protoreflect.EnumType {
	return &file_spire_api_types_logger_proto_enumTypes[0]
}

func (x LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogLevel.Descriptor instead.
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_spire_api_types_logger_proto_rawDescGZIP(), []int{0}
}

// Represents the current Logger settings.
type Logger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The logger's current log level.
	CurrentLevel LogLevel `protobuf:"varint,1,opt,name=current_level,json=currentLevel,proto3,enum=spire.api.types.LogLevel" json:"current_level,omitempty"`
	// Output only. The logger's log level at process launch.
	LaunchLevel LogLevel `protobuf:"varint,2,opt,name=launch_level,json=launchLevel,proto3,enum=spire.api.types.LogLevel" json:"launch_level,omitempty"`
}

func (x *Logger) Reset() {
	*x = Logger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_types_logger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger) ProtoMessage() {}

func (x *Logger) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_types_logger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger.ProtoReflect.Descriptor instead.
func (*Logger) Descriptor() ([]byte, []int) {
	return file_spire_api_types_logger_proto_rawDescGZIP(), []int{0}
}

func (x *Logger) GetCurrentLevel() LogLevel {
	if x != nil {
		return x.CurrentLevel
	}
	return LogLevel_UNSPECIFIED
}

func (x *Logger) GetLaunchLevel() LogLevel {
	if x != nil {
		return x.LaunchLevel
	}
	return LogLevel_UNSPECIFIED
}

var File_spire_api_types_logger_proto protoreflect.FileDescriptor

var file_spire_api_types_logger_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x86, 0x01, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0d, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x3c, 0x0a, 0x0c, 0x6c, 0x61,
	0x75, 0x6e, 0x63, 0x68, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x0b, 0x6c, 0x61, 0x75,
	0x6e, 0x63, 0x68, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x2a, 0x66, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x4e, 0x49, 0x43, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x52, 0x4e, 0x10, 0x04,
	0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45,
	0x42, 0x55, 0x47, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x52, 0x41, 0x43, 0x45, 0x10, 0x07,
	0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x69, 0x61, 0x72, 0x65, 0x69, 0x63, 0x68, 0x65, 0x72, 0x74, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_api_types_logger_proto_rawDescOnce sync.Once
	file_spire_api_types_logger_proto_rawDescData = file_spire_api_types_logger_proto_rawDesc
)

func file_spire_api_types_logger_proto_rawDescGZIP() []byte {
	file_spire_api_types_logger_proto_rawDescOnce.Do(func() {
		file_spire_api_types_logger_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_api_types_logger_proto_rawDescData)
	})
	return file_spire_api_types_logger_proto_rawDescData
}

var file_spire_api_types_logger_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spire_api_types_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_spire_api_types_logger_proto_goTypes = []interface{}{
	(LogLevel)(0),  // 0: spire.api.types.LogLevel
	(*Logger)(nil), // 1: spire.api.types.Logger
}
var file_spire_api_types_logger_proto_depIdxs = []int32{
	0, // 0: spire.api.types.Logger.current_level:type_name -> spire.api.types.LogLevel
	0, // 1: spire.api.types.Logger.launch_level:type_name -> spire.api.types.LogLevel
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spire_api_types_logger_proto_init() }
func file_spire_api_types_logger_proto_init() {
	if File_spire_api_types_logger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_api_types_logger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logger); i {
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
			RawDescriptor: file_spire_api_types_logger_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spire_api_types_logger_proto_goTypes,
		DependencyIndexes: file_spire_api_types_logger_proto_depIdxs,
		EnumInfos:         file_spire_api_types_logger_proto_enumTypes,
		MessageInfos:      file_spire_api_types_logger_proto_msgTypes,
	}.Build()
	File_spire_api_types_logger_proto = out.File
	file_spire_api_types_logger_proto_rawDesc = nil
	file_spire_api_types_logger_proto_goTypes = nil
	file_spire_api_types_logger_proto_depIdxs = nil
}
