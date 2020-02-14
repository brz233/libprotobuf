package v1

import (
	pb "github.com/brz233/libprotobuf/v1/protocol"
	"github.com/go-playground/assert"
	"github.com/golang/protobuf/proto"
	"log"
	"math"
	"testing"
)

//func TestRecord_Bytes(t *testing.T) {
//	obj := Record{1: String("abc"), 2: Sint32(-56)}
//	bs, err := obj.Bytes()
//	if err != nil {
//		t.Error(err)
//	}
//	fmt.Println(bs)
//}

func TestDouble_Bytes(t *testing.T) {
	record := Record{
		1:  Double(7.7),
		2:  Float(7.7),
		3:  Int32(7),
		4:  Int64(7),
		5:  Uint32(7),
		6:  Uint64(7),
		7:  Sint32(-7),
		8:  Sint64(-7),
		9:  Fixed32(7),
		10: Fixed64(7),
		11: Sfixed32(7),
		12: Sfixed64(7),
		13: Bool(true),
		14: String("ABCD"),
		15: Bytes([]byte{1, 3, 4}),
	}
	bs, err := record.Bytes()
	assert.Equal(t, err, nil)

	test1 := pb.Test1{}
	err = proto.Unmarshal(bs, &test1)
	assert.Equal(t, err, nil)

	assert.Equal(t, math.Abs(test1.DoubleValue-7.7) < 1e-10, true)
	log.Println("test double pass")
	assert.Equal(t, math.Abs(float64(test1.FloatValue-7.7)) < 1e-10, true)
	log.Println("test float pass")

	assert.Equal(t, test1.I32Value, int32(7))
	log.Println("test int32 pass")
	assert.Equal(t, test1.I64Value, int64(7))
	log.Println("test int64 pass")

	assert.Equal(t, test1.U32Value, uint32(7))
	log.Println("test uint32 pass")
	assert.Equal(t, test1.U64Value, uint64(7))
	log.Println("test uint64 pass")

	assert.Equal(t, test1.S32Value, int32(-7))
	log.Println("test sint32 pass")
	assert.Equal(t, test1.S64Value, int64(-7))
	log.Println("test sint64 pass")

	assert.Equal(t, test1.F32Value, uint32(7))
	log.Println("test fixed32 pass")
	assert.Equal(t, test1.F64Value, uint64(7))
	log.Println("test fixed64 pass")

	assert.Equal(t, test1.Sf32Value, int32(7))
	log.Println("test sfixed32 pass")
	assert.Equal(t, test1.Sf64Value, int64(7))
	log.Println("test sfixed64 pass")

	assert.Equal(t, test1.BoolValue, true)
	log.Println("test bool pass")

	assert.Equal(t, test1.StringValue, "ABCD")
	log.Println("test string pass")

	assert.Equal(t, test1.BytesValue, []byte{1, 3, 4})
	log.Println("test string pass")
}
