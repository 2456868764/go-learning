package sync

import (
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
	q.cond.Broadcast()
}

func (q *Queue[T]) Dequeue() T {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.queue) == 0 {
		q.cond.Wait()
	}
	item := q.queue[len(q.queue)-1]
	q.queue = q.queue[:len(q.queue)-1]
	return item
}

func NewQueue[T any](initCap int) *Queue[T] {
	queue := &Queue[T]{
		queue: make([]T, 0, initCap),
		cond:  sync.NewCond(&sync.Mutex{}),
	}
	return queue
}
