// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sinks/grpsink/grpc_sink.proto

package grpsink

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	ssf "github.com/stripe/veneur/v14/ssf"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8c7d48a83a8c7b, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return m.Size()
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Empty)(nil), "grpsink.Empty")
}

func init() { proto.RegisterFile("sinks/grpsink/grpc_sink.proto", fileDescriptor_5c8c7d48a83a8c7b) }

var fileDescriptor_5c8c7d48a83a8c7b = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0xce, 0xcc, 0xcb,
	0x2e, 0xd6, 0x4f, 0x2f, 0x2a, 0x00, 0x31, 0x40, 0x74, 0x72, 0x3c, 0x88, 0xa5, 0x57, 0x50, 0x94,
	0x5f, 0x92, 0x2f, 0xc4, 0x0e, 0x95, 0x90, 0x12, 0x28, 0x2e, 0x4e, 0xd3, 0x2f, 0x4e, 0xcc, 0x2d,
	0xc8, 0x49, 0x85, 0x48, 0x29, 0xb1, 0x73, 0xb1, 0xba, 0xe6, 0x16, 0x94, 0x54, 0x1a, 0x99, 0x70,
	0x71, 0x04, 0x17, 0x24, 0xe6, 0x05, 0x67, 0xe6, 0x65, 0x0b, 0x69, 0x70, 0x71, 0x04, 0xa7, 0xe6,
	0xa5, 0x80, 0xf8, 0x42, 0x3c, 0x7a, 0xc5, 0xc5, 0x69, 0x7a, 0xc1, 0xc1, 0x6e, 0x20, 0x9e, 0x14,
	0x9f, 0x1e, 0xd4, 0x28, 0x3d, 0xb0, 0x2e, 0x27, 0x89, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c,
	0x96, 0x63, 0x48, 0x62, 0x03, 0x9b, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x54, 0xd6,
	0x2f, 0x9b, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SpanSinkClient is the client API for SpanSink service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpanSinkClient interface {
	SendSpan(ctx context.Context, in *ssf.SSFSpan, opts ...grpc.CallOption) (*Empty, error)
}

type spanSinkClient struct {
	cc *grpc.ClientConn
}

func NewSpanSinkClient(cc *grpc.ClientConn) SpanSinkClient {
	return &spanSinkClient{cc}
}

func (c *spanSinkClient) SendSpan(ctx context.Context, in *ssf.SSFSpan, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/grpsink.SpanSink/SendSpan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpanSinkServer is the server API for SpanSink service.
type SpanSinkServer interface {
	SendSpan(context.Context, *ssf.SSFSpan) (*Empty, error)
}

func RegisterSpanSinkServer(s *grpc.Server, srv SpanSinkServer) {
	s.RegisterService(&_SpanSink_serviceDesc, srv)
}

func _SpanSink_SendSpan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ssf.SSFSpan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpanSinkServer).SendSpan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpsink.SpanSink/SendSpan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpanSinkServer).SendSpan(ctx, req.(*ssf.SSFSpan))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpanSink_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpsink.SpanSink",
	HandlerType: (*SpanSinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSpan",
			Handler:    _SpanSink_SendSpan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sinks/grpsink/grpc_sink.proto",
}

func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintGrpcSink(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Empty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovGrpcSink(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGrpcSink(x uint64) (n int) {
	return sovGrpcSink(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpcSink
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGrpcSink(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpcSink
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGrpcSink
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGrpcSink(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGrpcSink
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGrpcSink
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGrpcSink
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGrpcSink
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthGrpcSink
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGrpcSink
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGrpcSink(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthGrpcSink
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGrpcSink = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGrpcSink   = fmt.Errorf("proto: integer overflow")
)
