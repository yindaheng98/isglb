// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/isglb.proto

package proto

import (
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	ion "github.com/pion/ion/proto/ion"
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

type Subscription_Layer int32

const (
	Subscription_Q Subscription_Layer = 0
	Subscription_H Subscription_Layer = 1
	Subscription_F Subscription_Layer = 2
)

// Enum value maps for Subscription_Layer.
var (
	Subscription_Layer_name = map[int32]string{
		0: "Q",
		1: "H",
		2: "F",
	}
	Subscription_Layer_value = map[string]int32{
		"Q": 0,
		"H": 1,
		"F": 2,
	}
)

func (x Subscription_Layer) Enum() *Subscription_Layer {
	p := new(Subscription_Layer)
	*p = x
	return p
}

func (x Subscription_Layer) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Subscription_Layer) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_isglb_proto_enumTypes[0].Descriptor()
}

func (Subscription_Layer) Type() protoreflect.EnumType {
	return &file_proto_isglb_proto_enumTypes[0]
}

func (x Subscription_Layer) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Subscription_Layer.Descriptor instead.
func (Subscription_Layer) EnumDescriptor() ([]byte, []int) {
	return file_proto_isglb_proto_rawDescGZIP(), []int{0, 0}
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackId string             `protobuf:"bytes,1,opt,name=TrackId,proto3" json:"TrackId,omitempty"`
	Mute    bool               `protobuf:"varint,2,opt,name=Mute,proto3" json:"Mute,omitempty"`
	Layer   Subscription_Layer `protobuf:"varint,3,opt,name=layer,proto3,enum=islb.Subscription_Layer" json:"layer,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_isglb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_proto_isglb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_proto_isglb_proto_rawDescGZIP(), []int{0}
}

func (x *Subscription) GetTrackId() string {
	if x != nil {
		return x.TrackId
	}
	return ""
}

func (x *Subscription) GetMute() bool {
	if x != nil {
		return x.Mute
	}
	return false
}

func (x *Subscription) GetLayer() Subscription_Layer {
	if x != nil {
		return x.Layer
	}
	return Subscription_Q
}

// The track in session `remoteSessionId` from `src` node was/should forwarded to local session `localSessionId`
type ForwardTrack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  src + remoteSessionId is unique, use it as key
	Src             *ion.Node       `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
	RemoteSessionId string          `protobuf:"bytes,2,opt,name=remoteSessionId,proto3" json:"remoteSessionId,omitempty"` // The video/audio has session id in different nodes
	LocalSessionId  string          `protobuf:"bytes,3,opt,name=localSessionId,proto3" json:"localSessionId,omitempty"`   // The video/audio has session id in different nodes
	Tracks          []*Subscription `protobuf:"bytes,4,rep,name=tracks,proto3" json:"tracks,omitempty"`
}

func (x *ForwardTrack) Reset() {
	*x = ForwardTrack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_isglb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForwardTrack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardTrack) ProtoMessage() {}

func (x *ForwardTrack) ProtoReflect() protoreflect.Message {
	mi := &file_proto_isglb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardTrack.ProtoReflect.Descriptor instead.
func (*ForwardTrack) Descriptor() ([]byte, []int) {
	return file_proto_isglb_proto_rawDescGZIP(), []int{1}
}

func (x *ForwardTrack) GetSrc() *ion.Node {
	if x != nil {
		return x.Src
	}
	return nil
}

func (x *ForwardTrack) GetRemoteSessionId() string {
	if x != nil {
		return x.RemoteSessionId
	}
	return ""
}

func (x *ForwardTrack) GetLocalSessionId() string {
	if x != nil {
		return x.LocalSessionId
	}
	return ""
}

func (x *ForwardTrack) GetTracks() []*Subscription {
	if x != nil {
		return x.Tracks
	}
	return nil
}

// The track from srcSessionId was/should be proceeded and output track's session id is `dstSessionId`
type ProceedTrack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcSessionIdList []string `protobuf:"bytes,1,rep,name=srcSessionIdList,proto3" json:"srcSessionIdList,omitempty"`
	DstSessionId     string   `protobuf:"bytes,2,opt,name=dstSessionId,proto3" json:"dstSessionId,omitempty"`
	Procedure        string   `protobuf:"bytes,3,opt,name=procedure,proto3" json:"procedure,omitempty"` //src track should be proceeded by what kind of procedure
}

func (x *ProceedTrack) Reset() {
	*x = ProceedTrack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_isglb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProceedTrack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProceedTrack) ProtoMessage() {}

