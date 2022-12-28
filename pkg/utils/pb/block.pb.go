// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: block.proto

package pb

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

type BlockInfo_BlockType int32

const (
	BlockInfo_UNKNOWN BlockInfo_BlockType = 0
	BlockInfo_DATA    BlockInfo_BlockType = 1
	BlockInfo_HEADER  BlockInfo_BlockType = 2
)

// Enum value maps for BlockInfo_BlockType.
var (
	BlockInfo_BlockType_name = map[int32]string{
		0: "UNKNOWN",
		1: "DATA",
		2: "HEADER",
	}
	BlockInfo_BlockType_value = map[string]int32{
		"UNKNOWN": 0,
		"DATA":    1,
		"HEADER":  2,
	}
)

func (x BlockInfo_BlockType) Enum() *BlockInfo_BlockType {
	p := new(BlockInfo_BlockType)
	*p = x
	return p
}

func (x BlockInfo_BlockType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockInfo_BlockType) Descriptor() protoreflect.EnumDescriptor {
	return file_block_proto_enumTypes[0].Descriptor()
}

func (BlockInfo_BlockType) Type() protoreflect.EnumType {
	return &file_block_proto_enumTypes[0]
}

func (x BlockInfo_BlockType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockInfo_BlockType.Descriptor instead.
func (BlockInfo_BlockType) EnumDescriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{0, 0}
}

type BlockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId   string              `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Env       string              `protobuf:"bytes,2,opt,name=env,proto3" json:"env,omitempty"`
	BlockNum  int64               `protobuf:"varint,3,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
	BlockHash string              `protobuf:"bytes,4,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	MsgOffset int64               `protobuf:"varint,5,opt,name=msg_offset,json=msgOffset,proto3" json:"msg_offset,omitempty"`
	BlockType BlockInfo_BlockType `protobuf:"varint,6,opt,name=block_type,json=blockType,proto3,enum=pb.BlockInfo_BlockType" json:"block_type,omitempty"`
	BlockSize int64               `protobuf:"varint,7,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
}

func (x *BlockInfo) Reset() {
	*x = BlockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockInfo) ProtoMessage() {}

func (x *BlockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockInfo.ProtoReflect.Descriptor instead.
func (*BlockInfo) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{0}
}

func (x *BlockInfo) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *BlockInfo) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *BlockInfo) GetBlockNum() int64 {
	if x != nil {
		return x.BlockNum
	}
	return 0
}

func (x *BlockInfo) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *BlockInfo) GetMsgOffset() int64 {
	if x != nil {
		return x.MsgOffset
	}
	return 0
}

func (x *BlockInfo) GetBlockType() BlockInfo_BlockType {
	if x != nil {
		return x.BlockType
	}
	return BlockInfo_UNKNOWN
}

func (x *BlockInfo) GetBlockSize() int64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

type BatchItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BatchItem) Reset() {
	*x = BatchItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchItem) ProtoMessage() {}

func (x *BatchItem) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchItem.ProtoReflect.Descriptor instead.
func (*BatchItem) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{1}
}

func (x *BatchItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BatchItem) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info       *BlockInfo   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	BatchItems []*BatchItem `protobuf:"bytes,2,rep,name=batch_items,json=batchItems,proto3" json:"batch_items,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{2}
}

func (x *Block) GetInfo() *BlockInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Block) GetBatchItems() []*BatchItem {
	if x != nil {
		return x.BatchItems
	}
	return nil
}

type DBInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DbType string `protobuf:"bytes,2,opt,name=db_type,json=dbType,proto3" json:"db_type,omitempty"`
	DbPath string `protobuf:"bytes,3,opt,name=db_path,json=dbPath,proto3" json:"db_path,omitempty"`
	IsMeta bool   `protobuf:"varint,4,opt,name=is_meta,json=isMeta,proto3" json:"is_meta,omitempty"`
}

func (x *DBInfo) Reset() {
	*x = DBInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBInfo) ProtoMessage() {}

func (x *DBInfo) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBInfo.ProtoReflect.Descriptor instead.
func (*DBInfo) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{3}
}

func (x *DBInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DBInfo) GetDbType() string {
	if x != nil {
		return x.DbType
	}
	return ""
}

func (x *DBInfo) GetDbPath() string {
	if x != nil {
		return x.DbPath
	}
	return ""
}

func (x *DBInfo) GetIsMeta() bool {
	if x != nil {
		return x.IsMeta
	}
	return false
}

type DBInfoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DbInfos []*DBInfo `protobuf:"bytes,1,rep,name=db_infos,json=dbInfos,proto3" json:"db_infos,omitempty"`
}

func (x *DBInfoList) Reset() {
	*x = DBInfoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBInfoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBInfoList) ProtoMessage() {}

func (x *DBInfoList) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBInfoList.ProtoReflect.Descriptor instead.
func (*DBInfoList) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{4}
}

func (x *DBInfoList) GetDbInfos() []*DBInfo {
	if x != nil {
		return x.DbInfos
	}
	return nil
}

var File_block_proto protoreflect.FileDescriptor

var file_block_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x9a, 0x02, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e,
	0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x5f,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x73,
	0x67, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x62,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x2e,
	0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x41,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x02, 0x22, 0x2f,
	0x0a, 0x09, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x5a, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x21, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x0b, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x63, 0x0a, 0x06, 0x44,
	0x42, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x64, 0x62, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x62, 0x50, 0x61, 0x74, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4d, 0x65, 0x74, 0x61,
	0x22, 0x33, 0x0a, 0x0a, 0x44, 0x42, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x08, 0x64, 0x62, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x42, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x64, 0x62,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x74, 0x69,
	0x6c, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_block_proto_rawDescOnce sync.Once
	file_block_proto_rawDescData = file_block_proto_rawDesc
)

func file_block_proto_rawDescGZIP() []byte {
	file_block_proto_rawDescOnce.Do(func() {
		file_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_block_proto_rawDescData)
	})
	return file_block_proto_rawDescData
}

var file_block_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_block_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_block_proto_goTypes = []interface{}{
	(BlockInfo_BlockType)(0), // 0: pb.BlockInfo.BlockType
	(*BlockInfo)(nil),        // 1: pb.BlockInfo
	(*BatchItem)(nil),        // 2: pb.BatchItem
	(*Block)(nil),            // 3: pb.Block
	(*DBInfo)(nil),           // 4: pb.DBInfo
	(*DBInfoList)(nil),       // 5: pb.DBInfoList
}
var file_block_proto_depIdxs = []int32{
	0, // 0: pb.BlockInfo.block_type:type_name -> pb.BlockInfo.BlockType
	1, // 1: pb.Block.info:type_name -> pb.BlockInfo
	2, // 2: pb.Block.batch_items:type_name -> pb.BatchItem
	4, // 3: pb.DBInfoList.db_infos:type_name -> pb.DBInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_block_proto_init() }
func file_block_proto_init() {
	if File_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockInfo); i {
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
		file_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchItem); i {
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
		file_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBInfo); i {
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
		file_block_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBInfoList); i {
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
			RawDescriptor: file_block_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_block_proto_goTypes,
		DependencyIndexes: file_block_proto_depIdxs,
		EnumInfos:         file_block_proto_enumTypes,
		MessageInfos:      file_block_proto_msgTypes,
	}.Build()
	File_block_proto = out.File
	file_block_proto_rawDesc = nil
	file_block_proto_goTypes = nil
	file_block_proto_depIdxs = nil
}
