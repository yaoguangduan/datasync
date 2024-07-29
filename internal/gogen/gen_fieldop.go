package gogen

import (
	"fmt"
	"github.com/yaoguangduan/datasync/internal/gen"
)

func generateFieldOperateFunc(fw *gen.FileWriter, sfd gen.SyncDef, msg gen.SyncMsgOrEnumDef) {
	for _, field := range msg.MsgFields {
		//GetXXX
		fw.PLF("func (x *%s) Get%s() %s {", msg.SyncName, field.CapitalName, field.GoName())
		if field.IsMsg() {
			fw.PLF("if x.%s == nil {", field.Name)
			fw.PLF("x.%s = New%s()", field.Name, field.MsgOrEnumRef.SyncName)
			fw.PLF("x.%s.SetParent(x,x.%s)", field.Name, goFieldIndexName(field.Name))
			fw.PL("}")
		}
		if field.IsMap() {
			fw.PLF("if x.%s == nil {", field.Name)
			fw.PLF("x.%s = syncdep.NewMapSync[%s,*%s]()", field.Name, field.MapKeyKind, field.MsgOrEnumRef.SyncName)
			fw.PLF("x.%s.SetParent(x,x.%s)", field.Name, goFieldIndexName(field.Name))
			fw.PL("}")
		}
		if field.IsList() {
			fw.PLF("if x.%s == nil {", field.Name)
			fw.PLF("x.%s = %s", field.Name, fmt.Sprintf("syncdep.NewArraySync[%s]()", gen.FloatConvert(field.ListType)))
			fw.PLF("x.%s.SetParent(x,x.%s)", field.Name, goFieldIndexName(field.Name))
			fw.PL("}")
		}
		fw.PLF("return x.%s", field.Name)
		fw.PL("}")

		//SetXXX
		if !field.IsList() && !field.IsMap() {
			if field.IsPrimary() || field.IsEnum() {
				fw.PLF("func (x *%s) Set%s(v %s) *%s{", msg.SyncName, field.CapitalName, field.GoName(), msg.SyncName)
				fw.PLF("if x.%s == v {", field.Name)
				fw.PLF("return x")
				fw.PLF("}")
			} else {
				fw.PLF("func (x *%s) Set%s(v %s) *%s{", msg.SyncName, field.CapitalName, field.GoName(), msg.SyncName)
			}

			if field.IsMsg() {
				fw.PLF("if v != nil {")
				fw.PLF("v.SetParent(x,x.%s)", goFieldIndexName(field.Name))
				fw.PL("}")
				fw.PLF("if x.%s != nil {", field.Name)
				fw.PLF("x.%s.SetParent(nil,-1)", field.Name)
				fw.PL("}")
			}
			fw.PLF("x.%s = v", field.Name)
			fw.PLF("x.set%sDirty(true,false)", field.CapitalName)
			fw.PLF("return x")

			fw.PL("}")
		}
	}
}
