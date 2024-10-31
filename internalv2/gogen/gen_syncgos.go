package gogen

import (
	"fmt"
	"github.com/yaoguangduan/protosync/internalv2/gen"
)

func generateProtoSyncOperateFunc(fw *gen.FileWriter, sfd gen.SyncDef, msg gen.SyncMsgOrEnumDef) {

	for _, field := range msg.MsgFields {
		if field.IsPrimary() || field.IsEnum() {
			fw.PLF("func (xs *%s) Set%s(v %s) {", msg.Name, field.CapitalName, ProtoName(field))
			fw.PLF("xs.%s = &v", field.CapitalName)
			fw.PLF("}")
		} else {
			fw.PLF("func (xs *%s) Set%s(v %s) {", msg.Name, field.CapitalName, ProtoName(field))
			fw.PLF("xs.%s = v", field.CapitalName)
			fw.PLF("}")
		}
		if field.IsList() {
			fw.PLF("func (xs *%s) Add%s(v %s) {", msg.Name, field.CapitalName, gen.FloatConvert(field.ListType))
			fw.PLF("xs.%s = append(xs.%s,v)", field.CapitalName, field.CapitalName)
			fw.PLF("}")
		}
		if field.IsMap() {
			fw.PLF("func (xs *%s) Add%s(v *%s) {", msg.Name, field.CapitalName, field.MsgOrEnumRef.Name)
			fw.PLF("xs.%s = append(xs.%s,v)", field.CapitalName, field.CapitalName)
			fw.PLF("}")
		}
	}
	// marshal unmarshal
	fw.PLF("func (xs *%s) Unmarshal(buf []byte) error {", msg.Name)
	fw.PLF("for len(buf) > 0 {")
	fw.PLF("number, _, n := protowire.ConsumeTag(buf)")
	fw.PLF("if n < 0 { return protowire.ParseError(n) }")
	fw.PLF("buf = buf[n:]")

	fw.PLF("switch number {")
	for _, field := range msg.MsgFields {
		fw.PLF("case %d:", field.Number)
		if (field.IsPrimary() || field.IsEnum()) && field.Kind != "string" && field.Kind != "float" && field.Kind != "double" {
			fw.PLF("v, n := protowire.ConsumeVarint(buf)")
			fw.PLF("if n < 0 { return protowire.ParseError(n) }")
			fw.PLF("buf = buf[n:]")
			if field.Kind == "bool" {
				fw.PLF("xs.Set%s(protowire.DecodeBool(v))", field.CapitalName)
			} else {
				fw.PLF("xs.Set%s(%s(v))", field.CapitalName, field.GoName())
			}
		}
		if field.Kind == "string" || field.Kind == "bytes" {
			fw.PLF("v,n := protowire.ConsumeBytes(buf)")
			fw.PLF("if n < 0 { return protowire.ParseError(n) }")
			fw.PLF("buf = buf[n:]")
			if field.Kind == "bytes" {
				fw.PLF("xs.Set%s(v)", field.CapitalName)
			} else {
				fw.PLF("xs.Set%s(syncdep.Bys2Str(v))", field.CapitalName)
			}
		}
		if field.Kind == "float" {
			fw.PLF("v,n := protowire.ConsumeFixed32(buf)")
			fw.PLF("if n < 0 { return protowire.ParseError(n) }")
			fw.PLF("buf = buf[n:]")
			fw.PLF("xs.Set%s(math.Float32frombits(v))", field.CapitalName)
		}
		if field.Kind == "double" {
			fw.PLF("v,n := protowire.ConsumeFixed64(buf)")
			fw.PLF("if n < 0 { return protowire.ParseError(n) }")
			fw.PLF("buf = buf[n:]")
			fw.PLF("xs.Set%s(math.Float64frombits(v))", field.CapitalName)
		}
		if field.IsList() && field.ListType != "string" {
			if field.ListType == "float" {
				fw.PLF("v,n := protowire.ConsumeBytes(buf)")
				fw.PLF("if n < 0 { return protowire.ParseError(n) }")
				fw.PLF("buf = buf[n:]")
				fw.PLF("for len(v) > 0 {")
				fw.PLF("vv,nn := protowire.ConsumeFixed32(v)")
				fw.PLF("if nn < 0 { return protowire.ParseError(n) }")
				fw.PLF("v = v[nn:]")
				fw.PLF("xs.Add%s(math.Float32frombits(vv))", field.CapitalName)
				fw.PLF("}")
			} else if field.ListType == "double" {
				fw.PLF("v,n := protowire.ConsumeBytes(buf)")
				fw.PLF("if n < 0 { return protowire.ParseError(n) }")
				fw.PLF("buf = buf[n:]")
				fw.PLF("for len(v) > 0 {")
				fw.PLF("vv,nn := protowire.ConsumeFixed64(v)")
				fw.PLF("if nn < 0 { return protowire.ParseError(n) }")
				fw.PLF("v = v[nn:]")
				fw.PLF("xs.Add%s(math.Float64frombits(vv))", field.CapitalName)
				fw.PLF("}")
			} else {
				fw.PLF("v,n := protowire.ConsumeBytes(buf)")
				fw.PLF("if n < 0 { return protowire.ParseError(n) }")
				fw.PLF("buf = buf[n:]")
				fw.PLF("for len(v) > 0 {")
				fw.PLF("vv,nn := protowire.ConsumeVarint(v)")
				fw.PLF("if nn < 0 { return protowire.ParseError(n) }")
				fw.PLF("v = v[nn:]")
				if field.ListType == "bool" {
					fw.PLF("xs.Add%s(protowire.DecodeBool(vv))", field.CapitalName)
				} else {
					fw.PLF("xs.Add%s(%s(vv))", field.CapitalName, gen.FloatConvert(field.ListType))
				}
				fw.PLF("}")
			}
		}
		if field.IsList() && field.ListType == "string" {
			fw.PLF("v,n := protowire.ConsumeBytes(buf)")
			fw.PLF("if n < 0 { return protowire.ParseError(n) }")
			fw.PLF("buf = buf[n:]")
			fw.PLF("xs.Add%s(syncdep.Bys2Str(v))", field.CapitalName)
		}
		if field.IsMap() || field.IsMsg() {
			fw.PLF("v,n := protowire.ConsumeBytes(buf)")
			fw.PLF("if n < 0 { return protowire.ParseError(n) }")
			fw.PLF("buf = buf[n:]")
			fw.PLF("tmp := &%s{}", field.MsgOrEnumRef.Name)
			fw.PLF("err := tmp.Unmarshal(v)")
			fw.PLF("if err != nil { return err }")
			if field.IsMsg() {
				fw.PLF("xs.Set%s(tmp)", field.CapitalName)
			} else {
				fw.PLF("xs.Add%s(tmp)", field.CapitalName)
			}
		}
		fw.PLF("break")
	}
	fw.PLF("}")

	fw.PLF("}")
	fw.PL("return nil")
	fw.PLF("}")

	fw.PLF("func (xs *%s) Marshal() []byte {", msg.Name)
	fw.PLF("var buf []byte")
	for _, field := range msg.MsgFields {
		fw.PLF("if xs.%s != nil {", field.CapitalName)
		if (field.IsPrimary() || field.IsEnum()) && field.Kind != "string" && field.Kind != "float" && field.Kind != "double" {
			fw.PLF("buf = protowire.AppendTag(buf,%d,protowire.VarintType)", field.Number)
			if field.Kind == "bool" {
				fw.PLF("buf = protowire.AppendVarint(buf,protowire.EncodeBool(*xs.%s))", field.CapitalName)
			} else {
				fw.PLF("buf = protowire.AppendVarint(buf,uint64(*xs.%s))", field.CapitalName)
			}
		}
		if field.Kind == "string" || field.Kind == "bytes" {
			fw.PLF("buf = protowire.AppendTag(buf,%d,protowire.BytesType)", field.Number)
			if field.Kind == "bytes" {
				fw.PLF("buf = protowire.AppendBytes(buf,xs.%s)", field.CapitalName)
			} else {
				fw.PLF("buf = protowire.AppendString(buf,*xs.%s)", field.CapitalName)
			}
		}
		if field.Kind == "float" {
			fw.PLF("buf = protowire.AppendTag(buf,%d,protowire.Fixed32Type)", field.Number)
			fw.PLF("buf = protowire.AppendFixed32(buf,math.Float32bits(*xs.%s))", field.CapitalName)
		}
		if field.Kind == "double" {
			fw.PLF("buf = protowire.AppendTag(buf,%d,protowire.Fixed64Type)", field.Number)
			fw.PLF("buf = protowire.AppendFixed64(buf,math.Float64bits(*xs.%s))", field.CapitalName)
		}
		if field.IsList() {
			if field.ListType == "string" {
				fw.PLF("for _,s := range xs.%s {", field.CapitalName)
				fw.PLF("buf = protowire.AppendTag(buf,%d,protowire.BytesType)", field.Number)
				fw.PLF("buf = protowire.AppendString(buf,s)")
				fw.PLF("}")
			} else if field.ListType == "float" {
				fw.PLF("for _,s := range xs.%s {", field.CapitalName)
				fw.PLF("buf = protowire.AppendTag(buf,%d,protowire.Fixed32Type)", field.Number)
				fw.PLF("buf = protowire.AppendFixed32(buf,math.Float32bits(s))")
				fw.PLF("}")
			} else if field.ListType == "double" {
				fw.PLF("for _,s := range xs.%s {", field.CapitalName)
				fw.PLF("buf = protowire.AppendTag(buf,%d,protowire.Fixed64Type)", field.Number)
				fw.PLF("buf = protowire.AppendFixed64(buf,math.Float64bits(s))")
				fw.PLF("}")
			} else {
				fw.PLF("for _,s := range xs.%s {", field.CapitalName)
				fw.PLF("buf = protowire.AppendTag(buf,%d,protowire.VarintType)", field.Number)
				if field.ListType == "bool" {
					fw.PLF("buf = protowire.AppendVarint(buf,protowire.EncodeBool(s))")
				} else {
					fw.PLF("buf = protowire.AppendVarint(buf,uint64(s))")
				}
				fw.PLF("}")
			}
		}
		if field.IsMap() {
			fw.PLF("for _,s := range xs.%s {", field.CapitalName)
			fw.PLF("bys := s.Marshal()")
			fw.PLF("buf = protowire.AppendTag(buf,%d,protowire.BytesType)", field.Number)
			fw.PLF("buf = protowire.AppendBytes(buf,bys)")
			fw.PLF("}")
		}
		if field.IsMsg() {
			fw.PLF("bys := (*xs.%s).Marshal()", field.CapitalName)
			fw.PLF("buf = protowire.AppendTag(buf,%d,protowire.BytesType)", field.Number)
			fw.PLF("buf = protowire.AppendBytes(buf,bys)")
		}
		fw.PLF("}")
	}
	fw.PLF("return buf")
	fw.PLF("}")
}

// ProtoName go type str
func ProtoName(sfd gen.SyncFieldDef) string {
	if gen.IsBuildInType(sfd.Kind) {
		return gen.FloatConvert(sfd.Kind)
	}

	if sfd.IsList() {
		return fmt.Sprintf("[]%s", gen.FloatConvert(sfd.ListType))
	}
	if sfd.IsMap() {
		return fmt.Sprintf("[]%s", "*"+sfd.MsgOrEnumRef.Name)
	}
	if sfd.IsEnum() {
		return sfd.MsgOrEnumRef.Name
	} else {
		return "*" + sfd.MsgOrEnumRef.Name
	}
}
