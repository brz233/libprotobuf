package v1

import (
	"fmt"
	pb "github.com/brz233/libprotobuf/v1/protocol"
	"github.com/go-playground/assert"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestRecord_Bytes(t *testing.T) {
	obj := Record{1: String("abc"), 2: Sint32(-56)}
	bs, err := obj.Bytes()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(bs)
}

func TestDouble_Bytes(t *testing.T) {
	record := Record{1: Double(9.7)}
	bs, err := record.Bytes()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, err, nil)

	test1 := pb.Test1{}
	err = proto.Unmarshal(bs, &test1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, err, nil)

	assert.Equal(t, test1.DoubleValue, 9.7)
	if !(test1.DoubleValue-9.7 < 10e-7) {
		t.Error(err, test1.DoubleValue)
	}
}
