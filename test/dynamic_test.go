package test

import (
	"fmt"
	"github.com/yaoguangduan/protosync/pbgen"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"testing"
)

func TestDynamic(t *testing.T) {

	act := pbgen.ActionInfoSync{}
	m := act.ProtoReflect()
	m.Set(m.Descriptor().Fields().ByName("act"), protoreflect.ValueOfString("fsadas"))
	m.Set(m.Descriptor().Fields().ByName("detail"), protoreflect.ValueOfString(""))

	fmt.Println("before : detail set:", m.Has(m.Descriptor().Fields().ByName("detail")))
	fmt.Println("before : act set:", m.Has(m.Descriptor().Fields().ByName("act")))
	marshal, err := proto.Marshal(m.Interface())
	if err != nil {
		panic(err)
	}
	sync := pbgen.ActionInfoSync{}
	mmm := sync.ProtoReflect()
	mm := mmm.Interface()
	err = proto.Unmarshal(marshal, mm)
	if err != nil {
		panic(err)
	}
	fmt.Println(prototext.Format(mm))
	fmt.Println("af : detail set:", mmm.Has(mmm.Descriptor().Fields().ByName("detail")))
	fmt.Println("af : act set:", mmm.Has(mmm.Descriptor().Fields().ByName("act")))

	p := pbgen.PersonSync{}
	//p.Actions = nil
	p.Actions["s"] = &pbgen.ActionInfoSync{}
	fmt.Println(p.Actions)

}
