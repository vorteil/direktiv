// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: pkg/ingress/update-workflow.proto

package ingress

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

type UpdateWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         *string `protobuf:"bytes,1,opt,name=uid,proto3,oneof" json:"uid,omitempty"`
	Revision    *int32  `protobuf:"varint,2,opt,name=revision,proto3,oneof" json:"revision,omitempty"`
	Active      *bool   `protobuf:"varint,3,opt,name=active,proto3,oneof" json:"active,omitempty"`
	Workflow    []byte  `protobuf:"bytes,4,opt,name=workflow,proto3,oneof" json:"workflow,omitempty"`
	LogToEvents *string `protobuf:"bytes,5,opt,name=logToEvents,proto3,oneof" json:"logToEvents,omitempty"`
}

func (x *UpdateWorkflowRequest) Reset() {
	*x = UpdateWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_update_workflow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkflowRequest) ProtoMessage() {}

func (x *UpdateWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_update_workflow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkflowRequest.ProtoReflect.Descriptor instead.
func (*UpdateWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_update_workflow_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateWorkflowRequest) GetUid() string {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return ""
}

func (x *UpdateWorkflowRequest) GetRevision() int32 {
	if x != nil && x.Revision != nil {
		return *x.Revision
	}
	return 0
}

func (x *UpdateWorkflowRequest) GetActive() bool {
	if x != nil && x.Active != nil {
		return *x.Active
	}
	return false
}

func (x *UpdateWorkflowRequest) GetWorkflow() []byte {
	if x != nil {
		return x.Workflow
	}
	return nil
}

func (x *UpdateWorkflowRequest) GetLogToEvents() string {
	if x != nil && x.LogToEvents != nil {
		return *x.LogToEvents
	}
	return ""
}

type UpdateWorkflowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       *string              `protobuf:"bytes,1,opt,name=uid,proto3,oneof" json:"uid,omitempty"`
	Id        *string              `protobuf:"bytes,2,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Revision  *int32               `protobuf:"varint,3,opt,name=revision,proto3,oneof" json:"revision,omitempty"`
	Active    *bool                `protobuf:"varint,4,opt,name=active,proto3,oneof" json:"active,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3,oneof" json:"createdAt,omitempty"`
}

func (x *UpdateWorkflowResponse) Reset() {
	*x = UpdateWorkflowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_update_workflow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWorkflowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkflowResponse) ProtoMessage() {}

func (x *UpdateWorkflowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_update_workflow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkflowResponse.ProtoReflect.Descriptor instead.
func (*UpdateWorkflowResponse) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_update_workflow_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateWorkflowResponse) GetUid() string {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return ""
}

func (x *UpdateWorkflowResponse) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *UpdateWorkflowResponse) GetRevision() int32 {
	if x != nil && x.Revision != nil {
		return *x.Revision
	}
	return 0
}

func (x *UpdateWorkflowResponse) GetActive() bool {
	if x != nil && x.Active != nil {
		return *x.Active
	}
	return false
}

func (x *UpdateWorkflowResponse) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_pkg_ingress_update_workflow_proto protoreflect.FileDescriptor

var file_pkg_ingress_update_workflow_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f,
	0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x01, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x1b, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x02, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x03,
	0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a,
	0x0b, 0x6c, 0x6f, 0x67, 0x54, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x04, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x54, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x6f, 0x67, 0x54, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0xf6, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x04, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x69, 0x64, 0x42, 0x05, 0x0a,
	0x03, 0x5f, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6c,
	0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_ingress_update_workflow_proto_rawDescOnce sync.Once
	file_pkg_ingress_update_workflow_proto_rawDescData = file_pkg_ingress_update_workflow_proto_rawDesc
)

func file_pkg_ingress_update_workflow_proto_rawDescGZIP() []byte {
	file_pkg_ingress_update_workflow_proto_rawDescOnce.Do(func() {
		file_pkg_ingress_update_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_ingress_update_workflow_proto_rawDescData)
	})
	return file_pkg_ingress_update_workflow_proto_rawDescData
}

var file_pkg_ingress_update_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_ingress_update_workflow_proto_goTypes = []interface{}{
	(*UpdateWorkflowRequest)(nil),  // 0: ingress.UpdateWorkflowRequest
	(*UpdateWorkflowResponse)(nil), // 1: ingress.UpdateWorkflowResponse
	(*timestamp.Timestamp)(nil),    // 2: google.protobuf.Timestamp
}
var file_pkg_ingress_update_workflow_proto_depIdxs = []int32{
	2, // 0: ingress.UpdateWorkflowResponse.createdAt:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_ingress_update_workflow_proto_init() }
func file_pkg_ingress_update_workflow_proto_init() {
	if File_pkg_ingress_update_workflow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_ingress_update_workflow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWorkflowRequest); i {
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
		file_pkg_ingress_update_workflow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWorkflowResponse); i {
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
	file_pkg_ingress_update_workflow_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_ingress_update_workflow_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_ingress_update_workflow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_ingress_update_workflow_proto_goTypes,
		DependencyIndexes: file_pkg_ingress_update_workflow_proto_depIdxs,
		MessageInfos:      file_pkg_ingress_update_workflow_proto_msgTypes,
	}.Build()
	File_pkg_ingress_update_workflow_proto = out.File
	file_pkg_ingress_update_workflow_proto_rawDesc = nil
	file_pkg_ingress_update_workflow_proto_goTypes = nil
	file_pkg_ingress_update_workflow_proto_depIdxs = nil
}
