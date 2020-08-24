// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lease.proto

/*
	Package leasepb is a generated protocol buffer package.

	It is generated from these files:
		lease.proto

	It has these top-level messages:
		Lease
		LeaseInternalRequest
		LeaseInternalResponse
*/
package leasepb

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	etcdserverpb "github.com/hanjm/etcd/etcdserver/etcdserverpb"

	io "io"
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

type Lease struct {
	ID           int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TTL          int64 `protobuf:"varint,2,opt,name=TTL,proto3" json:"TTL,omitempty"`
	RemainingTTL int64 `protobuf:"varint,3,opt,name=RemainingTTL,proto3" json:"RemainingTTL,omitempty"`
}

func (m *Lease) Reset()                    { *m = Lease{} }
func (m *Lease) String() string            { return proto.CompactTextString(m) }
func (*Lease) ProtoMessage()               {}
func (*Lease) Descriptor() ([]byte, []int) { return fileDescriptorLease, []int{0} }

type LeaseInternalRequest struct {
	LeaseTimeToLiveRequest *etcdserverpb.LeaseTimeToLiveRequest `protobuf:"bytes,1,opt,name=LeaseTimeToLiveRequest" json:"LeaseTimeToLiveRequest,omitempty"`
}

func (m *LeaseInternalRequest) Reset()                    { *m = LeaseInternalRequest{} }
func (m *LeaseInternalRequest) String() string            { return proto.CompactTextString(m) }
func (*LeaseInternalRequest) ProtoMessage()               {}
func (*LeaseInternalRequest) Descriptor() ([]byte, []int) { return fileDescriptorLease, []int{1} }

type LeaseInternalResponse struct {
	LeaseTimeToLiveResponse *etcdserverpb.LeaseTimeToLiveResponse `protobuf:"bytes,1,opt,name=LeaseTimeToLiveResponse" json:"LeaseTimeToLiveResponse,omitempty"`
}

func (m *LeaseInternalResponse) Reset()                    { *m = LeaseInternalResponse{} }
func (m *LeaseInternalResponse) String() string            { return proto.CompactTextString(m) }
func (*LeaseInternalResponse) ProtoMessage()               {}
func (*LeaseInternalResponse) Descriptor() ([]byte, []int) { return fileDescriptorLease, []int{2} }

func init() {
	proto.RegisterType((*Lease)(nil), "leasepb.Lease")
	proto.RegisterType((*LeaseInternalRequest)(nil), "leasepb.LeaseInternalRequest")
	proto.RegisterType((*LeaseInternalResponse)(nil), "leasepb.LeaseInternalResponse")
}
func (m *Lease) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lease) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLease(dAtA, i, uint64(m.ID))
	}
	if m.TTL != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLease(dAtA, i, uint64(m.TTL))
	}
	if m.RemainingTTL != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLease(dAtA, i, uint64(m.RemainingTTL))
	}
	return i, nil
}

func (m *LeaseInternalRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseInternalRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.LeaseTimeToLiveRequest != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLease(dAtA, i, uint64(m.LeaseTimeToLiveRequest.Size()))
		n1, err := m.LeaseTimeToLiveRequest.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *LeaseInternalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseInternalResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.LeaseTimeToLiveResponse != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLease(dAtA, i, uint64(m.LeaseTimeToLiveResponse.Size()))
		n2, err := m.LeaseTimeToLiveResponse.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeVarintLease(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Lease) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovLease(uint64(m.ID))
	}
	if m.TTL != 0 {
		n += 1 + sovLease(uint64(m.TTL))
	}
	if m.RemainingTTL != 0 {
		n += 1 + sovLease(uint64(m.RemainingTTL))
	}
	return n
}

func (m *LeaseInternalRequest) Size() (n int) {
	var l int
	_ = l
	if m.LeaseTimeToLiveRequest != nil {
		l = m.LeaseTimeToLiveRequest.Size()
		n += 1 + l + sovLease(uint64(l))
	}
	return n
}

func (m *LeaseInternalResponse) Size() (n int) {
	var l int
	_ = l
	if m.LeaseTimeToLiveResponse != nil {
		l = m.LeaseTimeToLiveResponse.Size()
		n += 1 + l + sovLease(uint64(l))
	}
	return n
}

