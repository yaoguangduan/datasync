package gogen

import (
	"github.com/stretchr/testify/assert"
	"github.com/yaoguangduan/protosync/pbgen"
	"testing"
)

func TestCopy(t *testing.T) {
	test := fullTestData()
	test.FlushDirty(false)
	ts := pbgen.Test{}
	test.CopyToPb(&ts)

	testNew := pbgen.NewTestSync()
	testNew.CopyFromPb(&ts)
	testNew.FlushDirty(false)

	assert.Equal(t, test, testNew)
}
