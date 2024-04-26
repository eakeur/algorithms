package array

// When an array is sorted, instead of iterating over until we find the number,
// we can split the array in half and repeat the process until we find it,
// ignoring several iterations. THis may not affect small arrays but it does help
// for large arrays
func findInSortedArray(a []int, num int) int {

	low, high := 0, len(a)-1

	for low <= high {

		m := low + (high-low)/2

		if a[m] == num {
			return m
		}

		// If target greater, ignore left half
		if a[m] < num {
			low = m + 1
		} else {
			// If target is smaller, ignore right half
			high = m - 1
		}
	}

	return -1
}

func bruteForceSearch(a []int, num int) int {
	for i, val := range a {
		if val == num {
			return i
		}
	}

	return -1
}
