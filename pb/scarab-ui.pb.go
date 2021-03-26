// This file is part of scarab.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: scarab-ui.proto

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

type ListJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListJobsRequest) Reset() {
	*x = ListJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_ui_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsRequest) ProtoMessage() {}

func (x *ListJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_ui_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsRequest.ProtoReflect.Descriptor instead.
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return file_scarab_ui_proto_rawDescGZIP(), []int{0}
}

type ListJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs []*Job `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *ListJobsResponse) Reset() {
	*x = ListJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_ui_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsResponse) ProtoMessage() {}

func (x *ListJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_ui_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsResponse.ProtoReflect.Descriptor instead.
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return file_scarab_ui_proto_rawDescGZIP(), []int{1}
}

func (x *ListJobsResponse) GetJobs() []*Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type ListProfilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProfilesRequest) Reset() {
	*x = ListProfilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_ui_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProfilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProfilesRequest) ProtoMessage() {}

func (x *ListProfilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_ui_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProfilesRequest.ProtoReflect.Descriptor instead.
func (*ListProfilesRequest) Descriptor() ([]byte, []int) {
	return file_scarab_ui_proto_rawDescGZIP(), []int{2}
}

type ListProfilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registered []*RegisteredProfile `protobuf:"bytes,1,rep,name=registered,proto3" json:"registered,omitempty"`
}

func (x *ListProfilesResponse) Reset() {
	*x = ListProfilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_ui_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProfilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProfilesResponse) ProtoMessage() {}

func (x *ListProfilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_ui_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProfilesResponse.ProtoReflect.Descriptor instead.
func (*ListProfilesResponse) Descriptor() ([]byte, []int) {
	return file_scarab_ui_proto_rawDescGZIP(), []int{3}
}

func (x *ListProfilesResponse) GetRegistered() []*RegisteredProfile {
	if x != nil {
		return x.Registered
	}
	return nil
}

type WatchActiveJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WatchActiveJobsRequest) Reset() {
	*x = WatchActiveJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_ui_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchActiveJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchActiveJobsRequest) ProtoMessage() {}

func (x *WatchActiveJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_ui_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchActiveJobsRequest.ProtoReflect.Descriptor instead.
func (*WatchActiveJobsRequest) Descriptor() ([]byte, []int) {
	return file_scarab_ui_proto_rawDescGZIP(), []int{4}
}

type WatchActiveJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WatchActiveJobsResponse) Reset() {
	*x = WatchActiveJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_ui_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchActiveJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchActiveJobsResponse) ProtoMessage() {}

func (x *WatchActiveJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_ui_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchActiveJobsResponse.ProtoReflect.Descriptor instead.
func (*WatchActiveJobsResponse) Descriptor() ([]byte, []int) {
	return file_scarab_ui_proto_rawDescGZIP(), []int{5}
}

type StopJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StopJobRequest) Reset() {
	*x = StopJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_ui_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopJobRequest) ProtoMessage() {}

func (x *StopJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_ui_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopJobRequest.ProtoReflect.Descriptor instead.
func (*StopJobRequest) Descriptor() ([]byte, []int) {
	return file_scarab_ui_proto_rawDescGZIP(), []int{6}
}

func (x *StopJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StopJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopJobResponse) Reset() {
	*x = StopJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scarab_ui_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopJobResponse) ProtoMessage() {}

func (x *StopJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scarab_ui_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopJobResponse.ProtoReflect.Descriptor instead.
func (*StopJobResponse) Descriptor() ([]byte, []int) {
	return file_scarab_ui_proto_rawDescGZIP(), []int{7}
}

var File_scarab_ui_proto protoreflect.FileDescriptor

var file_scarab_ui_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2d, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x1a, 0x13, 0x73, 0x63, 0x61, 0x72, 0x61,
	0x62, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x33, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x4a, 0x6f, 0x62,
	0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x63, 0x61, 0x72,
	0x61, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64,
	0x22, 0x18, 0x0a, 0x16, 0x57, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4a,
	0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x19, 0x0a, 0x17, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x70, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x70, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf0, 0x02, 0x0a, 0x09, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x55, 0x49, 0x12, 0x3f, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x4a, 0x6f, 0x62, 0x12, 0x17, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x07, 0x53, 0x74, 0x6f,
	0x70, 0x4a, 0x6f, 0x62, 0x12, 0x16, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x53, 0x74,
	0x6f, 0x70, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73,
	0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73,
	0x12, 0x17, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f,
	0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x63, 0x61, 0x72,
	0x61, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x1e, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61,
	0x62, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4a, 0x6f, 0x62,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x63, 0x61, 0x72, 0x61,
	0x62, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4a, 0x6f, 0x62,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x23, 0x5a,
	0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6f, 0x6c,
	0x67, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x63, 0x61, 0x72, 0x61, 0x62, 0x2f, 0x73, 0x63, 0x61, 0x72,
	0x61, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scarab_ui_proto_rawDescOnce sync.Once
	file_scarab_ui_proto_rawDescData = file_scarab_ui_proto_rawDesc
)

func file_scarab_ui_proto_rawDescGZIP() []byte {
	file_scarab_ui_proto_rawDescOnce.Do(func() {
		file_scarab_ui_proto_rawDescData = protoimpl.X.CompressGZIP(file_scarab_ui_proto_rawDescData)
	})
	return file_scarab_ui_proto_rawDescData
}

var file_scarab_ui_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_scarab_ui_proto_goTypes = []interface{}{
	(*ListJobsRequest)(nil),         // 0: scarab.ListJobsRequest
	(*ListJobsResponse)(nil),        // 1: scarab.ListJobsResponse
	(*ListProfilesRequest)(nil),     // 2: scarab.ListProfilesRequest
	(*ListProfilesResponse)(nil),    // 3: scarab.ListProfilesResponse
	(*WatchActiveJobsRequest)(nil),  // 4: scarab.WatchActiveJobsRequest
	(*WatchActiveJobsResponse)(nil), // 5: scarab.WatchActiveJobsResponse
	(*StopJobRequest)(nil),          // 6: scarab.StopJobRequest
	(*StopJobResponse)(nil),         // 7: scarab.StopJobResponse
	(*Job)(nil),                     // 8: scarab.Job
	(*RegisteredProfile)(nil),       // 9: scarab.RegisteredProfile
	(*StartJobRequest)(nil),         // 10: scarab.StartJobRequest
	(*StartJobResponse)(nil),        // 11: scarab.StartJobResponse
}
var file_scarab_ui_proto_depIdxs = []int32{
	8,  // 0: scarab.ListJobsResponse.jobs:type_name -> scarab.Job
	9,  // 1: scarab.ListProfilesResponse.registered:type_name -> scarab.RegisteredProfile
	10, // 2: scarab.ManagerUI.StartJob:input_type -> scarab.StartJobRequest
	6,  // 3: scarab.ManagerUI.StopJob:input_type -> scarab.StopJobRequest
	2,  // 4: scarab.ManagerUI.ListProfiles:input_type -> scarab.ListProfilesRequest
	0,  // 5: scarab.ManagerUI.ListJobs:input_type -> scarab.ListJobsRequest
	4,  // 6: scarab.ManagerUI.WatchActiveJobs:input_type -> scarab.WatchActiveJobsRequest
	11, // 7: scarab.ManagerUI.StartJob:output_type -> scarab.StartJobResponse
	7,  // 8: scarab.ManagerUI.StopJob:output_type -> scarab.StopJobResponse
	3,  // 9: scarab.ManagerUI.ListProfiles:output_type -> scarab.ListProfilesResponse
	1,  // 10: scarab.ManagerUI.ListJobs:output_type -> scarab.ListJobsResponse
	5,  // 11: scarab.ManagerUI.WatchActiveJobs:output_type -> scarab.WatchActiveJobsResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_scarab_ui_proto_init() }
func file_scarab_ui_proto_init() {
	if File_scarab_ui_proto != nil {
		return
	}
	file_scarab_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_scarab_ui_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobsRequest); i {
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
		file_scarab_ui_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobsResponse); i {
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
		file_scarab_ui_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProfilesRequest); i {
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
		file_scarab_ui_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProfilesResponse); i {
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
		file_scarab_ui_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchActiveJobsRequest); i {
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
		file_scarab_ui_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchActiveJobsResponse); i {
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
		file_scarab_ui_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopJobRequest); i {
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
		file_scarab_ui_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopJobResponse); i {
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
			RawDescriptor: file_scarab_ui_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scarab_ui_proto_goTypes,
		DependencyIndexes: file_scarab_ui_proto_depIdxs,
		MessageInfos:      file_scarab_ui_proto_msgTypes,
	}.Build()
	File_scarab_ui_proto = out.File
	file_scarab_ui_proto_rawDesc = nil
	file_scarab_ui_proto_goTypes = nil
	file_scarab_ui_proto_depIdxs = nil
}
