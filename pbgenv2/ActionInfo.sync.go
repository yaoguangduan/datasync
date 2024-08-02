package pbgenv2

import "github.com/yaoguangduan/datasync/syncdep"

import "google.golang.org/protobuf/encoding/protowire"

import "math"

// struct  ActionInfoSync start

type ActionInfoSync struct {
	act string

	detail string

	time int64

	dfm []uint8
	p   syncdep.Sync
	i   int
}

func NewActionInfoSync() *ActionInfoSync {
	return &ActionInfoSync{
		dfm: make([]uint8, 1),
	}
}

// struct ActionInfoSync Sync interface methods start
func (x *ActionInfoSync) SetDirty(i int, dirty bool, sync syncdep.Sync) {
	idx := i >> 3
	off := i & 7
	if dirty {
		x.dfm[idx] = x.dfm[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dfm[idx] = x.dfm[idx] & ^(1 << off)
	}
}

func (x *ActionInfoSync) SetParentDirty() {
	if x.p != nil {
		x.p.SetDirty(x.i, true, x)
	}
}

func (x *ActionInfoSync) SetParent(sync syncdep.Sync, i int) {
	x.p = sync
	x.i = i
}
func (x *ActionInfoSync) FlushDirty(dirty bool) {

	if dirty || x.isActDirty() {
		x.setActDirty(dirty, true)
	}

	if dirty || x.isDetailDirty() {
		x.setDetailDirty(dirty, true)
	}

	if dirty || x.isTimeDirty() {
		x.setTimeDirty(dirty, true)
	}

}

func (x *ActionInfoSync) setActDirty(dirty bool, recur bool) {
	x.SetDirty(1, dirty, x)

}
func (x *ActionInfoSync) isActDirty() bool {
	return (x.dfm[0] & (1 << 1)) != 0
}

func (x *ActionInfoSync) Key() interface{} {
	return x.act
}
func (x *ActionInfoSync) SetKey(v interface{}) {
	if x.p != nil {
		if _, ok := x.p.(*syncdep.MapSync[string, *ActionInfoSync]); ok {
			panic("ActionInfoSync in map ,cannot set key")
		}
	}
	x.act = v.(string)
}

func (x *ActionInfoSync) setDetailDirty(dirty bool, recur bool) {
	x.SetDirty(2, dirty, x)

}
func (x *ActionInfoSync) isDetailDirty() bool {
	return (x.dfm[0] & (1 << 2)) != 0
}

func (x *ActionInfoSync) setTimeDirty(dirty bool, recur bool) {
	x.SetDirty(3, dirty, x)

}
func (x *ActionInfoSync) isTimeDirty() bool {
	return (x.dfm[0] & (1 << 3)) != 0
}

// struct ActionInfoSync Sync interface methods end

// struct  ActionInfoSync method clear copy methods start

func (x *ActionInfoSync) Clear() *ActionInfoSync {

	x.SetAct("")

	x.SetDetail("")

	x.SetTime(0)

	return x
}

func (x *ActionInfoSync) CopyFromPb(r *ActionInfo) *ActionInfoSync {

	if r.Act != nil {
		x.SetAct(*r.Act)
	}

	if r.Detail != nil {
		x.SetDetail(*r.Detail)
	}

	if r.Time != nil {
		x.SetTime(*r.Time)
	}

	return x
}

func (x *ActionInfoSync) CopyToPb(r *ActionInfo) *ActionInfoSync {

	r.Act = &x.act

	r.Detail = &x.detail

	r.Time = &x.time

	return x

}

// struct  ActionInfoSync get set methods start

func (x *ActionInfoSync) GetAct() string {

	return x.act
}

func (x *ActionInfoSync) SetAct(v string) *ActionInfoSync {

	if x.act == v {
		return x
	}

	x.act = v
	x.setActDirty(true, false)
	return x
}

func (x *ActionInfoSync) GetDetail() string {

	return x.detail
}

func (x *ActionInfoSync) SetDetail(v string) *ActionInfoSync {

	if x.detail == v {
		return x
	}

	x.detail = v
	x.setDetailDirty(true, false)
	return x
}

func (x *ActionInfoSync) GetTime() int64 {

	return x.time
}

func (x *ActionInfoSync) SetTime(v int64) *ActionInfoSync {

	if x.time == v {
		return x
	}

	x.time = v
	x.setTimeDirty(true, false)
	return x
}

// struct  ActionInfoSync dirty operate methods start

func (x *ActionInfoSync) MergeDirtyToPb(r *ActionInfo) *ActionInfoSync {

	if x.isActDirty() {
		tmp := x.act
		r.Act = &tmp
	}

	if x.isDetailDirty() {
		tmp := x.detail
		r.Detail = &tmp
	}

	if x.isTimeDirty() {
		tmp := x.time
		r.Time = &tmp
	}

	return x
}

func (x *ActionInfoSync) MergeDirtyFromPb(r *ActionInfo) *ActionInfoSync {

	if r.Act != nil {
		x.SetAct(*r.Act)
	}

	if r.Detail != nil {
		x.SetDetail(*r.Detail)
	}

	if r.Time != nil {
		x.SetTime(*r.Time)
	}

	return x
}

func (x *ActionInfoSync) MergeDirtyFromBytes(buf []byte) {

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

			x.SetAct(syncdep.Bys2Str(v.([]byte)))

		case 2:

			x.SetDetail(syncdep.Bys2Str(v.([]byte)))

		case 3:

			x.SetTime(int64(v.(uint64)))

		}
	}
}

func (x *ActionInfoSync) MergeDirtyToBytes() []byte {
	var buf []byte

	if x.isActDirty() {
		buf = protowire.AppendTag(buf, 1, 2)

		buf = protowire.AppendString(buf, x.act)

	}

	if x.isDetailDirty() {
		buf = protowire.AppendTag(buf, 2, 2)

		buf = protowire.AppendString(buf, x.detail)

	}

	if x.isTimeDirty() {
		buf = protowire.AppendTag(buf, 3, 0)

		buf = protowire.AppendVarint(buf, uint64(x.time))

	}

	return buf
}
