package leetcode

import "strings"

// O(n) time complexity and O(n) space complexity
// create a string slice of words in the given string
// now create result string by iterating slice in reverse order.
func reverseWords(s string) string {
	strList := []string{}

	var aux strings.Builder
	for _, c := range s {
		if c != ' ' {
			aux.WriteRune(c)
		} else {
			if aux.String() != "" {
				strList = append(strList, aux.String())
			}
			aux.Reset()
		}
	}
	if aux.String() != "" {
		strList = append(strList, aux.String())
	}
	aux.Reset()
	for i := len(strList) - 1; i >= 0; i-- {
		aux.WriteString(strList[i])
		if i != 0 {
			aux.WriteString(" ")
		}
	}
	return aux.String()
}

// O(1) space complexity
// inplace concept
// reverse the string.
// initialize start = -1
// iterate string
// if current character is space and start = -1 then remove the character and continue
// if current character is space and start != -1 then reverse string from index start to i-1
// if current character is not space and start != -1 then set start = i
// after loop check if start != -1 then reverse string from index start to i-1
// also chcek if last character of string is space if so remove the last character.

// Since above solution not possible in go due to immutable strings we will use rune slice to implement above approach.
func reverseWords1(s string) string {
	aux := []rune(s)
	reverseRune(aux, 0, len(aux)-1)
	start := -1
	i := 0
	for i < len(aux) {
		if aux[i] == ' ' {
			if start == -1 {
				aux = append(aux[:i], aux[i+1:]...)
			} else {
				reverseRune(aux, start, i-1)
				start = -1

				i++
			}
		} else {
			if start == -1 {
				start = i
			}
			i++
		}
	}
	if start != -1 {
		reverseRune(aux, start, i-1)
	}
	if aux[len(aux)-1] == ' ' {
		aux = append(aux[:len(aux)-1], aux[len(aux):]...)
	}
	return string(aux)
}

func reverseRune(arr []rune, start, end int) {
	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
}

// inplace concept although not possible in go.
// reverse the string.
// initialize start = -1
// iterate string
// if current character is space and start = -1 then remove the character and continue
// if current character is space and start != -1 then reverse string from index start to i-1
// if current character is not space and start != -1 then set start = i
// after loop check if start != -1 then reverse string from index start to i-1
// also chcek if last character of string is space if so remove the last character.
