// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test.proto

package protocol

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Test1 struct {
	DoubleNumber         float64  `protobuf:"fixed64,1,opt,name=double_number,json=doubleNumber,proto3" json:"double_number,omitempty"`
	FloatNumber          float32  `protobuf:"fixed32,2,opt,name=float_number,json=floatNumber,proto3" json:"float_number,omitempty"`
	I32Number            int32    `protobuf:"varint,3,opt,name=i32_number,json=i32Number,proto3" json:"i32_number,omitempty"`
	I64Number            int64    `protobuf:"varint,4,opt,name=i64_number,json=i64Number,proto3" json:"i64_number,omitempty"`
	U32Number            uint32   `protobuf:"varint,5,opt,name=u32_number,json=u32Number,proto3" json:"u32_number,omitempty"`
	U64Number            uint64   `protobuf:"varint,6,opt,name=u64_number,json=u64Number,proto3" json:"u64_number,omitempty"`
	S32Number            int32    `protobuf:"zigzag32,7,opt,name=s32_number,json=s32Number,proto3" json:"s32_number,omitempty"`
	S64Number            int64    `protobuf:"zigzag64,8,opt,name=s64_number,json=s64Number,proto3" json:"s64_number,omitempty"`
	F32Number            uint32   `protobuf:"fixed32,9,opt,name=f32_number,json=f32Number,proto3" json:"f32_number,omitempty"`
	F64Number            uint64   `protobuf:"fixed64,10,opt,name=f64_number,json=f64Number,proto3" json:"f64_number,omitempty"`
	Sf32Number           int32    `protobuf:"fixed32,11,opt,name=sf32_number,json=sf32Number,proto3" json:"sf32_number,omitempty"`
	Sf64Number           int64    `protobuf:"fixed64,12,opt,name=sf64_number,json=sf64Number,proto3" json:"sf64_number,omitempty"`
	BoolValue            bool     `protobuf:"varint,13,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	StringValue          string   `protobuf:"bytes,14,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	BytesValue           []byte   `protobuf:"bytes,15,opt,name=bytes_value,json=bytesValue,proto3" json:"bytes_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test1) Reset()         { *m = Test1{} }
func (m *Test1) String() string { return proto.CompactTextString(m) }
func (*Test1) ProtoMessage()    {}
func (*Test1) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}
func (m *Test1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Test1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Test1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Test1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test1.Merge(m, src)
}
func (m *Test1) XXX_Size() int {
	return m.Size()
}
func (m *Test1) XXX_DiscardUnknown() {
	xxx_messageInfo_Test1.DiscardUnknown(m)
}

var xxx_messageInfo_Test1 proto.InternalMessageInfo

func (m *Test1) GetDoubleNumber() float64 {
	if m != nil {
		return m.DoubleNumber
	}
	return 0
}

func (m *Test1) GetFloatNumber() float32 {
	if m != nil {
		return m.FloatNumber
	}
	return 0
}

func (m *Test1) GetI32Number() int32 {
	if m != nil {
		return m.I32Number
	}
	return 0
}

func (m *Test1) GetI64Number() int64 {
	if m != nil {
		return m.I64Number
	}
	return 0
}

func (m *Test1) GetU32Number() uint32 {
	if m != nil {
		return m.U32Number
	}
	return 0
}

func (m *Test1) GetU64Number() uint64 {
	if m != nil {
		return m.U64Number
	}
	return 0
}

func (m *Test1) GetS32Number() int32 {
	if m != nil {
		return m.S32Number
	}
	return 0
}

func (m *Test1) GetS64Number() int64 {
	if m != nil {
		return m.S64Number
	}
	return 0
}

func (m *Test1) GetF32Number() uint32 {
	if m != nil {
		return m.F32Number
	}
	return 0
}

func (m *Test1) GetF64Number() uint64 {
	if m != nil {
		return m.F64Number
	}
	return 0
}

func (m *Test1) GetSf32Number() int32 {
	if m != nil {
		return m.Sf32Number
	}
	return 0
}

func (m *Test1) GetSf64Number() int64 {
	if m != nil {
		return m.Sf64Number
	}
	return 0
}

func (m *Test1) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *Test1) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *Test1) GetBytesValue() []byte {
	if m != nil {
		return m.BytesValue
	}
	return nil
}

