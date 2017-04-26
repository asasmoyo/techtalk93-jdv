package http

import (
	"context"
	"encoding/json"
	"fmt"
	libHttp "net/http"

	api "github.com/asasmoyo/techtalk93-jdv/example-api"
	"github.com/asasmoyo/techtalk93-jdv/example-api/sql"
	"github.com/asasmoyo/techtalk93-jdv/example-api/user"
)

// Server represents server http instance
type Server struct {
	ListenIP   string
	ListenPort int
	router     *libHttp.ServeMux
}

func (s *Server) Run() error {
	if s.ListenIP == "" {
		s.ListenIP = "127.0.0.1"
	}
	if s.ListenPort == 0 {
		s.ListenPort = 9000
	}

	s.init()

	return libHttp.ListenAndServe(fmt.Sprintf("%s:%d", s.ListenIP, s.ListenPort), s.router)
}

func (s *Server) init() {
	var userV1 = new(user.APIV1)

	s.router = libHttp.NewServeMux()
	s.router.HandleFunc("/v1/users", wrapper(userV1.GetAll))
}

func wrapper(h api.APIHandler) libHttp.HandlerFunc {
	return func(w libHttp.ResponseWriter, r *libHttp.Request) {
		var req = new(api.Request)

		req.Header = make(map[string]string)
		for k, v := range r.Header {
			req.Header[k] = v[0]
		}

		req.Params = make(map[string][]string)
		r.ParseForm()
		for k, v := range r.Form {
			req.Params[k] = v
		}

		var ctx = makeContext(r.Context())

		var res = h(ctx, req)
		if res == nil {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		for k, v := range res.Header {
			w.Header().Add(k, v)
		}

		p, err := json.Marshal(res.Payload)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(`{"error": "internal server error"}`))
		} else {
			w.WriteHeader(res.Status)
			w.Write(p)
		}
	}
}

func makeContext(ctx context.Context) context.Context {
	var userService api.UserService = new(sql.UserService)
	return context.WithValue(ctx, api.CtxUserService, userService)
}
