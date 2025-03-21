package leetcode

// O(n) time complexity and O(1) space complexity
// in go we can't reverse in-place.
func reverseOnlyLetters(s string) string {
	res := []rune(s)
	start := 0
	end := len(res) - 1
	for start < end {
		r := res[start]
		if !(r >= 65 && r <= 90) && !(r >= 97 && r <= 122) {
			start++
			continue
		}
		r = res[end]
		if !(r >= 65 && r <= 90) && !(r >= 97 && r <= 122) {
			end--
			continue
		}
		x := res[start]
		res[start] = res[end]
		res[end] = x
		start++
		end--
	}
	return string(res)
}
