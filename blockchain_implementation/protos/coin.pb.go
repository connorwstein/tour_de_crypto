// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coin.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Maybe look at adding the scripting functionality to
// unlock UTXOs
// Each transaction references as input a UTXO from the existing
// set of UTXOs
// But how does it start? Every client knows about the genesis block
// which when mined introduces new coin from thin air (no inputUTXO)
type Transaction struct {
	// UTXO is a transaction hash
	InputUTXO []byte `protobuf:"bytes,1,opt,name=inputUTXO,proto3" json:"inputUTXO,omitempty"`
	// Pub key associated with source account
	// account must be part of the UTXO
	// bytes senderPubKey = 2;
	// Destination of the UTXO
	ReceiverPubKey []byte `protobuf:"bytes,3,opt,name=receiverPubKey,proto3" json:"receiverPubKey,omitempty"`
	// Ensures the private key associated with the public key in the inputUTXO was actually used to sign this
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	Value                uint64   `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_2f53f782266e87bc, []int{0}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetInputUTXO() []byte {
	if m != nil {
		return m.InputUTXO
	}
	return nil
}

func (m *Transaction) GetReceiverPubKey() []byte {
	if m != nil {
		return m.ReceiverPubKey
	}
	return nil
}

func (m *Transaction) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Transaction) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Hash of the block is made by hashing the blockheader twice
type BlockHeader struct {
	PrevBlockHash        []byte   `protobuf:"bytes,1,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	MerkleRoot           []byte   `protobuf:"bytes,2,opt,name=merkleRoot,proto3" json:"merkleRoot,omitempty"`
	TimeStamp            uint32   `protobuf:"varint,3,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	DifficultyTarget     uint32   `protobuf:"varint,4,opt,name=difficultyTarget,proto3" json:"difficultyTarget,omitempty"`
	Nonce                uint32   `protobuf:"varint,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_2f53f782266e87bc, []int{1}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (dst *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(dst, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetPrevBlockHash() []byte {
	if m != nil {
		return m.PrevBlockHash
	}
	return nil
}

func (m *BlockHeader) GetMerkleRoot() []byte {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

func (m *BlockHeader) GetTimeStamp() uint32 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *BlockHeader) GetDifficultyTarget() uint32 {
	if m != nil {
		return m.DifficultyTarget
	}
	return 0
}

func (m *BlockHeader) GetNonce() uint32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_2f53f782266e87bc, []int{2}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Hello struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_2f53f782266e87bc, []int{3}
}
func (m *Hello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hello.Unmarshal(m, b)
}
func (m *Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hello.Marshal(b, m, deterministic)
}
func (dst *Hello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hello.Merge(dst, src)
}
func (m *Hello) XXX_Size() int {
	return xxx_messageInfo_Hello.Size(m)
}
func (m *Hello) XXX_DiscardUnknown() {
	xxx_messageInfo_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_Hello proto.InternalMessageInfo

type Ack struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ack) Reset()         { *m = Ack{} }
func (m *Ack) String() string { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()    {}
func (*Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_coin_2f53f782266e87bc, []int{4}
}
func (m *Ack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ack.Unmarshal(m, b)
}
func (m *Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ack.Marshal(b, m, deterministic)
}
func (dst *Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ack.Merge(dst, src)
}
func (m *Ack) XXX_Size() int {
	return xxx_messageInfo_Ack.Size(m)
}
func (m *Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_Ack proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Transaction)(nil), "protos.Transaction")
	proto.RegisterType((*BlockHeader)(nil), "protos.BlockHeader")
	proto.RegisterType((*Empty)(nil), "protos.Empty")
	proto.RegisterType((*Hello)(nil), "protos.Hello")
	proto.RegisterType((*Ack)(nil), "protos.Ack")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PeeringClient is the client API for Peering service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PeeringClient interface {
	// Could add version exchange during peer connection
	Connect(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Ack, error)
}

type peeringClient struct {
	cc *grpc.ClientConn
}

func NewPeeringClient(cc *grpc.ClientConn) PeeringClient {
	return &peeringClient{cc}
}

