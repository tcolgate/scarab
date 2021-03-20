// This file is part of scarab.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: scarab.proto

package scarab

import (
	proto "github.com/golang/protobuf/proto"
	_go "github.com/prometheus/client_model/go"
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

type ProfileArgOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Option []string `protobuf:"bytes,1,rep,name=option,proto3" json:"option,omitempty"`
}

func (x *ProfileArgOption) Reset() {
	*x = ProfileArgOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileArgOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileArgOption) ProtoMessage() {}

func (x *ProfileArgOption) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileArgOption.ProtoReflect.Descriptor instead.
func (*ProfileArgOption) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{0}
}

func (x *ProfileArgOption) GetOption() []string {
	if x != nil {
		return x.Option
	}
	return nil
}

type ProfileArg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Types that are assignable to Default:
	//	*ProfileArg_String_
	//	*ProfileArg_Float
	//	*ProfileArg_Option
	Default isProfileArg_Default `protobuf_oneof:"default"`
}

func (x *ProfileArg) Reset() {
	*x = ProfileArg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileArg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileArg) ProtoMessage() {}

func (x *ProfileArg) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileArg.ProtoReflect.Descriptor instead.
func (*ProfileArg) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{1}
}

func (x *ProfileArg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProfileArg) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (m *ProfileArg) GetDefault() isProfileArg_Default {
	if m != nil {
		return m.Default
	}
	return nil
}

func (x *ProfileArg) GetString_() string {
	if x, ok := x.GetDefault().(*ProfileArg_String_); ok {
		return x.String_
	}
	return ""
}

func (x *ProfileArg) GetFloat() float32 {
	if x, ok := x.GetDefault().(*ProfileArg_Float); ok {
		return x.Float
	}
	return 0
}

func (x *ProfileArg) GetOption() *ProfileArgOption {
	if x, ok := x.GetDefault().(*ProfileArg_Option); ok {
		return x.Option
	}
	return nil
}

type isProfileArg_Default interface {
	isProfileArg_Default()
}

type ProfileArg_String_ struct {
	String_ string `protobuf:"bytes,3,opt,name=string,proto3,oneof"`
}

type ProfileArg_Float struct {
	Float float32 `protobuf:"fixed32,4,opt,name=float,proto3,oneof"`
}

type ProfileArg_Option struct {
	Option *ProfileArgOption `protobuf:"bytes,5,opt,name=option,proto3,oneof"`
}

func (*ProfileArg_String_) isProfileArg_Default() {}

func (*ProfileArg_Float) isProfileArg_Default() {}

func (*ProfileArg_Option) isProfileArg_Default() {}

type RegisterProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile string        `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Addr    string        `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Name    string        `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Args    []*ProfileArg `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *RegisterProfileRequest) Reset() {
	*x = RegisterProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterProfileRequest) ProtoMessage() {}

func (x *RegisterProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterProfileRequest.ProtoReflect.Descriptor instead.
func (*RegisterProfileRequest) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterProfileRequest) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *RegisterProfileRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *RegisterProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterProfileRequest) GetArgs() []*ProfileArg {
	if x != nil {
		return x.Args
	}
	return nil
}

type KeepAlive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KeepAlive) Reset() {
	*x = KeepAlive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeepAlive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeepAlive) ProtoMessage() {}

func (x *KeepAlive) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeepAlive.ProtoReflect.Descriptor instead.
func (*KeepAlive) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{3}
}

type ReportLoadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interval float32 `protobuf:"fixed32,1,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *ReportLoadRequest) Reset() {
	*x = ReportLoadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportLoadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportLoadRequest) ProtoMessage() {}

func (x *ReportLoadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportLoadRequest.ProtoReflect.Descriptor instead.
func (*ReportLoadRequest) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{4}
}

func (x *ReportLoadRequest) GetInterval() float32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

type LoadMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*_go.MetricFamily `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *LoadMetrics) Reset() {
	*x = LoadMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadMetrics) ProtoMessage() {}

func (x *LoadMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadMetrics.ProtoReflect.Descriptor instead.
func (*LoadMetrics) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{5}
}

func (x *LoadMetrics) GetMetrics() []*_go.MetricFamily {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type JobArg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Value:
	//	*JobArg_String_
	//	*JobArg_Float
	Value isJobArg_Value `protobuf_oneof:"value"`
}

func (x *JobArg) Reset() {
	*x = JobArg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobArg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobArg) ProtoMessage() {}

func (x *JobArg) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobArg.ProtoReflect.Descriptor instead.
func (*JobArg) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{6}
}

func (x *JobArg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *JobArg) GetValue() isJobArg_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *JobArg) GetString_() string {
	if x, ok := x.GetValue().(*JobArg_String_); ok {
		return x.String_
	}
	return ""
}

func (x *JobArg) GetFloat() float32 {
	if x, ok := x.GetValue().(*JobArg_Float); ok {
		return x.Float
	}
	return 0
}

type isJobArg_Value interface {
	isJobArg_Value()
}

type JobArg_String_ struct {
	String_ string `protobuf:"bytes,2,opt,name=string,proto3,oneof"`
}

type JobArg_Float struct {
	Float float32 `protobuf:"fixed32,3,opt,name=float,proto3,oneof"`
}

func (*JobArg_String_) isJobArg_Value() {}

func (*JobArg_Float) isJobArg_Value() {}

type RunJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile  string    `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Interval float32   `protobuf:"fixed32,2,opt,name=interval,proto3" json:"interval,omitempty"`
	Maxrps   float32   `protobuf:"fixed32,3,opt,name=maxrps,proto3" json:"maxrps,omitempty"`
	Args     []*JobArg `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *RunJobRequest) Reset() {
	*x = RunJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunJobRequest) ProtoMessage() {}

