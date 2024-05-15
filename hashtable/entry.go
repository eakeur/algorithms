package hashtable

type entry struct {
	key   string
	value any
	next  *entry
}

func (e *entry) set(k string, v any) *entry {
	if e == nil {
		return &entry{
			key:   k,
			value: v,
		}
	}

	if e.key == k {
		e.value = v
		return e
	}

	e.next = e.next.set(k, v)
	return e
}

func (e *entry) get(k string) any {
	if e == nil {
		return nil
	}

	if e.key == k {
		return e.value
	}

	return e.next.get(k)
}

func (e *entry) delete(k string) *entry {
	if e == nil {
		return e
	}

	if e.key == k {
		return e.next
	}

	e.next = e.next.delete(k)
	return e
}
