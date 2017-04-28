// Code generated by protoc-gen-go.
// source: fhirServer.proto
// DO NOT EDIT!

package buffer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ID struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *ID) Reset()                    { *m = ID{} }
func (m *ID) String() string            { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()               {}
func (*ID) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *ID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ID) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type QueryString struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *QueryString) Reset()                    { *m = QueryString{} }
func (m *QueryString) String() string            { return proto.CompactTextString(m) }
func (*QueryString) ProtoMessage()               {}
func (*QueryString) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *QueryString) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type FHIRUpdateResource struct {
	SytemIdentifier *ID            `protobuf:"bytes,2,opt,name=sytemIdentifier" json:"sytemIdentifier,omitempty"`
	Resource        *OneOfResource `protobuf:"bytes,3,opt,name=resource" json:"resource,omitempty"`
}

func (m *FHIRUpdateResource) Reset()                    { *m = FHIRUpdateResource{} }
func (m *FHIRUpdateResource) String() string            { return proto.CompactTextString(m) }
func (*FHIRUpdateResource) ProtoMessage()               {}
func (*FHIRUpdateResource) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *FHIRUpdateResource) GetSytemIdentifier() *ID {
	if m != nil {
		return m.SytemIdentifier
	}
	return nil
}

func (m *FHIRUpdateResource) GetResource() *OneOfResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func init() {
	proto.RegisterType((*ID)(nil), "buffer.ID")
	proto.RegisterType((*QueryString)(nil), "buffer.QueryString")
	proto.RegisterType((*FHIRUpdateResource)(nil), "buffer.FHIRUpdateResource")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FHIR service

type FHIRClient interface {
	GetResource(ctx context.Context, in *ID, opts ...grpc.CallOption) (*OneOfResource, error)
	CreateResource(ctx context.Context, in *FHIRUpdateResource, opts ...grpc.CallOption) (*OneOfResource, error)
	UpdateResource(ctx context.Context, in *FHIRUpdateResource, opts ...grpc.CallOption) (*OneOfResource, error)
	DeleteResource(ctx context.Context, in *ID, opts ...grpc.CallOption) (*OneOfResource, error)
	SearchResource(ctx context.Context, in *QueryString, opts ...grpc.CallOption) (*Bundle, error)
	BatchInsertResource(ctx context.Context, in *Bundle, opts ...grpc.CallOption) (*Bundle, error)
}

type fHIRClient struct {
	cc *grpc.ClientConn
}

func NewFHIRClient(cc *grpc.ClientConn) FHIRClient {
	return &fHIRClient{cc}
}

func (c *fHIRClient) GetResource(ctx context.Context, in *ID, opts ...grpc.CallOption) (*OneOfResource, error) {
	out := new(OneOfResource)
	err := grpc.Invoke(ctx, "/buffer.FHIR/GetResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fHIRClient) CreateResource(ctx context.Context, in *FHIRUpdateResource, opts ...grpc.CallOption) (*OneOfResource, error) {
	out := new(OneOfResource)
	err := grpc.Invoke(ctx, "/buffer.FHIR/CreateResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fHIRClient) UpdateResource(ctx context.Context, in *FHIRUpdateResource, opts ...grpc.CallOption) (*OneOfResource, error) {
	out := new(OneOfResource)
	err := grpc.Invoke(ctx, "/buffer.FHIR/UpdateResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fHIRClient) DeleteResource(ctx context.Context, in *ID, opts ...grpc.CallOption) (*OneOfResource, error) {
	out := new(OneOfResource)
	err := grpc.Invoke(ctx, "/buffer.FHIR/DeleteResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fHIRClient) SearchResource(ctx context.Context, in *QueryString, opts ...grpc.CallOption) (*Bundle, error) {
	out := new(Bundle)
	err := grpc.Invoke(ctx, "/buffer.FHIR/SearchResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fHIRClient) BatchInsertResource(ctx context.Context, in *Bundle, opts ...grpc.CallOption) (*Bundle, error) {
	out := new(Bundle)
	err := grpc.Invoke(ctx, "/buffer.FHIR/BatchInsertResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FHIR service

type FHIRServer interface {
	GetResource(context.Context, *ID) (*OneOfResource, error)
	CreateResource(context.Context, *FHIRUpdateResource) (*OneOfResource, error)
	UpdateResource(context.Context, *FHIRUpdateResource) (*OneOfResource, error)
	DeleteResource(context.Context, *ID) (*OneOfResource, error)
	SearchResource(context.Context, *QueryString) (*Bundle, error)
	BatchInsertResource(context.Context, *Bundle) (*Bundle, error)
}

func RegisterFHIRServer(s *grpc.Server, srv FHIRServer) {
	s.RegisterService(&_FHIR_serviceDesc, srv)
}

func _FHIR_GetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FHIRServer).GetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buffer.FHIR/GetResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FHIRServer).GetResource(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FHIR_CreateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FHIRUpdateResource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FHIRServer).CreateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buffer.FHIR/CreateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FHIRServer).CreateResource(ctx, req.(*FHIRUpdateResource))
	}
	return interceptor(ctx, in, info, handler)
}

func _FHIR_UpdateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FHIRUpdateResource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FHIRServer).UpdateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buffer.FHIR/UpdateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FHIRServer).UpdateResource(ctx, req.(*FHIRUpdateResource))
	}
	return interceptor(ctx, in, info, handler)
}

