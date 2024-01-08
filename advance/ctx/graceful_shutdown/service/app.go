package service

import (
	"context"
	"time"
)

type ShutdownCallback func(ctx context.Context) error

type App struct {
	// 启动服务数组
	servers []*Server

	// 优雅退出 整个超时设置， 默认是30秒
	shutdownTimeout time.Duration

	// 优雅退出 等待处理请求时间设置，默认是5秒
	waitTime time.Duration

	// 优雅退出 等待回调超时设置，默认是5分钟
	callbackTimeout time.Duration

	// 退出回调函数数组
	callbacks []ShutdownCallback
}

func NewApp(opts ...Option) *App {
	app := &App{
		waitTime:        5 * time.Second,
		callbackTimeout: 5 * time.Second,
		shutdownTimeout: 30 * time.Second,
	}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

func (app *App) StartAndServe() {

}

func (app *App) Shutdown() {

}

func (app *App) Close() {

}
