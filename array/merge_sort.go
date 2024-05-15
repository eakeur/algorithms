package array

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left, right := mergeSort(arr[:mid]), mergeSort(arr[mid:])

	short, long := left, right
	result := make([]int, 0)

	if len(left) > len(right) {
		short, long = right, left
	}

	for len(short) > 0 && len(long) > 0 {
		if short[0] < long[0] {
			result = append(result, short[0])
			short = deleteHead(short)
		} else {
			result = append(result, long[0])
			long = deleteHead(long)
		}
	}

	return append(result, append(short, long...)...)

}

func deleteHead(a []int) []int {
	if len(a) <= 1 {
		return []int{}
	}

	return a[1:]
}
