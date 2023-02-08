package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 1234
	ty1 := reflect.TypeOf(a)
	va1 := reflect.ValueOf(a)
	fmt.Printf("typ1 name:%s, typ1 kind:%s\n", ty1.Name(), ty1.Kind())
	fmt.Printf("typ1 value:%+v\n", va1)

	book := Book{Name: "isito", Price: 108}
	ty2 := reflect.TypeOf(book)
	val2 :=reflect.ValueOf(book)

	fmt.Printf("typ2 name:%s, typ1 kind:%s\n",ty2.Name(), ty2.Kind())
	fmt.Printf("typ2 value:%+v\n", val2)

	bookp := &Book{Name: "envoy", Price: 100}
	ty3 := reflect.TypeOf(bookp)
	val3 := reflect.ValueOf(bookp)

	fmt.Printf("typ3 name:%s, typ1 kind:%s\n",ty3.Name(), ty3.Kind())
	fmt.Printf("typ3 value:%+v\n", val3)
}

type Book struct {
	Name string
	Price int32
}

