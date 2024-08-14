// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: gossip.proto

package hubble_grpc

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
	return file_gossip_proto_enumTypes[0].Descriptor()
}

func (GossipVersion) Type() protoreflect.EnumType {
	return &file_gossip_proto_enumTypes[0]
}

func (x GossipVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GossipVersion.Descriptor instead.
func (GossipVersion) EnumDescriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{0}
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
		mi := &file_gossip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipAddressInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipAddressInfo) ProtoMessage() {}

func (x *GossipAddressInfo) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[0]
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
	return file_gossip_proto_rawDescGZIP(), []int{0}
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

type ContactInfoContentBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GossipAddress  *GossipAddressInfo `protobuf:"bytes,1,opt,name=gossip_address,json=gossipAddress,proto3" json:"gossip_address,omitempty"`
	RpcAddress     *GossipAddressInfo `protobuf:"bytes,2,opt,name=rpc_address,json=rpcAddress,proto3" json:"rpc_address,omitempty"`
	ExcludedHashes []string           `protobuf:"bytes,3,rep,name=excluded_hashes,json=excludedHashes,proto3" json:"excluded_hashes,omitempty"`
	Count          uint32             `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	HubVersion     string             `protobuf:"bytes,5,opt,name=hub_version,json=hubVersion,proto3" json:"hub_version,omitempty"`
	Network        FarcasterNetwork   `protobuf:"varint,6,opt,name=network,proto3,enum=FarcasterNetwork" json:"network,omitempty"`
	AppVersion     string             `protobuf:"bytes,7,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	Timestamp      uint64             `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ContactInfoContentBody) Reset() {
	*x = ContactInfoContentBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactInfoContentBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactInfoContentBody) ProtoMessage() {}

func (x *ContactInfoContentBody) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactInfoContentBody.ProtoReflect.Descriptor instead.
func (*ContactInfoContentBody) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{1}
}

func (x *ContactInfoContentBody) GetGossipAddress() *GossipAddressInfo {
	if x != nil {
		return x.GossipAddress
	}
	return nil
}

func (x *ContactInfoContentBody) GetRpcAddress() *GossipAddressInfo {
	if x != nil {
		return x.RpcAddress
	}
	return nil
}

func (x *ContactInfoContentBody) GetExcludedHashes() []string {
	if x != nil {
		return x.ExcludedHashes
	}
	return nil
}

func (x *ContactInfoContentBody) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ContactInfoContentBody) GetHubVersion() string {
	if x != nil {
		return x.HubVersion
	}
	return ""
}

func (x *ContactInfoContentBody) GetNetwork() FarcasterNetwork {
	if x != nil {
		return x.Network
	}
	return FarcasterNetwork_FARCASTER_NETWORK_NONE
}

func (x *ContactInfoContentBody) GetAppVersion() string {
	if x != nil {
		return x.AppVersion
	}
	return ""
}

func (x *ContactInfoContentBody) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ContactInfoContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GossipAddress  *GossipAddressInfo      `protobuf:"bytes,1,opt,name=gossip_address,json=gossipAddress,proto3" json:"gossip_address,omitempty"`
	RpcAddress     *GossipAddressInfo      `protobuf:"bytes,2,opt,name=rpc_address,json=rpcAddress,proto3" json:"rpc_address,omitempty"`
	ExcludedHashes []string                `protobuf:"bytes,3,rep,name=excluded_hashes,json=excludedHashes,proto3" json:"excluded_hashes,omitempty"`
	Count          uint32                  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	HubVersion     string                  `protobuf:"bytes,5,opt,name=hub_version,json=hubVersion,proto3" json:"hub_version,omitempty"`
	Network        FarcasterNetwork        `protobuf:"varint,6,opt,name=network,proto3,enum=FarcasterNetwork" json:"network,omitempty"`
	AppVersion     string                  `protobuf:"bytes,7,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	Timestamp      uint64                  `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Body           *ContactInfoContentBody `protobuf:"bytes,9,opt,name=body,proto3" json:"body,omitempty"`
	Signature      []byte                  `protobuf:"bytes,10,opt,name=signature,proto3" json:"signature,omitempty"`                        // Signature of the message digest
	Signer         []byte                  `protobuf:"bytes,11,opt,name=signer,proto3" json:"signer,omitempty"`                              // Public key of the peer that originated the contact info
	DataBytes      []byte                  `protobuf:"bytes,12,opt,name=data_bytes,json=dataBytes,proto3,oneof" json:"data_bytes,omitempty"` // Optional alternative serialization used for signing
}

