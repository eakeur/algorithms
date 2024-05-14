package array

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot, i, j := arr[len(arr)-1], -1, 0

	for j < len(arr)-1 {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}

		j++
	}

	i++
	arr[i], arr[j] = arr[j], arr[i]

	quickSort(arr[:i])
	quickSort(arr[i+1:])

}
