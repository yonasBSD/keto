// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: ory/keto/opl/v1alpha1/syntax_service.proto

package opl

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

type CheckRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       []byte                 `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	mi := &file_ory_keto_opl_v1alpha1_syntax_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_opl_v1alpha1_syntax_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_ory_keto_opl_v1alpha1_syntax_service_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type CheckResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ParseErrors   []*ParseError          `protobuf:"bytes,1,rep,name=parse_errors,json=parseErrors,proto3" json:"parse_errors,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	mi := &file_ory_keto_opl_v1alpha1_syntax_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_opl_v1alpha1_syntax_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_ory_keto_opl_v1alpha1_syntax_service_proto_rawDescGZIP(), []int{1}
}

func (x *CheckResponse) GetParseErrors() []*ParseError {
	if x != nil {
		return x.ParseErrors
	}
	return nil
}

type ParseError struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Start         *SourcePosition        `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End           *SourcePosition        `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ParseError) Reset() {
	*x = ParseError{}
	mi := &file_ory_keto_opl_v1alpha1_syntax_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParseError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseError) ProtoMessage() {}

func (x *ParseError) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_opl_v1alpha1_syntax_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseError.ProtoReflect.Descriptor instead.
func (*ParseError) Descriptor() ([]byte, []int) {
	return file_ory_keto_opl_v1alpha1_syntax_service_proto_rawDescGZIP(), []int{2}
}

func (x *ParseError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ParseError) GetStart() *SourcePosition {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *ParseError) GetEnd() *SourcePosition {
	if x != nil {
		return x.End
	}
	return nil
}

type SourcePosition struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Line          uint32                 `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	Column        uint32                 `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SourcePosition) Reset() {
	*x = SourcePosition{}
	mi := &file_ory_keto_opl_v1alpha1_syntax_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SourcePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourcePosition) ProtoMessage() {}

func (x *SourcePosition) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_opl_v1alpha1_syntax_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourcePosition.ProtoReflect.Descriptor instead.
func (*SourcePosition) Descriptor() ([]byte, []int) {
	return file_ory_keto_opl_v1alpha1_syntax_service_proto_rawDescGZIP(), []int{3}
}

func (x *SourcePosition) GetLine() uint32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *SourcePosition) GetColumn() uint32 {
	if x != nil {
		return x.Column
	}
	return 0
}

var File_ory_keto_opl_v1alpha1_syntax_service_proto protoreflect.FileDescriptor

var file_ory_keto_opl_v1alpha1_syntax_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x6c, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6f, 0x72,
	0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x6f, 0x70, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x22, 0x28, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x55, 0x0a,
	0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44,
	0x0a, 0x0c, 0x70, 0x61, 0x72, 0x73, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e,
	0x6f, 0x70, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x72,
	0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x73, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x73, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f,
	0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x6f, 0x70, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x37, 0x0a, 0x03, 0x65, 0x6e,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65,
	0x74, 0x6f, 0x2e, 0x6f, 0x70, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x22, 0x3c, 0x0a, 0x0e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x32, 0x63, 0x0a, 0x0d, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x52, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x23, 0x2e, 0x6f, 0x72,
	0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x6f, 0x70, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x6f, 0x70, 0x6c, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x95, 0x01, 0x0a, 0x18, 0x73, 0x68, 0x2e, 0x6f, 0x72,
	0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x6f, 0x70, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x42, 0x12, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x6f, 0x70,
	0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x6f, 0x70, 0x6c, 0xaa, 0x02,
	0x15, 0x4f, 0x72, 0x79, 0x2e, 0x4b, 0x65, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x6c, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x15, 0x4f, 0x72, 0x79, 0x5c, 0x4b, 0x65, 0x74,
	0x6f, 0x5c, 0x4f, 0x70, 0x6c, 0x5c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ory_keto_opl_v1alpha1_syntax_service_proto_rawDescOnce sync.Once
	file_ory_keto_opl_v1alpha1_syntax_service_proto_rawDescData = file_ory_keto_opl_v1alpha1_syntax_service_proto_rawDesc
)

func file_ory_keto_opl_v1alpha1_syntax_service_proto_rawDescGZIP() []byte {
	file_ory_keto_opl_v1alpha1_syntax_service_proto_rawDescOnce.Do(func() {
		file_ory_keto_opl_v1alpha1_syntax_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_ory_keto_opl_v1alpha1_syntax_service_proto_rawDescData)
	})
	return file_ory_keto_opl_v1alpha1_syntax_service_proto_rawDescData
}

var file_ory_keto_opl_v1alpha1_syntax_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ory_keto_opl_v1alpha1_syntax_service_proto_goTypes = []any{
	(*CheckRequest)(nil),   // 0: ory.keto.opl.v1alpha1.CheckRequest
	(*CheckResponse)(nil),  // 1: ory.keto.opl.v1alpha1.CheckResponse
	(*ParseError)(nil),     // 2: ory.keto.opl.v1alpha1.ParseError
	(*SourcePosition)(nil), // 3: ory.keto.opl.v1alpha1.SourcePosition
}
var file_ory_keto_opl_v1alpha1_syntax_service_proto_depIdxs = []int32{
	2, // 0: ory.keto.opl.v1alpha1.CheckResponse.parse_errors:type_name -> ory.keto.opl.v1alpha1.ParseError
	3, // 1: ory.keto.opl.v1alpha1.ParseError.start:type_name -> ory.keto.opl.v1alpha1.SourcePosition
	3, // 2: ory.keto.opl.v1alpha1.ParseError.end:type_name -> ory.keto.opl.v1alpha1.SourcePosition
	0, // 3: ory.keto.opl.v1alpha1.SyntaxService.Check:input_type -> ory.keto.opl.v1alpha1.CheckRequest
	1, // 4: ory.keto.opl.v1alpha1.SyntaxService.Check:output_type -> ory.keto.opl.v1alpha1.CheckResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ory_keto_opl_v1alpha1_syntax_service_proto_init() }
func file_ory_keto_opl_v1alpha1_syntax_service_proto_init() {
	if File_ory_keto_opl_v1alpha1_syntax_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ory_keto_opl_v1alpha1_syntax_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ory_keto_opl_v1alpha1_syntax_service_proto_goTypes,
		DependencyIndexes: file_ory_keto_opl_v1alpha1_syntax_service_proto_depIdxs,
		MessageInfos:      file_ory_keto_opl_v1alpha1_syntax_service_proto_msgTypes,
	}.Build()
	File_ory_keto_opl_v1alpha1_syntax_service_proto = out.File
	file_ory_keto_opl_v1alpha1_syntax_service_proto_rawDesc = nil
	file_ory_keto_opl_v1alpha1_syntax_service_proto_goTypes = nil
	file_ory_keto_opl_v1alpha1_syntax_service_proto_depIdxs = nil
}
