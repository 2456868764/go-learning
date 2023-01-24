# main 
基础语法 —— main 函数要点 
* 无参数、无返回值 
* main 方法必须要在 main 包里面 
* `go run main.go` 就可以执行 
* 如果文件不叫 `main.go`，则需要 `go build` 之后再 `go run`

# package

语法：package xxxx 
* 字母和下划线的组合 
* 可以和文件夹不同名字 
* 同一个文件夹下的声明一致

引入包语法：import [alia] xxx 

* 如果一个包引入了但是没有使用，会报错 
* 匿名引入：前面多一个下划线

# string

string
* 双引号引起来，则内部双引号需要使用\转义 
* `引号引起来，则内部`需要\转义

[string长度
* 字节长度：和编码无关，用 len(str)获取 
* 字符数量：和编码有关，用编码库来计算]()

strings包

* string的拼接直接使用 + 号就可以。注意的是，某些语言支持 string 和别的类型拼接，但是 golang 不可以 
* strings 主要方法（你所需要的全部都可以找 到）： 
  * 查找和替换 
  * 大小写转换 
  * 子字符串相关 
  * 相等

```go
func main() {
	 fmt.Println(`表示核心线程池的大小。
      当提交一个任务时`)
	 fmt.Println("表示核心线程池的大小。\"当提交一个任务时\"")
	 fmt.Println(len("表示核心线程池的大小。当提交一个任务时，如果当前核心线程池的线程个数没有达到 corePoolSize"))
	 fmt.Println(utf8.RuneCountInString("表示核心线程池的大小。当提交一个任务时，如果当前核心线程池的线程个数没有达到 corePoolSize"))
	 strings.Contains("表示核心线程池的大小", "大小")

}

```

# rune类型

* rune，直观理解，就是字符 
* rune 不是 byte! 
* rune 本质是 int32，一个 rune 四个字节 
* rune 在很多语言里面是没有的，与之对应的 是，golang 没有 char 类型。rune 不是数字， 也不是 char，也不是 byte！ 
* 实际中不太常用

```go
type rune = int32
```

# 基础类型 ——bool, int, uint, float

* bool: true, false 
* int: int8, int16, int32, int64
* unit: uint8, uint16, uint32, uint64
* float: float32, float64

# byte类型
* byte，字节，本质是 uint8
* 对应的操作包在bytes上

```go

type byte = uint8

```

# 变量

var，语法：var name type = value 
* 局部变量 
* 包变量 
* 块声明 
* 驼峰命名 
* 首字符是否大写控制了访问性：大写包外可 访问； 
* golang 支持类型推断

变量声明 := 
* 只能用于局部变量，即方法内部 
* golang 使用类型推断来推断类型。数字会被理 解为int或者float64。（所以要其它类型的数字，就得用 var 来声明）

变量声明易错点
* 变量声明了没有使用 
* 类型不匹配 
* 同作用域下，变量只能声明一次

常量声明 const

首字符是否大写控制了访问性：大写包 外可访问； 
* 驼峰命名 
* 支持类型推断 
* 无法修改值


```go

package main

// Global 首字母大写，全局可以访问
var Global = "全局变量"

// 首字母小写，只能在这个包里面使用
// 其子包也不能用
var local = "包变量"

var (
	First string = "abc"
	second int32 = 16
)

func main() {
	// int 是灰色的，是因为 golang 自己可以做类型推断，它觉得你可以省略
	var a int = 13
	println(a)
	// 这里我们省略了类型
	var b = 14
	println(b)

	// 这里 uint 不可省略，因为生路之后，因为不加 uint 类型，15会被解释为 int 类型
	var c uint = 15
	println(c)

	// 这一句无法通过编译，因为 golang 是强类型语言，并且不会帮你做任何的转换
	// println(a == c)

	// 只声明不赋值，d 是默认值 0，类型不可以省略
	var d int
	println(d)
}



```

# func

```go

//一个返回值
func f0(name string) string {
	return "welcome " + name
}

// 多个参数，多个返回值
func f1(a string, name string) (int, string) {
	return 0, name
}

// 返回值命名
func f2(a int, b string) (age int, name string) {
	age = a + 1
	name = "Welcome " + b
	return
}

// 多个参数类型相同，可以写在一起
func f3(a, b string, a1, a2 int, c string) (d, e int, f string) {
	d = a1
	e = a2
	f = a + b
	return
}

// 不定参数
func f4(a string, b string, names ...string) {
	for _, name := range names {
		fmt.Printf("name = %s \n", name)
	}

}

```


# fmt

### demo
```go

package main
import "fmt"


type Model struct {
	Name string
	Age int32
}

func main() {
	model :=  Model{
		Name: "jun",
		Age : 22,
	}

	fmt.Println("model name =" + model.Name)
	fmt.Printf("model name = %s", model.Name)
	modelName :=  fmt.Sprintf("model name = %s", model.Name)
	println(modelName)

	fmt.Printf("v => %v \n", model)
	fmt.Printf("+v => %+v \n", model)
	fmt.Printf("#v => %#v \n", model)
	fmt.Printf("T => %T \n", model)

}

输出:
model name =jun
model name = junmodel name = jun
v => {jun 22}
+v => {Name:jun Age:22}
#v => main.Model{Name:"jun", Age:22}
T => main.Model


```
### 主要格式化
General:

	%v	the value in a default format
		when printing structs, the plus flag (%+v) adds field names
	%#v	a Go-syntax representation of the value
	%T	a Go-syntax representation of the type of the value
	%%	a literal percent sign; consumes no value

Boolean:

	%t	the word true or false

Integer:

	%b	base 2
	%c	the character represented by the corresponding Unicode code point
	%d	base 10
	%o	base 8
	%O	base 8 with 0o prefix
	%q	a single-quoted character literal safely escaped with Go syntax.
	%x	base 16, with lower-case letters for a-f
	%X	base 16, with upper-case letters for A-F
	%U	Unicode format: U+1234; same as "U+%04X"

Floating-point and complex constituents:

	%b	decimalless scientific notation with exponent a power of two,
		in the manner of strconv.FormatFloat with the 'b' format,
		e.g. -123456p-78
	%e	scientific notation, e.g. -1.234456e+78
	%E	scientific notation, e.g. -1.234456E+78
	%f	decimal point but no exponent, e.g. 123.456
	%F	synonym for %f
	%g	%e for large exponents, %f otherwise. Precision is discussed below.
	%G	%E for large exponents, %F otherwise
	%x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
	%X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

String and slice of bytes (treated equivalently with these verbs):

	%s	the uninterpreted bytes of the string or slice
	%q	a double-quoted string safely escaped with Go syntax
	%x	base 16, lower-case, two characters per byte
	%X	base 16, upper-case, two characters per byte

Slice:

	%p	address of 0th element in base 16 notation, with leading 0x

Pointer:

	%p	base 16 notation, with leading 0x
	The %b, %d, %o, %x and %X verbs also work with pointers,
	formatting the value exactly as if it were an integer.

The default format for %v is:

	bool:                    %t
	int, int8 etc.:          %d
	uint, uint8 etc.:        %d, %#x if printed with %#v
	float32, complex64, etc: %g
	string:                  %s
	chan:                    %p
	pointer:                 %p



# array
数组语法 是：[cap]type
* 初始化要指定长度（或者叫做容量） 
* 直接初始化 
* arr[i]的形式访问元素 
* len和cap操作用于获取数组长度








