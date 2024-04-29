// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: spire/api/types/federateswith.proto

package types

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

type FederatesWithMatch_MatchBehavior int32

const (
	// Indicates that the federated trust domains in this match are
	// equal to the candidate trust domains, independent of ordering.
	// Example:
	//   Given:
	//     - e1 { FederatesWith: ["spiffe://td1", "spiffe://td2", "spiffe://td3"]}
	//     - e2 { FederatesWith: ["spiffe://td1", "spiffe://td2"]}
	//     - e3 { FederatesWith: ["spiffe://td1"]}
	//   Operation:
	//     - MATCH_EXACT ["spiffe://td1", "spiffe://td2"]
	//   Entries that match:
	//     - 'e2'
	FederatesWithMatch_MATCH_EXACT FederatesWithMatch_MatchBehavior = 0
	// Indicates that all candidates which have a non-empty subset
	// of the provided set of trust domains will match.
	// Example:
	//   Given:
	//     - e1 { FederatesWith: ["spiffe://td1", "spiffe://td2", "spiffe://td3"]}
	//     - e2 { FederatesWith: ["spiffe://td1", "spiffe://td2"]}
	//     - e3 { FederatesWith: ["spiffe://td1"]}
	//   Operation:
	//     - MATCH_SUBSET ["spiffe://td1"]
	//   Entries that match:
	//     - 'e1'
	FederatesWithMatch_MATCH_SUBSET FederatesWithMatch_MatchBehavior = 1
	// Indicate that all candidates which are a superset
	// of the provided set of trust domains will match.
	// Example:
	//   Given:
	//     - e1 { FederatesWith: ["spiffe://td1", "spiffe://td2", "spiffe://td3"]}
	//     - e2 { FederatesWith: ["spiffe://td1", "spiffe://td2"]}
	//     - e3 { FederatesWith: ["spiffe://td1"]}
	//   Operation:
	//     - MATCH_SUPERSET ["spiffe://td1", "spiffe://td2"]
	//   Entries that match:
	//     - 'e1'
	//     - 'e2'
	FederatesWithMatch_MATCH_SUPERSET FederatesWithMatch_MatchBehavior = 2
	// Indicates that all candidates which have at least one
	// of the provided set of trust domains will match.
	// Example:
	//   Given:
	//     - e1 { FederatesWith: ["spiffe://td1", "spiffe://td2", "spiffe://td3"]}
	//     - e2 { FederatesWith: ["spiffe://td1", "spiffe://td2"]}
	//     - e3 { FederatesWith: ["spiffe://td1"]}
	//   Operation:
	//     - MATCH_ANY ["spiffe://td1"]
	//   Entries that match:
	//     - 'e1'
	//     - 'e2'
	//     - 'e3'
	FederatesWithMatch_MATCH_ANY FederatesWithMatch_MatchBehavior = 3
)

// Enum value maps for FederatesWithMatch_MatchBehavior.
var (
	FederatesWithMatch_MatchBehavior_name = map[int32]string{
		0: "MATCH_EXACT",
		1: "MATCH_SUBSET",
		2: "MATCH_SUPERSET",
		3: "MATCH_ANY",
	}
	FederatesWithMatch_MatchBehavior_value = map[string]int32{
		"MATCH_EXACT":    0,
		"MATCH_SUBSET":   1,
		"MATCH_SUPERSET": 2,
		"MATCH_ANY":      3,
	}
)

func (x FederatesWithMatch_MatchBehavior) Enum() *FederatesWithMatch_MatchBehavior {
	p := new(FederatesWithMatch_MatchBehavior)
	*p = x
	return p
}

func (x FederatesWithMatch_MatchBehavior) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FederatesWithMatch_MatchBehavior) Descriptor() protoreflect.EnumDescriptor {
	return file_spire_api_types_federateswith_proto_enumTypes[0].Descriptor()
}

func (FederatesWithMatch_MatchBehavior) Type() protoreflect.EnumType {
	return &file_spire_api_types_federateswith_proto_enumTypes[0]
}

func (x FederatesWithMatch_MatchBehavior) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FederatesWithMatch_MatchBehavior.Descriptor instead.
func (FederatesWithMatch_MatchBehavior) EnumDescriptor() ([]byte, []int) {
	return file_spire_api_types_federateswith_proto_rawDescGZIP(), []int{0, 0}
}

type FederatesWithMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The set of trust domain names to match on (e.g., "example.org").
	TrustDomains []string `protobuf:"bytes,1,rep,name=trust_domains,json=trustDomains,proto3" json:"trust_domains,omitempty"`
	// How to match the trust domains.
	Match FederatesWithMatch_MatchBehavior `protobuf:"varint,2,opt,name=match,proto3,enum=spire.api.types.FederatesWithMatch_MatchBehavior" json:"match,omitempty"`
}

func (x *FederatesWithMatch) Reset() {
	*x = FederatesWithMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_types_federateswith_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederatesWithMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederatesWithMatch) ProtoMessage() {}

func (x *FederatesWithMatch) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_types_federateswith_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederatesWithMatch.ProtoReflect.Descriptor instead.
func (*FederatesWithMatch) Descriptor() ([]byte, []int) {
	return file_spire_api_types_federateswith_proto_rawDescGZIP(), []int{0}
}

func (x *FederatesWithMatch) GetTrustDomains() []string {
	if x != nil {
		return x.TrustDomains
	}
	return nil
}

func (x *FederatesWithMatch) GetMatch() FederatesWithMatch_MatchBehavior {
	if x != nil {
		return x.Match
	}
	return FederatesWithMatch_MATCH_EXACT
}

var File_spire_api_types_federateswith_proto protoreflect.FileDescriptor

var file_spire_api_types_federateswith_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x73, 0x77, 0x69, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0xd9, 0x01, 0x0a, 0x12, 0x46, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x23, 0x0a,
	0x0d, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x75, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x73, 0x12, 0x47, 0x0a, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x31, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x73, 0x57, 0x69, 0x74,
	0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x22, 0x55, 0x0a, 0x0d, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b,
	0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x45, 0x58, 0x41, 0x43, 0x54, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x45, 0x54, 0x10, 0x01, 0x12,
	0x12, 0x0a, 0x0e, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x55, 0x50, 0x45, 0x52, 0x53, 0x45,
	0x54, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x41, 0x4e, 0x59,
	0x10, 0x03, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x69, 0x61, 0x72, 0x65, 0x69, 0x63, 0x68, 0x65, 0x72, 0x74, 0x2f, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_api_types_federateswith_proto_rawDescOnce sync.Once
	file_spire_api_types_federateswith_proto_rawDescData = file_spire_api_types_federateswith_proto_rawDesc
)

func file_spire_api_types_federateswith_proto_rawDescGZIP() []byte {
	file_spire_api_types_federateswith_proto_rawDescOnce.Do(func() {
		file_spire_api_types_federateswith_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_api_types_federateswith_proto_rawDescData)
	})
	return file_spire_api_types_federateswith_proto_rawDescData
}

var file_spire_api_types_federateswith_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spire_api_types_federateswith_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_spire_api_types_federateswith_proto_goTypes = []interface{}{
	(FederatesWithMatch_MatchBehavior)(0), // 0: spire.api.types.FederatesWithMatch.MatchBehavior
	(*FederatesWithMatch)(nil),            // 1: spire.api.types.FederatesWithMatch
}
var file_spire_api_types_federateswith_proto_depIdxs = []int32{
	0, // 0: spire.api.types.FederatesWithMatch.match:type_name -> spire.api.types.FederatesWithMatch.MatchBehavior
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spire_api_types_federateswith_proto_init() }
func file_spire_api_types_federateswith_proto_init() {
	if File_spire_api_types_federateswith_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_api_types_federateswith_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederatesWithMatch); i {
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
			RawDescriptor: file_spire_api_types_federateswith_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spire_api_types_federateswith_proto_goTypes,
		DependencyIndexes: file_spire_api_types_federateswith_proto_depIdxs,
		EnumInfos:         file_spire_api_types_federateswith_proto_enumTypes,
		MessageInfos:      file_spire_api_types_federateswith_proto_msgTypes,
	}.Build()
	File_spire_api_types_federateswith_proto = out.File
	file_spire_api_types_federateswith_proto_rawDesc = nil
	file_spire_api_types_federateswith_proto_goTypes = nil
	file_spire_api_types_federateswith_proto_depIdxs = nil
}
