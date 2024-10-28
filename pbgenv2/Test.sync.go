package pbgenv2

import "github.com/yaoguangduan/protosync/syncdep"

import "google.golang.org/protobuf/encoding/protowire"

import "math"

import "slices"

// struct  TestSync start

type TestSync struct {
	id int32

	u32 uint32

	i64 int64

	u64 uint64

	b bool

	e ColorType

	str string

	obj *PersonSync

	i32Arr *syncdep.ArraySync[int32]

	u32Arr *syncdep.ArraySync[uint32]

	i64Arr *syncdep.ArraySync[int64]

	u64Arr *syncdep.ArraySync[uint64]

	boolArr *syncdep.ArraySync[bool]

	enumArr *syncdep.ArraySync[ColorType]

	strArr *syncdep.ArraySync[string]

	i32Map *syncdep.MapSync[int32, *TestI32MapSync]

	u32Map *syncdep.MapSync[uint32, *TestU32MapSync]

	i64Map *syncdep.MapSync[int64, *TestI64MapSync]

	u64Map *syncdep.MapSync[uint64, *TestU64MapSync]

	boolMap *syncdep.MapSync[bool, *TestBoolMapSync]

	strMap *syncdep.MapSync[string, *TestStringMapSync]

	f32 float32

	f64 float64

	f32Arr *syncdep.ArraySync[float32]

	f64Arr *syncdep.ArraySync[float64]

	dfm []uint8
	p   syncdep.Sync
	i   int
}

func NewTestSync() *TestSync {
	return &TestSync{
		dfm: make([]uint8, 4),
	}
}