func init() {
	proto.RegisterType((*Test1)(nil), "protocol.test1")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0x3f, 0x4e, 0xf3, 0x40,
	0x10, 0x47, 0xbf, 0xc9, 0xff, 0x1d, 0x3b, 0x5f, 0x82, 0xab, 0x34, 0x38, 0x03, 0x34, 0x53, 0x21,
	0x91, 0x44, 0x39, 0x00, 0x07, 0xa0, 0x98, 0x82, 0x36, 0x8a, 0xc1, 0x46, 0x96, 0x4c, 0x16, 0x65,
	0xbd, 0x48, 0xdc, 0x84, 0x23, 0x51, 0x72, 0x04, 0x14, 0x5a, 0x0e, 0x81, 0xbc, 0x89, 0xd9, 0xad,
	0x2c, 0xbd, 0xdf, 0x7b, 0x5e, 0x69, 0xb5, 0x88, 0x75, 0x6e, 0xea, 0xeb, 0x97, 0xbd, 0xae, 0x75,
	0x32, 0x72, 0x9f, 0x07, 0x5d, 0x5d, 0xfe, 0x74, 0xb1, 0xdf, 0x0c, 0x37, 0xc9, 0x15, 0x8e, 0x1f,
	0xb5, 0xcd, 0xaa, 0x7c, 0xb3, 0xb3, 0xcf, 0x59, 0xbe, 0x9f, 0x01, 0x01, 0x83, 0xc4, 0x47, 0x78,
	0xe7, 0x58, 0x72, 0x81, 0x71, 0x51, 0xe9, 0x6d, 0xdd, 0x3a, 0x1d, 0x02, 0xee, 0x48, 0xe4, 0xd8,
	0x49, 0x39, 0x47, 0x2c, 0x97, 0x8b, 0x56, 0xe8, 0x12, 0x70, 0x5f, 0x54, 0xb9, 0x5c, 0x04, 0xf3,
	0x7a, 0xd5, 0xce, 0x3d, 0x02, 0xee, 0x8a, 0x2a, 0xd7, 0x2b, 0x3f, 0x5b, 0x5f, 0xf7, 0x09, 0x78,
	0x2c, 0xca, 0x86, 0xb5, 0xf5, 0xf5, 0x80, 0x80, 0x7b, 0xa2, 0x6c, 0x58, 0x1b, 0x5f, 0x0f, 0x09,
	0xf8, 0x4c, 0x94, 0x09, 0x6b, 0xe3, 0xeb, 0x11, 0x01, 0x27, 0xa2, 0x4c, 0x58, 0x17, 0xbe, 0x56,
	0x04, 0x3c, 0x14, 0x55, 0x84, 0x75, 0xe1, 0x6b, 0x24, 0xe0, 0x81, 0xa8, 0xe2, 0xaf, 0x9e, 0x63,
	0x64, 0x82, 0x3c, 0x22, 0xe0, 0x89, 0xa0, 0xf1, 0xbd, 0x13, 0xfc, 0x0f, 0x62, 0x02, 0x9e, 0x36,
	0x42, 0x78, 0x7e, 0xa6, 0x75, 0xb5, 0x79, 0xdd, 0x56, 0x36, 0x9f, 0x8d, 0x09, 0x78, 0x24, 0xaa,
	0x21, 0xf7, 0x0d, 0x68, 0xee, 0xde, 0xd4, 0xfb, 0x72, 0xf7, 0x74, 0x12, 0xfe, 0x13, 0xb0, 0x92,
	0xe8, 0xc8, 0x8e, 0xca, 0x1c, 0xa3, 0xec, 0xad, 0xce, 0xcd, 0xc9, 0x98, 0x10, 0x70, 0x2c, 0xe8,
	0x90, 0x13, 0x6e, 0xa7, 0x1f, 0x87, 0x14, 0x3e, 0x0f, 0x29, 0x7c, 0x1d, 0x52, 0x78, 0xff, 0x4e,
	0xff, 0x65, 0x03, 0xf7, 0x14, 0x96, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x09, 0xe0, 0x6f,
	0x1f, 0x02, 0x00, 0x00,
}

