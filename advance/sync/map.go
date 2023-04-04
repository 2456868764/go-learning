package sync

import "sync"

type SafeMap[K comparable, V any] struct {
	m     map[K]V
	mutex sync.RWMutex
}

func (s *SafeMap[K, V]) LoadOrStore(key K, newValue V) (val V, loaded bool) {
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()
	if ok {
		return val, true
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	// 加锁 和 double check
	val, ok = s.m[key]
	if ok {
		return val, true
	}
	s.m[key] = newValue
	return newValue, false
}

func (s *SafeMap[K, V]) LockDoSomething() {
	s.mutex.Lock()
	// 检查
	// Do something
	s.mutex.Unlock()
}

func (s *SafeMap[K, V]) RWLockDoSomething() {
	s.mutex.RLock()
	// 第一次检查
	// Do something
	s.mutex.RUnlock()

	s.mutex.Lock()
	// 第二次检查 , Double Check
	// Do something
	defer s.mutex.Unlock()
}
