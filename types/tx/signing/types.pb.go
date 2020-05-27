// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/tx/signing/types.proto

package signing

import (
	fmt "fmt"
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

// TxRaw is using for transaction decoding and verification
type TxRaw struct {
	BodyBytes     []byte   `protobuf:"bytes,1,opt,name=body_bytes,json=bodyBytes,proto3" json:"body_bytes,omitempty"`
	AuthInfoBytes []byte   `protobuf:"bytes,2,opt,name=auth_info_bytes,json=authInfoBytes,proto3" json:"auth_info_bytes,omitempty"`
	Signatures    [][]byte `protobuf:"bytes,3,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (m *TxRaw) Reset()         { *m = TxRaw{} }
func (m *TxRaw) String() string { return proto.CompactTextString(m) }
func (*TxRaw) ProtoMessage()    {}
func (*TxRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_52d5d5e2efced9ba, []int{0}
}
func (m *TxRaw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxRaw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxRaw.Merge(m, src)
}
func (m *TxRaw) XXX_Size() int {
	return m.Size()
}
func (m *TxRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_TxRaw.DiscardUnknown(m)
}

var xxx_messageInfo_TxRaw proto.InternalMessageInfo

func (m *TxRaw) GetBodyBytes() []byte {
	if m != nil {
		return m.BodyBytes
	}
	return nil
}

func (m *TxRaw) GetAuthInfoBytes() []byte {
	if m != nil {
		return m.AuthInfoBytes
	}
	return nil
}

func (m *TxRaw) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// SignDocRaw is used for verifying signatures with SIGN_MODE_DIRECT
type SignDocRaw struct {
	BodyBytes       []byte `protobuf:"bytes,1,opt,name=body_bytes,json=bodyBytes,proto3" json:"body_bytes,omitempty"`
	AuthInfoBytes   []byte `protobuf:"bytes,2,opt,name=auth_info_bytes,json=authInfoBytes,proto3" json:"auth_info_bytes,omitempty"`
	ChainId         string `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	AccountNumber   uint64 `protobuf:"varint,4,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	AccountSequence uint64 `protobuf:"varint,5,opt,name=account_sequence,json=accountSequence,proto3" json:"account_sequence,omitempty"`
}

func (m *SignDocRaw) Reset()         { *m = SignDocRaw{} }
func (m *SignDocRaw) String() string { return proto.CompactTextString(m) }
func (*SignDocRaw) ProtoMessage()    {}
func (*SignDocRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_52d5d5e2efced9ba, []int{1}
}
func (m *SignDocRaw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignDocRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignDocRaw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignDocRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignDocRaw.Merge(m, src)
}
func (m *SignDocRaw) XXX_Size() int {
	return m.Size()
}
func (m *SignDocRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_SignDocRaw.DiscardUnknown(m)
}

var xxx_messageInfo_SignDocRaw proto.InternalMessageInfo

func (m *SignDocRaw) GetBodyBytes() []byte {
	if m != nil {
		return m.BodyBytes
	}
	return nil
}

func (m *SignDocRaw) GetAuthInfoBytes() []byte {
	if m != nil {
		return m.AuthInfoBytes
	}
	return nil
}

func (m *SignDocRaw) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *SignDocRaw) GetAccountNumber() uint64 {
	if m != nil {
		return m.AccountNumber
	}
	return 0
}

func (m *SignDocRaw) GetAccountSequence() uint64 {
	if m != nil {
		return m.AccountSequence
	}
	return 0
}

func init() {
	proto.RegisterType((*TxRaw)(nil), "cosmos_sdk.tx.signing.v1.TxRaw")
	proto.RegisterType((*SignDocRaw)(nil), "cosmos_sdk.tx.signing.v1.SignDocRaw")
}

func init() { proto.RegisterFile("types/tx/signing/types.proto", fileDescriptor_52d5d5e2efced9ba) }

var fileDescriptor_52d5d5e2efced9ba = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd1, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc0, 0xf1, 0xc5, 0x6d, 0xea, 0x3e, 0x36, 0x27, 0x3d, 0x45, 0xd0, 0x50, 0x06, 0x4a, 0x3d,
	0xac, 0x45, 0x7c, 0x83, 0x21, 0xc8, 0x2e, 0x1e, 0x3a, 0x4f, 0x5e, 0x4a, 0x9b, 0x66, 0x5d, 0x18,
	0x4b, 0xe6, 0x92, 0xe8, 0xfa, 0x16, 0x3e, 0x92, 0x47, 0x8f, 0x3b, 0x7a, 0x94, 0xf6, 0x45, 0xa4,
	0x59, 0x04, 0xf1, 0xec, 0x29, 0xe4, 0xc7, 0x9f, 0x04, 0xbe, 0x0f, 0xce, 0x75, 0xb9, 0x66, 0x2a,
	0xd2, 0xdb, 0x48, 0xf1, 0x42, 0x70, 0x51, 0x44, 0x16, 0xc2, 0xf5, 0x46, 0x6a, 0xe9, 0x61, 0x2a,
	0xd5, 0x4a, 0xaa, 0x44, 0xe5, 0xcb, 0x50, 0x6f, 0x43, 0x97, 0x84, 0x2f, 0x37, 0x23, 0x01, 0xdd,
	0xc7, 0x6d, 0x9c, 0xbe, 0x7a, 0x17, 0x00, 0x99, 0xcc, 0xcb, 0x24, 0x2b, 0x35, 0x53, 0x18, 0xf9,
	0x28, 0xe8, 0xc7, 0xbd, 0x46, 0x26, 0x0d, 0x78, 0x57, 0x30, 0x4c, 0x8d, 0x5e, 0x24, 0x5c, 0xcc,
	0xa5, 0x6b, 0x0e, 0x6c, 0x33, 0x68, 0x78, 0x2a, 0xe6, 0x72, 0xdf, 0x11, 0x80, 0xe6, 0xf5, 0x54,
	0x9b, 0x0d, 0x53, 0xb8, 0xed, 0xb7, 0x83, 0x7e, 0xfc, 0x4b, 0x46, 0xef, 0x08, 0x60, 0xc6, 0x0b,
	0x71, 0x27, 0xe9, 0x3f, 0xfe, 0x7a, 0x06, 0xc7, 0x74, 0x91, 0x72, 0x91, 0xf0, 0x1c, 0xb7, 0x7d,
	0x14, 0xf4, 0xe2, 0x23, 0x7b, 0x9f, 0xe6, 0xde, 0x25, 0x9c, 0xa4, 0x94, 0x4a, 0x23, 0x74, 0x22,
	0xcc, 0x2a, 0x63, 0x1b, 0xdc, 0xf1, 0x51, 0xd0, 0x89, 0x07, 0x4e, 0x1f, 0x2c, 0x7a, 0xd7, 0x70,
	0xfa, 0x93, 0x29, 0xf6, 0x6c, 0x98, 0xa0, 0x0c, 0x77, 0x6d, 0x38, 0x74, 0x3e, 0x73, 0x3c, 0xb9,
	0xff, 0xa8, 0x08, 0xda, 0x55, 0x04, 0x7d, 0x55, 0x04, 0xbd, 0xd5, 0xa4, 0xb5, 0xab, 0x49, 0xeb,
	0xb3, 0x26, 0xad, 0xa7, 0x71, 0xc1, 0xf5, 0xc2, 0x64, 0x21, 0x95, 0xab, 0x68, 0x3f, 0x71, 0x77,
	0x8c, 0x55, 0xbe, 0x8c, 0xfe, 0x6e, 0x28, 0x3b, 0xb4, 0xcb, 0xb9, 0xfd, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x34, 0x9a, 0xbd, 0x5a, 0xbc, 0x01, 0x00, 0x00,
}

func (m *TxRaw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxRaw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxRaw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signatures) > 0 {
		for iNdEx := len(m.Signatures) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signatures[iNdEx])
			copy(dAtA[i:], m.Signatures[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Signatures[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AuthInfoBytes) > 0 {
		i -= len(m.AuthInfoBytes)
		copy(dAtA[i:], m.AuthInfoBytes)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.AuthInfoBytes)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BodyBytes) > 0 {
		i -= len(m.BodyBytes)
		copy(dAtA[i:], m.BodyBytes)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BodyBytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SignDocRaw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignDocRaw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignDocRaw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AccountSequence != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.AccountSequence))
		i--
		dAtA[i] = 0x28
	}
	if m.AccountNumber != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.AccountNumber))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AuthInfoBytes) > 0 {
		i -= len(m.AuthInfoBytes)
		copy(dAtA[i:], m.AuthInfoBytes)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.AuthInfoBytes)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BodyBytes) > 0 {
		i -= len(m.BodyBytes)
		copy(dAtA[i:], m.BodyBytes)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BodyBytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxRaw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BodyBytes)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.AuthInfoBytes)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Signatures) > 0 {
		for _, b := range m.Signatures {
			l = len(b)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *SignDocRaw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BodyBytes)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.AuthInfoBytes)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.AccountNumber != 0 {
		n += 1 + sovTypes(uint64(m.AccountNumber))
	}
	if m.AccountSequence != 0 {
		n += 1 + sovTypes(uint64(m.AccountSequence))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TxRaw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: TxRaw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxRaw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BodyBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BodyBytes = append(m.BodyBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.BodyBytes == nil {
				m.BodyBytes = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthInfoBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthInfoBytes = append(m.AuthInfoBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.AuthInfoBytes == nil {
				m.AuthInfoBytes = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, make([]byte, postIndex-iNdEx))
			copy(m.Signatures[len(m.Signatures)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *SignDocRaw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: SignDocRaw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignDocRaw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BodyBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BodyBytes = append(m.BodyBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.BodyBytes == nil {
				m.BodyBytes = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthInfoBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthInfoBytes = append(m.AuthInfoBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.AuthInfoBytes == nil {
				m.AuthInfoBytes = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountNumber", wireType)
			}
			m.AccountNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccountNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountSequence", wireType)
			}
			m.AccountSequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccountSequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)