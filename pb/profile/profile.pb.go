// Code generated by protoc-gen-go.
// source: pb/profile/profile.proto
// DO NOT EDIT!

/*
Package profile is a generated protocol buffer package.

It is generated from these files:
	pb/profile/profile.proto

It has these top-level messages:
	Request
	Result
	Hotel
	Address
	Image
*/
package profile

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	HotelIds []string `protobuf:"bytes,1,rep,name=hotelIds" json:"hotelIds,omitempty"`
	Locale   string   `protobuf:"bytes,2,opt,name=locale" json:"locale,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Result struct {
	Hotels []*Hotel `protobuf:"bytes,1,rep,name=hotels" json:"hotels,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetHotels() []*Hotel {
	if m != nil {
		return m.Hotels
	}
	return nil
}

type Hotel struct {
	Id          string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	PhoneNumber string   `protobuf:"bytes,3,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Address     *Address `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	Images      []*Image `protobuf:"bytes,6,rep,name=images" json:"images,omitempty"`
}

func (m *Hotel) Reset()                    { *m = Hotel{} }
func (m *Hotel) String() string            { return proto.CompactTextString(m) }
func (*Hotel) ProtoMessage()               {}
func (*Hotel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Hotel) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Hotel) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

type Address struct {
	StreetNumber string `protobuf:"bytes,1,opt,name=streetNumber" json:"streetNumber,omitempty"`
	StreetName   string `protobuf:"bytes,2,opt,name=streetName" json:"streetName,omitempty"`
	City         string `protobuf:"bytes,3,opt,name=city" json:"city,omitempty"`
	State        string `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	Country      string `protobuf:"bytes,5,opt,name=country" json:"country,omitempty"`
	PostalCode   string `protobuf:"bytes,6,opt,name=postalCode" json:"postalCode,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Image struct {
	Url     string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Default bool   `protobuf:"varint,2,opt,name=default" json:"default,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Request)(nil), "profile.Request")
	proto.RegisterType((*Result)(nil), "profile.Result")
	proto.RegisterType((*Hotel)(nil), "profile.Hotel")
	proto.RegisterType((*Address)(nil), "profile.Address")
	proto.RegisterType((*Image)(nil), "profile.Image")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Profile service

type ProfileClient interface {
	GetProfiles(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type profileClient struct {
	cc *grpc.ClientConn
}

func NewProfileClient(cc *grpc.ClientConn) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) GetProfiles(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/profile.Profile/GetProfiles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Profile service

type ProfileServer interface {
	GetProfiles(context.Context, *Request) (*Result, error)
}

func RegisterProfileServer(s *grpc.Server, srv ProfileServer) {
	s.RegisterService(&_Profile_serviceDesc, srv)
}

func _Profile_GetProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/GetProfiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfiles(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Profile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profile.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfiles",
			Handler:    _Profile_GetProfiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("pb/profile/profile.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x52, 0xc1, 0x6e, 0xea, 0x30,
	0x10, 0x54, 0x80, 0x24, 0xb0, 0x79, 0xe2, 0xa1, 0xd5, 0xd3, 0x93, 0xc5, 0xa1, 0x42, 0x39, 0x54,
	0x55, 0x0f, 0x14, 0xc1, 0xb1, 0xea, 0xa1, 0xea, 0xa1, 0xe5, 0x52, 0x55, 0xfe, 0x83, 0x90, 0x2c,
	0x25, 0x52, 0x88, 0xd3, 0xd8, 0x39, 0xf0, 0x59, 0xfd, 0x86, 0xfe, 0x58, 0x6d, 0xc7, 0x81, 0xb4,
	0xa7, 0xec, 0xcc, 0x8e, 0xbd, 0x33, 0xce, 0x02, 0xab, 0x76, 0x77, 0x55, 0x2d, 0xf6, 0x79, 0x41,
	0xdd, 0x77, 0xa9, 0xbf, 0x4a, 0x60, 0xe8, 0x60, 0xfc, 0x00, 0x21, 0xa7, 0x8f, 0x86, 0xa4, 0xc2,
	0x39, 0x8c, 0x0f, 0x42, 0x51, 0xb1, 0xcd, 0x24, 0xf3, 0x16, 0xc3, 0x9b, 0x09, 0x3f, 0x63, 0xfc,
	0x0f, 0x41, 0x21, 0xd2, 0xa4, 0x20, 0x36, 0x58, 0x78, 0xba, 0xe3, 0x50, 0xbc, 0x82, 0x80, 0x93,
	0x6c, 0x0a, 0x85, 0xd7, 0x10, 0x58, 0x75, 0x7b, 0x36, 0x5a, 0x4f, 0x97, 0xdd, 0xc4, 0x17, 0x43,
	0x73, 0xd7, 0x8d, 0xbf, 0x3c, 0xf0, 0x2d, 0x83, 0x53, 0x18, 0xe4, 0x99, 0x56, 0x9b, 0xfb, 0x74,
	0x85, 0x08, 0xa3, 0x32, 0x39, 0x76, 0x13, 0x6c, 0x8d, 0x0b, 0x88, 0xaa, 0x83, 0x28, 0xe9, 0xb5,
	0x39, 0xee, 0xa8, 0x66, 0x43, 0xdb, 0xea, 0x53, 0x46, 0x91, 0x91, 0x4c, 0xeb, 0xbc, 0x52, 0xb9,
	0x28, 0xd9, 0xa8, 0x55, 0xf4, 0x28, 0xbc, 0x85, 0x30, 0xc9, 0xb2, 0x9a, 0xa4, 0x64, 0xbe, 0xee,
	0x46, 0xeb, 0xd9, 0xd9, 0xda, 0x63, 0xcb, 0xf3, 0x4e, 0x60, 0x52, 0xe4, 0xc7, 0xe4, 0x9d, 0x24,
	0x0b, 0x7e, 0xa5, 0xd8, 0x1a, 0x9a, 0xbb, 0x6e, 0xfc, 0xe9, 0x41, 0xe8, 0x0e, 0x63, 0x0c, 0x7f,
	0xa4, 0xaa, 0x89, 0x94, 0x33, 0xd9, 0x26, 0xfa, 0xc1, 0xe1, 0x15, 0x80, 0xc3, 0x97, 0x84, 0x3d,
	0xc6, 0x64, 0x4f, 0x73, 0x75, 0x72, 0x01, 0x6d, 0x8d, 0xff, 0xc0, 0x97, 0x2a, 0x51, 0xe4, 0x32,
	0xb5, 0x00, 0x19, 0x84, 0xa9, 0x68, 0x4a, 0x55, 0x9f, 0x6c, 0x9a, 0x09, 0xef, 0xa0, 0x99, 0x51,
	0x09, 0x2d, 0x2a, 0x9e, 0x44, 0x46, 0xda, 0xbf, 0x9d, 0x71, 0x61, 0xe2, 0x0d, 0xf8, 0x36, 0x04,
	0xce, 0x60, 0xd8, 0xd4, 0x85, 0xf3, 0x69, 0x4a, 0x73, 0x69, 0x46, 0xfb, 0x44, 0xff, 0x47, 0xeb,
	0x6d, 0xcc, 0x3b, 0xb8, 0xbe, 0x87, 0xf0, 0xad, 0x7d, 0x01, 0x5c, 0x41, 0xf4, 0x4c, 0xca, 0x21,
	0x89, 0x97, 0x57, 0x74, 0x0b, 0x34, 0xff, 0xdb, 0x63, 0xcc, 0x4e, 0xec, 0x02, 0xbb, 0x6c, 0x9b,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0x0b, 0x14, 0x13, 0x88, 0x02, 0x00, 0x00,
}