package v1

import (
	"bytes"
	"github.com/golang/protobuf/proto"
)

type ProtoBuffer struct {
	buf bytes.Buffer
}

// Constants that identify the encoding of a value on the wire.
const (
	WireVarint     = 0
	WireFixed64    = 1
	WireBytes      = 2
	WireStartGroup = 3
	WireEndGroup   = 4
	WireFixed32    = 5
)

const ()

func (p *ProtoBuffer) varintWithWireType(x uint64, w int) (err error) {
	fByte := byte(int(x<<3) | w)
	x = x >> 3
	if err = p.buf.WriteByte(fByte); err != nil {
		return
	}
	for x > 0 {
		if err = p.buf.WriteByte(byte(x << 7)); err != nil {
			return err
		}
		x = x >> 7
	}
	return nil
}

//varint
func (p *ProtoBuffer) varint(x uint64) (err error) {
	_, err = p.buf.Write(proto.EncodeVarint(x))
	return
}

func (p *ProtoBuffer) AddBytes(x uint64, bs []byte) (err error) {
	if err = p.varintWithWireType(x, WireBytes); err != nil {
		return
	}
	if _, err = p.buf.Write(proto.EncodeVarint(uint64(len(bs)))); err != nil {
		return
	}
	_, err = p.buf.Write(bs)
	return
}
