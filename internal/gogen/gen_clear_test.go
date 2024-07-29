package gogen

import (
	"github.com/samber/lo"
	"github.com/yaoguangduan/datasync/pbgen"
	"google.golang.org/protobuf/proto"
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
}
