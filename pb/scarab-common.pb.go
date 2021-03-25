// This file is part of scarab.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: scarab-common.proto

package scarab

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type WorkerDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *WorkerDetails) Reset() {
	*x = WorkerDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerDetails) ProtoMessage() {}

func (x *WorkerDetails) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerDetails.ProtoReflect.Descriptor instead.
func (*WorkerDetails) Descriptor() ([]byte, []int) {
	return file_scarab_common_proto_rawDescGZIP(), []int{0}
}

func (x *WorkerDetails) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *WorkerDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RegisteredProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec              *ProfileSpec         `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	FirstRegistration *timestamp.Timestamp `protobuf:"bytes,2,opt,name=firstRegistration,proto3" json:"firstRegistration,omitempty"`
	Workers           []*WorkerDetails     `protobuf:"bytes,3,rep,name=workers,proto3" json:"workers,omitempty"`
}

func (x *RegisteredProfile) Reset() {
	*x = RegisteredProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisteredProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisteredProfile) ProtoMessage() {}

func (x *RegisteredProfile) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisteredProfile.ProtoReflect.Descriptor instead.
func (*RegisteredProfile) Descriptor() ([]byte, []int) {
	return file_scarab_common_proto_rawDescGZIP(), []int{1}
}

func (x *RegisteredProfile) GetSpec() *ProfileSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *RegisteredProfile) GetFirstRegistration() *timestamp.Timestamp {
	if x != nil {
		return x.FirstRegistration
	}
	return nil
}

func (x *RegisteredProfile) GetWorkers() []*WorkerDetails {
	if x != nil {
		return x.Workers
	}
	return nil
}

type ProfileArg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Default     *JobArgValue `protobuf:"bytes,3,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *ProfileArg) Reset() {
	*x = ProfileArg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileArg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileArg) ProtoMessage() {}

func (x *ProfileArg) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_common_proto_msgTypes[2]
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
	return file_scarab_common_proto_rawDescGZIP(), []int{2}
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

func (x *ProfileArg) GetDefault() *JobArgValue {
	if x != nil {
		return x.Default
	}
	return nil
}

type ProfileSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile     string        `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Version     string        `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Description string        `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Args        []*ProfileArg `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *ProfileSpec) Reset() {
	*x = ProfileSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileSpec) ProtoMessage() {}

func (x *ProfileSpec) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileSpec.ProtoReflect.Descriptor instead.
func (*ProfileSpec) Descriptor() ([]byte, []int) {
	return file_scarab_common_proto_rawDescGZIP(), []int{3}
}

func (x *ProfileSpec) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *ProfileSpec) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ProfileSpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProfileSpec) GetArgs() []*ProfileArg {
	if x != nil {
		return x.Args
	}
	return nil
}

type JobArgOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Option   []string `protobuf:"bytes,1,rep,name=option,proto3" json:"option,omitempty"`
	Multiple bool     `protobuf:"varint,2,opt,name=multiple,proto3" json:"multiple,omitempty"`
}

func (x *JobArgOption) Reset() {
	*x = JobArgOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobArgOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobArgOption) ProtoMessage() {}

func (x *JobArgOption) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobArgOption.ProtoReflect.Descriptor instead.
func (*JobArgOption) Descriptor() ([]byte, []int) {
	return file_scarab_common_proto_rawDescGZIP(), []int{4}
}

func (x *JobArgOption) GetOption() []string {
	if x != nil {
		return x.Option
	}
	return nil
}

func (x *JobArgOption) GetMultiple() bool {
	if x != nil {
		return x.Multiple
	}
	return false
}

type JobArgValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*JobArgValue_String_
	//	*JobArgValue_Float
	//	*JobArgValue_Bool
	//	*JobArgValue_Option
	Value isJobArgValue_Value `protobuf_oneof:"value"`
}

func (x *JobArgValue) Reset() {
	*x = JobArgValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobArgValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobArgValue) ProtoMessage() {}

func (x *JobArgValue) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobArgValue.ProtoReflect.Descriptor instead.
func (*JobArgValue) Descriptor() ([]byte, []int) {
	return file_scarab_common_proto_rawDescGZIP(), []int{5}
}

func (m *JobArgValue) GetValue() isJobArgValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *JobArgValue) GetString_() string {
	if x, ok := x.GetValue().(*JobArgValue_String_); ok {
		return x.String_
	}
	return ""
}

func (x *JobArgValue) GetFloat() float32 {
	if x, ok := x.GetValue().(*JobArgValue_Float); ok {
		return x.Float
	}
	return 0
}

func (x *JobArgValue) GetBool() bool {
	if x, ok := x.GetValue().(*JobArgValue_Bool); ok {
		return x.Bool
	}
	return false
}

func (x *JobArgValue) GetOption() *JobArgOption {
	if x, ok := x.GetValue().(*JobArgValue_Option); ok {
		return x.Option
	}
	return nil
}

type isJobArgValue_Value interface {
	isJobArgValue_Value()
}

type JobArgValue_String_ struct {
	String_ string `protobuf:"bytes,1,opt,name=string,proto3,oneof"`
}

type JobArgValue_Float struct {
	Float float32 `protobuf:"fixed32,2,opt,name=float,proto3,oneof"`
}

type JobArgValue_Bool struct {
	Bool bool `protobuf:"varint,3,opt,name=bool,proto3,oneof"`
}

type JobArgValue_Option struct {
	Option *JobArgOption `protobuf:"bytes,4,opt,name=option,proto3,oneof"`
}

func (*JobArgValue_String_) isJobArgValue_Value() {}

func (*JobArgValue_Float) isJobArgValue_Value() {}

