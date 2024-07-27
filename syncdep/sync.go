package syncdep

type Sync interface {
	SetParent(sync Sync, idx int)
	SetDirty(idx int, dirty bool, sync Sync)
	SetParentDirty()
	FlushDirty(dirty bool)
	Key() interface{}
	SetKey(interface{})
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

type ArrayType interface {
	~int32 | int64 | uint32 | uint64 | string | bool | float32 | float64
}

type MsgKey interface {
	~int32 | int64 | uint32 | uint64 | string | bool
}

type MapSync[K MsgKey, V Sync] struct {
	parent   Sync
	idxInPar int
	value    map[K]V
	dirtied  map[K]struct{}
	deleted  map[K]struct{}
}

func NewMapSync[K MsgKey, V Sync]() *MapSync[K, V] {
	return &MapSync[K, V]{value: make(map[K]V), dirtied: make(map[K]struct{}), deleted: make(map[K]struct{})}
}

func (ms *MapSync[K, V]) SetParent(sync Sync, idx int) {
	ms.parent = sync
	ms.idxInPar = idx
}
func (ms *MapSync[K, V]) SetParentDirty() {
	if ms.parent != nil {
		ms.parent.SetDirty(ms.idxInPar, true, ms)
	}
}
func (ms *MapSync[K, V]) SetDirty(idx int, dirty bool, sync Sync) {
	var k K = sync.Key().(K)
	if dirty {
		ms.dirtied[k] = struct{}{}
		ms.SetParentDirty()
	} else {
		delete(ms.dirtied, k)
	}

}
func (ms *MapSync[K, V]) FlushDirty(dirty bool) {
	if !dirty {
		ms.dirtied = map[K]struct{}{}
		ms.deleted = map[K]struct{}{}
	}
	for k := range ms.value {
		ms.value[k].FlushDirty(dirty)
	}
}

func (ms *MapSync[K, V]) Key() interface{} {
	return nil
}
func (ms *MapSync[K, V]) SetKey(i interface{}) {

}
func (ms *MapSync[K, V]) PutOne(s V) *MapSync[K, V] {
	return ms.Put(s.Key().(K), s)
}
func (ms *MapSync[K, V]) Put(k K, s V) *MapSync[K, V] {
	old, exist := ms.value[k]
	s.Key()
	if exist {
		old.SetParent(nil, -1)
		ms.deleted[k] = struct{}{}
	}
	s.SetKey(k)
	s.SetParent(ms, -1)
	ms.value[k] = s
	ms.dirtied[k] = struct{}{}
	ms.SetParentDirty()
	return ms
}

func (ms *MapSync[K, V]) PutAll(kv map[K]V) *MapSync[K, V] {
	for k, v := range kv {
		ms.Put(k, v)
	}
	return ms
}

func (ms *MapSync[K, V]) Len() int {
	return len(ms.value)
}

func (ms *MapSync[K, V]) Clear() *MapSync[K, V] {
	if ms.Len() <= 0 {
		return ms
	}
	for k, v := range ms.value {
		v.SetParent(nil, -1)
		ms.deleted[k] = struct{}{}
	}
	ms.value = map[K]V{}
	ms.dirtied = map[K]struct{}{}
	ms.SetParentDirty()
	return ms
}
func (ms *MapSync[K, V]) Get(k K) V {
	v := ms.value[k]
	return v
}

func (ms *MapSync[K, V]) Remove(k K) V {
	v, exist := ms.value[k]
	if !exist {
		return v
	}
	delete(ms.value, k)
	delete(ms.dirtied, k)
	if vPtr, ok := any(v).(interface {
		SetParent(sync Sync, idx int)
	}); ok && vPtr != nil {
		vPtr.SetParent(nil, -1)
	}
	ms.deleted[k] = struct{}{}
	ms.SetParentDirty()
	return v
}
func (ms *MapSync[K, V]) RemoveAll(k []K) {
	for _, kk := range k {
		ms.Remove(kk)
	}
}
func (ms *MapSync[K, V]) Each(f func(k K, v V) bool) {
	for k, v := range ms.value {
		if !f(k, v) {
			break
		}
	}
}

func (ms *MapSync[K, V]) Dirtied() map[K]struct{} {
	return ms.dirtied
}

func (ms *MapSync[K, V]) Deleted() map[K]struct{} {
	return ms.deleted
}

func (ms *MapSync[K, V]) ContainDirtied(kk K) bool {
	for k := range ms.dirtied {
		if k == kk {
			return true
		}
	}
	return false
}

func (ms *MapSync[K, V]) ContainDeleted(kk K) bool {
	for k := range ms.deleted {
		if k == kk {
			return true
		}
	}
	return false
}

type ArraySync[V ArrayType] struct {
	parent   Sync
	idxInPar int
	value    []V
}

func NewArraySync[V ArrayType]() *ArraySync[V] {
	return &ArraySync[V]{value: make([]V, 0)}
}

func (ms *ArraySync[V]) SetParent(sync Sync, idx int) {
	ms.parent = sync
	ms.idxInPar = idx
}
func (ms *ArraySync[V]) SetParentDirty() {
	if ms.parent != nil {
		ms.parent.SetDirty(ms.idxInPar, true, ms)
	}
}
func (ms *ArraySync[V]) SetDirty(idx int, dirty bool, sync Sync) {
	ms.SetParentDirty()
}
func (ms *ArraySync[V]) FlushDirty(dirty bool) {
}

func (ms *ArraySync[V]) Key() interface{} {
	return nil
}
func (ms *ArraySync[V]) SetKey(i interface{}) {

}
func (ms *ArraySync[V]) Add(t V) *ArraySync[V] {
	ms.value = append(ms.value, t)
	ms.SetParentDirty()
	return ms
}

func (ms *ArraySync[V]) AddAll(t []V) *ArraySync[V] {
	if len(t) <= 0 {
		return ms
	}
	ms.value = append(ms.value, t...)
	ms.SetParentDirty()
	return ms
}

func (ms *ArraySync[V]) Each(f func(i int, v V) bool) {
	for i := range ms.value {
		if !f(i, ms.value[i]) {
			break
		}
	}
}

func (ms *ArraySync[V]) Clear() {
	ms.value = make([]V, 0)
	ms.SetParentDirty()
}
func (ms *ArraySync[V]) Len() int {
	return len(ms.value)
}

func (ms *ArraySync[V]) Remove(v V) {
	var idx = -1
	for i := range ms.value {
		if ms.value[i] == v {
			idx = i
			break
		}
	}
	ms.RemoveByIdx(idx)
}

func (ms *ArraySync[V]) RemoveByIdx(idx int) {
	if idx != -1 {
		ms.value = append(ms.value[0:idx], ms.value[idx+1:]...)
		ms.SetParentDirty()
	}
}

func (ms *ArraySync[V]) ValueView() []V {
	vv := make([]V, ms.Len())
	copy(vv, ms.value)
	return vv
}
