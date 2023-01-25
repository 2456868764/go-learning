package main

import "fmt"

func main() {
	//初始化数组，必须包含3个初始值
	a1 := [3]int{1, 2, 3}
	fmt.Printf("a: %v, len: %d, cap: %d", a1, len(a1), cap(a1))

	// 初始化3个元素的空数组，所有元素都是默认值0
	var a2 [3]int
	fmt.Printf("a: %v, len: %d, cap: %d", a2, len(a2), cap(a2))

	// 按下标索引
	fmt.Printf("a1[1]: %d", a1[0])

}
