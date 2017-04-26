package api

import (
	"context"

	null "gopkg.in/guregu/null.v3"
)

type APIHandler func(ctx context.Context, req *Request) *Response

type (
	Request struct {
		Header map[string]string
		Params map[string][]string
	}

	Response struct {
		Status  int
		Header  map[string]string
		Payload map[string]interface{}
	}
)

// model definitions
type (
	User struct {
		ID       null.Int    `json:"id"`
		Username null.String `json:"username"`
		Password null.String `json:"password"`
	}
)

// service definitions
type (
	Server interface {
		Run() error
	}

	UserService interface {
		GetAll() []User
	}
)
