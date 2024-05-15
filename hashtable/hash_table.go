package hashtable

import "hash/maphash"

type table struct {
	mem  []*entry
	seed maphash.Seed
}

func NewTable(size int) *table {
	if size == 0 {
		size = 11
	}

	return &table{
		mem:  make([]*entry, size),
		seed: maphash.MakeSeed(),
	}
}

func (t *table) hash(key string) int {
	h := maphash.String(t.seed, key)
	return int(h % uint64(len(t.mem)))
}

func (t *table) set(key string, value any) {
	h := t.hash(key)
	t.mem[h] = t.mem[h].set(key, value)
}

func (t *table) delete(key string) {
	h := t.hash(key)
	t.mem[h] = t.mem[h].delete(key)
}

func (t *table) get(key string) any {
	return t.mem[t.hash(key)].get(key)
}
