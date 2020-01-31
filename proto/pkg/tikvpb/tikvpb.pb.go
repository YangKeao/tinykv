// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tikvpb.proto

package tikvpb

import (
	"fmt"
	"io"
	"math"

	proto "github.com/golang/protobuf/proto"

	_ "github.com/gogo/protobuf/gogoproto"

	kvrpcpb "github.com/pingcap-incubator/tinykv/proto/pkg/kvrpcpb"

	raft_serverpb "github.com/pingcap-incubator/tinykv/proto/pkg/raft_serverpb"

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

type BatchRaftMessage struct {
	Msgs                 []*raft_serverpb.RaftMessage `protobuf:"bytes,1,rep,name=msgs" json:"msgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *BatchRaftMessage) Reset()         { *m = BatchRaftMessage{} }
func (m *BatchRaftMessage) String() string { return proto.CompactTextString(m) }
func (*BatchRaftMessage) ProtoMessage()    {}
func (*BatchRaftMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_tikvpb_78b381e194c522ab, []int{0}
}
func (m *BatchRaftMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BatchRaftMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BatchRaftMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BatchRaftMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchRaftMessage.Merge(dst, src)
}
func (m *BatchRaftMessage) XXX_Size() int {
	return m.Size()
}
func (m *BatchRaftMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchRaftMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BatchRaftMessage proto.InternalMessageInfo

func (m *BatchRaftMessage) GetMsgs() []*raft_serverpb.RaftMessage {
	if m != nil {
		return m.Msgs
	}
	return nil
}

func init() {
	proto.RegisterType((*BatchRaftMessage)(nil), "tikvpb.BatchRaftMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Tikv service

type TikvClient interface {
	// KV commands with mvcc/txn supported.
	KvGet(ctx context.Context, in *kvrpcpb.GetRequest, opts ...grpc.CallOption) (*kvrpcpb.GetResponse, error)
	KvScan(ctx context.Context, in *kvrpcpb.ScanRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanResponse, error)
	KvPrewrite(ctx context.Context, in *kvrpcpb.PrewriteRequest, opts ...grpc.CallOption) (*kvrpcpb.PrewriteResponse, error)
	KvCommit(ctx context.Context, in *kvrpcpb.CommitRequest, opts ...grpc.CallOption) (*kvrpcpb.CommitResponse, error)
	KvCheckTxnStatus(ctx context.Context, in *kvrpcpb.CheckTxnStatusRequest, opts ...grpc.CallOption) (*kvrpcpb.CheckTxnStatusResponse, error)
	KvCleanup(ctx context.Context, in *kvrpcpb.CleanupRequest, opts ...grpc.CallOption) (*kvrpcpb.CleanupResponse, error)
	KvBatchGet(ctx context.Context, in *kvrpcpb.BatchGetRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchGetResponse, error)
	KvBatchRollback(ctx context.Context, in *kvrpcpb.BatchRollbackRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchRollbackResponse, error)
	KvScanLock(ctx context.Context, in *kvrpcpb.ScanLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanLockResponse, error)
	KvResolveLock(ctx context.Context, in *kvrpcpb.ResolveLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ResolveLockResponse, error)
	// RawKV commands.
	RawGet(ctx context.Context, in *kvrpcpb.RawGetRequest, opts ...grpc.CallOption) (*kvrpcpb.RawGetResponse, error)
	RawPut(ctx context.Context, in *kvrpcpb.RawPutRequest, opts ...grpc.CallOption) (*kvrpcpb.RawPutResponse, error)
	RawDelete(ctx context.Context, in *kvrpcpb.RawDeleteRequest, opts ...grpc.CallOption) (*kvrpcpb.RawDeleteResponse, error)
	RawScan(ctx context.Context, in *kvrpcpb.RawScanRequest, opts ...grpc.CallOption) (*kvrpcpb.RawScanResponse, error)
	// Raft commands (tinykv <-> tinykv).
	Raft(ctx context.Context, opts ...grpc.CallOption) (Tikv_RaftClient, error)
	Snapshot(ctx context.Context, opts ...grpc.CallOption) (Tikv_SnapshotClient, error)
}

type tikvClient struct {
	cc *grpc.ClientConn
}

func NewTikvClient(cc *grpc.ClientConn) TikvClient {
	return &tikvClient{cc}
}

func (c *tikvClient) KvGet(ctx context.Context, in *kvrpcpb.GetRequest, opts ...grpc.CallOption) (*kvrpcpb.GetResponse, error) {
	out := new(kvrpcpb.GetResponse)
	err := c.cc.Invoke(ctx, "/tikvpb.Tikv/KvGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvScan(ctx context.Context, in *kvrpcpb.ScanRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanResponse, error) {
	out := new(kvrpcpb.ScanResponse)
	err := c.cc.Invoke(ctx, "/tikvpb.Tikv/KvScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvPrewrite(ctx context.Context, in *kvrpcpb.PrewriteRequest, opts ...grpc.CallOption) (*kvrpcpb.PrewriteResponse, error) {
	out := new(kvrpcpb.PrewriteResponse)
	err := c.cc.Invoke(ctx, "/tikvpb.Tikv/KvPrewrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvCommit(ctx context.Context, in *kvrpcpb.CommitRequest, opts ...grpc.CallOption) (*kvrpcpb.CommitResponse, error) {
	out := new(kvrpcpb.CommitResponse)
	err := c.cc.Invoke(ctx, "/tikvpb.Tikv/KvCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvCheckTxnStatus(ctx context.Context, in *kvrpcpb.CheckTxnStatusRequest, opts ...grpc.CallOption) (*kvrpcpb.CheckTxnStatusResponse, error) {
	out := new(kvrpcpb.CheckTxnStatusResponse)
	err := c.cc.Invoke(ctx, "/tikvpb.Tikv/KvCheckTxnStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvCleanup(ctx context.Context, in *kvrpcpb.CleanupRequest, opts ...grpc.CallOption) (*kvrpcpb.CleanupResponse, error) {
	out := new(kvrpcpb.CleanupResponse)
	err := c.cc.Invoke(ctx, "/tikvpb.Tikv/KvCleanup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvBatchGet(ctx context.Context, in *kvrpcpb.BatchGetRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchGetResponse, error) {
	out := new(kvrpcpb.BatchGetResponse)
	err := c.cc.Invoke(ctx, "/tikvpb.Tikv/KvBatchGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvBatchRollback(ctx context.Context, in *kvrpcpb.BatchRollbackRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchRollbackResponse, error) {
	out := new(kvrpcpb.BatchRollbackResponse)
	err := c.cc.Invoke(ctx, "/tikvpb.Tikv/KvBatchRollback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvScanLock(ctx context.Context, in *kvrpcpb.ScanLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanLockResponse, error) {
	out := new(kvrpcpb.ScanLockResponse)
	err := c.cc.Invoke(ctx, "/tikvpb.Tikv/KvScanLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) KvResolveLock(ctx context.Context, in *kvrpcpb.ResolveLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ResolveLockResponse, error) {
	out := new(kvrpcpb.ResolveLockResponse)
	err := c.cc.Invoke(ctx, "/tikvpb.Tikv/KvResolveLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) RawGet(ctx context.Context, in *kvrpcpb.RawGetRequest, opts ...grpc.CallOption) (*kvrpcpb.RawGetResponse, error) {
	out := new(kvrpcpb.RawGetResponse)
	err := c.cc.Invoke(ctx, "/tikvpb.Tikv/RawGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) RawPut(ctx context.Context, in *kvrpcpb.RawPutRequest, opts ...grpc.CallOption) (*kvrpcpb.RawPutResponse, error) {
	out := new(kvrpcpb.RawPutResponse)
	err := c.cc.Invoke(ctx, "/tikvpb.Tikv/RawPut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) RawDelete(ctx context.Context, in *kvrpcpb.RawDeleteRequest, opts ...grpc.CallOption) (*kvrpcpb.RawDeleteResponse, error) {
	out := new(kvrpcpb.RawDeleteResponse)
	err := c.cc.Invoke(ctx, "/tikvpb.Tikv/RawDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) RawScan(ctx context.Context, in *kvrpcpb.RawScanRequest, opts ...grpc.CallOption) (*kvrpcpb.RawScanResponse, error) {
	out := new(kvrpcpb.RawScanResponse)
	err := c.cc.Invoke(ctx, "/tikvpb.Tikv/RawScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tikvClient) Raft(ctx context.Context, opts ...grpc.CallOption) (Tikv_RaftClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tikv_serviceDesc.Streams[0], "/tikvpb.Tikv/Raft", opts...)
	if err != nil {
		return nil, err
	}
	x := &tikvRaftClient{stream}
	return x, nil
}

type Tikv_RaftClient interface {
	Send(*raft_serverpb.RaftMessage) error
	CloseAndRecv() (*raft_serverpb.Done, error)
	grpc.ClientStream
}

type tikvRaftClient struct {
	grpc.ClientStream
}

func (x *tikvRaftClient) Send(m *raft_serverpb.RaftMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tikvRaftClient) CloseAndRecv() (*raft_serverpb.Done, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(raft_serverpb.Done)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tikvClient) Snapshot(ctx context.Context, opts ...grpc.CallOption) (Tikv_SnapshotClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tikv_serviceDesc.Streams[1], "/tikvpb.Tikv/Snapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &tikvSnapshotClient{stream}
	return x, nil
}

type Tikv_SnapshotClient interface {
	Send(*raft_serverpb.SnapshotChunk) error
	CloseAndRecv() (*raft_serverpb.Done, error)
	grpc.ClientStream
}

type tikvSnapshotClient struct {
	grpc.ClientStream
}

func (x *tikvSnapshotClient) Send(m *raft_serverpb.SnapshotChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tikvSnapshotClient) CloseAndRecv() (*raft_serverpb.Done, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(raft_serverpb.Done)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Tikv service

type TikvServer interface {
	// KV commands with mvcc/txn supported.
	KvGet(context.Context, *kvrpcpb.GetRequest) (*kvrpcpb.GetResponse, error)
	KvScan(context.Context, *kvrpcpb.ScanRequest) (*kvrpcpb.ScanResponse, error)
	KvPrewrite(context.Context, *kvrpcpb.PrewriteRequest) (*kvrpcpb.PrewriteResponse, error)
	KvCommit(context.Context, *kvrpcpb.CommitRequest) (*kvrpcpb.CommitResponse, error)
	KvCheckTxnStatus(context.Context, *kvrpcpb.CheckTxnStatusRequest) (*kvrpcpb.CheckTxnStatusResponse, error)
	KvCleanup(context.Context, *kvrpcpb.CleanupRequest) (*kvrpcpb.CleanupResponse, error)
	KvBatchGet(context.Context, *kvrpcpb.BatchGetRequest) (*kvrpcpb.BatchGetResponse, error)
	KvBatchRollback(context.Context, *kvrpcpb.BatchRollbackRequest) (*kvrpcpb.BatchRollbackResponse, error)
	KvScanLock(context.Context, *kvrpcpb.ScanLockRequest) (*kvrpcpb.ScanLockResponse, error)
	KvResolveLock(context.Context, *kvrpcpb.ResolveLockRequest) (*kvrpcpb.ResolveLockResponse, error)
	// RawKV commands.
	RawGet(context.Context, *kvrpcpb.RawGetRequest) (*kvrpcpb.RawGetResponse, error)
	RawPut(context.Context, *kvrpcpb.RawPutRequest) (*kvrpcpb.RawPutResponse, error)
	RawDelete(context.Context, *kvrpcpb.RawDeleteRequest) (*kvrpcpb.RawDeleteResponse, error)
	RawScan(context.Context, *kvrpcpb.RawScanRequest) (*kvrpcpb.RawScanResponse, error)
	// Raft commands (tinykv <-> tinykv).
	Raft(Tikv_RaftServer) error
	Snapshot(Tikv_SnapshotServer) error
}

func RegisterTikvServer(s *grpc.Server, srv TikvServer) {
	s.RegisterService(&_Tikv_serviceDesc, srv)
}

func _Tikv_KvGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvGet(ctx, req.(*kvrpcpb.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvScan(ctx, req.(*kvrpcpb.ScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvPrewrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.PrewriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvPrewrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvPrewrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvPrewrite(ctx, req.(*kvrpcpb.PrewriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvCommit(ctx, req.(*kvrpcpb.CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvCheckTxnStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.CheckTxnStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvCheckTxnStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvCheckTxnStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvCheckTxnStatus(ctx, req.(*kvrpcpb.CheckTxnStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvCleanup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.CleanupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvCleanup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvCleanup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvCleanup(ctx, req.(*kvrpcpb.CleanupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvBatchGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.BatchGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvBatchGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvBatchGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvBatchGet(ctx, req.(*kvrpcpb.BatchGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvBatchRollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.BatchRollbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvBatchRollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvBatchRollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvBatchRollback(ctx, req.(*kvrpcpb.BatchRollbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvScanLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ScanLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvScanLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvScanLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvScanLock(ctx, req.(*kvrpcpb.ScanLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_KvResolveLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ResolveLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).KvResolveLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/KvResolveLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).KvResolveLock(ctx, req.(*kvrpcpb.ResolveLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_RawGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).RawGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/RawGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).RawGet(ctx, req.(*kvrpcpb.RawGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_RawPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).RawPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/RawPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).RawPut(ctx, req.(*kvrpcpb.RawPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_RawDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).RawDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/RawDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).RawDelete(ctx, req.(*kvrpcpb.RawDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_RawScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikvServer).RawScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.Tikv/RawScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikvServer).RawScan(ctx, req.(*kvrpcpb.RawScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tikv_Raft_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TikvServer).Raft(&tikvRaftServer{stream})
}

type Tikv_RaftServer interface {
	SendAndClose(*raft_serverpb.Done) error
	Recv() (*raft_serverpb.RaftMessage, error)
	grpc.ServerStream
}

type tikvRaftServer struct {
	grpc.ServerStream
}

func (x *tikvRaftServer) SendAndClose(m *raft_serverpb.Done) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tikvRaftServer) Recv() (*raft_serverpb.RaftMessage, error) {
	m := new(raft_serverpb.RaftMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tikv_Snapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TikvServer).Snapshot(&tikvSnapshotServer{stream})
}

type Tikv_SnapshotServer interface {
	SendAndClose(*raft_serverpb.Done) error
	Recv() (*raft_serverpb.SnapshotChunk, error)
	grpc.ServerStream
}

type tikvSnapshotServer struct {
	grpc.ServerStream
}

func (x *tikvSnapshotServer) SendAndClose(m *raft_serverpb.Done) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tikvSnapshotServer) Recv() (*raft_serverpb.SnapshotChunk, error) {
	m := new(raft_serverpb.SnapshotChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Tikv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tikvpb.Tikv",
	HandlerType: (*TikvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KvGet",
			Handler:    _Tikv_KvGet_Handler,
		},
		{
			MethodName: "KvScan",
			Handler:    _Tikv_KvScan_Handler,
		},
		{
			MethodName: "KvPrewrite",
			Handler:    _Tikv_KvPrewrite_Handler,
		},
		{
			MethodName: "KvCommit",
			Handler:    _Tikv_KvCommit_Handler,
		},
		{
			MethodName: "KvCheckTxnStatus",
			Handler:    _Tikv_KvCheckTxnStatus_Handler,
		},
		{
			MethodName: "KvCleanup",
			Handler:    _Tikv_KvCleanup_Handler,
		},
		{
			MethodName: "KvBatchGet",
			Handler:    _Tikv_KvBatchGet_Handler,
		},
		{
			MethodName: "KvBatchRollback",
			Handler:    _Tikv_KvBatchRollback_Handler,
		},
		{
			MethodName: "KvScanLock",
			Handler:    _Tikv_KvScanLock_Handler,
		},
		{
			MethodName: "KvResolveLock",
			Handler:    _Tikv_KvResolveLock_Handler,
		},
		{
			MethodName: "RawGet",
			Handler:    _Tikv_RawGet_Handler,
		},
		{
			MethodName: "RawPut",
			Handler:    _Tikv_RawPut_Handler,
		},
		{
			MethodName: "RawDelete",
			Handler:    _Tikv_RawDelete_Handler,
		},
		{
			MethodName: "RawScan",
			Handler:    _Tikv_RawScan_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Raft",
			Handler:       _Tikv_Raft_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Snapshot",
			Handler:       _Tikv_Snapshot_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "tikvpb.proto",
}

func (m *BatchRaftMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BatchRaftMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for _, msg := range m.Msgs {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTikvpb(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTikvpb(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BatchRaftMessage) Size() (n int) {
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovTikvpb(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTikvpb(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTikvpb(x uint64) (n int) {
	return sovTikvpb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BatchRaftMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTikvpb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BatchRaftMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BatchRaftMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTikvpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTikvpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msgs = append(m.Msgs, &raft_serverpb.RaftMessage{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTikvpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTikvpb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTikvpb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTikvpb
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTikvpb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTikvpb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTikvpb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTikvpb
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTikvpb(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTikvpb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTikvpb   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("tikvpb.proto", fileDescriptor_tikvpb_78b381e194c522ab) }

var fileDescriptor_tikvpb_78b381e194c522ab = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x6b, 0x11, 0x42, 0x3a, 0x50, 0x11, 0x6d, 0x0a, 0xa4, 0xa6, 0x18, 0x94, 0x53, 0x4f,
	0x46, 0x2a, 0x48, 0x1c, 0x10, 0x08, 0xe2, 0x48, 0x3d, 0xb8, 0x48, 0x91, 0x53, 0xce, 0x68, 0x63,
	0x4d, 0x9d, 0xc8, 0x8e, 0xd7, 0x78, 0xd7, 0x1b, 0x1e, 0x85, 0xa7, 0xe1, 0xcc, 0x91, 0x47, 0x40,
	0xe1, 0x45, 0x90, 0x6d, 0x76, 0xfd, 0x2f, 0xe1, 0x14, 0xfb, 0xfb, 0xbe, 0xf9, 0x79, 0xb2, 0xbb,
	0xb3, 0xf0, 0x40, 0xac, 0x43, 0x99, 0x2c, 0xed, 0x24, 0x65, 0x82, 0x91, 0x7e, 0xf9, 0x66, 0x9e,
	0x84, 0x32, 0x4d, 0x7c, 0x25, 0x9b, 0xa3, 0x94, 0xde, 0x8a, 0x2f, 0x1c, 0x53, 0x89, 0xa9, 0x16,
	0x4f, 0x03, 0x16, 0xb0, 0xe2, 0xf1, 0x65, 0xfe, 0x54, 0xaa, 0x93, 0x29, 0x0c, 0xa7, 0x54, 0xf8,
	0x2b, 0x8f, 0xde, 0x8a, 0x4f, 0xc8, 0x39, 0x0d, 0x90, 0xd8, 0xd0, 0xdb, 0xf0, 0x80, 0x8f, 0x8d,
	0x17, 0x77, 0x2e, 0xee, 0x5f, 0x9a, 0x76, 0x93, 0x56, 0x4b, 0x7a, 0x45, 0xee, 0xf2, 0xc7, 0x00,
	0x7a, 0x37, 0xeb, 0x50, 0x92, 0xd7, 0x70, 0xd7, 0x95, 0x57, 0x28, 0xc8, 0xc8, 0x56, 0x0d, 0x5d,
	0xa1, 0xf0, 0xf0, 0x6b, 0x86, 0x5c, 0x98, 0xa7, 0x4d, 0x91, 0x27, 0x2c, 0xe6, 0x38, 0x39, 0x22,
	0x6f, 0xa0, 0xef, 0xca, 0x85, 0x4f, 0x63, 0x52, 0x25, 0xf2, 0x57, 0x55, 0xf7, 0xa8, 0xa5, 0xea,
	0x42, 0x07, 0xc0, 0x95, 0xf3, 0x14, 0xb7, 0xe9, 0x5a, 0x20, 0x19, 0xeb, 0x98, 0x92, 0x14, 0xe0,
	0x6c, 0x8f, 0xa3, 0x21, 0xef, 0x60, 0xe0, 0x4a, 0x87, 0x6d, 0x36, 0x6b, 0x41, 0x1e, 0xeb, 0x60,
	0x29, 0x28, 0xc0, 0x93, 0x8e, 0xae, 0xcb, 0x3f, 0xc3, 0xd0, 0x95, 0xce, 0x0a, 0xfd, 0xf0, 0xe6,
	0x5b, 0xbc, 0x10, 0x54, 0x64, 0x9c, 0x58, 0x55, 0xbc, 0x61, 0x28, 0xdc, 0xf3, 0x83, 0xbe, 0xc6,
	0x7e, 0x80, 0x63, 0x57, 0x3a, 0x11, 0xd2, 0x38, 0x4b, 0x48, 0xed, 0xf3, 0xa5, 0xa2, 0x40, 0xe3,
	0xae, 0xd1, 0x5c, 0x9c, 0x62, 0x6b, 0xf3, 0x0d, 0xa9, 0x92, 0x4a, 0xea, 0x2e, 0x4e, 0xe5, 0x68,
	0x88, 0x07, 0x0f, 0xff, 0x41, 0x3c, 0x16, 0x45, 0x4b, 0xea, 0x87, 0xe4, 0x59, 0x33, 0xaf, 0x74,
	0x85, 0xb3, 0x0e, 0xd9, 0xcd, 0xc6, 0xf2, 0x9d, 0xbc, 0x66, 0x7e, 0x58, 0x6b, 0x4c, 0x49, 0xdd,
	0xc6, 0x2a, 0x47, 0x43, 0xae, 0xe1, 0xc4, 0x95, 0x1e, 0x72, 0x16, 0x49, 0x2c, 0x38, 0x4f, 0x75,
	0xba, 0xa6, 0x2a, 0xd4, 0xf9, 0x7e, 0x53, 0xd3, 0xde, 0x42, 0xdf, 0xa3, 0xdb, 0x7c, 0x9d, 0xaa,
	0x13, 0x50, 0x0a, 0xdd, 0x13, 0xa0, 0xf4, 0x56, 0xf1, 0x3c, 0x6b, 0x15, 0xcf, 0xb3, 0xfd, 0xc5,
	0x85, 0xae, 0x8b, 0x67, 0x70, 0xec, 0xd1, 0xed, 0x0c, 0x23, 0x14, 0x48, 0xce, 0xea, 0xb9, 0x52,
	0x53, 0x08, 0x73, 0x9f, 0xa5, 0x29, 0xef, 0xe1, 0x9e, 0x47, 0xb7, 0xc5, 0x08, 0x35, 0xbe, 0x55,
	0x9f, 0xa2, 0x71, 0xd7, 0xa8, 0xfd, 0x85, 0x5e, 0x3e, 0xd5, 0xe4, 0x3f, 0xa3, 0x6e, 0x8e, 0x5a,
	0xde, 0x8c, 0xc5, 0x38, 0x39, 0xba, 0x30, 0xc8, 0x47, 0x18, 0x2c, 0x62, 0x9a, 0xf0, 0x15, 0x13,
	0xe4, 0xbc, 0x15, 0x52, 0x86, 0xb3, 0xca, 0xe2, 0xf0, 0x20, 0x62, 0x3a, 0xf9, 0xb9, 0xb3, 0x8c,
	0x5f, 0x3b, 0xcb, 0xf8, 0xbd, 0xb3, 0x8c, 0xef, 0x7f, 0xac, 0x23, 0x18, 0xb2, 0x34, 0xb0, 0xf3,
	0xcb, 0xcd, 0x0e, 0x65, 0x71, 0x51, 0x2d, 0xfb, 0xc5, 0xcf, 0xab, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x8c, 0x38, 0xab, 0xfa, 0x01, 0x05, 0x00, 0x00,
}