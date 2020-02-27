package libproto

import (
	"encoding/binary"
	"errors"
	"math"
)

type Message interface {
	Bytes() ([]byte, error)
}

type Record map[uint64]Message
type Repeat []Message
type Double float64
type Float float32
type Int32 int32
type Int64 int64
type Uint32 uint32
type Uint64 uint64
type Sint32 int32
type Sint64 int64
type Fixed32 uint32
type Fixed64 uint64
type Sfixed32 int32
type Sfixed64 int64
type Bool bool
type String string
type Bytes []byte

func (r Record) Bytes() ([]byte, error) {
	buf := make([]byte, 0, 256)
	for k, v := range r {
		bs, err := v.Bytes()
		if err != nil {
			return nil, err
		}

		if len(bs) == 0 {
			continue
		}

		switch v.(type) {
		case Record:
			buf = append(buf, GenVarintWithByteSize(k, len(bs))...)
		case Repeat:
			buf = append(buf, GenVarintWithByteSize(k, len(bs))...)
		case Double:
			buf = append(buf, GenVarintWithWireType(k, WireFixed64)...)
		case Float:
			buf = append(buf, GenVarintWithWireType(k, WireFixed32)...)
		case Int32:
			buf = append(buf, GenVarintWithWireType(k, WireVarint)...)
		case Int64:
			buf = append(buf, GenVarintWithWireType(k, WireVarint)...)
		case Uint32:
			buf = append(buf, GenVarintWithWireType(k, WireVarint)...)
		case Uint64:
			buf = append(buf, GenVarintWithWireType(k, WireVarint)...)
		case Sint32:
			buf = append(buf, GenVarintWithWireType(k, WireVarint)...)
		case Sint64:
			buf = append(buf, GenVarintWithWireType(k, WireVarint)...)
		case Fixed32:
			buf = append(buf, GenVarintWithWireType(k, WireFixed32)...)
		case Fixed64:
			buf = append(buf, GenVarintWithWireType(k, WireFixed64)...)
		case Sfixed32:
			buf = append(buf, GenVarintWithWireType(k, WireFixed32)...)
		case Sfixed64:
			buf = append(buf, GenVarintWithWireType(k, WireFixed64)...)
		case Bool:
			buf = append(buf, GenVarintWithWireType(k, WireVarint)...)
		case String:
			buf = append(buf, GenVarintWithByteSize(k, len(bs))...)
		case Bytes:
			buf = append(buf, GenVarintWithByteSize(k, len(bs))...)
		default:
			return nil, errors.New("unsupported type")
		}
		buf = append(buf, bs...)
	}
	return buf, nil
}

func (m Repeat) Bytes() (bs []byte, err error) {
	arrayBuffer := make([]byte, 0, 1024)
	for _, obj := range m {
		bs, err := obj.Bytes()
		if err != nil {
			return nil, err
		}
		arrayBuffer = append(arrayBuffer, bs...)
	}
	return arrayBuffer, nil
}

func (m Double) Bytes() (bs []byte, err error) {
	if m == 0 {
		return
	}

	n := math.Float64bits(float64(m))
	bs = make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, n)
	return bs, nil
}

func (m Float) Bytes() (bs []byte, err error) {
	if m == 0 {
		return
	}
	n := math.Float32bits(float32(m))
	bs = make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, n)
	return bs, nil
}

func (m Int32) Bytes() (bs []byte, err error) {
	if m == 0 {
		return
	}
	return GenVarint(uint64(m)), nil
}

func (m Int64) Bytes() (bs []byte, err error) {
	if m == 0 {
		return
	}
	return GenVarint(uint64(m)), nil
}

func (m Uint32) Bytes() (bs []byte, err error) {
	if m == 0 {
		return
	}
	return GenVarint(uint64(m)), nil
}

func (m Uint64) Bytes() (bs []byte, err error) {
	if m == 0 {
		return
	}
	return GenVarint(uint64(m)), nil
}

func (m Sint32) Bytes() (bs []byte, err error) {
	if m == 0 {
		return
	}
	return GenVarint(Zigzag32(int32(m))), nil
}

func (m Sint64) Bytes() (bs []byte, err error) {
	if m == 0 {
		return
	}
	return GenVarint(Zigzag64(int64(m))), nil
}

func (m Fixed32) Bytes() (bs []byte, err error) {
	if m == 0 {
		return
	}
	bs = make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(m))
	return bs, nil
}

func (m Fixed64) Bytes() (bs []byte, err error) {
	if m == 0 {
		return
	}
	bs = make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(m))
	return bs, nil
}

func (m Sfixed32) Bytes() (bs []byte, err error) {
	if m == 0 {
		return
	}
	bs = make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(m))
	return bs, nil
}

func (m Sfixed64) Bytes() (bs []byte, err error) {
	if m == 0 {
		return
	}
	bs = make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(m))
	return bs, nil
}

func (m Bool) Bytes() (bs []byte, err error) {
	b := func() byte {
		if m {
			return 1
		}
		return 0
	}()
	return []byte{b}, nil
}

func (m String) Bytes() (bs []byte, err error) {
	return []byte(m), nil
}

func (m Bytes) Bytes() (bs []byte, err error) {
	return m, nil
}
