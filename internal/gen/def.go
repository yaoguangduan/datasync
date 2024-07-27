package gen

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

type MetaKVDef struct {
	Key  string
	Val  string
	Anno string
}

type SyncDef struct {
	Name     string
	Anno     string
	Defs     map[string]string
	Messages []SyncMsgOrEnumDef
}

func NewSyncDef() *SyncDef {
	return &SyncDef{
		Defs:     make(map[string]string),
		Messages: make([]SyncMsgOrEnumDef, 0),
	}
}

func (def *SyncDef) ParseSyncDefine(file string) {
	fr := lo.Must(os.Open(file))

	scan := bufio.NewScanner(fr)
	lines := make([]string, 0)
	for scan.Scan() {
		line := scan.Text()
		if line == "" || strings.HasPrefix(strings.TrimSpace(line), "//") {
			continue
		}
		if strings.TrimSpace(line) != "" {
			lines = append(lines, strings.TrimSpace(line))
		}
	}

	var syncMsg *SyncMsgOrEnumDef
	var enum = false
	var isMsgMap = make(map[string]bool)
	for _, line := range lines {
		if strings.HasPrefix(line, "enum ") || strings.HasPrefix(line, "message ") {
			info, _ := splitAnno(line)
			typeAndName := strings.Fields(strings.TrimSpace(info))
			if len(typeAndName) != 2 {
				panic(fmt.Sprintf("invalid define file:%s", line))
			}
			isMsgMap[strings.Split(typeAndName[1][:len(typeAndName[1])-1], "[")[0]] = typeAndName[0] == "message"
		}
	}
	for _, line := range lines {
		if strings.HasPrefix(line, "def ") {
			dkv := strings.Fields(line)
			def.Defs[dkv[1]] = strings.Join(dkv[2:], " ")
			continue
		}
		if strings.HasPrefix(line, "enum ") || strings.HasPrefix(line, "message ") {
			//new msg
			if syncMsg != nil {
				def.Messages = append(def.Messages, *syncMsg)
			}
			syncMsg = &SyncMsgOrEnumDef{EnumValues: []SyncEnumFieldDef{}}
			info, anno := splitAnno(line)
			syncMsg.Anno = anno
			typeAndName := strings.Fields(strings.TrimSpace(info))
			if len(typeAndName) != 2 {
				panic(fmt.Sprintf("invalid define file:%s", line))
			}
			syncMsg.Name = typeAndName[1]
			syncMsg.IsEnum = typeAndName[0] == "enum"
			enum = typeAndName[0] == "enum"
			if !enum {
				if strings.Contains(typeAndName[1], "[") && strings.HasSuffix(typeAndName[1], "]") {
					syncMsg.Name = strings.Split(typeAndName[1][:len(typeAndName[1])-1], "[")[0]
					syncMsg.MsgKey = &SyncFieldDef{Name: strings.Split(typeAndName[1][:len(typeAndName[1])-1], "[")[1]}
				} else {
					syncMsg.Name = typeAndName[1]
				}
			}
		} else {
			info, anno := splitAnno(line)
			if strings.TrimSpace(info) == "" {
				continue
			}
			field := SyncFieldDef{Anno: anno}
			fields := strings.Fields(strings.TrimSpace(info))
			if (enum && len(fields) != 2) || (!enum && len(fields) != 3) {
				panic(fmt.Sprintf("invalid define field:%s", line))
			}
			if enum {
				ef := SyncEnumFieldDef{Anno: anno, Value: fields[0], Number: str2int(fields[1])}
				syncMsg.EnumValues = append(syncMsg.EnumValues, ef)
			} else {
				field.Name = fields[1]
				field.Number = str2int(fields[2])
				field.Kind = fields[0]
				if strings.HasPrefix(field.Kind, "map[") && strings.HasSuffix(field.Kind, "]") {
					mapVal := field.Kind[4 : len(field.Kind)-1]
					if IsBuildInType(mapVal) {
						panic(fmt.Sprintf("invalid msg field define:(%s) -- map not support base value type!", mapVal))
					}
					field.MsgOrEnumRef = &SyncMsgOrEnumDef{Name: mapVal}
					field.Kind = "map"
				} else if strings.HasPrefix(field.Kind, "list[") && strings.HasSuffix(field.Kind, "]") {
					field.ListType = field.Kind[5 : len(field.Kind)-1]
					if !IsBuildInType(field.ListType) {
						field.MsgOrEnumRef = &SyncMsgOrEnumDef{Name: field.ListType}
					}
					field.Kind = "list"
				} else {
					if !IsBuildInType(field.Kind) {
						field.MsgOrEnumRef = &SyncMsgOrEnumDef{Name: field.Kind}
					}
				}
			}
			syncMsg.MsgFields = append(syncMsg.MsgFields, field)
		}
	}
	if syncMsg != nil {
		def.Messages = append(def.Messages, *syncMsg)
	}
	for i := range def.Messages {
		if def.Messages[i].IsEnum {
			continue
		}
		def.Messages[i].SyncName = def.Messages[i].Name + "Sync"
		var fieldKey *SyncFieldDef
		for j := range def.Messages[i].MsgFields {
			def.Messages[i].MsgFields[j].CapitalName = lo.Capitalize(def.Messages[i].MsgFields[j].Name[0:1]) + def.Messages[i].MsgFields[j].Name[1:]
			if def.Messages[i].MsgKey != nil && def.Messages[i].MsgFields[j].Name == def.Messages[i].MsgKey.Name {
				fieldKey = &def.Messages[i].MsgFields[j]
			}
			dep := def.Messages[i].MsgFields[j].DependencyName()
			if dep != nil {
				msgOrEnum := def.GetMsgOrEnumByName(*dep)
				def.Messages[i].MsgFields[j].MsgOrEnumRef = msgOrEnum
				if !def.Messages[i].MsgFields[j].IsMap() && !def.Messages[i].MsgFields[j].IsList() {
					def.Messages[i].MsgFields[j].Kind = lo.If(msgOrEnum.IsEnum, "enum").Else("message")
				}
			}
		}
		if fieldKey != nil {
			def.Messages[i].MsgKey = fieldKey
		}
	}
	for i := range def.Messages {
		if def.Messages[i].IsEnum {
			continue
		}
		for j := range def.Messages[i].MsgFields {
			if def.Messages[i].MsgFields[j].IsMap() {
				if def.Messages[i].MsgFields[j].MsgOrEnumRef.MsgKey == nil {
					panic(fmt.Sprintf("invalid msg field define:(%s[%s]) -- key missing!", def.Messages[i].Name, def.Messages[i].MsgFields[j].Name))
				}
				def.Messages[i].MsgFields[j].MapKeyKind = def.Messages[i].MsgFields[j].MsgOrEnumRef.MsgKey.Kind
			}
			if def.Messages[i].MsgFields[j].IsList() && def.Messages[i].MsgFields[j].MsgOrEnumRef != nil && !def.Messages[i].MsgFields[j].MsgOrEnumRef.IsEnum {
				panic(fmt.Sprintf("invalid msg field define:(%s[%s]) -- array not support message!", def.Messages[i].Name, def.Messages[i].MsgFields[j].Name))
			}
			if def.Messages[i].MsgFields[j].IsMap() && def.Messages[i].MsgFields[j].MsgOrEnumRef.IsEnum {
				panic(fmt.Sprintf("invalid msg field define:(%s[%s]) -- map not support enum!", def.Messages[i].Name, def.Messages[i].MsgFields[j].Name))
			}
		}
	}
	for k := range def.Defs {
		if def.Defs[k] == "go_out" || def.Defs[k] == "proto_out" {
			def.Defs[k] = lo.Must(filepath.Rel(lo.Must(os.Getwd()), def.Defs[k]))
		}

	}
}

