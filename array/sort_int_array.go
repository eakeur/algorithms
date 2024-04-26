package array

func SortIntArray(arr []int, way int) []int {

	const Ascending = 0
	const Descending = 1

	for i := 0; i < len(arr); i++ {
		// Range over the same array to compare arr[i]
		// with all the values in arr
		for j := 0; j < len(arr); j++ {

			// If the algorith desired is ascending
			// Check if arr[i] is smaller than arr[j].
			// If it is, arr[j] should swap places with arr[i]
			// so arr[j] can come first
			if way == Ascending && arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}

			if way == Descending && arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}
