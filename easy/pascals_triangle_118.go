package leetcode

// brute force approach
// use this approach when whole triangle(matrix) is required
func generate(numRows int) [][]int {
	matrix := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		matrix[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				matrix[i][j] = 1
				continue
			}
			matrix[i][j] = matrix[i-1][j-1] + matrix[i-1][j]
		}
	}
	return matrix
}

// for fetching single element use formula (r-1)C(c-1) (combination formula, here r=row number and c=column number)
