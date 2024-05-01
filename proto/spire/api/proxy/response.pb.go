// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: spire/api/proxy/response.proto

package proxy

import (
	types "github.com/biareichert/spire-api-sdk/proto/spire/api/types"
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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp       string `protobuf:"bytes,1,opt,name=resp,proto3" json:"resp,omitempty"`
	Exitstatus int32  `protobuf:"varint,2,opt,name=exitstatus,proto3" json:"exitstatus,omitempty"`
	//string trust_domain = 3;
	//string path = 4;
	CertChain [][]byte `protobuf:"bytes,5,rep,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	//int64 expires_at = 6;
	//string hint = 7;
	Reattestable bool          `protobuf:"varint,8,opt,name=reattestable,proto3" json:"reattestable,omitempty"`
	AgentID      string        `protobuf:"bytes,9,opt,name=agentID,proto3" json:"agentID,omitempty"` //bytes svid = 10;
	Bundle       *types.Bundle `protobuf:"bytes,19,opt,name=bundle,proto3" json:"bundle,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_proxy_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_proxy_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_spire_api_proxy_response_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetResp() string {
	if x != nil {
		return x.Resp
	}
	return ""
}

func (x *Response) GetExitstatus() int32 {
	if x != nil {
		return x.Exitstatus
	}
	return 0
}

func (x *Response) GetCertChain() [][]byte {
	if x != nil {
		return x.CertChain
	}
	return nil
}

func (x *Response) GetReattestable() bool {
	if x != nil {
		return x.Reattestable
	}
	return false
}

func (x *Response) GetAgentID() string {
	if x != nil {
		return x.AgentID
	}
	return ""
}

func (x *Response) GetBundle() *types.Bundle {
	if x != nil {
		return x.Bundle
	}
	return nil
}

var File_spire_api_proxy_response_proto protoreflect.FileDescriptor

var file_spire_api_proxy_response_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x1a, 0x1c, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcc, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x73, 0x70,
	0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x69, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x78, 0x69, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x65, 0x72, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12,
	0x22, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x65, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x2f, 0x0a,
	0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x42, 0x3c,
	0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x61,
	0x72, 0x65, 0x69, 0x63, 0x68, 0x65, 0x72, 0x74, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_api_proxy_response_proto_rawDescOnce sync.Once
	file_spire_api_proxy_response_proto_rawDescData = file_spire_api_proxy_response_proto_rawDesc
)

func file_spire_api_proxy_response_proto_rawDescGZIP() []byte {
	file_spire_api_proxy_response_proto_rawDescOnce.Do(func() {
		file_spire_api_proxy_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_api_proxy_response_proto_rawDescData)
	})
	return file_spire_api_proxy_response_proto_rawDescData
}

var file_spire_api_proxy_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_spire_api_proxy_response_proto_goTypes = []interface{}{
	(*Response)(nil),     // 0: spire.api.proxy.Response
	(*types.Bundle)(nil), // 1: spire.api.types.Bundle
}
var file_spire_api_proxy_response_proto_depIdxs = []int32{
	1, // 0: spire.api.proxy.Response.bundle:type_name -> spire.api.types.Bundle
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spire_api_proxy_response_proto_init() }
func file_spire_api_proxy_response_proto_init() {
	if File_spire_api_proxy_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_api_proxy_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_spire_api_proxy_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spire_api_proxy_response_proto_goTypes,
		DependencyIndexes: file_spire_api_proxy_response_proto_depIdxs,
		MessageInfos:      file_spire_api_proxy_response_proto_msgTypes,
	}.Build()
	File_spire_api_proxy_response_proto = out.File
	file_spire_api_proxy_response_proto_rawDesc = nil
	file_spire_api_proxy_response_proto_goTypes = nil
	file_spire_api_proxy_response_proto_depIdxs = nil
}
