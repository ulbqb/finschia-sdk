// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/connection/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the ibc connection submodule's genesis state.
type GenesisState struct {
	Connections           []IdentifiedConnection `protobuf:"bytes,1,rep,name=connections,proto3" json:"connections"`
	ClientConnectionPaths []ConnectionPaths      `protobuf:"bytes,2,rep,name=client_connection_paths,json=clientConnectionPaths,proto3" json:"client_connection_paths" yaml:"client_connection_paths"`
	// the sequence for the next generated connection identifier
	NextConnectionSequence uint64 `protobuf:"varint,3,opt,name=next_connection_sequence,json=nextConnectionSequence,proto3" json:"next_connection_sequence,omitempty" yaml:"next_connection_sequence"`
	Params                 Params `protobuf:"bytes,4,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1879d34bc6ac3cd7, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetConnections() []IdentifiedConnection {
	if m != nil {
		return m.Connections
	}
	return nil
}

func (m *GenesisState) GetClientConnectionPaths() []ConnectionPaths {
	if m != nil {
		return m.ClientConnectionPaths
	}
	return nil
}

func (m *GenesisState) GetNextConnectionSequence() uint64 {
	if m != nil {
		return m.NextConnectionSequence
	}
	return 0
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ibc.core.connection.v1.GenesisState")
}

func init() {
	proto.RegisterFile("ibc/core/connection/v1/genesis.proto", fileDescriptor_1879d34bc6ac3cd7)
}

var fileDescriptor_1879d34bc6ac3cd7 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x31, 0x4f, 0xf2, 0x40,
	0x18, 0xc7, 0xdb, 0x17, 0xc2, 0x50, 0xde, 0xa9, 0x51, 0x6c, 0x18, 0xae, 0xa4, 0x1a, 0x61, 0x90,
	0x9e, 0x40, 0xe2, 0x60, 0x9c, 0xea, 0x60, 0xdc, 0x08, 0x38, 0x99, 0x18, 0xd2, 0x1e, 0x8f, 0xe5,
	0x62, 0x7b, 0x57, 0xb9, 0x83, 0xc0, 0x27, 0x70, 0x70, 0xf1, 0x63, 0x31, 0x32, 0x3a, 0x11, 0x03,
	0xdf, 0x80, 0x4f, 0x60, 0xda, 0x12, 0x8a, 0xc6, 0x6e, 0xcd, 0xf3, 0xfc, 0xfe, 0xbf, 0x7f, 0x7a,
	0x8f, 0x76, 0x46, 0x3d, 0x82, 0x09, 0x1f, 0x03, 0x26, 0x9c, 0x31, 0x20, 0x92, 0x72, 0x86, 0xa7,
	0x2d, 0xec, 0x03, 0x03, 0x41, 0x85, 0x1d, 0x8d, 0xb9, 0xe4, 0x7a, 0x85, 0x7a, 0xc4, 0x8e, 0x29,
	0x3b, 0xa3, 0xec, 0x69, 0xab, 0x7a, 0xe4, 0x73, 0x9f, 0x27, 0x08, 0x8e, 0xbf, 0x52, 0xba, 0x5a,
	0xcf, 0x71, 0x1e, 0x64, 0x13, 0xd0, 0x7a, 0x2f, 0x68, 0xff, 0xef, 0xd2, 0xa2, 0xbe, 0x74, 0x25,
	0xe8, 0x0f, 0x5a, 0x39, 0x83, 0x84, 0xa1, 0xd6, 0x0a, 0x8d, 0x72, 0xfb, 0xc2, 0xfe, 0xbb, 0xdd,
	0xbe, 0x1f, 0x02, 0x93, 0xf4, 0x99, 0xc2, 0xf0, 0x76, 0x3f, 0x77, 0x8a, 0x8b, 0x95, 0xa9, 0xf4,
	0x0e, 0x35, 0xfa, 0x9b, 0xaa, 0x9d, 0x90, 0x80, 0x02, 0x93, 0x83, 0x6c, 0x3c, 0x88, 0x5c, 0x39,
	0x12, 0xc6, 0xbf, 0xa4, 0xa2, 0x9e, 0x57, 0x91, 0x89, 0xbb, 0x31, 0xee, 0x9c, 0xc7, 0xf6, 0xed,
	0xca, 0x44, 0x73, 0x37, 0x0c, 0xae, 0xad, 0x1c, 0xab, 0xd5, 0x3b, 0x4e, 0x37, 0xbf, 0xe2, 0xfa,
	0x93, 0x66, 0x30, 0x98, 0xfd, 0x08, 0x08, 0x78, 0x9d, 0x00, 0x23, 0x60, 0x14, 0x6a, 0x6a, 0xa3,
	0xe8, 0x9c, 0x6e, 0x57, 0xa6, 0x99, 0xca, 0xf3, 0x48, 0xab, 0x57, 0x89, 0x57, 0x99, 0xbb, 0xbf,
	0x5b, 0xe8, 0x37, 0x5a, 0x29, 0x72, 0xc7, 0x6e, 0x28, 0x8c, 0x62, 0x4d, 0x6d, 0x94, 0xdb, 0x28,
	0xef, 0xb7, 0xba, 0x09, 0xb5, 0x7b, 0xab, 0x5d, 0xc6, 0xe9, 0x2e, 0xd6, 0x48, 0x5d, 0xae, 0x91,
	0xfa, 0xb5, 0x46, 0xea, 0xc7, 0x06, 0x29, 0xcb, 0x0d, 0x52, 0x3e, 0x37, 0x48, 0x79, 0xbc, 0xf2,
	0xa9, 0x1c, 0x4d, 0x3c, 0x9b, 0xf0, 0x10, 0x07, 0x94, 0x01, 0x0e, 0xbc, 0xb0, 0x29, 0x86, 0x2f,
	0x78, 0x86, 0xf7, 0xa7, 0xbe, 0xec, 0x34, 0x0f, 0xae, 0x2d, 0xe7, 0x11, 0x08, 0xaf, 0x94, 0x9c,
	0xb9, 0xf3, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x44, 0xd5, 0x6f, 0x29, 0x65, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.NextConnectionSequence != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextConnectionSequence))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ClientConnectionPaths) > 0 {
		for iNdEx := len(m.ClientConnectionPaths) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClientConnectionPaths[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Connections) > 0 {
		for iNdEx := len(m.Connections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Connections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Connections) > 0 {
		for _, e := range m.Connections {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClientConnectionPaths) > 0 {
		for _, e := range m.ClientConnectionPaths {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.NextConnectionSequence != 0 {
		n += 1 + sovGenesis(uint64(m.NextConnectionSequence))
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Connections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Connections = append(m.Connections, IdentifiedConnection{})
			if err := m.Connections[len(m.Connections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientConnectionPaths", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientConnectionPaths = append(m.ClientConnectionPaths, ConnectionPaths{})
			if err := m.ClientConnectionPaths[len(m.ClientConnectionPaths)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextConnectionSequence", wireType)
			}
			m.NextConnectionSequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextConnectionSequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
