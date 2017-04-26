package api

import (
	"context"
)

// key definitions
const (
	CtxUserService = "user_service"
)

func WithUserService(ctx context.Context) UserService {
	return ctx.Value(CtxUserService).(UserService)
}
