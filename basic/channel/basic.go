package main

import (
	"fmt"
)

// 这里只能读
func read(c <-chan int) {
	fmt.Println("read:", <-c)
}

// 这里只能写
func write(c chan<- int) {
	c <- 0
}

func main() {
	c := make(chan int)
	go read(c)
	write(c)
}
