package gogen

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/yaoguangduan/protosync/pbgen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protopack"
	"testing"
)

func TestClear(t *testing.T) {
	test := fullTestData()
	test.FlushDirty(false)
	test.Clear()
	rs := pbgen.Test{}
	test.MergeDirtyToPb(&rs)
	t.Log(&rs)

	testNew := pbgen.NewTestSync()
	testNew.MergeDirtyFromPb(&rs)

	test.FlushDirty(false)
	testNew.FlushDirty(false)
	//assert.Equal(t, test, testNew)
	ti := &pbgen.Test{}
	msg := ti.ProtoReflect()

	bys := lo.Must(proto.Marshal(ti))

	tii := &pbgen.Test{}
	lo.Must0(proto.Unmarshal(bys, tii))

	t.Log(msg.Has(tii.ProtoReflect().Descriptor().Fields().ByName("b")))
	t.Log(msg.Has(tii.ProtoReflect().Descriptor().Fields().ByName("id")))

	newt := fullTestData()
	bys = newt.MergeDirtyToBytes()

	var packMsg protopack.Message
	packMsg = append(packMsg, protopack.Tag{Number: 1012, Type: 2}, protopack.Bool(false))

	// 遍历未知字段
	for _, field := range packMsg {
		fmt.Printf("Field Number: %+v\n",
			field)
	}
}
