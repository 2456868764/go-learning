package unsafe

import (
	"errors"
	"reflect"
	"unsafe"
)

type FieldAccessor interface {
	FieldAny(field string) (any, error)
	FieldInt(field string) (int, error)
	SetFieldAny(field string, val any) error
	SetFieldInt(field string, val int) error
}

type FieldMeta struct {
	name   string
	typ    reflect.Type
	offset uintptr
}

var _ FieldAccessor = &UnSafeAccessor{}

type UnSafeAccessor struct {
	fieldMetas map[string]FieldMeta
	address    unsafe.Pointer
}

func NewUnsafeAccessor(object any) (FieldAccessor, error) {
	typ := reflect.TypeOf(object)
	val := reflect.ValueOf(object)
	if typ.Kind() != reflect.Pointer {
		return nil, errors.New("unsupported kind")
	}

	typ = typ.Elem()
	numFields := typ.NumField()
	fieldMetas := make(map[string]FieldMeta, numFields)
	for i := 0; i < numFields; i++ {
		fieldMeta := FieldMeta{
			name: typ.Field(i).Name,
			typ: typ.Field(i).Type,
			offset: typ.Field(i).Offset,
		}
		fieldMetas[typ.Field(i).Name] = fieldMeta
	}
	return &UnSafeAccessor{
		fieldMetas: fieldMetas,
		address: val.UnsafePointer(),
	}, nil
}

func (u *UnSafeAccessor) FieldAny(field string) (any, error) {
	fieldMeta, ok:= u.fieldMetas[field]
	if !ok {
		return nil, errors.New("field not existed")
	}
	ptr := unsafe.Pointer(uintptr(u.address) + fieldMeta.offset)
	fieldVal := reflect.NewAt(fieldMeta.typ, ptr)
	return fieldVal.Interface(), nil

}

func (u *UnSafeAccessor) FieldInt(field string) (int, error) {
	fieldMeta, ok:= u.fieldMetas[field]
	if !ok {
		return 0, errors.New("field not existed")
	}
	ptr := unsafe.Pointer(uintptr(u.address) + fieldMeta.offset)
	fieldVal := *(*int)(ptr)
	return fieldVal, nil
}

func (u *UnSafeAccessor) SetFieldAny(field string, val any) error {
	fieldMeta, ok:= u.fieldMetas[field]
	if !ok {
		return errors.New("field not existed")
	}
	ptr := unsafe.Pointer(uintptr(u.address) + fieldMeta.offset)
	fieldVal := reflect.NewAt(fieldMeta.typ, ptr)
	if fieldVal.CanSet() {
		fieldVal.Set(reflect.ValueOf(val))
	}
	return nil
}

func (u *UnSafeAccessor) SetFieldInt(field string, val int) error {
	fieldMeta, ok:= u.fieldMetas[field]
	if !ok {
		return errors.New("field not existed")
	}
	ptr := unsafe.Pointer(uintptr(u.address) + fieldMeta.offset)
	*(*int)(ptr) = val
	return nil
}