func (x *ContactInfoContent) Reset() {
	*x = ContactInfoContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactInfoContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactInfoContent) ProtoMessage() {}

func (x *ContactInfoContent) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[2]
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
	return file_gossip_proto_rawDescGZIP(), []int{2}
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

func (x *ContactInfoContent) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ContactInfoContent) GetBody() *ContactInfoContentBody {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *ContactInfoContent) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *ContactInfoContent) GetSigner() []byte {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *ContactInfoContent) GetDataBytes() []byte {
	if x != nil {
		return x.DataBytes
	}
	return nil
}

type PingMessageBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PingOriginPeerId []byte `protobuf:"bytes,1,opt,name=ping_origin_peer_id,json=pingOriginPeerId,proto3" json:"ping_origin_peer_id,omitempty"`
	PingTimestamp    uint64 `protobuf:"varint,2,opt,name=ping_timestamp,json=pingTimestamp,proto3" json:"ping_timestamp,omitempty"`
}

func (x *PingMessageBody) Reset() {
	*x = PingMessageBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingMessageBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingMessageBody) ProtoMessage() {}

func (x *PingMessageBody) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingMessageBody.ProtoReflect.Descriptor instead.
func (*PingMessageBody) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{3}
}

func (x *PingMessageBody) GetPingOriginPeerId() []byte {
	if x != nil {
		return x.PingOriginPeerId
	}
	return nil
}

func (x *PingMessageBody) GetPingTimestamp() uint64 {
	if x != nil {
		return x.PingTimestamp
	}
	return 0
}

type AckMessageBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PingOriginPeerId []byte `protobuf:"bytes,1,opt,name=ping_origin_peer_id,json=pingOriginPeerId,proto3" json:"ping_origin_peer_id,omitempty"`
	AckOriginPeerId  []byte `protobuf:"bytes,2,opt,name=ack_origin_peer_id,json=ackOriginPeerId,proto3" json:"ack_origin_peer_id,omitempty"`
	PingTimestamp    uint64 `protobuf:"varint,3,opt,name=ping_timestamp,json=pingTimestamp,proto3" json:"ping_timestamp,omitempty"`
	AckTimestamp     uint64 `protobuf:"varint,4,opt,name=ack_timestamp,json=ackTimestamp,proto3" json:"ack_timestamp,omitempty"`
}

func (x *AckMessageBody) Reset() {
	*x = AckMessageBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckMessageBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckMessageBody) ProtoMessage() {}

func (x *AckMessageBody) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckMessageBody.ProtoReflect.Descriptor instead.
func (*AckMessageBody) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{4}
}

func (x *AckMessageBody) GetPingOriginPeerId() []byte {
	if x != nil {
		return x.PingOriginPeerId
	}
	return nil
}

func (x *AckMessageBody) GetAckOriginPeerId() []byte {
	if x != nil {
		return x.AckOriginPeerId
	}
	return nil
}

func (x *AckMessageBody) GetPingTimestamp() uint64 {
	if x != nil {
		return x.PingTimestamp
	}
	return 0
}

func (x *AckMessageBody) GetAckTimestamp() uint64 {
	if x != nil {
		return x.AckTimestamp
	}
	return 0
}

type NetworkLatencyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//
	//	*NetworkLatencyMessage_PingMessage
	//	*NetworkLatencyMessage_AckMessage
	Body isNetworkLatencyMessage_Body `protobuf_oneof:"body"`
}

func (x *NetworkLatencyMessage) Reset() {
	*x = NetworkLatencyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkLatencyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkLatencyMessage) ProtoMessage() {}

func (x *NetworkLatencyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkLatencyMessage.ProtoReflect.Descriptor instead.
func (*NetworkLatencyMessage) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{5}
}

func (m *NetworkLatencyMessage) GetBody() isNetworkLatencyMessage_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *NetworkLatencyMessage) GetPingMessage() *PingMessageBody {
	if x, ok := x.GetBody().(*NetworkLatencyMessage_PingMessage); ok {
		return x.PingMessage
	}
	return nil
}

func (x *NetworkLatencyMessage) GetAckMessage() *AckMessageBody {
	if x, ok := x.GetBody().(*NetworkLatencyMessage_AckMessage); ok {
		return x.AckMessage
	}
	return nil
}

type isNetworkLatencyMessage_Body interface {
	isNetworkLatencyMessage_Body()
}

type NetworkLatencyMessage_PingMessage struct {
	PingMessage *PingMessageBody `protobuf:"bytes,2,opt,name=ping_message,json=pingMessage,proto3,oneof"`
}

type NetworkLatencyMessage_AckMessage struct {
	AckMessage *AckMessageBody `protobuf:"bytes,3,opt,name=ack_message,json=ackMessage,proto3,oneof"`
}

func (*NetworkLatencyMessage_PingMessage) isNetworkLatencyMessage_Body() {}

func (*NetworkLatencyMessage_AckMessage) isNetworkLatencyMessage_Body() {}

type MessageBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash     []byte     `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Messages []*Message `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *MessageBundle) Reset() {
	*x = MessageBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageBundle) ProtoMessage() {}

func (x *MessageBundle) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageBundle.ProtoReflect.Descriptor instead.
func (*MessageBundle) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{6}
}

func (x *MessageBundle) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *MessageBundle) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type GossipMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//
	//	*GossipMessage_Message
	//	*GossipMessage_ContactInfoContent
	//	*GossipMessage_NetworkLatencyMessage
	//	*GossipMessage_MessageBundle
	Content   isGossipMessage_Content `protobuf_oneof:"content"`
	Topics    []string                `protobuf:"bytes,4,rep,name=topics,proto3" json:"topics,omitempty"`
	PeerId    []byte                  `protobuf:"bytes,5,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Version   GossipVersion           `protobuf:"varint,6,opt,name=version,proto3,enum=GossipVersion" json:"version,omitempty"`
	Timestamp uint32                  `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // Farcaster epoch timestamp in seconds when this message was first created
}

func (x *GossipMessage) Reset() {
	*x = GossipMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipMessage) ProtoMessage() {}

