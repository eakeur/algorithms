package matrix

func rotate270(m [][]int) {
	last := len(m) - 1

	// Range over half of the array not to
	// undo the transpositions
	for i := 0; i < len(m)/2; i++ {

		for j := i; j < last-i; j++ {
			// topLeft, topRight, bottomRight, bottomLeft = topRight, bottomRight, bottomLeft, topRight
			m[i][j], m[j][last-i], m[last-i][last-j], m[last-j][i] = m[j][last-i], m[last-i][last-j], m[last-j][i], m[i][j]
		}
	}
}
