package sync

import "sync"

type List[T any] interface {
	Get(index int) T
	Set(index int, t T)
	DeleteAt(index int) T
	Append(t T)
	Len() int
	Cap() int
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

func (a *ArrayList[T]) Len() int {
	return len(a.vals)
}

func (a *ArrayList[T]) Cap() int {
	return cap(a.vals)
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

func (s *SafeWapperArrayList[T]) Set(index int, t T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.l.Set(index, t)
}

func (s *SafeWapperArrayList[T]) DeleteAt(index int) T {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.l.DeleteAt(index)
}

func (s *SafeWapperArrayList[T]) Append(t T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.l.Append(t)
}

func (s *SafeWapperArrayList[T]) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.l.Len()
}

func (s *SafeWapperArrayList[T]) Cap() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.l.Cap()
}

func NewSafeWapperArrayList[T any](initCap int) *SafeWapperArrayList[T] {
	return &SafeWapperArrayList[T]{
		l: NewArrayList[T](initCap),
	}
}

type ConcurentArrayList[T any] struct {
	mu   sync.RWMutex
	vals []T
}

func (c *ConcurentArrayList[T]) Get(index int) T {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.vals[index]
}

func (c *ConcurentArrayList[T]) Set(index int, t T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.vals[index] = t
}

func (c *ConcurentArrayList[T]) DeleteAt(index int) T {
	c.mu.Lock()
	defer c.mu.Unlock()
	val := c.vals[index]
	c.vals = append(c.vals[:index], c.vals[index+1:]...)
	return val
}

func (c *ConcurentArrayList[T]) Append(t T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.vals = append(c.vals, t)
}

func (c *ConcurentArrayList[T]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.vals)
}

func (c *ConcurentArrayList[T]) Cap() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return cap(c.vals)
}

func NewConcurrentArrayList[T any](initCap int) *ConcurentArrayList[T] {
	return &ConcurentArrayList[T]{
		vals: make([]T, 0, initCap),
	}
}
