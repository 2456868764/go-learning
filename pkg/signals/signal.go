package signals

import (
	"context"
	"os"
	"os/signal"
)

var onlyOneSignalHandler = make(chan struct{})

func SetupSignalHandler() context.Context {
	close(onlyOneSignalHandler) // 调用两次就会 panic

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		cancel() // 第一次 优雅 shutdown
		<-c
		os.Exit(1) // 第二次kill就直接退出
	}()

	return ctx
}
