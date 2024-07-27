package gogen

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"gpsync/pbgen"
	"testing"
)

func TestZeroSet(t *testing.T) {
	action := pbgen.NewActionInfoSync()
	action.SetDetail("onestring")

	var r = &pbgen.ActionInfo{}
	action.MergeDirtyToPb(r)
	t.Log(r)
	assert.Equal(t, action.GetDetail(), r.GetDetail())
	assert.Equal(t, action.GetTime(), r.GetTime())

	action.FlushDirty(false)
	r = &pbgen.ActionInfo{}
	action.MergeDirtyToPb(r)
	t.Log(r)
	assert.Equal(t, "", r.GetDetail())
	assert.Equal(t, int64(0), r.GetTime())

	action.SetAct("q")
	r = &pbgen.ActionInfo{}
	action.MergeDirtyToPb(r)
	t.Log(r)

	//通过bitfield避免字段有值 再设置成空值，这个时候merge，无法将原始字段设置为空了
	action.SetAct("new val")
	action.FlushDirty(false)
	action.MergeDirtyFromPb(r)
	assert.Equal(t, "q", action.GetAct())

	p := pbgen.NewPersonSync()
	p.SetIsGirl(false)
	var ps = pbgen.Person{}
	p.MergeDirtyToPb(&ps)
	p.SetIsGirl(true)
	p.FlushDirty(false)

	var psc = pbgen.Person{}
	p.MergeDirtyToPb(&psc)

	t.Log(&ps)
	p.MergeDirtyFromPb(&ps)
	assert.True(t, p.GetIsGirl())
}
func TestList(t *testing.T) {
	p := pbgen.NewPersonSync()
	p.GetFavor().Add("apple")
	p.GetFavor().Add("bnn")

	var ps1 = pbgen.Person{}
	p.MergeDirtyToPb(&ps1)
	t.Log(prototext.Format(&ps1))
	t.Log(1 << 5)
	assert.Equal(t, []string{"apple", "bnn"}, ps1.GetFavor())

	p.GetFavor().Add("bnn1")
	p.GetFavor().Clear()
	ps1 = pbgen.Person{}
	p.MergeDirtyToPb(&ps1)
	t.Log(prototext.Format(&ps1))
	p.FlushDirty(false)
	p.MergeDirtyFromPb(&ps1)
	t.Log(p.GetFavor())
}

func mockPersonData() pbgen.PersonSync {
	p := pbgen.NewPersonSync()
	p.SetIsGirl(true).SetName("john").SetAge(22).GetFavor().Add("basket")
	p.GetDetail().SetMoney(2912).SetAddress("bj")
	p.GetLoveSeq().Add(pbgen.ColorType_Green)
	p.GetLoveSeq().Add(pbgen.ColorType_Red)
	a := pbgen.NewActionInfoSync()
	a.SetAct("sleep").SetTime(24).SetDetail("sleep in bed")
	p.GetActions().Put(a.GetAct(), a)
	p.FlushDirty(false)
	return *p
}

func TestMockTimeLine(t *testing.T) {
	p := mockPersonData()

	// 以上是从db加载的最原始的数据
	// 一次操作，修改它的爱好
	p.GetFavor().Add("swim")
	p.GetDetail().SetMoney(p.GetDetail().GetMoney() + 2000)
	// 操作好以后，1.将脏数据及时入db；2.将脏数据下发客户端
	dirty1 := &pbgen.Person{}
	p.MergeDirtyToPb(dirty1)
	p.FlushDirty(false)
	t.Log(prototext.Format(dirty1))

	// 二次操作，修改
	p.SetAge(p.GetAge() + 1).SetName(p.GetName() + ".jjj")
	p.GetActions().Get("sleep").SetDetail("sleep real in room")
	// 操作好以后，1.将脏数据及时入db；2.将脏数据下发客户端
	p.MergeDirtyToPb(dirty1)
	p.FlushDirty(false)
	t.Log(prototext.Format(dirty1))

	// 3
	p.SetAge(p.GetAge() + 1).SetName(p.GetName() + ".jjj")
	a := pbgen.NewActionInfoSync()
	a.SetAct("eat").SetTime(1).SetDetail("not very e")
	p.GetActions().Put(a.GetAct(), a)
	// 操作好以后，1.将脏数据及时入db；2.将脏数据下发客户端
	p.MergeDirtyToPb(dirty1)
	p.FlushDirty(false)
	t.Log(prototext.Format(dirty1))

	//4
	p.GetActions().Remove("sleep")
	// 操作好以后，1.将脏数据及时入db；2.将脏数据下发客户端
	p.MergeDirtyToPb(dirty1)
	p.FlushDirty(false)
	t.Log("fourth modify:", prototext.Format(dirty1))

	//5
	aa := pbgen.NewActionInfoSync()
	aa.SetAct("sleep").SetTime(21).SetDetail("sssssssssss")
	p.GetActions().Put(aa.GetAct(), aa)
	// 操作好以后，1.将脏数据及时入db；2.将脏数据下发客户端
	p.MergeDirtyToPb(dirty1)
	p.FlushDirty(false)
	t.Log("five modify:", prototext.Format(dirty1))
	//6
	p.GetActions().Clear()
	p.GetDetail().Clear()
	// 操作好以后，1.将脏数据及时入db；2.将脏数据下发客户端
	p.MergeDirtyToPb(dirty1)
	p.FlushDirty(false)
	t.Log("six modify:", prototext.Format(dirty1))

	//系统挂掉，只保存了dirty数据和原始数据
	po := mockPersonData()
	po.MergeDirtyFromPb(dirty1)
	po.FlushDirty(false)
	t.Log(p.GetFavor())
	t.Log(po.GetFavor())
	assert.Equal(t, p, po)

	pd := &pbgen.Person{}
	pod := &pbgen.Person{}
	p.CopyToPb(pd)
	po.CopyToPb(pod)
	assert.Equal(t, pd, pod)
}

