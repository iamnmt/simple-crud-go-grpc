// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: server/types/tx.proto

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

type MsgCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *MsgCreate) Reset() {
	*x = MsgCreate{}
	mi := &file_server_types_tx_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MsgCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreate) ProtoMessage() {}

func (x *MsgCreate) ProtoReflect() protoreflect.Message {
	mi := &file_server_types_tx_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCreate.ProtoReflect.Descriptor instead.
func (*MsgCreate) Descriptor() ([]byte, []int) {
	return file_server_types_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgCreate) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MsgCreate) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type MsgCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MsgCreateResponse) Reset() {
	*x = MsgCreateResponse{}
	mi := &file_server_types_tx_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MsgCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateResponse) ProtoMessage() {}

func (x *MsgCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_types_tx_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCreateResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateResponse) Descriptor() ([]byte, []int) {
	return file_server_types_tx_proto_rawDescGZIP(), []int{1}
}

func (x *MsgCreateResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type MsgRead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MsgRead) Reset() {
	*x = MsgRead{}
	mi := &file_server_types_tx_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MsgRead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRead) ProtoMessage() {}

func (x *MsgRead) ProtoReflect() protoreflect.Message {
	mi := &file_server_types_tx_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgRead.ProtoReflect.Descriptor instead.
func (*MsgRead) Descriptor() ([]byte, []int) {
	return file_server_types_tx_proto_rawDescGZIP(), []int{2}
}

func (x *MsgRead) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MsgReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *MsgReadResponse) Reset() {
	*x = MsgReadResponse{}
	mi := &file_server_types_tx_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MsgReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgReadResponse) ProtoMessage() {}

func (x *MsgReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_types_tx_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgReadResponse.ProtoReflect.Descriptor instead.
func (*MsgReadResponse) Descriptor() ([]byte, []int) {
	return file_server_types_tx_proto_rawDescGZIP(), []int{3}
}

func (x *MsgReadResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type MsgUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *MsgUpdate) Reset() {
	*x = MsgUpdate{}
	mi := &file_server_types_tx_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MsgUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdate) ProtoMessage() {}

func (x *MsgUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_server_types_tx_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgUpdate.ProtoReflect.Descriptor instead.
func (*MsgUpdate) Descriptor() ([]byte, []int) {
	return file_server_types_tx_proto_rawDescGZIP(), []int{4}
}

func (x *MsgUpdate) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MsgUpdate) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type MsgUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MsgUpdateResponse) Reset() {
	*x = MsgUpdateResponse{}
	mi := &file_server_types_tx_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MsgUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateResponse) ProtoMessage() {}

func (x *MsgUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_types_tx_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgUpdateResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateResponse) Descriptor() ([]byte, []int) {
	return file_server_types_tx_proto_rawDescGZIP(), []int{5}
}

func (x *MsgUpdateResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type MsgDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MsgDelete) Reset() {
	*x = MsgDelete{}
	mi := &file_server_types_tx_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MsgDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDelete) ProtoMessage() {}

func (x *MsgDelete) ProtoReflect() protoreflect.Message {
	mi := &file_server_types_tx_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDelete.ProtoReflect.Descriptor instead.
func (*MsgDelete) Descriptor() ([]byte, []int) {
	return file_server_types_tx_proto_rawDescGZIP(), []int{6}
}

func (x *MsgDelete) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MsgDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MsgDeleteResponse) Reset() {
	*x = MsgDeleteResponse{}
	mi := &file_server_types_tx_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MsgDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDeleteResponse) ProtoMessage() {}

func (x *MsgDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_types_tx_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDeleteResponse.ProtoReflect.Descriptor instead.
func (*MsgDeleteResponse) Descriptor() ([]byte, []int) {
	return file_server_types_tx_proto_rawDescGZIP(), []int{7}
}

func (x *MsgDeleteResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_server_types_tx_proto protoreflect.FileDescriptor

var file_server_types_tx_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x35,
	0x0a, 0x09, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x19, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a,
	0x0f, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x35, 0x0a, 0x09, 0x4d, 0x73,
	0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x2b, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1b,
	0x0a, 0x09, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x11, 0x4d,
	0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xd7, 0x01, 0x0a, 0x03, 0x4d, 0x73, 0x67,
	0x12, 0x34, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x18, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0e,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x61, 0x64, 0x1a, 0x16,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x1a, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d,
	0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x6d, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_types_tx_proto_rawDescOnce sync.Once
	file_server_types_tx_proto_rawDescData = file_server_types_tx_proto_rawDesc
)

func file_server_types_tx_proto_rawDescGZIP() []byte {
	file_server_types_tx_proto_rawDescOnce.Do(func() {
		file_server_types_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_types_tx_proto_rawDescData)
	})
	return file_server_types_tx_proto_rawDescData
}

var file_server_types_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_server_types_tx_proto_goTypes = []any{
	(*MsgCreate)(nil),         // 0: types.MsgCreate
	(*MsgCreateResponse)(nil), // 1: types.MsgCreateResponse
	(*MsgRead)(nil),           // 2: types.MsgRead
	(*MsgReadResponse)(nil),   // 3: types.MsgReadResponse
	(*MsgUpdate)(nil),         // 4: types.MsgUpdate
	(*MsgUpdateResponse)(nil), // 5: types.MsgUpdateResponse
	(*MsgDelete)(nil),         // 6: types.MsgDelete
	(*MsgDeleteResponse)(nil), // 7: types.MsgDeleteResponse
}
var file_server_types_tx_proto_depIdxs = []int32{
	0, // 0: types.Msg.Create:input_type -> types.MsgCreate
	2, // 1: types.Msg.Read:input_type -> types.MsgRead
	4, // 2: types.Msg.Update:input_type -> types.MsgUpdate
	6, // 3: types.Msg.Delete:input_type -> types.MsgDelete
	1, // 4: types.Msg.Create:output_type -> types.MsgCreateResponse
	3, // 5: types.Msg.Read:output_type -> types.MsgReadResponse
	5, // 6: types.Msg.Update:output_type -> types.MsgUpdateResponse
	7, // 7: types.Msg.Delete:output_type -> types.MsgDeleteResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_types_tx_proto_init() }
func file_server_types_tx_proto_init() {
	if File_server_types_tx_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_types_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_types_tx_proto_goTypes,
		DependencyIndexes: file_server_types_tx_proto_depIdxs,
		MessageInfos:      file_server_types_tx_proto_msgTypes,
	}.Build()
	File_server_types_tx_proto = out.File
	file_server_types_tx_proto_rawDesc = nil
	file_server_types_tx_proto_goTypes = nil
	file_server_types_tx_proto_depIdxs = nil
}