func (x *RunJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunJobRequest.ProtoReflect.Descriptor instead.
func (*RunJobRequest) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{7}
}

func (x *RunJobRequest) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *RunJobRequest) GetInterval() float32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *RunJobRequest) GetMaxrps() float32 {
	if x != nil {
		return x.Maxrps
	}
	return 0
}

func (x *RunJobRequest) GetArgs() []*JobArg {
	if x != nil {
		return x.Args
	}
	return nil
}

type JobMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*_go.MetricFamily `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *JobMetrics) Reset() {
	*x = JobMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobMetrics) ProtoMessage() {}

func (x *JobMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobMetrics.ProtoReflect.Descriptor instead.
func (*JobMetrics) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{8}
}

func (x *JobMetrics) GetMetrics() []*_go.MetricFamily {
	if x != nil {
		return x.Metrics
	}
	return nil
}

var File_scarab_proto protoreflect.FileDescriptor

var file_scarab_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x1a, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x41, 0x72, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xb3, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x72, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x12, 0x16, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x48,
	0x00, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61,
	0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x72, 0x67, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x41, 0x72, 0x67, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x0b, 0x0a, 0x09,
	0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x22, 0x2f, 0x0a, 0x11, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x4b, 0x0a, 0x0b, 0x4c, 0x6f,
	0x61, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x07,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x57, 0x0a, 0x06, 0x4a, 0x6f, 0x62, 0x41, 0x72,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00,
	0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x81, 0x01, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x72,
	0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x72, 0x70, 0x73,
	0x12, 0x22, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x41, 0x72, 0x67, 0x52, 0x04,
	0x61, 0x72, 0x67, 0x73, 0x22, 0x4a, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68,
	0x65, 0x75, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x32, 0x53, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x0f, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1e,
	0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x32, 0x85, 0x01, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x12, 0x40, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x19,
	0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x63, 0x61, 0x72,
	0x61, 0x62, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x39, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x15, 0x2e, 0x73,
	0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x4a, 0x6f, 0x62,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x23, 0x5a,
	0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6f, 0x6c,
	0x67, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2f, 0x73, 0x63, 0x61, 0x72,
	0x61, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scarab_proto_rawDescOnce sync.Once
	file_scarab_proto_rawDescData = file_scarab_proto_rawDesc
)

func file_scarab_proto_rawDescGZIP() []byte {
	file_scarab_proto_rawDescOnce.Do(func() {
		file_scarab_proto_rawDescData = protoimpl.X.CompressGZIP(file_scarab_proto_rawDescData)
	})
	return file_scarab_proto_rawDescData
}

var file_scarab_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_scarab_proto_goTypes = []interface{}{
	(*ProfileArgOption)(nil),       // 0: scarab.ProfileArgOption
	(*ProfileArg)(nil),             // 1: scarab.ProfileArg
	(*RegisterProfileRequest)(nil), // 2: scarab.RegisterProfileRequest
	(*KeepAlive)(nil),              // 3: scarab.KeepAlive
	(*ReportLoadRequest)(nil),      // 4: scarab.ReportLoadRequest
	(*LoadMetrics)(nil),            // 5: scarab.LoadMetrics
	(*JobArg)(nil),                 // 6: scarab.JobArg
	(*RunJobRequest)(nil),          // 7: scarab.RunJobRequest
	(*JobMetrics)(nil),             // 8: scarab.JobMetrics
	(*_go.MetricFamily)(nil),       // 9: io.prometheus.client.MetricFamily
}
var file_scarab_proto_depIdxs = []int32{
	0, // 0: scarab.ProfileArg.option:type_name -> scarab.ProfileArgOption
	1, // 1: scarab.RegisterProfileRequest.args:type_name -> scarab.ProfileArg
	9, // 2: scarab.LoadMetrics.metrics:type_name -> io.prometheus.client.MetricFamily
	6, // 3: scarab.RunJobRequest.args:type_name -> scarab.JobArg
	9, // 4: scarab.JobMetrics.metrics:type_name -> io.prometheus.client.MetricFamily
	2, // 5: scarab.Manager.RegisterProfile:input_type -> scarab.RegisterProfileRequest
	4, // 6: scarab.Worker.ReportLoad:input_type -> scarab.ReportLoadRequest
	7, // 7: scarab.Worker.RunJob:input_type -> scarab.RunJobRequest
	3, // 8: scarab.Manager.RegisterProfile:output_type -> scarab.KeepAlive
	5, // 9: scarab.Worker.ReportLoad:output_type -> scarab.LoadMetrics
	8, // 10: scarab.Worker.RunJob:output_type -> scarab.JobMetrics
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_scarab_proto_init() }
func file_scarab_proto_init() {
	if File_scarab_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scarab_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileArgOption); i {
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
		file_scarab_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileArg); i {
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
		file_scarab_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterProfileRequest); i {
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
		file_scarab_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeepAlive); i {
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
		file_scarab_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportLoadRequest); i {
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
		file_scarab_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadMetrics); i {
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
		file_scarab_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobArg); i {
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
		file_scarab_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunJobRequest); i {
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
		file_scarab_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobMetrics); i {
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
	file_scarab_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ProfileArg_String_)(nil),
		(*ProfileArg_Float)(nil),
		(*ProfileArg_Option)(nil),
	}
	file_scarab_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*JobArg_String_)(nil),
		(*JobArg_Float)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scarab_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_scarab_proto_goTypes,
		DependencyIndexes: file_scarab_proto_depIdxs,
		MessageInfos:      file_scarab_proto_msgTypes,
	}.Build()
	File_scarab_proto = out.File
	file_scarab_proto_rawDesc = nil
	file_scarab_proto_goTypes = nil
	file_scarab_proto_depIdxs = nil
}
