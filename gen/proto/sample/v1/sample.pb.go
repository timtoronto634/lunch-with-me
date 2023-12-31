// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/sample/v1/sample.proto

package samplev1

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

type GreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GreetRequest) Reset() {
	*x = GreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_v1_sample_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequest) ProtoMessage() {}

func (x *GreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_v1_sample_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequest.ProtoReflect.Descriptor instead.
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_v1_sample_proto_rawDescGZIP(), []int{0}
}

func (x *GreetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting string `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetResponse) Reset() {
	*x = GreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_v1_sample_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponse) ProtoMessage() {}

func (x *GreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_v1_sample_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponse.ProtoReflect.Descriptor instead.
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return file_proto_sample_v1_sample_proto_rawDescGZIP(), []int{1}
}

func (x *GreetResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

type ErrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idx int64 `protobuf:"varint,1,opt,name=idx,proto3" json:"idx,omitempty"`
}

func (x *ErrRequest) Reset() {
	*x = ErrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_v1_sample_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrRequest) ProtoMessage() {}

func (x *ErrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_v1_sample_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrRequest.ProtoReflect.Descriptor instead.
func (*ErrRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_v1_sample_proto_rawDescGZIP(), []int{2}
}

func (x *ErrRequest) GetIdx() int64 {
	if x != nil {
		return x.Idx
	}
	return 0
}

type ErrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ans string `protobuf:"bytes,1,opt,name=ans,proto3" json:"ans,omitempty"`
}

func (x *ErrResponse) Reset() {
	*x = ErrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_v1_sample_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrResponse) ProtoMessage() {}

func (x *ErrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_v1_sample_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrResponse.ProtoReflect.Descriptor instead.
func (*ErrResponse) Descriptor() ([]byte, []int) {
	return file_proto_sample_v1_sample_proto_rawDescGZIP(), []int{3}
}

func (x *ErrResponse) GetAns() string {
	if x != nil {
		return x.Ans
	}
	return ""
}

type DetaildErrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DetaildErrRequest) Reset() {
	*x = DetaildErrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_v1_sample_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetaildErrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetaildErrRequest) ProtoMessage() {}

func (x *DetaildErrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_v1_sample_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetaildErrRequest.ProtoReflect.Descriptor instead.
func (*DetaildErrRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_v1_sample_proto_rawDescGZIP(), []int{4}
}

type DetaildErrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DetaildErrResponse) Reset() {
	*x = DetaildErrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_v1_sample_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetaildErrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetaildErrResponse) ProtoMessage() {}

func (x *DetaildErrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_v1_sample_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetaildErrResponse.ProtoReflect.Descriptor instead.
func (*DetaildErrResponse) Descriptor() ([]byte, []int) {
	return file_proto_sample_v1_sample_proto_rawDescGZIP(), []int{5}
}

var File_proto_sample_v1_sample_proto protoreflect.FileDescriptor

var file_proto_sample_v1_sample_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22,
	0x22, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x0d, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x22, 0x1e, 0x0a, 0x0a, 0x45, 0x72, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x78,
	0x22, 0x1f, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6e,
	0x73, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x64, 0x45, 0x72, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x64, 0x45, 0x72, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf6, 0x01, 0x0a,
	0x0d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48,
	0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x03, 0x45, 0x72, 0x72, 0x12,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x72, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x72, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0a,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x64, 0x45, 0x72, 0x72, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x64, 0x45, 0x72, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x64, 0x45, 0x72, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x65, 0x63, 0x68, 0x6f, 0x2d, 0x6d, 0x65,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_sample_v1_sample_proto_rawDescOnce sync.Once
	file_proto_sample_v1_sample_proto_rawDescData = file_proto_sample_v1_sample_proto_rawDesc
)

func file_proto_sample_v1_sample_proto_rawDescGZIP() []byte {
	file_proto_sample_v1_sample_proto_rawDescOnce.Do(func() {
		file_proto_sample_v1_sample_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sample_v1_sample_proto_rawDescData)
	})
	return file_proto_sample_v1_sample_proto_rawDescData
}

var file_proto_sample_v1_sample_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_sample_v1_sample_proto_goTypes = []interface{}{
	(*GreetRequest)(nil),       // 0: proto.sample.v1.GreetRequest
	(*GreetResponse)(nil),      // 1: proto.sample.v1.GreetResponse
	(*ErrRequest)(nil),         // 2: proto.sample.v1.ErrRequest
	(*ErrResponse)(nil),        // 3: proto.sample.v1.ErrResponse
	(*DetaildErrRequest)(nil),  // 4: proto.sample.v1.DetaildErrRequest
	(*DetaildErrResponse)(nil), // 5: proto.sample.v1.DetaildErrResponse
}
var file_proto_sample_v1_sample_proto_depIdxs = []int32{
	0, // 0: proto.sample.v1.SampleService.Greet:input_type -> proto.sample.v1.GreetRequest
	2, // 1: proto.sample.v1.SampleService.Err:input_type -> proto.sample.v1.ErrRequest
	4, // 2: proto.sample.v1.SampleService.DetaildErr:input_type -> proto.sample.v1.DetaildErrRequest
	1, // 3: proto.sample.v1.SampleService.Greet:output_type -> proto.sample.v1.GreetResponse
	3, // 4: proto.sample.v1.SampleService.Err:output_type -> proto.sample.v1.ErrResponse
	5, // 5: proto.sample.v1.SampleService.DetaildErr:output_type -> proto.sample.v1.DetaildErrResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_sample_v1_sample_proto_init() }
func file_proto_sample_v1_sample_proto_init() {
	if File_proto_sample_v1_sample_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sample_v1_sample_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRequest); i {
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
		file_proto_sample_v1_sample_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetResponse); i {
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
		file_proto_sample_v1_sample_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrRequest); i {
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
		file_proto_sample_v1_sample_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrResponse); i {
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
		file_proto_sample_v1_sample_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetaildErrRequest); i {
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
		file_proto_sample_v1_sample_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetaildErrResponse); i {
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
			RawDescriptor: file_proto_sample_v1_sample_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sample_v1_sample_proto_goTypes,
		DependencyIndexes: file_proto_sample_v1_sample_proto_depIdxs,
		MessageInfos:      file_proto_sample_v1_sample_proto_msgTypes,
	}.Build()
	File_proto_sample_v1_sample_proto = out.File
	file_proto_sample_v1_sample_proto_rawDesc = nil
	file_proto_sample_v1_sample_proto_goTypes = nil
	file_proto_sample_v1_sample_proto_depIdxs = nil
}
