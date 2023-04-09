//go:build app

package service

var signals = []os.Signal{
	os.Interrupt, os.Kill, syscall.SIGKILL, syscall.SIGSTOP,
	syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTRAP,
	syscall.SIGABRT, syscall.SIGSYS, syscall.SIGTERM,
}
