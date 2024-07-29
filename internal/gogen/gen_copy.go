package gogen

import (
	"github.com/yaoguangduan/datasync/internal/gen"
)

func generateFuncCopyFromPb(fw *gen.FileWriter, msg gen.SyncMsgOrEnumDef) {
	fw.PLF("func (x *%s) CopyFromPb(r *%s) *%s{", msg.SyncName, msg.Name, msg.SyncName)
	for _, field := range msg.MsgFields {
		if field.IsPrimary() || field.IsEnum() {
			fw.PLF("if r.%s != nil {", field.CapitalName)
			fw.PLF("x.Set%s(*r.%s)", field.CapitalName, field.CapitalName)
			fw.PLF("}")
		}
		if field.IsBytes() {
			fw.PLF("if len(r.%s) > 0 {", field.CapitalName)
			fw.PLF("x.Set%s(slices.Clone(r.%s))", field.CapitalName, field.CapitalName)
			fw.PLF("}")
		}
		if field.IsMsg() {
			fw.PLF("if r.%s != nil {", field.CapitalName)
			fw.PLF("x.Get%s().CopyFromPb(r.%s)", field.CapitalName, field.CapitalName)
			fw.PLF("}")
		}
		if field.IsList() {
			fw.PLF("if r.%s != nil {", field.CapitalName)
			fw.PLF("x.Get%s().AddAll(r.%s)", field.CapitalName, field.CapitalName)
			fw.PLF("}")
		}
		if field.IsMap() {
			fw.PLF("for k,v := range r.%s {", field.CapitalName)
			fw.PLF("if v != nil {")
			fw.PLF("vv := New%s()", field.MsgOrEnumRef.SyncName)
			fw.PLF("vv.CopyFromPb(v)")
			fw.PLF("x.Get%s().Put(k,vv)", field.CapitalName)
			fw.PLF("}")
			fw.PLF("}")
		}
	}
	fw.PLF("return x")
	fw.PLF("}")
}

func generateFuncCopyToPb(fw *gen.FileWriter, msg gen.SyncMsgOrEnumDef) {
	fw.PLF("func (x *%s) CopyToPb(r *%s) *%s {", msg.SyncName, msg.Name, msg.SyncName)
	for _, field := range msg.MsgFields {
		if field.IsPrimary() || field.IsEnum() {
			fw.PLF("r.Set%s(x.%s)", field.CapitalName, field.Name)
		}
		if field.IsBytes() {
			fw.PLF("r.Set%s(slices.Clone(x.%s))", field.CapitalName, field.Name)
		}
		if field.IsMsg() {
			fw.PLF("if x.%s != nil {", field.Name)
			fw.PLF("tmp := &%s{}", field.MsgOrEnumRef.Name)
			fw.PLF("x.%s.CopyToPb(tmp)", field.Name)
			fw.PLF("r.Set%s(tmp)", field.CapitalName)
			fw.PLF("}")
		}
		if field.IsList() {
			fw.PLF("if x.%s != nil && x.%s.Len() > 0 {", field.Name, field.Name)
			fw.PLF("r.Set%s(x.%s.ValueView())", field.CapitalName, field.Name)
			fw.PLF("}")
		}
		if field.IsMap() {
			fw.PLF("if x.%s != nil && x.%s.Len() > 0 {", field.Name, field.Name)
			fw.PLF("tmp := make(map[%s]*%s)", field.MapKeyKind, field.MsgOrEnumRef.Name)
			fw.PLF("x.%s.Each(func (k %s,v *%s) bool {", field.Name, field.MapKeyKind, field.MsgOrEnumRef.SyncName)
			fw.PLF("tmpV := &%s{}", field.MsgOrEnumRef.Name)
			fw.PLF("v.CopyToPb(tmpV)")
			fw.PLF("tmp[k] = tmpV")
			fw.PLF("return true")
			fw.PLF("})")
			fw.PLF("r.Set%s(tmp)", field.CapitalName)
			fw.PLF("}")
		}
	}
	fw.PLF("return x")
	fw.PLF("}")
}
