package main

import "fmt"

func main() {
	checkGrade(90)
	checkGrade(50)
	checkDistance(100, 200)
	checkDistance(100, 120)
}

func checkGrade(grade int) {
	if grade >= 80 {
		fmt.Printf("Grade A\n")
	} else if grade < 80 && grade >= 60 {
		fmt.Printf("Grade B\n")
	} else {
		fmt.Printf("Grade C\n")
	}
}

func checkDistance(begin int, end int) {

	//distance变量只在if，else语句作用域里可以使用
	if distance := end - begin; distance > 50 {
		fmt.Printf("distance: %d is far away\n", distance)
	} else {
		fmt.Printf("distance: %d is so closed\n", distance)
	}

	//这里不能访问distance变量
	//fmt.Printf("distance is： %d\n", distance)
}
