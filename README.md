# libprotobuf

该项目的目的是用于动态生成protobuf的数据格式的数据。

```go
//main.go
package main

import (
	"fmt"
	libproto "github.com/brz233/libprotobuf/v1"
	"github.com/brz233/libprotobuf/v1/protocol"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)


func main() {
	record := libproto.Record{
		1: libproto.Uint32(1),
		2: libproto.String("Tony"),
		3: libproto.String("man"),
		4: libproto.Record{
			1: libproto.Uint32(1999),
			2: libproto.Uint32(8),
			3: libproto.Uint32(8),
			4: libproto.String("china"),
			5: libproto.String("guangzhou"),
		},
	}

	bs, err := record.Bytes()
	if err != nil {
		log.Fatal(err)
	}

	var user protocol.User
	err = proto.Unmarshal(bs, &user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)

}

```
