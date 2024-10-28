package pbgenv1

import "github.com/yaoguangduan/protosync/syncdep"

import "google.golang.org/protobuf/encoding/protowire"

import "math"

import "slices"

// struct  PersonSync start

type PersonSync struct {
	age int32

	vipLevel VipLevel

	name string

	actions *syncdep.MapSync[string, *ActionInfoSync]

	favor *syncdep.ArraySync[string]

	loveSeq *syncdep.ArraySync[ColorType]

	isGirl bool

	detail *IntroDetailSync

	data []byte

	dfm []uint8
	p   syncdep.Sync
	i   int
}

func NewPersonSync() *PersonSync {
	return &PersonSync{
		dfm: make([]uint8, 2),
	}
}

// struct PersonSync Sync interface methods start
func (x *PersonSync) SetDirty(i int, dirty bool, sync syncdep.Sync) {
	idx := i >> 3
	off := i & 7
	if dirty {
		x.dfm[idx] = x.dfm[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dfm[idx] = x.dfm[idx] & ^(1 << off)
	}
}

func (x *PersonSync) SetParentDirty() {
	if x.p != nil {
		x.p.SetDirty(x.i, true, x)
	}
}

func (x *PersonSync) SetParent(sync syncdep.Sync, i int) {
	x.p = sync
	x.i = i
}
func (x *PersonSync) FlushDirty(dirty bool) {

	if dirty || x.isAgeDirty() {
		x.setAgeDirty(dirty, true)
	}

	if dirty || x.isVipLevelDirty() {
		x.setVipLevelDirty(dirty, true)
	}

	if dirty || x.isNameDirty() {
		x.setNameDirty(dirty, true)
	}

	if dirty || x.isActionsDirty() {
		x.setActionsDirty(dirty, true)
	}

	if dirty || x.isFavorDirty() {
		x.setFavorDirty(dirty, true)
	}

	if dirty || x.isLoveSeqDirty() {
		x.setLoveSeqDirty(dirty, true)
	}

	if dirty || x.isIsGirlDirty() {
		x.setIsGirlDirty(dirty, true)
	}

	if dirty || x.isDetailDirty() {
		x.setDetailDirty(dirty, true)
	}

	if dirty || x.isDataDirty() {
		x.setDataDirty(dirty, true)
	}

}

func (x *PersonSync) setAgeDirty(dirty bool, recur bool) {
	x.SetDirty(1, dirty, x)

}
func (x *PersonSync) isAgeDirty() bool {
	return (x.dfm[0] & (1 << 1)) != 0
}

func (x *PersonSync) setVipLevelDirty(dirty bool, recur bool) {
	x.SetDirty(2, dirty, x)

}
func (x *PersonSync) isVipLevelDirty() bool {
	return (x.dfm[0] & (1 << 2)) != 0
}

func (x *PersonSync) setNameDirty(dirty bool, recur bool) {
	x.SetDirty(3, dirty, x)

}
func (x *PersonSync) isNameDirty() bool {
	return (x.dfm[0] & (1 << 3)) != 0
}

func (x *PersonSync) setActionsDirty(dirty bool, recur bool) {
	x.SetDirty(4, dirty, x)

	if recur && x.actions != nil {
		x.actions.FlushDirty(dirty)
	}

}
func (x *PersonSync) isActionsDirty() bool {
	return (x.dfm[0] & (1 << 4)) != 0
}

func (x *PersonSync) setFavorDirty(dirty bool, recur bool) {
	x.SetDirty(5, dirty, x)

}
func (x *PersonSync) isFavorDirty() bool {
	return (x.dfm[0] & (1 << 5)) != 0
}

func (x *PersonSync) setLoveSeqDirty(dirty bool, recur bool) {
	x.SetDirty(6, dirty, x)

}
func (x *PersonSync) isLoveSeqDirty() bool {
	return (x.dfm[0] & (1 << 6)) != 0
}

func (x *PersonSync) setIsGirlDirty(dirty bool, recur bool) {
	x.SetDirty(7, dirty, x)

}
func (x *PersonSync) isIsGirlDirty() bool {
	return (x.dfm[0] & (1 << 7)) != 0
}

func (x *PersonSync) setDetailDirty(dirty bool, recur bool) {
	x.SetDirty(8, dirty, x)

	if recur && x.detail != nil {
		x.detail.FlushDirty(dirty)
	}

}
func (x *PersonSync) isDetailDirty() bool {
	return (x.dfm[1] & (1 << 0)) != 0
}

func (x *PersonSync) setDataDirty(dirty bool, recur bool) {
	x.SetDirty(9, dirty, x)

}
func (x *PersonSync) isDataDirty() bool {
	return (x.dfm[1] & (1 << 1)) != 0
}

func (x *PersonSync) Key() interface{} {
	return nil
}
func (x *PersonSync) SetKey(v interface{}) {
}

// struct PersonSync Sync interface methods end

// struct  PersonSync method clear copy methods start

func (x *PersonSync) Clear() *PersonSync {

	x.SetAge(0)

	x.SetVipLevel(VipLevel_Level1)

	x.SetName("")

	if x.actions != nil {
		x.actions.Clear()
	}

	if x.favor != nil {
		x.favor.Clear()
	}

	if x.loveSeq != nil {
		x.loveSeq.Clear()
	}

	x.SetIsGirl(false)

	if x.detail != nil {
		x.detail.Clear()
	}

	x.SetData(make([]byte, 0))

	return x
}

func (x *PersonSync) CopyFromPb(r *Person) *PersonSync {

	if r.Age != nil {
		x.SetAge(*r.Age)
	}

	if r.VipLevel != nil {
		x.SetVipLevel(*r.VipLevel)
	}

	if r.Name != nil {
		x.SetName(*r.Name)
	}

	for _, v := range r.Actions {
		if v != nil {
			vv := NewActionInfoSync()
			vv.CopyFromPb(v)
			x.GetActions().Put(vv)
		}
	}

	if len(r.Favor) > 0 {
		x.GetFavor().AddAll(r.Favor)
	}

	if len(r.LoveSeq) > 0 {
		x.GetLoveSeq().AddAll(r.LoveSeq)
	}

	if r.IsGirl != nil {
		x.SetIsGirl(*r.IsGirl)
	}

	if r.Detail != nil {
		x.GetDetail().CopyFromPb(r.Detail)
	}

	x.SetData(slices.Clone(r.Data))

	return x
}

func (x *PersonSync) CopyToPb(r *Person) *PersonSync {

	r.Age = &x.age

	r.VipLevel = &x.vipLevel

	r.Name = &x.name

	if x.actions != nil && x.actions.Len() > 0 {
		tmp := make(map[string]*ActionInfo)
		x.actions.Each(func(k string, v *ActionInfoSync) bool {
			tmpV := &ActionInfo{}
			v.CopyToPb(tmpV)
			tmp[k] = tmpV
			return true
		})
		r.Actions = tmp
	}

	if x.favor != nil && x.favor.Len() > 0 {
		r.Favor = x.favor.ValueView()
	}

	if x.loveSeq != nil && x.loveSeq.Len() > 0 {
		r.LoveSeq = x.loveSeq.ValueView()
	}

	r.IsGirl = &x.isGirl

	if x.detail != nil {
		tmpV := &IntroDetail{}
		x.detail.CopyToPb(tmpV)
		r.Detail = tmpV
	}

	r.Data = slices.Clone(x.data)

	return x

}

// struct  PersonSync get set methods start

func (x *PersonSync) GetAge() int32 {

	return x.age
}

func (x *PersonSync) SetAge(v int32) *PersonSync {

	if x.age == v {
		return x
	}

	x.age = v
	x.setAgeDirty(true, false)
	return x
}

func (x *PersonSync) GetVipLevel() VipLevel {

	return x.vipLevel
}

func (x *PersonSync) SetVipLevel(v VipLevel) *PersonSync {

	if x.vipLevel == v {
		return x
	}

	x.vipLevel = v
	x.setVipLevelDirty(true, false)
	return x
}

func (x *PersonSync) GetName() string {

	return x.name
}

func (x *PersonSync) SetName(v string) *PersonSync {

	if x.name == v {
		return x
	}

	x.name = v
	x.setNameDirty(true, false)
	return x
}

func (x *PersonSync) GetActions() *syncdep.MapSync[string, *ActionInfoSync] {

	if x.actions == nil {
		x.actions = syncdep.NewMapSync[string, *ActionInfoSync]()
		x.actions.SetParent(x, 4)
	}

	return x.actions
}

func (x *PersonSync) GetFavor() *syncdep.ArraySync[string] {

	if x.favor == nil {
		x.favor = syncdep.NewArraySync[string]()
		x.favor.SetParent(x, 5)
	}

	return x.favor
}

func (x *PersonSync) GetLoveSeq() *syncdep.ArraySync[ColorType] {

	if x.loveSeq == nil {
		x.loveSeq = syncdep.NewArraySync[ColorType]()
		x.loveSeq.SetParent(x, 6)
	}

	return x.loveSeq
}

func (x *PersonSync) GetIsGirl() bool {

	return x.isGirl
}

func (x *PersonSync) SetIsGirl(v bool) *PersonSync {

	if x.isGirl == v {
		return x
	}

	x.isGirl = v
	x.setIsGirlDirty(true, false)
	return x
}

func (x *PersonSync) GetDetail() *IntroDetailSync {

	if x.detail == nil {
		x.detail = NewIntroDetailSync()
		x.detail.SetParent(x, 8)
	}

	return x.detail
}

func (x *PersonSync) SetDetail(v *IntroDetailSync) *PersonSync {

	v.SetParent(x, 8)
	if x.detail != nil {
		x.detail.SetParent(nil, -1)
	}

	x.detail = v
	x.setDetailDirty(true, false)
	return x
}

func (x *PersonSync) GetData() []byte {

	return x.data
}

func (x *PersonSync) SetData(v []byte) *PersonSync {

	x.data = v
	x.setDataDirty(true, false)
	return x
}

// struct  PersonSync dirty operate methods start

func (x *PersonSync) MergeDirtyToPb(r *Person) *PersonSync {

	var raw = syncdep.ToRawMessage(r.ProtoReflect().GetUnknown())

	if x.isAgeDirty() {
		tmp := x.age
		r.Age = &tmp
	}

	if x.isVipLevelDirty() {
		tmp := x.vipLevel
		r.VipLevel = &tmp
	}

	if x.isNameDirty() {
		tmp := x.name
		r.Name = &tmp
	}

	if x.isActionsDirty() {
		updated := make([]string, 0)
		if r.Actions != nil {
			for k := range r.Actions {
				if x.actions.ContainDeleted(k) {
					delete(r.Actions, k)
				}
				if x.actions.ContainDirtied(k) {
					updated = append(updated, k)
					tmp := x.actions.Get(k)
					if r.Actions[k] == nil {
						r.Actions[k] = &ActionInfo{}
					}
					tmp.MergeDirtyToPb(r.Actions[k])
				}
			}
		} else {
			r.Actions = make(map[string]*ActionInfo)
		}
		for k := range x.actions.Dirtied() {
			if !slices.Contains(updated, k) {
				tmp := x.actions.Get(k)
				if r.Actions[k] == nil {
					r.Actions[k] = &ActionInfo{}
				}
				tmp.MergeDirtyToPb(r.Actions[k])
			}
		}

		for k := range x.actions.Deleted() {
			raw.AddString(1004, k)
		}

	}

	if x.isFavorDirty() {
		count := x.favor.Len()
		r.Favor = make([]string, 0)
		if count > 0 {
			raw.ClearBool(1005)
			r.Favor = append(r.Favor, x.favor.ValueView()...)
		} else {
			raw.SetBool(1005)
		}
	}

	if x.isLoveSeqDirty() {
		count := x.loveSeq.Len()
		r.LoveSeq = make([]ColorType, 0)
		if count > 0 {
			raw.ClearBool(1006)
			r.LoveSeq = append(r.LoveSeq, x.loveSeq.ValueView()...)
		} else {
			raw.SetBool(1006)
		}
	}

	if x.isIsGirlDirty() {
		tmp := x.isGirl
		r.IsGirl = &tmp
	}

	if x.isDetailDirty() {
		if r.Detail == nil {
			r.Detail = &IntroDetail{}
		}
		x.detail.MergeDirtyToPb(r.Detail)
	}

	if x.isDataDirty() {
		r.Data = slices.Clone(x.data)
	}

	r.ProtoReflect().SetUnknown(raw.Marshal())

	return x
}

func (x *PersonSync) MergeDirtyFromPb(r *Person) *PersonSync {

	var raw = syncdep.ToRawMessage(r.ProtoReflect().GetUnknown())

	if r.Age != nil {
		x.SetAge(*r.Age)
	}

	if r.VipLevel != nil {
		x.SetVipLevel(*r.VipLevel)
	}

	if r.Name != nil {
		x.SetName(*r.Name)
	}

	if x.actions != nil {
		x.actions.RemoveAll(raw.GetStringList(1004))
	}
	for k, v := range r.Actions {
		var tmp = x.GetActions().Get(k)
		if tmp == nil {
			tmp = NewActionInfoSync()
			tmp.MergeDirtyFromPb(v)
			x.GetActions().Put(tmp)
		} else {
			tmp.MergeDirtyFromPb(v)
		}
	}

	if len(r.Favor) > 0 || raw.GetBool(1005) {
		x.GetFavor().Clear()
		x.favor.AddAll(r.Favor)
	}

	if len(r.LoveSeq) > 0 || raw.GetBool(1006) {
		x.GetLoveSeq().Clear()
		x.loveSeq.AddAll(r.LoveSeq)
	}

	if r.IsGirl != nil {
		x.SetIsGirl(*r.IsGirl)
	}

	if r.Detail != nil {
		x.GetDetail().MergeDirtyFromPb(r.Detail)
	}

	if len(r.Data) > 0 {
		x.SetData(slices.Clone(r.Data))
	}

	return x
}

func (x *PersonSync) MergeDirtyFromBytes(buf []byte) {

	vn := make([]int32, 0)
	un := make([]interface{}, 0) // uint64  []byte fixed32 fixed64
	for len(buf) > 0 {
		num, typ, n := protowire.ConsumeTag(buf)
		if n < 0 {
			panic(syncdep.ErrParseRawFields)
		}
		buf = buf[n:]
		switch num {

		case 1004:
			if x.actions != nil {
				bys, n := protowire.ConsumeBytes(buf)
				if n < 0 {
					panic(syncdep.ErrParseRawFields)
				}
				buf = buf[n:]

				x.GetActions().Remove(syncdep.Bys2Str(bys))

			}

		case 1005:
			_, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(syncdep.ErrParseRawFields)
			}
			buf = buf[n:]
			x.GetFavor().Clear()

		case 1006:
			_, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				panic(syncdep.ErrParseRawFields)
			}
			buf = buf[n:]
			x.GetLoveSeq().Clear()

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

	var favorCleared = false

	for i, num := range vn {
		v := un[i]

		switch num {

		case 1:

			x.SetAge(int32(v.(uint64)))

		case 2:

			x.SetVipLevel(VipLevel(v.(uint64)))

		case 3:

			x.SetName(syncdep.Bys2Str(v.([]byte)))

		case 4:

			k, bys := syncdep.ParseMap[string](v.([]byte))
			var tmp = x.GetActions().Get(k)
			if tmp == nil {
				tmp = NewActionInfoSync()
			}
			tmp.MergeDirtyFromBytes(bys)
			x.GetActions().Put(tmp)

		case 5:

			if !favorCleared {
				favorCleared = true
				x.GetFavor().Clear()
			}
			x.GetFavor().Add(syncdep.Bys2Str(v.([]byte)))

		case 6:

			x.GetLoveSeq().Clear()

			syncdep.VarintRange(v.([]byte), func(val uint64) {
				x.GetLoveSeq().Add(ColorType(val))
			})

		case 7:

			x.SetIsGirl(v.(uint64) > 0)

		case 8:

			x.GetDetail().MergeDirtyFromBytes(v.([]byte))

		case 9:

			x.SetData(v.([]byte))

		}
	}
}

func (x *PersonSync) MergeDirtyToBytes() []byte {
	var buf []byte

	if x.isAgeDirty() {
		buf = protowire.AppendTag(buf, 1, 0)

		buf = protowire.AppendVarint(buf, uint64(x.age))

	}

	if x.isVipLevelDirty() {
		buf = protowire.AppendTag(buf, 2, 0)

		buf = protowire.AppendVarint(buf, uint64(int32(x.vipLevel)))

	}

	if x.isNameDirty() {
		buf = protowire.AppendTag(buf, 3, 2)

		buf = protowire.AppendString(buf, x.name)

	}

	if x.isActionsDirty() {
		if len(x.actions.Deleted()) > 0 {

			for del := range x.actions.Deleted() {
				buf = protowire.AppendTag(buf, 1004, protowire.BytesType)
				buf = protowire.AppendString(buf, del)
			}

		}
		if x.actions.Len() > 0 {
			x.actions.Each(func(k string, v *ActionInfoSync) bool {
				if !x.actions.ContainDirtied(k) {
					return true
				}
				buf = syncdep.AppendMapFieldKeyValue(buf, 4, k, v.MergeDirtyToBytes())
				return true
			})
		}
	}

	if x.isFavorDirty() {
		if x.favor != nil && x.favor.Len() > 0 {

			x.favor.Each(func(i int, v string) bool {
				buf = protowire.AppendTag(buf, 5, 2)
				buf = protowire.AppendString(buf, v)
				return true
			})

		} else {
			buf = protowire.AppendTag(buf, 1005, protowire.VarintType)
			buf = protowire.AppendVarint(buf, uint64(1))
		}
	}

	if x.isLoveSeqDirty() {
		if x.loveSeq != nil && x.loveSeq.Len() > 0 {

			var packedBuf []byte
			x.loveSeq.Each(func(i int, v ColorType) bool {

				packedBuf = protowire.AppendVarint(packedBuf, uint64(int32(v)))

				return true
			})
			buf = protowire.AppendTag(buf, 6, 2)
			buf = protowire.AppendBytes(buf, packedBuf)

		} else {
			buf = protowire.AppendTag(buf, 1006, protowire.VarintType)
			buf = protowire.AppendVarint(buf, uint64(1))
		}
	}

	if x.isIsGirlDirty() {
		buf = protowire.AppendTag(buf, 7, 0)

		var v uint64 = 0
		if x.isGirl {
			v = 1
		}
		buf = protowire.AppendVarint(buf, v)

	}

	if x.isDetailDirty() {
		if x.detail != nil {
			bytes := x.detail.MergeDirtyToBytes()
			buf = protowire.AppendTag(buf, 8, 2)
			buf = protowire.AppendBytes(buf, bytes)
		}
	}

	if x.isDataDirty() {
		buf = protowire.AppendTag(buf, 9, 2)

		buf = protowire.AppendBytes(buf, x.data)

	}

	return buf
}
