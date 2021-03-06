// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/ms.proto

package memoryserver

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

// The request message containing the user's name.
type NumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *NumberRequest) Reset() {
	*x = NumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberRequest) ProtoMessage() {}

func (x *NumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberRequest.ProtoReflect.Descriptor instead.
func (*NumberRequest) Descriptor() ([]byte, []int) {
	return file_proto_ms_proto_rawDescGZIP(), []int{0}
}

func (x *NumberRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

// The response message containing the greetings
type DoIRemember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DoIRemember) Reset() {
	*x = DoIRemember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoIRemember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoIRemember) ProtoMessage() {}

func (x *DoIRemember) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoIRemember.ProtoReflect.Descriptor instead.
func (*DoIRemember) Descriptor() ([]byte, []int) {
	return file_proto_ms_proto_rawDescGZIP(), []int{1}
}

func (x *DoIRemember) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_ms_proto protoreflect.FileDescriptor

var file_proto_ms_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x27,
	0x0a, 0x0d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x0b, 0x44, 0x6f, 0x49, 0x52, 0x65,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x61, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x51, 0x0a, 0x15, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x49, 0x52, 0x65, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x00, 0x42, 0x5d, 0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x42, 0x11, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ms_proto_rawDescOnce sync.Once
	file_proto_ms_proto_rawDescData = file_proto_ms_proto_rawDesc
)

func file_proto_ms_proto_rawDescGZIP() []byte {
	file_proto_ms_proto_rawDescOnce.Do(func() {
		file_proto_ms_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ms_proto_rawDescData)
	})
	return file_proto_ms_proto_rawDescData
}

var file_proto_ms_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_ms_proto_goTypes = []interface{}{
	(*NumberRequest)(nil), // 0: memoryserver.NumberRequest
	(*DoIRemember)(nil),   // 1: memoryserver.DoIRemember
}
var file_proto_ms_proto_depIdxs = []int32{
	0, // 0: memoryserver.MemoryServer.checkWithMemoryServer:input_type -> memoryserver.NumberRequest
	1, // 1: memoryserver.MemoryServer.checkWithMemoryServer:output_type -> memoryserver.DoIRemember
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ms_proto_init() }
func file_proto_ms_proto_init() {
	if File_proto_ms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumberRequest); i {
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
		file_proto_ms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoIRemember); i {
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
			RawDescriptor: file_proto_ms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ms_proto_goTypes,
		DependencyIndexes: file_proto_ms_proto_depIdxs,
		MessageInfos:      file_proto_ms_proto_msgTypes,
	}.Build()
	File_proto_ms_proto = out.File
	file_proto_ms_proto_rawDesc = nil
	file_proto_ms_proto_goTypes = nil
	file_proto_ms_proto_depIdxs = nil
}
