package leetcode

import "strings"

// O(n) time complexity and O(1) space complexity
// Important don't use += operation on string it will make program very slow. Use strings.Builder instead.
// logic : iterate the string
// 1. In case of ( add to result string only if counter is not 0 and increment counter
// 2. In case of ) add to result string only if counter is not 1 and decrement counter
func removeOuterParentheses(s string) string {
	var res strings.Builder
	cnt := 0
	for _, c := range s {
		if c == '(' {
			if cnt != 0 {
				res.WriteRune('(')
			}
			cnt++
		} else {
			if cnt != 1 {
				res.WriteRune(')')
			}
			cnt--
		}
	}
	return res.String()
}
