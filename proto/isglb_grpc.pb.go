// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/isglb.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ISGLBClient is the client API for ISGLB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ISGLBClient interface {
	// When forward path changed, upload the SFUStatus from client to ISGLB
	// When forward path should change, send expected SFUStatus from ISGLB to client
	SyncSFUStatus(ctx context.Context, opts ...grpc.CallOption) (ISGLB_SyncSFUStatusClient, error)
	// Report the communication quality or computation quality of the edge
	QualityReport(ctx context.Context, opts ...grpc.CallOption) (ISGLB_QualityReportClient, error)
}

type iSGLBClient struct {
	cc grpc.ClientConnInterface
}

func NewISGLBClient(cc grpc.ClientConnInterface) ISGLBClient {
	return &iSGLBClient{cc}
}

func (c *iSGLBClient) SyncSFUStatus(ctx context.Context, opts ...grpc.CallOption) (ISGLB_SyncSFUStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &ISGLB_ServiceDesc.Streams[0], "/islb.ISGLB/SyncSFUStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &iSGLBSyncSFUStatusClient{stream}
	return x, nil
}

type ISGLB_SyncSFUStatusClient interface {
	Send(*SFUStatus) error
	Recv() (*SFUStatus, error)
	grpc.ClientStream
}

type iSGLBSyncSFUStatusClient struct {
	grpc.ClientStream
}

func (x *iSGLBSyncSFUStatusClient) Send(m *SFUStatus) error {
	return x.ClientStream.SendMsg(m)
}

func (x *iSGLBSyncSFUStatusClient) Recv() (*SFUStatus, error) {
	m := new(SFUStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *iSGLBClient) QualityReport(ctx context.Context, opts ...grpc.CallOption) (ISGLB_QualityReportClient, error) {
	stream, err := c.cc.NewStream(ctx, &ISGLB_ServiceDesc.Streams[1], "/islb.ISGLB/QualityReport", opts...)
	if err != nil {
		return nil, err
	}
	x := &iSGLBQualityReportClient{stream}
	return x, nil
}

type ISGLB_QualityReportClient interface {
	Send(*Report) error
	Recv() (*SFUStatus, error)
	grpc.ClientStream
}

type iSGLBQualityReportClient struct {
	grpc.ClientStream
}

func (x *iSGLBQualityReportClient) Send(m *Report) error {
	return x.ClientStream.SendMsg(m)
}

func (x *iSGLBQualityReportClient) Recv() (*SFUStatus, error) {
	m := new(SFUStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ISGLBServer is the server API for ISGLB service.
// All implementations must embed UnimplementedISGLBServer
// for forward compatibility
type ISGLBServer interface {
	// When forward path changed, upload the SFUStatus from client to ISGLB
	// When forward path should change, send expected SFUStatus from ISGLB to client
	SyncSFUStatus(ISGLB_SyncSFUStatusServer) error
	// Report the communication quality or computation quality of the edge
	QualityReport(ISGLB_QualityReportServer) error
	mustEmbedUnimplementedISGLBServer()
}

// UnimplementedISGLBServer must be embedded to have forward compatible implementations.
type UnimplementedISGLBServer struct {
}

func (UnimplementedISGLBServer) SyncSFUStatus(ISGLB_SyncSFUStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncSFUStatus not implemented")
}
func (UnimplementedISGLBServer) QualityReport(ISGLB_QualityReportServer) error {
	return status.Errorf(codes.Unimplemented, "method QualityReport not implemented")
}
func (UnimplementedISGLBServer) mustEmbedUnimplementedISGLBServer() {}

// UnsafeISGLBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ISGLBServer will
// result in compilation errors.
type UnsafeISGLBServer interface {
	mustEmbedUnimplementedISGLBServer()
}

func RegisterISGLBServer(s grpc.ServiceRegistrar, srv ISGLBServer) {
	s.RegisterService(&ISGLB_ServiceDesc, srv)
}

func _ISGLB_SyncSFUStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ISGLBServer).SyncSFUStatus(&iSGLBSyncSFUStatusServer{stream})
}

type ISGLB_SyncSFUStatusServer interface {
	Send(*SFUStatus) error
	Recv() (*SFUStatus, error)
	grpc.ServerStream
}

type iSGLBSyncSFUStatusServer struct {
	grpc.ServerStream
}

func (x *iSGLBSyncSFUStatusServer) Send(m *SFUStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *iSGLBSyncSFUStatusServer) Recv() (*SFUStatus, error) {
	m := new(SFUStatus)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ISGLB_QualityReport_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ISGLBServer).QualityReport(&iSGLBQualityReportServer{stream})
}

type ISGLB_QualityReportServer interface {
	Send(*SFUStatus) error
	Recv() (*Report, error)
	grpc.ServerStream
}

type iSGLBQualityReportServer struct {
	grpc.ServerStream
}

func (x *iSGLBQualityReportServer) Send(m *SFUStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *iSGLBQualityReportServer) Recv() (*Report, error) {
	m := new(Report)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ISGLB_ServiceDesc is the grpc.ServiceDesc for ISGLB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ISGLB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "islb.ISGLB",
	HandlerType: (*ISGLBServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SyncSFUStatus",
			Handler:       _ISGLB_SyncSFUStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "QualityReport",
			Handler:       _ISGLB_QualityReport_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/isglb.proto",
}
