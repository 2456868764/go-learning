package main

import (
	"fmt"
	"reflect"
)

type Model struct {
	Name string
	Age  int32
}

func main() {
	model := Model{
		Name: "jun",
		Age:  22,
	}

	fmt.Println("model name =" + model.Name)
	fmt.Printf("model name = %s", model.Name)
	modelName := fmt.Sprintf("model name = %s", model.Name)
	println(modelName)

	fmt.Printf("v => %v \n", model)
	fmt.Printf("+v => %+v \n", model)
	fmt.Printf("#v => %#v \n", model)
	fmt.Printf("T => %T \n", model)

	i1 := 12
	i2 := float32(12.0)
	i3 := float64(12.0)
	fmt.Printf("result=%t", reflect.DeepEqual(i1, i2))
	fmt.Printf("result=%t", reflect.DeepEqual(i2, i3))
}
