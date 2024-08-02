package syncdep

import (
	"errors"
	"fmt"
	"google.golang.org/protobuf/encoding/protowire"
	"math"
)

var ErrParseRawFields = errors.New("parse raw fields error")

// RawField 内部的Field只会是bool或者repeated int32 int64 uint32 uint64 string bool
type RawField struct {
	Number int32
	Type   protowire.Type
	Bytes  []byte
}

type RawMessage struct {
	RawFields map[int32][]*RawField
}

func (rm *RawMessage) Marshal() []byte {
	var buf []byte
	for _, f := range rm.RawFields {
		for _, field := range f {
			buf = protowire.AppendTag(buf, protowire.Number(field.Number), field.Type)
			if field.Type == protowire.VarintType {
				buf = protowire.AppendVarint(buf, uint64(1))
			} else {
				buf = protowire.AppendBytes(buf, field.Bytes)
			}

		}
	}
	return buf
}

func (rm *RawMessage) AddString(num int32, val string) {
	bys := []byte(val)
	rm.RawFields[num] = append(rm.RawFields[num], &RawField{Number: num, Type: protowire.BytesType, Bytes: bys})
}

func (rm *RawMessage) GetStringList(num int32) []string {
	_, exist := rm.RawFields[num]
	if exist {
		ret := make([]string, len(rm.RawFields[num]))
		for i, f := range rm.RawFields[num] {
			ret[i] = string(f.Bytes)
		}
		return ret
	} else {
		return make([]string, 0)
	}
}
func (rm *RawMessage) GetInt32List(num int32) []int32 {
	f, exist := rm.RawFields[num]
	if exist {
		ret := make([]int32, 0)
		tmp := f[0].Bytes
		for len(tmp) > 0 {
			val, n := protowire.ConsumeVarint(tmp)
			if n < 0 {
				panic(n)
			}
			tmp = tmp[n:]
			ret = append(ret, int32(val))
		}
		return ret
	} else {
		return make([]int32, 0)
	}
}
func (rm *RawMessage) GetUint32List(num int32) []uint32 {
	f, exist := rm.RawFields[num]
	if exist {
		ret := make([]uint32, 0)
		tmp := f[0].Bytes
		for len(tmp) > 0 {
			val, n := protowire.ConsumeVarint(tmp)
			if n < 0 {
				panic(n)
			}
			tmp = tmp[n:]
			ret = append(ret, uint32(val))
		}
		return ret
	} else {
		return make([]uint32, 0)
	}
}
func (rm *RawMessage) GetInt64List(num int32) []int64 {
	f, exist := rm.RawFields[num]
	if exist {
		ret := make([]int64, 0)
		tmp := f[0].Bytes
		for len(tmp) > 0 {
			val, n := protowire.ConsumeVarint(tmp)
			if n < 0 {
				panic(n)
			}
			tmp = tmp[n:]
			ret = append(ret, int64(val))
		}
		return ret
	} else {
		return make([]int64, 0)
	}
}
func (rm *RawMessage) GetUint64List(num int32) []uint64 {
	f, exist := rm.RawFields[num]
	if exist {
		ret := make([]uint64, 0)
		tmp := f[0].Bytes
		for len(tmp) > 0 {
			val, n := protowire.ConsumeVarint(tmp)
			if n < 0 {
				panic(n)
			}
			tmp = tmp[n:]
			ret = append(ret, val)
		}
		return ret
	} else {
		return make([]uint64, 0)
	}
}
func (rm *RawMessage) GetBoolList(num int32) []bool {
	f, exist := rm.RawFields[num]
	if exist {
		ret := make([]bool, 0)
		tmp := f[0].Bytes
		for len(tmp) > 0 {
			val, n := protowire.ConsumeVarint(tmp)
			if n < 0 {
				panic(n)
			}
			tmp = tmp[n:]
			ret = append(ret, val > 0)
		}
		return ret
	} else {
		return make([]bool, 0)
	}
}
func (rm *RawMessage) SetBool(num int32) {
	_, exist := rm.RawFields[num]
	if !exist {
		rm.RawFields[num] = append(rm.RawFields[num], &RawField{Number: num, Type: protowire.VarintType})
	}
}
func (rm *RawMessage) GetBool(num int32) bool {
	_, exist := rm.RawFields[num]
	return exist
}

