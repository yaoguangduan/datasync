package gen

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/yaoguangduan/protosync/pbgenv2"
	"github.com/yaoguangduan/protosync/syncdep"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protopack"
	"log"
	"testing"
)

func TestAllTypeMergePB(t *testing.T) {
	test := fullTestData()
	test.FlushDirty(false)
	another := fullTestData()
	another.FlushDirty(false)
	modifyAll(test)
	ts := &pbgenv2.Test{}
	test.MergeDirtyToPb(ts)
	t.Log(ts)
	another.MergeDirtyFromPb(ts)
	//修改后的数据 = 原始数据 + 脏数据
	testD := &pbgenv2.Test{}
	anotherD := &pbgenv2.Test{}
	test.CopyToPb(testD)
	t.Log(testD)
	another.CopyToPb(anotherD)
	t.Log(anotherD)
	assert.Equal(t, test, another)
}
func TestCopy(t *testing.T) {
	test := fullTestData()
	test.FlushDirty(false)
	ts := pbgenv2.Test{}
	test.CopyToPb(&ts)

	testNew := pbgenv2.NewTestSync()
	testNew.CopyFromPb(&ts)
	testNew.FlushDirty(false)

	assert.Equal(t, test, testNew)
}

func TestClear(t *testing.T) {
	test := fullTestData()
	test.FlushDirty(false)
	test.Clear()
	rs := pbgenv2.Test{}
	test.MergeDirtyToPb(&rs)
	t.Log(&rs)

	testNew := pbgenv2.NewTestSync()
	testNew.MergeDirtyFromPb(&rs)

	test.FlushDirty(false)
	testNew.FlushDirty(false)
	//assert.Equal(t, test, testNew)
	ti := &pbgenv2.Test{}
	msg := ti.ProtoReflect()

	bys := lo.Must(proto.Marshal(ti))

	tii := &pbgenv2.Test{}
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

func TestAllTypeMergeBytes(t *testing.T) {
	test := fullTestData()
	bytes := test.MergeDirtyToBytes()

	testSync := &pbgenv2.Test{}
	err := proto.Unmarshal(bytes, testSync)
	assert.NoError(t, err)
	t.Log(testSync)

	testSyncPb := &pbgenv2.Test{}
	test.MergeDirtyToPb(testSyncPb)
	t.Log(testSyncPb)

	testMsgUn := protopack.Message{}
	testMsgUn.Unmarshal(testSync.ProtoReflect().GetUnknown())
	testMsgUnNew := protopack.Message{}
	testMsgUnNew.Unmarshal(testSyncPb.ProtoReflect().GetUnknown())
	log.Println(testMsgUn)
	log.Println(testMsgUnNew)
	assert.Equal(t, testSync.ProtoReflect().GetUnknown(), testSyncPb.ProtoReflect().GetUnknown())

}

func TestAllTypeMergeFromBytes(t *testing.T) {
	test := fullTestData()
	bytes := test.MergeDirtyToBytes()

	testSync := &pbgenv2.Test{}
	err := proto.Unmarshal(bytes, testSync)
	assert.NoError(t, err)

	testNew := pbgenv2.NewTestSync()
	message := protopack.Message{}
	message.Unmarshal(bytes)
	t.Log(message)
	testNew.MergeDirtyFromBytes(bytes)

	testD := &pbgenv2.Test{}
	testNewD := &pbgenv2.Test{}
	test.CopyToPb(testD)
	testNew.CopyToPb(testNewD)
	t.Log(testSync)
	t.Log(testD)
	t.Log(testNewD)
	assert.Equal(t, test, testNew)

}

func TestMerge(t *testing.T) {
	ps := pbgenv2.NewPersonSync()
	ps.GetActions().Put("act", pbgenv2.NewActionInfoSync().SetTime(23132))
	ps.FlushDirty(false)
	ps.GetActions().Remove("act")
	r := &pbgenv2.Person{}
	ps.MergeDirtyToPb(r)

	fmt.Println(protojson.Format(r))
	fmt.Println(r.ProtoReflect().GetUnknown())
	bys, err := proto.Marshal(r)
	if err != nil {
		panic(err)
	}
	rr := &pbgenv2.Person{}
	err = proto.Unmarshal(bys, rr)
	if err != nil {
		panic(err)
	}
	fmt.Println(protojson.Format(rr))
	fmt.Println(rr.ProtoReflect().GetUnknown())
}

func TestZeroSet(t *testing.T) {
	action := pbgenv2.NewActionInfoSync()
	action.SetDetail("onestring")

	var r = &pbgenv2.ActionInfo{}
	action.MergeDirtyToPb(r)
	t.Log(r)
	assert.Equal(t, action.GetDetail(), r.GetDetail())
	assert.Equal(t, action.GetTime(), r.GetTime())

	action.FlushDirty(false)
	r = &pbgenv2.ActionInfo{}
	action.MergeDirtyToPb(r)
	t.Log(r)
	assert.Equal(t, "", r.GetDetail())
	assert.Equal(t, int64(0), r.GetTime())

	action.SetAct("q")
	r = &pbgenv2.ActionInfo{}
	action.MergeDirtyToPb(r)
	t.Log(r)

	//通过bitfield避免字段有值 再设置成空值，这个时候merge，无法将原始字段设置为空了
	action.SetAct("new val")
	action.FlushDirty(false)
	action.MergeDirtyFromPb(r)
	assert.Equal(t, "q", action.GetAct())

	p := pbgenv2.NewPersonSync()
	p.SetIsGirl(false)
	var ps = pbgenv2.Person{}
	p.MergeDirtyToPb(&ps)
	p.SetIsGirl(true)
	p.FlushDirty(false)

	var psc = pbgenv2.Person{}
	p.MergeDirtyToPb(&psc)

	t.Log(&ps)
	p.MergeDirtyFromPb(&ps)
	assert.True(t, p.GetIsGirl())
}
func TestList(t *testing.T) {
	p := pbgenv2.NewPersonSync()
	p.GetFavor().Add("apple")
	p.GetFavor().Add("bnn")

	var ps1 = pbgenv2.Person{}
	p.MergeDirtyToPb(&ps1)
	t.Log(prototext.Format(&ps1))
	t.Log(1 << 5)
	assert.Equal(t, []string{"apple", "bnn"}, ps1.GetFavor())

	p.GetFavor().Add("bnn1")
	p.GetFavor().Clear()
	ps1 = pbgenv2.Person{}
	p.MergeDirtyToPb(&ps1)
	t.Log(prototext.Format(&ps1))
	p.FlushDirty(false)
	p.MergeDirtyFromPb(&ps1)
	t.Log(p.GetFavor())
}

func mockPersonData() pbgenv2.PersonSync {
	p := pbgenv2.NewPersonSync()
	p.SetIsGirl(true).SetName("john").SetAge(22).GetFavor().Add("basket")
	p.GetDetail().SetMoney(2912).SetAddress("bj")
	p.GetLoveSeq().Add(pbgenv2.ColorType_Green)
	p.GetLoveSeq().Add(pbgenv2.ColorType_Red)
	a := pbgenv2.NewActionInfoSync()
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
	dirty1 := &pbgenv2.Person{}
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
	a := pbgenv2.NewActionInfoSync()
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
	aa := pbgenv2.NewActionInfoSync()
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

	pd := &pbgenv2.Person{}
	pod := &pbgenv2.Person{}
	p.CopyToPb(pd)
	po.CopyToPb(pod)
	assert.Equal(t, pd, pod)
}

func TestBytes(t *testing.T) {
	p := pbgenv2.NewPersonSync()
	p.SetName("name")
	p.SetData([]byte("hello-world"))
	p.GetDetail().SetMoney(223344)
	p.GetFavor().Add("sleep")
	p.GetLoveSeq().Add(pbgenv2.ColorType_Green)
	p.GetLoveSeq().Add(pbgenv2.ColorType_Red)

	a := pbgenv2.NewActionInfoSync().SetTime(111)
	p.GetActions().Put("act", a)

	bytes := p.MergeDirtyToBytes()

	pc := pbgenv2.Person{}
	err := proto.Unmarshal(bytes, &pc)
	if err != nil {
		panic(err)
	}
	t.Log(&pc)

	assert.Equal(t, "hello-world", string(pc.GetData()))

}

func fullTestData() *pbgenv2.TestSync {

	test := pbgenv2.NewTestSync()
	test.SetId(-32).SetI64(-64).SetU64(64).SetU32(32).SetF32(12.23).SetF64(64.23).SetStr("str").SetB(true).SetObj(pbgenv2.NewPersonSync().SetAge(11111))
	test.GetStrArr().Add("arr1")
	test.GetStrArr().Add("arr2")
	test.GetEnumArr().Add(pbgenv2.ColorType_Blue)
	test.GetBoolArr().Add(false)
	test.GetI32Arr().Add(-32)
	test.GetI32Arr().Add(-16)
	test.GetU32Arr().Add(32)
	test.GetU32Arr().Add(16)
	test.GetI64Arr().Add(-64)
	test.GetU64Arr().Add(64)
	test.GetI32Map().PutOne(pbgenv2.NewTestI32MapSync().SetId(-23).SetAddition("i32map"))
	test.GetStrMap().PutOne(pbgenv2.NewTestStringMapSync().SetId("sm").SetAddition("sm"))
	test.GetI64Map().PutOne(pbgenv2.NewTestI64MapSync().SetId(-64).SetAddition("i64map"))
	test.GetBoolMap().PutOne(pbgenv2.NewTestBoolMapSync().SetId(true).SetAddition("i32map"))
	test.GetU64Map().PutOne(pbgenv2.NewTestU64MapSync().SetId(64).SetAddition("i32map"))
	test.GetU64Map().PutOne(pbgenv2.NewTestU64MapSync().SetId(640).SetAddition("i32map2"))
	test.GetU32Map().PutOne(pbgenv2.NewTestU32MapSync().SetId(32).SetAddition("i32map"))
	return test
}
func modifyAll(test *pbgenv2.TestSync) {
	test.SetId(-312).SetI64(-624).SetU64(694).SetU32(328).SetStr("stro").SetB(false).SetObj(pbgenv2.NewPersonSync().SetName("11111").SetAge(1))
	test.GetStrArr().Add("arr12")
	test.GetEnumArr().Add(pbgenv2.ColorType_Blue)
	test.GetBoolArr().Add(false)
	test.GetI32Arr().Clear()
	test.GetU32Arr().Add(132)
	test.GetU32Arr().Add(126)
	test.GetI64Arr().Clear()
	test.GetU64Arr().Add(624)
	test.GetI32Map().PutOne(pbgenv2.NewTestI32MapSync().SetId(344).SetAddition("i32map2"))
	test.GetStrMap().PutOne(pbgenv2.NewTestStringMapSync().SetId("smm").SetAddition("sm1"))
	test.GetI64Map().PutOne(pbgenv2.NewTestI64MapSync().SetId(-64).SetAddition("i64map2"))
	test.GetBoolMap().PutOne(pbgenv2.NewTestBoolMapSync().SetId(false).SetAddition("i32map3"))
	test.GetU64Map().PutOne(pbgenv2.NewTestU64MapSync().SetId(624).SetAddition("i32map1"))
	test.GetU32Map().PutOne(pbgenv2.NewTestU32MapSync().SetId(32).SetAddition("i32map8"))
	test.GetU64Map().Remove(640)
}

func TestDirtyOp(t *testing.T) {
	test := pbgenv2.NewTestSync()
	test.GetObj().GetDetail().SetMoney(12)
	testPb := pbgenv2.Test{}
	test.MergeDirtyToPb(&testPb)
	fmt.Println(protojson.Format(&testPb))
	testCopy := pbgenv2.NewTestSync()
	testCopy.MergeDirtyFromPb(&testPb)

	testResult := pbgenv2.Test{}
	testCopy.CopyToPb(&testResult)
	fmt.Println(protojson.Format(&testResult))

	testDirty := pbgenv2.NewTestSync()
	testDirty.MergeDirtyFromPb(&testResult)
	testDirty.GetI32Map().Put(12, pbgenv2.NewTestI32MapSync().SetId(12).SetAddition("test"))
	testDirty.FlushDirty(false)
	testDirty.GetI32Map().Remove(12)
	dirtyResult := pbgenv2.Test{}
	testDirty.MergeDirtyToPb(&dirtyResult)
	pm := (&dirtyResult).ProtoReflect()
	raw := syncdep.ToRawMessage(pm.GetUnknown())
	for _, rf := range raw.RawFields {
		for _, f := range rf {
			v, n := protowire.ConsumeVarint(f.Bytes)
			fmt.Println(f.Number, v, n, len(f.Bytes))
		}
	}
	fmt.Println(raw)
}
