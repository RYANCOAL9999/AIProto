// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.4
// source: general_expression.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//The response for header and describe.
type DataFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *DataFrame) Reset() {
	*x = DataFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_expression_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataFrame) ProtoMessage() {}

func (x *DataFrame) ProtoReflect() protoreflect.Message {
	mi := &file_general_expression_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataFrame.ProtoReflect.Descriptor instead.
func (*DataFrame) Descriptor() ([]byte, []int) {
	return file_general_expression_proto_rawDescGZIP(), []int{0}
}

func (x *DataFrame) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

// The request message containing the header's Data.
type HeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeaderRequest) Reset() {
	*x = HeaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_expression_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderRequest) ProtoMessage() {}

func (x *HeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_general_expression_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderRequest.ProtoReflect.Descriptor instead.
func (*HeaderRequest) Descriptor() ([]byte, []int) {
	return file_general_expression_proto_rawDescGZIP(), []int{1}
}

// The request message containing the info's Data.
type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_expression_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_general_expression_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_general_expression_proto_rawDescGZIP(), []int{2}
}

// The request message containing the describe's Data.
type DescribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DescribeRequest) Reset() {
	*x = DescribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_expression_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRequest) ProtoMessage() {}

func (x *DescribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_general_expression_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRequest.ProtoReflect.Descriptor instead.
func (*DescribeRequest) Descriptor() ([]byte, []int) {
	return file_general_expression_proto_rawDescGZIP(), []int{3}
}

var File_general_expression_proto protoreflect.FileDescriptor

var file_general_expression_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x41, 0x49, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43,
	0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48,
	0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x0f, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xc9, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x41, 0x49, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x41, 0x49, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x09, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x41, 0x49, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x41, 0x49, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x6c, 0x65, 0x72, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x2e, 0x41, 0x49, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x41,
	0x49, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_general_expression_proto_rawDescOnce sync.Once
	file_general_expression_proto_rawDescData = file_general_expression_proto_rawDesc
)

func file_general_expression_proto_rawDescGZIP() []byte {
	file_general_expression_proto_rawDescOnce.Do(func() {
		file_general_expression_proto_rawDescData = protoimpl.X.CompressGZIP(file_general_expression_proto_rawDescData)
	})
	return file_general_expression_proto_rawDescData
}

var file_general_expression_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_general_expression_proto_goTypes = []interface{}{
	(*DataFrame)(nil),       // 0: AIProto.DataFrame
	(*HeaderRequest)(nil),   // 1: AIProto.HeaderRequest
	(*InfoRequest)(nil),     // 2: AIProto.InfoRequest
	(*DescribeRequest)(nil), // 3: AIProto.DescribeRequest
	(*anypb.Any)(nil),       // 4: google.protobuf.Any
}
var file_general_expression_proto_depIdxs = []int32{
	4, // 0: AIProto.DataFrame.data:type_name -> google.protobuf.Any
	1, // 1: AIProto.GeneralService.HeaderEvent:input_type -> AIProto.HeaderRequest
	2, // 2: AIProto.GeneralService.InfoEvent:input_type -> AIProto.InfoRequest
	3, // 3: AIProto.GeneralService.DescriblerEvent:input_type -> AIProto.DescribeRequest
	0, // 4: AIProto.GeneralService.HeaderEvent:output_type -> AIProto.DataFrame
	0, // 5: AIProto.GeneralService.InfoEvent:output_type -> AIProto.DataFrame
	0, // 6: AIProto.GeneralService.DescriblerEvent:output_type -> AIProto.DataFrame
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_general_expression_proto_init() }
func file_general_expression_proto_init() {
	if File_general_expression_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_general_expression_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataFrame); i {
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
		file_general_expression_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderRequest); i {
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
		file_general_expression_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
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
		file_general_expression_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRequest); i {
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
	file_general_expression_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_general_expression_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_general_expression_proto_goTypes,
		DependencyIndexes: file_general_expression_proto_depIdxs,
		MessageInfos:      file_general_expression_proto_msgTypes,
	}.Build()
	File_general_expression_proto = out.File
	file_general_expression_proto_rawDesc = nil
	file_general_expression_proto_goTypes = nil
	file_general_expression_proto_depIdxs = nil
}
