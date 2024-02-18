package strings

import "strings"

/*
*
https://leetcode.cn/problems/reverse-words-in-a-string/
*/
func reverseWords(s string) string {
	words := strings.Fields(s)
	val := ""
	for i := len(words) - 1; i >= 0; i-- {
		val = val + words[i] + " "
	}
	return strings.TrimSpace(val)
}
