package gen

import (
	"fmt"
	"go/format"
	"os"
	"strings"
)

type FileWriter struct {
	sb  *strings.Builder
	idt int
}

func NewFileWriter() *FileWriter {
	return &FileWriter{
		sb:  &strings.Builder{},
		idt: 0,
	}
}

func (g *FileWriter) internalWriteOneLine(wrap bool, sep string, is ...interface{}) {
	ss := make([]string, 0)
	for _, s := range is {
		ss = append(ss, fmt.Sprintf("%v", s))
	}
	g.sb.WriteString(strings.Repeat("    ", g.idt))
	g.sb.WriteString(strings.Join(ss, sep))
	if wrap {
		g.sb.WriteString("\r\n")
	}

}
func (g *FileWriter) P(is ...interface{}) {
	g.internalWriteOneLine(false, "", is...)
}

func (g *FileWriter) PL(is ...interface{}) {
	g.internalWriteOneLine(true, "", is...)
}

func (g *FileWriter) PLF(f string, is ...interface{}) {
	g.internalWriteOneLine(true, "", fmt.Sprintf(f, is...))
}

func (g *FileWriter) Save(join string) {
	var s = []byte(g.sb.String())
	if strings.HasSuffix(join, ".go") {
		source, err := format.Source([]byte(g.sb.String()))
		if err == nil {
			s = source
		}
	}
	err := os.WriteFile(join, s, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func (g *FileWriter) String() string {
	return g.sb.String()
}