func (rm *RawMessage) ClearBool(num int32) {
	delete(rm.RawFields, num)
}

func (rm *RawMessage) AddVarint(num int32, val uint64) {
	rf, exist := rm.RawFields[num]
	if exist {
		rf[0].Bytes = protowire.AppendVarint(rf[0].Bytes, val)
	} else {
		bys := protowire.AppendVarint([]byte{}, val)
		rm.RawFields[num] = append(rm.RawFields[num], &RawField{Number: num, Type: protowire.BytesType, Bytes: bys})
	}
}

func ToRawMessage(buf []byte) RawMessage {
	rm := RawMessage{RawFields: make(map[int32][]*RawField)}
	for len(buf) > 0 {
		numb, typ, n := protowire.ConsumeTag(buf)
		if n < 0 {
			panic(fmt.Sprintf("parse error:consumeTag:%d", n))
		}
		num := int32(numb)
		buf = buf[n:]
		switch typ {
		case protowire.VarintType:
			_, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(fmt.Sprintf("parse error:ConsumeVarint:%d", n))
			}
			buf = buf[n:]
			rm.RawFields[num] = append(rm.RawFields[num], &RawField{Number: num, Type: protowire.VarintType})
		case protowire.BytesType:
			bytes, n := protowire.ConsumeBytes(buf)
			if n < 0 {
				panic(fmt.Sprintf("parse error:ConsumeBytes:%d", n))
			}
			rm.RawFields[num] = append(rm.RawFields[num], &RawField{Number: num, Type: protowire.BytesType, Bytes: bytes})
			buf = buf[n:]
		default:
			panic(fmt.Sprintf("unsupported type: %d", typ))
		}
	}
	return rm
}

func VarintRange(tmp []byte, f func(v uint64)) {
	for len(tmp) > 0 {
		val, n := protowire.ConsumeVarint(tmp)
		if n < 0 {
			panic(ErrParseRawFields)
		}
		tmp = tmp[n:]
		f(val)
	}
}

func Fixed32Range(tmp []byte, f func(v float32)) {
	for len(tmp) > 0 {
		val, n := protowire.ConsumeFixed32(tmp)
		if n < 0 {
			panic(ErrParseRawFields)
		}
		tmp = tmp[n:]
		f(math.Float32frombits(val))
	}
}

func Fixed64Range(tmp []byte, f func(v float64)) {
	for len(tmp) > 0 {
		val, n := protowire.ConsumeFixed64(tmp)
		if n < 0 {
			panic(ErrParseRawFields)
		}
		tmp = tmp[n:]
		f(math.Float64frombits(val))
	}
}

func ParseMap[K int32 | uint32 | int64 | uint64 | string | bool](tmp []byte) (K, []byte) {
	_, _, n := protowire.ConsumeTag(tmp)
	if n < 0 {
		panic(ErrParseRawFields)
	}
	tmp = tmp[n:]
	var k K
	if _, ok := any(k).(string); ok {
		bys, n := protowire.ConsumeBytes(tmp)
		if n < 0 {
			panic(ErrParseRawFields)
		}
		tmp = tmp[n:]
		k = any(string(bys)).(K)
	} else {
		v, n := protowire.ConsumeVarint(tmp)
		if n < 0 {
			panic(ErrParseRawFields)
		}
		tmp = tmp[n:]
		switch any(k).(type) {
		case bool:
			k = any(v > 0).(K)
		case int32:
			k = any(int32(v)).(K)
		case uint32:
			k = any(uint32(v)).(K)
		case int64:
			k = any(int64(v)).(K)
		case uint64:
			k = any(v).(K)
		}
	}
	_, _, n = protowire.ConsumeTag(tmp)
	if n < 0 {
		panic(ErrParseRawFields)
	}
	tmp = tmp[n:]
	bys, n := protowire.ConsumeBytes(tmp)
	if n < 0 {
		panic(ErrParseRawFields)
	}
	return k, bys
}
