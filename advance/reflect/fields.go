package reflect

import (
	"errors"
	"reflect"
)

func IterateFields(object any) (map[string]any, error) {
	typ := indirectReflectType(reflect.TypeOf(object))
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("type is not supported")
	}

	val := indirectReflectValue(reflect.ValueOf(object))
	numFields := typ.NumField()
	fields := make(map[string]any, numFields)
	for i := 0; i < numFields; i++ {
		fieldVal := val.Field(i)
		fieldTyp := typ.Field(i)

		if fieldTyp.IsExported() {
			//公开字段
			fields[fieldTyp.Name] = fieldVal.Interface()
		} else {
			fields[fieldTyp.Name] = reflect.Zero(fieldTyp.Type).Interface()
		}
	}
	return fields, nil
}