func (m *Test1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Test1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Test1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.BytesValue) > 0 {
		i -= len(m.BytesValue)
		copy(dAtA[i:], m.BytesValue)
		i = encodeVarintTest(dAtA, i, uint64(len(m.BytesValue)))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.StringValue) > 0 {
		i -= len(m.StringValue)
		copy(dAtA[i:], m.StringValue)
		i = encodeVarintTest(dAtA, i, uint64(len(m.StringValue)))
		i--
		dAtA[i] = 0x72
	}
	if m.BoolValue {
		i--
		if m.BoolValue {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x68
	}
	if m.Sf64Number != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Sf64Number))
		i--
		dAtA[i] = 0x61
	}
	if m.Sf32Number != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.Sf32Number))
		i--
		dAtA[i] = 0x5d
	}
	if m.F64Number != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.F64Number))
		i--
		dAtA[i] = 0x51
	}
	if m.F32Number != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.F32Number))
		i--
		dAtA[i] = 0x4d
	}
	if m.S64Number != 0 {
		i = encodeVarintTest(dAtA, i, uint64((uint64(m.S64Number)<<1)^uint64((m.S64Number>>63))))
		i--
		dAtA[i] = 0x40
	}
	if m.S32Number != 0 {
		i = encodeVarintTest(dAtA, i, uint64((uint32(m.S32Number)<<1)^uint32((m.S32Number>>31))))
		i--
		dAtA[i] = 0x38
	}
	if m.U64Number != 0 {
		i = encodeVarintTest(dAtA, i, uint64(m.U64Number))
		i--
		dAtA[i] = 0x30
	}
	if m.U32Number != 0 {
		i = encodeVarintTest(dAtA, i, uint64(m.U32Number))
		i--
		dAtA[i] = 0x28
	}
	if m.I64Number != 0 {
		i = encodeVarintTest(dAtA, i, uint64(m.I64Number))
		i--
		dAtA[i] = 0x20
	}
	if m.I32Number != 0 {
		i = encodeVarintTest(dAtA, i, uint64(m.I32Number))
		i--
		dAtA[i] = 0x18
	}
	if m.FloatNumber != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.FloatNumber))))
		i--
		dAtA[i] = 0x15
	}
	if m.DoubleNumber != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.DoubleNumber))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func encodeVarintTest(dAtA []byte, offset int, v uint64) int {
	offset -= sovTest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Test1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DoubleNumber != 0 {
		n += 9
	}
	if m.FloatNumber != 0 {
		n += 5
	}
	if m.I32Number != 0 {
		n += 1 + sovTest(uint64(m.I32Number))
	}
	if m.I64Number != 0 {
		n += 1 + sovTest(uint64(m.I64Number))
	}
	if m.U32Number != 0 {
		n += 1 + sovTest(uint64(m.U32Number))
	}
	if m.U64Number != 0 {
		n += 1 + sovTest(uint64(m.U64Number))
	}
	if m.S32Number != 0 {
		n += 1 + sozTest(uint64(m.S32Number))
	}
	if m.S64Number != 0 {
		n += 1 + sozTest(uint64(m.S64Number))
	}
	if m.F32Number != 0 {
		n += 5
	}
	if m.F64Number != 0 {
		n += 9
	}
	if m.Sf32Number != 0 {
		n += 5
	}
	if m.Sf64Number != 0 {
		n += 9
	}
	if m.BoolValue {
		n += 2
	}
	l = len(m.StringValue)
	if l > 0 {
		n += 1 + l + sovTest(uint64(l))
	}
	l = len(m.BytesValue)
	if l > 0 {
		n += 1 + l + sovTest(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTest(x uint64) (n int) {
	return sovTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Test1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
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
			return fmt.Errorf("proto: test1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: test1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field DoubleNumber", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.DoubleNumber = float64(math.Float64frombits(v))
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloatNumber", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.FloatNumber = float32(math.Float32frombits(v))
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field I32Number", wireType)
			}
			m.I32Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.I32Number |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field I64Number", wireType)
			}
			m.I64Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.I64Number |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field U32Number", wireType)
			}
			m.U32Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.U32Number |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field U64Number", wireType)
			}
			m.U64Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.U64Number |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field S32Number", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.S32Number = v
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field S64Number", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.S64Number = int64(v)
		case 9:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field F32Number", wireType)
			}
			m.F32Number = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.F32Number = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 10:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field F64Number", wireType)
			}
			m.F64Number = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.F64Number = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 11:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sf32Number", wireType)
			}
			m.Sf32Number = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.Sf32Number = int32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 12:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sf64Number", wireType)
			}
			m.Sf64Number = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Sf64Number = int64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoolValue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BoolValue = bool(v != 0)
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StringValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BytesValue", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BytesValue = append(m.BytesValue[:0], dAtA[iNdEx:postIndex]...)
			if m.BytesValue == nil {
				m.BytesValue = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTest
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
					return 0, ErrIntOverflowTest
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
					return 0, ErrIntOverflowTest
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
				return 0, ErrInvalidLengthTest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTest = fmt.Errorf("proto: unexpected end of group")
)
