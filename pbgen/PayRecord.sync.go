package pbgen

//this file is generated by gsync, do not modify it manually !!!

import "github.com/yaoguangduan/protosync/syncdep"
import "google.golang.org/protobuf/encoding/protowire"

type PayRecordSync struct {
	timestamp      int64
	timestampINDEX int
	classic        string
	classicINDEX   int
	dirtyFieldMark []uint8
	parent         syncdep.Sync
	indexInParent  int
}

func NewPayRecordSync() *PayRecordSync {
	return &PayRecordSync{
		timestampINDEX: 0,
		classicINDEX:   1,
		dirtyFieldMark: make([]uint8, 1),
	}
}
func (x *PayRecordSync) Clear() *PayRecordSync {
	x.SetTimestamp(0)
	x.SetClassic("")
	return x
}
func (x *PayRecordSync) CopyToPb(r *PayRecord) *PayRecordSync {
	r.SetTimestamp(x.timestamp)
	r.SetClassic(x.classic)
	return x
}
func (x *PayRecordSync) CopyFromPb(r *PayRecord) *PayRecordSync {
	if r.Timestamp != nil {
		x.SetTimestamp(*r.Timestamp)
	}
	if r.Classic != nil {
		x.SetClassic(*r.Classic)
	}
	return x
}
func (x *PayRecordSync) MergeDirtyFromPb(r *PayRecord) {
	if r.Timestamp != nil {
		x.SetTimestamp(*r.Timestamp)
	}
	if r.Classic != nil {
		x.SetClassic(*r.Classic)
	}
}
func (x *PayRecordSync) MergeDirtyFromBytes(buf []byte) *PayRecordSync {
	fds := syncdep.PreParseProtoBytes(buf)
	for _, rawF := range fds.Values {
		switch rawF.Number {
		case 1:
			x.SetTimestamp(int64(rawF.Value.(uint64)))
		case 2:
			x.SetClassic(syncdep.Bys2Str(rawF.Value.([]byte)))
		}
	}
	return x
}
func (x *PayRecordSync) MergeDirtyToBytes() []byte {
	var buf []byte
	if x.isTimestampDirty() {
		buf = protowire.AppendTag(buf, 1, 0)
		buf = syncdep.AppendFieldValue(buf, x.timestamp)
	}
	if x.isClassicDirty() {
		buf = protowire.AppendTag(buf, 2, 2)
		buf = syncdep.AppendFieldValue(buf, x.classic)
	}
	return buf
}
func (x *PayRecordSync) MergeDirtyToPb(r *PayRecord) {
	if x.isTimestampDirty() {
		r.SetTimestamp(x.timestamp)
	}
	if x.isClassicDirty() {
		r.SetClassic(x.classic)
	}
}
func (x *PayRecordSync) SetDirty(index int, dirty bool, sync syncdep.Sync) {
	idx := index >> 3
	off := index & 7
	if dirty {
		x.dirtyFieldMark[idx] = x.dirtyFieldMark[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dirtyFieldMark[idx] = x.dirtyFieldMark[idx] & ^(1 << off)
	}
}
func (x *PayRecordSync) SetParentDirty() {
	if x.parent != nil {
		x.parent.SetDirty(x.indexInParent, true, x)
	}
}
func (x *PayRecordSync) SetParent(sync syncdep.Sync, idx int) {
	x.parent = sync
	x.indexInParent = idx
}
func (x *PayRecordSync) FlushDirty(dirty bool) {
	if dirty || x.isTimestampDirty() {
		x.setTimestampDirty(dirty, true)
	}
	if dirty || x.isClassicDirty() {
		x.setClassicDirty(dirty, true)
	}
}
func (x *PayRecordSync) setTimestampDirty(dirty bool, recur bool) {
	x.SetDirty(x.timestampINDEX, dirty, x)
}
func (x *PayRecordSync) isTimestampDirty() bool {
	idx := x.timestampINDEX >> 3
	off := x.timestampINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *PayRecordSync) setClassicDirty(dirty bool, recur bool) {
	x.SetDirty(x.classicINDEX, dirty, x)
}
func (x *PayRecordSync) isClassicDirty() bool {
	idx := x.classicINDEX >> 3
	off := x.classicINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *PayRecordSync) Key() interface{} {
	return x.timestamp
}
func (x *PayRecordSync) SetKey(v interface{}) {
	if x.parent != nil {
		if _, ok := x.parent.(*syncdep.MapSync[int64, *PayRecordSync]); ok {
			panic("PayRecordSync key can not set")
		}
	}
	x.timestamp = v.(int64)
}
func (x *PayRecordSync) GetTimestamp() int64 {
	return x.timestamp
}
func (x *PayRecordSync) SetTimestamp(v int64) *PayRecordSync {
	if x.timestamp == v {
		return x
	}
	x.timestamp = v
	x.setTimestampDirty(true, false)
	return x
}
func (x *PayRecordSync) GetClassic() string {
	return x.classic
}
func (x *PayRecordSync) SetClassic(v string) *PayRecordSync {
	if x.classic == v {
		return x
	}
	x.classic = v
	x.setClassicDirty(true, false)
	return x
}
func (xs *PayRecord) SetTimestamp(v int64) {
	xs.Timestamp = &v
}
func (xs *PayRecord) SetClassic(v string) {
	xs.Classic = &v
}
