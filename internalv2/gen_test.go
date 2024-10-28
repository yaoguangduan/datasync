package internalv2

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/yaoguangduan/protosync/pbgen"
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

func TestAllTypeMergeBytes(t *testing.T) {
	test := fullTestData()
	bytes := test.MergeDirtyToBytes()

	testSync := &pbgen.Test{}
	err := proto.Unmarshal(bytes, testSync)
	t.Log(testSync)
	assert.NoError(t, err)

	testSyncPb := &pbgen.Test{}
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

	testSync := &pbgen.Test{}
	err := proto.Unmarshal(bytes, testSync)
	assert.NoError(t, err)

	testNew := pbgen.NewTestSync()
	message := protopack.Message{}
	message.Unmarshal(bytes)
	t.Log(message)
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

func TestMerge(t *testing.T) {
	ps := pbgen.NewPersonSync()
	ps.GetActions().Put(pbgen.NewActionInfoSync().SetTime(23132))
	ps.FlushDirty(false)
	ps.GetActions().Remove("act")
	r := &pbgen.Person{}
	ps.MergeDirtyToPb(r)

	fmt.Println(protojson.Format(r))
	fmt.Println(r.ProtoReflect().GetUnknown())
	bys, err := proto.Marshal(r)
	if err != nil {
		panic(err)
	}
	rr := &pbgen.Person{}
	err = proto.Unmarshal(bys, rr)
	if err != nil {
		panic(err)
	}
	fmt.Println(protojson.Format(rr))
	fmt.Println(rr.ProtoReflect().GetUnknown())
}

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
	p.GetActions().Put(a)
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
	p.GetActions().Put(a)
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
	p.GetActions().Put(aa)
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
	p.GetActions().Put(a)

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
	test.GetI32Map().Put(pbgen.NewTestI32MapSync().SetId(-23).SetAddition("i32map"))
	test.GetStrMap().Put(pbgen.NewTestStringMapSync().SetId("sm").SetAddition("sm"))
	test.GetI64Map().Put(pbgen.NewTestI64MapSync().SetId(-64).SetAddition("i64map"))
	test.GetBoolMap().Put(pbgen.NewTestBoolMapSync().SetId(true).SetAddition("i32map"))
	test.GetU64Map().Put(pbgen.NewTestU64MapSync().SetId(64).SetAddition("i32map"))
	test.GetU64Map().Put(pbgen.NewTestU64MapSync().SetId(640).SetAddition("i32map2"))
	test.GetU32Map().Put(pbgen.NewTestU32MapSync().SetId(32).SetAddition("i32map"))
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
	test.GetI32Map().Put(pbgen.NewTestI32MapSync().SetId(344).SetAddition("i32map2"))
	test.GetStrMap().Put(pbgen.NewTestStringMapSync().SetId("smm").SetAddition("sm1"))
	test.GetI64Map().Put(pbgen.NewTestI64MapSync().SetId(-64).SetAddition("i64map2"))
	test.GetBoolMap().Put(pbgen.NewTestBoolMapSync().SetId(false).SetAddition("i32map3"))
	test.GetU64Map().Put(pbgen.NewTestU64MapSync().SetId(624).SetAddition("i32map1"))
	test.GetU32Map().Put(pbgen.NewTestU32MapSync().SetId(32).SetAddition("i32map8"))
	test.GetU64Map().Remove(640)
}

func TestDirtyOp(t *testing.T) {
	test := pbgen.NewTestSync()
	test.GetObj().GetDetail().SetMoney(12)
	testPb := pbgen.Test{}
	test.MergeDirtyToPb(&testPb)
	fmt.Println(protojson.Format(&testPb))
	testCopy := pbgen.NewTestSync()
	testCopy.MergeDirtyFromPb(&testPb)

	testResult := pbgen.Test{}
	testCopy.CopyToPb(&testResult)
	fmt.Println(protojson.Format(&testResult))

	testDirty := pbgen.NewTestSync()
	testDirty.MergeDirtyFromPb(&testResult)
	testDirty.GetI32Map().Put(pbgen.NewTestI32MapSync().SetId(12).SetAddition("test"))
	testDirty.FlushDirty(false)
	testDirty.GetI32Map().Remove(12)
	dirtyResult := pbgen.Test{}
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