func sovLease(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLease(x uint64) (n int) {
	return sovLease(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Lease) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLease
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
			return fmt.Errorf("proto: Lease: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lease: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TTL |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingTTL", wireType)
			}
			m.RemainingTTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RemainingTTL |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLease(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLease
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
func (m *LeaseInternalRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLease
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
			return fmt.Errorf("proto: LeaseInternalRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseInternalRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseTimeToLiveRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LeaseTimeToLiveRequest == nil {
				m.LeaseTimeToLiveRequest = &etcdserverpb.LeaseTimeToLiveRequest{}
			}
			if err := m.LeaseTimeToLiveRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLease(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLease
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
func (m *LeaseInternalResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLease
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
			return fmt.Errorf("proto: LeaseInternalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseInternalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseTimeToLiveResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LeaseTimeToLiveResponse == nil {
				m.LeaseTimeToLiveResponse = &etcdserverpb.LeaseTimeToLiveResponse{}
			}
			if err := m.LeaseTimeToLiveResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLease(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLease
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
func skipLease(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLease
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
					return 0, ErrIntOverflowLease
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
					return 0, ErrIntOverflowLease
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
				return 0, ErrInvalidLengthLease
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLease
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
				next, err := skipLease(dAtA[start:])
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
	ErrInvalidLengthLease = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLease   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("lease.proto", fileDescriptorLease) }

var fileDescriptorLease = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x49, 0x4d, 0x2c,
	0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x07, 0x73, 0x0a, 0x92, 0xa4, 0x44, 0xd2,
	0xf3, 0xd3, 0xf3, 0xc1, 0x62, 0xfa, 0x20, 0x16, 0x44, 0x5a, 0x4a, 0x2d, 0xb5, 0x24, 0x39, 0x45,
	0x1f, 0x44, 0x14, 0xa7, 0x16, 0x95, 0xa5, 0x16, 0x21, 0x31, 0x0b, 0x92, 0xf4, 0x8b, 0x0a, 0x92,
	0x21, 0xea, 0x94, 0x7c, 0xb9, 0x58, 0x7d, 0x40, 0x06, 0x09, 0xf1, 0x71, 0x31, 0x79, 0xba, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x31, 0x79, 0xba, 0x08, 0x09, 0x70, 0x31, 0x87, 0x84, 0xf8,
	0x48, 0x30, 0x81, 0x05, 0x40, 0x4c, 0x21, 0x25, 0x2e, 0x9e, 0xa0, 0xd4, 0xdc, 0xc4, 0xcc, 0xbc,
	0xcc, 0xbc, 0x74, 0x90, 0x14, 0x33, 0x58, 0x0a, 0x45, 0x4c, 0xa9, 0x84, 0x4b, 0x04, 0x6c, 0x9c,
	0x67, 0x5e, 0x49, 0x6a, 0x51, 0x5e, 0x62, 0x4e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x50,
	0x0c, 0x97, 0x18, 0x58, 0x3c, 0x24, 0x33, 0x37, 0x35, 0x24, 0xdf, 0x27, 0xb3, 0x2c, 0x15, 0x2a,
	0x03, 0xb6, 0x91, 0xdb, 0x48, 0x45, 0x0f, 0xd9, 0x7d, 0x7a, 0xd8, 0xd5, 0x06, 0xe1, 0x30, 0x43,
	0xa9, 0x82, 0x4b, 0x14, 0xcd, 0xd6, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa1, 0x78, 0x2e, 0x71,
	0x0c, 0x2d, 0x10, 0x29, 0xa8, 0xbd, 0xaa, 0x04, 0xec, 0x85, 0x28, 0x0e, 0xc2, 0x65, 0x8a, 0x93,
	0xc4, 0x89, 0x87, 0x72, 0x0c, 0x17, 0x1e, 0xca, 0x31, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81, 0xc3, 0xd7, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x9f, 0x8b, 0x6c, 0xb5, 0x01, 0x00, 0x00,
}
