package sync

import (
	"fmt"
	"sync"
)

type Queue[T any] struct {
	queue []T
	cond  *sync.Cond
}

func (q *Queue[T]) Enqueue(item T) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, item)
	fmt.Printf("Enqueue(): putting %+v to queue, notify all\n", item)
	q.cond.Broadcast()
}

func (q *Queue[T]) Dequeue() T {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.queue) == 0 {
		fmt.Println("no data available, wait")
		q.cond.Wait()
	}
	item := q.queue[0]
	fmt.Printf("Dequeue(): getting itme %+v from queue\n", item)
	q.queue = q.queue[1:]
	return item
}

func NewQueue[T any](initCap int) *Queue[T] {
	queue := &Queue[T]{
		queue: make([]T, 0, initCap),
		cond:  sync.NewCond(&sync.Mutex{}),
	}
	return queue
}
