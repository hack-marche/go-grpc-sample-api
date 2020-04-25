// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: popularity.proto

package popularity

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

type PopularityListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemCode string `protobuf:"bytes,1,opt,name=system_code,json=systemCode,proto3" json:"system_code,omitempty"`
}

func (x *PopularityListRequest) Reset() {
	*x = PopularityListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_popularity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopularityListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopularityListRequest) ProtoMessage() {}

func (x *PopularityListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_popularity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopularityListRequest.ProtoReflect.Descriptor instead.
func (*PopularityListRequest) Descriptor() ([]byte, []int) {
	return file_popularity_proto_rawDescGZIP(), []int{0}
}

func (x *PopularityListRequest) GetSystemCode() string {
	if x != nil {
		return x.SystemCode
	}
	return ""
}

type PopularityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemCode   string `protobuf:"bytes,1,opt,name=system_code,json=systemCode,proto3" json:"system_code,omitempty"`
	PopularityId string `protobuf:"bytes,2,opt,name=popularity_id,json=popularityId,proto3" json:"popularity_id,omitempty"`
}

func (x *PopularityRequest) Reset() {
	*x = PopularityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_popularity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopularityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopularityRequest) ProtoMessage() {}

func (x *PopularityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_popularity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopularityRequest.ProtoReflect.Descriptor instead.
func (*PopularityRequest) Descriptor() ([]byte, []int) {
	return file_popularity_proto_rawDescGZIP(), []int{1}
}

func (x *PopularityRequest) GetSystemCode() string {
	if x != nil {
		return x.SystemCode
	}
	return ""
}

func (x *PopularityRequest) GetPopularityId() string {
	if x != nil {
		return x.PopularityId
	}
	return ""
}

type PopularityListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Popularitys []*PopularityResponse `protobuf:"bytes,1,rep,name=popularitys,proto3" json:"popularitys,omitempty"`
}

func (x *PopularityListResponse) Reset() {
	*x = PopularityListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_popularity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopularityListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopularityListResponse) ProtoMessage() {}

func (x *PopularityListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_popularity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopularityListResponse.ProtoReflect.Descriptor instead.
func (*PopularityListResponse) Descriptor() ([]byte, []int) {
	return file_popularity_proto_rawDescGZIP(), []int{2}
}

func (x *PopularityListResponse) GetPopularitys() []*PopularityResponse {
	if x != nil {
		return x.Popularitys
	}
	return nil
}

type PopularityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *PopularityResponse) Reset() {
	*x = PopularityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_popularity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopularityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopularityResponse) ProtoMessage() {}

