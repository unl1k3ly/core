// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/wasm/v1beta1/stargatelayer/feegrant/query.proto

package feegrant

import (
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	feegrant "github.com/cosmos/cosmos-sdk/x/feegrant"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryAllowanceResponse is the response type for the Query/Allowance RPC method.
type QueryAllowanceResponse struct {
	// allowance is a allowance granted for grantee by granter.
	Allowance *feegrant.Grant `protobuf:"bytes,1,opt,name=allowance,proto3" json:"allowance,omitempty"`
}

func (m *QueryAllowanceResponse) Reset()         { *m = QueryAllowanceResponse{} }
func (m *QueryAllowanceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllowanceResponse) ProtoMessage()    {}
func (*QueryAllowanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4a9a739b0394bbd, []int{0}
}
func (m *QueryAllowanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllowanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllowanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllowanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllowanceResponse.Merge(m, src)
}
func (m *QueryAllowanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllowanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllowanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllowanceResponse proto.InternalMessageInfo

func (m *QueryAllowanceResponse) GetAllowance() *feegrant.Grant {
	if m != nil {
		return m.Allowance
	}
	return nil
}

// QueryAllowancesResponse is the response type for the Query/Allowances RPC method.
type QueryAllowancesResponse struct {
	// allowances are allowance's granted for grantee by granter.
	Allowances []*feegrant.Grant `protobuf:"bytes,1,rep,name=allowances,proto3" json:"allowances,omitempty"`
	// pagination defines an pagination for the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllowancesResponse) Reset()         { *m = QueryAllowancesResponse{} }
func (m *QueryAllowancesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllowancesResponse) ProtoMessage()    {}
func (*QueryAllowancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4a9a739b0394bbd, []int{1}
}
func (m *QueryAllowancesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllowancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllowancesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllowancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllowancesResponse.Merge(m, src)
}
func (m *QueryAllowancesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllowancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllowancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllowancesResponse proto.InternalMessageInfo

func (m *QueryAllowancesResponse) GetAllowances() []*feegrant.Grant {
	if m != nil {
		return m.Allowances
	}
	return nil
}

func (m *QueryAllowancesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAllowanceResponse)(nil), "terra.wasm.v1beta1.stargatelayer.feegrant.QueryAllowanceResponse")
	proto.RegisterType((*QueryAllowancesResponse)(nil), "terra.wasm.v1beta1.stargatelayer.feegrant.QueryAllowancesResponse")
}

func init() {
	proto.RegisterFile("terra/wasm/v1beta1/stargatelayer/feegrant/query.proto", fileDescriptor_e4a9a739b0394bbd)
}

var fileDescriptor_e4a9a739b0394bbd = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbd, 0x4a, 0x03, 0x41,
	0x10, 0xc7, 0xb3, 0x0a, 0x82, 0x6b, 0x97, 0x42, 0x43, 0x90, 0x25, 0xa4, 0xf0, 0x0b, 0xdc, 0x25,
	0x8a, 0x85, 0x20, 0x82, 0x36, 0x69, 0x35, 0x82, 0x85, 0xdd, 0xdc, 0x31, 0xae, 0x07, 0x77, 0x3b,
	0xe7, 0xee, 0xc6, 0x98, 0xb7, 0xf0, 0x19, 0x7c, 0x1a, 0xcb, 0x94, 0x96, 0x92, 0xbc, 0x88, 0xdc,
	0x5d, 0x6e, 0x2f, 0x01, 0x0b, 0xcb, 0x9d, 0xf9, 0xff, 0xfe, 0xf3, 0xb5, 0xfc, 0xc2, 0xa3, 0xb5,
	0xa0, 0x26, 0xe0, 0x32, 0xf5, 0x36, 0x88, 0xd0, 0xc3, 0x40, 0x39, 0x0f, 0x56, 0x83, 0xc7, 0x14,
	0xa6, 0x68, 0xd5, 0x33, 0xa2, 0xb6, 0x60, 0xbc, 0x7a, 0x1d, 0xa3, 0x9d, 0xca, 0xdc, 0x92, 0xa7,
	0xf6, 0x71, 0x89, 0xc9, 0x02, 0x93, 0x4b, 0x4c, 0xae, 0x61, 0xb2, 0xc6, 0xba, 0x07, 0x31, 0xb9,
	0x8c, 0x5c, 0xe3, 0x53, 0x97, 0xa9, 0x03, 0x95, 0x65, 0xf7, 0x64, 0xa9, 0x8b, 0xc0, 0x61, 0x55,
	0x2b, 0x28, 0x73, 0xd0, 0x89, 0x01, 0x9f, 0x90, 0x59, 0x6a, 0xf7, 0x35, 0x91, 0x4e, 0x51, 0x41,
	0x9e, 0x28, 0x30, 0x86, 0x7c, 0x99, 0x74, 0x55, 0xb6, 0xff, 0xc8, 0x77, 0xef, 0x0b, 0xfe, 0x26,
	0x4d, 0x69, 0x02, 0x26, 0xc6, 0x11, 0xba, 0x9c, 0x8c, 0xc3, 0xf6, 0x15, 0xdf, 0x86, 0x3a, 0xd8,
	0x61, 0x3d, 0x76, 0xb4, 0x73, 0x26, 0x64, 0x55, 0x37, 0x34, 0x1c, 0xe6, 0x19, 0x16, 0xaf, 0x51,
	0x03, 0xf4, 0x3f, 0x19, 0xdf, 0x5b, 0x37, 0x76, 0xc1, 0xf9, 0x9a, 0xf3, 0x20, 0x74, 0x1d, 0xd6,
	0xdb, 0xfc, 0x87, 0xf5, 0x0a, 0xd1, 0x1e, 0x72, 0xde, 0x4c, 0xd9, 0xd9, 0x28, 0x5b, 0x3b, 0xac,
	0xf9, 0x62, 0x25, 0xb2, 0x5a, 0x7f, 0xed, 0x70, 0x07, 0x3a, 0x8c, 0x35, 0x5a, 0x41, 0x6f, 0x1f,
	0xbe, 0xe6, 0x82, 0xcd, 0xe6, 0x82, 0xfd, 0xcc, 0x05, 0xfb, 0x58, 0x88, 0xd6, 0x6c, 0x21, 0x5a,
	0xdf, 0x0b, 0xd1, 0x7a, 0xba, 0xd4, 0x89, 0x7f, 0x19, 0x47, 0x32, 0xa6, 0x4c, 0x95, 0xe7, 0x3b,
	0xcd, 0xc8, 0xe0, 0x54, 0xc5, 0x64, 0x51, 0xbd, 0x57, 0x5f, 0xe0, 0xef, 0xd3, 0x47, 0x5b, 0xe5,
	0x62, 0xcf, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x78, 0x38, 0x6e, 0x13, 0x2e, 0x02, 0x00, 0x00,
}

func (m *QueryAllowanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllowanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllowanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Allowance != nil {
		{
			size, err := m.Allowance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllowancesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllowancesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllowancesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Allowances) > 0 {
		for iNdEx := len(m.Allowances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Allowances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAllowanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Allowance != nil {
		l = m.Allowance.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllowancesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Allowances) > 0 {
		for _, e := range m.Allowances {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAllowanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllowanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllowanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Allowance == nil {
				m.Allowance = &feegrant.Grant{}
			}
			if err := m.Allowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllowancesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllowancesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllowancesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Allowances = append(m.Allowances, &feegrant.Grant{})
			if err := m.Allowances[len(m.Allowances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)