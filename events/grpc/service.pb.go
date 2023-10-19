// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.10
// source: api/events/service.proto

package grpc

import (
	context "context"
	pb "github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	pb1 "github.com/GoogleCloudPlatform/retail-data-model/events/pb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_events_service_proto protoreflect.FileDescriptor

var file_api_events_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x6d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x63, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12,
	0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0xd1, 0x05, 0x0a, 0x11, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x62, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x30, 0x01, 0x12, 0x7d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x22, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x27, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x97, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x27, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x11, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x97, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x28, 0x01, 0x30, 0x01, 0x12, 0x97, 0x01, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f,
	0x3a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2a, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x28,
	0x01, 0x30, 0x01, 0x32, 0x8e, 0x02, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x78, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x1d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x30, 0x01, 0x12, 0x83,
	0x01, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x12, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x62,
	0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x32, 0x7c, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x72,
	0x0a, 0x04, 0x45, 0x6d, 0x69, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x62,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x0e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x5f, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2d, 0x64, 0x61, 0x74,
	0x61, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_events_service_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil),          // 0: google.protobuf.Empty
	(*pb.IDRequest)(nil),           // 1: google.retail.common.pb.IDRequest
	(*pb1.EventDescription)(nil),   // 2: google.retail.events.pb.EventDescription
	(*pb.TimeBoundRequest)(nil),    // 3: google.retail.common.pb.TimeBoundRequest
	(*pb1.Event)(nil),              // 4: google.retail.events.pb.Event
	(*pb.HealthCheckResponse)(nil), // 5: google.retail.common.pb.HealthCheckResponse
	(*pb.StatusResponse)(nil),      // 6: google.retail.common.pb.StatusResponse
	(*pb1.EventRecord)(nil),        // 7: google.retail.events.pb.EventRecord
}
var file_api_events_service_proto_depIdxs = []int32{
	0, // 0: google.retail.events.grpc.Status.Get:input_type -> google.protobuf.Empty
	0, // 1: google.retail.events.grpc.EventDescriptions.List:input_type -> google.protobuf.Empty
	1, // 2: google.retail.events.grpc.EventDescriptions.Get:input_type -> google.retail.common.pb.IDRequest
	2, // 3: google.retail.events.grpc.EventDescriptions.Create:input_type -> google.retail.events.pb.EventDescription
	2, // 4: google.retail.events.grpc.EventDescriptions.Update:input_type -> google.retail.events.pb.EventDescription
	2, // 5: google.retail.events.grpc.EventDescriptions.Delete:input_type -> google.retail.events.pb.EventDescription
	3, // 6: google.retail.events.grpc.EventRecords.List:input_type -> google.retail.common.pb.TimeBoundRequest
	1, // 7: google.retail.events.grpc.EventRecords.FindTransactionById:input_type -> google.retail.common.pb.IDRequest
	4, // 8: google.retail.events.grpc.Events.Emit:input_type -> google.retail.events.pb.Event
	5, // 9: google.retail.events.grpc.Status.Get:output_type -> google.retail.common.pb.HealthCheckResponse
	2, // 10: google.retail.events.grpc.EventDescriptions.List:output_type -> google.retail.events.pb.EventDescription
	2, // 11: google.retail.events.grpc.EventDescriptions.Get:output_type -> google.retail.events.pb.EventDescription
	6, // 12: google.retail.events.grpc.EventDescriptions.Create:output_type -> google.retail.common.pb.StatusResponse
	6, // 13: google.retail.events.grpc.EventDescriptions.Update:output_type -> google.retail.common.pb.StatusResponse
	6, // 14: google.retail.events.grpc.EventDescriptions.Delete:output_type -> google.retail.common.pb.StatusResponse
	7, // 15: google.retail.events.grpc.EventRecords.List:output_type -> google.retail.events.pb.EventRecord
	7, // 16: google.retail.events.grpc.EventRecords.FindTransactionById:output_type -> google.retail.events.pb.EventRecord
	6, // 17: google.retail.events.grpc.Events.Emit:output_type -> google.retail.common.pb.StatusResponse
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_events_service_proto_init() }
func file_api_events_service_proto_init() {
	if File_api_events_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_events_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_api_events_service_proto_goTypes,
		DependencyIndexes: file_api_events_service_proto_depIdxs,
	}.Build()
	File_api_events_service_proto = out.File
	file_api_events_service_proto_rawDesc = nil
	file_api_events_service_proto_goTypes = nil
	file_api_events_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StatusClient is the client API for Status service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatusClient interface {
	Get(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pb.HealthCheckResponse, error)
}

type statusClient struct {
	cc grpc.ClientConnInterface
}

