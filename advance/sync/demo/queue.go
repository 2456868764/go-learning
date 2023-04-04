package main

import (
	"github.com/2456868764/go-learning/advance/sync"
	"time"
)

func main() {
	q := sync.NewQueue[string](4)
	go func() {
		for {
			q.Enqueue("a")
			time.Sleep(time.Second * 2)
		}
	}()
	for {
		q.Dequeue()
		time.Sleep(time.Second)
	}
}
