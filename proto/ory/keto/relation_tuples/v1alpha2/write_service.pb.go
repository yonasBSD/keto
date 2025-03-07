// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: ory/keto/relation_tuples/v1alpha2/write_service.proto

package rts

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RelationTupleDelta_Action int32

const (
	// Unspecified.
	// The `TransactRelationTuples` RPC ignores this
	// RelationTupleDelta if an action was unspecified.
	RelationTupleDelta_ACTION_UNSPECIFIED RelationTupleDelta_Action = 0
	// Insertion of a new RelationTuple.
	// It is ignored if already existing.
	RelationTupleDelta_ACTION_INSERT RelationTupleDelta_Action = 1
	// Deletion of the RelationTuple.
	// It is ignored if it does not exist.
	RelationTupleDelta_ACTION_DELETE RelationTupleDelta_Action = 2
)

// Enum value maps for RelationTupleDelta_Action.
var (
	RelationTupleDelta_Action_name = map[int32]string{
		0: "ACTION_UNSPECIFIED",
		1: "ACTION_INSERT",
		2: "ACTION_DELETE",
	}
	RelationTupleDelta_Action_value = map[string]int32{
		"ACTION_UNSPECIFIED": 0,
		"ACTION_INSERT":      1,
		"ACTION_DELETE":      2,
	}
)

func (x RelationTupleDelta_Action) Enum() *RelationTupleDelta_Action {
	p := new(RelationTupleDelta_Action)
	*p = x
	return p
}

func (x RelationTupleDelta_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RelationTupleDelta_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_enumTypes[0].Descriptor()
}

func (RelationTupleDelta_Action) Type() protoreflect.EnumType {
	return &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_enumTypes[0]
}

func (x RelationTupleDelta_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RelationTupleDelta_Action.Descriptor instead.
func (RelationTupleDelta_Action) EnumDescriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP(), []int{1, 0}
}

// The request of a WriteService.TransactRelationTuples RPC.
type TransactRelationTuplesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The write delta for the relationships operated in one single transaction.
	// Either all actions succeed or no change takes effect on error.
	RelationTupleDeltas []*RelationTupleDelta `protobuf:"bytes,1,rep,name=relation_tuple_deltas,json=relationTupleDeltas,proto3" json:"relation_tuple_deltas,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *TransactRelationTuplesRequest) Reset() {
	*x = TransactRelationTuplesRequest{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactRelationTuplesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactRelationTuplesRequest) ProtoMessage() {}

func (x *TransactRelationTuplesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactRelationTuplesRequest.ProtoReflect.Descriptor instead.
func (*TransactRelationTuplesRequest) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP(), []int{0}
}

func (x *TransactRelationTuplesRequest) GetRelationTupleDeltas() []*RelationTupleDelta {
	if x != nil {
		return x.RelationTupleDeltas
	}
	return nil
}

// Write-delta for a TransactRelationTuplesRequest.
type RelationTupleDelta struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The action to do on the RelationTuple.
	Action RelationTupleDelta_Action `protobuf:"varint,1,opt,name=action,proto3,enum=ory.keto.relation_tuples.v1alpha2.RelationTupleDelta_Action" json:"action,omitempty"`
	// The target RelationTuple.
	RelationTuple *RelationTuple `protobuf:"bytes,2,opt,name=relation_tuple,json=relationTuple,proto3" json:"relation_tuple,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RelationTupleDelta) Reset() {
	*x = RelationTupleDelta{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelationTupleDelta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationTupleDelta) ProtoMessage() {}

func (x *RelationTupleDelta) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationTupleDelta.ProtoReflect.Descriptor instead.
func (*RelationTupleDelta) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP(), []int{1}
}

func (x *RelationTupleDelta) GetAction() RelationTupleDelta_Action {
	if x != nil {
		return x.Action
	}
	return RelationTupleDelta_ACTION_UNSPECIFIED
}

func (x *RelationTupleDelta) GetRelationTuple() *RelationTuple {
	if x != nil {
		return x.RelationTuple
	}
	return nil
}

