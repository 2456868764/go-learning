package main

import "fmt"

func main() {

	// 因为m是结构体，所以方法调用的时候它数据是不会变的
	m := Model{
		Name: "John",
		Age: 20,
	}
	m.ChangeName("Tome!")
	m.ChangeAge(30)
	fmt.Printf("%v \n", m)

	// mp是指针，所以内部数据是可以被改变的
	mp := &Model{
		Name: "Ross",
		Age: 32,
	}

	// 因为ChangeName的接收器是结构体,所以mp的数据还是不会变
	mp.ChangeName("Emma Changed!")
	mp.ChangeAge(70)

	fmt.Printf("%v \n", mp)
}

type Model struct {
	Name string
	Age int
}

// 结构体接收器
func (m Model) ChangeName(newName string)  {
	m.Name = newName
}

// 指针接收器
func (m *Model) ChangeAge(newAge int) {
	m.Age = newAge
}
