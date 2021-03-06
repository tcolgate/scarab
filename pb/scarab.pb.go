// This file is part of scarab.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: scarab.proto

package scarab

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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profiles []string `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty"`
	Addr     string   `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Name     string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetProfiles() []string {
	if x != nil {
		return x.Profiles
	}
	return nil
}

func (x *RegisterRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type KeepAlive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KeepAlive) Reset() {
	*x = KeepAlive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeepAlive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeepAlive) ProtoMessage() {}

func (x *KeepAlive) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use KeepAlive.ProtoReflect.Descriptor instead.
func (*KeepAlive) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{1}
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
		mi := &file_scarab_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportLoadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportLoadRequest) ProtoMessage() {}

func (x *ReportLoadRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ReportLoadRequest.ProtoReflect.Descriptor instead.
func (*ReportLoadRequest) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{2}
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

	Processrate float32 `protobuf:"fixed32,1,opt,name=processrate,proto3" json:"processrate,omitempty"`
	Serveridle  float32 `protobuf:"fixed32,2,opt,name=serveridle,proto3" json:"serveridle,omitempty"`
	Serverload  float32 `protobuf:"fixed32,3,opt,name=serverload,proto3" json:"serverload,omitempty"`
	Processmem  int64   `protobuf:"varint,4,opt,name=processmem,proto3" json:"processmem,omitempty"`
	Memdree     int64   `protobuf:"varint,5,opt,name=memdree,proto3" json:"memdree,omitempty"`
}

func (x *LoadMetrics) Reset() {
	*x = LoadMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadMetrics) ProtoMessage() {}

func (x *LoadMetrics) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use LoadMetrics.ProtoReflect.Descriptor instead.
func (*LoadMetrics) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{3}
}

func (x *LoadMetrics) GetProcessrate() float32 {
	if x != nil {
		return x.Processrate
	}
	return 0
}

func (x *LoadMetrics) GetServeridle() float32 {
	if x != nil {
		return x.Serveridle
	}
	return 0
}

func (x *LoadMetrics) GetServerload() float32 {
	if x != nil {
		return x.Serverload
	}
	return 0
}

func (x *LoadMetrics) GetProcessmem() int64 {
	if x != nil {
		return x.Processmem
	}
	return 0
}

func (x *LoadMetrics) GetMemdree() int64 {
	if x != nil {
		return x.Memdree
	}
	return 0
}

type RunJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile  string  `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Interval float32 `protobuf:"fixed32,2,opt,name=interval,proto3" json:"interval,omitempty"`
	Maxrps   float32 `protobuf:"fixed32,3,opt,name=maxrps,proto3" json:"maxrps,omitempty"`
	Host     string  `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *RunJobRequest) Reset() {
	*x = RunJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunJobRequest) ProtoMessage() {}

func (x *RunJobRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RunJobRequest.ProtoReflect.Descriptor instead.
func (*RunJobRequest) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{4}
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

func (x *RunJobRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type Counter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float32 `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Counter) Reset() {
	*x = Counter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Counter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Counter) ProtoMessage() {}

func (x *Counter) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Counter.ProtoReflect.Descriptor instead.
func (*Counter) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{5}
}

func (x *Counter) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Gauge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float32 `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Gauge) Reset() {
	*x = Gauge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gauge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gauge) ProtoMessage() {}

func (x *Gauge) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Gauge.ProtoReflect.Descriptor instead.
func (*Gauge) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{6}
}

func (x *Gauge) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Rate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value    float32 `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	Duration float32 `protobuf:"fixed32,2,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Rate) Reset() {
	*x = Rate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rate) ProtoMessage() {}

func (x *Rate) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Rate.ProtoReflect.Descriptor instead.
func (*Rate) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{7}
}

func (x *Rate) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Rate) GetDuration() float32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type Histogram struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Histogram) Reset() {
	*x = Histogram{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Histogram) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Histogram) ProtoMessage() {}

func (x *Histogram) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Histogram.ProtoReflect.Descriptor instead.
func (*Histogram) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{8}
}

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Types that are assignable to U:
	//	*Metric_Counter
	//	*Metric_Gauge
	//	*Metric_Rate
	//	*Metric_Histogram
	U isMetric_U `protobuf_oneof:"U"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_scarab_proto_rawDescGZIP(), []int{9}
}

func (x *Metric) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metric) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (m *Metric) GetU() isMetric_U {
	if m != nil {
		return m.U
	}
	return nil
}

func (x *Metric) GetCounter() *Counter {
	if x, ok := x.GetU().(*Metric_Counter); ok {
		return x.Counter
	}
	return nil
}

func (x *Metric) GetGauge() *Gauge {
	if x, ok := x.GetU().(*Metric_Gauge); ok {
		return x.Gauge
	}
	return nil
}

func (x *Metric) GetRate() *Rate {
	if x, ok := x.GetU().(*Metric_Rate); ok {
		return x.Rate
	}
	return nil
}

func (x *Metric) GetHistogram() *Histogram {
	if x, ok := x.GetU().(*Metric_Histogram); ok {
		return x.Histogram
	}
	return nil
}

