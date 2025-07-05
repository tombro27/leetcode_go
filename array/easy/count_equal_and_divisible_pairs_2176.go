package leetcode

// brute force solution
// O(n^2) time complexity and O(1) space complexity
func countPairs(nums []int, k int) int {
	c := 0
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				c++
			}
		}
	}
	return c
}
