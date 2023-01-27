package main

import "fmt"

func main() {
	fGrade("A")
	fGrade("D")
	fDay(1)
	fDay(6)
	fScore(90)
	fScore(60)
	fFallThrough(100)
	fFallThrough(300)
}

func fGrade(grade string) {
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	case "C":
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

}

func fDay(day int) {
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")

	}
}

func fScore(score int) {
	switch {
	case score >= 90:
		fmt.Println("享受假期")
	case score < 90 && score >= 80:
		fmt.Println("努力学习")
	default:
		fmt.Println("拼命学习")
	}
}

func fFallThrough(a int){
	switch a {
	case 100:
		fmt.Println("100")
		fallthrough
	case 200:
		fmt.Println("200")
	case 300:
		fmt.Println("300")
	default:
		fmt.Println("other")
	}
}