func str2int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func splitAnno(line string) (string, string) {
	idx := strings.Index(line, "//")
	if idx < 0 {
		return line, ""
	}
	return strings.TrimSpace(line[0:idx]), strings.TrimSpace(line[idx:])
}

func (def *SyncDef) GetMsgOrEnumByName(s string) *SyncMsgOrEnumDef {
	msgOrEnum, find := lo.Find(def.Messages, func(item SyncMsgOrEnumDef) bool {
		return item.Name == s
	})
	if !find {
		panic(fmt.Sprintf("missing message or enum define:%s", s))
	}
	return &msgOrEnum
}

func (def *SyncDef) FormatAndWrite(file string) {
	fw := NewFileWriter()
	for k, v := range def.Defs {
		fw.PL("def ", k, " ", v)
	}
	fw.PL()
	for _, msg := range def.Messages {
		lines := make([][]string, 0)
		if msg.IsEnum {
			fw.P("enum ")
			fw.PL(msg.Name, " ", msg.Anno)
			for _, e := range msg.EnumValues {
				lines = append(lines, []string{e.Value, strconv.Itoa(e.Number), e.Anno})
			}
		} else {
			fw.P("message ")
			if msg.MsgKey != nil {
				fw.PL(msg.Name, fmt.Sprintf("[%s]", msg.MsgKey.Name), " ", msg.Anno)
			} else {
				fw.PL(msg.Name)
			}
			for _, m := range msg.MsgFields {
				if m.IsPrimary() || m.IsBytes() {
					lines = append(lines, []string{m.Kind, m.Name, strconv.Itoa(m.Number) + m.Anno})
				} else if m.IsMsg() || m.IsEnum() {
					lines = append(lines, []string{m.MsgOrEnumRef.Name, m.Name, strconv.Itoa(m.Number) + m.Anno})
				} else if m.IsList() {
					lines = append(lines, []string{fmt.Sprintf("list[%v]", m.ListType), m.Name, strconv.Itoa(m.Number) + m.Anno})
				} else {
					lines = append(lines, []string{fmt.Sprintf("map[%s]", m.MsgOrEnumRef.Name), m.Name, strconv.Itoa(m.Number) + m.Anno})
				}
			}
		}

		colMaxMap := make(map[int]int)
		for _, line := range lines {
			for idx := range line {
				colMaxMap[idx] = max(colMaxMap[idx], len(line[idx]))
			}
		}
		fmt.Println(lines)
		for _, line := range lines {
			newLine := make([]interface{}, 0)
			for idx, word := range line {
				newLine = append(newLine, word+strings.Repeat(" ", colMaxMap[idx]-len(word)+10))
			}
			ls := fmt.Sprintf("    "+strings.Repeat("%s", len(newLine)), newLine...)
			fw.PLF(strings.TrimSuffix(ls, " "))
		}
		fw.PL()
	}
	fw.Save(file)
}