// struct TestSync Sync interface methods start
func (x *TestSync) SetDirty(i int, dirty bool, sync syncdep.Sync) {
	idx := i >> 3
	off := i & 7
	if dirty {
		x.dfm[idx] = x.dfm[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dfm[idx] = x.dfm[idx] & ^(1 << off)
	}
}

func (x *TestSync) SetParentDirty() {
	if x.p != nil {
		x.p.SetDirty(x.i, true, x)
	}
}

func (x *TestSync) SetParent(sync syncdep.Sync, i int) {
	x.p = sync
	x.i = i
}
func (x *TestSync) FlushDirty(dirty bool) {

	if dirty || x.isIdDirty() {
		x.setIdDirty(dirty, true)
	}

	if dirty || x.isU32Dirty() {
		x.setU32Dirty(dirty, true)
	}

	if dirty || x.isI64Dirty() {
		x.setI64Dirty(dirty, true)
	}

	if dirty || x.isU64Dirty() {
		x.setU64Dirty(dirty, true)
	}

	if dirty || x.isBDirty() {
		x.setBDirty(dirty, true)
	}

	if dirty || x.isEDirty() {
		x.setEDirty(dirty, true)
	}

	if dirty || x.isStrDirty() {
		x.setStrDirty(dirty, true)
	}

	if dirty || x.isObjDirty() {
		x.setObjDirty(dirty, true)
	}

	if dirty || x.isI32ArrDirty() {
		x.setI32ArrDirty(dirty, true)
	}

	if dirty || x.isU32ArrDirty() {
		x.setU32ArrDirty(dirty, true)
	}

	if dirty || x.isI64ArrDirty() {
		x.setI64ArrDirty(dirty, true)
	}

	if dirty || x.isU64ArrDirty() {
		x.setU64ArrDirty(dirty, true)
	}

	if dirty || x.isBoolArrDirty() {
		x.setBoolArrDirty(dirty, true)
	}

	if dirty || x.isEnumArrDirty() {
		x.setEnumArrDirty(dirty, true)
	}

	if dirty || x.isStrArrDirty() {
		x.setStrArrDirty(dirty, true)
	}

	if dirty || x.isI32MapDirty() {
		x.setI32MapDirty(dirty, true)
	}

	if dirty || x.isU32MapDirty() {
		x.setU32MapDirty(dirty, true)
	}

	if dirty || x.isI64MapDirty() {
		x.setI64MapDirty(dirty, true)
	}

	if dirty || x.isU64MapDirty() {
		x.setU64MapDirty(dirty, true)
	}

	if dirty || x.isBoolMapDirty() {
		x.setBoolMapDirty(dirty, true)
	}

	if dirty || x.isStrMapDirty() {
		x.setStrMapDirty(dirty, true)
	}

	if dirty || x.isF32Dirty() {
		x.setF32Dirty(dirty, true)
	}

	if dirty || x.isF64Dirty() {
		x.setF64Dirty(dirty, true)
	}

	if dirty || x.isF32ArrDirty() {
		x.setF32ArrDirty(dirty, true)
	}

	if dirty || x.isF64ArrDirty() {
		x.setF64ArrDirty(dirty, true)
	}

}

func (x *TestSync) setIdDirty(dirty bool, recur bool) {
	x.SetDirty(1, dirty, x)

}
func (x *TestSync) isIdDirty() bool {
	return (x.dfm[0] & (1 << 1)) != 0
}

func (x *TestSync) setU32Dirty(dirty bool, recur bool) {
	x.SetDirty(2, dirty, x)

}
func (x *TestSync) isU32Dirty() bool {
	return (x.dfm[0] & (1 << 2)) != 0
}

func (x *TestSync) setI64Dirty(dirty bool, recur bool) {
	x.SetDirty(3, dirty, x)

}
func (x *TestSync) isI64Dirty() bool {
	return (x.dfm[0] & (1 << 3)) != 0
}

func (x *TestSync) setU64Dirty(dirty bool, recur bool) {
	x.SetDirty(4, dirty, x)

}
func (x *TestSync) isU64Dirty() bool {
	return (x.dfm[0] & (1 << 4)) != 0
}

func (x *TestSync) setBDirty(dirty bool, recur bool) {
	x.SetDirty(5, dirty, x)

}
func (x *TestSync) isBDirty() bool {
	return (x.dfm[0] & (1 << 5)) != 0
}

func (x *TestSync) setEDirty(dirty bool, recur bool) {
	x.SetDirty(6, dirty, x)

}
func (x *TestSync) isEDirty() bool {
	return (x.dfm[0] & (1 << 6)) != 0
}

func (x *TestSync) setStrDirty(dirty bool, recur bool) {
	x.SetDirty(7, dirty, x)

}
func (x *TestSync) isStrDirty() bool {
	return (x.dfm[0] & (1 << 7)) != 0
}

func (x *TestSync) setObjDirty(dirty bool, recur bool) {
	x.SetDirty(22, dirty, x)

	if recur && x.obj != nil {
		x.obj.FlushDirty(dirty)
	}

}
func (x *TestSync) isObjDirty() bool {
	return (x.dfm[2] & (1 << 6)) != 0
}

func (x *TestSync) setI32ArrDirty(dirty bool, recur bool) {
	x.SetDirty(8, dirty, x)

}
func (x *TestSync) isI32ArrDirty() bool {
	return (x.dfm[1] & (1 << 0)) != 0
}

func (x *TestSync) setU32ArrDirty(dirty bool, recur bool) {
	x.SetDirty(9, dirty, x)

}
func (x *TestSync) isU32ArrDirty() bool {
	return (x.dfm[1] & (1 << 1)) != 0
}

func (x *TestSync) setI64ArrDirty(dirty bool, recur bool) {
	x.SetDirty(10, dirty, x)

}
func (x *TestSync) isI64ArrDirty() bool {
	return (x.dfm[1] & (1 << 2)) != 0
}

func (x *TestSync) setU64ArrDirty(dirty bool, recur bool) {
	x.SetDirty(11, dirty, x)

}
func (x *TestSync) isU64ArrDirty() bool {
	return (x.dfm[1] & (1 << 3)) != 0
}

func (x *TestSync) setBoolArrDirty(dirty bool, recur bool) {
	x.SetDirty(12, dirty, x)

}
func (x *TestSync) isBoolArrDirty() bool {
	return (x.dfm[1] & (1 << 4)) != 0
}

func (x *TestSync) setEnumArrDirty(dirty bool, recur bool) {
	x.SetDirty(13, dirty, x)

}
func (x *TestSync) isEnumArrDirty() bool {
	return (x.dfm[1] & (1 << 5)) != 0
}

func (x *TestSync) setStrArrDirty(dirty bool, recur bool) {
	x.SetDirty(14, dirty, x)

}
func (x *TestSync) isStrArrDirty() bool {
	return (x.dfm[1] & (1 << 6)) != 0
}

func (x *TestSync) setI32MapDirty(dirty bool, recur bool) {
	x.SetDirty(15, dirty, x)

	if recur && x.i32Map != nil {
		x.i32Map.FlushDirty(dirty)
	}

}
func (x *TestSync) isI32MapDirty() bool {
	return (x.dfm[1] & (1 << 7)) != 0
}

func (x *TestSync) setU32MapDirty(dirty bool, recur bool) {
	x.SetDirty(16, dirty, x)

	if recur && x.u32Map != nil {
		x.u32Map.FlushDirty(dirty)
	}

}
func (x *TestSync) isU32MapDirty() bool {
	return (x.dfm[2] & (1 << 0)) != 0
}

func (x *TestSync) setI64MapDirty(dirty bool, recur bool) {
	x.SetDirty(17, dirty, x)

	if recur && x.i64Map != nil {
		x.i64Map.FlushDirty(dirty)
	}

}
func (x *TestSync) isI64MapDirty() bool {
	return (x.dfm[2] & (1 << 1)) != 0
}

func (x *TestSync) setU64MapDirty(dirty bool, recur bool) {
	x.SetDirty(18, dirty, x)

	if recur && x.u64Map != nil {
		x.u64Map.FlushDirty(dirty)
	}

}
func (x *TestSync) isU64MapDirty() bool {
	return (x.dfm[2] & (1 << 2)) != 0
}

func (x *TestSync) setBoolMapDirty(dirty bool, recur bool) {
	x.SetDirty(19, dirty, x)

	if recur && x.boolMap != nil {
		x.boolMap.FlushDirty(dirty)
	}

}
func (x *TestSync) isBoolMapDirty() bool {
	return (x.dfm[2] & (1 << 3)) != 0
}

func (x *TestSync) setStrMapDirty(dirty bool, recur bool) {
	x.SetDirty(21, dirty, x)

	if recur && x.strMap != nil {
		x.strMap.FlushDirty(dirty)
	}

}
func (x *TestSync) isStrMapDirty() bool {
	return (x.dfm[2] & (1 << 5)) != 0
}

func (x *TestSync) setF32Dirty(dirty bool, recur bool) {
	x.SetDirty(23, dirty, x)

}
func (x *TestSync) isF32Dirty() bool {
	return (x.dfm[2] & (1 << 7)) != 0
}

func (x *TestSync) setF64Dirty(dirty bool, recur bool) {
	x.SetDirty(24, dirty, x)

}
func (x *TestSync) isF64Dirty() bool {
	return (x.dfm[3] & (1 << 0)) != 0
}

func (x *TestSync) setF32ArrDirty(dirty bool, recur bool) {
	x.SetDirty(25, dirty, x)

}
func (x *TestSync) isF32ArrDirty() bool {
	return (x.dfm[3] & (1 << 1)) != 0
}

func (x *TestSync) setF64ArrDirty(dirty bool, recur bool) {
	x.SetDirty(26, dirty, x)

}
func (x *TestSync) isF64ArrDirty() bool {
	return (x.dfm[3] & (1 << 2)) != 0
}

func (x *TestSync) Key() interface{} {
	return nil
}
func (x *TestSync) SetKey(v interface{}) {
}

// struct TestSync Sync interface methods end

// struct  TestSync method clear copy methods start

func (x *TestSync) Clear() *TestSync {

	x.SetId(0)

	x.SetU32(0)

	x.SetI64(0)

	x.SetU64(0)

	x.SetB(false)

	x.SetE(ColorType_Red)

	x.SetStr("")

	if x.obj != nil {
		x.obj.Clear()
	}

	if x.i32Arr != nil {
		x.i32Arr.Clear()
	}

	if x.u32Arr != nil {
		x.u32Arr.Clear()
	}

	if x.i64Arr != nil {
		x.i64Arr.Clear()
	}

	if x.u64Arr != nil {
		x.u64Arr.Clear()
	}

	if x.boolArr != nil {
		x.boolArr.Clear()
	}

	if x.enumArr != nil {
		x.enumArr.Clear()
	}

	if x.strArr != nil {
		x.strArr.Clear()
	}

	if x.i32Map != nil {
		x.i32Map.Clear()
	}

	if x.u32Map != nil {
		x.u32Map.Clear()
	}

	if x.i64Map != nil {
		x.i64Map.Clear()
	}

	if x.u64Map != nil {
		x.u64Map.Clear()
	}

	if x.boolMap != nil {
		x.boolMap.Clear()
	}

	if x.strMap != nil {
		x.strMap.Clear()
	}

	x.SetF32(0)

	x.SetF64(0)

	if x.f32Arr != nil {
		x.f32Arr.Clear()
	}

	if x.f64Arr != nil {
		x.f64Arr.Clear()
	}

	return x
}

func (x *TestSync) CopyFromPb(r *Test) *TestSync {

	if r.Id != nil {
		x.SetId(*r.Id)
	}

	if r.U32 != nil {
		x.SetU32(*r.U32)
	}

	if r.I64 != nil {
		x.SetI64(*r.I64)
	}

	if r.U64 != nil {
		x.SetU64(*r.U64)
	}

	if r.B != nil {
		x.SetB(*r.B)
	}

	if r.E != nil {
		x.SetE(*r.E)
	}

	if r.Str != nil {
		x.SetStr(*r.Str)
	}

	if r.Obj != nil {
		x.GetObj().CopyFromPb(r.Obj)
	}

	if len(r.I32Arr) > 0 {
		x.GetI32Arr().AddAll(r.I32Arr)
	}

	if len(r.U32Arr) > 0 {
		x.GetU32Arr().AddAll(r.U32Arr)
	}

	if len(r.I64Arr) > 0 {
		x.GetI64Arr().AddAll(r.I64Arr)
	}

	if len(r.U64Arr) > 0 {
		x.GetU64Arr().AddAll(r.U64Arr)
	}

	if len(r.BoolArr) > 0 {
		x.GetBoolArr().AddAll(r.BoolArr)
	}

	if len(r.EnumArr) > 0 {
		x.GetEnumArr().AddAll(r.EnumArr)
	}

	if len(r.StrArr) > 0 {
		x.GetStrArr().AddAll(r.StrArr)
	}

	for k, v := range r.I32Map {
		if v != nil {
			vv := NewTestI32MapSync()
			vv.CopyFromPb(v)
			x.GetI32Map().Put(k, vv)
		}
	}

	for k, v := range r.U32Map {
		if v != nil {
			vv := NewTestU32MapSync()
			vv.CopyFromPb(v)
			x.GetU32Map().Put(k, vv)
		}
	}

	for k, v := range r.I64Map {
		if v != nil {
			vv := NewTestI64MapSync()
			vv.CopyFromPb(v)
			x.GetI64Map().Put(k, vv)
		}
	}

	for k, v := range r.U64Map {
		if v != nil {
			vv := NewTestU64MapSync()
			vv.CopyFromPb(v)
			x.GetU64Map().Put(k, vv)
		}
	}

	for k, v := range r.BoolMap {
		if v != nil {
			vv := NewTestBoolMapSync()
			vv.CopyFromPb(v)
			x.GetBoolMap().Put(k, vv)
		}
	}

	for k, v := range r.StrMap {
		if v != nil {
			vv := NewTestStringMapSync()
			vv.CopyFromPb(v)
			x.GetStrMap().Put(k, vv)
		}
	}

	if r.F32 != nil {
		x.SetF32(*r.F32)
	}

	if r.F64 != nil {
		x.SetF64(*r.F64)
	}

	if len(r.F32Arr) > 0 {
		x.GetF32Arr().AddAll(r.F32Arr)
	}

	if len(r.F64Arr) > 0 {
		x.GetF64Arr().AddAll(r.F64Arr)
	}

	return x
}

func (x *TestSync) CopyToPb(r *Test) *TestSync {

	r.Id = &x.id

	r.U32 = &x.u32

	r.I64 = &x.i64

	r.U64 = &x.u64

	r.B = &x.b

	r.E = &x.e

	r.Str = &x.str

	if x.obj != nil {
		tmpV := &Person{}
		x.obj.CopyToPb(tmpV)
		r.Obj = tmpV
	}

	if x.i32Arr != nil && x.i32Arr.Len() > 0 {
		r.I32Arr = x.i32Arr.ValueView()
	}

	if x.u32Arr != nil && x.u32Arr.Len() > 0 {
		r.U32Arr = x.u32Arr.ValueView()
	}

	if x.i64Arr != nil && x.i64Arr.Len() > 0 {
		r.I64Arr = x.i64Arr.ValueView()
	}

	if x.u64Arr != nil && x.u64Arr.Len() > 0 {
		r.U64Arr = x.u64Arr.ValueView()
	}

	if x.boolArr != nil && x.boolArr.Len() > 0 {
		r.BoolArr = x.boolArr.ValueView()
	}

	if x.enumArr != nil && x.enumArr.Len() > 0 {
		r.EnumArr = x.enumArr.ValueView()
	}

	if x.strArr != nil && x.strArr.Len() > 0 {
		r.StrArr = x.strArr.ValueView()
	}

	if x.i32Map != nil && x.i32Map.Len() > 0 {
		tmp := make(map[int32]*TestI32Map)
		x.i32Map.Each(func(k int32, v *TestI32MapSync) bool {
			tmpV := &TestI32Map{}
			v.CopyToPb(tmpV)
			tmp[k] = tmpV
			return true
		})
		r.I32Map = tmp
	}

	if x.u32Map != nil && x.u32Map.Len() > 0 {
		tmp := make(map[uint32]*TestU32Map)
		x.u32Map.Each(func(k uint32, v *TestU32MapSync) bool {
			tmpV := &TestU32Map{}
			v.CopyToPb(tmpV)
			tmp[k] = tmpV
			return true
		})
		r.U32Map = tmp
	}

	if x.i64Map != nil && x.i64Map.Len() > 0 {
		tmp := make(map[int64]*TestI64Map)
		x.i64Map.Each(func(k int64, v *TestI64MapSync) bool {
			tmpV := &TestI64Map{}
			v.CopyToPb(tmpV)
			tmp[k] = tmpV
			return true
		})
		r.I64Map = tmp
	}

	if x.u64Map != nil && x.u64Map.Len() > 0 {
		tmp := make(map[uint64]*TestU64Map)
		x.u64Map.Each(func(k uint64, v *TestU64MapSync) bool {
			tmpV := &TestU64Map{}
			v.CopyToPb(tmpV)
			tmp[k] = tmpV
			return true
		})
		r.U64Map = tmp
	}

	if x.boolMap != nil && x.boolMap.Len() > 0 {
		tmp := make(map[bool]*TestBoolMap)
		x.boolMap.Each(func(k bool, v *TestBoolMapSync) bool {
			tmpV := &TestBoolMap{}
			v.CopyToPb(tmpV)
			tmp[k] = tmpV
			return true
		})
		r.BoolMap = tmp
	}

	if x.strMap != nil && x.strMap.Len() > 0 {
		tmp := make(map[string]*TestStringMap)
		x.strMap.Each(func(k string, v *TestStringMapSync) bool {
			tmpV := &TestStringMap{}
			v.CopyToPb(tmpV)
			tmp[k] = tmpV
			return true
		})
		r.StrMap = tmp
	}

	r.F32 = &x.f32

	r.F64 = &x.f64

	if x.f32Arr != nil && x.f32Arr.Len() > 0 {
		r.F32Arr = x.f32Arr.ValueView()
	}

	if x.f64Arr != nil && x.f64Arr.Len() > 0 {
		r.F64Arr = x.f64Arr.ValueView()
	}

	return x

}

// struct  TestSync get set methods start

func (x *TestSync) GetId() int32 {

	return x.id
}

func (x *TestSync) SetId(v int32) *TestSync {

	if x.id == v {
		return x
	}

	x.id = v
	x.setIdDirty(true, false)
	return x
}

func (x *TestSync) GetU32() uint32 {

	return x.u32
}

func (x *TestSync) SetU32(v uint32) *TestSync {

	if x.u32 == v {
		return x
	}

	x.u32 = v
	x.setU32Dirty(true, false)
	return x
}

func (x *TestSync) GetI64() int64 {

	return x.i64
}

func (x *TestSync) SetI64(v int64) *TestSync {

	if x.i64 == v {
		return x
	}

	x.i64 = v
	x.setI64Dirty(true, false)
	return x
}

func (x *TestSync) GetU64() uint64 {

	return x.u64
}

func (x *TestSync) SetU64(v uint64) *TestSync {

	if x.u64 == v {
		return x
	}

	x.u64 = v
	x.setU64Dirty(true, false)
	return x
}

func (x *TestSync) GetB() bool {

	return x.b
}

func (x *TestSync) SetB(v bool) *TestSync {

	if x.b == v {
		return x
	}

	x.b = v
	x.setBDirty(true, false)
	return x
}

func (x *TestSync) GetE() ColorType {

	return x.e
}

func (x *TestSync) SetE(v ColorType) *TestSync {

	if x.e == v {
		return x
	}

	x.e = v
	x.setEDirty(true, false)
	return x
}

func (x *TestSync) GetStr() string {

	return x.str
}

func (x *TestSync) SetStr(v string) *TestSync {

	if x.str == v {
		return x
	}

	x.str = v
	x.setStrDirty(true, false)
	return x
}

func (x *TestSync) GetObj() *PersonSync {

	if x.obj == nil {
		x.obj = NewPersonSync()
		x.obj.SetParent(x, 22)
	}

	return x.obj
}

func (x *TestSync) SetObj(v *PersonSync) *TestSync {

	v.SetParent(x, 22)
	if x.obj != nil {
		x.obj.SetParent(nil, -1)
	}

	x.obj = v
	x.setObjDirty(true, false)
	return x
}

func (x *TestSync) GetI32Arr() *syncdep.ArraySync[int32] {

	if x.i32Arr == nil {
		x.i32Arr = syncdep.NewArraySync[int32]()
		x.i32Arr.SetParent(x, 8)
	}

	return x.i32Arr
}

func (x *TestSync) GetU32Arr() *syncdep.ArraySync[uint32] {

	if x.u32Arr == nil {
		x.u32Arr = syncdep.NewArraySync[uint32]()
		x.u32Arr.SetParent(x, 9)
	}

	return x.u32Arr
}

func (x *TestSync) GetI64Arr() *syncdep.ArraySync[int64] {

	if x.i64Arr == nil {
		x.i64Arr = syncdep.NewArraySync[int64]()
		x.i64Arr.SetParent(x, 10)
	}

	return x.i64Arr
}

func (x *TestSync) GetU64Arr() *syncdep.ArraySync[uint64] {

	if x.u64Arr == nil {
		x.u64Arr = syncdep.NewArraySync[uint64]()
		x.u64Arr.SetParent(x, 11)
	}

	return x.u64Arr
}

func (x *TestSync) GetBoolArr() *syncdep.ArraySync[bool] {

	if x.boolArr == nil {
		x.boolArr = syncdep.NewArraySync[bool]()
		x.boolArr.SetParent(x, 12)
	}

	return x.boolArr
}

func (x *TestSync) GetEnumArr() *syncdep.ArraySync[ColorType] {

	if x.enumArr == nil {
		x.enumArr = syncdep.NewArraySync[ColorType]()
		x.enumArr.SetParent(x, 13)
	}

	return x.enumArr
}

func (x *TestSync) GetStrArr() *syncdep.ArraySync[string] {

	if x.strArr == nil {
		x.strArr = syncdep.NewArraySync[string]()
		x.strArr.SetParent(x, 14)
	}

	return x.strArr
}

func (x *TestSync) GetI32Map() *syncdep.MapSync[int32, *TestI32MapSync] {

	if x.i32Map == nil {
		x.i32Map = syncdep.NewMapSync[int32, *TestI32MapSync]()
		x.i32Map.SetParent(x, 15)
	}

	return x.i32Map
}

func (x *TestSync) GetU32Map() *syncdep.MapSync[uint32, *TestU32MapSync] {

	if x.u32Map == nil {
		x.u32Map = syncdep.NewMapSync[uint32, *TestU32MapSync]()
		x.u32Map.SetParent(x, 16)
	}

	return x.u32Map
}

func (x *TestSync) GetI64Map() *syncdep.MapSync[int64, *TestI64MapSync] {

	if x.i64Map == nil {
		x.i64Map = syncdep.NewMapSync[int64, *TestI64MapSync]()
		x.i64Map.SetParent(x, 17)
	}

	return x.i64Map
}

func (x *TestSync) GetU64Map() *syncdep.MapSync[uint64, *TestU64MapSync] {

	if x.u64Map == nil {
		x.u64Map = syncdep.NewMapSync[uint64, *TestU64MapSync]()
		x.u64Map.SetParent(x, 18)
	}

	return x.u64Map
}

func (x *TestSync) GetBoolMap() *syncdep.MapSync[bool, *TestBoolMapSync] {

	if x.boolMap == nil {
		x.boolMap = syncdep.NewMapSync[bool, *TestBoolMapSync]()
		x.boolMap.SetParent(x, 19)
	}

	return x.boolMap
}

func (x *TestSync) GetStrMap() *syncdep.MapSync[string, *TestStringMapSync] {

	if x.strMap == nil {
		x.strMap = syncdep.NewMapSync[string, *TestStringMapSync]()
		x.strMap.SetParent(x, 21)
	}

	return x.strMap
}

func (x *TestSync) GetF32() float32 {

	return x.f32
}

func (x *TestSync) SetF32(v float32) *TestSync {

	if x.f32 == v {
		return x
	}

	x.f32 = v
	x.setF32Dirty(true, false)
	return x
}

func (x *TestSync) GetF64() float64 {

	return x.f64
}

func (x *TestSync) SetF64(v float64) *TestSync {

	if x.f64 == v {
		return x
	}

	x.f64 = v
	x.setF64Dirty(true, false)
	return x
}

func (x *TestSync) GetF32Arr() *syncdep.ArraySync[float32] {

	if x.f32Arr == nil {
		x.f32Arr = syncdep.NewArraySync[float32]()
		x.f32Arr.SetParent(x, 25)
	}

	return x.f32Arr
}

func (x *TestSync) GetF64Arr() *syncdep.ArraySync[float64] {

	if x.f64Arr == nil {
		x.f64Arr = syncdep.NewArraySync[float64]()
		x.f64Arr.SetParent(x, 26)
	}

	return x.f64Arr
}

// struct  TestSync dirty operate methods start

func (x *TestSync) MergeDirtyToPb(r *Test) *TestSync {

	var raw = syncdep.ToRawMessage(r.ProtoReflect().GetUnknown())

	if x.isIdDirty() {
		tmp := x.id
		r.Id = &tmp
	}

	if x.isU32Dirty() {
		tmp := x.u32
		r.U32 = &tmp
	}

	if x.isI64Dirty() {
		tmp := x.i64
		r.I64 = &tmp
	}

	if x.isU64Dirty() {
		tmp := x.u64
		r.U64 = &tmp
	}

	if x.isBDirty() {
		tmp := x.b
		r.B = &tmp
	}

	if x.isEDirty() {
		tmp := x.e
		r.E = &tmp
	}

	if x.isStrDirty() {
		tmp := x.str
		r.Str = &tmp
	}

	if x.isObjDirty() {
		if r.Obj == nil {
			r.Obj = &Person{}
		}
		x.obj.MergeDirtyToPb(r.Obj)
	}

	if x.isI32ArrDirty() {
		count := x.i32Arr.Len()
		r.I32Arr = make([]int32, 0)
		if count > 0 {
			raw.ClearBool(1008)
			r.I32Arr = append(r.I32Arr, x.i32Arr.ValueView()...)
		} else {
			raw.SetBool(1008)
		}
	}

	if x.isU32ArrDirty() {
		count := x.u32Arr.Len()
		r.U32Arr = make([]uint32, 0)
		if count > 0 {
			raw.ClearBool(1009)
			r.U32Arr = append(r.U32Arr, x.u32Arr.ValueView()...)
		} else {
			raw.SetBool(1009)
		}
	}

	if x.isI64ArrDirty() {
		count := x.i64Arr.Len()
		r.I64Arr = make([]int64, 0)
		if count > 0 {
			raw.ClearBool(1010)
			r.I64Arr = append(r.I64Arr, x.i64Arr.ValueView()...)
		} else {
			raw.SetBool(1010)
		}
	}

	if x.isU64ArrDirty() {
		count := x.u64Arr.Len()
		r.U64Arr = make([]uint64, 0)
		if count > 0 {
			raw.ClearBool(1011)
			r.U64Arr = append(r.U64Arr, x.u64Arr.ValueView()...)
		} else {
			raw.SetBool(1011)
		}
	}

	if x.isBoolArrDirty() {
		count := x.boolArr.Len()
		r.BoolArr = make([]bool, 0)
		if count > 0 {
			raw.ClearBool(1012)
			r.BoolArr = append(r.BoolArr, x.boolArr.ValueView()...)
		} else {
			raw.SetBool(1012)
		}
	}

	if x.isEnumArrDirty() {
		count := x.enumArr.Len()
		r.EnumArr = make([]ColorType, 0)
		if count > 0 {
			raw.ClearBool(1013)
			r.EnumArr = append(r.EnumArr, x.enumArr.ValueView()...)
		} else {
			raw.SetBool(1013)
		}
	}

	if x.isStrArrDirty() {
		count := x.strArr.Len()
		r.StrArr = make([]string, 0)
		if count > 0 {
			raw.ClearBool(1014)
			r.StrArr = append(r.StrArr, x.strArr.ValueView()...)
		} else {
			raw.SetBool(1014)
		}
	}

	if x.isI32MapDirty() {
		updated := make([]int32, 0)
		if r.I32Map != nil {
			for k := range r.I32Map {
				if x.i32Map.ContainDeleted(k) {
					delete(r.I32Map, k)
				}
				if x.i32Map.ContainDirtied(k) {
					updated = append(updated, k)
					tmp := x.i32Map.Get(k)
					if r.I32Map[k] == nil {
						r.I32Map[k] = &TestI32Map{}
					}
					tmp.MergeDirtyToPb(r.I32Map[k])
				}
			}
		} else {
			r.I32Map = make(map[int32]*TestI32Map)
		}
		for k := range x.i32Map.Dirtied() {
			if !slices.Contains(updated, k) {
				tmp := x.i32Map.Get(k)
				if r.I32Map[k] == nil {
					r.I32Map[k] = &TestI32Map{}
				}
				tmp.MergeDirtyToPb(r.I32Map[k])
			}
		}

		for k := range x.i32Map.Deleted() {
			raw.AddVarint(1015, uint64(k))
		}

	}

	if x.isU32MapDirty() {
		updated := make([]uint32, 0)
		if r.U32Map != nil {
			for k := range r.U32Map {
				if x.u32Map.ContainDeleted(k) {
					delete(r.U32Map, k)
				}
				if x.u32Map.ContainDirtied(k) {
					updated = append(updated, k)
					tmp := x.u32Map.Get(k)
					if r.U32Map[k] == nil {
						r.U32Map[k] = &TestU32Map{}
					}
					tmp.MergeDirtyToPb(r.U32Map[k])
				}
			}
		} else {
			r.U32Map = make(map[uint32]*TestU32Map)
		}
		for k := range x.u32Map.Dirtied() {
			if !slices.Contains(updated, k) {
				tmp := x.u32Map.Get(k)
				if r.U32Map[k] == nil {
					r.U32Map[k] = &TestU32Map{}
				}
				tmp.MergeDirtyToPb(r.U32Map[k])
			}
		}

		for k := range x.u32Map.Deleted() {
			raw.AddVarint(1016, uint64(k))
		}

	}

	if x.isI64MapDirty() {
		updated := make([]int64, 0)
		if r.I64Map != nil {
			for k := range r.I64Map {
				if x.i64Map.ContainDeleted(k) {
					delete(r.I64Map, k)
				}
				if x.i64Map.ContainDirtied(k) {
					updated = append(updated, k)
					tmp := x.i64Map.Get(k)
					if r.I64Map[k] == nil {
						r.I64Map[k] = &TestI64Map{}
					}
					tmp.MergeDirtyToPb(r.I64Map[k])
				}
			}
		} else {
			r.I64Map = make(map[int64]*TestI64Map)
		}
		for k := range x.i64Map.Dirtied() {
			if !slices.Contains(updated, k) {
				tmp := x.i64Map.Get(k)
				if r.I64Map[k] == nil {
					r.I64Map[k] = &TestI64Map{}
				}
				tmp.MergeDirtyToPb(r.I64Map[k])
			}
		}

		for k := range x.i64Map.Deleted() {
			raw.AddVarint(1017, uint64(k))
		}

	}

	if x.isU64MapDirty() {
		updated := make([]uint64, 0)
		if r.U64Map != nil {
			for k := range r.U64Map {
				if x.u64Map.ContainDeleted(k) {
					delete(r.U64Map, k)
				}
				if x.u64Map.ContainDirtied(k) {
					updated = append(updated, k)
					tmp := x.u64Map.Get(k)
					if r.U64Map[k] == nil {
						r.U64Map[k] = &TestU64Map{}
					}
					tmp.MergeDirtyToPb(r.U64Map[k])
				}
			}
		} else {
			r.U64Map = make(map[uint64]*TestU64Map)
		}
		for k := range x.u64Map.Dirtied() {
			if !slices.Contains(updated, k) {
				tmp := x.u64Map.Get(k)
				if r.U64Map[k] == nil {
					r.U64Map[k] = &TestU64Map{}
				}
				tmp.MergeDirtyToPb(r.U64Map[k])
			}
		}

		for k := range x.u64Map.Deleted() {
			raw.AddVarint(1018, uint64(k))
		}

	}

	if x.isBoolMapDirty() {
		updated := make([]bool, 0)
		if r.BoolMap != nil {
			for k := range r.BoolMap {
				if x.boolMap.ContainDeleted(k) {
					delete(r.BoolMap, k)
				}
				if x.boolMap.ContainDirtied(k) {
					updated = append(updated, k)
					tmp := x.boolMap.Get(k)
					if r.BoolMap[k] == nil {
						r.BoolMap[k] = &TestBoolMap{}
					}
					tmp.MergeDirtyToPb(r.BoolMap[k])
				}
			}
		} else {
			r.BoolMap = make(map[bool]*TestBoolMap)
		}
		for k := range x.boolMap.Dirtied() {
			if !slices.Contains(updated, k) {
				tmp := x.boolMap.Get(k)
				if r.BoolMap[k] == nil {
					r.BoolMap[k] = &TestBoolMap{}
				}
				tmp.MergeDirtyToPb(r.BoolMap[k])
			}
		}

		for k := range x.boolMap.Deleted() {
			var tmp = 0
			if k {
				tmp = 1
			}
			raw.AddVarint(1019, uint64(tmp))
		}

	}

	if x.isStrMapDirty() {
		updated := make([]string, 0)
		if r.StrMap != nil {
			for k := range r.StrMap {
				if x.strMap.ContainDeleted(k) {
					delete(r.StrMap, k)
				}
				if x.strMap.ContainDirtied(k) {
					updated = append(updated, k)
					tmp := x.strMap.Get(k)
					if r.StrMap[k] == nil {
						r.StrMap[k] = &TestStringMap{}
					}
					tmp.MergeDirtyToPb(r.StrMap[k])
				}
			}
		} else {
			r.StrMap = make(map[string]*TestStringMap)
		}
		for k := range x.strMap.Dirtied() {
			if !slices.Contains(updated, k) {
				tmp := x.strMap.Get(k)
				if r.StrMap[k] == nil {
					r.StrMap[k] = &TestStringMap{}
				}
				tmp.MergeDirtyToPb(r.StrMap[k])
			}
		}

		for k := range x.strMap.Deleted() {
			raw.AddString(1021, k)
		}

	}

	if x.isF32Dirty() {
		tmp := x.f32
		r.F32 = &tmp
	}

	if x.isF64Dirty() {
		tmp := x.f64
		r.F64 = &tmp
	}

	if x.isF32ArrDirty() {
		count := x.f32Arr.Len()
		r.F32Arr = make([]float32, 0)
		if count > 0 {
			raw.ClearBool(1025)
			r.F32Arr = append(r.F32Arr, x.f32Arr.ValueView()...)
		} else {
			raw.SetBool(1025)
		}
	}

	if x.isF64ArrDirty() {
		count := x.f64Arr.Len()
		r.F64Arr = make([]float64, 0)
		if count > 0 {
			raw.ClearBool(1026)
			r.F64Arr = append(r.F64Arr, x.f64Arr.ValueView()...)
		} else {
			raw.SetBool(1026)
		}
	}

	r.ProtoReflect().SetUnknown(raw.Marshal())

	return x
}

func (x *TestSync) MergeDirtyFromPb(r *Test) *TestSync {

	var raw = syncdep.ToRawMessage(r.ProtoReflect().GetUnknown())

	if r.Id != nil {
		x.SetId(*r.Id)
	}

	if r.U32 != nil {
		x.SetU32(*r.U32)
	}

	if r.I64 != nil {
		x.SetI64(*r.I64)
	}

	if r.U64 != nil {
		x.SetU64(*r.U64)
	}

	if r.B != nil {
		x.SetB(*r.B)
	}

	if r.E != nil {
		x.SetE(*r.E)
	}

	if r.Str != nil {
		x.SetStr(*r.Str)
	}

	if r.Obj != nil {
		x.GetObj().MergeDirtyFromPb(r.Obj)
	}

	if len(r.I32Arr) > 0 || raw.GetBool(1008) {
		x.GetI32Arr().Clear()
		x.i32Arr.AddAll(r.I32Arr)
	}

	if len(r.U32Arr) > 0 || raw.GetBool(1009) {
		x.GetU32Arr().Clear()
		x.u32Arr.AddAll(r.U32Arr)
	}

	if len(r.I64Arr) > 0 || raw.GetBool(1010) {
		x.GetI64Arr().Clear()
		x.i64Arr.AddAll(r.I64Arr)
	}

	if len(r.U64Arr) > 0 || raw.GetBool(1011) {
		x.GetU64Arr().Clear()
		x.u64Arr.AddAll(r.U64Arr)
	}

	if len(r.BoolArr) > 0 || raw.GetBool(1012) {
		x.GetBoolArr().Clear()
		x.boolArr.AddAll(r.BoolArr)
	}

	if len(r.EnumArr) > 0 || raw.GetBool(1013) {
		x.GetEnumArr().Clear()
		x.enumArr.AddAll(r.EnumArr)
	}

	if len(r.StrArr) > 0 || raw.GetBool(1014) {
		x.GetStrArr().Clear()
		x.strArr.AddAll(r.StrArr)
	}

	if x.i32Map != nil {
		x.i32Map.RemoveAll(raw.GetInt32List(1015))
	}
	for k, v := range r.I32Map {
		var tmp = x.GetI32Map().Get(k)
		if tmp == nil {
			tmp = NewTestI32MapSync()
			tmp.MergeDirtyFromPb(v)
			x.GetI32Map().Put(k, tmp)
		} else {
			tmp.MergeDirtyFromPb(v)
		}
	}

	if x.u32Map != nil {
		x.u32Map.RemoveAll(raw.GetUint32List(1016))
	}
	for k, v := range r.U32Map {
		var tmp = x.GetU32Map().Get(k)
		if tmp == nil {
			tmp = NewTestU32MapSync()
			tmp.MergeDirtyFromPb(v)
			x.GetU32Map().Put(k, tmp)
		} else {
			tmp.MergeDirtyFromPb(v)
		}
	}

	if x.i64Map != nil {
		x.i64Map.RemoveAll(raw.GetInt64List(1017))
	}
	for k, v := range r.I64Map {
		var tmp = x.GetI64Map().Get(k)
		if tmp == nil {
			tmp = NewTestI64MapSync()
			tmp.MergeDirtyFromPb(v)
			x.GetI64Map().Put(k, tmp)
		} else {
			tmp.MergeDirtyFromPb(v)
		}
	}

	if x.u64Map != nil {
		x.u64Map.RemoveAll(raw.GetUint64List(1018))
	}
	for k, v := range r.U64Map {
		var tmp = x.GetU64Map().Get(k)
		if tmp == nil {
			tmp = NewTestU64MapSync()
			tmp.MergeDirtyFromPb(v)
			x.GetU64Map().Put(k, tmp)
		} else {
			tmp.MergeDirtyFromPb(v)
		}
	}

	if x.boolMap != nil {
		x.boolMap.RemoveAll(raw.GetBoolList(1019))
	}
	for k, v := range r.BoolMap {
		var tmp = x.GetBoolMap().Get(k)
		if tmp == nil {
			tmp = NewTestBoolMapSync()
			tmp.MergeDirtyFromPb(v)
			x.GetBoolMap().Put(k, tmp)
		} else {
			tmp.MergeDirtyFromPb(v)
		}
	}

	if x.strMap != nil {
		x.strMap.RemoveAll(raw.GetStringList(1021))
	}
	for k, v := range r.StrMap {
		var tmp = x.GetStrMap().Get(k)
		if tmp == nil {
			tmp = NewTestStringMapSync()
			tmp.MergeDirtyFromPb(v)
			x.GetStrMap().Put(k, tmp)
		} else {
			tmp.MergeDirtyFromPb(v)
		}
	}

	if r.F32 != nil {
		x.SetF32(*r.F32)
	}

	if r.F64 != nil {
		x.SetF64(*r.F64)
	}

	if len(r.F32Arr) > 0 || raw.GetBool(1025) {
		x.GetF32Arr().Clear()
		x.f32Arr.AddAll(r.F32Arr)
	}

	if len(r.F64Arr) > 0 || raw.GetBool(1026) {
		x.GetF64Arr().Clear()
		x.f64Arr.AddAll(r.F64Arr)
	}

	return x
}

func (x *TestSync) MergeDirtyFromBytes(buf []byte) {

	vn := make([]int32, 0)
	un := make([]interface{}, 0) // uint64  []byte fixed32 fixed64
	for len(buf) > 0 {
		num, typ, n := protowire.ConsumeTag(buf)
		if n < 0 {
			panic(syncdep.ErrParseRawFields)
		}
		buf = buf[n:]
		switch num {

		case 1008:
			_, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(syncdep.ErrParseRawFields)
			}
			buf = buf[n:]
			x.GetI32Arr().Clear()

		case 1009:
			_, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(syncdep.ErrParseRawFields)
			}
			buf = buf[n:]
			x.GetU32Arr().Clear()

		case 1010:
			_, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(syncdep.ErrParseRawFields)
			}
			buf = buf[n:]
			x.GetI64Arr().Clear()

		case 1011:
			_, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(syncdep.ErrParseRawFields)
			}
			buf = buf[n:]
			x.GetU64Arr().Clear()

		case 1012:
			_, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(syncdep.ErrParseRawFields)
			}
			buf = buf[n:]
			x.GetBoolArr().Clear()

		case 1013:
			_, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(syncdep.ErrParseRawFields)
			}
			buf = buf[n:]
			x.GetEnumArr().Clear()

		case 1014:
			_, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(syncdep.ErrParseRawFields)
			}
			buf = buf[n:]
			x.GetStrArr().Clear()

		case 1015:
			if x.i32Map != nil {
				bys, n := protowire.ConsumeBytes(buf)
				if n < 0 {
					panic(syncdep.ErrParseRawFields)
				}
				buf = buf[n:]

				syncdep.VarintRange(bys, func(val uint64) {

					x.GetI32Map().Remove(int32(val))

				})

			}

		case 1016:
			if x.u32Map != nil {
				bys, n := protowire.ConsumeBytes(buf)
				if n < 0 {
					panic(syncdep.ErrParseRawFields)
				}
				buf = buf[n:]

				syncdep.VarintRange(bys, func(val uint64) {

					x.GetU32Map().Remove(uint32(val))

				})

			}

		case 1017:
			if x.i64Map != nil {
				bys, n := protowire.ConsumeBytes(buf)
				if n < 0 {
					panic(syncdep.ErrParseRawFields)
				}
				buf = buf[n:]

				syncdep.VarintRange(bys, func(val uint64) {

					x.GetI64Map().Remove(int64(val))

				})

			}

		case 1018:
			if x.u64Map != nil {
				bys, n := protowire.ConsumeBytes(buf)
				if n < 0 {
					panic(syncdep.ErrParseRawFields)
				}
				buf = buf[n:]

				syncdep.VarintRange(bys, func(val uint64) {

					x.GetU64Map().Remove(uint64(val))

				})

			}

		case 1019:
			if x.boolMap != nil {
				bys, n := protowire.ConsumeBytes(buf)
				if n < 0 {
					panic(syncdep.ErrParseRawFields)
				}
				buf = buf[n:]

				syncdep.VarintRange(bys, func(val uint64) {

					x.GetBoolMap().Remove(val > 0)

				})

			}

		case 1021:
			if x.strMap != nil {
				bys, n := protowire.ConsumeBytes(buf)
				if n < 0 {
					panic(syncdep.ErrParseRawFields)
				}
				buf = buf[n:]

				x.GetStrMap().Remove(syncdep.Bys2Str(bys))

			}

		case 1025:
			_, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(syncdep.ErrParseRawFields)
			}
			buf = buf[n:]
			x.GetF32Arr().Clear()

		case 1026:
			_, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(syncdep.ErrParseRawFields)
			}
			buf = buf[n:]
			x.GetF64Arr().Clear()

		default:
			vn = append(vn, int32(num))
			var v interface{}
			switch typ {
			case protowire.VarintType:
				v, n = protowire.ConsumeVarint(buf)
			case protowire.Fixed32Type:
				var f32 uint32
				f32, n = protowire.ConsumeFixed32(buf)
				v = math.Float32frombits(f32)
			case protowire.Fixed64Type:
				var f64 uint64
				f64, n = protowire.ConsumeFixed64(buf)
				v = math.Float64frombits(f64)
			case protowire.BytesType:
				v, n = protowire.ConsumeBytes(buf)
			}
			if n < 0 {
				panic(syncdep.ErrParseRawFields)
			}
			buf = buf[n:]
			un = append(un, v)
		}
	}

	var strArrCleared = false

	for i, num := range vn {
		v := un[i]

		switch num {

		case 1:

			x.SetId(int32(v.(uint64)))

		case 2:

			x.SetU32(uint32(v.(uint64)))

		case 3:

			x.SetI64(int64(v.(uint64)))

		case 4:

			x.SetU64(uint64(v.(uint64)))

		case 5:

			x.SetB(v.(uint64) > 0)

		case 6:

			x.SetE(ColorType(v.(uint64)))

		case 7:

			x.SetStr(syncdep.Bys2Str(v.([]byte)))

		case 22:

			x.GetObj().MergeDirtyFromBytes(v.([]byte))

		case 8:

			x.GetI32Arr().Clear()

			syncdep.VarintRange(v.([]byte), func(val uint64) {
				x.GetI32Arr().Add(int32(val))
			})

		case 9:

			x.GetU32Arr().Clear()

			syncdep.VarintRange(v.([]byte), func(val uint64) {
				x.GetU32Arr().Add(uint32(val))
			})

		case 10:

			x.GetI64Arr().Clear()

			syncdep.VarintRange(v.([]byte), func(val uint64) {
				x.GetI64Arr().Add(int64(val))
			})

		case 11:

			x.GetU64Arr().Clear()

			syncdep.VarintRange(v.([]byte), func(val uint64) {
				x.GetU64Arr().Add(uint64(val))
			})

		case 12:

			x.GetBoolArr().Clear()

			syncdep.VarintRange(v.([]byte), func(val uint64) {
				x.GetBoolArr().Add(val > 0)
			})

		case 13:

			x.GetEnumArr().Clear()

			syncdep.VarintRange(v.([]byte), func(val uint64) {
				x.GetEnumArr().Add(ColorType(val))
			})

		case 14:

			if !strArrCleared {
				strArrCleared = true
				x.GetStrArr().Clear()
			}
			x.GetStrArr().Add(syncdep.Bys2Str(v.([]byte)))

		case 15:

			k, bys := syncdep.ParseMap[int32](v.([]byte))
			var tmp = x.GetI32Map().Get(k)
			if tmp == nil {
				tmp = NewTestI32MapSync()
			}
			tmp.MergeDirtyFromBytes(bys)
			x.GetI32Map().Put(k, tmp)

		case 16:

			k, bys := syncdep.ParseMap[uint32](v.([]byte))
			var tmp = x.GetU32Map().Get(k)
			if tmp == nil {
				tmp = NewTestU32MapSync()
			}
			tmp.MergeDirtyFromBytes(bys)
			x.GetU32Map().Put(k, tmp)

		case 17:

			k, bys := syncdep.ParseMap[int64](v.([]byte))
			var tmp = x.GetI64Map().Get(k)
			if tmp == nil {
				tmp = NewTestI64MapSync()
			}
			tmp.MergeDirtyFromBytes(bys)
			x.GetI64Map().Put(k, tmp)

		case 18:

			k, bys := syncdep.ParseMap[uint64](v.([]byte))
			var tmp = x.GetU64Map().Get(k)
			if tmp == nil {
				tmp = NewTestU64MapSync()
			}
			tmp.MergeDirtyFromBytes(bys)
			x.GetU64Map().Put(k, tmp)

		case 19:

			k, bys := syncdep.ParseMap[bool](v.([]byte))
			var tmp = x.GetBoolMap().Get(k)
			if tmp == nil {
				tmp = NewTestBoolMapSync()
			}
			tmp.MergeDirtyFromBytes(bys)
			x.GetBoolMap().Put(k, tmp)

		case 21:

			k, bys := syncdep.ParseMap[string](v.([]byte))
			var tmp = x.GetStrMap().Get(k)
			if tmp == nil {
				tmp = NewTestStringMapSync()
			}
			tmp.MergeDirtyFromBytes(bys)
			x.GetStrMap().Put(k, tmp)

		case 23:

			x.SetF32(v.(float32))

		case 24:

			x.SetF64(v.(float64))

		case 25:

			x.GetF32Arr().Clear()

			syncdep.Fixed32Range(v.([]byte), func(val float32) {
				x.GetF32Arr().Add(val)
			})

		case 26:

			x.GetF64Arr().Clear()

			syncdep.Fixed64Range(v.([]byte), func(val float64) {
				x.GetF64Arr().Add(val)
			})

		}
	}
}

