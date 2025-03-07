package leetcode

// better than brute force
// O(n*n) time complexit and O(n*n) space complexity
func findMissingAndRepeatedValues(grid [][]int) []int {
	res := []int{}
	mymap := make(map[int]int)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if _, exists := mymap[grid[i][j]]; exists {
				res = append(res, grid[i][j])
				mymap[grid[i][j]]++
			} else {
				mymap[grid[i][j]] = 1
			}
		}
	}
	for i := 1; i <= len(grid)*len(grid); i++ {
		if _, exists := mymap[i]; exists == false {
			res = append(res, i)
			break
		}
	}
	return res
}

// better solution
// no need to iterate map for finding missing element
// use formula sum of 1 to n = n*(n+1)/2
func findMissingAndRepeatedValues1(grid [][]int) []int {
	res := []int{}
	n := len(grid)
	nsqr := n * n
	gridsum := 0
	mymap := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			gridsum += grid[i][j]
			if _, exists := mymap[grid[i][j]]; exists {
				res = append(res, grid[i][j])
				mymap[grid[i][j]]++
			} else {
				mymap[grid[i][j]] = 1
			}
		}
	}

	nsum := (nsqr * (nsqr + 1)) / 2
	missing := nsum - gridsum + res[0]
	res = append(res, missing)
	return res
}

// better solution
// O(n*n) time complexity and O(n*n) space complexity
// using array instead of map
func findMissingAndRepeatedValues2(grid [][]int) []int {
	res := []int{}
	n := len(grid)
	nsqr := n * n
	m := make([]int, nsqr+1)
	gridsum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			gridsum += grid[i][j]
			if m[grid[i][j]] != 1 {
				m[grid[i][j]] = 1
			} else {
				res = append(res, grid[i][j])
				m[grid[i][j]]++
			}
		}
	}

	nsum := (nsqr * (nsqr + 1)) / 2
	missing := nsum - gridsum + res[0]
	res = append(res, missing)
	return res
}
