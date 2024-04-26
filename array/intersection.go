package array

// The returning array should contain the numbers existing
// in both the arrays and repeated them for s many times as
// they appear in the origin arrayss
func IntersectRepeated(x, y []int) []int {

	set := map[int]int{}

	for _, valx := range x {

		for _, valy := range y {
			if valx == valy {
				set[valx] = set[valx] + 1
			}
		}
	}

	result := []int{}

	for num, occ := range set {
		for i := 0; i < occ; i++ {
			result = append(result, num)
		}
	}

	return SortIntArray(result, 0)
}