func (*JobArgValue_Bool) isJobArgValue_Value() {}

func (*JobArgValue_Option) isJobArgValue_Value() {}

type JobArg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value *JobArgValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *JobArg) Reset() {
	*x = JobArg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobArg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobArg) ProtoMessage() {}

func (x *JobArg) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_common_proto_msgTypes[6]
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
	return file_scarab_common_proto_rawDescGZIP(), []int{6}
}

func (x *JobArg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobArg) GetValue() *JobArgValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type StartJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile string    `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Version string    `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Args    []*JobArg `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *StartJobRequest) Reset() {
	*x = StartJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartJobRequest) ProtoMessage() {}

func (x *StartJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartJobRequest.ProtoReflect.Descriptor instead.
func (*StartJobRequest) Descriptor() ([]byte, []int) {
	return file_scarab_common_proto_rawDescGZIP(), []int{7}
}

func (x *StartJobRequest) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *StartJobRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *StartJobRequest) GetArgs() []*JobArg {
	if x != nil {
		return x.Args
	}
	return nil
}

type StartJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StartJobResponse) Reset() {
	*x = StartJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartJobResponse) ProtoMessage() {}

func (x *StartJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartJobResponse.ProtoReflect.Descriptor instead.
func (*StartJobResponse) Descriptor() ([]byte, []int) {
	return file_scarab_common_proto_rawDescGZIP(), []int{8}
}

func (x *StartJobResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_scarab_common_proto protoreflect.FileDescriptor

var file_scarab_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37,
	0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x27, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x63,
	0x61, 0x72, 0x61, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x48, 0x0a, 0x11, 0x66, 0x69, 0x72, 0x73, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2f, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x73, 0x22, 0x71, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x72, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e,
	0x4a, 0x6f, 0x62, 0x41, 0x72, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x22, 0x8b, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x04, 0x61, 0x72,
	0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61,
	0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x72, 0x67, 0x52, 0x04, 0x61, 0x72,
	0x67, 0x73, 0x22, 0x42, 0x0a, 0x0c, 0x4a, 0x6f, 0x62, 0x41, 0x72, 0x67, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x41, 0x72,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x12, 0x16, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x48,
	0x00, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x2e,
	0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x41, 0x72, 0x67, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x47, 0x0a, 0x06, 0x4a, 0x6f, 0x62, 0x41, 0x72,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x4a, 0x6f,
	0x62, 0x41, 0x72, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x69, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x4a,
	0x6f, 0x62, 0x41, 0x72, 0x67, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42,
	0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63,
	0x6f, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2f, 0x73, 0x63,
	0x61, 0x72, 0x61, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scarab_common_proto_rawDescOnce sync.Once
	file_scarab_common_proto_rawDescData = file_scarab_common_proto_rawDesc
)

func file_scarab_common_proto_rawDescGZIP() []byte {
	file_scarab_common_proto_rawDescOnce.Do(func() {
		file_scarab_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_scarab_common_proto_rawDescData)
	})
	return file_scarab_common_proto_rawDescData
}

var file_scarab_common_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_scarab_common_proto_goTypes = []interface{}{
	(*WorkerDetails)(nil),       // 0: scarab.WorkerDetails
	(*RegisteredProfile)(nil),   // 1: scarab.RegisteredProfile
	(*ProfileArg)(nil),          // 2: scarab.ProfileArg
	(*ProfileSpec)(nil),         // 3: scarab.ProfileSpec
	(*JobArgOption)(nil),        // 4: scarab.JobArgOption
	(*JobArgValue)(nil),         // 5: scarab.JobArgValue
	(*JobArg)(nil),              // 6: scarab.JobArg
	(*StartJobRequest)(nil),     // 7: scarab.StartJobRequest
	(*StartJobResponse)(nil),    // 8: scarab.StartJobResponse
	(*timestamp.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_scarab_common_proto_depIdxs = []int32{
	3, // 0: scarab.RegisteredProfile.spec:type_name -> scarab.ProfileSpec
	9, // 1: scarab.RegisteredProfile.firstRegistration:type_name -> google.protobuf.Timestamp
	0, // 2: scarab.RegisteredProfile.workers:type_name -> scarab.WorkerDetails
	5, // 3: scarab.ProfileArg.default:type_name -> scarab.JobArgValue
	2, // 4: scarab.ProfileSpec.args:type_name -> scarab.ProfileArg
	4, // 5: scarab.JobArgValue.option:type_name -> scarab.JobArgOption
	5, // 6: scarab.JobArg.value:type_name -> scarab.JobArgValue
	6, // 7: scarab.StartJobRequest.args:type_name -> scarab.JobArg
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_scarab_common_proto_init() }
func file_scarab_common_proto_init() {
	if File_scarab_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scarab_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerDetails); i {
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
		file_scarab_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisteredProfile); i {
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
		file_scarab_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_scarab_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileSpec); i {
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
		file_scarab_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobArgOption); i {
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
		file_scarab_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobArgValue); i {
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
		file_scarab_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_scarab_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartJobRequest); i {
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
		file_scarab_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartJobResponse); i {
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
	file_scarab_common_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*JobArgValue_String_)(nil),
		(*JobArgValue_Float)(nil),
		(*JobArgValue_Bool)(nil),
		(*JobArgValue_Option)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scarab_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scarab_common_proto_goTypes,
		DependencyIndexes: file_scarab_common_proto_depIdxs,
		MessageInfos:      file_scarab_common_proto_msgTypes,
	}.Build()
	File_scarab_common_proto = out.File
	file_scarab_common_proto_rawDesc = nil
	file_scarab_common_proto_goTypes = nil
	file_scarab_common_proto_depIdxs = nil
}
