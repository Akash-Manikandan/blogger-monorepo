// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: health/v1/health.proto

package healthv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	mi := &file_health_v1_health_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_health_v1_health_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_health_v1_health_proto_rawDescGZIP(), []int{0}
}

type CheckResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	DbStatus      string                 `protobuf:"bytes,2,opt,name=db_status,json=dbStatus,proto3" json:"db_status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	mi := &file_health_v1_health_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_health_v1_health_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_health_v1_health_proto_rawDescGZIP(), []int{1}
}

func (x *CheckResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CheckResponse) GetDbStatus() string {
	if x != nil {
		return x.DbStatus
	}
	return ""
}

var File_health_v1_health_proto protoreflect.FileDescriptor

var file_health_v1_health_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x22, 0x0e, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x4b, 0x0a, 0x0d, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa2, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6b, 0x61, 0x73, 0x68, 0x2d, 0x4d, 0x61, 0x6e, 0x69, 0x6b, 0x61,
	0x6e, 0x64, 0x61, 0x6e, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2d, 0x62, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x58, 0x58, 0xaa, 0x02,
	0x09, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0a, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_health_v1_health_proto_rawDescOnce sync.Once
	file_health_v1_health_proto_rawDescData []byte
)

func file_health_v1_health_proto_rawDescGZIP() []byte {
	file_health_v1_health_proto_rawDescOnce.Do(func() {
		file_health_v1_health_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_health_v1_health_proto_rawDesc), len(file_health_v1_health_proto_rawDesc)))
	})
	return file_health_v1_health_proto_rawDescData
}

var file_health_v1_health_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_health_v1_health_proto_goTypes = []any{
	(*CheckRequest)(nil),  // 0: health.v1.CheckRequest
	(*CheckResponse)(nil), // 1: health.v1.CheckResponse
}
var file_health_v1_health_proto_depIdxs = []int32{
	0, // 0: health.v1.HealthService.Check:input_type -> health.v1.CheckRequest
	1, // 1: health.v1.HealthService.Check:output_type -> health.v1.CheckResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_health_v1_health_proto_init() }
func file_health_v1_health_proto_init() {
	if File_health_v1_health_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_health_v1_health_proto_rawDesc), len(file_health_v1_health_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_health_v1_health_proto_goTypes,
		DependencyIndexes: file_health_v1_health_proto_depIdxs,
		MessageInfos:      file_health_v1_health_proto_msgTypes,
	}.Build()
	File_health_v1_health_proto = out.File
	file_health_v1_health_proto_goTypes = nil
	file_health_v1_health_proto_depIdxs = nil
}
