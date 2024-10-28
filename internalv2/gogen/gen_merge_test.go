package gogen

import (
	"github.com/stretchr/testify/assert"
	"github.com/yaoguangduan/protosync/pbgen"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestAllTypeMergePB(t *testing.T) {
	test := fullTestData()
	test.FlushDirty(false)
	another := fullTestData()
	another.FlushDirty(false)
	modifyAll(test)
	ts := &pbgen.Test{}
	test.MergeDirtyToPb(ts)
	t.Log(ts)
	another.MergeDirtyFromPb(ts)
	//修改后的数据 = 原始数据 + 脏数据
	testD := &pbgen.Test{}
	anotherD := &pbgen.Test{}
	test.CopyToPb(testD)
	t.Log(testD)
	another.CopyToPb(anotherD)
	t.Log(anotherD)
	assert.Equal(t, test, another)
}

func TestAllTypeMergeBytes(t *testing.T) {
	test := fullTestData()
	bytes := test.MergeDirtyToBytes()

	testSync := &pbgen.Test{}
	err := proto.Unmarshal(bytes, testSync)
	assert.NoError(t, err)
	t.Log(testSync)

	testSyncPb := &pbgen.Test{}
	test.MergeDirtyToPb(testSyncPb)
	t.Log(testSyncPb)
	assert.Equal(t, testSync, testSyncPb)

}

func TestAllTypeMergeFromBytes(t *testing.T) {
	test := fullTestData()
	bytes := test.MergeDirtyToBytes()

	testSync := &pbgen.Test{}
	err := proto.Unmarshal(bytes, testSync)
	assert.NoError(t, err)

	testNew := pbgen.NewTestSync()
	testNew.MergeDirtyFromBytes(bytes)

	testD := &pbgen.Test{}
	testNewD := &pbgen.Test{}
	test.CopyToPb(testD)
	testNew.CopyToPb(testNewD)
	t.Log(testSync)
	t.Log(testD)
	t.Log(testNewD)
	assert.Equal(t, test, testNew)

}
