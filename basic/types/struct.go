package main

import "fmt"

func main() {
	// fish1 是 *Fish
	fish1 := &Fish{Color: "Red", Price: 100}
	fish1.Swim()

	fish2 := Fish{Color: "Blue", Price: 120}
	fish2.Swim()

	// fish3 是 *Fish
	fish3 := new(Fish)
	fish3.Swim()


	// 当你声明这样的时候，GO分配好内存
	var fish4 Fish
	fish4.Swim()

	// fish5 就是一个指针了
	//var fish5 *Fish
	// 这边会直接panic 掉
	//fish5.Swim()


	// 赋值,初始化按字段名字赋值
	fish6 := Fish{
		Color: "Green",
		Price: 200,
	}
	fish6.Swim()


	// 后面再单独赋值
	fish7 := Fish{}
	fish7.Color = "Purple"
	fish7.Swim()

}

type Fish struct {
	Color string
	Price uint32
}

func (f *Fish) Swim() {
	fmt.Printf("fish swim color:%s price:%d\n", f.Color, f.Price)
}



