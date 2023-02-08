package main

import (
	"fmt"
	"reflect"
)

func main() {
	f1()
	f2()
}

func f1() {
	fmt.Println("f1()================================")
	a := 1234
	ty1 := reflect.TypeOf(a)
	va1 := reflect.ValueOf(a)
	fmt.Printf("typ1 name:%s, typ1 kind:%s\n", ty1.Name(), ty1.Kind())
	fmt.Printf("typ1 value:%+v\n", va1)

	book := Book{Name: "isito", Price: 108}
	ty2 := reflect.TypeOf(book)
	val2 := reflect.ValueOf(book)

	fmt.Printf("typ2 name:%s, typ1 kind:%s\n", ty2.Name(), ty2.Kind())
	fmt.Printf("typ2 value:%+v\n", val2)

	bookp := &Book{Name: "envoy", Price: 100}
	ty3 := reflect.TypeOf(bookp)
	val3 := reflect.ValueOf(bookp)

	fmt.Printf("typ3 name:%s, typ1 kind:%s\n", ty3.Name(), ty3.Kind())
	fmt.Printf("typ3 value:%+v\n", val3)
}

func f2() {
	fmt.Println("f12()================================")
	a := 126.54
	val := reflect.ValueOf(a)
	b := val.Interface().(float64)
	fmt.Printf("a :%f b:%f\n", a, b)
	book := Book{
		Name:  "eBRF",
		Price: 129,
	}
	fmt.Printf("book :%#v \n", book)
	valb := reflect.ValueOf(book)
	rbook := valb.Interface().(Book)
	fmt.Printf("book :%#v \n", rbook)

}

func f3() {
	var a float64 = 2.4
	v := reflect.ValueOf(a)
	v.SetFloat(76.1) // Error: will panic.
}

func f4() {
	var a float64 = 2.4
	v := reflect.ValueOf(&a)
	p := v.Elem()
	p.SetFloat(76.1)
}



type Book struct {
	Name  string
	Price int32
}
