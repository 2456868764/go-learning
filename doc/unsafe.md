# unsafe 
## 对象内存布局
go 访问内存按照字节长数倍数访问，对象内存布局按照字节倍数对齐
* 32位机器按照4个字节直接对齐
* 64位机器按照8个直接对齐

看一下测试数据 BlogV1、BlogV2、BlogV3类型
```go

type BlogV1 struct {
	Name string
	Id int32
	Content []byte
	Title string
	Created time.Time
	Updated time.Time
}

type BlogV2 struct {
	Name string
	Id int32
	ParentId int32
	Content []byte
	Title string
	Created time.Time
	Updated time.Time
}

type BlogV3 struct {
	Name string
	Id int32
	Content []byte
	Title string
	Created time.Time
	Updated time.Time
	ParentId int32
}
```

```go
func OutputFieldLayout(object any) {
	typ := reflect.TypeOf(object)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fmt.Printf("%s offset %d\n", field.Name, field.Offset)
	}
}
```

BlogV1输出结果如下:
```shell
size: 112
Name offset 0
Id offset 16   // 这里Id是4个字节，这里后面空了4个字节
Content offset 24 // 这里按照8个字节对齐
Title offset 48
Created offset 64
Updated offset 88
```
BlogV2输出结果
```shell
size: 112 
Name offset 0
Id offset 16  // 这里Id是4个字节 
ParentId offset 20 // 这里也是4个字节，接着ID后面, 内存布局尽可能紧凑
Content offset 24  // 这里按照8个字节对齐
Title offset 48
Created offset 64
Updated offset 88
```

BlogV2输出结果:
```shell
size: 112 
Name offset 0
Id offset 16  // 这里Id是4个字节 
ParentId offset 20 // 这里也是4个字节，接着Id后面， 所以整体BlogV2比BlogV1字节数都是120个字节
Content offset 24  // 这里按照8个字节对齐
Title offset 48
Created offset 64
Updated offset 88
```
BlogV3输出结果:
```shell
size: 120
Name offset 0
Id offset 16 // 这里Id是4个字节 ,这里后面空了4个字节
Content offset 24 // 这里按照8个字节对齐
Title offset 48
Created offset 64
Updated offset 88
ParentId offset 112 // 这里按照8个字节对齐, 这里后面空了4个字节, 整体BlogV3比BlogV2多了8个字节
```
BlogV3 和 BlogV2 差异是ParentId 次序不同， 

对象内存布局按照8个字节对齐，内存布局不会改变次序，但是尽可能紧凑。

## unsafe 操作

unsafe操作是对象内存地址，本质上对象内存的起始地址和偏移量

ptr = 结构体内存起始地址 + 字段偏移量， 读写操作如下：

1. 读操作

*(*T)(ptr) T是目标对象类型，如果不清楚类型，可以用反射拿到 Type， 可以用 reflect.NewAt(typ, ptr).Elem()

2. 写操作

*(*T)(ptr) = T T是目标对象类型


## unsafe.Pointer 和 uintptr

unsafe.Pointer 表示一个对象起始地址，用 uintptr 表示偏移量

unsafe.Pointer 和 uintptr  两者都是指针，区别如下：
1. unsafe.Pointer是层面指针，对象 在GC标记复杂后， 会维护unsafe.Pointer值
2. uintptr 直接就是一个内存地址数字，GC后不会变化

## 用unsafe.Pointer 读写结构体

1. 定义访问接口 和 保存字段 类型和偏移量内容元数据

```go
type FieldAccessor interface {
	FieldAny(field string) (any, error)
	FieldInt(field string) (int, error)
	SetFieldAny(field string, val any) error
	SetFieldInt(field string, val int) error
}

type FieldMeta struct {
	name   string
	typ    reflect.Type  // 字段类型
	offset uintptr       // 字段偏移量 用 uintptr类型
}
```

2. 定义接口实现结构体

```go

var _ FieldAccessor = &UnSafeAccessor{}  // 确保编译情况 UnSafeAccessor 实现了 FieldAccessor 

type UnSafeAccessor struct {
	fieldMetas map[string]FieldMeta  // 字段元数据 map
	address    unsafe.Pointer        // 结构体起始地址 用unsafe.Pointer 类型
}

```

3. 接口实现

```go

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

```

在不清楚字段具体类型情况， 可以用 reflect 拿到对应字段 type, 用 reflect.NewAt 把对应指针地址数据 转成 Type 类型的 Value
```go
    ptr := unsafe.Pointer(uintptr(u.address) + fieldMeta.offset)
	fieldVal := reflect.NewAt(fieldMeta.typ, ptr)
```
在清楚知道字段具体类型情况下， 用具体类型强制转换就可以

```go

    ptr := unsafe.Pointer(uintptr(u.address) + fieldMeta.offset)
	fieldVal := *(*int)(ptr)

```

4. 构建 UnSafeAccessor

```go

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
		address: val.UnsafePointer(), // 这里用 Value 的 起始地址
	}, nil
}

```


## unsafe 操作性能 和 []byte 和 string 转换

1. 用unsafe 操作 []byte 和 string 转换， 不需要内存分配
```go

// StringToBytes converts string to byte slice without a memory allocation.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

```

2. 用类型转换，需要分配内存

```go
func rawBytesToStr(b []byte) string {
	return string(b)
}

func rawStrToBytes(s string) []byte {
	return []byte(s)
}
```

3. 性能对比

```shell
BenchmarkBytesConvBytesToStrRaw
BenchmarkBytesConvBytesToStrRaw-8   	65310120	        17.80 ns/op
BenchmarkBytesConvBytesToStr
BenchmarkBytesConvBytesToStr-8      	1000000000	         0.3186 ns/op
BenchmarkBytesConvStrToBytesRaw
BenchmarkBytesConvStrToBytesRaw-8   	61284420	        19.57 ns/op
BenchmarkBytesConvStrToBytes
BenchmarkBytesConvStrToBytes-8      	1000000000	         0.3174 ns/op
```





