// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: protocomponent/component1.0.proto

package protocomponent

import (
	protopkg "github.com/4rchr4y/goray/internal/proto/protopkg"
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

type Heartbeat_Status int32

const (
	Heartbeat_UNKNOWN Heartbeat_Status = 0
	Heartbeat_OK      Heartbeat_Status = 1
)

// Enum value maps for Heartbeat_Status.
var (
	Heartbeat_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "OK",
	}
	Heartbeat_Status_value = map[string]int32{
		"UNKNOWN": 0,
		"OK":      1,
	}
)

func (x Heartbeat_Status) Enum() *Heartbeat_Status {
	p := new(Heartbeat_Status)
	*p = x
	return p
}

func (x Heartbeat_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Heartbeat_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protocomponent_component1_0_proto_enumTypes[0].Descriptor()
}

func (Heartbeat_Status) Type() protoreflect.EnumType {
	return &file_protocomponent_component1_0_proto_enumTypes[0]
}

func (x Heartbeat_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Heartbeat_Status.Descriptor instead.
func (Heartbeat_Status) EnumDescriptor() ([]byte, []int) {
	return file_protocomponent_component1_0_proto_rawDescGZIP(), []int{0, 0}
}

type Heartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocomponent_component1_0_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_protocomponent_component1_0_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_protocomponent_component1_0_proto_rawDescGZIP(), []int{0}
}

type Configure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Configure) Reset() {
	*x = Configure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocomponent_component1_0_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configure) ProtoMessage() {}

func (x *Configure) ProtoReflect() protoreflect.Message {
	mi := &file_protocomponent_component1_0_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configure.ProtoReflect.Descriptor instead.
func (*Configure) Descriptor() ([]byte, []int) {
	return file_protocomponent_component1_0_proto_rawDescGZIP(), []int{1}
}

type DescribeSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DescribeSchema) Reset() {
	*x = DescribeSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocomponent_component1_0_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSchema) ProtoMessage() {}

func (x *DescribeSchema) ProtoReflect() protoreflect.Message {
	mi := &file_protocomponent_component1_0_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSchema.ProtoReflect.Descriptor instead.
func (*DescribeSchema) Descriptor() ([]byte, []int) {
	return file_protocomponent_component1_0_proto_rawDescGZIP(), []int{2}
}

type Stop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Stop) Reset() {
	*x = Stop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocomponent_component1_0_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stop) ProtoMessage() {}

func (x *Stop) ProtoReflect() protoreflect.Message {
	mi := &file_protocomponent_component1_0_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stop.ProtoReflect.Descriptor instead.
func (*Stop) Descriptor() ([]byte, []int) {
	return file_protocomponent_component1_0_proto_rawDescGZIP(), []int{3}
}

type Heartbeat_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Heartbeat_Request) Reset() {
	*x = Heartbeat_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocomponent_component1_0_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat_Request) ProtoMessage() {}

func (x *Heartbeat_Request) ProtoReflect() protoreflect.Message {
	mi := &file_protocomponent_component1_0_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat_Request.ProtoReflect.Descriptor instead.
func (*Heartbeat_Request) Descriptor() ([]byte, []int) {
	return file_protocomponent_component1_0_proto_rawDescGZIP(), []int{0, 0}
}

type Heartbeat_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      Heartbeat_Status       `protobuf:"varint,1,opt,name=status,proto3,enum=protocomponent.Heartbeat_Status" json:"status,omitempty"`
	Diagnostics []*protopkg.Diagnostic `protobuf:"bytes,2,rep,name=diagnostics,proto3" json:"diagnostics,omitempty"`
}

func (x *Heartbeat_Response) Reset() {
	*x = Heartbeat_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocomponent_component1_0_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat_Response) ProtoMessage() {}

