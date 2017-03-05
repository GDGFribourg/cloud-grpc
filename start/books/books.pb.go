//   Copyright 2016, Google, Inc.
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//       http://www.apache.org/licenses/LICENSE-2.0
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by protoc-gen-go.
// source: books.proto
// DO NOT EDIT!

/*
Package books is a generated protocol buffer package.

It is generated from these files:
	books.proto

It has these top-level messages:
	Empty
	Book
	BookList
	BookIdRequest
*/
package books

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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Book struct {
	Id     int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
}

func (m *Book) Reset()                    { *m = Book{} }
func (m *Book) String() string            { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()               {}
func (*Book) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Book) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type BookList struct {
	Books []*Book `protobuf:"bytes,1,rep,name=books" json:"books,omitempty"`
}

func (m *BookList) Reset()                    { *m = BookList{} }
func (m *BookList) String() string            { return proto.CompactTextString(m) }
func (*BookList) ProtoMessage()               {}
func (*BookList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BookList) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

type BookIdRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *BookIdRequest) Reset()                    { *m = BookIdRequest{} }
func (m *BookIdRequest) String() string            { return proto.CompactTextString(m) }
func (*BookIdRequest) ProtoMessage()               {}
func (*BookIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BookIdRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "books.Empty")
	proto.RegisterType((*Book)(nil), "books.Book")
	proto.RegisterType((*BookList)(nil), "books.BookList")
	proto.RegisterType((*BookIdRequest)(nil), "books.BookIdRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BookService service

type BookServiceClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BookList, error)
	Insert(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Empty, error)
	Get(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*Book, error)
	Delete(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*Empty, error)
	Watch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (BookService_WatchClient, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := grpc.Invoke(ctx, "/books.BookService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Insert(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/books.BookService/Insert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Get(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/books.BookService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Delete(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/books.BookService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Watch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (BookService_WatchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BookService_serviceDesc.Streams[0], c.cc, "/books.BookService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookService_WatchClient interface {
	Recv() (*Book, error)
	grpc.ClientStream
}

type bookServiceWatchClient struct {
	grpc.ClientStream
}

func (x *bookServiceWatchClient) Recv() (*Book, error) {
	m := new(Book)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for BookService service

type BookServiceServer interface {
	List(context.Context, *Empty) (*BookList, error)
	Insert(context.Context, *Book) (*Empty, error)
	Get(context.Context, *BookIdRequest) (*Book, error)
	Delete(context.Context, *BookIdRequest) (*Empty, error)
	Watch(*Empty, BookService_WatchServer) error
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.BookService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.BookService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Insert(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.BookService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Get(ctx, req.(*BookIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.BookService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Delete(ctx, req.(*BookIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookServiceServer).Watch(m, &bookServiceWatchServer{stream})
}

type BookService_WatchServer interface {
	Send(*Book) error
	grpc.ServerStream
}

type bookServiceWatchServer struct {
	grpc.ServerStream
}

func (x *bookServiceWatchServer) Send(m *Book) error {
	return x.ServerStream.SendMsg(m)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "books.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _BookService_List_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _BookService_Insert_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BookService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BookService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _BookService_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "books.proto",
}

func init() { proto.RegisterFile("books.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xb3, 0x49, 0x37, 0xea, 0xc4, 0x3f, 0x30, 0x14, 0x09, 0xb9, 0x18, 0x17, 0xd4, 0x20,
	0x18, 0xa4, 0xbe, 0x81, 0x54, 0xa4, 0xe0, 0x29, 0x1e, 0x3c, 0xb7, 0xcd, 0x40, 0x97, 0x56, 0xb7,
	0x26, 0x53, 0xc1, 0xc7, 0xf5, 0x4d, 0x64, 0x37, 0x41, 0xb6, 0x01, 0x8f, 0xdf, 0xce, 0xcc, 0xef,
	0xf7, 0xc1, 0x42, 0xb2, 0x30, 0x66, 0xdd, 0x96, 0xdb, 0xc6, 0xb0, 0x41, 0xe9, 0x82, 0x3a, 0x00,
	0xf9, 0xf4, 0xbe, 0xe5, 0x6f, 0x35, 0x85, 0xd1, 0xa3, 0x31, 0x6b, 0x3c, 0x85, 0x50, 0xd7, 0xa9,
	0xc8, 0x45, 0x21, 0xab, 0x50, 0xd7, 0x38, 0x06, 0xc9, 0x9a, 0x37, 0x94, 0x86, 0xb9, 0x28, 0x8e,
	0xaa, 0x2e, 0xe0, 0x39, 0xc4, 0xf3, 0x1d, 0xaf, 0x4c, 0x93, 0x46, 0xee, 0xb9, 0x4f, 0xea, 0x0e,
	0x0e, 0x2d, 0xe5, 0x45, 0xb7, 0x8c, 0x97, 0xd0, 0x39, 0x52, 0x91, 0x47, 0x45, 0x32, 0x49, 0xca,
	0x4e, 0x6f, 0xe7, 0x55, 0x6f, 0xbf, 0x80, 0x13, 0x1b, 0x67, 0x75, 0x45, 0x9f, 0x3b, 0x6a, 0x79,
	0x68, 0x9f, 0xfc, 0x08, 0x48, 0xec, 0xc6, 0x2b, 0x35, 0x5f, 0x7a, 0x49, 0x78, 0x03, 0x23, 0xc7,
	0x3e, 0xee, 0x61, 0xae, 0x7b, 0x76, 0xe6, 0xa1, 0xed, 0x58, 0x05, 0x78, 0x05, 0xf1, 0xec, 0xa3,
	0xa5, 0x86, 0xd1, 0xf7, 0x66, 0x7b, 0x77, 0x2a, 0xc0, 0x5b, 0x88, 0x9e, 0x89, 0x71, 0xec, 0xed,
	0xfc, 0x95, 0xc9, 0xfc, 0x4b, 0x15, 0x60, 0x09, 0xf1, 0x94, 0x36, 0xc4, 0xf4, 0xcf, 0xfa, 0x90,
	0x7d, 0x0d, 0xf2, 0x6d, 0xce, 0xcb, 0xd5, 0xa0, 0xec, 0x3e, 0xf5, 0x5e, 0x2c, 0x62, 0xf7, 0x21,
	0x0f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xef, 0xf7, 0x61, 0x7c, 0x9f, 0x01, 0x00, 0x00,
}
