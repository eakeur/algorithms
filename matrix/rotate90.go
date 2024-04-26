package matrix

func rotate90(m [][]int) {
	last := len(m) - 1

	// Range over half of the array not to
	// undo the transpositions
	for i := 0; i < len(m)/2; i++ {

		for j := i; j < last-i; j++ {

			// m[i][j] = top-left
			// m[j][last-i] = top-right
			// m[last-j][last-i] = bottom-right
			// m[last-j][i] = bottom-left

			m[i][j], m[j][last-i], m[last-i][last-j], m[last-j][i] = m[last-j][i], m[i][j], m[j][last-i], m[last-i][last-j]
		}
	}
}