func (c *peeringClient) Connect(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/protos.Peering/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeeringServer is the server API for Peering service.
type PeeringServer interface {
	// Could add version exchange during peer connection
	Connect(context.Context, *Hello) (*Ack, error)
}

func RegisterPeeringServer(s *grpc.Server, srv PeeringServer) {
	s.RegisterService(&_Peering_serviceDesc, srv)
}

func _Peering_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Peering/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringServer).Connect(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

var _Peering_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Peering",
	HandlerType: (*PeeringServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _Peering_Connect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coin.proto",
}

// TransactionsClient is the client API for Transactions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionsClient interface {
	ReceiveTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Empty, error)
	SendTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Empty, error)
}

type transactionsClient struct {
	cc *grpc.ClientConn
}

func NewTransactionsClient(cc *grpc.ClientConn) TransactionsClient {
	return &transactionsClient{cc}
}

func (c *transactionsClient) ReceiveTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.Transactions/ReceiveTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsClient) SendTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.Transactions/SendTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionsServer is the server API for Transactions service.
type TransactionsServer interface {
	ReceiveTransaction(context.Context, *Transaction) (*Empty, error)
	SendTransaction(context.Context, *Transaction) (*Empty, error)
}

func RegisterTransactionsServer(s *grpc.Server, srv TransactionsServer) {
	s.RegisterService(&_Transactions_serviceDesc, srv)
}

func _Transactions_ReceiveTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServer).ReceiveTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Transactions/ReceiveTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServer).ReceiveTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transactions_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Transactions/SendTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServer).SendTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transactions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Transactions",
	HandlerType: (*TransactionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveTransaction",
			Handler:    _Transactions_ReceiveTransaction_Handler,
		},
		{
			MethodName: "SendTransaction",
			Handler:    _Transactions_SendTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coin.proto",
}

// StateClient is the client API for State service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StateClient interface {
	// Could be a huge number of blocks and transactions
	// lets use a stream
	GetTransactions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (State_GetTransactionsClient, error)
}

type stateClient struct {
	cc *grpc.ClientConn
}

func NewStateClient(cc *grpc.ClientConn) StateClient {
	return &stateClient{cc}
}

func (c *stateClient) GetTransactions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (State_GetTransactionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_State_serviceDesc.Streams[0], "/protos.State/GetTransactions", opts...)
	if err != nil {
		return nil, err
	}
	x := &stateGetTransactionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type State_GetTransactionsClient interface {
	Recv() (*Transaction, error)
	grpc.ClientStream
}

type stateGetTransactionsClient struct {
	grpc.ClientStream
}

func (x *stateGetTransactionsClient) Recv() (*Transaction, error) {
	m := new(Transaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StateServer is the server API for State service.
type StateServer interface {
	// Could be a huge number of blocks and transactions
	// lets use a stream
	GetTransactions(*Empty, State_GetTransactionsServer) error
}

func RegisterStateServer(s *grpc.Server, srv StateServer) {
	s.RegisterService(&_State_serviceDesc, srv)
}

func _State_GetTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StateServer).GetTransactions(m, &stateGetTransactionsServer{stream})
}

type State_GetTransactionsServer interface {
	Send(*Transaction) error
	grpc.ServerStream
}

type stateGetTransactionsServer struct {
	grpc.ServerStream
}

func (x *stateGetTransactionsServer) Send(m *Transaction) error {
	return x.ServerStream.SendMsg(m)
}

var _State_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.State",
	HandlerType: (*StateServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTransactions",
			Handler:       _State_GetTransactions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "coin.proto",
}

func init() { proto.RegisterFile("coin.proto", fileDescriptor_coin_2f53f782266e87bc) }

