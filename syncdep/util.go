package syncdep

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/yaoguangduan/protosync/internalv2/gen"
	"google.golang.org/protobuf/encoding/protowire"
	"math"
	"unsafe"
)

//go:inline
func Value[T any](v T) *T {
	vv := new(T)
	*vv = v
	return vv
}

func AppendFieldValue[V int32 | uint32 | int64 | uint64 | string | bool | []byte | float32 | float64](buf []byte, vv V) []byte {
	var value uint64
	switch v := any(vv).(type) {
	case int32:
		value = uint64(v)
		return protowire.AppendVarint(buf, value)
	case int64:
		value = uint64(v)
		return protowire.AppendVarint(buf, value)
	case uint32:
		value = uint64(v)
		return protowire.AppendVarint(buf, value)
	case uint64:
		return protowire.AppendVarint(buf, v)
	case bool:
		if v {
			value = 1
		} else {
			value = 0
		}
		return protowire.AppendVarint(buf, value)
	case float32:
		return protowire.AppendFixed32(buf, math.Float32bits(v))
	case float64:
		return protowire.AppendFixed64(buf, math.Float64bits(v))
	case string:
		return protowire.AppendString(buf, v)
	case []byte:
		return protowire.AppendBytes(buf, v)
	default:
		panic("unsupported type")
	}
}

func AppendListFieldValue[V int32 | uint32 | int64 | uint64 | string | bool | float32 | float64](buf []byte, a []V) []byte {
	// Create a buffer for the packed array
	var packedBuf []byte
	for _, v := range a {
		packedBuf = AppendFieldValue(packedBuf, v)
	}
	// Append the packed array itself
	buf = protowire.AppendBytes(buf, packedBuf)

	return buf
}

func AppendMapFieldKeyValue[K int32 | uint32 | int64 | uint64 | string | bool](buf []byte, fieldNumber int, k K, v []byte) []byte {
	var entryBuf []byte
	switch any(k).(type) {
	case string:
		entryBuf = protowire.AppendTag(entryBuf, 1, protowire.BytesType)
	default:
		entryBuf = protowire.AppendTag(entryBuf, 1, protowire.VarintType)
	}
	entryBuf = AppendFieldValue(entryBuf, k)
	entryBuf = protowire.AppendTag(entryBuf, 2, protowire.BytesType)
	entryBuf = protowire.AppendBytes(entryBuf, v)

	buf = protowire.AppendTag(buf, protowire.Number(fieldNumber), protowire.BytesType)
	buf = protowire.AppendBytes(buf, entryBuf)
	return buf
}

func FieldTypeToWireType(field gen.SyncFieldDef) protowire.Type {
	if field.Kind == "enum" || field.Kind == "bool" || field.Kind == "int32" || field.Kind == "uint32" || field.Kind == "int64" || field.Kind == "uint64" {
		return protowire.VarintType
	}
	if field.Kind == "double" {
		return protowire.Fixed64Type
	}
	if field.Kind == "float" {
		return protowire.Fixed32Type
	}
	return protowire.BytesType
}

type PreParsedProto struct {
	Values  []ProtoRawField
	Numbers []int
}

func (p PreParsedProto) NumberSlice() []int {
	if p.Numbers == nil {
		p.Numbers = lo.Map[ProtoRawField, int](p.Values, func(item ProtoRawField, index int) int {
			return item.Number
		})
	}
	return p.Numbers
}

func (p PreParsedProto) RawFieldByNumber(number int) *ProtoRawField {
	for _, f := range p.Values {
		if f.Number == number {
			return &f
		}
	}
	return nil
}

type ProtoRawField struct {
	Number int
	Type   protowire.Type
	Value  interface{}
}

func GetMapKey[T any](p *ProtoRawField) T {
	var t T
	switch any(t).(type) {
	case int32:
		return any(int32(p.Value.(uint64))).(T)
	case uint32:
		return any(uint32(p.Value.(uint64))).(T)
	case int64:
		return any(int64(p.Value.(uint64))).(T)
	case uint64:
		return any(p.Value.(uint64)).(T)
	case bool:
		return any(p.Value.(uint64) > 0).(T)
	case string:
		return any(string(p.Value.([]byte))).(T)
	default:
		panic(fmt.Sprintf("unkonwn raw type %v", p))
	}
}

// MustParseVarintArr int32 uint32 int64 uint64 bool enum
func MustParseVarintArr[T interface{}](p *ProtoRawField) []T {
	var t T
	buf := p.Value.([]byte)
	ret := make([]T, 0)
	for len(buf) > 0 {
		switch any(t).(type) {
		case int32:
			val, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(fmt.Sprintf("parse varint arr error:%d", n))
			}
			buf = buf[n:]
			ret = append(ret, any(int32(val)).(T))
		case uint32:
			val, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(fmt.Sprintf("parse varint arr error:%d", n))
			}
			buf = buf[n:]
			ret = append(ret, any(uint32(val)).(T))
		case int64:
			val, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(fmt.Sprintf("parse varint arr error:%d", n))
			}
			buf = buf[n:]
			ret = append(ret, any(int64(val)).(T))
		case uint64:
			val, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(fmt.Sprintf("parse varint arr error:%d", n))
			}
			buf = buf[n:]
			ret = append(ret, any(val).(T))
		case bool:
			val, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(fmt.Sprintf("parse varint arr error:%d", n))
			}
			buf = buf[n:]
			ret = append(ret, any(val > 0).(T))
		default:
			val, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(fmt.Sprintf("parse varint arr error:%d", n))
			}
			buf = buf[n:]
			ret = append(ret, any(int32(val)).(T))

		}
	}
	return ret
}
func Str2Bys(str string) []byte {
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

func Bys2Str(bys []byte) string {
	return unsafe.String(unsafe.SliceData(bys), len(bys))
}

/**
const (
	VarintType     Type = 0
	Fixed32Type    Type = 5
	Fixed64Type    Type = 1
	BytesType      Type = 2
	EndGroupType   Type = 4
)

*/

func PreParseProtoBytes(buf []byte) PreParsedProto {
	fields := make([]ProtoRawField, 0)
	for len(buf) > 0 {
		num, typ, n := protowire.ConsumeTag(buf)
		if n < 0 {
			panic(fmt.Sprintf("parse error:consumeTag:%d", n))
		}
		buf = buf[n:]
		var v interface{}
		switch typ {
		case protowire.VarintType:
			u64, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(fmt.Sprintf("parse error:ConsumeVarint:%d", n))
			}
			v = u64
			buf = buf[n:]
		case protowire.BytesType:
			bytes, n := protowire.ConsumeBytes(buf)
			if n < 0 {
				panic(fmt.Sprintf("parse error:ConsumeBytes:%d", n))
			}
			v = bytes
			buf = buf[n:]
		case protowire.Fixed32Type:
			u32, n := protowire.ConsumeFixed32(buf)
			if n < 0 {
				panic(fmt.Sprintf("parse error:ConsumeFixed32:%d", n))
			}
			v = math.Float32frombits(u32)
			buf = buf[n:]
		case protowire.Fixed64Type:
			u64, n := protowire.ConsumeFixed64(buf)
			if n < 0 {
				panic(fmt.Sprintf("parse error:ConsumeFixed64:%d", n))
			}
			v = math.Float64frombits(u64)
			buf = buf[n:]
		default:
			panic(fmt.Sprintf("unsupported type: %d", typ))
		}
		fields = append(fields, ProtoRawField{Number: int(num), Type: typ, Value: v})
	}
	return PreParsedProto{Values: fields}
}
