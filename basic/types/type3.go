package main

import (
	"fmt"
	"unsafe"
)

func main() {
	f1()
	f2()
	f3()
	f4()
}

func f1() {
	var a int32 = 10
	var b int64 = int64(a)
	var c float32 = 12.3
	var d float64 = float64(c)
	fmt.Printf("a:%d, b:%d, c:%f, d:%f\n", a, b, c, d)
}

func f2() {
	var a int = 10
	var b *int = &a
	var c *int64 = (*int64)(unsafe.Pointer(b))
	fmt.Println(*c)
}

func f3()  {
	var a interface{} = 10
	// 进行类型的断言的变量必须是空接口
	if _, ok := a.(int); ok {
		fmt.Println(a)
	}
}

type interf interface {
	Run()
}

type Server struct {
	
}

func (s Server) Run() {
}

var _ interf = (*Server)(nil)

func f4() {
	var server interface{} =  Server{}
	if _ , ok:= server.(interf); ok {
		fmt.Println("run as server\n")
	}

}
