package reflect

import "reflect"

func indirectReflectType(reflectType reflect.Type) reflect.Type {
	if reflectType.Kind() == reflect.Pointer {
		return reflectType.Elem()
	}
	return reflectType
}

func indirectReflectValue(reflectValue reflect.Value) reflect.Value {
	if reflectValue.Kind() == reflect.Pointer {
		return reflectValue.Elem()
	}
	return reflectValue
}
