package leetcode

// brute force approach
// O(nlogn) time complexity and O(1) space complexity
func missingNumber(nums []int) int {
	aux := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		aux[nums[i]] = -1
	}
	for i, num := range aux {
		if num != -1 {
			return i
		}
	}
	return 0
}

// Better approach
// O(n) time and space complexity
// since original slice nums will contain number in range of [0,n] we can create an auxiliary array aux with n+1 capacity
// for all the number present in nums we will set element at that index in aux as -1
// after that iterate aux array and check for index where value stored is not -1 that is our answer
func missingNumber1(nums []int) int {
	aux := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		aux[nums[i]] = -1
	}
	for i, num := range aux {
		if num != -1 {
			return i
		}
	}
	return 0
}

// optimal approach
// O(n) time complexity and O(1) space complexity
// logic : sum of first n natural numbers is (n*(n+1))/2
func missingNumber2(nums []int) int {
	s := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		s += nums[i]
	}
	return (n*(n+1))/2 - s
}

// best approach
// O(n) time complexity and O(1) space complexity
// Two important properties of XOR are the following:
// XOR of two same numbers is always 0 i.e. a ^ a = 0. ←Property 1.
// XOR of a number with 0 will result in the number itself i.e. 0 ^ a = a.  ←Property 2
func missingNumber3(nums []int) int {
	xor1 := 0
	xor2 := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		xor1 ^= nums[i]
		xor2 ^= i + 1
	}
	return xor1 ^ xor2
}
