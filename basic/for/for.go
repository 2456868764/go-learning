package main

import "fmt"

func main() {
	forLoop()
	forI()
	forRange()
}

func forLoop() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	index := 0
	fmt.Println(" for loop start \n ")
	for {
		if index == 4 {
			// break跳出for循环
			break
		}
		fmt.Printf("%d = %d\n", index, arr[index])
		index++
	}

	fmt.Println(" for loop end \n ")
}

func forI() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("for i loop start \n ")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d = %d \n", i, arr[i])
	}
	fmt.Println("for i loop end \n ")
}

func forRange() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("for range loop start \n ")

	for index, value := range arr {
		fmt.Printf("%d = %d\n", index, value)
	}

	// 只需要value, 可以用 _ 代替index
	for _, value := range arr {
		fmt.Printf("value: %d \n", value)
	}

	// 只需要index也可以去掉写成for index := range arr
	for index := range arr {
		fmt.Printf("index: %d \n", index)
	}

	fmt.Println("for range loop end \n ")
}
