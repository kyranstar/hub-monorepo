// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: hubble/v1/rpc.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_hubble_v1_rpc_proto protoreflect.FileDescriptor

var file_hubble_v1_rpc_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x17, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x68, 0x75, 0x62, 0x62, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x68, 0x75,
	0x62, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x75, 0x62,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x68, 0x75,
	0x62, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa2,
	0x15, 0x0a, 0x0a, 0x48, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a,
	0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x12, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x12, 0x1b, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x75, 0x62,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x68,
	0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x75, 0x62, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x30, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x68,
	0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x73, 0x74, 0x49, 0x64, 0x1a,
	0x12, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x74, 0x73, 0x42,
	0x79, 0x46, 0x69, 0x64, 0x12, 0x15, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x75,
	0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x73, 0x74, 0x73, 0x42, 0x79, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x68,
	0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x73, 0x74, 0x73, 0x42, 0x79,
	0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x73, 0x74, 0x73, 0x42, 0x79, 0x4d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x15, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x52, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x79, 0x46, 0x69, 0x64, 0x12, 0x20, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x46,
	0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x75, 0x62, 0x62,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x43, 0x61, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x68,
	0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x23, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x75,
	0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x46, 0x69, 0x64, 0x12, 0x15, 0x2e, 0x68, 0x75,
	0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x59, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68,
	0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e,
	0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x4b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x46, 0x69, 0x64, 0x12, 0x15, 0x2e, 0x68, 0x75, 0x62,
	0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x68, 0x75,
	0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x42, 0x79, 0x46, 0x69, 0x64, 0x12, 0x15, 0x2e, 0x68,
	0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x53, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x75, 0x62, 0x62,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x65, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x46, 0x69, 0x64, 0x73, 0x12, 0x16, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x12, 0x16, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x68, 0x75,
	0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x4a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x79, 0x46, 0x69, 0x64,
	0x12, 0x1c, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e,
	0x6b, 0x73, 0x42, 0x79, 0x46, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x1f, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x42, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x42, 0x79, 0x46, 0x69, 0x64, 0x12, 0x15, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x46, 0x69, 0x64, 0x12, 0x15, 0x2e, 0x68, 0x75,
	0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x55, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x46,
	0x69, 0x64, 0x12, 0x15, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x75, 0x62, 0x62,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79,
	0x46, 0x69, 0x64, 0x12, 0x15, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x75, 0x62,
	0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x42, 0x79, 0x46, 0x69, 0x64, 0x12, 0x15, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x42, 0x79, 0x46, 0x69, 0x64, 0x12, 0x15, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68,
	0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x75, 0x62, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x2e, 0x68,
	0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x68, 0x75, 0x62,
	0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x53, 0x79, 0x6e, 0x63, 0x49, 0x64, 0x73, 0x42, 0x79, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x12, 0x19, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x69, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x1a, 0x12, 0x2e,
	0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x49, 0x64,
	0x73, 0x12, 0x4a, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x42, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x49, 0x64, 0x73, 0x12, 0x12, 0x2e, 0x68,
	0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x49, 0x64, 0x73,
	0x1a, 0x1b, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x19, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x1a, 0x23, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x69, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53,
	0x79, 0x6e, 0x63, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x42, 0x79, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x19, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x69, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x1a, 0x23,
	0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xac, 0x02, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0f, 0x52, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x53,
	0x79, 0x6e, 0x63, 0x54, 0x72, 0x69, 0x65, 0x12, 0x10, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x68, 0x75, 0x62, 0x62,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x17, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x46, 0x72, 0x6f, 0x6d, 0x44, 0x62, 0x12, 0x10, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4f, 0x0a, 0x15, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a,
	0x1a, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x55, 0x0a, 0x17, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x79, 0x72, 0x61, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2f, 0x68, 0x75, 0x62, 0x2d, 0x6d,
	0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68,
	0x75, 0x62, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_hubble_v1_rpc_proto_goTypes = []interface{}{
	(*Message)(nil),                         // 0: hubble.v1.Message
	(*SubscribeRequest)(nil),                // 1: hubble.v1.SubscribeRequest
	(*EventRequest)(nil),                    // 2: hubble.v1.EventRequest
	(*CastId)(nil),                          // 3: hubble.v1.CastId
	(*FidRequest)(nil),                      // 4: hubble.v1.FidRequest
	(*CastsByParentRequest)(nil),            // 5: hubble.v1.CastsByParentRequest
	(*ReactionRequest)(nil),                 // 6: hubble.v1.ReactionRequest
	(*ReactionsByFidRequest)(nil),           // 7: hubble.v1.ReactionsByFidRequest
	(*ReactionsByTargetRequest)(nil),        // 8: hubble.v1.ReactionsByTargetRequest
	(*UserDataRequest)(nil),                 // 9: hubble.v1.UserDataRequest
	(*NameRegistryEventRequest)(nil),        // 10: hubble.v1.NameRegistryEventRequest
	(*VerificationRequest)(nil),             // 11: hubble.v1.VerificationRequest
	(*SignerRequest)(nil),                   // 12: hubble.v1.SignerRequest
	(*IdRegistryEventRequest)(nil),          // 13: hubble.v1.IdRegistryEventRequest
	(*IdRegistryEventByAddressRequest)(nil), // 14: hubble.v1.IdRegistryEventByAddressRequest
	(*FidsRequest)(nil),                     // 15: hubble.v1.FidsRequest
	(*LinkRequest)(nil),                     // 16: hubble.v1.LinkRequest
	(*LinksByFidRequest)(nil),               // 17: hubble.v1.LinksByFidRequest
	(*LinksByTargetRequest)(nil),            // 18: hubble.v1.LinksByTargetRequest
	(*HubInfoRequest)(nil),                  // 19: hubble.v1.HubInfoRequest
	(*SyncStatusRequest)(nil),               // 20: hubble.v1.SyncStatusRequest
	(*TrieNodePrefix)(nil),                  // 21: hubble.v1.TrieNodePrefix
	(*SyncIds)(nil),                         // 22: hubble.v1.SyncIds
	(*Empty)(nil),                           // 23: hubble.v1.Empty
	(*IdRegistryEvent)(nil),                 // 24: hubble.v1.IdRegistryEvent
	(*NameRegistryEvent)(nil),               // 25: hubble.v1.NameRegistryEvent
	(*HubEvent)(nil),                        // 26: hubble.v1.HubEvent
	(*MessagesResponse)(nil),                // 27: hubble.v1.MessagesResponse
	(*FidsResponse)(nil),                    // 28: hubble.v1.FidsResponse
	(*HubInfoResponse)(nil),                 // 29: hubble.v1.HubInfoResponse
	(*SyncStatusResponse)(nil),              // 30: hubble.v1.SyncStatusResponse
	(*TrieNodeMetadataResponse)(nil),        // 31: hubble.v1.TrieNodeMetadataResponse
	(*TrieNodeSnapshotResponse)(nil),        // 32: hubble.v1.TrieNodeSnapshotResponse
}
var file_hubble_v1_rpc_proto_depIdxs = []int32{
	0,  // 0: hubble.v1.HubService.SubmitMessage:input_type -> hubble.v1.Message
	1,  // 1: hubble.v1.HubService.Subscribe:input_type -> hubble.v1.SubscribeRequest
	2,  // 2: hubble.v1.HubService.GetEvent:input_type -> hubble.v1.EventRequest
	3,  // 3: hubble.v1.HubService.GetCast:input_type -> hubble.v1.CastId
	4,  // 4: hubble.v1.HubService.GetCastsByFid:input_type -> hubble.v1.FidRequest
	5,  // 5: hubble.v1.HubService.GetCastsByParent:input_type -> hubble.v1.CastsByParentRequest
	4,  // 6: hubble.v1.HubService.GetCastsByMention:input_type -> hubble.v1.FidRequest
	6,  // 7: hubble.v1.HubService.GetReaction:input_type -> hubble.v1.ReactionRequest
	7,  // 8: hubble.v1.HubService.GetReactionsByFid:input_type -> hubble.v1.ReactionsByFidRequest
	8,  // 9: hubble.v1.HubService.GetReactionsByCast:input_type -> hubble.v1.ReactionsByTargetRequest
	8,  // 10: hubble.v1.HubService.GetReactionsByTarget:input_type -> hubble.v1.ReactionsByTargetRequest
	9,  // 11: hubble.v1.HubService.GetUserData:input_type -> hubble.v1.UserDataRequest
	4,  // 12: hubble.v1.HubService.GetUserDataByFid:input_type -> hubble.v1.FidRequest
	10, // 13: hubble.v1.HubService.GetNameRegistryEvent:input_type -> hubble.v1.NameRegistryEventRequest
	11, // 14: hubble.v1.HubService.GetVerification:input_type -> hubble.v1.VerificationRequest
	4,  // 15: hubble.v1.HubService.GetVerificationsByFid:input_type -> hubble.v1.FidRequest
	12, // 16: hubble.v1.HubService.GetSigner:input_type -> hubble.v1.SignerRequest
	4,  // 17: hubble.v1.HubService.GetSignersByFid:input_type -> hubble.v1.FidRequest
	13, // 18: hubble.v1.HubService.GetIdRegistryEvent:input_type -> hubble.v1.IdRegistryEventRequest
	14, // 19: hubble.v1.HubService.GetIdRegistryEventByAddress:input_type -> hubble.v1.IdRegistryEventByAddressRequest
	15, // 20: hubble.v1.HubService.GetFids:input_type -> hubble.v1.FidsRequest
	16, // 21: hubble.v1.HubService.GetLink:input_type -> hubble.v1.LinkRequest
	17, // 22: hubble.v1.HubService.GetLinksByFid:input_type -> hubble.v1.LinksByFidRequest
	18, // 23: hubble.v1.HubService.GetLinksByTarget:input_type -> hubble.v1.LinksByTargetRequest
	4,  // 24: hubble.v1.HubService.GetAllCastMessagesByFid:input_type -> hubble.v1.FidRequest
	4,  // 25: hubble.v1.HubService.GetAllReactionMessagesByFid:input_type -> hubble.v1.FidRequest
	4,  // 26: hubble.v1.HubService.GetAllVerificationMessagesByFid:input_type -> hubble.v1.FidRequest
	4,  // 27: hubble.v1.HubService.GetAllSignerMessagesByFid:input_type -> hubble.v1.FidRequest
	4,  // 28: hubble.v1.HubService.GetAllUserDataMessagesByFid:input_type -> hubble.v1.FidRequest
	4,  // 29: hubble.v1.HubService.GetAllLinkMessagesByFid:input_type -> hubble.v1.FidRequest
	19, // 30: hubble.v1.HubService.GetInfo:input_type -> hubble.v1.HubInfoRequest
	20, // 31: hubble.v1.HubService.GetSyncStatus:input_type -> hubble.v1.SyncStatusRequest
	21, // 32: hubble.v1.HubService.GetAllSyncIdsByPrefix:input_type -> hubble.v1.TrieNodePrefix
	22, // 33: hubble.v1.HubService.GetAllMessagesBySyncIds:input_type -> hubble.v1.SyncIds
	21, // 34: hubble.v1.HubService.GetSyncMetadataByPrefix:input_type -> hubble.v1.TrieNodePrefix
	21, // 35: hubble.v1.HubService.GetSyncSnapshotByPrefix:input_type -> hubble.v1.TrieNodePrefix
	23, // 36: hubble.v1.AdminService.RebuildSyncTrie:input_type -> hubble.v1.Empty
	23, // 37: hubble.v1.AdminService.DeleteAllMessagesFromDb:input_type -> hubble.v1.Empty
	24, // 38: hubble.v1.AdminService.SubmitIdRegistryEvent:input_type -> hubble.v1.IdRegistryEvent
	25, // 39: hubble.v1.AdminService.SubmitNameRegistryEvent:input_type -> hubble.v1.NameRegistryEvent
	0,  // 40: hubble.v1.HubService.SubmitMessage:output_type -> hubble.v1.Message
	26, // 41: hubble.v1.HubService.Subscribe:output_type -> hubble.v1.HubEvent
	26, // 42: hubble.v1.HubService.GetEvent:output_type -> hubble.v1.HubEvent
	0,  // 43: hubble.v1.HubService.GetCast:output_type -> hubble.v1.Message
	27, // 44: hubble.v1.HubService.GetCastsByFid:output_type -> hubble.v1.MessagesResponse
	27, // 45: hubble.v1.HubService.GetCastsByParent:output_type -> hubble.v1.MessagesResponse
	27, // 46: hubble.v1.HubService.GetCastsByMention:output_type -> hubble.v1.MessagesResponse
	0,  // 47: hubble.v1.HubService.GetReaction:output_type -> hubble.v1.Message
	27, // 48: hubble.v1.HubService.GetReactionsByFid:output_type -> hubble.v1.MessagesResponse
	27, // 49: hubble.v1.HubService.GetReactionsByCast:output_type -> hubble.v1.MessagesResponse
	27, // 50: hubble.v1.HubService.GetReactionsByTarget:output_type -> hubble.v1.MessagesResponse
	0,  // 51: hubble.v1.HubService.GetUserData:output_type -> hubble.v1.Message
	27, // 52: hubble.v1.HubService.GetUserDataByFid:output_type -> hubble.v1.MessagesResponse
	25, // 53: hubble.v1.HubService.GetNameRegistryEvent:output_type -> hubble.v1.NameRegistryEvent
	0,  // 54: hubble.v1.HubService.GetVerification:output_type -> hubble.v1.Message
	27, // 55: hubble.v1.HubService.GetVerificationsByFid:output_type -> hubble.v1.MessagesResponse
	0,  // 56: hubble.v1.HubService.GetSigner:output_type -> hubble.v1.Message
	27, // 57: hubble.v1.HubService.GetSignersByFid:output_type -> hubble.v1.MessagesResponse
	24, // 58: hubble.v1.HubService.GetIdRegistryEvent:output_type -> hubble.v1.IdRegistryEvent
	24, // 59: hubble.v1.HubService.GetIdRegistryEventByAddress:output_type -> hubble.v1.IdRegistryEvent
	28, // 60: hubble.v1.HubService.GetFids:output_type -> hubble.v1.FidsResponse
	0,  // 61: hubble.v1.HubService.GetLink:output_type -> hubble.v1.Message
	27, // 62: hubble.v1.HubService.GetLinksByFid:output_type -> hubble.v1.MessagesResponse
	27, // 63: hubble.v1.HubService.GetLinksByTarget:output_type -> hubble.v1.MessagesResponse
	27, // 64: hubble.v1.HubService.GetAllCastMessagesByFid:output_type -> hubble.v1.MessagesResponse
	27, // 65: hubble.v1.HubService.GetAllReactionMessagesByFid:output_type -> hubble.v1.MessagesResponse
	27, // 66: hubble.v1.HubService.GetAllVerificationMessagesByFid:output_type -> hubble.v1.MessagesResponse
	27, // 67: hubble.v1.HubService.GetAllSignerMessagesByFid:output_type -> hubble.v1.MessagesResponse
	27, // 68: hubble.v1.HubService.GetAllUserDataMessagesByFid:output_type -> hubble.v1.MessagesResponse
	27, // 69: hubble.v1.HubService.GetAllLinkMessagesByFid:output_type -> hubble.v1.MessagesResponse
	29, // 70: hubble.v1.HubService.GetInfo:output_type -> hubble.v1.HubInfoResponse
	30, // 71: hubble.v1.HubService.GetSyncStatus:output_type -> hubble.v1.SyncStatusResponse
	22, // 72: hubble.v1.HubService.GetAllSyncIdsByPrefix:output_type -> hubble.v1.SyncIds
	27, // 73: hubble.v1.HubService.GetAllMessagesBySyncIds:output_type -> hubble.v1.MessagesResponse
	31, // 74: hubble.v1.HubService.GetSyncMetadataByPrefix:output_type -> hubble.v1.TrieNodeMetadataResponse
	32, // 75: hubble.v1.HubService.GetSyncSnapshotByPrefix:output_type -> hubble.v1.TrieNodeSnapshotResponse
	23, // 76: hubble.v1.AdminService.RebuildSyncTrie:output_type -> hubble.v1.Empty
	23, // 77: hubble.v1.AdminService.DeleteAllMessagesFromDb:output_type -> hubble.v1.Empty
	24, // 78: hubble.v1.AdminService.SubmitIdRegistryEvent:output_type -> hubble.v1.IdRegistryEvent
	25, // 79: hubble.v1.AdminService.SubmitNameRegistryEvent:output_type -> hubble.v1.NameRegistryEvent
	40, // [40:80] is the sub-list for method output_type
	0,  // [0:40] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_hubble_v1_rpc_proto_init() }
func file_hubble_v1_rpc_proto_init() {
	if File_hubble_v1_rpc_proto != nil {
		return
	}
	file_hubble_v1_message_proto_init()
	file_hubble_v1_id_registry_event_proto_init()
	file_hubble_v1_name_registry_event_proto_init()
	file_hubble_v1_hub_event_proto_init()
	file_hubble_v1_request_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hubble_v1_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_hubble_v1_rpc_proto_goTypes,
		DependencyIndexes: file_hubble_v1_rpc_proto_depIdxs,
	}.Build()
	File_hubble_v1_rpc_proto = out.File
	file_hubble_v1_rpc_proto_rawDesc = nil
	file_hubble_v1_rpc_proto_goTypes = nil
	file_hubble_v1_rpc_proto_depIdxs = nil
}
