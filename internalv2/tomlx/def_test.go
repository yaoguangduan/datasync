package tomlx

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	bytes, err := os.ReadFile("../../define.toml")
	if err != nil {
		panic(err)
	}
	pd := &Config{}
	toml.Decode([]byte(string(bytes)), pd)
	fmt.Printf("Config: %+v\n", pd)

}
