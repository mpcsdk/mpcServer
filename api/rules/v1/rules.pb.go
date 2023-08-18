// protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: rules/v1/rules.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RiskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contract string `protobuf:"bytes,1,opt,name=Contract,proto3" json:"Contract,omitempty" v:"required"` // v: required
	Method   string `protobuf:"bytes,2,opt,name=Method,proto3" json:"Method,omitempty" v:"required"`     // v: required
	Data     string `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty" v:"required"`         // v: required
}

func (x *RiskReq) Reset() {
	*x = RiskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rules_v1_rules_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiskReq) ProtoMessage() {}

func (x *RiskReq) ProtoReflect() protoreflect.Message {
	mi := &file_rules_v1_rules_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiskReq.ProtoReflect.Descriptor instead.
func (*RiskReq) Descriptor() ([]byte, []int) {
	return file_rules_v1_rules_proto_rawDescGZIP(), []int{0}
}

func (x *RiskReq) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *RiskReq) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *RiskReq) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type RiskRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RiskRes) Reset() {
	*x = RiskRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rules_v1_rules_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiskRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiskRes) ProtoMessage() {}

func (x *RiskRes) ProtoReflect() protoreflect.Message {
	mi := &file_rules_v1_rules_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiskRes.ProtoReflect.Descriptor instead.
func (*RiskRes) Descriptor() ([]byte, []int) {
	return file_rules_v1_rules_proto_rawDescGZIP(), []int{1}
}

var File_rules_v1_rules_proto protoreflect.FileDescriptor

var file_rules_v1_rules_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x51, 0x0a,
	0x07, 0x52, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x09, 0x0a, 0x07, 0x52, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x32, 0x37, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x0b, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x69,
	0x73, 0x6b, 0x12, 0x0e, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x52, 0x69, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x52, 0x69, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rules_v1_rules_proto_rawDescOnce sync.Once
	file_rules_v1_rules_proto_rawDescData = file_rules_v1_rules_proto_rawDesc
)

func file_rules_v1_rules_proto_rawDescGZIP() []byte {
	file_rules_v1_rules_proto_rawDescOnce.Do(func() {
		file_rules_v1_rules_proto_rawDescData = protoimpl.X.CompressGZIP(file_rules_v1_rules_proto_rawDescData)
	})
	return file_rules_v1_rules_proto_rawDescData
}

var file_rules_v1_rules_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rules_v1_rules_proto_goTypes = []interface{}{
	(*RiskReq)(nil), // 0: rules.RiskReq
	(*RiskRes)(nil), // 1: rules.RiskRes
}
var file_rules_v1_rules_proto_depIdxs = []int32{
	0, // 0: rules.User.PerformRisk:input_type -> rules.RiskReq
	1, // 1: rules.User.PerformRisk:output_type -> rules.RiskRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rules_v1_rules_proto_init() }
func file_rules_v1_rules_proto_init() {
	if File_rules_v1_rules_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rules_v1_rules_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiskReq); i {
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
		file_rules_v1_rules_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiskRes); i {
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
			RawDescriptor: file_rules_v1_rules_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rules_v1_rules_proto_goTypes,
		DependencyIndexes: file_rules_v1_rules_proto_depIdxs,
		MessageInfos:      file_rules_v1_rules_proto_msgTypes,
	}.Build()
	File_rules_v1_rules_proto = out.File
	file_rules_v1_rules_proto_rawDesc = nil
	file_rules_v1_rules_proto_goTypes = nil
	file_rules_v1_rules_proto_depIdxs = nil
}
