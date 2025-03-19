package leetcode

import "sort"

// Better approach
// O(nlogn) time complexity and O(1) space complexity
// Logic - sort the array and then find the longest subsequence with consecutive numbers
func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	sort.Ints(nums)
	cnt := 1
	max := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			cnt++
			if max < cnt {
				max = cnt
			}
		} else if nums[i] == nums[i-1] {
			continue
		} else {
			cnt = 1
		}
	}
	if max < cnt {
		max = cnt
	}
	return max
}

// So called optimal approach but works slower in real world
// We will follow a similar approach to the brute-force method but with optimizations in the search process. Instead of checking sequences for every array element, we will only identify sequences starting from numbers that can serve as the beginning of a sequence. This targeted approach reduces unnecessary computations and improves efficiency.

// To achieve this, we will utilize the Map data structure.

// Identifying the Starting Number of a Sequence:
// Insert all array elements into a map to enable quick lookups.
// A number, num, can be a starting number if num - 1 is not present in the map.
// If x - 1 exists in the map, x cannot be a starting number, so we move to the next element.
// If x - 1 does not exist, x is the start of a sequence, and we begin finding its consecutive elements.
// This approach efficiently identifies and processes only valid sequences, leading to significant performance improvements.
func longestConsecutive1(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	mymap := make(map[int]int)
	cnt := 1
	max := 0
	for i := 0; i < len(nums); i++ {
		mymap[nums[i]] = 1
	}
	for k, _ := range mymap {
		if _, exists := mymap[k-1]; exists {
			continue
		} else {
			n := k + 1
			flag := true
			for flag {
				if _, exists := mymap[n]; exists {
					n++
					cnt++
				} else {
					flag = false
					if cnt > max {
						max = cnt
					}
					cnt = 1
				}
			}
		}
	}
	if max < cnt {
		max = cnt
	}
	return max
}
