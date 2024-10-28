package internalv2

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/yaoguangduan/protosync/internalv2/gen"
	"github.com/yaoguangduan/protosync/internalv2/gogen"
	"github.com/yaoguangduan/protosync/internalv2/proto_file_gen"
	"github.com/yaoguangduan/protosync/syncproto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"log"
	"os"
	"path/filepath"
)

func Gen(input string) {

	sdf := gen.NewSyncDef()
	sdf.ParseSyncDefine(input)
	sdf.FormatAndWrite(input)

	fmt.Println(sdf)
	proto_file_gen.GenerateProto(*sdf)
	gogen.GenerateGo(*sdf)
}

func clearDir(out string) {
	dir := lo.Must(os.ReadDir(out))
	for _, e := range dir {
		lo.Must0(os.RemoveAll(filepath.Join(out, e.Name())))
	}
}

func GenFromPlugin(plugin *protogen.Plugin) {
	msgMap := make(map[string]lo.Tuple2[*protogen.Message, string])
	enumMap := make(map[string]lo.Tuple2[*protogen.Enum, string])
	for _, file := range plugin.Files {
		for _, msg := range file.Messages {
			msgMap[string(msg.Desc.Name())] = lo.Tuple2[*protogen.Message, string]{A: msg, B: file.GeneratedFilenamePrefix}
		}
		for _, enum := range file.Enums {
			enumMap[string(enum.Desc.Name())] = lo.Tuple2[*protogen.Enum, string]{A: enum, B: file.GeneratedFilenamePrefix}
		}
	}

	fileMessagesMap := make(map[string]map[string]*protogen.Message)
	for _, tuple := range msgMap {
		msg := tuple.A
		if proto.HasExtension(msg.Desc.Options(), syncproto.E_SyncGen) {
			genIt := proto.GetExtension(msg.Desc.Options(), syncproto.E_SyncGen).(bool)
			if genIt {
				findMessageDep(fileMessagesMap, tuple, msgMap)
			}
		}
	}
	log.Printf("%+v\n", enumMap)
	nameFileDefMap := make(map[string]*gen.SyncDef)
	for prefix, msgs := range fileMessagesMap {
		sfd := gen.NewSyncDef()
		f, _ := lo.Find(plugin.Files, func(item *protogen.File) bool {
			return item.GeneratedFilenamePrefix == prefix
		})
		sfd.PluginFile = f
		sfd.Defs["go_out"] = prefix
		for name, msg := range msgs {
			msgDef := gen.SyncMsgOrEnumDef{}
			msgDef.IsEnum = false
			msgDef.Name = name
			msgDef.SyncName = name + "Sync"
			for _, field := range msg.Fields {
				fieldDef := gen.SyncFieldDef{}
				fieldDef.Number = int(field.Desc.Number())
				fieldDef.Kind = toDefKind(field.Desc)
				if field.Desc.IsList() {
					fieldDef.Kind = "list"
				}
				if field.Desc.IsMap() {
					fieldDef.Kind = "map"
				}
				fieldDef.Name = string(field.Desc.Name())
				fieldDef.CapitalName = lo.Capitalize(fieldDef.Name[0:1]) + fieldDef.Name[1:]
				if fieldDef.IsList() {
					if field.Desc.Kind() == protoreflect.EnumKind {
						fieldDef.ListType = string(field.Enum.Desc.Name())
						fieldDef.MsgOrEnumRef = &gen.SyncMsgOrEnumDef{Name: fieldDef.ListType}
					} else {
						fieldDef.ListType = toDefKind(field.Desc)
					}
				}
				if field.Desc.IsMap() {
					fieldDef.MapKeyKind = toDefKind(field.Message.Fields[0].Desc)
					fieldDef.MsgOrEnumRef = &gen.SyncMsgOrEnumDef{Name: string(field.Message.Fields[1].Desc.Message().Name())}
				}
				if fieldDef.Kind == "message" {
					fieldDef.MsgOrEnumRef = &gen.SyncMsgOrEnumDef{Name: string(field.Message.Desc.Name())}
				}
				if fieldDef.Kind == "enum" {
					fieldDef.MsgOrEnumRef = &gen.SyncMsgOrEnumDef{Name: string(field.Enum.Desc.Name())}
				}
				msgDef.MsgFields = append(msgDef.MsgFields, fieldDef)
				if proto.HasExtension(msg.Desc.Options(), syncproto.E_SyncKey) {
					if int32(field.Desc.Number()) == proto.GetExtension(msg.Desc.Options(), syncproto.E_SyncKey).(int32) {
						msgDef.MsgKey = &fieldDef
					}
				}
			}
			sfd.Messages = append(sfd.Messages, msgDef)
		}
		nameFileDefMap[prefix] = sfd
	}
	for _, sfd := range nameFileDefMap {
		for i := range sfd.Messages {
			for j := range sfd.Messages[i].MsgFields {
				if sfd.Messages[i].MsgFields[j].MsgOrEnumRef != nil {
					relRef := findRefByName(sfd.Messages[i].MsgFields[j].MsgOrEnumRef.Name, nameFileDefMap, enumMap)
					sfd.Messages[i].MsgFields[j].MsgOrEnumRef = relRef
				}
			}
		}
	}
	for _, sfd := range nameFileDefMap {
		gogen.GenerateGoFromPlugin(plugin, *sfd)
	}
}

