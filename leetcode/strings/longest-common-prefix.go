package strings

/*
* easy: https://leetcode.cn/problems/longest-common-prefix/
 */
func longestCommonPrefix(strs []string) string {
	index := -1
	stop := false
	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				stop = true
				break
			}
		}
		if stop == true {
			break
		}
		index++
	}
	if index >= 0 {
		return strs[0][0 : index+1]
	}
	return ""
}
