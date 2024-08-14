// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: sync_trie.proto

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

type DbTrieNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ChildChars []uint32 `protobuf:"varint,2,rep,packed,name=childChars,proto3" json:"childChars,omitempty"`
	Items      uint32   `protobuf:"varint,3,opt,name=items,proto3" json:"items,omitempty"`
	Hash       []byte   `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *DbTrieNode) Reset() {
	*x = DbTrieNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_trie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DbTrieNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbTrieNode) ProtoMessage() {}

func (x *DbTrieNode) ProtoReflect() protoreflect.Message {
	mi := &file_sync_trie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbTrieNode.ProtoReflect.Descriptor instead.
func (*DbTrieNode) Descriptor() ([]byte, []int) {
	return file_sync_trie_proto_rawDescGZIP(), []int{0}
}

func (x *DbTrieNode) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *DbTrieNode) GetChildChars() []uint32 {
	if x != nil {
		return x.ChildChars
	}
	return nil
}

func (x *DbTrieNode) GetItems() uint32 {
	if x != nil {
		return x.Items
	}
	return 0
}

func (x *DbTrieNode) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

var File_sync_trie_proto protoreflect.FileDescriptor

var file_sync_trie_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x74, 0x72, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x68, 0x0a, 0x0a, 0x44, 0x62, 0x54, 0x72, 0x69, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x43, 0x68, 0x61, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x43, 0x68, 0x61, 0x72,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x42, 0x23, 0x5a, 0x21, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x69, 0x63, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sync_trie_proto_rawDescOnce sync.Once
	file_sync_trie_proto_rawDescData = file_sync_trie_proto_rawDesc
)

func file_sync_trie_proto_rawDescGZIP() []byte {
	file_sync_trie_proto_rawDescOnce.Do(func() {
		file_sync_trie_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_trie_proto_rawDescData)
	})
	return file_sync_trie_proto_rawDescData
}

var file_sync_trie_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sync_trie_proto_goTypes = []any{
	(*DbTrieNode)(nil), // 0: DbTrieNode
}
var file_sync_trie_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sync_trie_proto_init() }
func file_sync_trie_proto_init() {
	if File_sync_trie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sync_trie_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DbTrieNode); i {
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
			RawDescriptor: file_sync_trie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sync_trie_proto_goTypes,
		DependencyIndexes: file_sync_trie_proto_depIdxs,
		MessageInfos:      file_sync_trie_proto_msgTypes,
	}.Build()
	File_sync_trie_proto = out.File
	file_sync_trie_proto_rawDesc = nil
	file_sync_trie_proto_goTypes = nil
	file_sync_trie_proto_depIdxs = nil
}