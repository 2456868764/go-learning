package reflect

import (
	"errors"
	"reflect"
)

type MethodInfo struct {
	Name      string
	ParamsIn  []reflect.Type
	ParamsOut []reflect.Type
}

func IterateMethods(object any) (map[string]*MethodInfo, error) {
	typ := reflect.TypeOf(object)
	if !(typ.Kind() == reflect.Struct || typ.Kind() == reflect.Pointer) {
		return nil, errors.New("type is not supported")
	}

	numMethods := typ.NumMethod()
	methodInfos := make(map[string]*MethodInfo, numMethods)
	for i := 0; i < numMethods; i++ {
		method := typ.Method(i)
		numIns := method.Type.NumIn()
		numOuts := method.Type.NumOut()
		paramsIn := make([]reflect.Type, 0, numIns)
		paramsOut := make([]reflect.Type, 0, numOuts)
		// 第一个参数是接收器
		for j := 0; j < numIns; j++ {
			paramsIn = append(paramsIn, method.Type.In(j))
		}

		for j := 0; j < numOuts; j++ {
			paramsOut = append(paramsOut, method.Type.Out(j))
		}

		methodInfos[method.Name] = &MethodInfo{
			Name:      method.Name,
			ParamsIn:  paramsIn,
			ParamsOut: paramsOut,
		}
	}

	return methodInfos, nil

}
