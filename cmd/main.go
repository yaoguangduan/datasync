package main

import (
	"gpsync/internal"
	"os"
)

func main() {
	var file = "define.sync"
	if len(os.Args) >= 2 {
		file = os.Args[1]
	}
	internal.Gen(file)
}
