package gogen

import (
	"github.com/samber/lo"
	"github.com/yaoguangduan/datasync/internal/gen"
	"github.com/yaoguangduan/datasync/internal/proto_file_gen"
	"github.com/yaoguangduan/datasync/syncdep"
)

func generateFuncMergeDirtyToPb(fw *gen.FileWriter, msg gen.SyncMsgOrEnumDef) {
	fw.PLF("func (x *%s) MergeDirtyToPb(r *%s) {", msg.SyncName, msg.Name)
	for _, field := range msg.MsgFields {
		if field.IsPrimary() || field.IsEnum() {
			fw.PLF("if x.is%sDirty() {", field.CapitalName)
			fw.PLF("r.Set%s(x.%s)", field.CapitalName, field.Name)
			fw.PL("}")
		}
		if field.IsBytes() {
			fw.PLF("if x.is%sDirty() {", field.CapitalName)
			fw.PLF("r.Set%s(slices.Clone(x.%s))", field.CapitalName, field.Name)
			fw.PL("}")
		}
		if field.IsMsg() {
			fw.PLF("if x.is%sDirty() {", field.CapitalName)
			fw.PLF("if r.%s == nil {", field.CapitalName)
			fw.PLF("r.%s = &%s{}", field.CapitalName, (*field.MsgOrEnumRef).Name)
			fw.PL("}")
			fw.PLF("x.%s.MergeDirtyToPb(r.%s)", field.Name, field.CapitalName)
			fw.PL("}")
		}
		if field.IsList() {
			fw.PLF("if x.is%sDirty() {", field.CapitalName)
			fw.PLF("count := x.%s.Len()", field.Name)
			fw.PLF("r.%s = make([]%s,0)", field.CapitalName, lo.If(field.ListType != "enum", gen.FloatConvert(field.ListType)).Else(field.MsgOrEnumRef.Name))
			fw.PLF("if count > 0 {")
			fw.PLF("r.%s = false", proto_file_gen.ProtoClearedName(field.CapitalName))
			fw.PLF("r.%s = append(r.%s,x.%s.ValueView()...)", field.CapitalName, field.CapitalName, field.Name)
			fw.PLF("} else {")
			fw.PLF("r.%s = true", proto_file_gen.ProtoClearedName(field.CapitalName))
			fw.PLF("}")
			fw.PL("}")
		}
		if field.IsMap() {
			fw.PLF("if x.is%sDirty() {", field.CapitalName)
			fw.PLF("updated := make([]%s,0)", field.MapKeyKind)
			fw.PLF("if r.%s != nil {", field.CapitalName)
			fw.PLF("for k := range r.%s {", field.CapitalName)

			fw.PLF("if x.%s.ContainDeleted(k) {", field.Name)
			fw.PLF("delete(r.%s,k)", field.CapitalName)
			fw.PLF("}")
			fw.PLF("if x.%s.ContainDirtied(k) {", field.Name)
			fw.PLF("updated = append(updated,k)")

			fw.PLF("tmp := x.%s.Get(k)", field.Name)
			fw.PLF("if r.%s[k] == nil{", field.CapitalName)
			fw.PLF("r.%s[k] = &%s{}", field.CapitalName, field.MsgOrEnumRef.Name)
			fw.PLF("}")
			fw.PLF("tmp.MergeDirtyToPb(r.%s[k])", field.CapitalName)

			fw.PLF("}")
			fw.PLF("}")
			fw.PLF("} else {")
			fw.PLF("r.%s = make(map[%s]*%s)", field.CapitalName, field.MapKeyKind, field.MsgOrEnumRef.Name)
			fw.PLF("}")

			fw.PLF("for k := range x.%s.Dirtied() {", field.Name)
			fw.PLF("if !slices.Contains(updated,k) {")
			fw.PLF("tmp := x.%s.Get(k)", field.Name)
			fw.PLF("if r.%s[k] == nil {", field.CapitalName)
			fw.PLF("r.%s[k] = &%s{}", field.CapitalName, field.MsgOrEnumRef.Name)
			fw.PLF("}")
			fw.PLF("tmp.MergeDirtyToPb(r.%s[k])", field.CapitalName)

			fw.PLF("}")
			fw.PLF("}")

			fw.PLF("if r.%s == nil && len(x.%s.Deleted()) > 0{", proto_file_gen.ProtoDeletedName(field.CapitalName), field.Name)
			fw.PLF("r.%s = make([]%s,0)", proto_file_gen.ProtoDeletedName(field.CapitalName), field.MapKeyKind)
			fw.PLF("}")

			fw.PLF("for k := range x.%s.Deleted() {", field.Name)
			fw.PLF("if !slices.Contains(r.%s,k) {", proto_file_gen.ProtoDeletedName(field.CapitalName))
			fw.PLF("r.%s = append(r.%s,k)", proto_file_gen.ProtoDeletedName(field.CapitalName), proto_file_gen.ProtoDeletedName(field.CapitalName))
			fw.PLF("}")
			fw.PLF("}")

			fw.PL("}")
		}
	}
	fw.PL("}")
}

