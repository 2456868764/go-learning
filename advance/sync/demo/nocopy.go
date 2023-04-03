package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	copywg := wg
	fmt.Println(wg, copywg)

	url := URL{Name: "hello"}
	url2 := url
	fmt.Println(url2, url)
}

type noCopy struct {
}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type URL struct {
	noCopy noCopy
	Name   string
}
