package leetcode

// O(n) time complexity solution
// traverse the array and check if an element is smaller than the previous element
// if no such element exists array is sorted return true
// else array might be rotated - to check if array is rotated we need to check if this particular element is smaller than the first element of array if not array is neither rotated nor sorted
// if this element is smaller than the first element check if remaining array is sorted and all elements are smaller thant the first element if not array is neither rotated nor sorted else it is rotated
func check(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if nums[i] > nums[0] {
				return false
			}
			for j := i + 1; j < len(nums); j++ {
				if nums[j] < nums[j-1] || nums[j] > nums[0] {
					return false
				}
			}
			break
		}
	}
	return true
}
