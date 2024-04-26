package binarytree

import "fmt"

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func (t *Tree) Add(val int) *Tree {
	if t == nil {

		return &Tree{
			Value: val,
		}
	}

	if val == t.Value {
		return t
	}

	if val < t.Value {
		t.Left = t.Left.Add(val)

		return t
	}

	t.Right = t.Right.Add(val)

	return t
}

func (t *Tree) Search(val int) int {

	index, _ := t.search(val, 0)
	return index
}

func (t *Tree) search(val int, index int) (int, bool) {
	if t == nil {
		return index, false
	}

	if val == t.Value {
		return index, true
	}

	index++

	searchAt := t.Left

	if val > t.Value {
		searchAt = t.Right
	}

	return searchAt.search(val, index)
}

func (t *Tree) String() {
	if t == nil {
		return
	}

	t.Left.String()
	fmt.Println(t.Value)
	t.Right.String()
}