func TestBytes(t *testing.T) {
	p := pbgen.NewPersonSync()
	p.SetName("name")
	p.SetData([]byte("hello-world"))
	p.GetDetail().SetMoney(223344)
	p.GetFavor().Add("sleep")
	p.GetLoveSeq().Add(pbgen.ColorType_Green)
	p.GetLoveSeq().Add(pbgen.ColorType_Red)

	a := pbgen.NewActionInfoSync().SetTime(111)
	p.GetActions().Put("act", a)

	bytes := p.MergeDirtyToBytes()

	pc := pbgen.Person{}
	err := proto.Unmarshal(bytes, &pc)
	if err != nil {
		panic(err)
	}
	t.Log(&pc)

	assert.Equal(t, "hello-world", string(pc.GetData()))

}

func fullTestData() *pbgen.TestSync {

	test := pbgen.NewTestSync()
	test.SetId(-32).SetI64(-64).SetU64(64).SetU32(32).SetF32(12.23).SetF64(64.23).SetStr("str").SetB(true).SetObj(pbgen.NewPersonSync().SetAge(11111))
	test.GetStrArr().Add("arr1")
	test.GetStrArr().Add("arr2")
	test.GetEnumArr().Add(pbgen.ColorType_Blue)
	test.GetBoolArr().Add(false)
	test.GetI32Arr().Add(-32)
	test.GetI32Arr().Add(-16)
	test.GetU32Arr().Add(32)
	test.GetU32Arr().Add(16)
	test.GetI64Arr().Add(-64)
	test.GetU64Arr().Add(64)
	test.GetI32Map().PutOne(pbgen.NewTestI32MapSync().SetId(-23).SetAddition("i32map"))
	test.GetStrMap().PutOne(pbgen.NewTestStringMapSync().SetId("sm").SetAddition("sm"))
	test.GetI64Map().PutOne(pbgen.NewTestI64MapSync().SetId(-64).SetAddition("i64map"))
	test.GetBoolMap().PutOne(pbgen.NewTestBoolMapSync().SetId(true).SetAddition("i32map"))
	test.GetU64Map().PutOne(pbgen.NewTestU64MapSync().SetId(64).SetAddition("i32map"))
	test.GetU64Map().PutOne(pbgen.NewTestU64MapSync().SetId(640).SetAddition("i32map2"))
	test.GetU32Map().PutOne(pbgen.NewTestU32MapSync().SetId(32).SetAddition("i32map"))
	return test
}
func modifyAll(test *pbgen.TestSync) {
	test.SetId(-312).SetI64(-624).SetU64(694).SetU32(328).SetStr("stro").SetB(false).SetObj(pbgen.NewPersonSync().SetName("11111").SetAge(1))
	test.GetStrArr().Add("arr12")
	test.GetEnumArr().Add(pbgen.ColorType_Blue)
	test.GetBoolArr().Add(false)
	test.GetI32Arr().Clear()
	test.GetU32Arr().Add(132)
	test.GetU32Arr().Add(126)
	test.GetI64Arr().Clear()
	test.GetU64Arr().Add(624)
	test.GetI32Map().PutOne(pbgen.NewTestI32MapSync().SetId(344).SetAddition("i32map2"))
	test.GetStrMap().PutOne(pbgen.NewTestStringMapSync().SetId("smm").SetAddition("sm1"))
	test.GetI64Map().PutOne(pbgen.NewTestI64MapSync().SetId(-64).SetAddition("i64map2"))
	test.GetBoolMap().PutOne(pbgen.NewTestBoolMapSync().SetId(false).SetAddition("i32map3"))
	test.GetU64Map().PutOne(pbgen.NewTestU64MapSync().SetId(624).SetAddition("i32map1"))
	test.GetU32Map().PutOne(pbgen.NewTestU32MapSync().SetId(32).SetAddition("i32map8"))
	test.GetU64Map().Remove(640)
}
