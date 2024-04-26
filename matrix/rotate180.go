package matrix

func rotate180(m [][]int) {

	lastEl := len(m) - 1

	// Iterating over half of the matrix rows
	// so we don't undo the transpositions we did
	// for the first half
	for i := 0; i < len(m)/2; i++ {
		for j := 0; j < len(m); j++ {
			// len(m) - 1 - j gets the last element
			// excluding the ones processed
			// For example, in a row containing 4 elements
			// if m[i][j] is the first element, m[i][len(m) - 1 - j]
			// is the last one. If m[i][j] is the second one,
			// m[i][len(m) - 1 - j] is the penultimate, and so on
			m[i][j], m[lastEl-i][lastEl-j] = m[lastEl-i][lastEl-j], m[i][j]
		}
	}

	if len(m)%2 == 1 {
		middleRow := len(m) / 2
		for j := 0; j < middleRow; j++ {
			m[middleRow][j], m[middleRow][lastEl-j] = m[middleRow][lastEl-j], m[middleRow][j]
		}
	}
}
