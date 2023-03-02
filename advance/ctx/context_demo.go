package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//contextWithTimeOut()
	//contextWithCancel()
	//contextWithDeadline()
	contextWithValue()
	exampleContextTimeout()
	exampleContextTimeoutParentSon()
}
func contextWithTimeOut() {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()
	start := time.Now().Unix()
	<- ctx.Done()
	end := time.Now().Unix()
	fmt.Printf("total time = %d\n" , end -start)
	err := ctx.Err()
	switch err {
	case context.Canceled:
		fmt.Println("context canceled")
	case context.DeadlineExceeded:
		fmt.Println("context timeout")
	default:
		fmt.Println("context done")
	}

}

func contextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(){
		fmt.Println("go routine started")
		<- ctx.Done()
		fmt.Println("go routine canceled")
	}()

	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)
	err := ctx.Err()
	switch err {
	case context.Canceled:
		fmt.Println("context canceled")
	case context.DeadlineExceeded:
		fmt.Println("context timeout")
	default:
		fmt.Println("context done")
	}

}

func contextWithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2 * time.Second))
	defer cancel()
	start := time.Now().Unix()
	<- ctx.Done()
	end := time.Now().Unix()
	fmt.Printf("total time = %d\n" , end -start)
	err := ctx.Err()
	switch err {
	case context.Canceled:
		fmt.Println("context canceled")
	case context.DeadlineExceeded:
		fmt.Println("context timeout")
	default:
		fmt.Println("context done")
	}
}

func contextWithValue() {
	parentKey := "parent_key"
	parentValue := "parent_value"
	parentCtx := context.WithValue(context.Background(), parentKey, parentValue)
	sonKey := "son_key"
	sonValue := "son_value"

	sonCtx := context.WithValue(parentCtx, sonKey, sonValue)

	if parentCtx.Value(sonKey) == nil {
		fmt.Println("parent can not get son context")
	}

	if sonCtx.Value(parentKey) != nil {
		fmt.Println("son can get parent context")
	}
}



func exampleContextTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2* time.Second)
	defer cancel()

	bsCh := make(chan struct{})
	go func() {
		processBusiness()
		bsCh <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("TimeoutExample context timeout")
	case <-bsCh:
		fmt.Println("TimeoutExample process business end")
	}
}

func processBusiness() {
	time.Sleep(1 * time.Second)
}


func exampleContextTimeoutParentSon() {
	bg := context.Background()
	timeoutCtx, cancel1 := context.WithTimeout(bg, time.Second)
	subCtx, cancel2 := context.WithTimeout(timeoutCtx, 4*time.Second)
	defer cancel2()
	defer cancel1()
	go func() {
		// 一秒钟之后就会过期
		<-subCtx.Done()
		fmt.Printf("subctx timout")
	}()

	time.Sleep(2 * time.Second)
}


