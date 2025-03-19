package leetcode

// O(n) time complexity
func moveZeroes(nums []int) {
	i := 0
	j := 1
	for j < len(nums) && i < len(nums) {
		if nums[i] == 0 && nums[j] != 0 {
			nums[i] = nums[j]
			nums[j] = 0
			i++
			j++
		} else {
			if nums[i] != 0 {
				i++
			}
			if nums[j] == 0 {
				j++
			}
		}
		if i == j {
			j++
		}
	}

}
