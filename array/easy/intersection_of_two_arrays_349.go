package leetcode

import "sort"

// time complexity -> O(nlogn) + O(mlogm) + 	O(n+m) 							~		 O(nlogn) (or O(mlogm), which ever is greater)
//						^ 			^				^
// 						|			|				|
// 				sorting nums1  	sorting nums2   iterating for intersection

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	if len(nums1) < len(nums2) {
		return findIntersection(nums1, nums2)
	} else {
		return findIntersection(nums2, nums1)
	}
}
func findIntersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	i := 0
	j := 0
	last := -1
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == last {
			i++
		} else {
			if nums1[i] == nums2[j] {
				last = nums1[i]
				res = append(res, last)
				i++
				j++
			} else if nums1[i] < nums2[j] {
				i++
			} else {
				j++
			}
		}
	}
	return res
}