func (x *TestSync) MergeDirtyToBytes() []byte {
	var buf []byte

	if x.isIdDirty() {
		buf = protowire.AppendTag(buf, 1, 0)

		buf = protowire.AppendVarint(buf, uint64(x.id))

	}

	if x.isU32Dirty() {
		buf = protowire.AppendTag(buf, 2, 0)

		buf = protowire.AppendVarint(buf, uint64(x.u32))

	}

	if x.isI64Dirty() {
		buf = protowire.AppendTag(buf, 3, 0)

		buf = protowire.AppendVarint(buf, uint64(x.i64))

	}

	if x.isU64Dirty() {
		buf = protowire.AppendTag(buf, 4, 0)

		buf = protowire.AppendVarint(buf, uint64(x.u64))

	}

	if x.isBDirty() {
		buf = protowire.AppendTag(buf, 5, 0)

		var v uint64 = 0
		if x.b {
			v = 1
		}
		buf = protowire.AppendVarint(buf, v)

	}

	if x.isEDirty() {
		buf = protowire.AppendTag(buf, 6, 0)

		buf = protowire.AppendVarint(buf, uint64(int32(x.e)))

	}

	if x.isStrDirty() {
		buf = protowire.AppendTag(buf, 7, 2)

		buf = protowire.AppendString(buf, x.str)

	}

	if x.isObjDirty() {
		if x.obj != nil {
			bytes := x.obj.MergeDirtyToBytes()
			buf = protowire.AppendTag(buf, 22, 2)
			buf = protowire.AppendBytes(buf, bytes)
		}
	}

	if x.isI32ArrDirty() {
		if x.i32Arr != nil && x.i32Arr.Len() > 0 {

			var packedBuf []byte
			x.i32Arr.Each(func(i int, v int32) bool {

				packedBuf = protowire.AppendVarint(packedBuf, uint64(v))

				return true
			})
			buf = protowire.AppendTag(buf, 8, 2)
			buf = protowire.AppendBytes(buf, packedBuf)

		} else {
			buf = protowire.AppendTag(buf, 1008, protowire.VarintType)
			buf = protowire.AppendVarint(buf, uint64(1))
		}
	}

	if x.isU32ArrDirty() {
		if x.u32Arr != nil && x.u32Arr.Len() > 0 {

			var packedBuf []byte
			x.u32Arr.Each(func(i int, v uint32) bool {

				packedBuf = protowire.AppendVarint(packedBuf, uint64(v))

				return true
			})
			buf = protowire.AppendTag(buf, 9, 2)
			buf = protowire.AppendBytes(buf, packedBuf)

		} else {
			buf = protowire.AppendTag(buf, 1009, protowire.VarintType)
			buf = protowire.AppendVarint(buf, uint64(1))
		}
	}

	if x.isI64ArrDirty() {
		if x.i64Arr != nil && x.i64Arr.Len() > 0 {

			var packedBuf []byte
			x.i64Arr.Each(func(i int, v int64) bool {

				packedBuf = protowire.AppendVarint(packedBuf, uint64(v))

				return true
			})
			buf = protowire.AppendTag(buf, 10, 2)
			buf = protowire.AppendBytes(buf, packedBuf)

		} else {
			buf = protowire.AppendTag(buf, 1010, protowire.VarintType)
			buf = protowire.AppendVarint(buf, uint64(1))
		}
	}

	if x.isU64ArrDirty() {
		if x.u64Arr != nil && x.u64Arr.Len() > 0 {

			var packedBuf []byte
			x.u64Arr.Each(func(i int, v uint64) bool {

				packedBuf = protowire.AppendVarint(packedBuf, uint64(v))

				return true
			})
			buf = protowire.AppendTag(buf, 11, 2)
			buf = protowire.AppendBytes(buf, packedBuf)

		} else {
			buf = protowire.AppendTag(buf, 1011, protowire.VarintType)
			buf = protowire.AppendVarint(buf, uint64(1))
		}
	}

	if x.isBoolArrDirty() {
		if x.boolArr != nil && x.boolArr.Len() > 0 {

			var packedBuf []byte
			x.boolArr.Each(func(i int, v bool) bool {

				var vv uint64 = 0
				if v {
					vv = 1
				}
				packedBuf = protowire.AppendVarint(packedBuf, vv)

				return true
			})
			buf = protowire.AppendTag(buf, 12, 2)
			buf = protowire.AppendBytes(buf, packedBuf)

		} else {
			buf = protowire.AppendTag(buf, 1012, protowire.VarintType)
			buf = protowire.AppendVarint(buf, uint64(1))
		}
	}

	if x.isEnumArrDirty() {
		if x.enumArr != nil && x.enumArr.Len() > 0 {

			var packedBuf []byte
			x.enumArr.Each(func(i int, v ColorType) bool {

				packedBuf = protowire.AppendVarint(packedBuf, uint64(int32(v)))

				return true
			})
			buf = protowire.AppendTag(buf, 13, 2)
			buf = protowire.AppendBytes(buf, packedBuf)

		} else {
			buf = protowire.AppendTag(buf, 1013, protowire.VarintType)
			buf = protowire.AppendVarint(buf, uint64(1))
		}
	}

	if x.isStrArrDirty() {
		if x.strArr != nil && x.strArr.Len() > 0 {

			x.strArr.Each(func(i int, v string) bool {
				buf = protowire.AppendTag(buf, 14, 2)
				buf = protowire.AppendString(buf, v)
				return true
			})

		} else {
			buf = protowire.AppendTag(buf, 1014, protowire.VarintType)
			buf = protowire.AppendVarint(buf, uint64(1))
		}
	}

	if x.isI32MapDirty() {
		if len(x.i32Map.Deleted()) > 0 {

			var packedBuf []byte
			for del := range x.i32Map.Deleted() {
				packedBuf = syncdep.AppendFieldValue(packedBuf, del)
			}
			buf = protowire.AppendTag(buf, 1015, 2)
			buf = protowire.AppendBytes(buf, packedBuf)

		}
		if x.i32Map.Len() > 0 {
			x.i32Map.Each(func(k int32, v *TestI32MapSync) bool {
				if !x.i32Map.ContainDirtied(k) {
					return true
				}
				buf = syncdep.AppendMapFieldKeyValue(buf, 15, k, v.MergeDirtyToBytes())
				return true
			})
		}
	}

	if x.isU32MapDirty() {
		if len(x.u32Map.Deleted()) > 0 {

			var packedBuf []byte
			for del := range x.u32Map.Deleted() {
				packedBuf = syncdep.AppendFieldValue(packedBuf, del)
			}
			buf = protowire.AppendTag(buf, 1016, 2)
			buf = protowire.AppendBytes(buf, packedBuf)

		}
		if x.u32Map.Len() > 0 {
			x.u32Map.Each(func(k uint32, v *TestU32MapSync) bool {
				if !x.u32Map.ContainDirtied(k) {
					return true
				}
				buf = syncdep.AppendMapFieldKeyValue(buf, 16, k, v.MergeDirtyToBytes())
				return true
			})
		}
	}

	if x.isI64MapDirty() {
		if len(x.i64Map.Deleted()) > 0 {

			var packedBuf []byte
			for del := range x.i64Map.Deleted() {
				packedBuf = syncdep.AppendFieldValue(packedBuf, del)
			}
			buf = protowire.AppendTag(buf, 1017, 2)
			buf = protowire.AppendBytes(buf, packedBuf)

		}
		if x.i64Map.Len() > 0 {
			x.i64Map.Each(func(k int64, v *TestI64MapSync) bool {
				if !x.i64Map.ContainDirtied(k) {
					return true
				}
				buf = syncdep.AppendMapFieldKeyValue(buf, 17, k, v.MergeDirtyToBytes())
				return true
			})
		}
	}

	if x.isU64MapDirty() {
		if len(x.u64Map.Deleted()) > 0 {

			var packedBuf []byte
			for del := range x.u64Map.Deleted() {
				packedBuf = syncdep.AppendFieldValue(packedBuf, del)
			}
			buf = protowire.AppendTag(buf, 1018, 2)
			buf = protowire.AppendBytes(buf, packedBuf)

		}
		if x.u64Map.Len() > 0 {
			x.u64Map.Each(func(k uint64, v *TestU64MapSync) bool {
				if !x.u64Map.ContainDirtied(k) {
					return true
				}
				buf = syncdep.AppendMapFieldKeyValue(buf, 18, k, v.MergeDirtyToBytes())
				return true
			})
		}
	}

	if x.isBoolMapDirty() {
		if len(x.boolMap.Deleted()) > 0 {

			var packedBuf []byte
			for del := range x.boolMap.Deleted() {
				packedBuf = syncdep.AppendFieldValue(packedBuf, del)
			}
			buf = protowire.AppendTag(buf, 1019, 2)
			buf = protowire.AppendBytes(buf, packedBuf)

		}
		if x.boolMap.Len() > 0 {
			x.boolMap.Each(func(k bool, v *TestBoolMapSync) bool {
				if !x.boolMap.ContainDirtied(k) {
					return true
				}
				buf = syncdep.AppendMapFieldKeyValue(buf, 19, k, v.MergeDirtyToBytes())
				return true
			})
		}
	}

	if x.isStrMapDirty() {
		if len(x.strMap.Deleted()) > 0 {

			for del := range x.strMap.Deleted() {
				buf = protowire.AppendTag(buf, 1021, protowire.BytesType)
				buf = protowire.AppendString(buf, del)
			}

		}
		if x.strMap.Len() > 0 {
			x.strMap.Each(func(k string, v *TestStringMapSync) bool {
				if !x.strMap.ContainDirtied(k) {
					return true
				}
				buf = syncdep.AppendMapFieldKeyValue(buf, 21, k, v.MergeDirtyToBytes())
				return true
			})
		}
	}

	if x.isF32Dirty() {
		buf = protowire.AppendTag(buf, 23, 5)

		buf = protowire.AppendFixed32(buf, math.Float32bits(x.f32))

	}

	if x.isF64Dirty() {
		buf = protowire.AppendTag(buf, 24, 1)

		buf = protowire.AppendFixed64(buf, math.Float64bits(x.f64))

	}

	if x.isF32ArrDirty() {
		if x.f32Arr != nil && x.f32Arr.Len() > 0 {

			var packedBuf []byte
			x.f32Arr.Each(func(i int, v float32) bool {

				packedBuf = protowire.AppendFixed32(packedBuf, math.Float32bits(v))

				return true
			})
			buf = protowire.AppendTag(buf, 25, 2)
			buf = protowire.AppendBytes(buf, packedBuf)

		} else {
			buf = protowire.AppendTag(buf, 1025, protowire.VarintType)
			buf = protowire.AppendVarint(buf, uint64(1))
		}
	}

	if x.isF64ArrDirty() {
		if x.f64Arr != nil && x.f64Arr.Len() > 0 {

			var packedBuf []byte
			x.f64Arr.Each(func(i int, v float64) bool {

				packedBuf = protowire.AppendFixed64(packedBuf, math.Float64bits(v))

				return true
			})
			buf = protowire.AppendTag(buf, 26, 2)
			buf = protowire.AppendBytes(buf, packedBuf)

		} else {
			buf = protowire.AppendTag(buf, 1026, protowire.VarintType)
			buf = protowire.AppendVarint(buf, uint64(1))
		}
	}

	return buf
}

