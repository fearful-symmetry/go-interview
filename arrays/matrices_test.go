package arrays

import (
	"testing"
)

var matrixRotated = [][]int{
	{12, 8, 4, 0},
	{13, 9, 5, 1},
	{14, 10, 6, 2},
	{15, 11, 7, 3},
}

func compareMatrix(m1, m2 [][]int) bool {

	size := len(m1[0])

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if m1[x][y] != m2[x][y] {
				return false
			}
		}
	}

	return true

}

func TestRotateNewMatrix(t *testing.T) {

	a := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
		{12, 13, 14, 15},
	}
	ar := RotateNewMatrix(a)
	if !compareMatrix(matrixRotated, ar) {
		t.Fatalf("wrong result: %v \n", ar)
	}
}

func TestRotateInPlace(t *testing.T) {

	a := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
		{12, 13, 14, 15},
	}

	ar := RotateMatrixInPlace(a)
	if !compareMatrix(matrixRotated, ar) {
		t.Fatalf("wrong result: %v \n", ar)
	}
}
