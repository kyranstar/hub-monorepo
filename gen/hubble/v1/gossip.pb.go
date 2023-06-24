// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: hubble/v1/gossip.proto

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

type GossipVersion int32

const (
	GossipVersion_GOSSIP_VERSION_V1   GossipVersion = 0
	GossipVersion_GOSSIP_VERSION_V1_1 GossipVersion = 1
)

// Enum value maps for GossipVersion.
var (
	GossipVersion_name = map[int32]string{
		0: "GOSSIP_VERSION_V1",
		1: "GOSSIP_VERSION_V1_1",
	}
	GossipVersion_value = map[string]int32{
		"GOSSIP_VERSION_V1":   0,
		"GOSSIP_VERSION_V1_1": 1,
	}
)

func (x GossipVersion) Enum() *GossipVersion {
	p := new(GossipVersion)
	*p = x
	return p
}

func (x GossipVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GossipVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_hubble_v1_gossip_proto_enumTypes[0].Descriptor()
}

func (GossipVersion) Type() protoreflect.EnumType {
	return &file_hubble_v1_gossip_proto_enumTypes[0]
}

func (x GossipVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GossipVersion.Descriptor instead.
func (GossipVersion) EnumDescriptor() ([]byte, []int) {
	return file_hubble_v1_gossip_proto_rawDescGZIP(), []int{0}
}

type GossipAddressInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Family  uint32 `protobuf:"varint,2,opt,name=family,proto3" json:"family,omitempty"`
	Port    uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	DnsName string `protobuf:"bytes,4,opt,name=dns_name,json=dnsName,proto3" json:"dns_name,omitempty"`
}

func (x *GossipAddressInfo) Reset() {
	*x = GossipAddressInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hubble_v1_gossip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipAddressInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipAddressInfo) ProtoMessage() {}

func (x *GossipAddressInfo) ProtoReflect() protoreflect.Message {
	mi := &file_hubble_v1_gossip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GossipAddressInfo.ProtoReflect.Descriptor instead.
func (*GossipAddressInfo) Descriptor() ([]byte, []int) {
	return file_hubble_v1_gossip_proto_rawDescGZIP(), []int{0}
}

func (x *GossipAddressInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GossipAddressInfo) GetFamily() uint32 {
	if x != nil {
		return x.Family
	}
	return 0
}

func (x *GossipAddressInfo) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *GossipAddressInfo) GetDnsName() string {
	if x != nil {
		return x.DnsName
	}
	return ""
}

type ContactInfoContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GossipAddress  *GossipAddressInfo `protobuf:"bytes,1,opt,name=gossip_address,json=gossipAddress,proto3" json:"gossip_address,omitempty"`
	RpcAddress     *GossipAddressInfo `protobuf:"bytes,2,opt,name=rpc_address,json=rpcAddress,proto3" json:"rpc_address,omitempty"`
	ExcludedHashes []string           `protobuf:"bytes,3,rep,name=excluded_hashes,json=excludedHashes,proto3" json:"excluded_hashes,omitempty"`
	Count          uint32             `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	HubVersion     string             `protobuf:"bytes,5,opt,name=hub_version,json=hubVersion,proto3" json:"hub_version,omitempty"`
	Network        FarcasterNetwork   `protobuf:"varint,6,opt,name=network,proto3,enum=hubble.v1.FarcasterNetwork" json:"network,omitempty"`
	AppVersion     string             `protobuf:"bytes,7,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
}

func (x *ContactInfoContent) Reset() {
	*x = ContactInfoContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hubble_v1_gossip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactInfoContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactInfoContent) ProtoMessage() {}

func (x *ContactInfoContent) ProtoReflect() protoreflect.Message {
	mi := &file_hubble_v1_gossip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactInfoContent.ProtoReflect.Descriptor instead.
func (*ContactInfoContent) Descriptor() ([]byte, []int) {
	return file_hubble_v1_gossip_proto_rawDescGZIP(), []int{1}
}

func (x *ContactInfoContent) GetGossipAddress() *GossipAddressInfo {
	if x != nil {
		return x.GossipAddress
	}
	return nil
}

func (x *ContactInfoContent) GetRpcAddress() *GossipAddressInfo {
	if x != nil {
		return x.RpcAddress
	}
	return nil
}

func (x *ContactInfoContent) GetExcludedHashes() []string {
	if x != nil {
		return x.ExcludedHashes
	}
	return nil
}

func (x *ContactInfoContent) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ContactInfoContent) GetHubVersion() string {
	if x != nil {
		return x.HubVersion
	}
	return ""
}

