package internal

import (
	"fmt"
	"github.com/samber/lo"
	"gpsync/internal/gen"
	"gpsync/internal/gogen"
	"gpsync/internal/protogen"
	"os"
	"path/filepath"
)

func Gen(input string) {

	sdf := gen.NewSyncDef()
	sdf.ParseSyncDefine(input)
	sdf.FormatAndWrite(input)

	clearDir(sdf.Defs["proto_out"])
	clearDir(sdf.Defs["go_out"])
	fmt.Println(sdf)
	protogen.GenerateProto(*sdf)
	gogen.GenerateGo(*sdf)
}

func clearDir(out string) {
	dir := lo.Must(os.ReadDir(out))
	for _, e := range dir {
		lo.Must0(os.RemoveAll(filepath.Join(out, e.Name())))
	}
}
