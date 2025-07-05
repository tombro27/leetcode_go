package leetcode

// O(n) time complexity and O(1) space complexity
// simple approach iterate and count number of negatives and zero
// break as soon as positive integer occurs
func maximumCount(nums []int) int {
	neg := 0
	zero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			neg++
		} else if nums[i] == 0 {
			zero++
		} else {
			break
		}
	}
	if len(nums)-neg-zero > neg {
		return len(nums) - neg - zero
	}
	return neg
}