func (x *ContactInfoContent) GetNetwork() FarcasterNetwork {
	if x != nil {
		return x.Network
	}
	return FarcasterNetwork_FARCASTER_NETWORK_NONE
}

func (x *ContactInfoContent) GetAppVersion() string {
	if x != nil {
		return x.AppVersion
	}
	return ""
}

type GossipMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//
	//	*GossipMessage_Message
	//	*GossipMessage_IdRegistryEvent
	//	*GossipMessage_ContactInfoContent
	Content isGossipMessage_Content `protobuf_oneof:"content"`
	Topics  []string                `protobuf:"bytes,4,rep,name=topics,proto3" json:"topics,omitempty"`
	PeerId  []byte                  `protobuf:"bytes,5,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Version GossipVersion           `protobuf:"varint,6,opt,name=version,proto3,enum=hubble.v1.GossipVersion" json:"version,omitempty"`
}

func (x *GossipMessage) Reset() {
	*x = GossipMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hubble_v1_gossip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipMessage) ProtoMessage() {}

func (x *GossipMessage) ProtoReflect() protoreflect.Message {
	mi := &file_hubble_v1_gossip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GossipMessage.ProtoReflect.Descriptor instead.
func (*GossipMessage) Descriptor() ([]byte, []int) {
	return file_hubble_v1_gossip_proto_rawDescGZIP(), []int{2}
}

func (m *GossipMessage) GetContent() isGossipMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *GossipMessage) GetMessage() *Message {
	if x, ok := x.GetContent().(*GossipMessage_Message); ok {
		return x.Message
	}
	return nil
}

func (x *GossipMessage) GetIdRegistryEvent() *IdRegistryEvent {
	if x, ok := x.GetContent().(*GossipMessage_IdRegistryEvent); ok {
		return x.IdRegistryEvent
	}
	return nil
}

func (x *GossipMessage) GetContactInfoContent() *ContactInfoContent {
	if x, ok := x.GetContent().(*GossipMessage_ContactInfoContent); ok {
		return x.ContactInfoContent
	}
	return nil
}

func (x *GossipMessage) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *GossipMessage) GetPeerId() []byte {
	if x != nil {
		return x.PeerId
	}
	return nil
}

func (x *GossipMessage) GetVersion() GossipVersion {
	if x != nil {
		return x.Version
	}
	return GossipVersion_GOSSIP_VERSION_V1
}

type isGossipMessage_Content interface {
	isGossipMessage_Content()
}

type GossipMessage_Message struct {
	Message *Message `protobuf:"bytes,1,opt,name=message,proto3,oneof"`
}

type GossipMessage_IdRegistryEvent struct {
	IdRegistryEvent *IdRegistryEvent `protobuf:"bytes,2,opt,name=id_registry_event,json=idRegistryEvent,proto3,oneof"`
}

type GossipMessage_ContactInfoContent struct {
	ContactInfoContent *ContactInfoContent `protobuf:"bytes,3,opt,name=contact_info_content,json=contactInfoContent,proto3,oneof"`
}

func (*GossipMessage_Message) isGossipMessage_Content() {}

func (*GossipMessage_IdRegistryEvent) isGossipMessage_Content() {}

func (*GossipMessage_ContactInfoContent) isGossipMessage_Content() {}

var File_hubble_v1_gossip_proto protoreflect.FileDescriptor

var file_hubble_v1_gossip_proto_rawDesc = []byte{
	0x0a, 0x16, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x73, 0x73,
	0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x17, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x68, 0x75,
	0x62, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x74, 0x0a, 0x11, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x6e,
	0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6e,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd0, 0x02, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x0e,
	0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0d, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x72, 0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x64, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x68, 0x75, 0x62, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x75, 0x62, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x35, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61,
	0x72, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70,
	0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xcc, 0x02, 0x0a, 0x0d, 0x47, 0x6f, 0x73,
	0x73, 0x69, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x75,
	0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x48, 0x0a, 0x11, 0x69, 0x64,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x48, 0x00, 0x52, 0x0f, 0x69, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x51, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x48, 0x00, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x68, 0x75, 0x62, 0x62,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2a, 0x3f, 0x0a, 0x0d, 0x47, 0x6f, 0x73, 0x73, 0x69,
	0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x4f, 0x53, 0x53,
	0x49, 0x50, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x31, 0x10, 0x00, 0x12,
	0x17, 0x0a, 0x13, 0x47, 0x4f, 0x53, 0x53, 0x49, 0x50, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x56, 0x31, 0x5f, 0x31, 0x10, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x79, 0x72, 0x61, 0x6e, 0x73, 0x74, 0x61, 0x72,
	0x2f, 0x68, 0x75, 0x62, 0x2d, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hubble_v1_gossip_proto_rawDescOnce sync.Once
	file_hubble_v1_gossip_proto_rawDescData = file_hubble_v1_gossip_proto_rawDesc
)

func file_hubble_v1_gossip_proto_rawDescGZIP() []byte {
	file_hubble_v1_gossip_proto_rawDescOnce.Do(func() {
		file_hubble_v1_gossip_proto_rawDescData = protoimpl.X.CompressGZIP(file_hubble_v1_gossip_proto_rawDescData)
	})
	return file_hubble_v1_gossip_proto_rawDescData
}

var file_hubble_v1_gossip_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hubble_v1_gossip_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_hubble_v1_gossip_proto_goTypes = []interface{}{
	(GossipVersion)(0),         // 0: hubble.v1.GossipVersion
	(*GossipAddressInfo)(nil),  // 1: hubble.v1.GossipAddressInfo
	(*ContactInfoContent)(nil), // 2: hubble.v1.ContactInfoContent
	(*GossipMessage)(nil),      // 3: hubble.v1.GossipMessage
	(FarcasterNetwork)(0),      // 4: hubble.v1.FarcasterNetwork
	(*Message)(nil),            // 5: hubble.v1.Message
	(*IdRegistryEvent)(nil),    // 6: hubble.v1.IdRegistryEvent
}
var file_hubble_v1_gossip_proto_depIdxs = []int32{
	1, // 0: hubble.v1.ContactInfoContent.gossip_address:type_name -> hubble.v1.GossipAddressInfo
	1, // 1: hubble.v1.ContactInfoContent.rpc_address:type_name -> hubble.v1.GossipAddressInfo
	4, // 2: hubble.v1.ContactInfoContent.network:type_name -> hubble.v1.FarcasterNetwork
	5, // 3: hubble.v1.GossipMessage.message:type_name -> hubble.v1.Message
	6, // 4: hubble.v1.GossipMessage.id_registry_event:type_name -> hubble.v1.IdRegistryEvent
	2, // 5: hubble.v1.GossipMessage.contact_info_content:type_name -> hubble.v1.ContactInfoContent
	0, // 6: hubble.v1.GossipMessage.version:type_name -> hubble.v1.GossipVersion
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_hubble_v1_gossip_proto_init() }
func file_hubble_v1_gossip_proto_init() {
	if File_hubble_v1_gossip_proto != nil {
		return
	}
	file_hubble_v1_message_proto_init()
	file_hubble_v1_id_registry_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_hubble_v1_gossip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GossipAddressInfo); i {
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
		file_hubble_v1_gossip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactInfoContent); i {
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
		file_hubble_v1_gossip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GossipMessage); i {
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
	file_hubble_v1_gossip_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*GossipMessage_Message)(nil),
		(*GossipMessage_IdRegistryEvent)(nil),
		(*GossipMessage_ContactInfoContent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hubble_v1_gossip_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hubble_v1_gossip_proto_goTypes,
		DependencyIndexes: file_hubble_v1_gossip_proto_depIdxs,
		EnumInfos:         file_hubble_v1_gossip_proto_enumTypes,
		MessageInfos:      file_hubble_v1_gossip_proto_msgTypes,
	}.Build()
	File_hubble_v1_gossip_proto = out.File
	file_hubble_v1_gossip_proto_rawDesc = nil
	file_hubble_v1_gossip_proto_goTypes = nil
	file_hubble_v1_gossip_proto_depIdxs = nil
}
