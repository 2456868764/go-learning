package main

import (
	"fmt"
	"sync"
)

func worker(i int) {
	fmt.Println("worker: ", i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		// 开始 goroutine 前 +1
		wg.Add(1)
		go func(i int) {
			// 完成后减 1
			defer wg.Done()
			worker(i)
		}(i)
	}
	// 等待
	wg.Wait()
	fmt.Println("main done")
}
