package leetcode

// O(n) time complexity and O(1) space complexity
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	cnt := 0
	for _, num := range nums {
		if num == 1 {
			cnt++
		} else {
			if max < cnt {
				max = cnt
			}
			cnt = 0
		}
	}
	if max < cnt {
		max = cnt
	}
	return max

}