// The response of a WriteService.TransactRelationTuples rpc.
type TransactRelationTuplesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field is not implemented yet and has no effect.
	// <!--
	// The list of the new latest snapshot tokens of the affected RelationTuple,
	// with the same index as specified in the `relation_tuple_deltas` field of
	// the TransactRelationTuplesRequest request.
	//
	// If the RelationTupleDelta_Action was DELETE
	// the snaptoken is empty at the same index.
	// -->
	Snaptokens    []string `protobuf:"bytes,1,rep,name=snaptokens,proto3" json:"snaptokens,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransactRelationTuplesResponse) Reset() {
	*x = TransactRelationTuplesResponse{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactRelationTuplesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactRelationTuplesResponse) ProtoMessage() {}

func (x *TransactRelationTuplesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactRelationTuplesResponse.ProtoReflect.Descriptor instead.
func (*TransactRelationTuplesResponse) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP(), []int{2}
}

func (x *TransactRelationTuplesResponse) GetSnaptokens() []string {
	if x != nil {
		return x.Snaptokens
	}
	return nil
}

type DeleteRelationTuplesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in ory/keto/relation_tuples/v1alpha2/write_service.proto.
	Query         *DeleteRelationTuplesRequest_Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	RelationQuery *RelationQuery                     `protobuf:"bytes,2,opt,name=relation_query,json=relationQuery,proto3" json:"relation_query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRelationTuplesRequest) Reset() {
	*x = DeleteRelationTuplesRequest{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRelationTuplesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRelationTuplesRequest) ProtoMessage() {}

func (x *DeleteRelationTuplesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRelationTuplesRequest.ProtoReflect.Descriptor instead.
func (*DeleteRelationTuplesRequest) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP(), []int{3}
}

// Deprecated: Marked as deprecated in ory/keto/relation_tuples/v1alpha2/write_service.proto.
func (x *DeleteRelationTuplesRequest) GetQuery() *DeleteRelationTuplesRequest_Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *DeleteRelationTuplesRequest) GetRelationQuery() *RelationQuery {
	if x != nil {
		return x.RelationQuery
	}
	return nil
}

type DeleteRelationTuplesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRelationTuplesResponse) Reset() {
	*x = DeleteRelationTuplesResponse{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRelationTuplesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRelationTuplesResponse) ProtoMessage() {}

func (x *DeleteRelationTuplesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRelationTuplesResponse.ProtoReflect.Descriptor instead.
func (*DeleteRelationTuplesResponse) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP(), []int{4}
}

// The query for deleting relationships
type DeleteRelationTuplesRequest_Query struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. The namespace to query.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Optional. The object to query for.
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	// Optional. The relation to query for.
	Relation string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
	// Optional. The subject to query for.
	Subject       *Subject `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRelationTuplesRequest_Query) Reset() {
	*x = DeleteRelationTuplesRequest_Query{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRelationTuplesRequest_Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRelationTuplesRequest_Query) ProtoMessage() {}

func (x *DeleteRelationTuplesRequest_Query) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRelationTuplesRequest_Query.ProtoReflect.Descriptor instead.
func (*DeleteRelationTuplesRequest_Query) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP(), []int{3, 0}
}

func (x *DeleteRelationTuplesRequest_Query) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeleteRelationTuplesRequest_Query) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *DeleteRelationTuplesRequest_Query) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *DeleteRelationTuplesRequest_Query) GetSubject() *Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

var File_ory_keto_relation_tuples_v1alpha2_write_service_proto protoreflect.FileDescriptor

var file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDesc = string([]byte{
	0x0a, 0x35, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74,
	0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x37, 0x6f, 0x72, 0x79, 0x2f,
	0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75,
	0x70, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x1d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x69, 0x0a, 0x15, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x52, 0x13, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x73,
	0x22, 0x8b, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70,
	0x6c, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x54, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65,
	0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x57, 0x0a,
	0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x22, 0x46, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x22, 0x40,
	0x0a, 0x1e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6e, 0x61, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6e, 0x61, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x22, 0xf8, 0x02, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x5e, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x44, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x02, 0x18, 0x01, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x57, 0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b,
	0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70,
	0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x9f, 0x01, 0x0a, 0x05, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74,
	0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x1e, 0x0a, 0x1c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc8, 0x02, 0x0a, 0x0c,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9d, 0x01, 0x0a,
	0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x40, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65,
	0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x6f, 0x72, 0x79, 0x2e,
	0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75,
	0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x97, 0x01, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x75, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x3e, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xc2, 0x01, 0x0a, 0x24, 0x73, 0x68, 0x2e, 0x6f, 0x72,
	0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42,
	0x11, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x3b, 0x72, 0x74, 0x73, 0xaa, 0x02, 0x20, 0x4f, 0x72, 0x79, 0x2e, 0x4b, 0x65, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xca, 0x02, 0x20, 0x4f, 0x72, 0x79, 0x5c, 0x4b,
	0x65, 0x74, 0x6f, 0x5c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c,
	0x65, 0x73, 0x5c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescOnce sync.Once
	file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescData []byte
)

func file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP() []byte {
	file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescOnce.Do(func() {
		file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDesc), len(file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDesc)))
	})
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescData
}

var file_ory_keto_relation_tuples_v1alpha2_write_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ory_keto_relation_tuples_v1alpha2_write_service_proto_goTypes = []any{
	(RelationTupleDelta_Action)(0),            // 0: ory.keto.relation_tuples.v1alpha2.RelationTupleDelta.Action
	(*TransactRelationTuplesRequest)(nil),     // 1: ory.keto.relation_tuples.v1alpha2.TransactRelationTuplesRequest
	(*RelationTupleDelta)(nil),                // 2: ory.keto.relation_tuples.v1alpha2.RelationTupleDelta
	(*TransactRelationTuplesResponse)(nil),    // 3: ory.keto.relation_tuples.v1alpha2.TransactRelationTuplesResponse
	(*DeleteRelationTuplesRequest)(nil),       // 4: ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest
	(*DeleteRelationTuplesResponse)(nil),      // 5: ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesResponse
	(*DeleteRelationTuplesRequest_Query)(nil), // 6: ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest.Query
	(*RelationTuple)(nil),                     // 7: ory.keto.relation_tuples.v1alpha2.RelationTuple
	(*RelationQuery)(nil),                     // 8: ory.keto.relation_tuples.v1alpha2.RelationQuery
	(*Subject)(nil),                           // 9: ory.keto.relation_tuples.v1alpha2.Subject
}
var file_ory_keto_relation_tuples_v1alpha2_write_service_proto_depIdxs = []int32{
	2, // 0: ory.keto.relation_tuples.v1alpha2.TransactRelationTuplesRequest.relation_tuple_deltas:type_name -> ory.keto.relation_tuples.v1alpha2.RelationTupleDelta
	0, // 1: ory.keto.relation_tuples.v1alpha2.RelationTupleDelta.action:type_name -> ory.keto.relation_tuples.v1alpha2.RelationTupleDelta.Action
	7, // 2: ory.keto.relation_tuples.v1alpha2.RelationTupleDelta.relation_tuple:type_name -> ory.keto.relation_tuples.v1alpha2.RelationTuple
	6, // 3: ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest.query:type_name -> ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest.Query
	8, // 4: ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest.relation_query:type_name -> ory.keto.relation_tuples.v1alpha2.RelationQuery
	9, // 5: ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest.Query.subject:type_name -> ory.keto.relation_tuples.v1alpha2.Subject
	1, // 6: ory.keto.relation_tuples.v1alpha2.WriteService.TransactRelationTuples:input_type -> ory.keto.relation_tuples.v1alpha2.TransactRelationTuplesRequest
	4, // 7: ory.keto.relation_tuples.v1alpha2.WriteService.DeleteRelationTuples:input_type -> ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest
	3, // 8: ory.keto.relation_tuples.v1alpha2.WriteService.TransactRelationTuples:output_type -> ory.keto.relation_tuples.v1alpha2.TransactRelationTuplesResponse
	5, // 9: ory.keto.relation_tuples.v1alpha2.WriteService.DeleteRelationTuples:output_type -> ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ory_keto_relation_tuples_v1alpha2_write_service_proto_init() }
func file_ory_keto_relation_tuples_v1alpha2_write_service_proto_init() {
	if File_ory_keto_relation_tuples_v1alpha2_write_service_proto != nil {
		return
	}
	file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDesc), len(file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ory_keto_relation_tuples_v1alpha2_write_service_proto_goTypes,
		DependencyIndexes: file_ory_keto_relation_tuples_v1alpha2_write_service_proto_depIdxs,
		EnumInfos:         file_ory_keto_relation_tuples_v1alpha2_write_service_proto_enumTypes,
		MessageInfos:      file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes,
	}.Build()
	File_ory_keto_relation_tuples_v1alpha2_write_service_proto = out.File
	file_ory_keto_relation_tuples_v1alpha2_write_service_proto_goTypes = nil
	file_ory_keto_relation_tuples_v1alpha2_write_service_proto_depIdxs = nil
}