func NewStatusClient(cc grpc.ClientConnInterface) StatusClient {
	return &statusClient{cc}
}

func (c *statusClient) Get(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pb.HealthCheckResponse, error) {
	out := new(pb.HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/google.retail.events.grpc.Status/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatusServer is the server API for Status service.
type StatusServer interface {
	Get(context.Context, *emptypb.Empty) (*pb.HealthCheckResponse, error)
}

// UnimplementedStatusServer can be embedded to have forward compatible implementations.
type UnimplementedStatusServer struct {
}

func (*UnimplementedStatusServer) Get(context.Context, *emptypb.Empty) (*pb.HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterStatusServer(s *grpc.Server, srv StatusServer) {
	s.RegisterService(&_Status_serviceDesc, srv)
}

func _Status_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.retail.events.grpc.Status/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServer).Get(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Status_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.retail.events.grpc.Status",
	HandlerType: (*StatusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Status_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/events/service.proto",
}

// EventDescriptionsClient is the client API for EventDescriptions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventDescriptionsClient interface {
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (EventDescriptions_ListClient, error)
	Get(ctx context.Context, in *pb.IDRequest, opts ...grpc.CallOption) (*pb1.EventDescription, error)
	Create(ctx context.Context, opts ...grpc.CallOption) (EventDescriptions_CreateClient, error)
	Update(ctx context.Context, opts ...grpc.CallOption) (EventDescriptions_UpdateClient, error)
	Delete(ctx context.Context, opts ...grpc.CallOption) (EventDescriptions_DeleteClient, error)
}

type eventDescriptionsClient struct {
	cc grpc.ClientConnInterface
}

func NewEventDescriptionsClient(cc grpc.ClientConnInterface) EventDescriptionsClient {
	return &eventDescriptionsClient{cc}
}

func (c *eventDescriptionsClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (EventDescriptions_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventDescriptions_serviceDesc.Streams[0], "/google.retail.events.grpc.EventDescriptions/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventDescriptionsListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventDescriptions_ListClient interface {
	Recv() (*pb1.EventDescription, error)
	grpc.ClientStream
}

type eventDescriptionsListClient struct {
	grpc.ClientStream
}

func (x *eventDescriptionsListClient) Recv() (*pb1.EventDescription, error) {
	m := new(pb1.EventDescription)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventDescriptionsClient) Get(ctx context.Context, in *pb.IDRequest, opts ...grpc.CallOption) (*pb1.EventDescription, error) {
	out := new(pb1.EventDescription)
	err := c.cc.Invoke(ctx, "/google.retail.events.grpc.EventDescriptions/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventDescriptionsClient) Create(ctx context.Context, opts ...grpc.CallOption) (EventDescriptions_CreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventDescriptions_serviceDesc.Streams[1], "/google.retail.events.grpc.EventDescriptions/Create", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventDescriptionsCreateClient{stream}
	return x, nil
}

type EventDescriptions_CreateClient interface {
	Send(*pb1.EventDescription) error
	Recv() (*pb.StatusResponse, error)
	grpc.ClientStream
}

type eventDescriptionsCreateClient struct {
	grpc.ClientStream
}

func (x *eventDescriptionsCreateClient) Send(m *pb1.EventDescription) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventDescriptionsCreateClient) Recv() (*pb.StatusResponse, error) {
	m := new(pb.StatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventDescriptionsClient) Update(ctx context.Context, opts ...grpc.CallOption) (EventDescriptions_UpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventDescriptions_serviceDesc.Streams[2], "/google.retail.events.grpc.EventDescriptions/Update", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventDescriptionsUpdateClient{stream}
	return x, nil
}

type EventDescriptions_UpdateClient interface {
	Send(*pb1.EventDescription) error
	Recv() (*pb.StatusResponse, error)
	grpc.ClientStream
}

type eventDescriptionsUpdateClient struct {
	grpc.ClientStream
}

func (x *eventDescriptionsUpdateClient) Send(m *pb1.EventDescription) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventDescriptionsUpdateClient) Recv() (*pb.StatusResponse, error) {
	m := new(pb.StatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventDescriptionsClient) Delete(ctx context.Context, opts ...grpc.CallOption) (EventDescriptions_DeleteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventDescriptions_serviceDesc.Streams[3], "/google.retail.events.grpc.EventDescriptions/Delete", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventDescriptionsDeleteClient{stream}
	return x, nil
}

type EventDescriptions_DeleteClient interface {
	Send(*pb1.EventDescription) error
	Recv() (*pb.StatusResponse, error)
	grpc.ClientStream
}

type eventDescriptionsDeleteClient struct {
	grpc.ClientStream
}

func (x *eventDescriptionsDeleteClient) Send(m *pb1.EventDescription) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventDescriptionsDeleteClient) Recv() (*pb.StatusResponse, error) {
	m := new(pb.StatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventDescriptionsServer is the server API for EventDescriptions service.
type EventDescriptionsServer interface {
	List(*emptypb.Empty, EventDescriptions_ListServer) error
	Get(context.Context, *pb.IDRequest) (*pb1.EventDescription, error)
	Create(EventDescriptions_CreateServer) error
	Update(EventDescriptions_UpdateServer) error
	Delete(EventDescriptions_DeleteServer) error
}

// UnimplementedEventDescriptionsServer can be embedded to have forward compatible implementations.
type UnimplementedEventDescriptionsServer struct {
}

func (*UnimplementedEventDescriptionsServer) List(*emptypb.Empty, EventDescriptions_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedEventDescriptionsServer) Get(context.Context, *pb.IDRequest) (*pb1.EventDescription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedEventDescriptionsServer) Create(EventDescriptions_CreateServer) error {
	return status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedEventDescriptionsServer) Update(EventDescriptions_UpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedEventDescriptionsServer) Delete(EventDescriptions_DeleteServer) error {
	return status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterEventDescriptionsServer(s *grpc.Server, srv EventDescriptionsServer) {
	s.RegisterService(&_EventDescriptions_serviceDesc, srv)
}

func _EventDescriptions_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventDescriptionsServer).List(m, &eventDescriptionsListServer{stream})
}

type EventDescriptions_ListServer interface {
	Send(*pb1.EventDescription) error
	grpc.ServerStream
}

type eventDescriptionsListServer struct {
	grpc.ServerStream
}

func (x *eventDescriptionsListServer) Send(m *pb1.EventDescription) error {
	return x.ServerStream.SendMsg(m)
}

func _EventDescriptions_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventDescriptionsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.retail.events.grpc.EventDescriptions/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventDescriptionsServer).Get(ctx, req.(*pb.IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventDescriptions_Create_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventDescriptionsServer).Create(&eventDescriptionsCreateServer{stream})
}

type EventDescriptions_CreateServer interface {
	Send(*pb.StatusResponse) error
	Recv() (*pb1.EventDescription, error)
	grpc.ServerStream
}

type eventDescriptionsCreateServer struct {
	grpc.ServerStream
}

func (x *eventDescriptionsCreateServer) Send(m *pb.StatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventDescriptionsCreateServer) Recv() (*pb1.EventDescription, error) {
	m := new(pb1.EventDescription)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EventDescriptions_Update_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventDescriptionsServer).Update(&eventDescriptionsUpdateServer{stream})
}

type EventDescriptions_UpdateServer interface {
	Send(*pb.StatusResponse) error
	Recv() (*pb1.EventDescription, error)
	grpc.ServerStream
}

type eventDescriptionsUpdateServer struct {
	grpc.ServerStream
}

func (x *eventDescriptionsUpdateServer) Send(m *pb.StatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventDescriptionsUpdateServer) Recv() (*pb1.EventDescription, error) {
	m := new(pb1.EventDescription)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EventDescriptions_Delete_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventDescriptionsServer).Delete(&eventDescriptionsDeleteServer{stream})
}

type EventDescriptions_DeleteServer interface {
	Send(*pb.StatusResponse) error
	Recv() (*pb1.EventDescription, error)
	grpc.ServerStream
}

type eventDescriptionsDeleteServer struct {
	grpc.ServerStream
}

func (x *eventDescriptionsDeleteServer) Send(m *pb.StatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventDescriptionsDeleteServer) Recv() (*pb1.EventDescription, error) {
	m := new(pb1.EventDescription)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EventDescriptions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.retail.events.grpc.EventDescriptions",
	HandlerType: (*EventDescriptionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _EventDescriptions_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _EventDescriptions_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Create",
			Handler:       _EventDescriptions_Create_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Update",
			Handler:       _EventDescriptions_Update_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Delete",
			Handler:       _EventDescriptions_Delete_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/events/service.proto",
}

// EventRecordsClient is the client API for EventRecords service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventRecordsClient interface {
	List(ctx context.Context, in *pb.TimeBoundRequest, opts ...grpc.CallOption) (EventRecords_ListClient, error)
	FindTransactionById(ctx context.Context, in *pb.IDRequest, opts ...grpc.CallOption) (*pb1.EventRecord, error)
}

type eventRecordsClient struct {
	cc grpc.ClientConnInterface
}

func NewEventRecordsClient(cc grpc.ClientConnInterface) EventRecordsClient {
	return &eventRecordsClient{cc}
}

func (c *eventRecordsClient) List(ctx context.Context, in *pb.TimeBoundRequest, opts ...grpc.CallOption) (EventRecords_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventRecords_serviceDesc.Streams[0], "/google.retail.events.grpc.EventRecords/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventRecordsListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventRecords_ListClient interface {
	Recv() (*pb1.EventRecord, error)
	grpc.ClientStream
}

type eventRecordsListClient struct {
	grpc.ClientStream
}

func (x *eventRecordsListClient) Recv() (*pb1.EventRecord, error) {
	m := new(pb1.EventRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventRecordsClient) FindTransactionById(ctx context.Context, in *pb.IDRequest, opts ...grpc.CallOption) (*pb1.EventRecord, error) {
	out := new(pb1.EventRecord)
	err := c.cc.Invoke(ctx, "/google.retail.events.grpc.EventRecords/FindTransactionById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventRecordsServer is the server API for EventRecords service.
type EventRecordsServer interface {
	List(*pb.TimeBoundRequest, EventRecords_ListServer) error
	FindTransactionById(context.Context, *pb.IDRequest) (*pb1.EventRecord, error)
}

// UnimplementedEventRecordsServer can be embedded to have forward compatible implementations.
type UnimplementedEventRecordsServer struct {
}

func (*UnimplementedEventRecordsServer) List(*pb.TimeBoundRequest, EventRecords_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedEventRecordsServer) FindTransactionById(context.Context, *pb.IDRequest) (*pb1.EventRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTransactionById not implemented")
}

func RegisterEventRecordsServer(s *grpc.Server, srv EventRecordsServer) {
	s.RegisterService(&_EventRecords_serviceDesc, srv)
}

func _EventRecords_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(pb.TimeBoundRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventRecordsServer).List(m, &eventRecordsListServer{stream})
}

type EventRecords_ListServer interface {
	Send(*pb1.EventRecord) error
	grpc.ServerStream
}

type eventRecordsListServer struct {
	grpc.ServerStream
}

func (x *eventRecordsListServer) Send(m *pb1.EventRecord) error {
	return x.ServerStream.SendMsg(m)
}

func _EventRecords_FindTransactionById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventRecordsServer).FindTransactionById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.retail.events.grpc.EventRecords/FindTransactionById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventRecordsServer).FindTransactionById(ctx, req.(*pb.IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventRecords_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.retail.events.grpc.EventRecords",
	HandlerType: (*EventRecordsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindTransactionById",
			Handler:    _EventRecords_FindTransactionById_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _EventRecords_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/events/service.proto",
}

// EventsClient is the client API for Events service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventsClient interface {
	Emit(ctx context.Context, opts ...grpc.CallOption) (Events_EmitClient, error)
}

type eventsClient struct {
	cc grpc.ClientConnInterface
}

func NewEventsClient(cc grpc.ClientConnInterface) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) Emit(ctx context.Context, opts ...grpc.CallOption) (Events_EmitClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Events_serviceDesc.Streams[0], "/google.retail.events.grpc.Events/Emit", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventsEmitClient{stream}
	return x, nil
}

type Events_EmitClient interface {
	Send(*pb1.Event) error
	Recv() (*pb.StatusResponse, error)
	grpc.ClientStream
}

type eventsEmitClient struct {
	grpc.ClientStream
}

func (x *eventsEmitClient) Send(m *pb1.Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventsEmitClient) Recv() (*pb.StatusResponse, error) {
	m := new(pb.StatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventsServer is the server API for Events service.
type EventsServer interface {
	Emit(Events_EmitServer) error
}

// UnimplementedEventsServer can be embedded to have forward compatible implementations.
type UnimplementedEventsServer struct {
}

func (*UnimplementedEventsServer) Emit(Events_EmitServer) error {
	return status.Errorf(codes.Unimplemented, "method Emit not implemented")
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_Emit_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventsServer).Emit(&eventsEmitServer{stream})
}

type Events_EmitServer interface {
	Send(*pb.StatusResponse) error
	Recv() (*pb1.Event, error)
	grpc.ServerStream
}

type eventsEmitServer struct {
	grpc.ServerStream
}

func (x *eventsEmitServer) Send(m *pb.StatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventsEmitServer) Recv() (*pb1.Event, error) {
	m := new(pb1.Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.retail.events.grpc.Events",
	HandlerType: (*EventsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Emit",
			Handler:       _Events_Emit_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/events/service.proto",
}