func (x *PopularityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_popularity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopularityResponse.ProtoReflect.Descriptor instead.
func (*PopularityResponse) Descriptor() ([]byte, []int) {
	return file_popularity_proto_rawDescGZIP(), []int{3}
}

func (x *PopularityResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PopularityResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_popularity_proto protoreflect.FileDescriptor

var file_popularity_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x22, 0x38,
	0x0a, 0x15, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x59, 0x0a, 0x11, 0x50, 0x6f, 0x70, 0x75,
	0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x16, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a,
	0x0b, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0b, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x73, 0x22,
	0x42, 0x0a, 0x12, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x32, 0xbc, 0x01, 0x0a, 0x0a, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x5c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72,
	0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x6f, 0x70,
	0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x50, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x50,
	0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x6f,
	0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_popularity_proto_rawDescOnce sync.Once
	file_popularity_proto_rawDescData = file_popularity_proto_rawDesc
)

func file_popularity_proto_rawDescGZIP() []byte {
	file_popularity_proto_rawDescOnce.Do(func() {
		file_popularity_proto_rawDescData = protoimpl.X.CompressGZIP(file_popularity_proto_rawDescData)
	})
	return file_popularity_proto_rawDescData
}

var file_popularity_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_popularity_proto_goTypes = []interface{}{
	(*PopularityListRequest)(nil),  // 0: popularity.PopularityListRequest
	(*PopularityRequest)(nil),      // 1: popularity.PopularityRequest
	(*PopularityListResponse)(nil), // 2: popularity.PopularityListResponse
	(*PopularityResponse)(nil),     // 3: popularity.PopularityResponse
}
var file_popularity_proto_depIdxs = []int32{
	3, // 0: popularity.PopularityListResponse.popularitys:type_name -> popularity.PopularityResponse
	0, // 1: popularity.Popularity.GetPopularityList:input_type -> popularity.PopularityListRequest
	1, // 2: popularity.Popularity.GetPopularity:input_type -> popularity.PopularityRequest
	2, // 3: popularity.Popularity.GetPopularityList:output_type -> popularity.PopularityListResponse
	3, // 4: popularity.Popularity.GetPopularity:output_type -> popularity.PopularityResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_popularity_proto_init() }
func file_popularity_proto_init() {
	if File_popularity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_popularity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopularityListRequest); i {
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
		file_popularity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopularityRequest); i {
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
		file_popularity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopularityListResponse); i {
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
		file_popularity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopularityResponse); i {
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
			RawDescriptor: file_popularity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_popularity_proto_goTypes,
		DependencyIndexes: file_popularity_proto_depIdxs,
		MessageInfos:      file_popularity_proto_msgTypes,
	}.Build()
	File_popularity_proto = out.File
	file_popularity_proto_rawDesc = nil
	file_popularity_proto_goTypes = nil
	file_popularity_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PopularityClient is the client API for Popularity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PopularityClient interface {
	GetPopularityList(ctx context.Context, in *PopularityListRequest, opts ...grpc.CallOption) (*PopularityListResponse, error)
	GetPopularity(ctx context.Context, in *PopularityRequest, opts ...grpc.CallOption) (*PopularityResponse, error)
}

type popularityClient struct {
	cc grpc.ClientConnInterface
}

func NewPopularityClient(cc grpc.ClientConnInterface) PopularityClient {
	return &popularityClient{cc}
}

func (c *popularityClient) GetPopularityList(ctx context.Context, in *PopularityListRequest, opts ...grpc.CallOption) (*PopularityListResponse, error) {
	out := new(PopularityListResponse)
	err := c.cc.Invoke(ctx, "/popularity.Popularity/GetPopularityList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *popularityClient) GetPopularity(ctx context.Context, in *PopularityRequest, opts ...grpc.CallOption) (*PopularityResponse, error) {
	out := new(PopularityResponse)
	err := c.cc.Invoke(ctx, "/popularity.Popularity/GetPopularity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PopularityServer is the server API for Popularity service.
type PopularityServer interface {
	GetPopularityList(context.Context, *PopularityListRequest) (*PopularityListResponse, error)
	GetPopularity(context.Context, *PopularityRequest) (*PopularityResponse, error)
}

// UnimplementedPopularityServer can be embedded to have forward compatible implementations.
type UnimplementedPopularityServer struct {
}

func (*UnimplementedPopularityServer) GetPopularityList(context.Context, *PopularityListRequest) (*PopularityListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPopularityList not implemented")
}
func (*UnimplementedPopularityServer) GetPopularity(context.Context, *PopularityRequest) (*PopularityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPopularity not implemented")
}

func RegisterPopularityServer(s *grpc.Server, srv PopularityServer) {
	s.RegisterService(&_Popularity_serviceDesc, srv)
}

func _Popularity_GetPopularityList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PopularityListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PopularityServer).GetPopularityList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/popularity.Popularity/GetPopularityList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PopularityServer).GetPopularityList(ctx, req.(*PopularityListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Popularity_GetPopularity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PopularityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PopularityServer).GetPopularity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/popularity.Popularity/GetPopularity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PopularityServer).GetPopularity(ctx, req.(*PopularityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Popularity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "popularity.Popularity",
	HandlerType: (*PopularityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPopularityList",
			Handler:    _Popularity_GetPopularityList_Handler,
		},
		{
			MethodName: "GetPopularity",
			Handler:    _Popularity_GetPopularity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "popularity.proto",
}