func generateFuncMergeDirtyToBytes(fw *gen.FileWriter, msg gen.SyncMsgOrEnumDef) {
	fw.PLF("func (x *%s) MergeDirtyToBytes() []byte {", msg.SyncName)
	fw.PLF("var buf []byte")
	for _, field := range msg.MsgFields {
		fw.PLF("if x.is%sDirty() {", field.CapitalName)

		if gen.IsBuildInType(field.Kind) || field.IsEnum() {
			fw.PLF("buf = protowire.AppendTag(buf, %d, %v)", field.Number, syncdep.FieldTypeToWireType(field))
			if field.IsEnum() {
				fw.PLF("buf = syncdep.AppendFieldValue(buf, int32(x.%s))", field.Name)
			} else {
				fw.PLF("buf = syncdep.AppendFieldValue(buf, x.%s)", field.Name)
			}
		}
		if field.IsMsg() {
			fw.PLF("if x.%s != nil {", field.Name)
			fw.PLF("bytes := x.%s.MergeDirtyToBytes()", field.Name)
			fw.PLF("buf = protowire.AppendTag(buf, %d, %v)", field.Number, syncdep.FieldTypeToWireType(field))
			fw.PLF("buf = syncdep.AppendFieldValue(buf, bytes)")
			fw.PL("}")
		}
		if field.IsList() {
			fw.PLF("if x.%s != nil && x.%s.Len() > 0 {", field.Name, field.Name)
			fw.PLF("buf = protowire.AppendTag(buf, %d, protowire.VarintType)", field.ProtoClearNumber())
			fw.PLF("buf = syncdep.AppendFieldValue(buf, false)")
			if field.ListType == "string" {
				fw.PLF("x.%s.Each(func(i int,v string) bool {", field.Name)
				fw.PLF("buf = protowire.AppendTag(buf, %d, %v)", field.Number, syncdep.FieldTypeToWireType(field))
				fw.PLF("buf = protowire.AppendString(buf, v)")
				fw.PLF("return true")
				fw.PLF("})")
			} else {
				fw.PLF("var packedBuf []byte")
				fw.PLF("x.%s.Each(func(i int,v %s)bool {", field.Name, gen.FloatConvert(field.ListType))
				if field.MsgOrEnumRef != nil {
					fw.PLF("packedBuf = syncdep.AppendFieldValue(packedBuf, int32(v))")
				} else {
					fw.PLF("packedBuf = syncdep.AppendFieldValue(packedBuf, v)")
				}
				fw.PLF("return true")
				fw.PLF("})")
				fw.PLF("buf = protowire.AppendTag(buf, %d, %v)", field.Number, syncdep.FieldTypeToWireType(field))
				fw.PLF("buf = protowire.AppendBytes(buf, packedBuf)")
			}
			fw.PL("} else {")
			fw.PLF("buf = protowire.AppendTag(buf, %d, protowire.VarintType)", field.ProtoClearNumber())
			fw.PLF("buf = syncdep.AppendFieldValue(buf, true)")
			fw.PLF("}")
		}
		if field.IsMap() {
			fw.PLF("if len(x.%s.Deleted()) > 0 {", field.Name)
			fw.PLF("deleted := x.%s.Deleted()", field.Name)

			if field.MapKeyKind == "string" {
				fw.PLF("for del := range deleted {")
				fw.PLF("buf = protowire.AppendTag(buf, %d, protowire.BytesType)", field.ProtoDelNumber())
				fw.PLF("buf = protowire.AppendString(buf, del)")
				fw.PLF("}")
			} else {
				fw.PLF("var packedBuf []byte")
				fw.PLF("for v := range deleted {")
				if field.MapKeyKind == "enum" {
					fw.PLF("packedBuf = syncdep.AppendFieldValue(packedBuf, int32(v))")
				} else {
					fw.PLF("packedBuf = syncdep.AppendFieldValue(packedBuf, v)")
				}
				fw.PLF("}")
				fw.PLF("buf = protowire.AppendTag(buf, %d, %v)", field.ProtoDelNumber(), syncdep.FieldTypeToWireType(field))
				fw.PLF("buf = protowire.AppendBytes(buf, packedBuf)")
			}

			fw.PLF("}")
			fw.PLF("if x.%s.Len() > 0 {", field.Name)
			fw.PLF("x.%s.Each(func(k %s,v *%s) bool {", field.Name, field.MapKeyKind, field.MsgOrEnumRef.SyncName)
			fw.PLF("if !x.%s.ContainDirtied(k) {", field.Name)
			fw.PLF("return true")
			fw.PLF("}")
			fw.PLF("buf = syncdep.AppendMapFieldKeyValue(buf,%d, k, v.MergeDirtyToBytes())", field.Number)
			fw.PLF("return true")
			fw.PLF("})")
			fw.PL("}")
		}
		fw.PL("}")
	}
	fw.PLF("return buf")
	fw.PLF("}")
}

