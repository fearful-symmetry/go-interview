package arrays

//QUESTION: Roate a NxN array by 90 degrees

// 1 2 3 4
// 5 6 7 8
// 9 a b c
// d e f 10

// d 9 5 1
// e a 6 2
// f b 7 3
// 10 c 8 4

//RotateNewMatrix solves the problem with a new array, that is, not in-place
//It will rotate the matrix 90º to the right
//This will assume an NxN (square) array
func RotateNewMatrix(m [][]int) [][]int {

	//get the matrix length
	size := len(m[0])

	//initialize a new 2d array, which is very annoying
	rotated := make([][]int, size)
	for i := range rotated {
		rotated[i] = make([]int, size)
	}

	for index := 0; index < size; index++ {
		//take the nth row of m, turn it into the len(m)-nth column of rotated
		col := size - (index + 1)
		for inner := 0; inner < size; inner++ {
			v := m[index][inner]
			rotated[inner][col] = v
		}
	}

	return rotated
}

//RotateMatrixInPlace rotates the new array 90º in-place
func RotateMatrixInPlace(m [][]int) [][]int {
	//rotate a given value from one edge to the other
	size := len(m[0])
	//iterate over the layers. Outer, inner
	for layer := 0; layer < size/2; layer++ {
		first := layer
		//iterate among elements in the layer
		last := size - 1 - layer
		for y := first; y < last; y++ {

			//make our algebra easier
			offset := y - first

			//save off the first edge
			top := m[first][y]
			//left to top
			m[first][y] = m[last-offset][first]
			//bottom to left
			m[last-offset][first] = m[last][last-offset]
			//right to bottom
			m[last][last-offset] = m[y][last]
			//top to right
			m[y][last] = top
		}
	}

	return m
}

// QUESTION: If an element in an MxN matrix is 0, zero the entire row and column

//ZeroRowColumn searches the matrix for a zero, then zeros out the row and column of the zero
func ZeroRowColumn(m [][]int) [][]int {

	//make one pass through the array, mark all rows and colums that we are zeroing
	var rows []int
	var columns []int

	for x, col := range m {
		for y := range col {
			if m[x][y] == 0 {
				rows = append(rows, x)
				columns = append(columns, y)
			}
		}
	}

	//Now that we have all the rows and colums, zero out each
	for _, col := range columns {
		for index := 0; index < len(m); index++ {
			m[index][col] = 0
		}
	}

	for _, row := range rows {
		for index := 0; index < len(m[0]); index++ {
			m[row][index] = 0
		}
	}

	return m
}
