// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/wasm/v1beta1/stargatelayer/market/query.proto

package market

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types1 "github.com/terra-money/core/x/market/types"
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

// QuerySwapResponse is the response type for the Query/Swap RPC method.
type QuerySwapResponse struct {
	// return_coin defines the coin returned as a result of the swap simulation.
	ReturnCoin types.Coin `protobuf:"bytes,1,opt,name=return_coin,json=returnCoin,proto3" json:"return_coin"`
}

func (m *QuerySwapResponse) Reset()         { *m = QuerySwapResponse{} }
func (m *QuerySwapResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySwapResponse) ProtoMessage()    {}
func (*QuerySwapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e802064116374db0, []int{0}
}
func (m *QuerySwapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySwapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySwapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySwapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySwapResponse.Merge(m, src)
}
func (m *QuerySwapResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySwapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySwapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySwapResponse proto.InternalMessageInfo

func (m *QuerySwapResponse) GetReturnCoin() types.Coin {
	if m != nil {
		return m.ReturnCoin
	}
	return types.Coin{}
}

// QueryTerraPoolDeltaResponse is the response type for the Query/TerraPoolDelta RPC method.
type QueryTerraPoolDeltaResponse struct {
	// terra_pool_delta defines the gap between the TerraPool and the TerraBasePool
	TerraPoolDelta github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=terra_pool_delta,json=terraPoolDelta,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"terra_pool_delta"`
}

func (m *QueryTerraPoolDeltaResponse) Reset()         { *m = QueryTerraPoolDeltaResponse{} }
func (m *QueryTerraPoolDeltaResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTerraPoolDeltaResponse) ProtoMessage()    {}
func (*QueryTerraPoolDeltaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e802064116374db0, []int{1}
}
func (m *QueryTerraPoolDeltaResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTerraPoolDeltaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTerraPoolDeltaResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTerraPoolDeltaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTerraPoolDeltaResponse.Merge(m, src)
}
func (m *QueryTerraPoolDeltaResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTerraPoolDeltaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTerraPoolDeltaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTerraPoolDeltaResponse proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params types1.Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e802064116374db0, []int{2}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() types1.Params {
	if m != nil {
		return m.Params
	}
	return types1.Params{}
}

func init() {
	proto.RegisterType((*QuerySwapResponse)(nil), "terra.wasm.v1beta1.stargatelayer.market.QuerySwapResponse")
	proto.RegisterType((*QueryTerraPoolDeltaResponse)(nil), "terra.wasm.v1beta1.stargatelayer.market.QueryTerraPoolDeltaResponse")
	proto.RegisterType((*QueryParamsResponse)(nil), "terra.wasm.v1beta1.stargatelayer.market.QueryParamsResponse")
}

func init() {
	proto.RegisterFile("terra/wasm/v1beta1/stargatelayer/market/query.proto", fileDescriptor_e802064116374db0)
}

var fileDescriptor_e802064116374db0 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0x6f, 0xda, 0x30,
	0x18, 0xc7, 0x13, 0x69, 0xe2, 0x60, 0xa6, 0x69, 0x63, 0x3b, 0x6c, 0x0c, 0x85, 0x8d, 0xc3, 0xb6,
	0x0b, 0xb6, 0x18, 0x87, 0x49, 0x3b, 0x4d, 0x8c, 0x0f, 0x00, 0xb4, 0x95, 0xaa, 0x5e, 0x90, 0x13,
	0x1e, 0xa5, 0x11, 0x89, 0x9f, 0xd4, 0x36, 0xa5, 0xf9, 0x16, 0xfd, 0x58, 0x1c, 0x39, 0x56, 0x3d,
	0xa0, 0x0a, 0xbe, 0x48, 0xe5, 0x17, 0xa2, 0x22, 0xf5, 0x94, 0xc4, 0x7e, 0x7e, 0xff, 0x17, 0xc7,
	0x64, 0xa8, 0x41, 0x4a, 0xce, 0xd6, 0x5c, 0x15, 0xec, 0x76, 0x10, 0x83, 0xe6, 0x03, 0xa6, 0x34,
	0x97, 0x29, 0xd7, 0x90, 0xf3, 0x0a, 0x24, 0x2b, 0xb8, 0x5c, 0x82, 0x66, 0x37, 0x2b, 0x90, 0x15,
	0x2d, 0x25, 0x6a, 0x6c, 0xfd, 0xb4, 0x10, 0x35, 0x10, 0xf5, 0x10, 0x3d, 0x81, 0xa8, 0x83, 0xda,
	0x9f, 0x52, 0x4c, 0xd1, 0x32, 0xcc, 0xbc, 0x39, 0xbc, 0xdd, 0x49, 0x11, 0xd3, 0x1c, 0x18, 0x2f,
	0x33, 0xc6, 0x85, 0x40, 0xcd, 0x75, 0x86, 0x42, 0xf9, 0xdd, 0xef, 0x2e, 0x91, 0xb7, 0x3d, 0x66,
	0x72, 0x9f, 0x7e, 0x24, 0x4a, 0x50, 0x15, 0xa8, 0x58, 0xcc, 0x15, 0xd4, 0x13, 0x09, 0x66, 0xc2,
	0xed, 0xf7, 0x2e, 0xc8, 0x87, 0xa9, 0x89, 0x7b, 0xb6, 0xe6, 0xe5, 0x0c, 0x54, 0x89, 0x42, 0x41,
	0xeb, 0x1f, 0x69, 0x4a, 0xd0, 0x2b, 0x29, 0xe6, 0x66, 0xf2, 0x73, 0xf8, 0x2d, 0xfc, 0xd5, 0xfc,
	0xfd, 0x85, 0x3a, 0x29, 0x6a, 0xa4, 0xea, 0x2e, 0xff, 0x31, 0x13, 0xa3, 0x37, 0x9b, 0x5d, 0x37,
	0x98, 0x11, 0xc7, 0x98, 0x95, 0xde, 0x9a, 0x7c, 0xb5, 0xb2, 0xe7, 0x26, 0xe0, 0x04, 0x31, 0x1f,
	0x43, 0xae, 0x79, 0x6d, 0x70, 0x49, 0xde, 0xdb, 0xe8, 0xf3, 0x12, 0x31, 0x9f, 0x2f, 0xcc, 0x9e,
	0x75, 0x79, 0x3b, 0xa2, 0x46, 0xea, 0x71, 0xd7, 0xfd, 0x91, 0x66, 0xfa, 0x7a, 0x15, 0xd3, 0x04,
	0x0b, 0xe6, 0x2b, 0xb8, 0x47, 0x5f, 0x2d, 0x96, 0x4c, 0x57, 0x25, 0x28, 0x3a, 0x86, 0x64, 0xf6,
	0x4e, 0x9f, 0x38, 0xf4, 0xa6, 0xe4, 0xa3, 0x35, 0x9e, 0x70, 0xc9, 0x0b, 0x55, 0x1b, 0xfe, 0x25,
	0x8d, 0xd2, 0xae, 0xf8, 0x32, 0x1d, 0xea, 0xfe, 0x8b, 0x3f, 0xab, 0x63, 0x1b, 0x47, 0xf9, 0x3e,
	0x9e, 0x18, 0x4d, 0x37, 0xfb, 0x28, 0xdc, 0xee, 0xa3, 0xf0, 0x69, 0x1f, 0x85, 0xf7, 0x87, 0x28,
	0xd8, 0x1e, 0xa2, 0xe0, 0xe1, 0x10, 0x05, 0x57, 0x7f, 0x5e, 0x84, 0xb4, 0x7a, 0xfd, 0x02, 0x05,
	0x54, 0x2c, 0x41, 0x09, 0xec, 0xce, 0xdd, 0x94, 0xd7, 0x6e, 0x48, 0xdc, 0xb0, 0x87, 0x3f, 0x7c,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xa1, 0x8e, 0xcb, 0x53, 0x02, 0x00, 0x00,
}

func (m *QuerySwapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySwapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySwapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ReturnCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryTerraPoolDeltaResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTerraPoolDeltaResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTerraPoolDeltaResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TerraPoolDelta.Size()
		i -= size
		if _, err := m.TerraPoolDelta.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *QuerySwapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ReturnCoin.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryTerraPoolDeltaResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TerraPoolDelta.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QuerySwapResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySwapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySwapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReturnCoin", wireType)
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
			if err := m.ReturnCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryTerraPoolDeltaResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryTerraPoolDeltaResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTerraPoolDeltaResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TerraPoolDelta", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TerraPoolDelta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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