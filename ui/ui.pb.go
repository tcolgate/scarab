// This file is part of scarab.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: ui.proto

package ui

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type StartJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartJobRequest) Reset() {
	*x = StartJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ui_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartJobRequest) ProtoMessage() {}

func (x *StartJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ui_proto_msgTypes[0]
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
	return file_ui_proto_rawDescGZIP(), []int{0}
}

type StartJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartJobResponse) Reset() {
	*x = StartJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ui_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartJobResponse) ProtoMessage() {}

func (x *StartJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ui_proto_msgTypes[1]
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
	return file_ui_proto_rawDescGZIP(), []int{1}
}

type ListProfilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProfilesRequest) Reset() {
	*x = ListProfilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ui_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProfilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProfilesRequest) ProtoMessage() {}

func (x *ListProfilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ui_proto_msgTypes[2]
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
	return file_ui_proto_rawDescGZIP(), []int{2}
}

type ListProfilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProfilesResponse) Reset() {
	*x = ListProfilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ui_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProfilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProfilesResponse) ProtoMessage() {}

func (x *ListProfilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ui_proto_msgTypes[3]
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
	return file_ui_proto_rawDescGZIP(), []int{3}
}

type WatchActiveJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WatchActiveJobsRequest) Reset() {
	*x = WatchActiveJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ui_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchActiveJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchActiveJobsRequest) ProtoMessage() {}

func (x *WatchActiveJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ui_proto_msgTypes[4]
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
	return file_ui_proto_rawDescGZIP(), []int{4}
}

type WatchActiveJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WatchActiveJobsResponse) Reset() {
	*x = WatchActiveJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ui_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchActiveJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchActiveJobsResponse) ProtoMessage() {}

func (x *WatchActiveJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ui_proto_msgTypes[5]
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
	return file_ui_proto_rawDescGZIP(), []int{5}
}

var File_ui_proto protoreflect.FileDescriptor

var file_ui_proto_rawDesc = []byte{
	0x0a, 0x08, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x75, 0x69, 0x22, 0x11,
	0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x57, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x19,
	0x0a, 0x17, 0x57, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4a, 0x6f, 0x62,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd9, 0x01, 0x0a, 0x09, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x55, 0x49, 0x12, 0x37, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x4a, 0x6f, 0x62, 0x12, 0x13, 0x2e, 0x75, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x75, 0x69, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x12, 0x17, 0x2e, 0x75, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75, 0x69, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x1a, 0x2e, 0x75, 0x69, 0x2e, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x69, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6f, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x63, 0x61,
	0x72, 0x61, 0x62, 0x2f, 0x75, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ui_proto_rawDescOnce sync.Once
	file_ui_proto_rawDescData = file_ui_proto_rawDesc
)

func file_ui_proto_rawDescGZIP() []byte {
	file_ui_proto_rawDescOnce.Do(func() {
		file_ui_proto_rawDescData = protoimpl.X.CompressGZIP(file_ui_proto_rawDescData)
	})
	return file_ui_proto_rawDescData
}

var file_ui_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ui_proto_goTypes = []interface{}{
	(*StartJobRequest)(nil),         // 0: ui.StartJobRequest
	(*StartJobResponse)(nil),        // 1: ui.StartJobResponse
	(*ListProfilesRequest)(nil),     // 2: ui.ListProfilesRequest
	(*ListProfilesResponse)(nil),    // 3: ui.ListProfilesResponse
	(*WatchActiveJobsRequest)(nil),  // 4: ui.WatchActiveJobsRequest
	(*WatchActiveJobsResponse)(nil), // 5: ui.WatchActiveJobsResponse
}
var file_ui_proto_depIdxs = []int32{
	0, // 0: ui.ManagerUI.StartJob:input_type -> ui.StartJobRequest
	2, // 1: ui.ManagerUI.ListProfiles:input_type -> ui.ListProfilesRequest
	4, // 2: ui.ManagerUI.WatchActiveJobs:input_type -> ui.WatchActiveJobsRequest
	1, // 3: ui.ManagerUI.StartJob:output_type -> ui.StartJobResponse
	3, // 4: ui.ManagerUI.ListProfiles:output_type -> ui.ListProfilesResponse
	5, // 5: ui.ManagerUI.WatchActiveJobs:output_type -> ui.WatchActiveJobsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ui_proto_init() }
func file_ui_proto_init() {
	if File_ui_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ui_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_ui_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_ui_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_ui_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_ui_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_ui_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ui_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ui_proto_goTypes,
		DependencyIndexes: file_ui_proto_depIdxs,
		MessageInfos:      file_ui_proto_msgTypes,
	}.Build()
	File_ui_proto = out.File
	file_ui_proto_rawDesc = nil
	file_ui_proto_goTypes = nil
	file_ui_proto_depIdxs = nil
}