func (x *ProceedTrack) ProtoReflect() protoreflect.Message {
	mi := &file_proto_isglb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProceedTrack.ProtoReflect.Descriptor instead.
func (*ProceedTrack) Descriptor() ([]byte, []int) {
	return file_proto_isglb_proto_rawDescGZIP(), []int{2}
}

func (x *ProceedTrack) GetSrcSessionIdList() []string {
	if x != nil {
		return x.SrcSessionIdList
	}
	return nil
}

func (x *ProceedTrack) GetDstSessionId() string {
	if x != nil {
		return x.DstSessionId
	}
	return ""
}

func (x *ProceedTrack) GetProcedure() string {
	if x != nil {
		return x.Procedure
	}
	return ""
}

// Which session do the client need. May be a lot of track in one session
type ClientNeededSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session string `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	User    string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ClientNeededSession) Reset() {
	*x = ClientNeededSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_isglb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientNeededSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientNeededSession) ProtoMessage() {}

func (x *ClientNeededSession) ProtoReflect() protoreflect.Message {
	mi := &file_proto_isglb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientNeededSession.ProtoReflect.Descriptor instead.
func (*ClientNeededSession) Descriptor() ([]byte, []int) {
	return file_proto_isglb_proto_rawDescGZIP(), []int{3}
}

func (x *ClientNeededSession) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

func (x *ClientNeededSession) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

// TODO: move SFUStatus to SFU module
type SFUStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SFU           *ion.Node              `protobuf:"bytes,1,opt,name=SFU,proto3" json:"SFU,omitempty"`
	ForwardTracks []*ForwardTrack        `protobuf:"bytes,2,rep,name=forwardTracks,proto3" json:"forwardTracks,omitempty"`
	ProceedTracks []*ProceedTrack        `protobuf:"bytes,3,rep,name=proceedTracks,proto3" json:"proceedTracks,omitempty"`
	Clients       []*ClientNeededSession `protobuf:"bytes,4,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *SFUStatus) Reset() {
	*x = SFUStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_isglb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SFUStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SFUStatus) ProtoMessage() {}

func (x *SFUStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_isglb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SFUStatus.ProtoReflect.Descriptor instead.
func (*SFUStatus) Descriptor() ([]byte, []int) {
	return file_proto_isglb_proto_rawDescGZIP(), []int{4}
}

func (x *SFUStatus) GetSFU() *ion.Node {
	if x != nil {
		return x.SFU
	}
	return nil
}

func (x *SFUStatus) GetForwardTracks() []*ForwardTrack {
	if x != nil {
		return x.ForwardTracks
	}
	return nil
}

func (x *SFUStatus) GetProceedTracks() []*ProceedTrack {
	if x != nil {
		return x.ProceedTracks
	}
	return nil
}

func (x *SFUStatus) GetClients() []*ClientNeededSession {
	if x != nil {
		return x.Clients
	}
	return nil
}

type TransmissionReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src    *ion.Node `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
	Dst    *ion.Node `protobuf:"bytes,2,opt,name=dst,proto3" json:"dst,omitempty"`
	Report *any.Any  `protobuf:"bytes,3,opt,name=report,proto3" json:"report,omitempty"` //TODO: Add a more specific Report type
}

func (x *TransmissionReport) Reset() {
	*x = TransmissionReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_isglb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransmissionReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransmissionReport) ProtoMessage() {}

func (x *TransmissionReport) ProtoReflect() protoreflect.Message {
	mi := &file_proto_isglb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransmissionReport.ProtoReflect.Descriptor instead.
func (*TransmissionReport) Descriptor() ([]byte, []int) {
	return file_proto_isglb_proto_rawDescGZIP(), []int{5}
}

func (x *TransmissionReport) GetSrc() *ion.Node {
	if x != nil {
		return x.Src
	}
	return nil
}

func (x *TransmissionReport) GetDst() *ion.Node {
	if x != nil {
		return x.Dst
	}
	return nil
}

func (x *TransmissionReport) GetReport() *any.Any {
	if x != nil {
		return x.Report
	}
	return nil
}

type ComputationReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node   *ion.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Report *any.Any  `protobuf:"bytes,2,opt,name=report,proto3" json:"report,omitempty"` //TODO: Add a more specific Report type
}

func (x *ComputationReport) Reset() {
	*x = ComputationReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_isglb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputationReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputationReport) ProtoMessage() {}

func (x *ComputationReport) ProtoReflect() protoreflect.Message {
	mi := &file_proto_isglb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputationReport.ProtoReflect.Descriptor instead.
func (*ComputationReport) Descriptor() ([]byte, []int) {
	return file_proto_isglb_proto_rawDescGZIP(), []int{6}
}

func (x *ComputationReport) GetNode() *ion.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *ComputationReport) GetReport() *any.Any {
	if x != nil {
		return x.Report
	}
	return nil
}

type QualityReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp    *timestamp.Timestamp  `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Transmission []*TransmissionReport `protobuf:"bytes,2,rep,name=transmission,proto3" json:"transmission,omitempty"`
	Computation  []*ComputationReport  `protobuf:"bytes,3,rep,name=computation,proto3" json:"computation,omitempty"`
}

