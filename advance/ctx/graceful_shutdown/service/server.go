package service

import "context"

type Server interface {
	Run(ctx context.Context)
	Stop(ctx context.Context)
}
