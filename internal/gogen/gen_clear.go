package gogen

import (
	"github.com/yaoguangduan/datasync/internal/gen"
)

func generateFuncClear(fw *gen.FileWriter, msg gen.SyncMsgOrEnumDef) {
	fw.PLF("func (x *%s) Clear() *%s {", msg.SyncName, msg.SyncName)
	for _, field := range msg.MsgFields {
		if gen.IsBuildInType(field.Kind) || field.IsEnum() {
			fw.PLF("x.Set%s(%v)", field.CapitalName, fieldDefaultVal(field))
		}
		if field.IsMsg() {
			fw.PLF("if x.%s != nil {", field.Name)
			fw.PLF("x.%s.Clear()", field.Name)
			fw.PLF("}")
		}
		if field.IsList() || field.IsMap() {
			fw.PLF("if x.%s != nil {", field.Name)
			fw.PLF("x.%s.Clear()", field.Name)
			fw.PLF("}")
		}
	}
	fw.PLF("return x")
	fw.PLF("}")
}