type SyncEnumFieldDef struct {
	Value  string
	Number int
	Anno   string
}

type SyncMsgOrEnumDef struct {
	Pkg        string
	Name       string
	SyncName   string
	IsEnum     bool
	EnumValues []SyncEnumFieldDef
	MsgKey     *SyncFieldDef
	MsgFields  []SyncFieldDef
	Anno       string
}

func (d SyncMsgOrEnumDef) MaxFieldNumber() int {
	var i = -1
	for _, f := range d.MsgFields {
		i = max(f.Number, i)
	}
	return i
}

func (d SyncMsgOrEnumDef) BitfieldNumber() int {
	var numMax = -1
	for _, field := range d.MsgFields {
		numMax = max(numMax, field.Number)
	}
	for _, field := range d.MsgFields {
		numMax = max(numMax, field.Number)
		if field.IsMap() {
			numMax = max(numMax, field.Number+1000)
		}
	}
	return numMax + 1
}

type SyncFieldDef struct {
	Name         string
	CapitalName  string
	Number       int
	Anno         string
	Kind         string //"int32", "uint32", "int64", "uint64", "float32", "float64", "string","message","enum","bool","list"
	ListType     string // Kind == "list"
	MapKeyKind   string //"int32", "uint32", "int64", "uint64","string","bool"
	MsgOrEnumRef *SyncMsgOrEnumDef
}

func (sfd SyncFieldDef) GoName() string {
	if sfd.IsMap() {
		return fmt.Sprintf("*syncdep.MapSync[%s,*%s]", sfd.MapKeyKind, sfd.MsgOrEnumRef.SyncName)
	} else if sfd.IsList() {
		return fmt.Sprintf("*syncdep.ArraySync[%s]", FloatConvert(sfd.ListType))
	} else if IsBuildInType(sfd.Kind) {
		return FloatConvert(sfd.Kind)
	} else if sfd.IsEnum() {
		return sfd.MsgOrEnumRef.Name
	} else {
		return "*" + sfd.MsgOrEnumRef.SyncName
	}
}

func (sfd SyncFieldDef) IsMap() bool {
	return sfd.Kind == "map"
}

func (sfd SyncFieldDef) IsList() bool {
	return sfd.Kind == "list"
}

var BuildInTypes = []string{"int32", "uint32", "int64", "uint64", "float", "double", "string", "bool", "bytes"}

func (sfd SyncFieldDef) IsMsg() bool {
	return sfd.Kind == "message"
}

func (sfd SyncFieldDef) IsEnum() bool {
	return sfd.Kind == "enum"
}

func (sfd SyncFieldDef) IsBytes() bool {
	return sfd.Kind == "bytes"
}

func (sfd SyncFieldDef) IsPrimary() bool {
	return IsBuildInType(sfd.Kind) && !sfd.IsEnum() && !sfd.IsBytes()
}

// DependencyName nil if no dep
func (sfd SyncFieldDef) DependencyName() *string {
	if sfd.MsgOrEnumRef != nil {
		return &sfd.MsgOrEnumRef.Name
	}
	return nil
}

func (sfd SyncFieldDef) ProtoDelNumber() int {
	return sfd.Number + 1000
}
func (sfd SyncFieldDef) ProtoClearNumber() int {
	return sfd.Number + 2000
}

func FloatConvert(s string) string {
	if s == "float" {
		return "float32"
	}
	if s == "double" {
		return "float64"
	}
	if s == "bytes" {
		return "[]byte"
	}
	return s
}

func IsBuildInType(s string) bool {
	return slices.Contains(BuildInTypes, s)
}
