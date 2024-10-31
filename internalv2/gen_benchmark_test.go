package internalv2

import (
	"github.com/stretchr/testify/assert"
	"github.com/yaoguangduan/protosync/pbgen"
	"google.golang.org/protobuf/proto"
	"testing"
)

/**
BenchmarkUnmarshalStandard
BenchmarkUnmarshalStandard-8   	  400316	      3193 ns/op
BenchmarkUnmarshalStandard-8   	  342280	      3007 ns/op
BenchmarkUnmarshalStandard-8   	  400479	      2927 ns/op

BenchmarkUnmarshalSync
BenchmarkUnmarshalSync-8      586020	      2087 ns/op
BenchmarkUnmarshalSync-8   	  540828	      1915 ns/op
BenchmarkUnmarshalSync-8   	  562594	      1911 ns/op

BenchmarkMarshalStandard
BenchmarkMarshalStandard-8   	 1000000	      1125 ns/op
BenchmarkMarshalStandard-8   	 1000000	      1014 ns/op
BenchmarkMarshalStandard-8   	 1000000	      1115 ns/op

BenchmarkMarshalSync
BenchmarkMarshalSync-8   	  991932	      1090 ns/op
BenchmarkMarshalSync-8   	 1000000	      1091 ns/op
BenchmarkMarshalSync-8   	 1124191	      1004 ns/op
*/

func BenchmarkUnmarshalStandard(b *testing.B) {
	test := fullTestData()
	testDto := &pbgen.Test{}
	test.CopyToPb(testDto)
	bys, err := proto.Marshal(testDto)
	assert.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tu := &pbgen.Test{}
		err = proto.Unmarshal(bys, tu)
	}
}

func BenchmarkUnmarshalSync(b *testing.B) {
	test := fullTestData()
	testDto := &pbgen.Test{}
	test.CopyToPb(testDto)
	bys, err := proto.Marshal(testDto)
	assert.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tu := &pbgen.Test{}
		err = tu.Unmarshal(bys)
	}
}

func BenchmarkMarshalStandard(b *testing.B) {
	test := fullTestData()
	testDto := &pbgen.Test{}
	test.CopyToPb(testDto)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(testDto)
	}
}

func BenchmarkMarshalSync(b *testing.B) {
	test := fullTestData()
	testDto := &pbgen.Test{}
	test.CopyToPb(testDto)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = testDto.Marshal()
	}
}
