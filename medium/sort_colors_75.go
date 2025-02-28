package leetcode

// O(n) time complexity and O(1) space complexity
func sortColors(nums []int) {
	r := 0
	w := 0
	for _, num := range nums {
		if num == 0 {
			r++
		} else if num == 1 {
			w++
		}
	}
	for i := 0; i < len(nums); i++ {
		if r > 0 {
			nums[i] = 0
			r--
		} else if w > 0 {
			nums[i] = 1
			w--
		} else {
			nums[i] = 2
		}
	}
}

// The steps will be the following:
// First, we will run a loop that will continue until mid <= high.
// There can be three different values of mid pointer i.e. arr[mid]
// If arr[mid] == 0, we will swap arr[low] and arr[mid] and will increment both low and mid. Now the subarray from index 0 to (low-1) only contains 0.
// If arr[mid] == 1, we will just increment the mid pointer and then the index (mid-1) will point to 1 as it should according to the rules.
// If arr[mid] == 2, we will swap arr[mid] and arr[high] and will decrement high. Now the subarray from index high+1 to (n-1) only contains 2.
// In this step, we will do nothing to the mid-pointer as even after swapping, the subarray from mid to high(after decrementing high) might be unsorted. So, we will check the value of mid again in the next iteration.
// Finally, our array should be sorted.
// time complexity O(n)
func sortColors1(nums []int) {
	low := 0
	mid := 0
	high := len(nums) - 1
	for mid <= high {
		if nums[mid] == 0 {
			t := nums[low]
			nums[low] = nums[mid]
			nums[mid] = t
			low++
			mid++
		} else if nums[mid] == 1 {
			mid++
		} else {
			t := nums[high]
			nums[high] = nums[mid]
			nums[mid] = t
			high--
		}
	}
}
