package unsafe

import (
	"fmt"
	"reflect"
)

func OutputFieldLayout(object any) {
	typ := reflect.TypeOf(object)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fmt.Printf("%s offset %d\n", field.Name, field.Offset)
	}
}
