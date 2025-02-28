package leetcode

// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]

// wrong implementation -> rotation in wrong direction
// my output [4 5 6 7 1 2 3]
func rotateWrong(nums []int, k int) {
	k %= len(nums)
	aux := []int{}
	for i := 0; i+k < len(nums); i++ {
		aux = append(aux, nums[i+k])
	}
	for i := 0; i < k; i++ {
		aux = append(aux, nums[i])
	}
	nums = aux
}

// Brute force approach
// O(kn) time complexity and O(1) space complexity
func rotate(nums []int, k int) {
	k %= len(nums)
	// aux:=[]int{}
	for k != 0 {
		last := nums[len(nums)-1]
		for i := len(nums) - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = last
		k--
	}
}

// Better approach
// O(n) space and time complexity
// iterating slice twice
func rotate1(nums []int, k int) {
	k %= len(nums)
	aux := []int{}
	for i := len(nums) - k; i < len(nums); i++ {
		aux = append(aux, nums[i])
	}
	for i := 0; i < len(nums)-k; i++ {
		aux = append(aux, nums[i])
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = aux[i]
	}
}

// Best approach - inplace
// O(n) time complexity and O(1) space complexity
// Step 1: Reverse the first n-k elements of the array.
// Step 2: Reverse the last k elements of the array
// Step 3: Reverse the whole array.
func rotate2(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-k-1)
	reverse(nums, len(nums)-k, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		t := nums[i]
		nums[i] = nums[j]
		nums[j] = t
		i++
		j--
	}
}
