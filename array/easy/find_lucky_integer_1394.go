package leetcode

import "sort"

// O(NLogN) time complexit and O(1) space complexity
// sort array and check each number's frequency by iterating on array
// number whos frequency is equal to that number itself is luck number
// last lucky number will be the answer
func findLucky(arr []int) int {
	res := -1
	cur := arr[0]
	freq := 0
	sort.Ints(arr)
	for _, x := range arr {
		if x != cur {
			if cur == freq {
				res = cur
			}
			cur = x
			freq = 1
		} else {
			freq++
		}
	}
	if cur == freq {
		res = cur
	}
	return res
}

// O(N)time complexity and O(1) space complexity
func findLucky1(arr []int) int {
	res := -1
	var counter [501]int
	for _, num := range arr {
		counter[num]++
	}
	for i := 1; i < 501; i++ {
		if i == counter[i] {
			res = i
		}
	}
	return res
}
