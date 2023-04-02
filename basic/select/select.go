package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	done := make(chan struct{})

	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("go routine1 wake up")
		ch1 <- struct{}{}

	}()

	go func() {

		time.Sleep(time.Second)
		fmt.Println("go routine2 wake up")
		ch2 <- struct{}{}

	}()

	go func() {
		for {

			select {
			case <-ch1:
				fmt.Println("select go routine1")
			case <-ch2:
				fmt.Println("select go routine2")
			case <-time.After(2 * time.Second):
				done <- struct{}{}
				fmt.Println("select timeout")
				return
			}
		}
	}()

	<-done
	fmt.Println("main done")

}
