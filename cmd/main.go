package main

import (
	"github.com/yaoguangduan/protosync/internalv2"
	"os"
)

func main() {
	var file = "define.sync"
	if len(os.Args) >= 2 {
		file = os.Args[1]
	}
	internalv2.Gen(file)
}
