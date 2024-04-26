package invert

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntDigitsIntoStringArrayWithInvertedNumbersWithItsIndexes(t *testing.T) {

	assert.Equal(t, "27 26 95 84 23 02 01 40", IntDigitsIntoStringArrayWithInvertedNumbersWithItsIndexes(40028922))
}

func IntDigitsIntoStringArrayWithInvertedNumbersWithItsIndexes(number int) (output string) {
	str := strconv.Itoa(number)

	for i, st := range str {
		output = fmt.Sprintf("%s%d %s", string(st), i, output)
	}

	return strings.Trim(output, " ")
}

func TestIntArrIntoStringArrayWithPowedNumbersWithItsIndexes(t *testing.T) {

	assert.Equal(t, "1-0 4-1 9-2 16-3 25-4 36-5 49-6 64-7 81-8 0-9 2500-10", IntArrIntoStringArrayWithPowedNumbersWithItsIndexes([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, -50}))
}

func IntArrIntoStringArrayWithPowedNumbersWithItsIndexes(numbers []int) (output string) {
	for i, st := range numbers {
		output = fmt.Sprintf("%s %d-%d", output, int(math.Pow(float64(st), 2)), i)
	}

	return strings.Trim(output, " ")
}

func TestEsseAqui(t *testing.T) {

	var input = [][]int{
		[]int{1, 2, 3, 4},
		[]int{2, 2, 2, 2},
		[]int{0, 0, 0, 13},
	}

	var output = map[int]int{
		0: 30,
		1: 16,
		2: 169,
	}

	assert.Equal(t, output, FazIsso(input))
}

func FazIsso(arrs [][]int) map[int]int {
	output := make(map[int]int, len(arrs))

	for index, arr := range arrs {

		for _, n := range arr {
			output[index] += n * n
		}
	}

	return output
}