func generateFuncMergeDirtyFromBytes(fw *gen.FileWriter, msg gen.SyncMsgOrEnumDef) {
	fw.PLF("func (x *%s) MergeDirtyFromBytes(buf []byte) *%s{", msg.SyncName, msg.SyncName)
	fw.PLF("fds := syncdep.PreParseProtoBytes(buf)")

	var hasArrOrMap = lo.CountBy(msg.MsgFields, func(item gen.SyncFieldDef) bool {
		return item.IsMap() || item.IsList()
	}) > 0
	if hasArrOrMap {
		fw.PLF("for _,rawF := range fds.Values {")
		fw.PLF("switch rawF.Number {")
		for _, field := range msg.MsgFields {
			if field.IsList() {
				fw.PLF("case %d:", field.ProtoClearNumber())
				fw.PLF("if rawF.Value.(uint64) > 0 {")
				fw.PLF("x.%s.Clear()", field.Name)
				fw.PLF("}")
			}
			if field.IsMap() {
				fw.PLF("case %d:", field.Number+1000)
				fw.PLF("if x.%s != nil {", field.Name)
				if field.MapKeyKind == "string" {
					fw.PLF("x.%s.Remove(syncdep.Bys2Str(rawF.Value.([]byte)))", field.Name)
				} else {
					fw.PLF("x.%s.RemoveAll(syncdep.MustParseVarintArr[%s](&rawF))", field.Name, field.MapKeyKind)
				}
				fw.PLF("}")
			}
		}
		fw.PLF("}")
		fw.PLF("}")
	}

	fw.PLF("for _,rawF := range fds.Values {")
	fw.PLF("switch rawF.Number {")
	for _, field := range msg.MsgFields {
		fw.PLF("case %d:", field.Number)
		if field.IsPrimary() || field.IsEnum() || field.IsBytes() {
			if field.Kind == "float" || field.Kind == "double" {
				fw.PLF("x.Set%s(rawF.Value.(%s))", field.CapitalName, gen.FloatConvert(field.Kind))
			} else if field.Kind == "bool" {
				fw.PLF("x.Set%s(rawF.Value.(uint64) > 0)", field.CapitalName)
			} else if field.Kind == "enum" {
				fw.PLF("x.Set%s(%s(rawF.Value.(uint64)))", field.CapitalName, field.MsgOrEnumRef.Name)
			} else if field.Kind == "string" {
				fw.PLF("x.Set%s(syncdep.Bys2Str(rawF.Value.([]byte)))", field.CapitalName)
			} else if field.Kind == "bytes" {
				fw.PLF("x.Set%s(rawF.Value.([]byte))", field.CapitalName)
			} else {
				fw.PLF("x.Set%s(%s(rawF.Value.(uint64)))", field.CapitalName, field.Kind)
			}
		}
		if field.IsMsg() {
			fw.PLF("x.Get%s().MergeDirtyFromBytes(rawF.Value.([]byte))", field.CapitalName)
		}
		if field.IsList() {
			if field.ListType == "string" {
				fw.PLF("x.Get%s().Add(syncdep.Bys2Str(rawF.Value.([]byte)))", field.CapitalName)
			} else {
				fw.PLF("tmp := rawF.Value.([]byte)")
				fw.PLF("for len(tmp) > 0 {")
				if field.ListType == "float" {
					fw.PLF("val, n := protowire.ConsumeFixed32(tmp)")
					fw.PLF("if n <0 {")
					fw.PLF("panic(n)")
					fw.PLF("}")
					fw.PLF("tmp = tmp[n:]")
					fw.PLF("x.Get%s().Add(math.Float32frombits(val))", field.CapitalName)
				} else if field.ListType == "double" {
					fw.PLF("val, n := protowire.ConsumeFixed64(tmp)")
					fw.PLF("if n <0 {")
					fw.PLF("panic(n)")
					fw.PLF("}")
					fw.PLF("tmp = tmp[n:]")
					fw.PLF("x.Get%s().Add(math.Float64frombits(val))", field.CapitalName)

				} else if field.ListType == "enum" {
					fw.PLF("val, n := protowire.ConsumeVarint(tmp)")
					fw.PLF("if n <0 {")
					fw.PLF("panic(n)")
					fw.PLF("}")
					fw.PLF("tmp = tmp[n:]")
					fw.PLF("x.Get%s().Add(%s(val))", field.CapitalName, field.MsgOrEnumRef.SyncName)
				} else if field.ListType == "bool" {
					fw.PLF("val, n := protowire.ConsumeVarint(tmp)")
					fw.PLF("if n <0 {")
					fw.PLF("panic(n)")
					fw.PLF("}")
					fw.PLF("tmp = tmp[n:]")
					fw.PLF("x.Get%s().Add(val> 0)", field.CapitalName)
				} else {
					fw.PLF("val, n := protowire.ConsumeVarint(tmp)")
					fw.PLF("if n <0 {")
					fw.PLF("panic(n)")
					fw.PLF("}")
					fw.PLF("tmp = tmp[n:]")
					fw.PLF("x.Get%s().Add(%s(val))", field.CapitalName, field.ListType)

				}
				fw.PLF("}")
			}
		}
		if field.IsMap() {
			fw.PLF("mapKV := syncdep.PreParseProtoBytes(rawF.Value.([]byte)).Values")
			fw.PLF("k := syncdep.GetMapKey[%s](&mapKV[0])", field.MapKeyKind)
			fw.PLF("var tmp = x.Get%s().Get(k)", field.CapitalName)
			fw.PLF("if tmp == nil {")
			fw.PLF("tmp = New%s()", field.MsgOrEnumRef.SyncName)
			fw.PLF("}")
			fw.PLF("tmp.MergeDirtyFromBytes(mapKV[1].Value.([]byte))")
			fw.PLF("x.Get%s().Put(k,tmp)", field.CapitalName)
		}
	}
	fw.PLF("}")
	fw.PLF("}")
	fw.PLF("return x")
	fw.PLF("}")
}

