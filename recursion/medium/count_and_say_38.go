package leetcode

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return rle(countAndSay(n - 1))
}
func rle(s string) string {
	var aux strings.Builder
	cur := s[0]
	cnt := 1
	i := 1
	for ; i < len(s); i++ {
		if s[i] == cur {
			cnt++
		} else {
			aux.WriteString(strconv.Itoa(cnt))
			aux.WriteByte(cur)
			cur = s[i]
			cnt = 1
		}
	}
	aux.WriteString(strconv.Itoa(cnt))
	aux.WriteByte(cur)
	return aux.String()
}
