package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	copywg := wg
	fmt.Println(wg, copywg)
}
