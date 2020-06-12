//
//Copyright (c) 2020 Docker Inc.
//
//Permission is hereby granted, free of charge, to any person
//obtaining a copy of this software and associated documentation
//files (the "Software"), to deal in the Software without
//restriction, including without limitation the rights to use, copy,
//modify, merge, publish, distribute, sublicense, and/or sell copies
//of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be
//included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
//EXPRESS OR IMPLIED,
//INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
//IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
//HOLDERS BE LIABLE FOR ANY CLAIM,
//DAMAGES OR OTHER LIABILITY,
//WHETHER IN AN ACTION OF CONTRACT,
//TORT OR OTHERWISE,
//ARISING FROM, OUT OF OR IN CONNECTION WITH
//THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.2
// source: protos/compose/v1/compose.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ComposeUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string   `protobuf:"bytes,1,opt,name=projectName,proto3" json:"projectName,omitempty"`
	WorkDir     string   `protobuf:"bytes,2,opt,name=workDir,proto3" json:"workDir,omitempty"`
	Files       []string `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *ComposeUpRequest) Reset() {
	*x = ComposeUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_compose_v1_compose_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeUpRequest) ProtoMessage() {}

func (x *ComposeUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_compose_v1_compose_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeUpRequest.ProtoReflect.Descriptor instead.
func (*ComposeUpRequest) Descriptor() ([]byte, []int) {
	return file_protos_compose_v1_compose_proto_rawDescGZIP(), []int{0}
}

func (x *ComposeUpRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *ComposeUpRequest) GetWorkDir() string {
	if x != nil {
		return x.WorkDir
	}
	return ""
}

func (x *ComposeUpRequest) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

type ComposeUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ComposeUpResponse) Reset() {
	*x = ComposeUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_compose_v1_compose_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeUpResponse) ProtoMessage() {}

func (x *ComposeUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_compose_v1_compose_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeUpResponse.ProtoReflect.Descriptor instead.
func (*ComposeUpResponse) Descriptor() ([]byte, []int) {
	return file_protos_compose_v1_compose_proto_rawDescGZIP(), []int{1}
}

type ComposeDownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ComposeDownRequest) Reset() {
	*x = ComposeDownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_compose_v1_compose_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeDownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeDownRequest) ProtoMessage() {}

func (x *ComposeDownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_compose_v1_compose_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeDownRequest.ProtoReflect.Descriptor instead.
func (*ComposeDownRequest) Descriptor() ([]byte, []int) {
	return file_protos_compose_v1_compose_proto_rawDescGZIP(), []int{2}
}

type ComposeDownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ComposeDownResponse) Reset() {
	*x = ComposeDownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_compose_v1_compose_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeDownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeDownResponse) ProtoMessage() {}

func (x *ComposeDownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_compose_v1_compose_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeDownResponse.ProtoReflect.Descriptor instead.
func (*ComposeDownResponse) Descriptor() ([]byte, []int) {
	return file_protos_compose_v1_compose_proto_rawDescGZIP(), []int{3}
}

var File_protos_compose_v1_compose_proto protoreflect.FileDescriptor

var file_protos_compose_v1_compose_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x22, 0x64, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x55, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72,
	0x6b, 0x44, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b,
	0x44, 0x69, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x65, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14,
	0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x44,
	0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xed, 0x01, 0x0a, 0x07,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x02, 0x55, 0x70, 0x12, 0x32, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x55, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x04, 0x44, 0x6f, 0x77, 0x6e, 0x12, 0x34,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x44,
	0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_compose_v1_compose_proto_rawDescOnce sync.Once
	file_protos_compose_v1_compose_proto_rawDescData = file_protos_compose_v1_compose_proto_rawDesc
)

func file_protos_compose_v1_compose_proto_rawDescGZIP() []byte {
	file_protos_compose_v1_compose_proto_rawDescOnce.Do(func() {
		file_protos_compose_v1_compose_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_compose_v1_compose_proto_rawDescData)
	})
	return file_protos_compose_v1_compose_proto_rawDescData
}

var file_protos_compose_v1_compose_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_compose_v1_compose_proto_goTypes = []interface{}{
	(*ComposeUpRequest)(nil),    // 0: com.docker.api.protos.compose.v1.ComposeUpRequest
	(*ComposeUpResponse)(nil),   // 1: com.docker.api.protos.compose.v1.ComposeUpResponse
	(*ComposeDownRequest)(nil),  // 2: com.docker.api.protos.compose.v1.ComposeDownRequest
	(*ComposeDownResponse)(nil), // 3: com.docker.api.protos.compose.v1.ComposeDownResponse
}
var file_protos_compose_v1_compose_proto_depIdxs = []int32{
	0, // 0: com.docker.api.protos.compose.v1.Compose.Up:input_type -> com.docker.api.protos.compose.v1.ComposeUpRequest
	2, // 1: com.docker.api.protos.compose.v1.Compose.Down:input_type -> com.docker.api.protos.compose.v1.ComposeDownRequest
	1, // 2: com.docker.api.protos.compose.v1.Compose.Up:output_type -> com.docker.api.protos.compose.v1.ComposeUpResponse
	3, // 3: com.docker.api.protos.compose.v1.Compose.Down:output_type -> com.docker.api.protos.compose.v1.ComposeDownResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_compose_v1_compose_proto_init() }
func file_protos_compose_v1_compose_proto_init() {
	if File_protos_compose_v1_compose_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_compose_v1_compose_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeUpRequest); i {
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
		file_protos_compose_v1_compose_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeUpResponse); i {
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
		file_protos_compose_v1_compose_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeDownRequest); i {
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
		file_protos_compose_v1_compose_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeDownResponse); i {
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
			RawDescriptor: file_protos_compose_v1_compose_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_compose_v1_compose_proto_goTypes,
		DependencyIndexes: file_protos_compose_v1_compose_proto_depIdxs,
		MessageInfos:      file_protos_compose_v1_compose_proto_msgTypes,
	}.Build()
	File_protos_compose_v1_compose_proto = out.File
	file_protos_compose_v1_compose_proto_rawDesc = nil
	file_protos_compose_v1_compose_proto_goTypes = nil
	file_protos_compose_v1_compose_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ComposeClient is the client API for Compose service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ComposeClient interface {
	Up(ctx context.Context, in *ComposeUpRequest, opts ...grpc.CallOption) (*ComposeUpResponse, error)
	Down(ctx context.Context, in *ComposeDownRequest, opts ...grpc.CallOption) (*ComposeDownResponse, error)
}

type composeClient struct {
	cc grpc.ClientConnInterface
}

func NewComposeClient(cc grpc.ClientConnInterface) ComposeClient {
	return &composeClient{cc}
}

func (c *composeClient) Up(ctx context.Context, in *ComposeUpRequest, opts ...grpc.CallOption) (*ComposeUpResponse, error) {
	out := new(ComposeUpResponse)
	err := c.cc.Invoke(ctx, "/com.docker.api.protos.compose.v1.Compose/Up", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *composeClient) Down(ctx context.Context, in *ComposeDownRequest, opts ...grpc.CallOption) (*ComposeDownResponse, error) {
	out := new(ComposeDownResponse)
	err := c.cc.Invoke(ctx, "/com.docker.api.protos.compose.v1.Compose/Down", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComposeServer is the server API for Compose service.
type ComposeServer interface {
	Up(context.Context, *ComposeUpRequest) (*ComposeUpResponse, error)
	Down(context.Context, *ComposeDownRequest) (*ComposeDownResponse, error)
}

// UnimplementedComposeServer can be embedded to have forward compatible implementations.
type UnimplementedComposeServer struct {
}

func (*UnimplementedComposeServer) Up(context.Context, *ComposeUpRequest) (*ComposeUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Up not implemented")
}
func (*UnimplementedComposeServer) Down(context.Context, *ComposeDownRequest) (*ComposeDownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Down not implemented")
}

func RegisterComposeServer(s *grpc.Server, srv ComposeServer) {
	s.RegisterService(&_Compose_serviceDesc, srv)
}

func _Compose_Up_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComposeUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComposeServer).Up(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.docker.api.protos.compose.v1.Compose/Up",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComposeServer).Up(ctx, req.(*ComposeUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Compose_Down_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComposeDownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComposeServer).Down(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.docker.api.protos.compose.v1.Compose/Down",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComposeServer).Down(ctx, req.(*ComposeDownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Compose_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.docker.api.protos.compose.v1.Compose",
	HandlerType: (*ComposeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Up",
			Handler:    _Compose_Up_Handler,
		},
		{
			MethodName: "Down",
			Handler:    _Compose_Down_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/compose/v1/compose.proto",
}