// struct  TestI32MapSync start

type TestI32MapSync struct {
	id int32

	addition string

	dfm []uint8
	p   syncdep.Sync
	i   int
}

func NewTestI32MapSync() *TestI32MapSync {
	return &TestI32MapSync{
		dfm: make([]uint8, 1),
	}
}

// struct TestI32MapSync Sync interface methods start
func (x *TestI32MapSync) SetDirty(i int, dirty bool, sync syncdep.Sync) {
	idx := i >> 3
	off := i & 7
	if dirty {
		x.dfm[idx] = x.dfm[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dfm[idx] = x.dfm[idx] & ^(1 << off)
	}
}

func (x *TestI32MapSync) SetParentDirty() {
	if x.p != nil {
		x.p.SetDirty(x.i, true, x)
	}
}

func (x *TestI32MapSync) SetParent(sync syncdep.Sync, i int) {
	x.p = sync
	x.i = i
}
func (x *TestI32MapSync) FlushDirty(dirty bool) {

	if dirty || x.isIdDirty() {
		x.setIdDirty(dirty, true)
	}

	if dirty || x.isAdditionDirty() {
		x.setAdditionDirty(dirty, true)
	}

}

func (x *TestI32MapSync) setIdDirty(dirty bool, recur bool) {
	x.SetDirty(1, dirty, x)

}
func (x *TestI32MapSync) isIdDirty() bool {
	return (x.dfm[0] & (1 << 1)) != 0
}

func (x *TestI32MapSync) Key() interface{} {
	return x.id
}
func (x *TestI32MapSync) SetKey(v interface{}) {
	if x.p != nil {
		if _, ok := x.p.(*syncdep.MapSync[int32, *TestI32MapSync]); ok {
			panic("TestI32MapSync in map ,cannot set key")
		}
	}
	x.id = v.(int32)
}

func (x *TestI32MapSync) setAdditionDirty(dirty bool, recur bool) {
	x.SetDirty(2, dirty, x)

}
func (x *TestI32MapSync) isAdditionDirty() bool {
	return (x.dfm[0] & (1 << 2)) != 0
}

// struct TestI32MapSync Sync interface methods end

// struct  TestI32MapSync method clear copy methods start

func (x *TestI32MapSync) Clear() *TestI32MapSync {

	x.SetId(0)

	x.SetAddition("")

	return x
}

func (x *TestI32MapSync) CopyFromPb(r *TestI32Map) *TestI32MapSync {

	if r.Id != nil {
		x.SetId(*r.Id)
	}

	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}

	return x
}

