// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scarab.proto

package scarab

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegisterRequest struct {
	Profile              string   `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d446afcdfbf09f, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

type KeepAlive struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeepAlive) Reset()         { *m = KeepAlive{} }
func (m *KeepAlive) String() string { return proto.CompactTextString(m) }
func (*KeepAlive) ProtoMessage()    {}
func (*KeepAlive) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d446afcdfbf09f, []int{1}
}

func (m *KeepAlive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeepAlive.Unmarshal(m, b)
}
func (m *KeepAlive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeepAlive.Marshal(b, m, deterministic)
}
func (m *KeepAlive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeepAlive.Merge(m, src)
}
func (m *KeepAlive) XXX_Size() int {
	return xxx_messageInfo_KeepAlive.Size(m)
}
func (m *KeepAlive) XXX_DiscardUnknown() {
	xxx_messageInfo_KeepAlive.DiscardUnknown(m)
}

var xxx_messageInfo_KeepAlive proto.InternalMessageInfo

type RunJobRequest struct {
	Profile              string   `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Interval             float32  `protobuf:"fixed32,2,opt,name=interval,proto3" json:"interval,omitempty"`
	Maxrps               float32  `protobuf:"fixed32,3,opt,name=maxrps,proto3" json:"maxrps,omitempty"`
	Host                 string   `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunJobRequest) Reset()         { *m = RunJobRequest{} }
func (m *RunJobRequest) String() string { return proto.CompactTextString(m) }
func (*RunJobRequest) ProtoMessage()    {}
func (*RunJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d446afcdfbf09f, []int{2}
}

func (m *RunJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunJobRequest.Unmarshal(m, b)
}
func (m *RunJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunJobRequest.Marshal(b, m, deterministic)
}
func (m *RunJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunJobRequest.Merge(m, src)
}
func (m *RunJobRequest) XXX_Size() int {
	return xxx_messageInfo_RunJobRequest.Size(m)
}
func (m *RunJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunJobRequest proto.InternalMessageInfo

func (m *RunJobRequest) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *RunJobRequest) GetInterval() float32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *RunJobRequest) GetMaxrps() float32 {
	if m != nil {
		return m.Maxrps
	}
	return 0
}

func (m *RunJobRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type Counter struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Counter) Reset()         { *m = Counter{} }
func (m *Counter) String() string { return proto.CompactTextString(m) }
func (*Counter) ProtoMessage()    {}
func (*Counter) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d446afcdfbf09f, []int{3}
}

func (m *Counter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Counter.Unmarshal(m, b)
}
func (m *Counter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Counter.Marshal(b, m, deterministic)
}
func (m *Counter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Counter.Merge(m, src)
}
func (m *Counter) XXX_Size() int {
	return xxx_messageInfo_Counter.Size(m)
}
func (m *Counter) XXX_DiscardUnknown() {
	xxx_messageInfo_Counter.DiscardUnknown(m)
}

var xxx_messageInfo_Counter proto.InternalMessageInfo

func (m *Counter) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Gauge struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Gauge) Reset()         { *m = Gauge{} }
func (m *Gauge) String() string { return proto.CompactTextString(m) }
func (*Gauge) ProtoMessage()    {}
func (*Gauge) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d446afcdfbf09f, []int{4}
}

func (m *Gauge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gauge.Unmarshal(m, b)
}
func (m *Gauge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gauge.Marshal(b, m, deterministic)
}
func (m *Gauge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gauge.Merge(m, src)
}
func (m *Gauge) XXX_Size() int {
	return xxx_messageInfo_Gauge.Size(m)
}
func (m *Gauge) XXX_DiscardUnknown() {
	xxx_messageInfo_Gauge.DiscardUnknown(m)
}

var xxx_messageInfo_Gauge proto.InternalMessageInfo

func (m *Gauge) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Rate struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	Duration             float32  `protobuf:"fixed32,2,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rate) Reset()         { *m = Rate{} }
func (m *Rate) String() string { return proto.CompactTextString(m) }
func (*Rate) ProtoMessage()    {}
func (*Rate) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d446afcdfbf09f, []int{5}
}

func (m *Rate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rate.Unmarshal(m, b)
}
func (m *Rate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rate.Marshal(b, m, deterministic)
}
func (m *Rate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rate.Merge(m, src)
}
func (m *Rate) XXX_Size() int {
	return xxx_messageInfo_Rate.Size(m)
}
func (m *Rate) XXX_DiscardUnknown() {
	xxx_messageInfo_Rate.DiscardUnknown(m)
}

var xxx_messageInfo_Rate proto.InternalMessageInfo

