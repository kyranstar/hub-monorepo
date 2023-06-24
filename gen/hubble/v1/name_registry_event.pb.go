// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: hubble/v1/name_registry_event.proto

package v1

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

type NameRegistryEventType int32

const (
	NameRegistryEventType_NAME_REGISTRY_EVENT_TYPE_NONE     NameRegistryEventType = 0
	NameRegistryEventType_NAME_REGISTRY_EVENT_TYPE_TRANSFER NameRegistryEventType = 1
	NameRegistryEventType_NAME_REGISTRY_EVENT_TYPE_RENEW    NameRegistryEventType = 2
)

// Enum value maps for NameRegistryEventType.
var (
	NameRegistryEventType_name = map[int32]string{
		0: "NAME_REGISTRY_EVENT_TYPE_NONE",
		1: "NAME_REGISTRY_EVENT_TYPE_TRANSFER",
		2: "NAME_REGISTRY_EVENT_TYPE_RENEW",
	}
	NameRegistryEventType_value = map[string]int32{
		"NAME_REGISTRY_EVENT_TYPE_NONE":     0,
		"NAME_REGISTRY_EVENT_TYPE_TRANSFER": 1,
		"NAME_REGISTRY_EVENT_TYPE_RENEW":    2,
	}
)

func (x NameRegistryEventType) Enum() *NameRegistryEventType {
	p := new(NameRegistryEventType)
	*p = x
	return p
}

func (x NameRegistryEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NameRegistryEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_hubble_v1_name_registry_event_proto_enumTypes[0].Descriptor()
}

func (NameRegistryEventType) Type() protoreflect.EnumType {
	return &file_hubble_v1_name_registry_event_proto_enumTypes[0]
}

func (x NameRegistryEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NameRegistryEventType.Descriptor instead.
func (NameRegistryEventType) EnumDescriptor() ([]byte, []int) {
	return file_hubble_v1_name_registry_event_proto_rawDescGZIP(), []int{0}
}

type NameRegistryEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber     uint32                `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	BlockHash       []byte                `protobuf:"bytes,2,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	TransactionHash []byte                `protobuf:"bytes,3,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	LogIndex        uint32                `protobuf:"varint,4,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
	Fname           []byte                `protobuf:"bytes,5,opt,name=fname,proto3" json:"fname,omitempty"`
	From            []byte                `protobuf:"bytes,6,opt,name=from,proto3" json:"from,omitempty"`
	To              []byte                `protobuf:"bytes,7,opt,name=to,proto3" json:"to,omitempty"`
	Type            NameRegistryEventType `protobuf:"varint,8,opt,name=type,proto3,enum=hubble.v1.NameRegistryEventType" json:"type,omitempty"`
	Expiry          uint32                `protobuf:"varint,9,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *NameRegistryEvent) Reset() {
	*x = NameRegistryEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hubble_v1_name_registry_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameRegistryEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameRegistryEvent) ProtoMessage() {}

func (x *NameRegistryEvent) ProtoReflect() protoreflect.Message {
	mi := &file_hubble_v1_name_registry_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameRegistryEvent.ProtoReflect.Descriptor instead.
func (*NameRegistryEvent) Descriptor() ([]byte, []int) {
	return file_hubble_v1_name_registry_event_proto_rawDescGZIP(), []int{0}
}

func (x *NameRegistryEvent) GetBlockNumber() uint32 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *NameRegistryEvent) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *NameRegistryEvent) GetTransactionHash() []byte {
	if x != nil {
		return x.TransactionHash
	}
	return nil
}

func (x *NameRegistryEvent) GetLogIndex() uint32 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *NameRegistryEvent) GetFname() []byte {
	if x != nil {
		return x.Fname
	}
	return nil
}

func (x *NameRegistryEvent) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *NameRegistryEvent) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *NameRegistryEvent) GetType() NameRegistryEventType {
	if x != nil {
		return x.Type
	}
	return NameRegistryEventType_NAME_REGISTRY_EVENT_TYPE_NONE
}

func (x *NameRegistryEvent) GetExpiry() uint32 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

var File_hubble_v1_name_registry_event_proto protoreflect.FileDescriptor

var file_hubble_v1_name_registry_event_proto_rawDesc = []byte{
	0x0a, 0x23, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x22, 0xa5, 0x02, 0x0a, 0x11, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x2a, 0x85, 0x01, 0x0a, 0x15, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x21, 0x0a, 0x1d, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53,
	0x54, 0x52, 0x59, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x52, 0x45,
	0x47, 0x49, 0x53, 0x54, 0x52, 0x59, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e,
	0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x59, 0x5f, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x4e, 0x45, 0x57, 0x10, 0x02,
	0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x79, 0x72, 0x61, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2f, 0x68, 0x75, 0x62, 0x2d, 0x6d, 0x6f, 0x6e,
	0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x75, 0x62,
	0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hubble_v1_name_registry_event_proto_rawDescOnce sync.Once
	file_hubble_v1_name_registry_event_proto_rawDescData = file_hubble_v1_name_registry_event_proto_rawDesc
)

func file_hubble_v1_name_registry_event_proto_rawDescGZIP() []byte {
	file_hubble_v1_name_registry_event_proto_rawDescOnce.Do(func() {
		file_hubble_v1_name_registry_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_hubble_v1_name_registry_event_proto_rawDescData)
	})
	return file_hubble_v1_name_registry_event_proto_rawDescData
}

var file_hubble_v1_name_registry_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hubble_v1_name_registry_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hubble_v1_name_registry_event_proto_goTypes = []interface{}{
	(NameRegistryEventType)(0), // 0: hubble.v1.NameRegistryEventType
	(*NameRegistryEvent)(nil),  // 1: hubble.v1.NameRegistryEvent
}
var file_hubble_v1_name_registry_event_proto_depIdxs = []int32{
	0, // 0: hubble.v1.NameRegistryEvent.type:type_name -> hubble.v1.NameRegistryEventType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hubble_v1_name_registry_event_proto_init() }
func file_hubble_v1_name_registry_event_proto_init() {
	if File_hubble_v1_name_registry_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hubble_v1_name_registry_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameRegistryEvent); i {
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
			RawDescriptor: file_hubble_v1_name_registry_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hubble_v1_name_registry_event_proto_goTypes,
		DependencyIndexes: file_hubble_v1_name_registry_event_proto_depIdxs,
		EnumInfos:         file_hubble_v1_name_registry_event_proto_enumTypes,
		MessageInfos:      file_hubble_v1_name_registry_event_proto_msgTypes,
	}.Build()
	File_hubble_v1_name_registry_event_proto = out.File
	file_hubble_v1_name_registry_event_proto_rawDesc = nil
	file_hubble_v1_name_registry_event_proto_goTypes = nil
	file_hubble_v1_name_registry_event_proto_depIdxs = nil
}
