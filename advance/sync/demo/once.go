package main

import (
	"fmt"
	"sync"
)

type Biz struct {
	onceClose sync.Once
}

func (b *Biz) Close() error {
	b.onceClose.Do(func() {
		fmt.Printf("close")
	})
	return nil
}

func init() {
	// 这里也是只执行一次
}

var singletone *Singleton
var instanceOnce sync.Once

type Singleton struct {
}

func NewInstance() *Singleton {
	instanceOnce.Do(func() {
		singletone = &Singleton{}
	})
	return singletone
}
