package matrix

import "fmt"

func printMatrix(m [][]int) {
	for i := 0; i < len(m); i++ {

		for j := 0; j < len(m); j++ {

			fmt.Printf("%d ", m[i][j])

		}

		fmt.Println("")

	}

	fmt.Println("")
}