func (x *TestI32MapSync) CopyToPb(r *TestI32Map) *TestI32MapSync {

	r.Id = &x.id

	r.Addition = &x.addition

	return x

}

// struct  TestI32MapSync get set methods start

func (x *TestI32MapSync) GetId() int32 {

	return x.id
}

func (x *TestI32MapSync) SetId(v int32) *TestI32MapSync {

	if x.id == v {
		return x
	}

	x.id = v
	x.setIdDirty(true, false)
	return x
}

func (x *TestI32MapSync) GetAddition() string {

	return x.addition
}

func (x *TestI32MapSync) SetAddition(v string) *TestI32MapSync {

	if x.addition == v {
		return x
	}

	x.addition = v
	x.setAdditionDirty(true, false)
	return x
}

// struct  TestI32MapSync dirty operate methods start

func (x *TestI32MapSync) MergeDirtyToPb(r *TestI32Map) *TestI32MapSync {

	if x.isIdDirty() {
		tmp := x.id
		r.Id = &tmp
	}

	if x.isAdditionDirty() {
		tmp := x.addition
		r.Addition = &tmp
	}

	return x
}

func (x *TestI32MapSync) MergeDirtyFromPb(r *TestI32Map) *TestI32MapSync {

	if r.Id != nil {
		x.SetId(*r.Id)
	}

	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}

	return x
}

