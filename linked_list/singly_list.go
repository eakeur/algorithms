package linkedlist

type node struct {
	value any
	next  *node
}

func (n *node) add(val any) *node {
	if n == nil {
		return &node{
			value: val,
		}
	}

	n.next = n.next.add(val)
	return n
}

func (n *node) delete(val any) *node {
	if n == nil {
		return nil
	}

	if n.value == val {
		return n.next
	}

	n.next = n.next.delete(val)
	return n
}

func (n *node) arr() (a []any) {
	if n == nil {
		return nil
	}

	a = append(append(a, n.value), n.next.arr()...)
	return a
}
