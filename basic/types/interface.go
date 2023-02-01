package main

// 首字母小写，是一个包私有的接口
type kid interface {
	name() string
	age() int
}


// 首字母大写，是一个包外可访问的接口
type Student interface {
	// 这里可以有任意多个方法，不过一般建议是小接口, 即接口里面不会有很多方法
	// 方法声明不需要func 关键字
	Eat()
}