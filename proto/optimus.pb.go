// Code generated by protoc-gen-go. DO NOT EDIT.
// source: optimus.proto

package sonm

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

type PredictSupplierRequest struct {
	Devices *DevicesReply `protobuf:"bytes,1,opt,name=devices" json:"devices,omitempty"`
}

func (m *PredictSupplierRequest) Reset()                    { *m = PredictSupplierRequest{} }
func (m *PredictSupplierRequest) String() string            { return proto.CompactTextString(m) }
func (*PredictSupplierRequest) ProtoMessage()               {}
func (*PredictSupplierRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *PredictSupplierRequest) GetDevices() *DevicesReply {
	if m != nil {
		return m.Devices
	}
	return nil
}

type PredictSupplierReply struct {
	Price    *Price    `protobuf:"bytes,1,opt,name=price" json:"price,omitempty"`
	OrderIDs []*BigInt `protobuf:"bytes,2,rep,name=orderIDs" json:"orderIDs,omitempty"`
}

func (m *PredictSupplierReply) Reset()                    { *m = PredictSupplierReply{} }
func (m *PredictSupplierReply) String() string            { return proto.CompactTextString(m) }
func (*PredictSupplierReply) ProtoMessage()               {}
func (*PredictSupplierReply) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *PredictSupplierReply) GetPrice() *Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *PredictSupplierReply) GetOrderIDs() []*BigInt {
	if m != nil {
		return m.OrderIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*PredictSupplierRequest)(nil), "sonm.PredictSupplierRequest")
	proto.RegisterType((*PredictSupplierReply)(nil), "sonm.PredictSupplierReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for OrderPredictor service

type OrderPredictorClient interface {
	Predict(ctx context.Context, in *BidResources, opts ...grpc.CallOption) (*Price, error)
	PredictSupplier(ctx context.Context, in *PredictSupplierRequest, opts ...grpc.CallOption) (*PredictSupplierReply, error)
}

type orderPredictorClient struct {
	cc *grpc.ClientConn
}

func NewOrderPredictorClient(cc *grpc.ClientConn) OrderPredictorClient {
	return &orderPredictorClient{cc}
}

func (c *orderPredictorClient) Predict(ctx context.Context, in *BidResources, opts ...grpc.CallOption) (*Price, error) {
	out := new(Price)
	err := grpc.Invoke(ctx, "/sonm.OrderPredictor/Predict", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderPredictorClient) PredictSupplier(ctx context.Context, in *PredictSupplierRequest, opts ...grpc.CallOption) (*PredictSupplierReply, error) {
	out := new(PredictSupplierReply)
	err := grpc.Invoke(ctx, "/sonm.OrderPredictor/PredictSupplier", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderPredictor service

type OrderPredictorServer interface {
	Predict(context.Context, *BidResources) (*Price, error)
	PredictSupplier(context.Context, *PredictSupplierRequest) (*PredictSupplierReply, error)
}

func RegisterOrderPredictorServer(s *grpc.Server, srv OrderPredictorServer) {
	s.RegisterService(&_OrderPredictor_serviceDesc, srv)
}

func _OrderPredictor_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidResources)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderPredictorServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.OrderPredictor/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderPredictorServer).Predict(ctx, req.(*BidResources))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderPredictor_PredictSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderPredictorServer).PredictSupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.OrderPredictor/PredictSupplier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderPredictorServer).PredictSupplier(ctx, req.(*PredictSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderPredictor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.OrderPredictor",
	HandlerType: (*OrderPredictorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _OrderPredictor_Predict_Handler,
		},
		{
			MethodName: "PredictSupplier",
			Handler:    _OrderPredictor_PredictSupplier_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "optimus.proto",
}

func init() { proto.RegisterFile("optimus.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xdd, 0x4a, 0xfc, 0x30,
	0x10, 0xc5, 0xb7, 0xff, 0xbf, 0x5f, 0xa4, 0xd5, 0xc5, 0x20, 0xb2, 0x14, 0x2f, 0xd6, 0x5e, 0xf5,
	0x42, 0x7a, 0x51, 0xdf, 0x40, 0x16, 0x61, 0x2f, 0xc4, 0x25, 0x3e, 0x41, 0x37, 0x1d, 0x96, 0x61,
	0xdb, 0x24, 0x4e, 0x52, 0xa5, 0x0f, 0xe1, 0x3b, 0x4b, 0x4c, 0x2a, 0xe2, 0xc7, 0xe5, 0xf9, 0x9d,
	0x33, 0x67, 0x86, 0x61, 0xa7, 0xda, 0x38, 0xec, 0x07, 0x5b, 0x19, 0xd2, 0x4e, 0xf3, 0x03, 0xab,
	0x55, 0x9f, 0x67, 0x5b, 0xdc, 0xa1, 0x72, 0x81, 0xe5, 0x73, 0x54, 0x9e, 0x2a, 0x6c, 0x22, 0x38,
	0xef, 0x1b, 0xda, 0x83, 0x33, 0x5d, 0x23, 0x21, 0xa2, 0xec, 0x55, 0xd3, 0x1e, 0x28, 0xa8, 0xe2,
	0x9e, 0x5d, 0x6e, 0x08, 0x5a, 0x94, 0xee, 0x69, 0x30, 0xa6, 0x43, 0x20, 0x01, 0xcf, 0x03, 0x58,
	0xc7, 0x6f, 0xd8, 0x71, 0x0b, 0x2f, 0x28, 0xc1, 0x2e, 0x92, 0x65, 0x52, 0xa6, 0x35, 0xaf, 0x7c,
	0x77, 0xb5, 0x0a, 0x50, 0x80, 0xe9, 0x46, 0x31, 0x45, 0x0a, 0xc9, 0x2e, 0x7e, 0xf4, 0x98, 0x6e,
	0xe4, 0xd7, 0xec, 0xd0, 0x10, 0x4a, 0x88, 0x1d, 0x69, 0xe8, 0xd8, 0x78, 0x24, 0x82, 0xc3, 0x4b,
	0x76, 0xa2, 0xa9, 0x05, 0x5a, 0xaf, 0xec, 0xe2, 0xdf, 0xf2, 0x7f, 0x99, 0xd6, 0x59, 0x48, 0xdd,
	0xe1, 0x6e, 0xad, 0x9c, 0xf8, 0x74, 0xeb, 0xb7, 0x84, 0x9d, 0x3d, 0x7a, 0x11, 0x57, 0x69, 0xf2,
	0x57, 0x46, 0xc1, 0xf9, 0x34, 0xd5, 0x0a, 0xb0, 0x7a, 0x20, 0x09, 0x36, 0xff, 0xba, 0xaf, 0x98,
	0xf1, 0x07, 0x36, 0xff, 0x76, 0x25, 0xbf, 0x9a, 0x12, 0xbf, 0x3d, 0x21, 0xcf, 0xff, 0x70, 0x4d,
	0x37, 0x16, 0xb3, 0xed, 0xd1, 0xc7, 0x0f, 0x6f, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xf4,
	0x18, 0x0d, 0x9a, 0x01, 0x00, 0x00,
}