func (x *GossipMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[7]
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
	return file_gossip_proto_rawDescGZIP(), []int{7}
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

func (x *GossipMessage) GetContactInfoContent() *ContactInfoContent {
	if x, ok := x.GetContent().(*GossipMessage_ContactInfoContent); ok {
		return x.ContactInfoContent
	}
	return nil
}

func (x *GossipMessage) GetNetworkLatencyMessage() *NetworkLatencyMessage {
	if x, ok := x.GetContent().(*GossipMessage_NetworkLatencyMessage); ok {
		return x.NetworkLatencyMessage
	}
	return nil
}

func (x *GossipMessage) GetMessageBundle() *MessageBundle {
	if x, ok := x.GetContent().(*GossipMessage_MessageBundle); ok {
		return x.MessageBundle
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

func (x *GossipMessage) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type isGossipMessage_Content interface {
	isGossipMessage_Content()
}

type GossipMessage_Message struct {
	Message *Message `protobuf:"bytes,1,opt,name=message,proto3,oneof"`
}

type GossipMessage_ContactInfoContent struct {
	// Deprecated
	// IdRegistryEvent id_registry_event = 2;
	ContactInfoContent *ContactInfoContent `protobuf:"bytes,3,opt,name=contact_info_content,json=contactInfoContent,proto3,oneof"`
}

type GossipMessage_NetworkLatencyMessage struct {
	NetworkLatencyMessage *NetworkLatencyMessage `protobuf:"bytes,7,opt,name=network_latency_message,json=networkLatencyMessage,proto3,oneof"`
}

type GossipMessage_MessageBundle struct {
	MessageBundle *MessageBundle `protobuf:"bytes,9,opt,name=message_bundle,json=messageBundle,proto3,oneof"`
}

func (*GossipMessage_Message) isGossipMessage_Content() {}

func (*GossipMessage_ContactInfoContent) isGossipMessage_Content() {}

func (*GossipMessage_NetworkLatencyMessage) isGossipMessage_Content() {}

func (*GossipMessage_MessageBundle) isGossipMessage_Content() {}

var File_gossip_proto protoreflect.FileDescriptor

var file_gossip_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a,
	0x11, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x6e, 0x73, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6e, 0x73, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0xd4, 0x02, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x39,
	0x0a, 0x0e, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x67, 0x6f, 0x73, 0x73,
	0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x0b, 0x72, 0x70, 0x63,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x72, 0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x27,
	0x0a, 0x0f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x64, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x68, 0x75, 0x62, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x68, 0x75, 0x62, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b,
	0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x11, 0x2e, 0x46, 0x61, 0x72, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x70, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xe6, 0x03, 0x0a, 0x12, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x39, 0x0a, 0x0e, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x47, 0x6f, 0x73, 0x73,
	0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x67,
	0x6f, 0x73, 0x73, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x0b,
	0x72, 0x70, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x72, 0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x64, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x75, 0x62, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x75, 0x62, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x2b, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x46, 0x61, 0x72, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2b, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72,
	0x12, 0x22, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x22, 0x67, 0x0a, 0x0f, 0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x2d, 0x0a, 0x13, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x10, 0x70, 0x69, 0x6e, 0x67, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x50,
	0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x70,
	0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xb8, 0x01, 0x0a,
	0x0e, 0x41, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12,
	0x2d, 0x0a, 0x13, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x70,
	0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x70, 0x69,
	0x6e, 0x67, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b,
	0x0a, 0x12, 0x61, 0x63, 0x6b, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x65, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x61, 0x63, 0x6b, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70,
	0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0d, 0x70, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x61, 0x63, 0x6b, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x8a, 0x01, 0x0a, 0x15, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x35, 0x0a, 0x0c, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x69, 0x6e,
	0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x61, 0x63, 0x6b, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x41, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x48, 0x00,
	0x52, 0x0a, 0x61, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x22, 0x49, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x08, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22,
	0x8d, 0x03, 0x0a, 0x0d, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x24, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x12, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x50, 0x0a, 0x17, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6c, 0x61, 0x74, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4c, 0x61, 0x74, 0x65, 0x6e,
	0x63, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x15, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x37, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e,
	0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2a,
	0x3f, 0x0a, 0x0d, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x15, 0x0a, 0x11, 0x47, 0x4f, 0x53, 0x53, 0x49, 0x50, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x56, 0x31, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x4f, 0x53, 0x53, 0x49,
	0x50, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x31, 0x5f, 0x31, 0x10, 0x01,
	0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a,
	0x75, 0x69, 0x63, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gossip_proto_rawDescOnce sync.Once
	file_gossip_proto_rawDescData = file_gossip_proto_rawDesc
)

func file_gossip_proto_rawDescGZIP() []byte {
	file_gossip_proto_rawDescOnce.Do(func() {
		file_gossip_proto_rawDescData = protoimpl.X.CompressGZIP(file_gossip_proto_rawDescData)
	})
	return file_gossip_proto_rawDescData
}

var file_gossip_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gossip_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_gossip_proto_goTypes = []any{
	(GossipVersion)(0),             // 0: GossipVersion
	(*GossipAddressInfo)(nil),      // 1: GossipAddressInfo
	(*ContactInfoContentBody)(nil), // 2: ContactInfoContentBody
	(*ContactInfoContent)(nil),     // 3: ContactInfoContent
	(*PingMessageBody)(nil),        // 4: PingMessageBody
	(*AckMessageBody)(nil),         // 5: AckMessageBody
	(*NetworkLatencyMessage)(nil),  // 6: NetworkLatencyMessage
	(*MessageBundle)(nil),          // 7: MessageBundle
	(*GossipMessage)(nil),          // 8: GossipMessage
	(FarcasterNetwork)(0),          // 9: FarcasterNetwork
	(*Message)(nil),                // 10: Message
}
var file_gossip_proto_depIdxs = []int32{
	1,  // 0: ContactInfoContentBody.gossip_address:type_name -> GossipAddressInfo
	1,  // 1: ContactInfoContentBody.rpc_address:type_name -> GossipAddressInfo
	9,  // 2: ContactInfoContentBody.network:type_name -> FarcasterNetwork
	1,  // 3: ContactInfoContent.gossip_address:type_name -> GossipAddressInfo
	1,  // 4: ContactInfoContent.rpc_address:type_name -> GossipAddressInfo
	9,  // 5: ContactInfoContent.network:type_name -> FarcasterNetwork
	2,  // 6: ContactInfoContent.body:type_name -> ContactInfoContentBody
	4,  // 7: NetworkLatencyMessage.ping_message:type_name -> PingMessageBody
	5,  // 8: NetworkLatencyMessage.ack_message:type_name -> AckMessageBody
	10, // 9: MessageBundle.messages:type_name -> Message
	10, // 10: GossipMessage.message:type_name -> Message
	3,  // 11: GossipMessage.contact_info_content:type_name -> ContactInfoContent
	6,  // 12: GossipMessage.network_latency_message:type_name -> NetworkLatencyMessage
	7,  // 13: GossipMessage.message_bundle:type_name -> MessageBundle
	0,  // 14: GossipMessage.version:type_name -> GossipVersion
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_gossip_proto_init() }
func file_gossip_proto_init() {
	if File_gossip_proto != nil {
		return
	}
	file_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gossip_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_gossip_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ContactInfoContentBody); i {
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
		file_gossip_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_gossip_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PingMessageBody); i {
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
		file_gossip_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AckMessageBody); i {
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
		file_gossip_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*NetworkLatencyMessage); i {
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
		file_gossip_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*MessageBundle); i {
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
		file_gossip_proto_msgTypes[7].Exporter = func(v any, i int) any {
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
	file_gossip_proto_msgTypes[2].OneofWrappers = []any{}
	file_gossip_proto_msgTypes[5].OneofWrappers = []any{
		(*NetworkLatencyMessage_PingMessage)(nil),
		(*NetworkLatencyMessage_AckMessage)(nil),
	}
	file_gossip_proto_msgTypes[7].OneofWrappers = []any{
		(*GossipMessage_Message)(nil),
		(*GossipMessage_ContactInfoContent)(nil),
		(*GossipMessage_NetworkLatencyMessage)(nil),
		(*GossipMessage_MessageBundle)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gossip_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gossip_proto_goTypes,
		DependencyIndexes: file_gossip_proto_depIdxs,
		EnumInfos:         file_gossip_proto_enumTypes,
		MessageInfos:      file_gossip_proto_msgTypes,
	}.Build()
	File_gossip_proto = out.File
	file_gossip_proto_rawDesc = nil
	file_gossip_proto_goTypes = nil
	file_gossip_proto_depIdxs = nil
}
