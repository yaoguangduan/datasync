package test

import (
	"fmt"
	"github.com/yaoguangduan/datasync/pbgen"
	"google.golang.org/protobuf/encoding/protojson"
	"testing"
)

func TestGen(t *testing.T) {
	p := pbgen.NewPerson()
	p.GetDetail().SetMoney(23)

	a1 := pbgen.NewActionInfo().SetAct("a1")
	p.GetActions().Put(a1.GetAct(), a1)

	a1.SetTime(222)
	a2 := pbgen.NewActionInfo().SetAct("a2")
	p.GetActions().Put(a2.GetAct(), a2)
	a2.SetTime(567)
	a3 := pbgen.NewActionInfo().SetAct("a3")
	p.GetActions().Put(a3.GetAct(), a3)
	a3.SetTime(789)

	ps := &pbgen.PersonSync{}
	p.CollectDirtyToPb(ps)
	marshal, err := protojson.Marshal(ps)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))

	p.FlushDirty(false)

	ps1 := &pbgen.PersonSync{}
	a1.SetTime(9999)
	p.GetActions().Get("a1").SetDetail("news")
	p.GetActions().Remove("a2")
	p.CollectDirtyToPb(ps1)
	marshal1, err := protojson.Marshal(ps1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal1))

	p.FlushDirty(false)
	ps2 := &pbgen.PersonSync{}
	p.CollectDirtyToPb(ps2)

	marshal2, err := protojson.Marshal(ps2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal2))
}

func TestDirty(t *testing.T) {
	d := make([]uint8, 2)
	index := 1
	idx := index >> 3
	off := index & 7
	d[idx] = d[idx] | (1 << off)
	fmt.Println(fmt.Sprintf("%b", d[0]))

	fmt.Println((d[idx] & (1 << off)) != 0)
}