func (x *QualityReport) Reset() {
	*x = QualityReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_isglb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QualityReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QualityReport) ProtoMessage() {}

func (x *QualityReport) ProtoReflect() protoreflect.Message {
	mi := &file_proto_isglb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QualityReport.ProtoReflect.Descriptor instead.
func (*QualityReport) Descriptor() ([]byte, []int) {
	return file_proto_isglb_proto_rawDescGZIP(), []int{7}
}

func (x *QualityReport) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *QualityReport) GetTransmission() []*TransmissionReport {
	if x != nil {
		return x.Transmission
	}
	return nil
}

func (x *QualityReport) GetComputation() []*ComputationReport {
	if x != nil {
		return x.Computation
	}
	return nil
}

type SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*SyncRequest_Status
	//	*SyncRequest_Report
	Request isSyncRequest_Request `protobuf_oneof:"request"`
}

func (x *SyncRequest) Reset() {
	*x = SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_isglb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRequest) ProtoMessage() {}

func (x *SyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_isglb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRequest.ProtoReflect.Descriptor instead.
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return file_proto_isglb_proto_rawDescGZIP(), []int{8}
}

func (m *SyncRequest) GetRequest() isSyncRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *SyncRequest) GetStatus() *SFUStatus {
	if x, ok := x.GetRequest().(*SyncRequest_Status); ok {
		return x.Status
	}
	return nil
}

func (x *SyncRequest) GetReport() *QualityReport {
	if x, ok := x.GetRequest().(*SyncRequest_Report); ok {
		return x.Report
	}
	return nil
}

type isSyncRequest_Request interface {
	isSyncRequest_Request()
}

type SyncRequest_Status struct {
	Status *SFUStatus `protobuf:"bytes,1,opt,name=status,proto3,oneof"`
}

type SyncRequest_Report struct {
	Report *QualityReport `protobuf:"bytes,2,opt,name=report,proto3,oneof"`
}

func (*SyncRequest_Status) isSyncRequest_Request() {}

func (*SyncRequest_Report) isSyncRequest_Request() {}

var File_proto_isglb_proto protoreflect.FileDescriptor

var file_proto_isglb_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x73, 0x67, 0x6c, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x69, 0x73, 0x6c, 0x62, 0x1a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a,
	0x01, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x75, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x4d, 0x75, 0x74, 0x65, 0x12, 0x2e, 0x0a,
	0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x69,
	0x73, 0x6c, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x1c, 0x0a,
	0x05, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x05, 0x0a, 0x01, 0x51, 0x10, 0x00, 0x12, 0x05, 0x0a,
	0x01, 0x48, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x46, 0x10, 0x02, 0x22, 0xa9, 0x01, 0x0a, 0x0c,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x1b, 0x0a, 0x03,
	0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x69, 0x6f, 0x6e, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x73,
	0x6c, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x7c, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x72, 0x63, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x10, 0x73, 0x72, 0x63, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x73, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x73, 0x74, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x64, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x64, 0x75, 0x72, 0x65, 0x22, 0x43, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e,
	0x65, 0x65, 0x64, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xd1, 0x01, 0x0a, 0x09, 0x53,
	0x46, 0x55, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x03, 0x53, 0x46, 0x55, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x03, 0x53, 0x46, 0x55, 0x12, 0x38, 0x0a, 0x0d, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69,
	0x73, 0x6c, 0x62, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x52, 0x0d, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x12,
	0x38, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x73, 0x6c, 0x62, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x73, 0x6c,
	0x62, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x7c,
	0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x03, 0x73, 0x72,
	0x63, 0x12, 0x1b, 0x0a, 0x03, 0x64, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x03, 0x64, 0x73, 0x74, 0x12, 0x2c,
	0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x60, 0x0a, 0x11,
	0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65,
	0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xc2,
	0x01, 0x0a, 0x0d, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3c, 0x0a, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x69, 0x73, 0x6c, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x69, 0x73, 0x6c, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x72, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x73, 0x6c, 0x62, 0x2e, 0x53, 0x46, 0x55, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a,
	0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x69, 0x73, 0x6c, 0x62, 0x2e, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x09, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x3c, 0x0a, 0x05, 0x49, 0x53, 0x47, 0x4c, 0x42,
	0x12, 0x33, 0x0a, 0x07, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x46, 0x55, 0x12, 0x11, 0x2e, 0x69, 0x73,
	0x6c, 0x62, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x69, 0x73, 0x6c, 0x62, 0x2e, 0x53, 0x46, 0x55, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x69, 0x6e, 0x64, 0x61, 0x68, 0x65, 0x6e, 0x67, 0x39, 0x38, 0x2f,
	0x64, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_isglb_proto_rawDescOnce sync.Once
	file_proto_isglb_proto_rawDescData = file_proto_isglb_proto_rawDesc
)

