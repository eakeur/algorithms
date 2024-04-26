package array

func SortIntArray(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		// Range over the same array to compare arr[i]
		// with all the values in arr
		for j := 0; j < len(arr); j++ {

			// If the algorith desired is ascending
			// Check if arr[i] is smaller than arr[j].
			// If it is, arr[j] should swap places with arr[i]
			// so arr[j] can come first

			// (arr = [5, 2, 4, 0], i = 0, arr[i] = 5) -> [5, 2, 4, 0]
			// after first iteration nothing changes because 5 > all numbers
			//
			// (arr = [5, 2, 4, 0], i = 1, arr[i] = 2) -> [2, 5, 4, 0]
			// arr[1] and arr[0] swaps because 2 < 5. After swapping,
			// arr[1] becomes 5 and 5 > all numbers to be evaluated next.
			//
			// (arr = [2, 5, 4, 0], i = 2, arr[i] = 4) -> [2, 4, 5, 0]
			// arr[2] and arr[1] swaps because 4 < 5. After swapping,
			// arr[2] becomes 5 and 5 > all numbers to be evaluated next.
			//
			// (arr = [2, 4, 5, 0], i = 3, arr[i] = 0) -> [0, 2, 4, 5]
			// arr[3] and arr[0] swaps because 0 < 2. After swapping,
			// arr[3] becomes 2.
			// arr[3] and arr[1] swaps because 2 < 4. After swapping,
			// arr[3] becomes 4.
			// arr[3] and arr[2] swaps because 4 < 5. After swapping,
			// arr[3] becomes 5.
			//
			// After all iterations, the final array is [0, 2, 4, 5]

			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}