func (x *TestI32MapSync) MergeDirtyFromBytes(buf []byte) {

	for len(buf) > 0 {
		num, typ, n := protowire.ConsumeTag(buf)
		if n < 0 {
			panic(syncdep.ErrParseRawFields)
		}
		buf = buf[n:]
		var v interface{}
		switch typ {
		case protowire.VarintType:
			v, n = protowire.ConsumeVarint(buf)
		case protowire.Fixed32Type:
			var f32 uint32
			f32, n = protowire.ConsumeFixed32(buf)
			v = math.Float32frombits(f32)
		case protowire.Fixed64Type:
			var f64 uint64
			f64, n = protowire.ConsumeFixed64(buf)
			v = math.Float64frombits(f64)
		case protowire.BytesType:
			v, n = protowire.ConsumeBytes(buf)
		}
		if n < 0 {
			panic(syncdep.ErrParseRawFields)
		}
		buf = buf[n:]

		switch num {

		case 1:

			x.SetId(int32(v.(uint64)))

		case 2:

			x.SetAddition(syncdep.Bys2Str(v.([]byte)))

		}
	}
}

func (x *TestI32MapSync) MergeDirtyToBytes() []byte {
	var buf []byte

	if x.isIdDirty() {
		buf = protowire.AppendTag(buf, 1, 0)

		buf = protowire.AppendVarint(buf, uint64(x.id))

	}

	if x.isAdditionDirty() {
		buf = protowire.AppendTag(buf, 2, 2)

		buf = protowire.AppendString(buf, x.addition)

	}

	return buf
}

// struct  TestU32MapSync start

type TestU32MapSync struct {
	id uint32

	addition string

	dfm []uint8
	p   syncdep.Sync
	i   int
}

func NewTestU32MapSync() *TestU32MapSync {
	return &TestU32MapSync{
		dfm: make([]uint8, 1),
	}
}

// struct TestU32MapSync Sync interface methods start
func (x *TestU32MapSync) SetDirty(i int, dirty bool, sync syncdep.Sync) {
	idx := i >> 3
	off := i & 7
	if dirty {
		x.dfm[idx] = x.dfm[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dfm[idx] = x.dfm[idx] & ^(1 << off)
	}
}

func (x *TestU32MapSync) SetParentDirty() {
	if x.p != nil {
		x.p.SetDirty(x.i, true, x)
	}
}

func (x *TestU32MapSync) SetParent(sync syncdep.Sync, i int) {
	x.p = sync
	x.i = i
}
func (x *TestU32MapSync) FlushDirty(dirty bool) {

	if dirty || x.isIdDirty() {
		x.setIdDirty(dirty, true)
	}

	if dirty || x.isAdditionDirty() {
		x.setAdditionDirty(dirty, true)
	}

}

func (x *TestU32MapSync) setIdDirty(dirty bool, recur bool) {
	x.SetDirty(1, dirty, x)

}
func (x *TestU32MapSync) isIdDirty() bool {
	return (x.dfm[0] & (1 << 1)) != 0
}

func (x *TestU32MapSync) Key() interface{} {
	return x.id
}
func (x *TestU32MapSync) SetKey(v interface{}) {
	if x.p != nil {
		if _, ok := x.p.(*syncdep.MapSync[uint32, *TestU32MapSync]); ok {
			panic("TestU32MapSync in map ,cannot set key")
		}
	}
	x.id = v.(uint32)
}

func (x *TestU32MapSync) setAdditionDirty(dirty bool, recur bool) {
	x.SetDirty(2, dirty, x)

}
func (x *TestU32MapSync) isAdditionDirty() bool {
	return (x.dfm[0] & (1 << 2)) != 0
}

// struct TestU32MapSync Sync interface methods end

// struct  TestU32MapSync method clear copy methods start

func (x *TestU32MapSync) Clear() *TestU32MapSync {

	x.SetId(0)

	x.SetAddition("")

	return x
}

func (x *TestU32MapSync) CopyFromPb(r *TestU32Map) *TestU32MapSync {

	if r.Id != nil {
		x.SetId(*r.Id)
	}

	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}

	return x
}

func (x *TestU32MapSync) CopyToPb(r *TestU32Map) *TestU32MapSync {

	r.Id = &x.id

	r.Addition = &x.addition

	return x

}

// struct  TestU32MapSync get set methods start

func (x *TestU32MapSync) GetId() uint32 {

	return x.id
}

func (x *TestU32MapSync) SetId(v uint32) *TestU32MapSync {

	if x.id == v {
		return x
	}

	x.id = v
	x.setIdDirty(true, false)
	return x
}

func (x *TestU32MapSync) GetAddition() string {

	return x.addition
}

func (x *TestU32MapSync) SetAddition(v string) *TestU32MapSync {

	if x.addition == v {
		return x
	}

	x.addition = v
	x.setAdditionDirty(true, false)
	return x
}

