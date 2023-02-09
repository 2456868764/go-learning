package reflect

import (
	"errors"
	"reflect"
)

func IterateList(object any) ([]any, error) {
	val := reflect.ValueOf(object)
	typ := reflect.ValueOf(object)
	kind := typ.Kind()
	if !(kind == reflect.Array || kind == reflect.String || kind == reflect.Slice) {
		return nil, errors.New("type is not supported")
	}
	len := val.Len()
	list := make([]any, 0, len)
	for i := 0; i < len; i++ {
		element := val.Index(i)
		list = append(list, element.Interface())
	}

	return list, nil

}

func IterateMap(object any) ([]any, []any, error) {
	val := reflect.ValueOf(object)
	typ := reflect.ValueOf(object)
	kind := typ.Kind()
	if kind != reflect.Map {
		return nil, nil, errors.New("type is not supported")
	}
	len := val.Len()
	keys := make([]any, 0, len)
	values := make([]any, 0, len)
	for _, key := range val.MapKeys() {
		keys = append(keys, key.Interface())
		values = append(values, val.MapIndex(key).Interface())
	}

	//另外一种写法
	//iterate := val.MapRange()
	//for iterate.Next() {
	//	keys = append(keys, iterate.Key().Interface())
	//	values = append(values, iterate.Value().Interface())
	//}

	return keys, values, nil

}
