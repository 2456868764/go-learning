package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	t1 := func() { fmt.Println("run task1") }
	t2 := func() { fmt.Println("run task2") }
	t3 := func() { fmt.Println("run task3") }
	t4 := func() { fmt.Println("run task4") }
	t5 := func() { fmt.Println("run task5") }
	t6 := func() { fmt.Println("run task6") }
	t7 := func() { fmt.Println("run task7") }
	t8 := func() { fmt.Println("run task8") }
	t9 := func() { fmt.Println("run task9") }
	t10 := func() { fmt.Println("run task10") }
	tasks := make([]func(), 0, 10)
	tasks = append(tasks, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10)
	maxFlight(5, tasks)

	<-ctx.Done()

}

func maxFlight(maxFlight int, tasks []func()) error {
	ch := make(chan int, maxFlight)
	for _, task := range tasks {
		go func(t func()) {
			ch <- 0
			fmt.Println("start run()")
			t()
			fmt.Println("end run()")
			<-ch
		}(task)
	}
	return nil
}
