package array

// getCycle iterates over a slice s times, reading back to the beginning if len(a) < s
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
