package main

import (
	"fmt"
	"reflect"
)

type Person interface {
	Name() string
	Age() int
}

type Martian struct {
	
}

func (m Martian) Name() string {
	return "name"
}

func (m Martian) Age() int {
	return 0
}

// 强制类型 Martian 必须实现接口 Person 的所有方法
var _ Person = (*Martian)(nil)
// 1. 声明一个 _ 变量 (不使用)
// 2. 把一个 nil 转换为 (*Martian)，然后再转换为 Person
// 3. 如果 Martian 没有实现 Person 的全部方法，则转换失败，编译器报错

func main() {
	fmt.Printf("implement validation\n")

	// 获取 Person 类型
	p := reflect.TypeOf((*Person)(nil)).Elem()

	// 获取 Martian 结构体指针类型
	martian := reflect.TypeOf(&Martian{})

	// 判断 Martian 结构体类型是否实现了 Person 接口
	fmt.Println(martian.Implements(p))


	// 变量必须声明为 interface 类型
	var m interface{}
	m = &Martian{}
	if v, ok := m.(Person); ok {
		fmt.Printf("name = %s, age = %d\n", v.Name(), v.Age())
	} else {
		fmt.Println("Martian does not implements Person")
	}
}