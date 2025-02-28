package leetcode

// O(n) time complexity and O(1) space complexity
func singleNumber(nums []int) int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor
}
