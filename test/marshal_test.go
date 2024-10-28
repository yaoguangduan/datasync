package test

import (
	"fmt"
	"github.com/yaoguangduan/protosync/pbgen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"testing"
)

/**
Wire Type 对应表
Protobuf 类型	Wire Type	Wire Type 编号
int32	Varint	0
int64	Varint	0
uint32	Varint	0
uint64	Varint	0
bool	Varint	0
enum	Varint	0
double	64-bit	1
string	Length-delimited	2
bytes	Length-delimited	2
embedded messages	Length-delimited	2
packed repeated fields	Length-delimited	2
float	32-bit	5
*/

func TestMarshal(t *testing.T) {
	act := pbgen.ActionInfoSync{
		Act:    "age",
		Detail: "2werw",
		Time:   222,
	}
	var buf []byte
	buf = protowire.AppendVarint(buf, protowire.EncodeTag(1, protowire.BytesType))
	buf = protowire.AppendString(buf, act.Act)
	buf = protowire.AppendVarint(buf, protowire.EncodeTag(2, protowire.BytesType))
	buf = protowire.AppendString(buf, act.Detail)

	buf = protowire.AppendVarint(buf, protowire.EncodeTag(3, protowire.VarintType))
	buf = protowire.AppendVarint(buf, uint64(act.Time))

	actp := &pbgen.ActionInfoSync{}
	err := proto.Unmarshal(buf, actp)
	if err != nil {
		panic(err)
	}
	fmt.Println(actp)
}
