package gen

import (
	"bytes"
	"fmt"
	"github.com/samber/lo"
	"github.com/yaoguangduan/datasync/internalv2/def"
	"github.com/yaoguangduan/datasync/syncproto"
	"go/format"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"log"
	"text/template"
)

func GenerateOneFile(prefix string, msgs map[string]*protogen.Message, templates *template.Template, plugin *protogen.Plugin, msgMap map[string]lo.Tuple2[*protogen.Message, string]) {
	filename := prefix + ".sync.go"
	f, _ := lo.Find(plugin.Files, func(item *protogen.File) bool {
		return item.GeneratedFilenamePrefix == prefix
	})
	log.Printf(filename)
	g := plugin.NewGeneratedFile(filename, f.GoImportPath)
	w := bytes.Buffer{}
	fileInfo := createFileInfo(msgs, f)
	err := templates.ExecuteTemplate(&w, "file.tmpl", &fileInfo)
	if err != nil {
		panic(err)
	}
	for _, msg := range msgs {
		structInfo := createMsgStructInfo(msg, msgMap)
		err = templates.ExecuteTemplate(&w, "struct.tmpl", &structInfo)
		if err != nil {
			panic(err)
		}
		err = templates.ExecuteTemplate(&w, "methods_interface.tmpl", &structInfo)
		if err != nil {
			panic(err)
		}
		err = templates.ExecuteTemplate(&w, "methods_op.tmpl", &structInfo)
		if err != nil {
			panic(err)
		}
		err = templates.ExecuteTemplate(&w, "methods_getset.tmpl", &structInfo)
		if err != nil {
			panic(err)
		}
		err = templates.ExecuteTemplate(&w, "methods_dirtyop.tmpl", &structInfo)
		if err != nil {
			panic(err)
		}
	}
	src, err := format.Source(w.Bytes())
	if err != nil {
		log.Println(err)
		g.P(w.String())
	} else {
		g.P(string(src))
	}
}

func createMsgStructInfo(msg *protogen.Message, msgMap map[string]lo.Tuple2[*protogen.Message, string]) def.GoStructInfo {
	structInfo := def.GoStructInfo{}
	structInfo.StructName = string(msg.Desc.Name() + "Sync")
	structInfo.MessageName = string(msg.Desc.Name())
	if proto.HasExtension(msg.Desc.Options(), syncproto.E_SyncKey) {
		structInfo.KeyFieldNum = int(proto.GetExtension(msg.Desc.Options(), syncproto.E_SyncKey).(int32))
	}
	var maxNumber = -1
	for i := 0; i < msg.Desc.Fields().Len(); i++ {
		fd := msg.Desc.Fields().Get(i)
		if fd.IsMap() || fd.IsList() {
			structInfo.HasMapOrListField = true
		}
		maxNumber = max(maxNumber, int(fd.Number()))
	}
	for _, field := range msg.Fields {
		fd := field.Desc
		fi := def.GoFieldInfo{}
		fi.FieldName = string(fd.Name())
		fi.OriFieldName = field.GoName
		fi.FieldType = fieldGoType(fd)
		fi.FieldNumber = int(fd.Number())
		fi.Kind = getKindFromFd(fd)
		if fi.FieldNumber == structInfo.KeyFieldNum {
			fi.IsKey = true
		}
		fi.DirtyIndex = fi.FieldNumber >> 3
		fi.DirtyOffset = fi.FieldNumber & 7
		fi.DefaultVal = fieldDefaultVal(fd)
		if fd.Kind() == protoreflect.MessageKind {
			if fd.IsMap() {
				fi.Struct = createMsgStructInfo(msgMap[string(fd.MapValue().Message().Name())].A, msgMap)
				fi.MapKeyType = fieldKindToGoType(fd.MapKey())
				fi.MapDelNumber = fi.FieldNumber + 1000
			} else {
				fi.Struct = createMsgStructInfo(msgMap[string(fd.Message().Name())].A, msgMap)
			}
		}
		if fd.IsList() {
			fi.ListType = fieldKindToGoType(fd)
			fi.ListClearNumber = fi.FieldNumber + 1000
		}
		fi.ProtoWireType = fieldTypeToWireType(fi.Kind)
		structInfo.Fields = append(structInfo.Fields, fi)
	}
	structInfo.DirtyArrSize = (maxNumber+1)/8 + 1
	return structInfo
}

func fieldTypeToWireType(kind string) protowire.Type {
	if kind == "enum" || kind == "bool" || kind == "int32" || kind == "uint32" || kind == "int64" || kind == "uint64" {
		return protowire.VarintType
	}
	if kind == "float64" {
		return protowire.Fixed64Type
	}
	if kind == "float32" {
		return protowire.Fixed32Type
	}
	return protowire.BytesType
}

func getKindFromFd(fd protoreflect.FieldDescriptor) string {
	var goType = fieldKindToGoType(fd)
	if fd.Kind() == protoreflect.MessageKind {
		goType = "message"
	}
	if fd.Kind() == protoreflect.EnumKind {
		goType = "enum"
	}
	if fd.IsMap() {
		goType = "map"
	}
	if fd.IsList() {
		goType = "list"
	}
	return goType
}

func fieldGoType(d protoreflect.FieldDescriptor) string {
	if d.IsMap() {
		return fmt.Sprintf("*syncdep.MapSync[%s,*%s]", fieldKindToGoType(d.MapKey()), fieldKindToGoType(d.MapValue()))
	} else if d.IsList() {
		return fmt.Sprintf("*syncdep.ArraySync[%s]", fieldKindToGoType(d))
	} else {
		if d.Kind() == protoreflect.MessageKind {
			return "*" + fieldKindToGoType(d)
		} else {
			return fieldKindToGoType(d)
		}
	}
}

func fieldDefaultVal(field protoreflect.FieldDescriptor) interface{} {
	kind := getKindFromFd(field)
	switch kind {
	case "int32", "int64", "uint32", "uint64", "float32", "float64":
		return 0
	case "string":
		return "\"\""
	case "bool":
		return false
	case "[]byte":
		return "make([]byte,0)"
	case "enum":
		return string(field.Enum().Name()) + "_" + string(field.Enum().Values().Get(0).FullName())
	}
	return nil
}

func fieldKindToGoType(d protoreflect.FieldDescriptor) string {
	switch d.Kind() {
	case protoreflect.StringKind:
		return "string"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "uint32"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return "int32"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "uint64"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return "int64"
	case protoreflect.BoolKind:
		return "bool"
	case protoreflect.DoubleKind:
		return "float64"
	case protoreflect.FloatKind:
		return "float32"
	case protoreflect.EnumKind:
		return string(d.Enum().Name())
	case protoreflect.BytesKind:
		return "[]byte"
	case protoreflect.MessageKind:
		return string(d.Message().Name() + "Sync")
	}
	panic(fmt.Sprintf("not supported field kind %s", d.Kind()))
}

func createFileInfo(msgs map[string]*protogen.Message, f *protogen.File) def.GoFileInfo {
	fi := def.GoFileInfo{PackageName: string(f.GoPackageName)}
	fi.Imports = append(fi.Imports, "github.com/yaoguangduan/datasync/syncdep")
	fi.Imports = append(fi.Imports, "google.golang.org/protobuf/encoding/protowire")
	fi.Imports = append(fi.Imports, "math")

	imports := map[string]struct{}{}
	for _, msg := range msgs {
		for _, field := range msg.Fields {
			if field.Desc.IsMap() {
				imports["slices"] = struct{}{}
				break
			}
		}
	}
	for s := range imports {
		fi.Imports = append(fi.Imports, s)
	}
	return fi
}
