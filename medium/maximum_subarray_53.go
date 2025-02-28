package leetcode

// brute force approach - checking sum for subarray
// O(n^2) time complexity and O(1) space complexity
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		s := 0
		for j := i; j < len(nums); j++ {
			s += nums[j]
			if s > max {
				max = s
			}
		}
	}
	return max
}

// better approach
// O(n) time complexity and O(1) space complexity
// Logic : A subarray with a sum less than 0 will always reduce our answer and so this type of subarray cannot be a part of the subarray with maximum sum.
// While iterating we will add the elements to the sum variable and consider the maximum one.
// If at any point the sum becomes negative we will set the sum to 0 as we are not going to consider it as a part of our answer.
func maxSubArray1(nums []int) int {
	maxSum := nums[0]
	largest := nums[0]
	s := 0
	for i := 0; i < len(nums); i++ {
		if largest < nums[i] {
			largest = nums[i]
		}
		s += nums[i]
		if s < 0 {
			s = 0
			continue
		}
		if s > maxSum {
			maxSum = s
		}
	}
	if maxSum > largest {
		return maxSum
	}
	return largest
}
