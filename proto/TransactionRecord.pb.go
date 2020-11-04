// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/TransactionRecord.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Response when the client sends the node TransactionGetRecordResponse
type TransactionRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Receipt            *TransactionReceipt `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`                       // The status (reach consensus, or failed, or is unknown) and the ID of any new account/file/instance created.
	TransactionHash    []byte              `protobuf:"bytes,2,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`       // The hash of the Transaction that executed (not the hash of any Transaction that failed for having a duplicate TransactionID)
	ConsensusTimestamp *Timestamp          `protobuf:"bytes,3,opt,name=consensusTimestamp,proto3" json:"consensusTimestamp,omitempty"` // The consensus timestamp (or null if didn't reach consensus yet)
	TransactionID      *TransactionID      `protobuf:"bytes,4,opt,name=transactionID,proto3" json:"transactionID,omitempty"`           // The ID of the transaction this record represents
	Memo               string              `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`                             // The memo that was submitted as part of the transaction (max 100 bytes)
	TransactionFee     uint64              `protobuf:"varint,6,opt,name=transactionFee,proto3" json:"transactionFee,omitempty"`        // The actual transaction fee charged, not the original transactionFee value from TransactionBody
	// Types that are assignable to Body:
	//	*TransactionRecord_ContractCallResult
	//	*TransactionRecord_ContractCreateResult
	Body               isTransactionRecord_Body `protobuf_oneof:"body"`
	TransferList       *TransferList            `protobuf:"bytes,10,opt,name=transferList,proto3" json:"transferList,omitempty"`             // All hbar transfers as a result of this transaction, such as fees, or transfers performed by the transaction, or by a smart contract it calls, or by the creation of threshold records that it triggers.
	TokenTransferLists []*TokenTransferList     `protobuf:"bytes,11,rep,name=tokenTransferLists,proto3" json:"tokenTransferLists,omitempty"` // All Token transfers as a result of this transaction
}

func (x *TransactionRecord) Reset() {
	*x = TransactionRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_TransactionRecord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRecord) ProtoMessage() {}

func (x *TransactionRecord) ProtoReflect() protoreflect.Message {
	mi := &file_proto_TransactionRecord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRecord.ProtoReflect.Descriptor instead.
func (*TransactionRecord) Descriptor() ([]byte, []int) {
	return file_proto_TransactionRecord_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionRecord) GetReceipt() *TransactionReceipt {
	if x != nil {
		return x.Receipt
	}
	return nil
}

func (x *TransactionRecord) GetTransactionHash() []byte {
	if x != nil {
		return x.TransactionHash
	}
	return nil
}

func (x *TransactionRecord) GetConsensusTimestamp() *Timestamp {
	if x != nil {
		return x.ConsensusTimestamp
	}
	return nil
}

func (x *TransactionRecord) GetTransactionID() *TransactionID {
	if x != nil {
		return x.TransactionID
	}
	return nil
}

func (x *TransactionRecord) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *TransactionRecord) GetTransactionFee() uint64 {
	if x != nil {
		return x.TransactionFee
	}
	return 0
}

func (m *TransactionRecord) GetBody() isTransactionRecord_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *TransactionRecord) GetContractCallResult() *ContractFunctionResult {
	if x, ok := x.GetBody().(*TransactionRecord_ContractCallResult); ok {
		return x.ContractCallResult
	}
	return nil
}

func (x *TransactionRecord) GetContractCreateResult() *ContractFunctionResult {
	if x, ok := x.GetBody().(*TransactionRecord_ContractCreateResult); ok {
		return x.ContractCreateResult
	}
	return nil
}

func (x *TransactionRecord) GetTransferList() *TransferList {
	if x != nil {
		return x.TransferList
	}
	return nil
}

func (x *TransactionRecord) GetTokenTransferLists() []*TokenTransferList {
	if x != nil {
		return x.TokenTransferLists
	}
	return nil
}

type isTransactionRecord_Body interface {
	isTransactionRecord_Body()
}