func _FHIR_DeleteResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FHIRServer).DeleteResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buffer.FHIR/DeleteResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FHIRServer).DeleteResource(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FHIR_SearchResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FHIRServer).SearchResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buffer.FHIR/SearchResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FHIRServer).SearchResource(ctx, req.(*QueryString))
	}
	return interceptor(ctx, in, info, handler)
}

func _FHIR_BatchInsertResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bundle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FHIRServer).BatchInsertResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buffer.FHIR/BatchInsertResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FHIRServer).BatchInsertResource(ctx, req.(*Bundle))
	}
	return interceptor(ctx, in, info, handler)
}

var _FHIR_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buffer.FHIR",
	HandlerType: (*FHIRServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResource",
			Handler:    _FHIR_GetResource_Handler,
		},
		{
			MethodName: "CreateResource",
			Handler:    _FHIR_CreateResource_Handler,
		},
		{
			MethodName: "UpdateResource",
			Handler:    _FHIR_UpdateResource_Handler,
		},
		{
			MethodName: "DeleteResource",
			Handler:    _FHIR_DeleteResource_Handler,
		},
		{
			MethodName: "SearchResource",
			Handler:    _FHIR_SearchResource_Handler,
		},
		{
			MethodName: "BatchInsertResource",
			Handler:    _FHIR_BatchInsertResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fhirServer.proto",
}

func init() { proto.RegisterFile("fhirServer.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x4e, 0x83, 0x40,
	0x10, 0xc7, 0x2d, 0xd6, 0x46, 0xa7, 0x06, 0xcd, 0x56, 0x13, 0xc2, 0xc9, 0xe0, 0xa5, 0x27, 0x12,
	0x51, 0xe3, 0xbd, 0xe2, 0x07, 0xa7, 0x46, 0x88, 0x0f, 0xc0, 0xc7, 0x20, 0x9b, 0x54, 0xc0, 0x61,
	0x31, 0xe1, 0xe0, 0x43, 0xfb, 0x06, 0xa6, 0x4b, 0x17, 0x09, 0x35, 0xc6, 0x78, 0x63, 0x27, 0xbf,
	0xff, 0x6f, 0x67, 0xd8, 0x81, 0xe3, 0x34, 0xe3, 0x14, 0x20, 0xbd, 0x23, 0xd9, 0x25, 0x15, 0xa2,
	0x60, 0x93, 0xa8, 0x4e, 0x53, 0x24, 0xf3, 0x30, 0xaa, 0xf3, 0x64, 0x85, 0x6d, 0xd5, 0x9a, 0x83,
	0xe6, 0xb9, 0x4c, 0x07, 0x8d, 0x27, 0xc6, 0xe8, 0x6c, 0x34, 0x3f, 0xf0, 0x35, 0x9e, 0x30, 0x06,
	0x63, 0xd1, 0x94, 0x68, 0x68, 0xb2, 0x22, 0xbf, 0xad, 0x73, 0x98, 0x3e, 0xd5, 0x48, 0x4d, 0x20,
	0x88, 0xe7, 0x2f, 0xec, 0x04, 0xf6, 0xde, 0xd6, 0xc7, 0x4d, 0xaa, 0x3d, 0x58, 0x1f, 0xc0, 0xee,
	0x1f, 0x3d, 0xff, 0xb9, 0x4c, 0x42, 0x81, 0x3e, 0x56, 0x45, 0x4d, 0x31, 0xb2, 0x2b, 0x38, 0xaa,
	0x1a, 0x81, 0xaf, 0x5e, 0x82, 0xb9, 0xe0, 0x29, 0x47, 0x92, 0xe6, 0xa9, 0x03, 0x76, 0xdb, 0x94,
	0xed, 0xb9, 0xfe, 0x10, 0x61, 0x17, 0xb0, 0x4f, 0x1b, 0x83, 0xb1, 0x2b, 0xf1, 0x53, 0x85, 0x2f,
	0x73, 0x5c, 0xa6, 0x4a, 0xef, 0x77, 0x98, 0xf3, 0xa9, 0xc1, 0x78, 0x7d, 0x3f, 0x73, 0x60, 0xfa,
	0x80, 0xa2, 0x6b, 0xa0, 0x77, 0x8f, 0xf9, 0xb3, 0xc4, 0xda, 0x61, 0x77, 0xa0, 0xdf, 0x12, 0xf6,
	0xfb, 0x36, 0x15, 0xba, 0x3d, 0xd3, 0xaf, 0x9a, 0xc1, 0xf8, 0xff, 0xd2, 0x5c, 0x83, 0xee, 0xe2,
	0x0a, 0x7b, 0x9a, 0x3f, 0x0d, 0x71, 0x03, 0x7a, 0x80, 0x21, 0xc5, 0x59, 0x17, 0x9b, 0x29, 0xb4,
	0xf7, 0x7a, 0xa6, 0xae, 0x8a, 0x0b, 0xb9, 0x0c, 0x32, 0x38, 0x5b, 0x84, 0x22, 0xce, 0xbc, 0xbc,
	0x42, 0xfa, 0xfe, 0x73, 0x03, 0x70, 0x3b, 0x18, 0x4d, 0xe4, 0x22, 0x5d, 0x7e, 0x05, 0x00, 0x00,
	0xff, 0xff, 0xc2, 0x51, 0xfc, 0x9f, 0x72, 0x02, 0x00, 0x00,
}