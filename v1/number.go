package libproto

import (
	"bytes"
	"github.com/golang/protobuf/proto"
)

type ProtoBuffer struct {
	buf *bytes.Buffer
}

func NewProtoBuffer(bs []byte) *ProtoBuffer {
	return &ProtoBuffer{buf: bytes.NewBuffer(bs)}
}

// Constants that identify the encoding of a value on the wire.
const (
	WireVarint  = 0
	WireFixed64 = 1
	WireBytes   = 2
	//WireStartGroup = 3
	//WireEndGroup   = 4
	WireFixed32 = 5
)

func Zigzag32(x int32) uint64 {
	return uint64(x>>31 ^ x<<1)
}

func Zigzag64(x int64) uint64 {
	return uint64(x>>63 ^ x<<1)
}

func Zigzag(x int64) uint64 {
	return uint64(x>>7 ^ x<<1)
}

func GenVarintWithWireType(x uint64, w int) (bs []byte) {
	bs = make([]byte, 0, 10)
	bs = append(bs, byte(int(x<<3)|w))
	x = x >> 4
	for x > 0 {
		tempX := byte(x << 7 >> 7)
		x = x >> 7
		if x > 0 {
			bs = append(bs, tempX|128)
			continue
		}
		bs = append(bs, tempX)
	}
	return bs
}

//GenVarintWithByteSize
func GenVarintWithByteSize(x uint64, size int) (bs []byte) {
	bs = make([]byte, 0, 10)
	bs = append(bs, GenVarintWithWireType(x, WireBytes)...)
	bs = append(bs, GenVarint(uint64(size))...)
	return bs
}

//varint
func GenVarint(x uint64) (bs []byte) {
	return proto.EncodeVarint(x)
}

func GenWireBytes(x uint64, content []byte) (bs []byte) {
	bs = make([]byte, 0, len(content)+10)
	bs = append(bs, GenVarintWithWireType(x, WireBytes)...)
	bs = append(bs, GenVarint(uint64(len(content)))...)
	return append(bs, content...)
}
