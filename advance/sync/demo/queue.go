package main

import (
	"fmt"
	"github.com/2456868764/go-learning/advance/sync"
	"time"
)

func main() {
	q := sync.NewQueue[string](4)
	go func() {
		for {
			item := "a"
			q.Enqueue(item)
			fmt.Printf("Enqueue(): putting %+v to queue, notify all\n", item)
			time.Sleep(time.Second * 2)
		}
	}()
	for {
		item := q.Dequeue()
		fmt.Printf("Dequeue(): getting itme %+v from queue\n", item)
		time.Sleep(time.Second)
	}
}
