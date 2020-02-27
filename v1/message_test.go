package v1

import (
	pb "github.com/brz233/libprotobuf/v1/protocol"
	"github.com/go-playground/assert"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestDouble_Bytes(t *testing.T) {
	item := Record{
		1: Double(0.3),
		2: Double(4.9),
		3: Double(1000.45),
		4: Double(-8889.45),
		5: Double(0),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pItem pb.DoubleItem
	err = proto.Unmarshal(bs, &pItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, 0.3, pItem.Value1)
	assert.Equal(t, 4.9, pItem.Value2)
	assert.Equal(t, 1000.45, pItem.Value3)
	assert.Equal(t, -8889.45, pItem.Value4)
	assert.Equal(t, float64(0), pItem.Value5)
}

func TestFloat_Bytes(t *testing.T) {
	item := Record{
		1: Float(0.3),
		2: Float(4.9),
		3: Float(1000.45),
		4: Float(-8889.45),
		5: Float(0),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pItem pb.FloatItem
	err = proto.Unmarshal(bs, &pItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, float32(0.3), pItem.Value1)
	assert.Equal(t, float32(4.9), pItem.Value2)
	assert.Equal(t, float32(1000.45), pItem.Value3)
	assert.Equal(t, float32(-8889.45), pItem.Value4)
	assert.Equal(t, float32(0), pItem.Value5)
}

func TestInt32_Bytes(t *testing.T) {
	item := Record{
		1: Int32(3),
		2: Int32(1000),
		3: Int32(20000000),
		4: Int32(-1000),
		5: Int32(0),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pItem pb.Int32Item
	err = proto.Unmarshal(bs, &pItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, int32(3), pItem.Value1)
	assert.Equal(t, int32(1000), pItem.Value2)
	assert.Equal(t, int32(20000000), pItem.Value3)
	assert.Equal(t, int32(-1000), pItem.Value4)
	assert.Equal(t, int32(0), pItem.Value5)
}

func TestInt64_Bytes(t *testing.T) {
	item := Record{
		1: Int64(3),
		2: Int64(1000),
		3: Int64(20000000),
		4: Int64(-1000),
		5: Int64(0),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pItem pb.Int64Item
	err = proto.Unmarshal(bs, &pItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, int64(3), pItem.Value1)
	assert.Equal(t, int64(1000), pItem.Value2)
	assert.Equal(t, int64(20000000), pItem.Value3)
	assert.Equal(t, int64(-1000), pItem.Value4)
	assert.Equal(t, int64(0), pItem.Value5)
}

func TestUint32_Bytes(t *testing.T) {
	item := Record{
		1: Uint32(3),
		2: Uint32(1000),
		3: Uint32(20000000),
		4: Uint32(889999566),
		5: Uint32(0),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pItem pb.Uint32Item
	err = proto.Unmarshal(bs, &pItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, uint32(3), pItem.Value1)
	assert.Equal(t, uint32(1000), pItem.Value2)
	assert.Equal(t, uint32(20000000), pItem.Value3)
	assert.Equal(t, uint32(889999566), pItem.Value4)
	assert.Equal(t, uint32(0), pItem.Value5)
}

func TestUint64_Bytes(t *testing.T) {
	item := Record{
		1: Uint64(3),
		2: Uint64(1000),
		3: Uint64(20000000),
		4: Uint64(889999566334444),
		5: Uint64(0),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pItem pb.Uint64Item
	err = proto.Unmarshal(bs, &pItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, uint64(3), pItem.Value1)
	assert.Equal(t, uint64(1000), pItem.Value2)
	assert.Equal(t, uint64(20000000), pItem.Value3)
	assert.Equal(t, uint64(889999566334444), pItem.Value4)
	assert.Equal(t, uint64(0), pItem.Value5)
}

func TestSint32_Bytes(t *testing.T) {
	item := Record{
		1: Sint32(4),
		2: Sint32(256),
		3: Sint32(100000000),
		4: Sint32(-512),
		5: Sint32(0),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pItem pb.Sint32Item
	err = proto.Unmarshal(bs, &pItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, int32(4), pItem.Value1)
	assert.Equal(t, int32(256), pItem.Value2)
	assert.Equal(t, int32(100000000), pItem.Value3)
	assert.Equal(t, int32(-512), pItem.Value4)
	assert.Equal(t, int32(0), pItem.Value5)
}

func TestSint64_Bytes(t *testing.T) {
	item := Record{
		1: Sint64(4),
		2: Sint64(256),
		3: Sint64(100000000),
		4: Sint64(-512),
		5: Sint64(0),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pItem pb.Sint64Item
	err = proto.Unmarshal(bs, &pItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, int64(4), pItem.Value1)
	assert.Equal(t, int64(256), pItem.Value2)
	assert.Equal(t, int64(100000000), pItem.Value3)
	assert.Equal(t, int64(-512), pItem.Value4)
	assert.Equal(t, int64(0), pItem.Value5)
}

func TestFixed32_Bytes(t *testing.T) {
	item := Record{
		1: Fixed32(4),
		2: Fixed32(256),
		3: Fixed32(100000000),
		4: Fixed32(512),
		5: Fixed32(0),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pItem pb.Fixed32Item
	err = proto.Unmarshal(bs, &pItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, uint32(4), pItem.Value1)
	assert.Equal(t, uint32(256), pItem.Value2)
	assert.Equal(t, uint32(100000000), pItem.Value3)
	assert.Equal(t, uint32(512), pItem.Value4)
	assert.Equal(t, uint32(0), pItem.Value5)
}

func TestFixed64_Bytes(t *testing.T) {
	item := Record{
		1: Fixed64(4),
		2: Fixed64(256),
		3: Fixed64(100000000),
		4: Fixed64(512),
		5: Fixed64(0),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pItem pb.Fixed64Item
	err = proto.Unmarshal(bs, &pItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, uint64(4), pItem.Value1)
	assert.Equal(t, uint64(256), pItem.Value2)
	assert.Equal(t, uint64(100000000), pItem.Value3)
	assert.Equal(t, uint64(512), pItem.Value4)
	assert.Equal(t, uint64(0), pItem.Value5)
}

func TestSfixed32_Bytes(t *testing.T) {
	item := Record{
		1: Sfixed32(4),
		2: Sfixed32(256),
		3: Sfixed32(100000000),
		4: Sfixed32(-512),
		5: Fixed32(0),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pItem pb.Sfixed32Item
	err = proto.Unmarshal(bs, &pItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, int32(4), pItem.Value1)
	assert.Equal(t, int32(256), pItem.Value2)
	assert.Equal(t, int32(100000000), pItem.Value3)
	assert.Equal(t, int32(-512), pItem.Value4)
	assert.Equal(t, int32(0), pItem.Value5)
}

func TestSfixed64_Bytes(t *testing.T) {
	item := Record{
		1: Sfixed64(4),
		2: Sfixed64(256),
		3: Sfixed64(100000000),
		4: Sfixed64(-512),
		5: Fixed64(0),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pItem pb.Sfixed64Item
	err = proto.Unmarshal(bs, &pItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, int64(4), pItem.Value1)
	assert.Equal(t, int64(256), pItem.Value2)
	assert.Equal(t, int64(100000000), pItem.Value3)
	assert.Equal(t, int64(-512), pItem.Value4)
	assert.Equal(t, int64(0), pItem.Value5)
}

func TestBool_Bytes(t *testing.T) {
	item := Record{
		1: Bool(true),
		2: Bool(false),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pItem pb.BoolItem
	err = proto.Unmarshal(bs, &pItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, true, pItem.Value1)
	assert.Equal(t, false, pItem.Value2)

}

func TestString_Bytes(t *testing.T) {
	item := Record{
		1: String("ABCDEFG"),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pItem pb.StringItem
	err = proto.Unmarshal(bs, &pItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, "ABCDEFG", pItem.Value1)
}

func TestBytes_Bytes(t *testing.T) {
	item := Record{
		1: Bytes("ABCDEFG"),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pItem pb.StringItem
	err = proto.Unmarshal(bs, &pItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, "ABCDEFG", pItem.Value1)

	// 测试长度
	item2 := Record{1: Bytes("")}
	bs, err = item2.Bytes()
	assert.Equal(t, nil, err)
}

func TestRepeat_Bytes(t *testing.T) {
	item := Record{
		1: Repeat{
			Double(1.2),
			Double(3.4),
			Double(-0.889),
		},
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pbItem pb.RepeatedItem
	err = proto.Unmarshal(bs, &pbItem)
	assert.Equal(t, nil, err)

	tmap := map[float64]struct{}{1.2: {}, 3.4: {}, -0.889: {}}
	for _, v := range pbItem.Values {
		_, has := tmap[v]
		assert.Equal(t, true, has)
	}
}

func TestRecord_Bytes(t *testing.T) {
	item := Record{
		1: Record{
			1: Int32(56),
			2: String("233"),
		},
		2: Int32(72),
	}

	bs, err := item.Bytes()
	assert.Equal(t, nil, err)

	var pbItem pb.RecordItem
	err = proto.Unmarshal(bs, &pbItem)
	assert.Equal(t, nil, err)

	assert.Equal(t, int32(72), pbItem.Num)
	assert.Equal(t, int32(56), pbItem.Inner.Id)
	assert.Equal(t, "233", pbItem.Inner.Name)
}
