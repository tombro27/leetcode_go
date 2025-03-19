package leetcode

// Brute force approach with O(n^3) time complexity
// *failed cases where elements contained -1*
// this approach can be successful in case we declared a new matrix i.e., O(mn) space complexity
func setZeroes(matrix [][]int) {
	for i, row := range matrix {
		for j, m := range row {
			if m == 0 {
				setMinusOne(matrix, i, j)
			}
		}
	}

	for i, row := range matrix {
		for j, m := range row {
			if m == -1 {
				matrix[i][j] = 0
			}
		}
	}
}

func setMinusOne(matrix [][]int, i, j int) {
	for x := 0; x < len(matrix); x++ {
		if matrix[x][j] != 0 {
			matrix[x][j] = -1
		}
	}
	for y := 0; y < len(matrix[0]); y++ {
		if matrix[i][y] != 0 {
			matrix[i][y] = -1
		}
	}
}

// better approach O(n^2) time complexity but O(m+n) space complexity
// Declare two slices row and col, denoting the index of row and column of matrix respectively for which elements need to be made zero
func setZeroes1(matrix [][]int) {
	var row = make([]int, len(matrix))
	var col = make([]int, len(matrix[0]))
	for i, tuple := range matrix {
		for j, element := range tuple {
			if element == 0 {
				row[i] = -1
				col[j] = -1
			}
		}
	}
	for i, m := range row {
		if m == -1 {
			setRowZero(matrix, i)
		}
	}
	for j, n := range col {
		if n == -1 {
			setColZero(matrix, j)
		}
	}

}

func setRowZero(matrix [][]int, x int) {
	for j := 0; j < len(matrix[x]); j++ {
		matrix[x][j] = 0
	}
}
func setColZero(matrix [][]int, y int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][y] = 0
	}
}

// Optimal approach
// Using first row and first column as the row and column array declared in previous approach and col0 var for first column zero
// set zeroes from bottom right to upper left (dont set zeroes for first row and column)
// now set all elements zero in first row if matrix[0][0] = 0
// set set all elements zero in first column if col0 = 0
func setZeroes2(matrix [][]int) {
	col0 := 1
	for i, tuple := range matrix {
		for j, element := range tuple {
			if element == 0 {
				matrix[i][0] = 0
				if j != 0 {
					matrix[0][j] = 0
				} else {
					col0 = 0
				}
			}
		}
	}
	for i := len(matrix) - 1; i > 0; i-- {
		for j := len(matrix[i]) - 1; j > 0; j-- {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if matrix[0][0] == 0 {
		for j, _ := range matrix[0] {
			matrix[0][j] = 0
		}
	}
	if col0 == 0 {
		for i, _ := range matrix {
			matrix[i][0] = 0
		}
	}

}
