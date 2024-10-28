package def

import "google.golang.org/protobuf/encoding/protowire"

type GoFileInfo struct {
	PackageName string
	Imports     []string
}

type GoStructInfo struct {
	StructName        string
	Fields            []GoFieldInfo
	KeyFieldNum       int
	DirtyArrSize      int
	MessageName       string
	HasMapOrListField bool
}
type GoFieldInfo struct {
	FieldName       string
	FieldType       string
	FieldNumber     int
	OriFieldName    string // proto生成的go文件的field名称
	Kind            string
	DirtyIndex      int
	DirtyOffset     int
	IsKey           bool
	DefaultVal      interface{}
	Struct          GoStructInfo
	MapKeyType      string
	ListType        string
	MapDelNumber    int
	ListClearNumber int
	ProtoWireType   protowire.Type
}
