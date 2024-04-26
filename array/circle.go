package array

// getCycle iterates over a slice s times, reading back to the beginning if len(a) < s
// Given a=[1, 2, 3, 4], start=3 and s=2, the result will be [4, 1].
// As we start reading the array from the position 3, we go back to the
// beginning everytime s >= len(a)
func getCycle(a []int, start int, s int) []int {
	n := len(a)
	if n == 0 {
		return nil
	}

	result := make([]int, s)

	for i := 0; i < s; i++ {
		result[i] = a[(start+i)%n]
	}
	return result
}