type TransactionRecord_ContractCallResult struct {
	ContractCallResult *ContractFunctionResult `protobuf:"bytes,7,opt,name=contractCallResult,proto3,oneof"` // Record of the value returned by the smart contract function (if it completed and didn't fail) from ContractCallTransaction
}

type TransactionRecord_ContractCreateResult struct {
	ContractCreateResult *ContractFunctionResult `protobuf:"bytes,8,opt,name=contractCreateResult,proto3,oneof"` // Record of the value returned by the smart contract constructor (if it completed and didn't fail) from ContractCreateTransaction
}

func (*TransactionRecord_ContractCallResult) isTransactionRecord_Body() {}

func (*TransactionRecord_ContractCreateResult) isTransactionRecord_Body() {}

var File_proto_TransactionRecord_proto protoreflect.FileDescriptor

var file_proto_TransactionRecord_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x04, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12,
	0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x40, 0x0a, 0x12, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73,
	0x75, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3a, 0x0a, 0x0d, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x65, 0x65, 0x12, 0x4f, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43,
	0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00,
	0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x53, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x48, 0x00, 0x52, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x37, 0x0a, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x48, 0x0a, 0x12, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x12, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x42, 0x06, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x42, 0x50, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_TransactionRecord_proto_rawDescOnce sync.Once
	file_proto_TransactionRecord_proto_rawDescData = file_proto_TransactionRecord_proto_rawDesc
)

func file_proto_TransactionRecord_proto_rawDescGZIP() []byte {
	file_proto_TransactionRecord_proto_rawDescOnce.Do(func() {
		file_proto_TransactionRecord_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_TransactionRecord_proto_rawDescData)
	})
	return file_proto_TransactionRecord_proto_rawDescData
}

var file_proto_TransactionRecord_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_TransactionRecord_proto_goTypes = []interface{}{
	(*TransactionRecord)(nil),      // 0: proto.TransactionRecord
	(*TransactionReceipt)(nil),     // 1: proto.TransactionReceipt
	(*Timestamp)(nil),              // 2: proto.Timestamp
	(*TransactionID)(nil),          // 3: proto.TransactionID
	(*ContractFunctionResult)(nil), // 4: proto.ContractFunctionResult
	(*TransferList)(nil),           // 5: proto.TransferList
	(*TokenTransferList)(nil),      // 6: proto.TokenTransferList
}
var file_proto_TransactionRecord_proto_depIdxs = []int32{
	1, // 0: proto.TransactionRecord.receipt:type_name -> proto.TransactionReceipt
	2, // 1: proto.TransactionRecord.consensusTimestamp:type_name -> proto.Timestamp
	3, // 2: proto.TransactionRecord.transactionID:type_name -> proto.TransactionID
	4, // 3: proto.TransactionRecord.contractCallResult:type_name -> proto.ContractFunctionResult
	4, // 4: proto.TransactionRecord.contractCreateResult:type_name -> proto.ContractFunctionResult
	5, // 5: proto.TransactionRecord.transferList:type_name -> proto.TransferList
	6, // 6: proto.TransactionRecord.tokenTransferLists:type_name -> proto.TokenTransferList
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_TransactionRecord_proto_init() }
func file_proto_TransactionRecord_proto_init() {
	if File_proto_TransactionRecord_proto != nil {
		return
	}
	file_proto_Timestamp_proto_init()
	file_proto_BasicTypes_proto_init()
	file_proto_TransactionReceipt_proto_init()
	file_proto_ContractCallLocal_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_TransactionRecord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRecord); i {
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
	file_proto_TransactionRecord_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TransactionRecord_ContractCallResult)(nil),
		(*TransactionRecord_ContractCreateResult)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_TransactionRecord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_TransactionRecord_proto_goTypes,
		DependencyIndexes: file_proto_TransactionRecord_proto_depIdxs,
		MessageInfos:      file_proto_TransactionRecord_proto_msgTypes,
	}.Build()
	File_proto_TransactionRecord_proto = out.File
	file_proto_TransactionRecord_proto_rawDesc = nil
	file_proto_TransactionRecord_proto_goTypes = nil
	file_proto_TransactionRecord_proto_depIdxs = nil
}
