package leetcode

// Brute force approach
// O(n*n) time and space complexity
func rotateImage(matrix [][]int) {
	res := make([][]int, 0)
	for j := 0; j < len(matrix[0]); j++ {
		aux := []int{}
		for i := len(matrix) - 1; i >= 0; i-- {
			aux = append(aux, matrix[i][j])
		}
		res = append(res, aux)
	}
	copy(matrix, res)
}

// Better approach
// O(n*n) time complexity and O(1) space complexity
// Transpose the matrix and then reverse all the rows
func rotateImage1(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[0]); j++ {
			matrix[i][j] = matrix[i][j] + matrix[j][i]
			matrix[j][i] = matrix[i][j] - matrix[j][i]
			matrix[i][j] = matrix[i][j] - matrix[j][i]
		}
	}
	for i := 0; i < len(matrix); i++ {
		a := 0
		b := len(matrix[i]) - 1
		for a < b {
			matrix[i][a] = matrix[i][a] + matrix[i][b]
			matrix[i][b] = matrix[i][a] - matrix[i][b]
			matrix[i][a] = matrix[i][a] - matrix[i][b]
			a++
			b--
		}
	}

}