var fileDescriptor_coin_2f53f782266e87bc = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4b, 0xfb, 0x30,
	0x14, 0xc7, 0xd7, 0xdf, 0xd6, 0xdf, 0xf0, 0x6d, 0x75, 0x12, 0x3d, 0x94, 0x21, 0x32, 0x8a, 0xe8,
	0xf0, 0x30, 0xa4, 0x1e, 0x44, 0x6f, 0x9b, 0x88, 0x03, 0x0f, 0x8e, 0x6e, 0x82, 0xd7, 0x2c, 0x7b,
	0x9b, 0xa1, 0x6d, 0x52, 0xd2, 0x74, 0xb0, 0xb3, 0x07, 0xff, 0x20, 0xff, 0x41, 0x69, 0xe3, 0xb6,
	0x4e, 0xbd, 0x78, 0x6a, 0xdf, 0x87, 0x6f, 0xbe, 0x7c, 0xc8, 0x0b, 0x00, 0x93, 0x5c, 0xf4, 0x12,
	0x25, 0xb5, 0x24, 0xff, 0x8b, 0x4f, 0xea, 0xbd, 0x5b, 0xd0, 0x98, 0x28, 0x2a, 0x52, 0xca, 0x34,
	0x97, 0x82, 0x1c, 0xc3, 0x1e, 0x17, 0x49, 0xa6, 0x9f, 0x27, 0x2f, 0x4f, 0xae, 0xd5, 0xb1, 0xba,
	0xcd, 0x60, 0x0b, 0xc8, 0x19, 0xec, 0x2b, 0x64, 0xc8, 0x97, 0xa8, 0x46, 0xd9, 0xf4, 0x11, 0x57,
	0x6e, 0xb5, 0x88, 0x7c, 0xa3, 0x79, 0x4b, 0xca, 0x17, 0x82, 0xea, 0x4c, 0xa1, 0x5b, 0x33, 0x2d,
	0x1b, 0x40, 0x8e, 0xc0, 0x5e, 0xd2, 0x28, 0x43, 0xd7, 0xee, 0x58, 0xdd, 0x5a, 0x60, 0x06, 0xef,
	0xc3, 0x82, 0xc6, 0x20, 0x92, 0x2c, 0x1c, 0x22, 0x9d, 0xa1, 0x22, 0xa7, 0xe0, 0x24, 0x0a, 0x97,
	0x06, 0xd1, 0xf4, 0xf5, 0xcb, 0x66, 0x17, 0x92, 0x13, 0x80, 0x18, 0x55, 0x18, 0x61, 0x20, 0xa5,
	0x76, 0xff, 0x15, 0x91, 0x12, 0xc9, 0x4d, 0x34, 0x8f, 0x71, 0xac, 0x69, 0x9c, 0x14, 0xb2, 0x4e,
	0xb0, 0x05, 0xe4, 0x02, 0x0e, 0x66, 0x7c, 0x3e, 0xe7, 0x2c, 0x8b, 0xf4, 0x6a, 0x42, 0xd5, 0x02,
	0x75, 0xa1, 0xeb, 0x04, 0x3f, 0x78, 0x6e, 0x2d, 0xa4, 0x60, 0xc6, 0xda, 0x09, 0xcc, 0xe0, 0xd5,
	0xc1, 0xbe, 0x8f, 0x13, 0xbd, 0xca, 0x7f, 0x86, 0x18, 0x45, 0xd2, 0xb3, 0xa1, 0xda, 0x67, 0xa1,
	0xef, 0x43, 0x7d, 0x84, 0xa8, 0xb8, 0x58, 0x90, 0x73, 0xa8, 0xdf, 0x49, 0x21, 0x90, 0x69, 0xe2,
	0x98, 0xeb, 0x4f, 0x7b, 0x45, 0xb6, 0xdd, 0x58, 0x8f, 0x7d, 0x16, 0x7a, 0x15, 0xff, 0xcd, 0x82,
	0x66, 0x69, 0x19, 0x29, 0xb9, 0x05, 0x12, 0x98, 0x9b, 0x2d, 0xef, 0xe8, 0x70, 0x7d, 0xaa, 0x04,
	0xdb, 0x9b, 0x66, 0xa3, 0x53, 0x21, 0xd7, 0xd0, 0x1a, 0xa3, 0x98, 0xfd, 0xf9, 0xa0, 0x3f, 0x00,
	0x7b, 0xac, 0xa9, 0x46, 0x72, 0x03, 0xad, 0x07, 0xd4, 0x3b, 0x42, 0xbb, 0xe1, 0xf6, 0x6f, 0x85,
	0x5e, 0xe5, 0xd2, 0x9a, 0x9a, 0xe7, 0x75, 0xf5, 0x19, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xd5, 0x45,
	0xca, 0x73, 0x02, 0x00, 0x00,
}
