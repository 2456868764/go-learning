package sync

import "sync"

type List[T any] interface {
	Get(index int) T
	Set(index int, t T)
	DeleteAt(index int) T
	Append(t T)
}

type ArrayList[T any] struct {
	vals []T
}

func (a *ArrayList[T]) Get(index int) T {
	return a.vals[index]
}

func (a *ArrayList[T]) Set(index int, t T) {
	if index >= len(a.vals) || index < 0 {
		panic("index out of range")
	}
	a.vals[index] = t
}

func (a *ArrayList[T]) DeleteAt(index int) T {
	if index >= len(a.vals) || index < 0 {
		panic("index out of range")
	}
	res := a.vals[index]
	a.vals = append(a.vals[:index], a.vals[index+1:]...)
	return res
}

func (a *ArrayList[T]) Append(t T) {
	a.vals = append(a.vals, t)
}

func NewArrayList[T any](initCap int) *ArrayList[T] {
	return &ArrayList[T]{
		vals: make([]T, 0, initCap),
	}
}

type SafeWapperArrayList[T any] struct {
	l     List[T]
	mutex sync.Mutex
}

func (s *SafeWapperArrayList[T]) Get(index int) T {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.l.Get(index)
}

func (s SafeWapperArrayList[T]) Set(index int, t T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.l.Set(index, t)
}

func (s SafeWapperArrayList[T]) DeleteAt(index int) T {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.l.DeleteAt(index)
}

func (s SafeWapperArrayList[T]) Append(t T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.l.Append(t)
}

func NewSafeWapperArrayList[T any](initCap int) *SafeWapperArrayList[T] {
	return &SafeWapperArrayList[T]{
		l: NewArrayList[T](initCap),
	}
}
