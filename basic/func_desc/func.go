package main

import "fmt"

func main() {
	s := f0("jun")
	println(s)

	age, name := f1("", "jun")
	println(age)
	println(name)

	age, name = f2(0, "jun")

	println(age)
	println(name)

	d, e, f := f3("a", "b", 0, 1, "jun")

	println(d)
	println(e)
	println(f)

	h := []string{"CI", "Da"}
	f4("hello", "", h...)

}

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