// struct  TestU32MapSync dirty operate methods start

func (x *TestU32MapSync) MergeDirtyToPb(r *TestU32Map) *TestU32MapSync {

	if x.isIdDirty() {
		tmp := x.id
		r.Id = &tmp
	}

	if x.isAdditionDirty() {
		tmp := x.addition
		r.Addition = &tmp
	}

	return x
}

func (x *TestU32MapSync) MergeDirtyFromPb(r *TestU32Map) *TestU32MapSync {

	if r.Id != nil {
		x.SetId(*r.Id)
	}

	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}

	return x
}

func (x *TestU32MapSync) MergeDirtyFromBytes(buf []byte) {

	for len(buf) > 0 {
		num, typ, n := protowire.ConsumeTag(buf)
		if n < 0 {
			panic(syncdep.ErrParseRawFields)
		}
		buf = buf[n:]
		var v interface{}
		switch typ {
		case protowire.VarintType:
			v, n = protowire.ConsumeVarint(buf)
		case protowire.Fixed32Type:
			var f32 uint32
			f32, n = protowire.ConsumeFixed32(buf)
			v = math.Float32frombits(f32)
		case protowire.Fixed64Type:
			var f64 uint64
			f64, n = protowire.ConsumeFixed64(buf)
			v = math.Float64frombits(f64)
		case protowire.BytesType:
			v, n = protowire.ConsumeBytes(buf)
		}
		if n < 0 {
			panic(syncdep.ErrParseRawFields)
		}
		buf = buf[n:]

		switch num {

		case 1:

			x.SetId(uint32(v.(uint64)))

		case 2:

			x.SetAddition(syncdep.Bys2Str(v.([]byte)))

		}
	}
}

func (x *TestU32MapSync) MergeDirtyToBytes() []byte {
	var buf []byte

	if x.isIdDirty() {
		buf = protowire.AppendTag(buf, 1, 0)

		buf = protowire.AppendVarint(buf, uint64(x.id))

	}

	if x.isAdditionDirty() {
		buf = protowire.AppendTag(buf, 2, 2)

		buf = protowire.AppendString(buf, x.addition)

	}

	return buf
}

// struct  TestI64MapSync start

type TestI64MapSync struct {
	id int64

	addition string

	dfm []uint8
	p   syncdep.Sync
	i   int
}

func NewTestI64MapSync() *TestI64MapSync {
	return &TestI64MapSync{
		dfm: make([]uint8, 1),
	}
}

// struct TestI64MapSync Sync interface methods start
func (x *TestI64MapSync) SetDirty(i int, dirty bool, sync syncdep.Sync) {
	idx := i >> 3
	off := i & 7
	if dirty {
		x.dfm[idx] = x.dfm[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dfm[idx] = x.dfm[idx] & ^(1 << off)
	}
}

func (x *TestI64MapSync) SetParentDirty() {
	if x.p != nil {
		x.p.SetDirty(x.i, true, x)
	}
}

func (x *TestI64MapSync) SetParent(sync syncdep.Sync, i int) {
	x.p = sync
	x.i = i
}
func (x *TestI64MapSync) FlushDirty(dirty bool) {

	if dirty || x.isIdDirty() {
		x.setIdDirty(dirty, true)
	}

	if dirty || x.isAdditionDirty() {
		x.setAdditionDirty(dirty, true)
	}

}

func (x *TestI64MapSync) setIdDirty(dirty bool, recur bool) {
	x.SetDirty(1, dirty, x)

}
func (x *TestI64MapSync) isIdDirty() bool {
	return (x.dfm[0] & (1 << 1)) != 0
}

func (x *TestI64MapSync) Key() interface{} {
	return x.id
}
func (x *TestI64MapSync) SetKey(v interface{}) {
	if x.p != nil {
		if _, ok := x.p.(*syncdep.MapSync[int64, *TestI64MapSync]); ok {
			panic("TestI64MapSync in map ,cannot set key")
		}
	}
	x.id = v.(int64)
}

func (x *TestI64MapSync) setAdditionDirty(dirty bool, recur bool) {
	x.SetDirty(2, dirty, x)

}
func (x *TestI64MapSync) isAdditionDirty() bool {
	return (x.dfm[0] & (1 << 2)) != 0
}

// struct TestI64MapSync Sync interface methods end

// struct  TestI64MapSync method clear copy methods start

func (x *TestI64MapSync) Clear() *TestI64MapSync {

	x.SetId(0)

	x.SetAddition("")

	return x
}

func (x *TestI64MapSync) CopyFromPb(r *TestI64Map) *TestI64MapSync {

	if r.Id != nil {
		x.SetId(*r.Id)
	}

	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}

	return x
}

func (x *TestI64MapSync) CopyToPb(r *TestI64Map) *TestI64MapSync {

	r.Id = &x.id

	r.Addition = &x.addition

	return x

}

// struct  TestI64MapSync get set methods start

func (x *TestI64MapSync) GetId() int64 {

	return x.id
}

func (x *TestI64MapSync) SetId(v int64) *TestI64MapSync {

	if x.id == v {
		return x
	}

	x.id = v
	x.setIdDirty(true, false)
	return x
}

func (x *TestI64MapSync) GetAddition() string {

	return x.addition
}

func (x *TestI64MapSync) SetAddition(v string) *TestI64MapSync {

	if x.addition == v {
		return x
	}

	x.addition = v
	x.setAdditionDirty(true, false)
	return x
}

// struct  TestI64MapSync dirty operate methods start

func (x *TestI64MapSync) MergeDirtyToPb(r *TestI64Map) *TestI64MapSync {

	if x.isIdDirty() {
		tmp := x.id
		r.Id = &tmp
	}

	if x.isAdditionDirty() {
		tmp := x.addition
		r.Addition = &tmp
	}

	return x
}

func (x *TestI64MapSync) MergeDirtyFromPb(r *TestI64Map) *TestI64MapSync {

	if r.Id != nil {
		x.SetId(*r.Id)
	}

	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}

	return x
}

func (x *TestI64MapSync) MergeDirtyFromBytes(buf []byte) {

	for len(buf) > 0 {
		num, typ, n := protowire.ConsumeTag(buf)
		if n < 0 {
			panic(syncdep.ErrParseRawFields)
		}
		buf = buf[n:]
		var v interface{}
		switch typ {
		case protowire.VarintType:
			v, n = protowire.ConsumeVarint(buf)
		case protowire.Fixed32Type:
			var f32 uint32
			f32, n = protowire.ConsumeFixed32(buf)
			v = math.Float32frombits(f32)
		case protowire.Fixed64Type:
			var f64 uint64
			f64, n = protowire.ConsumeFixed64(buf)
			v = math.Float64frombits(f64)
		case protowire.BytesType:
			v, n = protowire.ConsumeBytes(buf)
		}
		if n < 0 {
			panic(syncdep.ErrParseRawFields)
		}
		buf = buf[n:]

		switch num {

		case 1:

			x.SetId(int64(v.(uint64)))

		case 2:

			x.SetAddition(syncdep.Bys2Str(v.([]byte)))

		}
	}
}

func (x *TestI64MapSync) MergeDirtyToBytes() []byte {
	var buf []byte

	if x.isIdDirty() {
		buf = protowire.AppendTag(buf, 1, 0)

		buf = protowire.AppendVarint(buf, uint64(x.id))

	}

	if x.isAdditionDirty() {
		buf = protowire.AppendTag(buf, 2, 2)

		buf = protowire.AppendString(buf, x.addition)

	}

	return buf
}

// struct  TestU64MapSync start

type TestU64MapSync struct {
	id uint64

	addition string

	dfm []uint8
	p   syncdep.Sync
	i   int
}

func NewTestU64MapSync() *TestU64MapSync {
	return &TestU64MapSync{
		dfm: make([]uint8, 1),
	}
}

// struct TestU64MapSync Sync interface methods start
func (x *TestU64MapSync) SetDirty(i int, dirty bool, sync syncdep.Sync) {
	idx := i >> 3
	off := i & 7
	if dirty {
		x.dfm[idx] = x.dfm[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dfm[idx] = x.dfm[idx] & ^(1 << off)
	}
}

func (x *TestU64MapSync) SetParentDirty() {
	if x.p != nil {
		x.p.SetDirty(x.i, true, x)
	}
}

func (x *TestU64MapSync) SetParent(sync syncdep.Sync, i int) {
	x.p = sync
	x.i = i
}
func (x *TestU64MapSync) FlushDirty(dirty bool) {

	if dirty || x.isIdDirty() {
		x.setIdDirty(dirty, true)
	}

	if dirty || x.isAdditionDirty() {
		x.setAdditionDirty(dirty, true)
	}

}

func (x *TestU64MapSync) setIdDirty(dirty bool, recur bool) {
	x.SetDirty(1, dirty, x)

}
func (x *TestU64MapSync) isIdDirty() bool {
	return (x.dfm[0] & (1 << 1)) != 0
}

func (x *TestU64MapSync) Key() interface{} {
	return x.id
}
func (x *TestU64MapSync) SetKey(v interface{}) {
	if x.p != nil {
		if _, ok := x.p.(*syncdep.MapSync[uint64, *TestU64MapSync]); ok {
			panic("TestU64MapSync in map ,cannot set key")
		}
	}
	x.id = v.(uint64)
}

func (x *TestU64MapSync) setAdditionDirty(dirty bool, recur bool) {
	x.SetDirty(2, dirty, x)

}
func (x *TestU64MapSync) isAdditionDirty() bool {
	return (x.dfm[0] & (1 << 2)) != 0
}

// struct TestU64MapSync Sync interface methods end

// struct  TestU64MapSync method clear copy methods start

func (x *TestU64MapSync) Clear() *TestU64MapSync {

	x.SetId(0)

	x.SetAddition("")

	return x
}

func (x *TestU64MapSync) CopyFromPb(r *TestU64Map) *TestU64MapSync {

	if r.Id != nil {
		x.SetId(*r.Id)
	}

	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}

	return x
}

func (x *TestU64MapSync) CopyToPb(r *TestU64Map) *TestU64MapSync {

	r.Id = &x.id

	r.Addition = &x.addition

	return x

}

// struct  TestU64MapSync get set methods start

func (x *TestU64MapSync) GetId() uint64 {

	return x.id
}

func (x *TestU64MapSync) SetId(v uint64) *TestU64MapSync {

	if x.id == v {
		return x
	}

	x.id = v
	x.setIdDirty(true, false)
	return x
}

func (x *TestU64MapSync) GetAddition() string {

	return x.addition
}

func (x *TestU64MapSync) SetAddition(v string) *TestU64MapSync {

	if x.addition == v {
		return x
	}

	x.addition = v
	x.setAdditionDirty(true, false)
	return x
}

// struct  TestU64MapSync dirty operate methods start

func (x *TestU64MapSync) MergeDirtyToPb(r *TestU64Map) *TestU64MapSync {

	if x.isIdDirty() {
		tmp := x.id
		r.Id = &tmp
	}

	if x.isAdditionDirty() {
		tmp := x.addition
		r.Addition = &tmp
	}

	return x
}

func (x *TestU64MapSync) MergeDirtyFromPb(r *TestU64Map) *TestU64MapSync {

	if r.Id != nil {
		x.SetId(*r.Id)
	}

	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}

	return x
}