func generateFuncMergeDirtyFromPb(fw *gen.FileWriter, msg gen.SyncMsgOrEnumDef) {
	fw.PLF("func (x *%s) MergeDirtyFromPb(r *%s) {", msg.SyncName, msg.Name)
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
		if field.IsList() {
			fw.PLF("if len(r.%s) > 0 || r.%s {", field.CapitalName, proto_file_gen.ProtoClearedName(field.CapitalName))
			fw.PLF("x.Get%s().Clear()", field.CapitalName)
			fw.PLF("x.%s.AddAll(r.%s)", field.Name, field.CapitalName)
			fw.PLF("}")
		}
		if field.IsMsg() {
			fw.PLF("if r.%s != nil {", field.CapitalName)
			fw.PLF("x.Get%s().MergeDirtyFromPb(r.%s)", field.CapitalName, field.CapitalName)
			fw.PLF("}")
		}
		if field.IsMap() {
			fw.PLF("if x.%s != nil {", field.Name)
			fw.PLF("x.Get%s().RemoveAll(r.%s)", field.CapitalName, proto_file_gen.ProtoDeletedName(field.CapitalName))
			fw.PLF("}")

			fw.PLF("for k,v := range r.%s {", field.CapitalName)
			fw.PLF("var tmp = x.Get%s().Get(k)", field.CapitalName)
			fw.PLF("if tmp == nil {")
			fw.PLF("tmp = New%s()", field.MsgOrEnumRef.SyncName)
			fw.PLF("tmp.MergeDirtyFromPb(v)")
			fw.PLF("x.Get%s().Put(k,tmp)", field.CapitalName)
			fw.PLF("} else {")
			fw.PLF("tmp.MergeDirtyFromPb(v)")
			fw.PLF("}")
			fw.PLF("}")
		}
	}
	fw.PLF("}")
}
