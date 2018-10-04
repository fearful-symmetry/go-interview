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

//Just a helper function to compare our two 2D arrays
func compareMatrix(m1, m2 [][]int) bool {
	sizeX := len(m1)
	sizeY := len(m1[0])

	for x := 0; x < sizeX; x++ {
		for y := 0; y < sizeY; y++ {
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

func TestZeroRowColumn(t *testing.T) {

	a := [][]int{
		{4, 1, 2, 3},
		{4, 0, 6, 7},
		{8, 9, 10, 11},
		{12, 13, 0, 15},
		{12, 12, 5, 16},
	}
	aOut := [][]int{
		{4, 0, 0, 3},
		{0, 0, 0, 0},
		{8, 0, 0, 11},
		{0, 0, 0, 0},
		{12, 0, 0, 16},
	}
	ar := ZeroRowColumn(a)
	if !compareMatrix(a, aOut) {
		t.Fatalf("wrong result: %v \n", ar)
	}
}