func findRefByName(name string, defMap map[string]*gen.SyncDef, enumMap map[string]lo.Tuple2[*protogen.Enum, string]) *gen.SyncMsgOrEnumDef {
	for _, def := range defMap {
		msg := def.GetMsgOrEnumByName(name)
		if msg != nil {
			return msg
		}
	}
	t, exist := enumMap[name]
	if exist {
		e := t.A
		ed := gen.SyncMsgOrEnumDef{IsEnum: true}
		ed.Name = string(e.Desc.Name())
		for _, v := range e.Values {
			ef := gen.SyncEnumFieldDef{}
			ef.Value = string(v.Desc.Name())
			ef.Number = int(v.Desc.Number())
			ed.EnumValues = append(ed.EnumValues, ef)
		}
		return &ed
	}
	panic(fmt.Sprintf("msg %s cannot find", name))
}

func toDefKind(desc protoreflect.FieldDescriptor) string {
	switch desc.Kind() {
	case protoreflect.BoolKind:
		return "bool"
	case protoreflect.EnumKind:
		return "enum"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "uint64"
	case protoreflect.FloatKind:
		return "float"
	case protoreflect.DoubleKind:
		return "double"
	case protoreflect.StringKind:
		return "string"
	case protoreflect.BytesKind:
		return "bytes"
	case protoreflect.MessageKind:
		return "message"
	}
	panic(fmt.Sprintf("unknown enum kind %v", desc))
}

func findMessageDep(fileMessagesMap map[string]map[string]*protogen.Message, tuple lo.Tuple2[*protogen.Message, string], msgMap map[string]lo.Tuple2[*protogen.Message, string]) {
	msg := tuple.A
	msgName := string(msg.Desc.Name())
	if _, ok := fileMessagesMap[tuple.B]; !ok {
		fileMessagesMap[tuple.B] = make(map[string]*protogen.Message)
	}
	fileMessagesMap[tuple.B][msgName] = msg
	for _, field := range msg.Fields {
		if field.Desc.IsList() && field.Message != nil && field.Desc.Kind() != protoreflect.EnumKind {
			panic(fmt.Sprintf("msg %s can not contain message(%s) list", msgName, field.Desc.Name()))
		}
		if field.Desc.IsMap() && field.Message.Fields[1].Message == nil {
			panic(fmt.Sprintf("msg %s can not contain base(%s) map val", msgName, field.Message.Fields[1].Desc.Name()))
		}
		if field.Desc.IsMap() {
			mapVal := field.Message.Fields[1]
			if proto.HasExtension(mapVal.Message.Desc.Options(), syncproto.E_SyncKey) {
				genKey := proto.GetExtension(mapVal.Message.Desc.Options(), syncproto.E_SyncKey).(int32)
				fk := mapVal.Message.Desc.Fields().ByNumber(protoreflect.FieldNumber(genKey))
				if fk == nil {
					panic(fmt.Sprintf("map val %s must specify key field", mapVal.Message.Desc.Name()))
				}
				tp, exist := msgMap[string(mapVal.Message.Desc.Name())]
				if !exist {
					panic(fmt.Sprintf("map val type %s not exist", mapVal.Message.Desc.Name()))
				}
				findMessageDep(fileMessagesMap, tp, msgMap)
			} else {
				panic(fmt.Sprintf("map val %s must specify key field", mapVal.Message.Desc.Name()))
			}
		} else {
			if field.Message != nil {
				tp, exist := msgMap[string(field.Message.Desc.Name())]
				if !exist {
					panic(fmt.Sprintf("map val type %s not exist", field.Message.Desc.Name()))
				}
				findMessageDep(fileMessagesMap, tp, msgMap)
			}
		}

	}
}
