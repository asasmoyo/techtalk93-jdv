package user

import (
	"context"

	api "github.com/asasmoyo/techtalk93-jdv/example-api"
)

// APIV1 api version 1 handler
type APIV1 struct {
}

// GetAll returns all users
func (a *APIV1) GetAll(ctx context.Context, req *api.Request) (res *api.Response) {
	var users = api.WithUserService(ctx).GetAll()

	res = new(api.Response)
	res.Status = 200
	res.Payload = map[string]interface{}{
		"users": users,
	}
	return
}
