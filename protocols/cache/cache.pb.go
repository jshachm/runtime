// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cache.proto

/*
	Package cache is a generated protocol buffer package.

	It is generated from these files:
		cache.proto

	It has these top-level messages:
		GrpcVMConfig
		GrpcVM
*/
package cache

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type GrpcVMConfig struct {
	Data        []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	AgentConfig []byte `protobuf:"bytes,2,opt,name=AgentConfig,proto3" json:"AgentConfig,omitempty"`
}

func (m *GrpcVMConfig) Reset()                    { *m = GrpcVMConfig{} }
func (m *GrpcVMConfig) String() string            { return proto.CompactTextString(m) }
func (*GrpcVMConfig) ProtoMessage()               {}
func (*GrpcVMConfig) Descriptor() ([]byte, []int) { return fileDescriptorCache, []int{0} }

func (m *GrpcVMConfig) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GrpcVMConfig) GetAgentConfig() []byte {
	if m != nil {
		return m.AgentConfig
	}
	return nil
}

type GrpcVM struct {
	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hypervisor []byte `protobuf:"bytes,2,opt,name=hypervisor,proto3" json:"hypervisor,omitempty"`
	ProxyPid   int64  `protobuf:"varint,3,opt,name=proxyPid,proto3" json:"proxyPid,omitempty"`
	ProxyURL   string `protobuf:"bytes,4,opt,name=proxyURL,proto3" json:"proxyURL,omitempty"`
	Cpu        uint32 `protobuf:"varint,5,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory     uint32 `protobuf:"varint,6,opt,name=memory,proto3" json:"memory,omitempty"`
	CpuDelta   uint32 `protobuf:"varint,7,opt,name=cpuDelta,proto3" json:"cpuDelta,omitempty"`
}

func (m *GrpcVM) Reset()                    { *m = GrpcVM{} }
func (m *GrpcVM) String() string            { return proto.CompactTextString(m) }
func (*GrpcVM) ProtoMessage()               {}
func (*GrpcVM) Descriptor() ([]byte, []int) { return fileDescriptorCache, []int{1} }

func (m *GrpcVM) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GrpcVM) GetHypervisor() []byte {
	if m != nil {
		return m.Hypervisor
	}
	return nil
}

func (m *GrpcVM) GetProxyPid() int64 {
	if m != nil {
		return m.ProxyPid
	}
	return 0
}

func (m *GrpcVM) GetProxyURL() string {
	if m != nil {
		return m.ProxyURL
	}
	return ""
}

func (m *GrpcVM) GetCpu() uint32 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *GrpcVM) GetMemory() uint32 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *GrpcVM) GetCpuDelta() uint32 {
	if m != nil {
		return m.CpuDelta
	}
	return 0
}

func init() {
	proto.RegisterType((*GrpcVMConfig)(nil), "cache.GrpcVMConfig")
	proto.RegisterType((*GrpcVM)(nil), "cache.GrpcVM")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CacheService service

type CacheServiceClient interface {
	Config(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*GrpcVMConfig, error)
	GetBaseVM(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*GrpcVM, error)
}

type cacheServiceClient struct {
	cc *grpc.ClientConn
}

func NewCacheServiceClient(cc *grpc.ClientConn) CacheServiceClient {
	return &cacheServiceClient{cc}
}

func (c *cacheServiceClient) Config(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*GrpcVMConfig, error) {
	out := new(GrpcVMConfig)
	err := grpc.Invoke(ctx, "/cache.CacheService/Config", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) GetBaseVM(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*GrpcVM, error) {
	out := new(GrpcVM)
	err := grpc.Invoke(ctx, "/cache.CacheService/GetBaseVM", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CacheService service

type CacheServiceServer interface {
	Config(context.Context, *google_protobuf.Empty) (*GrpcVMConfig, error)
	GetBaseVM(context.Context, *google_protobuf.Empty) (*GrpcVM, error)
}

func RegisterCacheServiceServer(s *grpc.Server, srv CacheServiceServer) {
	s.RegisterService(&_CacheService_serviceDesc, srv)
}

func _CacheService_Config_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Config(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache.CacheService/Config",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Config(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_GetBaseVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).GetBaseVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache.CacheService/GetBaseVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).GetBaseVM(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _CacheService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cache.CacheService",
	HandlerType: (*CacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Config",
			Handler:    _CacheService_Config_Handler,
		},
		{
			MethodName: "GetBaseVM",
			Handler:    _CacheService_GetBaseVM_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache.proto",
}

func (m *GrpcVMConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcVMConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCache(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if len(m.AgentConfig) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCache(dAtA, i, uint64(len(m.AgentConfig)))
		i += copy(dAtA[i:], m.AgentConfig)
	}
	return i, nil
}

func (m *GrpcVM) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcVM) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCache(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Hypervisor) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCache(dAtA, i, uint64(len(m.Hypervisor)))
		i += copy(dAtA[i:], m.Hypervisor)
	}
	if m.ProxyPid != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCache(dAtA, i, uint64(m.ProxyPid))
	}
	if len(m.ProxyURL) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCache(dAtA, i, uint64(len(m.ProxyURL)))
		i += copy(dAtA[i:], m.ProxyURL)
	}
	if m.Cpu != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintCache(dAtA, i, uint64(m.Cpu))
	}
	if m.Memory != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintCache(dAtA, i, uint64(m.Memory))
	}
	if m.CpuDelta != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintCache(dAtA, i, uint64(m.CpuDelta))
	}
	return i, nil
}

func encodeVarintCache(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GrpcVMConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovCache(uint64(l))
	}
	l = len(m.AgentConfig)
	if l > 0 {
		n += 1 + l + sovCache(uint64(l))
	}
	return n
}

func (m *GrpcVM) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCache(uint64(l))
	}
	l = len(m.Hypervisor)
	if l > 0 {
		n += 1 + l + sovCache(uint64(l))
	}
	if m.ProxyPid != 0 {
		n += 1 + sovCache(uint64(m.ProxyPid))
	}
	l = len(m.ProxyURL)
	if l > 0 {
		n += 1 + l + sovCache(uint64(l))
	}
	if m.Cpu != 0 {
		n += 1 + sovCache(uint64(m.Cpu))
	}
	if m.Memory != 0 {
		n += 1 + sovCache(uint64(m.Memory))
	}
	if m.CpuDelta != 0 {
		n += 1 + sovCache(uint64(m.CpuDelta))
	}
	return n
}

func sovCache(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCache(x uint64) (n int) {
	return sovCache(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GrpcVMConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCache
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GrpcVMConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrpcVMConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgentConfig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AgentConfig = append(m.AgentConfig[:0], dAtA[iNdEx:postIndex]...)
			if m.AgentConfig == nil {
				m.AgentConfig = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCache(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCache
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
func (m *GrpcVM) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCache
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GrpcVM: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrpcVM: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hypervisor", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hypervisor = append(m.Hypervisor[:0], dAtA[iNdEx:postIndex]...)
			if m.Hypervisor == nil {
				m.Hypervisor = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProxyPid", wireType)
			}
			m.ProxyPid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProxyPid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProxyURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProxyURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cpu", wireType)
			}
			m.Cpu = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Cpu |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memory", wireType)
			}
			m.Memory = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Memory |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CpuDelta", wireType)
			}
			m.CpuDelta = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CpuDelta |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCache(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCache
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
func skipCache(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCache
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
					return 0, ErrIntOverflowCache
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
					return 0, ErrIntOverflowCache
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCache
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCache
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
				next, err := skipCache(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthCache = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCache   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cache.proto", fileDescriptorCache) }

var fileDescriptorCache = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xd1, 0x6a, 0xc2, 0x30,
	0x18, 0x85, 0x49, 0xab, 0xdd, 0xfc, 0xad, 0x63, 0x64, 0x20, 0xa1, 0x83, 0x52, 0xbc, 0xea, 0x55,
	0x05, 0xc7, 0x1e, 0x60, 0xda, 0xe1, 0xcd, 0x84, 0xd1, 0x31, 0xef, 0x6b, 0x1a, 0x6b, 0xc1, 0x9a,
	0x10, 0xd3, 0xb1, 0xbe, 0xd8, 0x9e, 0x61, 0x97, 0x7b, 0x84, 0xd1, 0x27, 0x19, 0x4d, 0xbb, 0xa2,
	0x17, 0xde, 0xe5, 0x9c, 0xff, 0x9c, 0x43, 0xf8, 0x60, 0x48, 0x63, 0xba, 0x63, 0x81, 0x90, 0x5c,
	0x71, 0xdc, 0xd7, 0xc2, 0xb9, 0x4f, 0x39, 0x4f, 0xf7, 0x6c, 0xaa, 0xcd, 0x4d, 0xb1, 0x9d, 0xb2,
	0x5c, 0xa8, 0xb2, 0xc9, 0x4c, 0x42, 0xb0, 0x97, 0x52, 0xd0, 0xf5, 0x6a, 0xc1, 0x0f, 0xdb, 0x2c,
	0xc5, 0x18, 0x7a, 0x61, 0xac, 0x62, 0x82, 0x3c, 0xe4, 0xdb, 0x91, 0x7e, 0x63, 0x0f, 0x86, 0x4f,
	0x29, 0x3b, 0xa8, 0x26, 0x42, 0x0c, 0x7d, 0x3a, 0xb5, 0x26, 0x5f, 0x08, 0xac, 0x66, 0x06, 0xdf,
	0x80, 0x91, 0x25, 0xba, 0x3e, 0x88, 0x8c, 0x2c, 0xc1, 0x2e, 0xc0, 0xae, 0x14, 0x4c, 0x7e, 0x64,
	0x47, 0x2e, 0xdb, 0xee, 0x89, 0x83, 0x1d, 0xb8, 0x16, 0x92, 0x7f, 0x96, 0xaf, 0x59, 0x42, 0x4c,
	0x0f, 0xf9, 0x66, 0xd4, 0xe9, 0xee, 0xf6, 0x1e, 0xbd, 0x90, 0x9e, 0x5e, 0xec, 0x34, 0xbe, 0x05,
	0x93, 0x8a, 0x82, 0xf4, 0x3d, 0xe4, 0x8f, 0xa2, 0xfa, 0x89, 0xc7, 0x60, 0xe5, 0x2c, 0xe7, 0xb2,
	0x24, 0x96, 0x36, 0x5b, 0x55, 0xaf, 0x50, 0x51, 0x84, 0x6c, 0xaf, 0x62, 0x72, 0xa5, 0x2f, 0x9d,
	0x9e, 0x95, 0x60, 0x2f, 0x6a, 0x48, 0x6f, 0xf5, 0x77, 0x28, 0xc3, 0x8f, 0x60, 0xb5, 0x20, 0xc6,
	0x41, 0x83, 0x2d, 0xf8, 0xc7, 0x16, 0x3c, 0xd7, 0xd8, 0x9c, 0xbb, 0xa0, 0x41, 0x7c, 0x46, 0x6d,
	0x06, 0x83, 0x25, 0x53, 0xf3, 0xf8, 0xc8, 0xd6, 0xab, 0x8b, 0xcd, 0xd1, 0x59, 0x73, 0x6e, 0x7f,
	0x57, 0x2e, 0xfa, 0xa9, 0x5c, 0xf4, 0x5b, 0xb9, 0x68, 0x63, 0xe9, 0xf0, 0xc3, 0x5f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb5, 0xef, 0x80, 0xc1, 0xc1, 0x01, 0x00, 0x00,
}