func file_proto_isglb_proto_rawDescGZIP() []byte {
	file_proto_isglb_proto_rawDescOnce.Do(func() {
		file_proto_isglb_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_isglb_proto_rawDescData)
	})
	return file_proto_isglb_proto_rawDescData
}

var file_proto_isglb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_isglb_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_isglb_proto_goTypes = []interface{}{
	(Subscription_Layer)(0),     // 0: islb.Subscription.Layer
	(*Subscription)(nil),        // 1: islb.Subscription
	(*ForwardTrack)(nil),        // 2: islb.ForwardTrack
	(*ProceedTrack)(nil),        // 3: islb.ProceedTrack
	(*ClientNeededSession)(nil), // 4: islb.ClientNeededSession
	(*SFUStatus)(nil),           // 5: islb.SFUStatus
	(*TransmissionReport)(nil),  // 6: islb.TransmissionReport
	(*ComputationReport)(nil),   // 7: islb.ComputationReport
	(*QualityReport)(nil),       // 8: islb.QualityReport
	(*SyncRequest)(nil),         // 9: islb.SyncRequest
	(*ion.Node)(nil),            // 10: ion.Node
	(*any.Any)(nil),             // 11: google.protobuf.Any
	(*timestamp.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_proto_isglb_proto_depIdxs = []int32{
	0,  // 0: islb.Subscription.layer:type_name -> islb.Subscription.Layer
	10, // 1: islb.ForwardTrack.src:type_name -> ion.Node
	1,  // 2: islb.ForwardTrack.tracks:type_name -> islb.Subscription
	10, // 3: islb.SFUStatus.SFU:type_name -> ion.Node
	2,  // 4: islb.SFUStatus.forwardTracks:type_name -> islb.ForwardTrack
	3,  // 5: islb.SFUStatus.proceedTracks:type_name -> islb.ProceedTrack
	4,  // 6: islb.SFUStatus.clients:type_name -> islb.ClientNeededSession
	10, // 7: islb.TransmissionReport.src:type_name -> ion.Node
	10, // 8: islb.TransmissionReport.dst:type_name -> ion.Node
	11, // 9: islb.TransmissionReport.report:type_name -> google.protobuf.Any
	10, // 10: islb.ComputationReport.node:type_name -> ion.Node
	11, // 11: islb.ComputationReport.report:type_name -> google.protobuf.Any
	12, // 12: islb.QualityReport.timestamp:type_name -> google.protobuf.Timestamp
	6,  // 13: islb.QualityReport.transmission:type_name -> islb.TransmissionReport
	7,  // 14: islb.QualityReport.computation:type_name -> islb.ComputationReport
	5,  // 15: islb.SyncRequest.status:type_name -> islb.SFUStatus
	8,  // 16: islb.SyncRequest.report:type_name -> islb.QualityReport
	9,  // 17: islb.ISGLB.SyncSFU:input_type -> islb.SyncRequest
	5,  // 18: islb.ISGLB.SyncSFU:output_type -> islb.SFUStatus
	18, // [18:19] is the sub-list for method output_type
	17, // [17:18] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_proto_isglb_proto_init() }
func file_proto_isglb_proto_init() {
	if File_proto_isglb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_isglb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
		file_proto_isglb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForwardTrack); i {
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
		file_proto_isglb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProceedTrack); i {
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
		file_proto_isglb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientNeededSession); i {
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
		file_proto_isglb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SFUStatus); i {
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
		file_proto_isglb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransmissionReport); i {
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
		file_proto_isglb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputationReport); i {
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
		file_proto_isglb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QualityReport); i {
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
		file_proto_isglb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRequest); i {
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
	file_proto_isglb_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*SyncRequest_Status)(nil),
		(*SyncRequest_Report)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_isglb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_isglb_proto_goTypes,
		DependencyIndexes: file_proto_isglb_proto_depIdxs,
		EnumInfos:         file_proto_isglb_proto_enumTypes,
		MessageInfos:      file_proto_isglb_proto_msgTypes,
	}.Build()
	File_proto_isglb_proto = out.File
	file_proto_isglb_proto_rawDesc = nil
	file_proto_isglb_proto_goTypes = nil
	file_proto_isglb_proto_depIdxs = nil
}
