package leetcode

// Optimal approach
// O(n) time complexit and O(1) aux space complexity
func rearrangeArray(nums []int) []int {
	res := make([]int, len(nums))
	i := 0
	j := 1
	for _, num := range nums {
		if num >= 0 {
			res[i] = num
			i += 2
		} else {
			res[j] = num
			j += 2
		}
	}
	return res
}
