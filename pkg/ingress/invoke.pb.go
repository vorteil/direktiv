// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: pkg/ingress/invoke.proto

package ingress

import (
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

type InvokeWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Name      *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Input     []byte  `protobuf:"bytes,3,opt,name=input,proto3,oneof" json:"input,omitempty"`
	Wait      *bool   `protobuf:"varint,4,opt,name=wait,proto3,oneof" json:"wait,omitempty"`
}

func (x *InvokeWorkflowRequest) Reset() {
	*x = InvokeWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_invoke_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeWorkflowRequest) ProtoMessage() {}

func (x *InvokeWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_invoke_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeWorkflowRequest.ProtoReflect.Descriptor instead.
func (*InvokeWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_invoke_proto_rawDescGZIP(), []int{0}
}

func (x *InvokeWorkflowRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *InvokeWorkflowRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *InvokeWorkflowRequest) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *InvokeWorkflowRequest) GetWait() bool {
	if x != nil && x.Wait != nil {
		return *x.Wait
	}
	return false
}

type InvokeWorkflowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId *string `protobuf:"bytes,1,opt,name=instanceId,proto3,oneof" json:"instanceId,omitempty"`
	Output     []byte  `protobuf:"bytes,2,opt,name=output,proto3,oneof" json:"output,omitempty"`
	ErrorCode  *string `protobuf:"bytes,3,opt,name=errorCode,proto3,oneof" json:"errorCode,omitempty"`
	ErrorMsg   *string `protobuf:"bytes,4,opt,name=errorMsg,proto3,oneof" json:"errorMsg,omitempty"`
}

func (x *InvokeWorkflowResponse) Reset() {
	*x = InvokeWorkflowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_invoke_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeWorkflowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeWorkflowResponse) ProtoMessage() {}

func (x *InvokeWorkflowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_invoke_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeWorkflowResponse.ProtoReflect.Descriptor instead.
func (*InvokeWorkflowResponse) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_invoke_proto_rawDescGZIP(), []int{1}
}

func (x *InvokeWorkflowResponse) GetInstanceId() string {
	if x != nil && x.InstanceId != nil {
		return *x.InstanceId
	}
	return ""
}

func (x *InvokeWorkflowResponse) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *InvokeWorkflowResponse) GetErrorCode() string {
	if x != nil && x.ErrorCode != nil {
		return *x.ErrorCode
	}
	return ""
}

func (x *InvokeWorkflowResponse) GetErrorMsg() string {
	if x != nil && x.ErrorMsg != nil {
		return *x.ErrorMsg
	}
	return ""
}

var File_pkg_ingress_invoke_proto protoreflect.FileDescriptor

var file_pkg_ingress_invoke_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x69, 0x6e,
	0x76, 0x6f, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x22, 0xb1, 0x01, 0x0a, 0x15, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x02, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x77, 0x61, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x03, 0x52, 0x04, 0x77, 0x61, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x77, 0x61, 0x69, 0x74, 0x22, 0xd3, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x76, 0x6f,
	0x6b, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x72, 0x74,
	0x65, 0x69, 0x6c, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_ingress_invoke_proto_rawDescOnce sync.Once
	file_pkg_ingress_invoke_proto_rawDescData = file_pkg_ingress_invoke_proto_rawDesc
)

func file_pkg_ingress_invoke_proto_rawDescGZIP() []byte {
	file_pkg_ingress_invoke_proto_rawDescOnce.Do(func() {
		file_pkg_ingress_invoke_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_ingress_invoke_proto_rawDescData)
	})
	return file_pkg_ingress_invoke_proto_rawDescData
}

var file_pkg_ingress_invoke_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_ingress_invoke_proto_goTypes = []interface{}{
	(*InvokeWorkflowRequest)(nil),  // 0: ingress.InvokeWorkflowRequest
	(*InvokeWorkflowResponse)(nil), // 1: ingress.InvokeWorkflowResponse
}
var file_pkg_ingress_invoke_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_ingress_invoke_proto_init() }
func file_pkg_ingress_invoke_proto_init() {
	if File_pkg_ingress_invoke_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_ingress_invoke_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeWorkflowRequest); i {
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
		file_pkg_ingress_invoke_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeWorkflowResponse); i {
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
	file_pkg_ingress_invoke_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_ingress_invoke_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_ingress_invoke_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_ingress_invoke_proto_goTypes,
		DependencyIndexes: file_pkg_ingress_invoke_proto_depIdxs,
		MessageInfos:      file_pkg_ingress_invoke_proto_msgTypes,
	}.Build()
	File_pkg_ingress_invoke_proto = out.File
	file_pkg_ingress_invoke_proto_rawDesc = nil
	file_pkg_ingress_invoke_proto_goTypes = nil
	file_pkg_ingress_invoke_proto_depIdxs = nil
}
