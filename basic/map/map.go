package main

import "fmt"

func main() {
	// 创建一个容量是4的map
	salary := make(map[string]int, 2)
	// 没有指定容量
	d1 := make(map[string]string)
	// 直接初始化
	d2 := map[string]string{
		"Tom": "Jerry",
	}

	// 赋值
	salary["steve"] = 1200
	salary["jun"] = 1800


	d1["hello"] = "world"
	// 赋值
	d2["hello"] = "world"
	// 取值
	val := salary["steve"]
	println(val)



	// 使用两个返回值，ok表示map有没有这个key
	val, ok := salary["tom"]
	if !ok {
		println("tom not found")
	}

	for name, num := range salary {
		fmt.Printf("%s => %d \n", name, num)
	}

	//删除
	delete(salary, "jun")
}
