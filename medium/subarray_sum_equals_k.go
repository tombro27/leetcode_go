package leetcode

// brute force approach
// checking sum for all sub sequences
func subarraySum(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		s := 0
		for j := i; j < len(nums); j++ {
			s += nums[j]
			if s == k {
				res++
			}
		}
	}
	return res
}

// Optimal approach
// time complexit O(n) and space complexity O(n)
// using concept of prefix sum
func subarraySum1(nums []int, k int) int {
	res := 0
	mymap := make(map[int]int, len(nums))
	s := 0
	mymap[s]++
	for i := 0; i < len(nums); i++ {
		s += nums[i]
		if val, exists := mymap[s-k]; exists {
			res += val
		}
		mymap[s]++
	}
	return res
}
