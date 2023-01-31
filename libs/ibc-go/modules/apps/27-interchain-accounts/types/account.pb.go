// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/interchain_accounts/v1/account.proto

package types

//
//import (
//	fmt "fmt"
//	io "io"
//	math "math"
//	math_bits "math/bits"
//
//	types "github.com/FiboChain/fbc/libs/cosmos-sdk/x/auth/typesadapter"
//
//	_ "github.com/gogo/protobuf/gogoproto"
//	proto "github.com/gogo/protobuf/proto"
//	_ "github.com/regen-network/cosmos-proto"
//)
//
//// Reference imports to suppress errors if they are not otherwise used.
//var _ = proto.Marshal
//var _ = fmt.Errorf
//var _ = math.Inf
//
//// This is a compile-time assertion to ensure that this generated file
//// is compatible with the proto package it is being compiled against.
//// A compilation error at this line likely means your copy of the
//// proto package needs to be updated.
//const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package
//
//// An InterchainAccount is defined as a BaseAccount & the address of the account owner on the controller chain
//type InterchainAccount struct {
//	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty" yaml:"base_account"`
//	AccountOwner       string `protobuf:"bytes,2,opt,name=account_owner,json=accountOwner,proto3" json:"account_owner,omitempty" yaml:"account_owner"`
//}
//
//func (m *InterchainAccount) Reset()      { *m = InterchainAccount{} }
//func (*InterchainAccount) ProtoMessage() {}
//func (*InterchainAccount) Descriptor() ([]byte, []int) {
//	return fileDescriptor_5561bd92625bf7da, []int{0}
//}
//func (m *InterchainAccount) XXX_Unmarshal(b []byte) error {
//	return m.Unmarshal(b)
//}
//func (m *InterchainAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
//	if deterministic {
//		return xxx_messageInfo_InterchainAccount.Marshal(b, m, deterministic)
//	} else {
//		b = b[:cap(b)]
//		n, err := m.MarshalToSizedBuffer(b)
//		if err != nil {
//			return nil, err
//		}
//		return b[:n], nil
//	}
//}
//func (m *InterchainAccount) XXX_Merge(src proto.Message) {
//	xxx_messageInfo_InterchainAccount.Merge(m, src)
//}
//func (m *InterchainAccount) XXX_Size() int {
//	return m.Size()
//}
//func (m *InterchainAccount) XXX_DiscardUnknown() {
//	xxx_messageInfo_InterchainAccount.DiscardUnknown(m)
//}
//
//var xxx_messageInfo_InterchainAccount proto.InternalMessageInfo
//
//func init() {
//	proto.RegisterType((*InterchainAccount)(nil), "ibc.applications.interchain_accounts.v1.InterchainAccount")
//}
//
//func init() {
//	proto.RegisterFile("ibc/applications/interchain_accounts/v1/account.proto", fileDescriptor_5561bd92625bf7da)
//}
//
//var fileDescriptor_5561bd92625bf7da = []byte{
//	// 341 bytes of a gzipped FileDescriptorProto
//	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xcd, 0x4c, 0x4a, 0xd6,
//	0x4f, 0x2c, 0x28, 0xc8, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0xcf, 0xcc, 0x2b,
//	0x49, 0x2d, 0x4a, 0xce, 0x48, 0xcc, 0xcc, 0x8b, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x29,
//	0xd6, 0x2f, 0x33, 0xd4, 0x87, 0xb2, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd4, 0x33, 0x93,
//	0x92, 0xf5, 0x90, 0xb5, 0xe9, 0x61, 0xd1, 0xa6, 0x57, 0x66, 0x28, 0x25, 0x99, 0x9c, 0x5f, 0x9c,
//	0x9b, 0x5f, 0x1c, 0x0f, 0xd6, 0xa6, 0x0f, 0xe1, 0x40, 0xcc, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
//	0x87, 0x88, 0x83, 0x58, 0x50, 0x51, 0x39, 0x88, 0x1a, 0xfd, 0xc4, 0xd2, 0x92, 0x0c, 0xfd, 0x32,
//	0xc3, 0xa4, 0xd4, 0x92, 0x44, 0x43, 0x30, 0x07, 0x22, 0xaf, 0x74, 0x85, 0x91, 0x4b, 0xd0, 0x13,
//	0x6e, 0x97, 0x23, 0xc4, 0x2a, 0xa1, 0x04, 0x2e, 0x9e, 0xa4, 0xc4, 0xe2, 0x54, 0x98, 0xd5, 0x12,
//	0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x0a, 0x7a, 0x50, 0x0b, 0xc1, 0xfa, 0xa1, 0x86, 0xe9, 0x39,
//	0x25, 0x16, 0xa7, 0x42, 0xf5, 0x39, 0x49, 0x5f, 0xb8, 0x27, 0xcf, 0xf8, 0xe9, 0x9e, 0xbc, 0x70,
//	0x65, 0x62, 0x6e, 0x8e, 0x95, 0x12, 0xb2, 0x19, 0x4a, 0x41, 0xdc, 0x49, 0x08, 0x95, 0x42, 0xb6,
//	0x5c, 0xbc, 0x50, 0x89, 0xf8, 0xfc, 0xf2, 0xbc, 0xd4, 0x22, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x4e,
//	0x27, 0x89, 0x4f, 0xf7, 0xe4, 0x45, 0x20, 0x9a, 0x51, 0xa4, 0x95, 0x82, 0x78, 0xa0, 0x7c, 0x7f,
//	0x10, 0xd7, 0x4a, 0xae, 0x63, 0x81, 0x3c, 0xc3, 0x8c, 0x05, 0xf2, 0x0c, 0x97, 0xb6, 0xe8, 0x0a,
//	0x61, 0xb8, 0xdf, 0xd3, 0x29, 0xfe, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
//	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2,
//	0x5c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xa1, 0xe1, 0xa7, 0x9f, 0x99,
//	0x94, 0xac, 0x9b, 0x9e, 0xaf, 0x5f, 0x66, 0xa2, 0x9f, 0x9b, 0x9f, 0x52, 0x9a, 0x93, 0x5a, 0x0c,
//	0x8a, 0xc1, 0x62, 0x7d, 0x23, 0x73, 0x5d, 0x44, 0x2c, 0xe8, 0xc2, 0x23, 0xaf, 0xa4, 0xb2, 0x20,
//	0xb5, 0x38, 0x89, 0x0d, 0x1c, 0x7c, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xd7, 0x23,
//	0x15, 0xf1, 0x01, 0x00, 0x00,
//}
//
//func (m *InterchainAccount) Marshal() (dAtA []byte, err error) {
//	size := m.Size()
//	dAtA = make([]byte, size)
//	n, err := m.MarshalToSizedBuffer(dAtA[:size])
//	if err != nil {
//		return nil, err
//	}
//	return dAtA[:n], nil
//}
//
//func (m *InterchainAccount) MarshalTo(dAtA []byte) (int, error) {
//	size := m.Size()
//	return m.MarshalToSizedBuffer(dAtA[:size])
//}
//
//func (m *InterchainAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
//	i := len(dAtA)
//	_ = i
//	var l int
//	_ = l
//	if len(m.AccountOwner) > 0 {
//		i -= len(m.AccountOwner)
//		copy(dAtA[i:], m.AccountOwner)
//		i = encodeVarintAccount(dAtA, i, uint64(len(m.AccountOwner)))
//		i--
//		dAtA[i] = 0x12
//	}
//	if m.BaseAccount != nil {
//		{
//			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
//			if err != nil {
//				return 0, err
//			}
//			i -= size
//			i = encodeVarintAccount(dAtA, i, uint64(size))
//		}
//		i--
//		dAtA[i] = 0xa
//	}
//	return len(dAtA) - i, nil
//}
//
//func encodeVarintAccount(dAtA []byte, offset int, v uint64) int {
//	offset -= sovAccount(v)
//	base := offset
//	for v >= 1<<7 {
//		dAtA[offset] = uint8(v&0x7f | 0x80)
//		v >>= 7
//		offset++
//	}
//	dAtA[offset] = uint8(v)
//	return base
//}
//func (m *InterchainAccount) Size() (n int) {
//	if m == nil {
//		return 0
//	}
//	var l int
//	_ = l
//	if m.BaseAccount != nil {
//		l = m.BaseAccount.Size()
//		n += 1 + l + sovAccount(uint64(l))
//	}
//	l = len(m.AccountOwner)
//	if l > 0 {
//		n += 1 + l + sovAccount(uint64(l))
//	}
//	return n
//}
//
//func sovAccount(x uint64) (n int) {
//	return (math_bits.Len64(x|1) + 6) / 7
//}
//func sozAccount(x uint64) (n int) {
//	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
//}
//func (m *InterchainAccount) Unmarshal(dAtA []byte) error {
//	l := len(dAtA)
//	iNdEx := 0
//	for iNdEx < l {
//		preIndex := iNdEx
//		var wire uint64
//		for shift := uint(0); ; shift += 7 {
//			if shift >= 64 {
//				return ErrIntOverflowAccount
//			}
//			if iNdEx >= l {
//				return io.ErrUnexpectedEOF
//			}
//			b := dAtA[iNdEx]
//			iNdEx++
//			wire |= uint64(b&0x7F) << shift
//			if b < 0x80 {
//				break
//			}
//		}
//		fieldNum := int32(wire >> 3)
//		wireType := int(wire & 0x7)
//		if wireType == 4 {
//			return fmt.Errorf("proto: InterchainAccount: wiretype end group for non-group")
//		}
//		if fieldNum <= 0 {
//			return fmt.Errorf("proto: InterchainAccount: illegal tag %d (wire type %d)", fieldNum, wire)
//		}
//		switch fieldNum {
//		case 1:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowAccount
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthAccount
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthAccount
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.BaseAccount == nil {
//				m.BaseAccount = &types.BaseAccount{}
//			}
//			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 2:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AccountOwner", wireType)
//			}
//			var stringLen uint64
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowAccount
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				stringLen |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			intStringLen := int(stringLen)
//			if intStringLen < 0 {
//				return ErrInvalidLengthAccount
//			}
//			postIndex := iNdEx + intStringLen
//			if postIndex < 0 {
//				return ErrInvalidLengthAccount
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			m.AccountOwner = string(dAtA[iNdEx:postIndex])
//			iNdEx = postIndex
//		default:
//			iNdEx = preIndex
//			skippy, err := skipAccount(dAtA[iNdEx:])
//			if err != nil {
//				return err
//			}
//			if (skippy < 0) || (iNdEx+skippy) < 0 {
//				return ErrInvalidLengthAccount
//			}
//			if (iNdEx + skippy) > l {
//				return io.ErrUnexpectedEOF
//			}
//			iNdEx += skippy
//		}
//	}
//
//	if iNdEx > l {
//		return io.ErrUnexpectedEOF
//	}
//	return nil
//}
//func skipAccount(dAtA []byte) (n int, err error) {
//	l := len(dAtA)
//	iNdEx := 0
//	depth := 0
//	for iNdEx < l {
//		var wire uint64
//		for shift := uint(0); ; shift += 7 {
//			if shift >= 64 {
//				return 0, ErrIntOverflowAccount
//			}
//			if iNdEx >= l {
//				return 0, io.ErrUnexpectedEOF
//			}
//			b := dAtA[iNdEx]
//			iNdEx++
//			wire |= (uint64(b) & 0x7F) << shift
//			if b < 0x80 {
//				break
//			}
//		}
//		wireType := int(wire & 0x7)
//		switch wireType {
//		case 0:
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return 0, ErrIntOverflowAccount
//				}
//				if iNdEx >= l {
//					return 0, io.ErrUnexpectedEOF
//				}
//				iNdEx++
//				if dAtA[iNdEx-1] < 0x80 {
//					break
//				}
//			}
//		case 1:
//			iNdEx += 8
//		case 2:
//			var length int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return 0, ErrIntOverflowAccount
//				}
//				if iNdEx >= l {
//					return 0, io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				length |= (int(b) & 0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if length < 0 {
//				return 0, ErrInvalidLengthAccount
//			}
//			iNdEx += length
//		case 3:
//			depth++
//		case 4:
//			if depth == 0 {
//				return 0, ErrUnexpectedEndOfGroupAccount
//			}
//			depth--
//		case 5:
//			iNdEx += 4
//		default:
//			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
//		}
//		if iNdEx < 0 {
//			return 0, ErrInvalidLengthAccount
//		}
//		if depth == 0 {
//			return iNdEx, nil
//		}
//	}
//	return 0, io.ErrUnexpectedEOF
//}
//
//var (
//	ErrInvalidLengthAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
//	ErrIntOverflowAccount          = fmt.Errorf("proto: integer overflow")
//	ErrUnexpectedEndOfGroupAccount = fmt.Errorf("proto: unexpected end of group")
//)
