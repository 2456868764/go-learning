package strings

/**
https://leetcode.cn/problems/reverse-string-ii
*/

func reverseStr(s string, k int) string {
	t := []byte(s)
	for i := 0; i < len(t)-1; i += 2 * k {
		p := i + k - 1
		if p+k > len(t)-1 {
			if len(t)-i < k {
				p = len(t) - 1
			}
		}
		left := i
		for left < p {
			t[left], t[p] = t[p], t[left]
			left++
			p--
		}
	}

	return string(t)
}
