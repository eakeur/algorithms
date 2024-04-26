package array

func deleteIndex(a []int, i int) []int {
	if i >= len(a) {
		return a
	}

	// fmt.Println(a[:i])   // a[:i] is not inclusive. If a = [1, 2, 4, 5], a[:2] is [1, 2]
	// fmt.Println(a[i:])   // a[i: is inclusive. If a = [1, 2, 4, 5], a[2:] = [4, 5]
	// fmt.Println(a[i+1:]) // Even when len(a) = 1, calling a[i+1] will not panic

	// For deleting an element from an array in Go,
	// append all the elements to the right of a[i] into a slice
	// with the element preceding a[i]
	//
	// a = [1, 2, 3, 4, 5, 6, 7]
	// i = 3
	// a[i] = 4
	//
	// a[:i] = [1, 2, 3] (i[3] is not included in the slice)
	// a[i+1:] = (i+1 = 4) -> a[4:] = [5, 6, 7]
	return append(a[:i], a[i+1:]...)
}

func deleteIndexInPlace(a []int, i int) {
	if i >= len(a) {
		return
	}

	// For deleting an element from an array in Go in place,
	// copy all the elements to the right of a[i] into a[i] on
	//
	//
	// a = [1, 2, 3, 4, 5, 6, 7]
	// i = 3
	// a[i] = 4
	//
	// a[i:] = [4, 5, 6, 7] (preceding elements are untouched)
	// a[i+1:] = (i+1 = 4) -> a[4:] = [5, 6, 7]
	// copy([4, 5, 6, 7], [5, 6, 7]) -> [1, 2, 3, 5, 6, 7, 7] (the last element is duplicated to fill the length)
	copy(a[i:], a[i+1:])
	a[len(a)-1] = 0
}
