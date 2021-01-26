//Implementation of the GRPC health package

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: protos/health.proto

package health

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type JoinNetworkResponse_ServerStatus int32

const (
	JoinNetworkResponse_HEAD JoinNetworkResponse_ServerStatus = 0
	JoinNetworkResponse_TAIL JoinNetworkResponse_ServerStatus = 1
)

// Enum value maps for JoinNetworkResponse_ServerStatus.
var (
	JoinNetworkResponse_ServerStatus_name = map[int32]string{
		0: "HEAD",
		1: "TAIL",
	}
	JoinNetworkResponse_ServerStatus_value = map[string]int32{
		"HEAD": 0,
		"TAIL": 1,
	}
)

func (x JoinNetworkResponse_ServerStatus) Enum() *JoinNetworkResponse_ServerStatus {
	p := new(JoinNetworkResponse_ServerStatus)
	*p = x
	return p
}

func (x JoinNetworkResponse_ServerStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JoinNetworkResponse_ServerStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_health_proto_enumTypes[0].Descriptor()
}

func (JoinNetworkResponse_ServerStatus) Type() protoreflect.EnumType {
	return &file_protos_health_proto_enumTypes[0]
}

func (x JoinNetworkResponse_ServerStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JoinNetworkResponse_ServerStatus.Descriptor instead.
func (JoinNetworkResponse_ServerStatus) EnumDescriptor() ([]byte, []int) {
	return file_protos_health_proto_rawDescGZIP(), []int{1, 0}
}

type JoinNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
}

func (x *JoinNetworkRequest) Reset() {
	*x = JoinNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinNetworkRequest) ProtoMessage() {}

func (x *JoinNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinNetworkRequest.ProtoReflect.Descriptor instead.
func (*JoinNetworkRequest) Descriptor() ([]byte, []int) {
	return file_protos_health_proto_rawDescGZIP(), []int{0}
}

func (x *JoinNetworkRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type JoinNetworkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status JoinNetworkResponse_ServerStatus `protobuf:"varint,1,opt,name=status,proto3,enum=JoinNetworkResponse_ServerStatus" json:"status,omitempty"`
}

func (x *JoinNetworkResponse) Reset() {
	*x = JoinNetworkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinNetworkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinNetworkResponse) ProtoMessage() {}

func (x *JoinNetworkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinNetworkResponse.ProtoReflect.Descriptor instead.
func (*JoinNetworkResponse) Descriptor() ([]byte, []int) {
	return file_protos_health_proto_rawDescGZIP(), []int{1}
}

func (x *JoinNetworkResponse) GetStatus() JoinNetworkResponse_ServerStatus {
	if x != nil {
		return x.Status
	}
	return JoinNetworkResponse_HEAD
}

var File_protos_health_proto protoreflect.FileDescriptor

var file_protos_health_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x12, 0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x74, 0x0a,
	0x13, 0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x22, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x08, 0x0a, 0x04, 0x48, 0x45, 0x41, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x41, 0x49,
	0x4c, 0x10, 0x01, 0x32, 0x42, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x38, 0x0a,
	0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x13, 0x2e, 0x4a,
	0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x44, 0x4a, 0x46, 0x69, 0x73, 0x68, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_health_proto_rawDescOnce sync.Once
	file_protos_health_proto_rawDescData = file_protos_health_proto_rawDesc
)

func file_protos_health_proto_rawDescGZIP() []byte {
	file_protos_health_proto_rawDescOnce.Do(func() {
		file_protos_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_health_proto_rawDescData)
	})
	return file_protos_health_proto_rawDescData
}

var file_protos_health_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_health_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_health_proto_goTypes = []interface{}{
	(JoinNetworkResponse_ServerStatus)(0), // 0: JoinNetworkResponse.ServerStatus
	(*JoinNetworkRequest)(nil),            // 1: JoinNetworkRequest
	(*JoinNetworkResponse)(nil),           // 2: JoinNetworkResponse
}
var file_protos_health_proto_depIdxs = []int32{
	0, // 0: JoinNetworkResponse.status:type_name -> JoinNetworkResponse.ServerStatus
	1, // 1: Health.JoinNetwork:input_type -> JoinNetworkRequest
	2, // 2: Health.JoinNetwork:output_type -> JoinNetworkResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_health_proto_init() }
func file_protos_health_proto_init() {
	if File_protos_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinNetworkRequest); i {
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
		file_protos_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinNetworkResponse); i {
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
			RawDescriptor: file_protos_health_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_health_proto_goTypes,
		DependencyIndexes: file_protos_health_proto_depIdxs,
		EnumInfos:         file_protos_health_proto_enumTypes,
		MessageInfos:      file_protos_health_proto_msgTypes,
	}.Build()
	File_protos_health_proto = out.File
	file_protos_health_proto_rawDesc = nil
	file_protos_health_proto_goTypes = nil
	file_protos_health_proto_depIdxs = nil
}
