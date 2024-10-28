package pbgenv1

import "github.com/yaoguangduan/protosync/syncdep"

import "google.golang.org/protobuf/encoding/protowire"

import "math"

// struct  IntroDetailSync start

type IntroDetailSync struct {
	address string

	money int32

	dfm []uint8
	p   syncdep.Sync
	i   int
}

func NewIntroDetailSync() *IntroDetailSync {
	return &IntroDetailSync{
		dfm: make([]uint8, 1),
	}
}

// struct IntroDetailSync Sync interface methods start
func (x *IntroDetailSync) SetDirty(i int, dirty bool, sync syncdep.Sync) {
	idx := i >> 3
	off := i & 7
	if dirty {
		x.dfm[idx] = x.dfm[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dfm[idx] = x.dfm[idx] & ^(1 << off)
	}
}

func (x *IntroDetailSync) SetParentDirty() {
	if x.p != nil {
		x.p.SetDirty(x.i, true, x)
	}
}

func (x *IntroDetailSync) SetParent(sync syncdep.Sync, i int) {
	x.p = sync
	x.i = i
}
func (x *IntroDetailSync) FlushDirty(dirty bool) {

	if dirty || x.isAddressDirty() {
		x.setAddressDirty(dirty, true)
	}

	if dirty || x.isMoneyDirty() {
		x.setMoneyDirty(dirty, true)
	}

}

func (x *IntroDetailSync) setAddressDirty(dirty bool, recur bool) {
	x.SetDirty(1, dirty, x)

}
func (x *IntroDetailSync) isAddressDirty() bool {
	return (x.dfm[0] & (1 << 1)) != 0
}

func (x *IntroDetailSync) setMoneyDirty(dirty bool, recur bool) {
	x.SetDirty(2, dirty, x)

}
func (x *IntroDetailSync) isMoneyDirty() bool {
	return (x.dfm[0] & (1 << 2)) != 0
}

func (x *IntroDetailSync) Key() interface{} {
	return nil
}
func (x *IntroDetailSync) SetKey(v interface{}) {
}

// struct IntroDetailSync Sync interface methods end

// struct  IntroDetailSync method clear copy methods start

func (x *IntroDetailSync) Clear() *IntroDetailSync {

	x.SetAddress("")

	x.SetMoney(0)

	return x
}

func (x *IntroDetailSync) CopyFromPb(r *IntroDetail) *IntroDetailSync {

	if r.Address != nil {
		x.SetAddress(*r.Address)
	}

	if r.Money != nil {
		x.SetMoney(*r.Money)
	}

	return x
}

func (x *IntroDetailSync) CopyToPb(r *IntroDetail) *IntroDetailSync {

	r.Address = &x.address

	r.Money = &x.money

	return x

}

// struct  IntroDetailSync get set methods start

func (x *IntroDetailSync) GetAddress() string {

	return x.address
}

func (x *IntroDetailSync) SetAddress(v string) *IntroDetailSync {

	if x.address == v {
		return x
	}

	x.address = v
	x.setAddressDirty(true, false)
	return x
}

func (x *IntroDetailSync) GetMoney() int32 {

	return x.money
}

func (x *IntroDetailSync) SetMoney(v int32) *IntroDetailSync {

	if x.money == v {
		return x
	}

	x.money = v
	x.setMoneyDirty(true, false)
	return x
}

// struct  IntroDetailSync dirty operate methods start

func (x *IntroDetailSync) MergeDirtyToPb(r *IntroDetail) *IntroDetailSync {

	if x.isAddressDirty() {
		tmp := x.address
		r.Address = &tmp
	}

	if x.isMoneyDirty() {
		tmp := x.money
		r.Money = &tmp
	}

	return x
}

func (x *IntroDetailSync) MergeDirtyFromPb(r *IntroDetail) *IntroDetailSync {

	if r.Address != nil {
		x.SetAddress(*r.Address)
	}

	if r.Money != nil {
		x.SetMoney(*r.Money)
	}

	return x
}

func (x *IntroDetailSync) MergeDirtyFromBytes(buf []byte) {

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

			x.SetAddress(syncdep.Bys2Str(v.([]byte)))

		case 2:

			x.SetMoney(int32(v.(uint64)))

		}
	}
}

func (x *IntroDetailSync) MergeDirtyToBytes() []byte {
	var buf []byte

	if x.isAddressDirty() {
		buf = protowire.AppendTag(buf, 1, 2)

		buf = protowire.AppendString(buf, x.address)

	}

	if x.isMoneyDirty() {
		buf = protowire.AppendTag(buf, 2, 0)

		buf = protowire.AppendVarint(buf, uint64(x.money))

	}

	return buf
}