func (m *Rate) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Rate) GetDuration() float32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type Histogram struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Histogram) Reset()         { *m = Histogram{} }
func (m *Histogram) String() string { return proto.CompactTextString(m) }
func (*Histogram) ProtoMessage()    {}
func (*Histogram) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d446afcdfbf09f, []int{6}
}

func (m *Histogram) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Histogram.Unmarshal(m, b)
}
func (m *Histogram) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Histogram.Marshal(b, m, deterministic)
}
func (m *Histogram) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Histogram.Merge(m, src)
}
func (m *Histogram) XXX_Size() int {
	return xxx_messageInfo_Histogram.Size(m)
}
func (m *Histogram) XXX_DiscardUnknown() {
	xxx_messageInfo_Histogram.DiscardUnknown(m)
}

var xxx_messageInfo_Histogram proto.InternalMessageInfo

type Metric struct {
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Types that are valid to be assigned to U:
	//	*Metric_Counter
	//	*Metric_Gauge
	//	*Metric_Rate
	//	*Metric_Histogram
	U                    isMetric_U `protobuf_oneof:"U"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d446afcdfbf09f, []int{7}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metric) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
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

func (m *Metric) GetU() isMetric_U {
	if m != nil {
		return m.U
	}
	return nil
}

func (m *Metric) GetCounter() *Counter {
	if x, ok := m.GetU().(*Metric_Counter); ok {
		return x.Counter
	}
	return nil
}

func (m *Metric) GetGauge() *Gauge {
	if x, ok := m.GetU().(*Metric_Gauge); ok {
		return x.Gauge
	}
	return nil
}

func (m *Metric) GetRate() *Rate {
	if x, ok := m.GetU().(*Metric_Rate); ok {
		return x.Rate
	}
	return nil
}

func (m *Metric) GetHistogram() *Histogram {
	if x, ok := m.GetU().(*Metric_Histogram); ok {
		return x.Histogram
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Metric) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Metric_Counter)(nil),
		(*Metric_Gauge)(nil),
		(*Metric_Rate)(nil),
		(*Metric_Histogram)(nil),
	}
}

type JobMetrics struct {
	Metrics              []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *JobMetrics) Reset()         { *m = JobMetrics{} }
func (m *JobMetrics) String() string { return proto.CompactTextString(m) }
func (*JobMetrics) ProtoMessage()    {}
func (*JobMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d446afcdfbf09f, []int{8}
}

func (m *JobMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobMetrics.Unmarshal(m, b)
}
func (m *JobMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobMetrics.Marshal(b, m, deterministic)
}
func (m *JobMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobMetrics.Merge(m, src)
}
func (m *JobMetrics) XXX_Size() int {
	return xxx_messageInfo_JobMetrics.Size(m)
}
func (m *JobMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_JobMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_JobMetrics proto.InternalMessageInfo

func (m *JobMetrics) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "scarab.RegisterRequest")
	proto.RegisterType((*KeepAlive)(nil), "scarab.KeepAlive")
	proto.RegisterType((*RunJobRequest)(nil), "scarab.RunJobRequest")
	proto.RegisterType((*Counter)(nil), "scarab.Counter")
	proto.RegisterType((*Gauge)(nil), "scarab.Gauge")
	proto.RegisterType((*Rate)(nil), "scarab.Rate")
	proto.RegisterType((*Histogram)(nil), "scarab.Histogram")
	proto.RegisterType((*Metric)(nil), "scarab.Metric")
	proto.RegisterType((*JobMetrics)(nil), "scarab.JobMetrics")
}

func init() {
	proto.RegisterFile("scarab.proto", fileDescriptor_e7d446afcdfbf09f)
}

var fileDescriptor_e7d446afcdfbf09f = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x5d, 0x8b, 0x13, 0x31,
	0x14, 0x9d, 0xe9, 0x47, 0xba, 0xbd, 0xb3, 0xeb, 0xe2, 0xc5, 0x8f, 0x50, 0x10, 0x4b, 0x40, 0x28,
	0x2c, 0x2c, 0x3a, 0x82, 0x8a, 0x6f, 0xab, 0x82, 0xc3, 0xaa, 0x2f, 0x01, 0xf1, 0x39, 0x6d, 0xaf,
	0xdd, 0x60, 0x3b, 0x99, 0x4d, 0x32, 0xc5, 0x5f, 0xed, 0x6f, 0x90, 0x49, 0x9a, 0xe9, 0xfa, 0xb0,
	0xf8, 0x76, 0x4f, 0xce, 0xc9, 0x9d, 0x73, 0x26, 0x07, 0x4e, 0xdd, 0x4a, 0x59, 0xb5, 0xbc, 0x6c,
	0xac, 0xf1, 0x06, 0x59, 0x44, 0xe2, 0x02, 0xce, 0x25, 0x6d, 0xb4, 0xf3, 0x64, 0x25, 0xdd, 0xb6,
	0xe4, 0x3c, 0x72, 0x98, 0x34, 0xd6, 0xfc, 0xd4, 0x5b, 0xe2, 0xf9, 0x3c, 0x5f, 0x4c, 0x65, 0x82,
	0xa2, 0x80, 0xe9, 0x17, 0xa2, 0xe6, 0x6a, 0xab, 0xf7, 0x24, 0x6e, 0xe1, 0x4c, 0xb6, 0xf5, 0xb5,
	0x59, 0xfe, 0xf7, 0x1e, 0xce, 0xe0, 0x44, 0xd7, 0x9e, 0xec, 0x5e, 0x6d, 0xf9, 0x60, 0x9e, 0x2f,
	0x06, 0xb2, 0xc7, 0xf8, 0x04, 0xd8, 0x4e, 0xfd, 0xb6, 0x8d, 0xe3, 0xc3, 0xc0, 0x1c, 0x10, 0x22,
	0x8c, 0x6e, 0x8c, 0xf3, 0x7c, 0x14, 0x56, 0x85, 0x59, 0x3c, 0x87, 0xc9, 0x47, 0xd3, 0x76, 0x37,
	0xf1, 0x11, 0x8c, 0xf7, 0x6a, 0xdb, 0xc6, 0x4f, 0x0d, 0x64, 0x04, 0xe2, 0x19, 0x8c, 0x3f, 0xab,
	0x76, 0x43, 0xf7, 0xd0, 0xef, 0x60, 0x24, 0x95, 0xbf, 0x87, 0xed, 0x5c, 0xae, 0x5b, 0xab, 0xbc,
	0x36, 0x75, 0x72, 0x99, 0x70, 0x97, 0xbc, 0xd2, 0xce, 0x9b, 0x8d, 0x55, 0x3b, 0xf1, 0x27, 0x07,
	0xf6, 0x8d, 0xbc, 0xd5, 0xab, 0xce, 0x65, 0xad, 0x76, 0x29, 0x70, 0x98, 0x71, 0x0e, 0xc5, 0x9a,
	0xdc, 0xca, 0xea, 0xa6, 0x5f, 0x35, 0x95, 0x77, 0x8f, 0xf0, 0x02, 0x26, 0xab, 0x98, 0x23, 0x84,
	0x2e, 0xca, 0xf3, 0xcb, 0xc3, 0xe3, 0x1c, 0xe2, 0x55, 0x99, 0x4c, 0x0a, 0x7c, 0x01, 0xe3, 0x4d,
	0x97, 0x29, 0xfc, 0x89, 0xa2, 0x3c, 0x4b, 0xd2, 0x10, 0xb4, 0xca, 0x64, 0x64, 0x51, 0xc0, 0xc8,
	0x2a, 0x4f, 0x7c, 0x1c, 0x54, 0xa7, 0x49, 0xd5, 0xe5, 0xad, 0x32, 0x19, 0x38, 0x7c, 0x05, 0xd3,
	0x9b, 0x94, 0x82, 0xb3, 0x20, 0x7c, 0x98, 0x84, 0x7d, 0xbc, 0x2a, 0x93, 0x47, 0xd5, 0x87, 0x21,
	0xe4, 0xdf, 0xc5, 0x1b, 0x80, 0x6b, 0xb3, 0x8c, 0x91, 0x1d, 0x2e, 0x60, 0xb2, 0x8b, 0x23, 0xcf,
	0xe7, 0xc3, 0x45, 0x51, 0x3e, 0x48, 0x3b, 0xa2, 0x42, 0x26, 0xba, 0xfc, 0x04, 0xec, 0x2b, 0xa9,
	0x35, 0x59, 0x7c, 0x0f, 0x27, 0xa9, 0x66, 0xf8, 0xb4, 0xf7, 0xf6, 0x6f, 0xf1, 0x66, 0xbd, 0x97,
	0x63, 0xc9, 0xb2, 0x97, 0x79, 0x79, 0x05, 0xec, 0x87, 0xb1, 0xbf, 0xc8, 0xe2, 0x5b, 0x60, 0xb1,
	0x72, 0xf8, 0xb8, 0xdf, 0x71, 0xb7, 0x82, 0x33, 0x4c, 0xc7, 0x47, 0xbb, 0xdd, 0x8a, 0x25, 0x0b,
	0xa5, 0x7f, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x86, 0x74, 0xff, 0x04, 0x03, 0x00, 0x00,
}
