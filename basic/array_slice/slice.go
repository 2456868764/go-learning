package main

import "fmt"

func main() {

	a1 := []int{1, 2, 3, 4, 5, 6} // 初始化了6个元素的切片
	fmt.Printf("a1: %v, len %d, cap: %d \n", a1, len(a1), cap(a1))

	a2 := make([]int, 3, 4) // 创建一个包含3个元素，容量为4的切片
	fmt.Printf("a2: %v, len %d, cap: %d \n", a2, len(a2), cap(a2))

	a2 = append(a2, 7) // 后边添加一个元素，没有超出容量限制，不会发生扩容
	fmt.Printf("a2: %v, len %d, cap: %d \n", a2, len(a2), cap(a2))

	a2 = append(a2, 8) // 后边添加了一个元素，触发扩容
	fmt.Printf("a2: %v, len %d, cap: %d \n", a2, len(a2), cap(a2))

	a3 := make([]int, 6) // 只传入1个参数，表示创建6个元素，容量也为4切片
	// 等价于 a3 := make([]int, 6, 6)
	fmt.Printf("a3: %v, len %d, cap: %d \n", a3, len(a3), cap(a3))

	// 按下标索引
	fmt.Printf("a3[2]: %d", a3[2])

	//子切片
	subSliceDemo()
	//共享底层数据
	shareSliceDemo()
	//迭代切片数据
	rangeSliceDemo()
	//切片扩容
	extendSliceDemo()
	//添加数据和切片参数
	appendSliceDemo()
}

func subSliceDemo() {
	fmt.Println("subSliceDemo()===========================")
	array := []int{10, 20, 30, 40, 50, 60}
	fmt.Printf("array: %v, len %d, cap: %d \n", array, len(array), cap(array))

	sliceA := array[2:5]
	fmt.Printf("sliceA: %v, len %d, cap: %d \n", sliceA, len(sliceA), cap(sliceA))

	sliceB := array[1:3]
	fmt.Printf("sliceB: %v, len %d, cap: %d \n", sliceB, len(sliceB), cap(sliceB))

	a3 := array[3:]
	fmt.Printf("a3: %v, len %d, cap: %d \n", a3, len(a3), cap(a3))

	a4 := array[:4]
	fmt.Printf("a4: %v, len %d, cap: %d \n", a4, len(a4), cap(a4))
}

func shareSliceDemo() {
	fmt.Println("shareSliceDemo()===========================")
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	a2 := a1[2:]
	fmt.Printf("a1: %v, pointer = %p, len %d, cap: %d \n", a1, &a1, len(a1), cap(a1))
	fmt.Printf("a2: %v, pointer = %p, len %d, cap: %d \n", a2, &a2, len(a2), cap(a2))
	//a1, a2共享底层数据
	a2[0] = 9
	fmt.Printf("a1: %v, pointer = %p, len %d, cap: %d \n", a1, &a1, len(a1), cap(a1))
	fmt.Printf("a2: %v, pointer = %p, len %d, cap: %d \n", a2, &a2, len(a2), cap(a2))

	//a2结构发生变化，a1,a2不共享底层数据
	a2 = append(a2, 19)
	fmt.Printf("a1: %v, pointer = %p, len %d, cap: %d \n", a1, &a1, len(a1), cap(a1))
	fmt.Printf("a2: %v, pointer = %p, len %d, cap: %d \n", a2, &a2, len(a2), cap(a2))

	a2[1] = 29
	fmt.Printf("a1: %v, pointer = %p, len %d, cap: %d \n", a1, &a1, len(a1), cap(a1))
	fmt.Printf("a2: %v, pointer = %p, len %d, cap: %d \n", a2, &a2, len(a2), cap(a2))
}

func rangeSliceDemo() {
	fmt.Println("rangeSliceDemo()===========================")
	a1 := []int{10, 20, 30, 40, 50}
	for _, value := range a1 {
		value *= 2
	}
	fmt.Printf("a1 %+v\n", a1)
	for index, _ := range a1 {
		a1[index] *= 2
	}
	fmt.Printf("a1 %+v\n", a1)
}

func extendSliceDemo() {
	fmt.Println("extendSliceDemo()===========================")
	slice := []int{10, 20, 30, 40}
	newSlice := append(slice, 50)
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))

}

func appendSliceDemo() {
	fmt.Println("appendSliceDemo()===========================")
	input := make([]int, 0)
	fmt.Println("Origianl:", input)
	fmt.Printf("Origianl address %p   %p;\n", &input, input)
	processSliceData(input)
	fmt.Println("Output:", input)
	fmt.Printf("Output address %p   %p;\n", &input, input)

}

func processSliceData(input []int) {

	for i := 0; i < 5; i++ {
		input = append(input, i)
		fmt.Printf("i = %d ,len = %d ,cap = %d ,Temp address is %p   %p\n", i, len(input), cap(input), &input, input)
	}

}
