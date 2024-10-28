package internalv1

import (
	"embed"
	"fmt"
	"github.com/Masterminds/sprig/v3"
	"github.com/samber/lo"
	"github.com/yaoguangduan/protosync/internalv1/gen"
	"github.com/yaoguangduan/protosync/syncproto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"log"
	"text/template"
)

//go:embed asserts/*.tmpl
var templatesFS embed.FS

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

	tmpl := template.New("").Funcs(sprig.FuncMap())
	templates, err := tmpl.ParseFS(templatesFS, "asserts/*.tmpl")
	if err != nil {
		panic(err)
	}
	for _, t := range templates.Templates() {
		log.Printf("gen template :%+v", t)
	}
	for prefix, msgs := range fileMessagesMap {
		gen.GenerateOneFile(prefix, msgs, templates, plugin, msgMap)
	}
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
