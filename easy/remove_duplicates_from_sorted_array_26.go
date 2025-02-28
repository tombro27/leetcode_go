package leetcode

// O(n) time complxity and space complexity
// cons - we traverse the array twice and allocate a new slice
func removeDuplicates(nums []int) int {
	res := []int{}
	for _, n := range nums {
		if len(res) == 0 {
			res = append(res, n)
		} else if res[len(res)-1] != n {
			res = append(res, n)
		}
	}
	for i, _ := range res {
		nums[i] = res[i]
	}

	return len(res)
}

// two pointer approach
// O(n) time complexity and O(1) space complexity
func removeDuplicates1(nums []int) int {
	i := 0
	j := 1
	for j < len(nums) {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
			j++
		} else {
			j++
		}
	}
	return i + 1
}
