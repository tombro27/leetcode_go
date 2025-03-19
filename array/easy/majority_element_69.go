package leetcode

import "sort"

// Using map
// O(n^2) time complexit and O(n) space complexity
func majorityElement(nums []int) int {
	n := len(nums) / 2
	if n == 0 {
		return nums[0]
	}
	mymap := make(map[int]int)
	for _, num := range nums {
		_, exists := mymap[num]
		if exists {
			mymap[num]++
			if mymap[num] > n {
				return num
			}
		} else {
			mymap[num] = 1
		}
	}
	return 0
}

// better approach O(nlogn) time complexity
func majorityElement1(nums []int) int {
	sort.Ints(nums)
	cnt := 1
	first := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[first] {
			cnt++
			if cnt > (len(nums) / 2) {
				return nums[i]
			}
		} else {
			first = i
		}
	}
	return nums[0]
}

// Optimal approach
// O(n) time complexity and  O(1) space complexity
func majorityElement2(nums []int) int {
	cnt := 1
	element := nums[0]
	for i := 1; i < len(nums); i++ {
		if cnt == 0 {
			element = nums[i]
			cnt++
		} else if cnt > len(nums)/2 {
			break
		} else {
			if nums[i] == element {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return element
}