func (x *Heartbeat_Response) ProtoReflect() protoreflect.Message {
	mi := &file_protocomponent_component1_0_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat_Response.ProtoReflect.Descriptor instead.
func (*Heartbeat_Response) Descriptor() ([]byte, []int) {
	return file_protocomponent_component1_0_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Heartbeat_Response) GetStatus() Heartbeat_Status {
	if x != nil {
		return x.Status
	}
	return Heartbeat_UNKNOWN
}

func (x *Heartbeat_Response) GetDiagnostics() []*protopkg.Diagnostic {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

type Configure_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msgpack []byte `protobuf:"bytes,1,opt,name=msgpack,proto3" json:"msgpack,omitempty"`
}

func (x *Configure_Request) Reset() {
	*x = Configure_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocomponent_component1_0_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configure_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configure_Request) ProtoMessage() {}

func (x *Configure_Request) ProtoReflect() protoreflect.Message {
	mi := &file_protocomponent_component1_0_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configure_Request.ProtoReflect.Descriptor instead.
func (*Configure_Request) Descriptor() ([]byte, []int) {
	return file_protocomponent_component1_0_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Configure_Request) GetMsgpack() []byte {
	if x != nil {
		return x.Msgpack
	}
	return nil
}

type Configure_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Diagnostics []*protopkg.Diagnostic `protobuf:"bytes,1,rep,name=diagnostics,proto3" json:"diagnostics,omitempty"`
}

func (x *Configure_Response) Reset() {
	*x = Configure_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocomponent_component1_0_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configure_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configure_Response) ProtoMessage() {}

func (x *Configure_Response) ProtoReflect() protoreflect.Message {
	mi := &file_protocomponent_component1_0_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configure_Response.ProtoReflect.Descriptor instead.
func (*Configure_Response) Descriptor() ([]byte, []int) {
	return file_protocomponent_component1_0_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Configure_Response) GetDiagnostics() []*protopkg.Diagnostic {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

type DescribeSchema_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DescribeSchema_Request) Reset() {
	*x = DescribeSchema_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocomponent_component1_0_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSchema_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSchema_Request) ProtoMessage() {}

func (x *DescribeSchema_Request) ProtoReflect() protoreflect.Message {
	mi := &file_protocomponent_component1_0_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSchema_Request.ProtoReflect.Descriptor instead.
func (*DescribeSchema_Request) Descriptor() ([]byte, []int) {
	return file_protocomponent_component1_0_proto_rawDescGZIP(), []int{2, 0}
}

type DescribeSchema_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver      *protopkg.Schema       `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Diagnostics []*protopkg.Diagnostic `protobuf:"bytes,2,rep,name=diagnostics,proto3" json:"diagnostics,omitempty"`
}

func (x *DescribeSchema_Response) Reset() {
	*x = DescribeSchema_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocomponent_component1_0_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSchema_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSchema_Response) ProtoMessage() {}

func (x *DescribeSchema_Response) ProtoReflect() protoreflect.Message {
	mi := &file_protocomponent_component1_0_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSchema_Response.ProtoReflect.Descriptor instead.
func (*DescribeSchema_Response) Descriptor() ([]byte, []int) {
	return file_protocomponent_component1_0_proto_rawDescGZIP(), []int{2, 1}
}

func (x *DescribeSchema_Response) GetDriver() *protopkg.Schema {
	if x != nil {
		return x.Driver
	}
	return nil
}

func (x *DescribeSchema_Response) GetDiagnostics() []*protopkg.Diagnostic {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

type Stop_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Stop_Request) Reset() {
	*x = Stop_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocomponent_component1_0_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stop_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stop_Request) ProtoMessage() {}

func (x *Stop_Request) ProtoReflect() protoreflect.Message {
	mi := &file_protocomponent_component1_0_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stop_Request.ProtoReflect.Descriptor instead.
func (*Stop_Request) Descriptor() ([]byte, []int) {
	return file_protocomponent_component1_0_proto_rawDescGZIP(), []int{3, 0}
}

type Stop_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Stop_Response) Reset() {
	*x = Stop_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocomponent_component1_0_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stop_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stop_Response) ProtoMessage() {}

func (x *Stop_Response) ProtoReflect() protoreflect.Message {
	mi := &file_protocomponent_component1_0_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stop_Response.ProtoReflect.Descriptor instead.
func (*Stop_Response) Descriptor() ([]byte, []int) {
	return file_protocomponent_component1_0_proto_rawDescGZIP(), []int{3, 1}
}

func (x *Stop_Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_protocomponent_component1_0_proto protoreflect.FileDescriptor

var file_protocomponent_component1_0_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x31, 0x2e, 0x30, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x7c,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x70, 0x6b, 0x67, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x52,
	0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x22, 0x1d, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x22, 0x74, 0x0a, 0x09, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x1a, 0x23, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x70, 0x61, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x70, 0x61, 0x63, 0x6b, 0x1a, 0x42, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x64, 0x69, 0x61,
	0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x6b, 0x67, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f,
	0x73, 0x74, 0x69, 0x63, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x22, 0x89, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x6c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x70, 0x6b, 0x67, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x06, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x70, 0x6b, 0x67, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63,
	0x52, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x22, 0x33, 0x0a,
	0x04, 0x53, 0x74, 0x6f, 0x70, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0xdb, 0x02, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x12, 0x52, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x26, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x53,
	0x74, 0x6f, 0x70, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x34,
	0x72, 0x63, 0x68, 0x72, 0x34, 0x79, 0x2f, 0x67, 0x6f, 0x72, 0x61, 0x79, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protocomponent_component1_0_proto_rawDescOnce sync.Once
	file_protocomponent_component1_0_proto_rawDescData = file_protocomponent_component1_0_proto_rawDesc
)

func file_protocomponent_component1_0_proto_rawDescGZIP() []byte {
	file_protocomponent_component1_0_proto_rawDescOnce.Do(func() {
		file_protocomponent_component1_0_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocomponent_component1_0_proto_rawDescData)
	})
	return file_protocomponent_component1_0_proto_rawDescData
}

var file_protocomponent_component1_0_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protocomponent_component1_0_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_protocomponent_component1_0_proto_goTypes = []interface{}{
	(Heartbeat_Status)(0),           // 0: protocomponent.Heartbeat.Status
	(*Heartbeat)(nil),               // 1: protocomponent.Heartbeat
	(*Configure)(nil),               // 2: protocomponent.Configure
	(*DescribeSchema)(nil),          // 3: protocomponent.DescribeSchema
	(*Stop)(nil),                    // 4: protocomponent.Stop
	(*Heartbeat_Request)(nil),       // 5: protocomponent.Heartbeat.Request
	(*Heartbeat_Response)(nil),      // 6: protocomponent.Heartbeat.Response
	(*Configure_Request)(nil),       // 7: protocomponent.Configure.Request
	(*Configure_Response)(nil),      // 8: protocomponent.Configure.Response
	(*DescribeSchema_Request)(nil),  // 9: protocomponent.DescribeSchema.Request
	(*DescribeSchema_Response)(nil), // 10: protocomponent.DescribeSchema.Response
	(*Stop_Request)(nil),            // 11: protocomponent.Stop.Request
	(*Stop_Response)(nil),           // 12: protocomponent.Stop.Response
	(*protopkg.Diagnostic)(nil),     // 13: protopkg.Diagnostic
	(*protopkg.Schema)(nil),         // 14: protopkg.Schema
}
var file_protocomponent_component1_0_proto_depIdxs = []int32{
	0,  // 0: protocomponent.Heartbeat.Response.status:type_name -> protocomponent.Heartbeat.Status
	13, // 1: protocomponent.Heartbeat.Response.diagnostics:type_name -> protopkg.Diagnostic
	13, // 2: protocomponent.Configure.Response.diagnostics:type_name -> protopkg.Diagnostic
	14, // 3: protocomponent.DescribeSchema.Response.driver:type_name -> protopkg.Schema
	13, // 4: protocomponent.DescribeSchema.Response.diagnostics:type_name -> protopkg.Diagnostic
	5,  // 5: protocomponent.Component.Heartbeat:input_type -> protocomponent.Heartbeat.Request
	7,  // 6: protocomponent.Component.Configure:input_type -> protocomponent.Configure.Request
	9,  // 7: protocomponent.Component.DescribeSchema:input_type -> protocomponent.DescribeSchema.Request
	11, // 8: protocomponent.Component.Stop:input_type -> protocomponent.Stop.Request
	6,  // 9: protocomponent.Component.Heartbeat:output_type -> protocomponent.Heartbeat.Response
	8,  // 10: protocomponent.Component.Configure:output_type -> protocomponent.Configure.Response
	10, // 11: protocomponent.Component.DescribeSchema:output_type -> protocomponent.DescribeSchema.Response
	12, // 12: protocomponent.Component.Stop:output_type -> protocomponent.Stop.Response
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_protocomponent_component1_0_proto_init() }
func file_protocomponent_component1_0_proto_init() {
	if File_protocomponent_component1_0_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocomponent_component1_0_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat); i {
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
		file_protocomponent_component1_0_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configure); i {
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
		file_protocomponent_component1_0_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSchema); i {
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
		file_protocomponent_component1_0_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stop); i {
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
		file_protocomponent_component1_0_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat_Request); i {
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
		file_protocomponent_component1_0_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat_Response); i {
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
		file_protocomponent_component1_0_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configure_Request); i {
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
		file_protocomponent_component1_0_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configure_Response); i {
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
		file_protocomponent_component1_0_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSchema_Request); i {
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
		file_protocomponent_component1_0_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSchema_Response); i {
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
		file_protocomponent_component1_0_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stop_Request); i {
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
		file_protocomponent_component1_0_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stop_Response); i {
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
			RawDescriptor: file_protocomponent_component1_0_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocomponent_component1_0_proto_goTypes,
		DependencyIndexes: file_protocomponent_component1_0_proto_depIdxs,
		EnumInfos:         file_protocomponent_component1_0_proto_enumTypes,
		MessageInfos:      file_protocomponent_component1_0_proto_msgTypes,
	}.Build()
	File_protocomponent_component1_0_proto = out.File
	file_protocomponent_component1_0_proto_rawDesc = nil
	file_protocomponent_component1_0_proto_goTypes = nil
	file_protocomponent_component1_0_proto_depIdxs = nil
}