type isMetric_U interface {
	isMetric_U()
}

type Metric_Counter struct {
	Counter *Counter `protobuf:"bytes,3,opt,name=counter,proto3,oneof"`
}

type Metric_Gauge struct {
	Gauge *Gauge `protobuf:"bytes,4,opt,name=gauge,proto3,oneof"`
}

type Metric_Rate struct {
	Rate *Rate `protobuf:"bytes,5,opt,name=rate,proto3,oneof"`
}

type Metric_Histogram struct {
	Histogram *Histogram `protobuf:"bytes,6,opt,name=histogram,proto3,oneof"`
}

func (*Metric_Counter) isMetric_U() {}

func (*Metric_Gauge) isMetric_U() {}

func (*Metric_Rate) isMetric_U() {}

func (*Metric_Histogram) isMetric_U() {}

type JobMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *JobMetrics) Reset() {
	*x = JobMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobMetrics) ProtoMessage() {}

func (x *JobMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_proto_msgTypes[10]
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
	return file_scarab_proto_rawDescGZIP(), []int{10}
}

func (x *JobMetrics) GetMetrics() []*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

var File_scarab_proto protoreflect.FileDescriptor

var file_scarab_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x22, 0x55, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0b, 0x0a,
	0x09, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x22, 0x2f, 0x0a, 0x11, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0xa9, 0x01, 0x0a, 0x0b,
	0x4c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x69, 0x64, 0x6c, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6d, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x6d, 0x64, 0x72, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x6d, 0x65, 0x6d, 0x64, 0x72, 0x65, 0x65, 0x22, 0x71, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x61, 0x78, 0x72, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x6d, 0x61, 0x78, 0x72, 0x70, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x07, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x47,
	0x61, 0x75, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x38, 0x0a, 0x04, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0b, 0x0a, 0x09, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x22, 0xee, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x25, 0x0a, 0x05, 0x67, 0x61, 0x75, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x47, 0x61, 0x75, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x05, 0x67, 0x61, 0x75, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x48, 0x00, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x48, 0x00, 0x52, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x42, 0x03, 0x0a,
	0x01, 0x55, 0x22, 0x36, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x12, 0x28, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x32, 0x44, 0x0a, 0x06, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x17, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x63, 0x61, 0x72,
	0x61, 0x62, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x32, 0x83, 0x01, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0a, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x19, 0x2e, 0x73, 0x63, 0x61, 0x72,
	0x61, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x4c, 0x6f,
	0x61, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x37, 0x0a,
	0x06, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x15, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62,
	0x2e, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x22, 0x00, 0x30, 0x01, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6f, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x63,
	0x61, 0x72, 0x61, 0x62, 0x2f, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_scarab_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_scarab_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),   // 0: scarab.RegisterRequest
	(*KeepAlive)(nil),         // 1: scarab.KeepAlive
	(*ReportLoadRequest)(nil), // 2: scarab.ReportLoadRequest
	(*LoadMetrics)(nil),       // 3: scarab.LoadMetrics
	(*RunJobRequest)(nil),     // 4: scarab.RunJobRequest
	(*Counter)(nil),           // 5: scarab.Counter
	(*Gauge)(nil),             // 6: scarab.Gauge
	(*Rate)(nil),              // 7: scarab.Rate
	(*Histogram)(nil),         // 8: scarab.Histogram
	(*Metric)(nil),            // 9: scarab.Metric
	(*JobMetrics)(nil),        // 10: scarab.JobMetrics
}
var file_scarab_proto_depIdxs = []int32{
	5,  // 0: scarab.Metric.counter:type_name -> scarab.Counter
	6,  // 1: scarab.Metric.gauge:type_name -> scarab.Gauge
	7,  // 2: scarab.Metric.rate:type_name -> scarab.Rate
	8,  // 3: scarab.Metric.histogram:type_name -> scarab.Histogram
	9,  // 4: scarab.JobMetrics.metrics:type_name -> scarab.Metric
	0,  // 5: scarab.Leader.Register:input_type -> scarab.RegisterRequest
	2,  // 6: scarab.Worker.ReportLoad:input_type -> scarab.ReportLoadRequest
	4,  // 7: scarab.Worker.RunJob:input_type -> scarab.RunJobRequest
	1,  // 8: scarab.Leader.Register:output_type -> scarab.KeepAlive
	3,  // 9: scarab.Worker.ReportLoad:output_type -> scarab.LoadMetrics
	10, // 10: scarab.Worker.RunJob:output_type -> scarab.JobMetrics
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_scarab_proto_init() }
func file_scarab_proto_init() {
	if File_scarab_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scarab_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_scarab_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_scarab_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_scarab_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_scarab_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Counter); i {
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
			switch v := v.(*Gauge); i {
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
			switch v := v.(*Rate); i {
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
			switch v := v.(*Histogram); i {
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
		file_scarab_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
		file_scarab_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
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
	file_scarab_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*Metric_Counter)(nil),
		(*Metric_Gauge)(nil),
		(*Metric_Rate)(nil),
		(*Metric_Histogram)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scarab_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
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
