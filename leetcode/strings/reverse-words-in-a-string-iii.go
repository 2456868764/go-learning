package strings

import "strings"

/*
*
https://leetcode.cn/problems/reverse-words-in-a-string-iii/
*/
func reverseWords3(s string) string {
	words := strings.Fields(s)
	val := ""
	for i := 0; i < len(words); i++ {
		bytes := []byte(words[i])
		left := 0
		right := len(bytes) - 1
		for left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
		val = val + string(bytes) + " "
	}
	return strings.TrimSpace(val)
}
