package strings

/**
easy: https://leetcode.cn/problems/jewels-and-stones/
*/

func numJewelsInStones(jewels string, stones string) int {
	jewelsMap := make(map[uint8]int, 0)
	for i := 0; i < len(jewels); i++ {
		jewelsMap[jewels[i]] = 1
	}
	sum := 0
	for j := 0; j < len(stones); j++ {
		if _, ok := jewelsMap[stones[j]]; ok {
			sum++
		}
	}
	return sum
}
