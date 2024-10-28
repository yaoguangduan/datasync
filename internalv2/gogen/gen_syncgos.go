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
	}
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
