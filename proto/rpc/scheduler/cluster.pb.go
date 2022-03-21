// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: proto/rpc/scheduler/cluster.proto

package rpcv3

import (
	v31 "github.com/RafaySystems/rcloud-base/proto/types/commonpb/v3"
	v3 "github.com/RafaySystems/rcloud-base/proto/types/infrapb/v3"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RegisterClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token          string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	SigningRequest []byte `protobuf:"bytes,2,opt,name=signingRequest,proto3" json:"signingRequest,omitempty"`
}

func (x *RegisterClusterRequest) Reset() {
	*x = RegisterClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_scheduler_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterClusterRequest) ProtoMessage() {}

func (x *RegisterClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_scheduler_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterClusterRequest.ProtoReflect.Descriptor instead.
func (*RegisterClusterRequest) Descriptor() ([]byte, []int) {
	return file_proto_rpc_scheduler_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterClusterRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RegisterClusterRequest) GetSigningRequest() []byte {
	if x != nil {
		return x.SigningRequest
	}
	return nil
}

type RegisterClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certificate   []byte `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	CaCertificate []byte `protobuf:"bytes,2,opt,name=caCertificate,proto3" json:"caCertificate,omitempty"`
}

func (x *RegisterClusterResponse) Reset() {
	*x = RegisterClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_scheduler_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterClusterResponse) ProtoMessage() {}

func (x *RegisterClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_scheduler_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterClusterResponse.ProtoReflect.Descriptor instead.
func (*RegisterClusterResponse) Descriptor() ([]byte, []int) {
	return file_proto_rpc_scheduler_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterClusterResponse) GetCertificate() []byte {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *RegisterClusterResponse) GetCaCertificate() []byte {
	if x != nil {
		return x.CaCertificate
	}
	return nil
}

type DeleteClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteClusterResponse) Reset() {
	*x = DeleteClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_scheduler_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClusterResponse) ProtoMessage() {}

func (x *DeleteClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_scheduler_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClusterResponse.ProtoReflect.Descriptor instead.
func (*DeleteClusterResponse) Descriptor() ([]byte, []int) {
	return file_proto_rpc_scheduler_cluster_proto_rawDescGZIP(), []int{2}
}

var File_proto_rpc_scheduler_cluster_proto protoreflect.FileDescriptor

var file_proto_rpc_scheduler_cluster_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x76, 0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2f,
	0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x56, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x26, 0x0a, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x61, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x63, 0x61, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xa8, 0x08, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0xc7, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x21, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x1a, 0x21, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x33, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x70, 0x92, 0x41, 0x36, 0x4a, 0x34, 0x0a, 0x03,
	0x32, 0x30, 0x31, 0x12, 0x2d, 0x0a, 0x2b, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20,
	0x77, 0x68, 0x65, 0x6e, 0x20, 0x65, 0x64, 0x67, 0x65, 0x20, 0x69, 0x73, 0x20, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c,
	0x79, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x22, 0x2c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x8a, 0x01, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x27, 0x2e, 0x72, 0x61, 0x66, 0x61,
	0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x25, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x25, 0x12, 0x23, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x98, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x33,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x21, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x44, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x3e, 0x12, 0x3c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x12, 0x9e, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x21, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76,
	0x33, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x47, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x41, 0x1a, 0x3c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f,
	0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a,
	0x01, 0x2a, 0x12, 0xde, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x76, 0x33, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x27, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x80, 0x01, 0x92, 0x41, 0x39, 0x4a, 0x37, 0x0a, 0x03, 0x32, 0x30, 0x34, 0x12, 0x30, 0x0a,
	0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x69, 0x73, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x2e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x3e, 0x2a, 0x3c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x76, 0x33,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x12, 0xa8, 0x01, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x76, 0x33, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x23, 0x2e, 0x72, 0x61, 0x66,
	0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22,
	0x4d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x47, 0x12, 0x45, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f,
	0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0xa4,
	0x04, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x33, 0x42, 0x0c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x66, 0x61, 0x79, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73,
	0x2f, 0x72, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x3b, 0x72, 0x70, 0x63, 0x76, 0x33, 0xa2, 0x02, 0x03, 0x52, 0x44, 0x52, 0xaa, 0x02, 0x10, 0x52,
	0x61, 0x66, 0x61, 0x79, 0x2e, 0x44, 0x65, 0x76, 0x2e, 0x52, 0x70, 0x63, 0x2e, 0x56, 0x33, 0xca,
	0x02, 0x10, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x52, 0x70, 0x63, 0x5c,
	0x56, 0x33, 0xe2, 0x02, 0x1c, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x52,
	0x70, 0x63, 0x5c, 0x56, 0x33, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x13, 0x52, 0x61, 0x66, 0x61, 0x79, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x3a, 0x3a,
	0x52, 0x70, 0x63, 0x3a, 0x3a, 0x56, 0x33, 0x92, 0x41, 0xda, 0x02, 0x12, 0x23, 0x0a, 0x0f, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x0b,
	0x0a, 0x09, 0x52, 0x61, 0x66, 0x61, 0x79, 0x20, 0x44, 0x65, 0x76, 0x32, 0x03, 0x33, 0x2e, 0x30,
	0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x79, 0x61, 0x6d, 0x6c, 0x52, 0x50, 0x0a, 0x03, 0x34,
	0x30, 0x33, 0x12, 0x49, 0x0a, 0x47, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77,
	0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65,
	0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x3b, 0x0a,
	0x03, 0x34, 0x30, 0x34, 0x12, 0x34, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64,
	0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73,
	0x74, 0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x5a, 0x38, 0x0a, 0x25, 0x0a, 0x0a,
	0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x17, 0x08, 0x02, 0x1a, 0x11,
	0x58, 0x2d, 0x52, 0x41, 0x46, 0x41, 0x59, 0x2d, 0x41, 0x50, 0x49, 0x2d, 0x4b, 0x45, 0x59, 0x49,
	0x44, 0x20, 0x02, 0x0a, 0x0f, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x02, 0x08, 0x01, 0x62, 0x1f, 0x0a, 0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rpc_scheduler_cluster_proto_rawDescOnce sync.Once
	file_proto_rpc_scheduler_cluster_proto_rawDescData = file_proto_rpc_scheduler_cluster_proto_rawDesc
)

func file_proto_rpc_scheduler_cluster_proto_rawDescGZIP() []byte {
	file_proto_rpc_scheduler_cluster_proto_rawDescOnce.Do(func() {
		file_proto_rpc_scheduler_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rpc_scheduler_cluster_proto_rawDescData)
	})
	return file_proto_rpc_scheduler_cluster_proto_rawDescData
}

var file_proto_rpc_scheduler_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_rpc_scheduler_cluster_proto_goTypes = []interface{}{
	(*RegisterClusterRequest)(nil),  // 0: rafay.dev.rpc.v3.RegisterClusterRequest
	(*RegisterClusterResponse)(nil), // 1: rafay.dev.rpc.v3.RegisterClusterResponse
	(*DeleteClusterResponse)(nil),   // 2: rafay.dev.rpc.v3.DeleteClusterResponse
	(*v3.Cluster)(nil),              // 3: rafay.dev.types.infra.v3.Cluster
	(*v31.QueryOptions)(nil),        // 4: rafay.dev.types.common.v3.QueryOptions
	(*v3.ClusterList)(nil),          // 5: rafay.dev.types.infra.v3.ClusterList
	(*v31.HttpBody)(nil),            // 6: rafay.dev.types.common.v3.HttpBody
}
var file_proto_rpc_scheduler_cluster_proto_depIdxs = []int32{
	3, // 0: rafay.dev.rpc.v3.Cluster.CreateCluster:input_type -> rafay.dev.types.infra.v3.Cluster
	4, // 1: rafay.dev.rpc.v3.Cluster.GetClusters:input_type -> rafay.dev.types.common.v3.QueryOptions
	3, // 2: rafay.dev.rpc.v3.Cluster.GetCluster:input_type -> rafay.dev.types.infra.v3.Cluster
	3, // 3: rafay.dev.rpc.v3.Cluster.UpdateCluster:input_type -> rafay.dev.types.infra.v3.Cluster
	3, // 4: rafay.dev.rpc.v3.Cluster.DeleteCluster:input_type -> rafay.dev.types.infra.v3.Cluster
	3, // 5: rafay.dev.rpc.v3.Cluster.DownloadCluster:input_type -> rafay.dev.types.infra.v3.Cluster
	3, // 6: rafay.dev.rpc.v3.Cluster.CreateCluster:output_type -> rafay.dev.types.infra.v3.Cluster
	5, // 7: rafay.dev.rpc.v3.Cluster.GetClusters:output_type -> rafay.dev.types.infra.v3.ClusterList
	3, // 8: rafay.dev.rpc.v3.Cluster.GetCluster:output_type -> rafay.dev.types.infra.v3.Cluster
	3, // 9: rafay.dev.rpc.v3.Cluster.UpdateCluster:output_type -> rafay.dev.types.infra.v3.Cluster
	2, // 10: rafay.dev.rpc.v3.Cluster.DeleteCluster:output_type -> rafay.dev.rpc.v3.DeleteClusterResponse
	6, // 11: rafay.dev.rpc.v3.Cluster.DownloadCluster:output_type -> rafay.dev.types.common.v3.HttpBody
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_rpc_scheduler_cluster_proto_init() }
func file_proto_rpc_scheduler_cluster_proto_init() {
	if File_proto_rpc_scheduler_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rpc_scheduler_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterClusterRequest); i {
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
		file_proto_rpc_scheduler_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterClusterResponse); i {
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
		file_proto_rpc_scheduler_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClusterResponse); i {
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
			RawDescriptor: file_proto_rpc_scheduler_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rpc_scheduler_cluster_proto_goTypes,
		DependencyIndexes: file_proto_rpc_scheduler_cluster_proto_depIdxs,
		MessageInfos:      file_proto_rpc_scheduler_cluster_proto_msgTypes,
	}.Build()
	File_proto_rpc_scheduler_cluster_proto = out.File
	file_proto_rpc_scheduler_cluster_proto_rawDesc = nil
	file_proto_rpc_scheduler_cluster_proto_goTypes = nil
	file_proto_rpc_scheduler_cluster_proto_depIdxs = nil
}
