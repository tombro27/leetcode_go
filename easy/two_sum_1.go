package leetcode

// brute force approach
// O(n^2) time complexity
func twoSum(nums []int, target int) []int {
	var res []int
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return append(res, i, j)
			}
		}
	}
	return res
}

// Optimal approach using map
// O(n^2) time complexity (worst case)
// Θ(n) time complexity (average case)
func twoSum1(nums []int, target int) []int {
	res := []int{}
	mymap := make(map[int]int)
	for i, num := range nums {
		mymap[num] = i
	}
	for i, num := range nums {
		val, exists := mymap[target-num]
		if exists && i != val {
			res = append(res, val, i)
			return res
		}
	}
	return res
}
