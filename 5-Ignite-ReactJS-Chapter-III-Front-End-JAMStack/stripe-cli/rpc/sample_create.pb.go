// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: sample_create.proto

package rpc

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

type SampleCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the sample, e.g. accept-a-card-payment. Use the `SamplesList` method to get a list of
	// available samples.
	SampleName string `protobuf:"bytes,1,opt,name=sample_name,json=sampleName,proto3" json:"sample_name,omitempty"`
	// Name of the particular integration, e.g. using-webhooks. Use the `SampleConfigs` method to get
	// the available options.
	IntegrationName string `protobuf:"bytes,2,opt,name=integration_name,json=integrationName,proto3" json:"integration_name,omitempty"`
	// Platform or language for the client, e.g. web. Use the `SampleConfigs` method to get the
	// available options.
	Client string `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	// Platform or language for the server, e.g. node. Use the `SampleConfigs` method to get the
	// available options.
	Server string `protobuf:"bytes,4,opt,name=server,proto3" json:"server,omitempty"`
	// Path to clone the repo to.
	Path string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	// If true, clear the local cache before creating the sample.
	ForceRefresh bool `protobuf:"varint,6,opt,name=force_refresh,json=forceRefresh,proto3" json:"force_refresh,omitempty"`
}

func (x *SampleCreateRequest) Reset() {
	*x = SampleCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleCreateRequest) ProtoMessage() {}

func (x *SampleCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sample_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleCreateRequest.ProtoReflect.Descriptor instead.
func (*SampleCreateRequest) Descriptor() ([]byte, []int) {
	return file_sample_create_proto_rawDescGZIP(), []int{0}
}

func (x *SampleCreateRequest) GetSampleName() string {
	if x != nil {
		return x.SampleName
	}
	return ""
}

func (x *SampleCreateRequest) GetIntegrationName() string {
	if x != nil {
		return x.IntegrationName
	}
	return ""
}

func (x *SampleCreateRequest) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *SampleCreateRequest) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *SampleCreateRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SampleCreateRequest) GetForceRefresh() bool {
	if x != nil {
		return x.ForceRefresh
	}
	return false
}

type SampleCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Additional instructions for the sample after install.
	PostInstall string `protobuf:"bytes,1,opt,name=post_install,json=postInstall,proto3" json:"post_install,omitempty"`
	// Path to the sample.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *SampleCreateResponse) Reset() {
	*x = SampleCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleCreateResponse) ProtoMessage() {}

func (x *SampleCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sample_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleCreateResponse.ProtoReflect.Descriptor instead.
func (*SampleCreateResponse) Descriptor() ([]byte, []int) {
	return file_sample_create_proto_rawDescGZIP(), []int{1}
}

func (x *SampleCreateResponse) GetPostInstall() string {
	if x != nil {
		return x.PostInstall
	}
	return ""
}

func (x *SampleCreateResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_sample_create_proto protoreflect.FileDescriptor

var file_sample_create_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0xca, 0x01, 0x0a, 0x13, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x22, 0x4d, 0x0a, 0x14, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x69,
	0x70, 0x65, 0x2d, 0x63, 0x6c, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sample_create_proto_rawDescOnce sync.Once
	file_sample_create_proto_rawDescData = file_sample_create_proto_rawDesc
)

func file_sample_create_proto_rawDescGZIP() []byte {
	file_sample_create_proto_rawDescOnce.Do(func() {
		file_sample_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_sample_create_proto_rawDescData)
	})
	return file_sample_create_proto_rawDescData
}

var file_sample_create_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sample_create_proto_goTypes = []interface{}{
	(*SampleCreateRequest)(nil),  // 0: rpc.SampleCreateRequest
	(*SampleCreateResponse)(nil), // 1: rpc.SampleCreateResponse
}
var file_sample_create_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sample_create_proto_init() }
func file_sample_create_proto_init() {
	if File_sample_create_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sample_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleCreateRequest); i {
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
		file_sample_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleCreateResponse); i {
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
			RawDescriptor: file_sample_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sample_create_proto_goTypes,
		DependencyIndexes: file_sample_create_proto_depIdxs,
		MessageInfos:      file_sample_create_proto_msgTypes,
	}.Build()
	File_sample_create_proto = out.File
	file_sample_create_proto_rawDesc = nil
	file_sample_create_proto_goTypes = nil
	file_sample_create_proto_depIdxs = nil
}
