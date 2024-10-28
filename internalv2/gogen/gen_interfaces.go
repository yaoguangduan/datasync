package gogen

import (
	"github.com/yaoguangduan/protosync/internalv2/gen"
)

func generateFuncInterfaces(fw *gen.FileWriter, msg gen.SyncMsgOrEnumDef) {
	fw.PLF("func (x *%s) SetDirty(index int, dirty bool,sync syncdep.Sync) {", msg.SyncName)
	fw.PL("idx := index >> 3")
	fw.PL("off := index & 7")
	fw.PL("if dirty {")
	fw.PLF("x.%s[idx] = x.%s[idx] | ( 1 << off)", dirtyFieldName, dirtyFieldName)
	fw.PL("x.SetParentDirty()")
	fw.PL("} else {")
	fw.PLF("x.%s[idx] = x.%s[idx] & ^( 1 << off)", dirtyFieldName, dirtyFieldName)
	fw.PL("}")
	fw.PL("}")

	fw.PLF("func (x *%s) SetParentDirty() {", msg.SyncName)
	fw.PL("if x.parent != nil {")
	fw.PLF("x.parent.SetDirty(x.%s,true,x)", indexInParName)
	//fw.PLF("x.parent.SetParentDirty()"))
	fw.PL("}")
	fw.PL("}")

	fw.PLF("func (x *%s) SetParent(sync syncdep.Sync, idx int) {", msg.SyncName)
	fw.PLF("x.parent = sync")
	fw.PLF("x.%s = idx", indexInParName)
	fw.PL("}")

	fw.PLF("func (x *%s) FlushDirty(dirty bool) {", msg.SyncName)
	for _, field := range msg.MsgFields {
		fw.PLF("if dirty || x.is%sDirty() {", field.CapitalName)
		fw.PLF("x.set%sDirty(dirty,true)", field.CapitalName)
		fw.PL("}")
	}
	fw.PL("}")

	for _, field := range msg.MsgFields {
		fw.PLF("func (x *%s) set%sDirty(dirty bool,recur bool) {", msg.SyncName, field.CapitalName)
		fw.PLF("x.SetDirty(x.%s,dirty,x)", goFieldIndexName(field.Name))
		if !gen.IsBuildInType(field.Kind) && !field.IsEnum() {
			if field.IsMsg() || field.IsMap() {
				fw.PLF("if recur && x.%s != nil {", field.Name)
				fw.PLF("x.%s.FlushDirty(dirty)", field.Name)
				fw.PL("}")
			}
		}
		fw.PL("}")

		fw.PLF("func (x *%s) is%sDirty() bool{", msg.SyncName, field.CapitalName)
		fw.PLF("idx := x.%s >> 3", goFieldIndexName(field.Name))
		fw.PLF("off := x.%s & 7", goFieldIndexName(field.Name))
		fw.PLF("return (x.%s[idx] & (1 << off)) != 0", dirtyFieldName)
		fw.PL("}")
	}

	//Keys
	fw.PLF("func (x *%s) Key() interface{} {", msg.SyncName)
	if msg.MsgKey != nil {
		fw.PLF("return x.%s", msg.MsgKey.Name)
	} else {
		fw.PLF("return nil")
	}
	fw.PLF("}")

	fw.PLF("func (x *%s) SetKey(v interface{}) {", msg.SyncName)
	if msg.MsgKey == nil {
		fw.PLF("return")
	} else {
		fw.PLF("if x.parent != nil {")
		fw.PLF("if _,ok := x.parent.(*syncdep.MapSync[%s,*%s]); ok {", msg.MsgKey.Kind, msg.SyncName)
		fw.PLF("panic(\"%s key can not set\")", msg.SyncName)
		fw.PLF("}")
		fw.PLF("}")
		fw.PLF("x.%s = v.(%s)", msg.MsgKey.Name, msg.MsgKey.Kind)
	}
	fw.PLF("}")
}