func (x *TestU64MapSync) MergeDirtyFromBytes(buf []byte) {

	for len(buf) > 0 {
		num, typ, n := protowire.ConsumeTag(buf)
		if n < 0 {
			panic(syncdep.ErrParseRawFields)
		}
		buf = buf[n:]
		var v interface{}
		switch typ {
		case protowire.VarintType:
			v, n = protowire.ConsumeVarint(buf)
		case protowire.Fixed32Type:
			var f32 uint32
			f32, n = protowire.ConsumeFixed32(buf)
			v = math.Float32frombits(f32)
		case protowire.Fixed64Type:
			var f64 uint64
			f64, n = protowire.ConsumeFixed64(buf)
			v = math.Float64frombits(f64)
		case protowire.BytesType:
			v, n = protowire.ConsumeBytes(buf)
		}
		if n < 0 {
			panic(syncdep.ErrParseRawFields)
		}
		buf = buf[n:]

		switch num {

		case 1:

			x.SetId(uint64(v.(uint64)))

		case 2:

			x.SetAddition(syncdep.Bys2Str(v.([]byte)))

		}
	}
}

func (x *TestU64MapSync) MergeDirtyToBytes() []byte {
	var buf []byte

	if x.isIdDirty() {
		buf = protowire.AppendTag(buf, 1, 0)

		buf = protowire.AppendVarint(buf, uint64(x.id))

	}

	if x.isAdditionDirty() {
		buf = protowire.AppendTag(buf, 2, 2)

		buf = protowire.AppendString(buf, x.addition)

	}

	return buf
}

// struct  TestBoolMapSync start

type TestBoolMapSync struct {
	id bool

	addition string

	dfm []uint8
	p   syncdep.Sync
	i   int
}

func NewTestBoolMapSync() *TestBoolMapSync {
	return &TestBoolMapSync{
		dfm: make([]uint8, 1),
	}
}

// struct TestBoolMapSync Sync interface methods start
func (x *TestBoolMapSync) SetDirty(i int, dirty bool, sync syncdep.Sync) {
	idx := i >> 3
	off := i & 7
	if dirty {
		x.dfm[idx] = x.dfm[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dfm[idx] = x.dfm[idx] & ^(1 << off)
	}
}

func (x *TestBoolMapSync) SetParentDirty() {
	if x.p != nil {
		x.p.SetDirty(x.i, true, x)
	}
}

func (x *TestBoolMapSync) SetParent(sync syncdep.Sync, i int) {
	x.p = sync
	x.i = i
}
func (x *TestBoolMapSync) FlushDirty(dirty bool) {

	if dirty || x.isIdDirty() {
		x.setIdDirty(dirty, true)
	}

	if dirty || x.isAdditionDirty() {
		x.setAdditionDirty(dirty, true)
	}

}

func (x *TestBoolMapSync) setIdDirty(dirty bool, recur bool) {
	x.SetDirty(1, dirty, x)

}
func (x *TestBoolMapSync) isIdDirty() bool {
	return (x.dfm[0] & (1 << 1)) != 0
}

func (x *TestBoolMapSync) Key() interface{} {
	return x.id
}
func (x *TestBoolMapSync) SetKey(v interface{}) {
	if x.p != nil {
		if _, ok := x.p.(*syncdep.MapSync[bool, *TestBoolMapSync]); ok {
			panic("TestBoolMapSync in map ,cannot set key")
		}
	}
	x.id = v.(bool)
}

func (x *TestBoolMapSync) setAdditionDirty(dirty bool, recur bool) {
	x.SetDirty(2, dirty, x)

}
func (x *TestBoolMapSync) isAdditionDirty() bool {
	return (x.dfm[0] & (1 << 2)) != 0
}

// struct TestBoolMapSync Sync interface methods end

// struct  TestBoolMapSync method clear copy methods start

func (x *TestBoolMapSync) Clear() *TestBoolMapSync {

	x.SetId(false)

	x.SetAddition("")

	return x
}

func (x *TestBoolMapSync) CopyFromPb(r *TestBoolMap) *TestBoolMapSync {

	if r.Id != nil {
		x.SetId(*r.Id)
	}

	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}

	return x
}

func (x *TestBoolMapSync) CopyToPb(r *TestBoolMap) *TestBoolMapSync {

	r.Id = &x.id

	r.Addition = &x.addition

	return x

}

// struct  TestBoolMapSync get set methods start

func (x *TestBoolMapSync) GetId() bool {

	return x.id
}

func (x *TestBoolMapSync) SetId(v bool) *TestBoolMapSync {

	if x.id == v {
		return x
	}

	x.id = v
	x.setIdDirty(true, false)
	return x
}

func (x *TestBoolMapSync) GetAddition() string {

	return x.addition
}

func (x *TestBoolMapSync) SetAddition(v string) *TestBoolMapSync {

	if x.addition == v {
		return x
	}

	x.addition = v
	x.setAdditionDirty(true, false)
	return x
}

// struct  TestBoolMapSync dirty operate methods start

func (x *TestBoolMapSync) MergeDirtyToPb(r *TestBoolMap) *TestBoolMapSync {

	if x.isIdDirty() {
		tmp := x.id
		r.Id = &tmp
	}

	if x.isAdditionDirty() {
		tmp := x.addition
		r.Addition = &tmp
	}

	return x
}

func (x *TestBoolMapSync) MergeDirtyFromPb(r *TestBoolMap) *TestBoolMapSync {

	if r.Id != nil {
		x.SetId(*r.Id)
	}

	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}

	return x
}

func (x *TestBoolMapSync) MergeDirtyFromBytes(buf []byte) {

	for len(buf) > 0 {
		num, typ, n := protowire.ConsumeTag(buf)
		if n < 0 {
			panic(syncdep.ErrParseRawFields)
		}
		buf = buf[n:]
		var v interface{}
		switch typ {
		case protowire.VarintType:
			v, n = protowire.ConsumeVarint(buf)
		case protowire.Fixed32Type:
			var f32 uint32
			f32, n = protowire.ConsumeFixed32(buf)
			v = math.Float32frombits(f32)
		case protowire.Fixed64Type:
			var f64 uint64
			f64, n = protowire.ConsumeFixed64(buf)
			v = math.Float64frombits(f64)
		case protowire.BytesType:
			v, n = protowire.ConsumeBytes(buf)
		}
		if n < 0 {
			panic(syncdep.ErrParseRawFields)
		}
		buf = buf[n:]

		switch num {

		case 1:

			x.SetId(v.(uint64) > 0)

		case 2:

			x.SetAddition(syncdep.Bys2Str(v.([]byte)))

		}
	}
}

func (x *TestBoolMapSync) MergeDirtyToBytes() []byte {
	var buf []byte

	if x.isIdDirty() {
		buf = protowire.AppendTag(buf, 1, 0)

		var v uint64 = 0
		if x.id {
			v = 1
		}
		buf = protowire.AppendVarint(buf, v)

	}

	if x.isAdditionDirty() {
		buf = protowire.AppendTag(buf, 2, 2)

		buf = protowire.AppendString(buf, x.addition)

	}

	return buf
}

// struct  TestStringMapSync start

type TestStringMapSync struct {
	id string

	addition string

	dfm []uint8
	p   syncdep.Sync
	i   int
}

func NewTestStringMapSync() *TestStringMapSync {
	return &TestStringMapSync{
		dfm: make([]uint8, 1),
	}
}

// struct TestStringMapSync Sync interface methods start
func (x *TestStringMapSync) SetDirty(i int, dirty bool, sync syncdep.Sync) {
	idx := i >> 3
	off := i & 7
	if dirty {
		x.dfm[idx] = x.dfm[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dfm[idx] = x.dfm[idx] & ^(1 << off)
	}
}

func (x *TestStringMapSync) SetParentDirty() {
	if x.p != nil {
		x.p.SetDirty(x.i, true, x)
	}
}

func (x *TestStringMapSync) SetParent(sync syncdep.Sync, i int) {
	x.p = sync
	x.i = i
}
func (x *TestStringMapSync) FlushDirty(dirty bool) {

	if dirty || x.isIdDirty() {
		x.setIdDirty(dirty, true)
	}

	if dirty || x.isAdditionDirty() {
		x.setAdditionDirty(dirty, true)
	}

}

func (x *TestStringMapSync) setIdDirty(dirty bool, recur bool) {
	x.SetDirty(1, dirty, x)

}
func (x *TestStringMapSync) isIdDirty() bool {
	return (x.dfm[0] & (1 << 1)) != 0
}

func (x *TestStringMapSync) Key() interface{} {
	return x.id
}
func (x *TestStringMapSync) SetKey(v interface{}) {
	if x.p != nil {
		if _, ok := x.p.(*syncdep.MapSync[string, *TestStringMapSync]); ok {
			panic("TestStringMapSync in map ,cannot set key")
		}
	}
	x.id = v.(string)
}

func (x *TestStringMapSync) setAdditionDirty(dirty bool, recur bool) {
	x.SetDirty(2, dirty, x)

}
func (x *TestStringMapSync) isAdditionDirty() bool {
	return (x.dfm[0] & (1 << 2)) != 0
}

// struct TestStringMapSync Sync interface methods end

// struct  TestStringMapSync method clear copy methods start

func (x *TestStringMapSync) Clear() *TestStringMapSync {

	x.SetId("")

	x.SetAddition("")

	return x
}

func (x *TestStringMapSync) CopyFromPb(r *TestStringMap) *TestStringMapSync {

	if r.Id != nil {
		x.SetId(*r.Id)
	}

	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}

	return x
}

func (x *TestStringMapSync) CopyToPb(r *TestStringMap) *TestStringMapSync {

	r.Id = &x.id

	r.Addition = &x.addition

	return x

}

// struct  TestStringMapSync get set methods start

func (x *TestStringMapSync) GetId() string {

	return x.id
}

func (x *TestStringMapSync) SetId(v string) *TestStringMapSync {

	if x.id == v {
		return x
	}

	x.id = v
	x.setIdDirty(true, false)
	return x
}

func (x *TestStringMapSync) GetAddition() string {

	return x.addition
}

func (x *TestStringMapSync) SetAddition(v string) *TestStringMapSync {

	if x.addition == v {
		return x
	}

	x.addition = v
	x.setAdditionDirty(true, false)
	return x
}

// struct  TestStringMapSync dirty operate methods start

func (x *TestStringMapSync) MergeDirtyToPb(r *TestStringMap) *TestStringMapSync {

	if x.isIdDirty() {
		tmp := x.id
		r.Id = &tmp
	}

	if x.isAdditionDirty() {
		tmp := x.addition
		r.Addition = &tmp
	}

	return x
}

func (x *TestStringMapSync) MergeDirtyFromPb(r *TestStringMap) *TestStringMapSync {

	if r.Id != nil {
		x.SetId(*r.Id)
	}

	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}

	return x
}

func (x *TestStringMapSync) MergeDirtyFromBytes(buf []byte) {

	for len(buf) > 0 {
		num, typ, n := protowire.ConsumeTag(buf)
		if n < 0 {
			panic(syncdep.ErrParseRawFields)
		}
		buf = buf[n:]
		var v interface{}
		switch typ {
		case protowire.VarintType:
			v, n = protowire.ConsumeVarint(buf)
		case protowire.Fixed32Type:
			var f32 uint32
			f32, n = protowire.ConsumeFixed32(buf)
			v = math.Float32frombits(f32)
		case protowire.Fixed64Type:
			var f64 uint64
			f64, n = protowire.ConsumeFixed64(buf)
			v = math.Float64frombits(f64)
		case protowire.BytesType:
			v, n = protowire.ConsumeBytes(buf)
		}
		if n < 0 {
			panic(syncdep.ErrParseRawFields)
		}
		buf = buf[n:]

		switch num {

		case 1:

			x.SetId(syncdep.Bys2Str(v.([]byte)))

		case 2:

			x.SetAddition(syncdep.Bys2Str(v.([]byte)))

		}
	}
}

func (x *TestStringMapSync) MergeDirtyToBytes() []byte {
	var buf []byte

	if x.isIdDirty() {
		buf = protowire.AppendTag(buf, 1, 2)

		buf = protowire.AppendString(buf, x.id)

	}

	if x.isAdditionDirty() {
		buf = protowire.AppendTag(buf, 2, 2)

		buf = protowire.AppendString(buf, x.addition)

	}

	return buf
}
