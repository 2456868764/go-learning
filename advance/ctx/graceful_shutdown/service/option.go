package service

import "time"

type Option func(*App)

func WithServer(s *Server) Option {
	return func(a *App) {
		a.servers = append(a.servers, s)
	}
}

func WithCallback(callback ShutdownCallback) Option {
	return func(a *App) {
		a.callbacks = append(a.callbacks, callback)
	}
}

func WithShutdownTimeout(timeout time.Duration) Option {
	return func(a *App) {
		a.shutdownTimeout = timeout
	}
}

func WithWaitTime(waitTime time.Duration) Option {
	return func(a *App) {
		a.waitTime = waitTime
	}
}

func WithCallbackTimeout(callbackTimeout time.Duration) Option {
	return func(a *App) {
		a.callbackTimeout = callbackTimeout
	}
}
