package strings

/**
easy: https://leetcode.cn/problems/reverse-only-letters/
*/

func reverseOnlyLetters(s string) string {
	i, j, bytes := 0, len(s)-1, []byte(s)
	for i < j {
		for i < j && !isLetter(bytes[i]) {
			i++
		}
		for i < j && !isLetter(bytes[j]) {
			j--
		}
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}

func isLetter(c byte) bool {
	if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}
