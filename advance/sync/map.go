package sync

import "sync"

type SafeMap[K comparable, V any] struct {
	m     map[K]V
	mutex sync.RWMutex
}

// LoadOrStore 读写锁要 double check
// goroutine1 设置： key1 => 1
// goroutine2 设置： key1 => 2
func (s *SafeMap[K, V]) LoadOrStore(key K, newValue V) (val V, loaded bool) {
	// 先获取读锁， goroutine1 和  goroutine2 都获得读锁进入
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()
	if ok {
		return val, true
	}

	// 假如 goroutine1 获得写锁，
	// goroutine2 等待 goroutine1 释放锁
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// 加锁 和 double check
	// goroutine1 修改 key1 => 1 后， 释放锁
	// goroutine2 进入后， 如果不 double check, 会把 key1 => 2, 而不是 goroutine1 设置 key1 == 1 的值
	// 到写锁 要 double check一下
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
