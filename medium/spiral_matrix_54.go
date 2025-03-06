package leetcode

// O(n*m) time complexity and O(1) space complexity
// Maintain for variables up,down,left,right
// run a loop in which there are four for loops
// loop one iterates M[up][left]->M[up][right] then up=up+1 and break outer loop if up > down
// loop two iterates M[up][right]->M[down][right] then right=right-1 and break outer loop if right < left
// loop three iterates M[down][right]->M[down][left] then down=down-1 and break outer loop if down < up
// loop four iterates M[down][left]->M[up][left] then left=left+1 and break outer loop if left > right
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	up := 0
	down := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1
	for true {
		for j := left; j <= right; j++ {
			res = append(res, matrix[up][j])
		}
		up++
		if up > down {
			break
		}
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		for j := right; j >= left; j-- {
			res = append(res, matrix[down][j])
		}
		down--
		if up > down {
			break
		}
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if right < left {
			break
		}
	}
	return res
}
