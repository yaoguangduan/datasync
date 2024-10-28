package main

import (
	"github.com/yaoguangduan/protosync/internalv1"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		internalv1.GenFromPlugin(plugin)
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return nil
	})
}
