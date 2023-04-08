package main

import (
	"context"
	"fmt"
	"github.com/2456868764/go-learning/pkg/signals"
	"time"
)

func main() {
	ctx := signals.SetupSignalHandler()
	go func(ctx context.Context) {
		for {
			// handle business here
			select {
			case <-ctx.Done():
				fmt.Printf("shutdown\n")
				return
			}
		}
	}(ctx)

	go func(ctx context.Context) {
		for {
			// handle business here
			select {
			case <-ctx.Done():
				fmt.Printf("shutdown\n")
				return
			}
		}
	}(ctx)

	<-ctx.Done()

	time.Sleep(5 * time.Second)

}
