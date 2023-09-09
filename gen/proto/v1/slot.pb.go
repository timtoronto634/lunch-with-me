// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/v1/slot.proto

package slotv1

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

type SaveSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SaveSlotRequest) Reset() {
	*x = SaveSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_slot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveSlotRequest) ProtoMessage() {}

func (x *SaveSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_slot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveSlotRequest.ProtoReflect.Descriptor instead.
func (*SaveSlotRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_slot_proto_rawDescGZIP(), []int{0}
}

func (x *SaveSlotRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SaveSlotRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SaveSlotRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SaveSlotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SaveSlotResponse) Reset() {
	*x = SaveSlotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_slot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveSlotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveSlotResponse) ProtoMessage() {}

func (x *SaveSlotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_slot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveSlotResponse.ProtoReflect.Descriptor instead.
func (*SaveSlotResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_slot_proto_rawDescGZIP(), []int{1}
}

func (x *SaveSlotResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_v1_slot_proto protoreflect.FileDescriptor

var file_proto_v1_slot_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6c, 0x6f, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x22,
	0x4b, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x22, 0x0a, 0x10,
	0x53, 0x61, 0x76, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x32, 0x52, 0x0a, 0x0b, 0x53, 0x6c, 0x6f, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x43, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x65, 0x63, 0x68, 0x6f, 0x2d, 0x6d, 0x65, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x6c, 0x6f,
	0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_slot_proto_rawDescOnce sync.Once
	file_proto_v1_slot_proto_rawDescData = file_proto_v1_slot_proto_rawDesc
)

func file_proto_v1_slot_proto_rawDescGZIP() []byte {
	file_proto_v1_slot_proto_rawDescOnce.Do(func() {
		file_proto_v1_slot_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_slot_proto_rawDescData)
	})
	return file_proto_v1_slot_proto_rawDescData
}

var file_proto_v1_slot_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_v1_slot_proto_goTypes = []interface{}{
	(*SaveSlotRequest)(nil),  // 0: proto.v1.SaveSlotRequest
	(*SaveSlotResponse)(nil), // 1: proto.v1.SaveSlotResponse
}
var file_proto_v1_slot_proto_depIdxs = []int32{
	0, // 0: proto.v1.SlotService.SaveSlot:input_type -> proto.v1.SaveSlotRequest
	1, // 1: proto.v1.SlotService.SaveSlot:output_type -> proto.v1.SaveSlotResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_v1_slot_proto_init() }
func file_proto_v1_slot_proto_init() {
	if File_proto_v1_slot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_slot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveSlotRequest); i {
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
		file_proto_v1_slot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveSlotResponse); i {
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
			RawDescriptor: file_proto_v1_slot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_slot_proto_goTypes,
		DependencyIndexes: file_proto_v1_slot_proto_depIdxs,
		MessageInfos:      file_proto_v1_slot_proto_msgTypes,
	}.Build()
	File_proto_v1_slot_proto = out.File
	file_proto_v1_slot_proto_rawDesc = nil
	file_proto_v1_slot_proto_goTypes = nil
	file_proto_v1_slot_proto_depIdxs = nil
}
