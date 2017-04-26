package sql

import (
	api "github.com/asasmoyo/techtalk93-jdv/example-api"
	"gopkg.in/guregu/null.v3"
)

type UserService struct {
}

func (s *UserService) GetAll() []api.User {
	return []api.User{
		api.User{ID: null.IntFrom(1), Username: null.StringFrom("user1"), Password: null.StringFrom("pass")},
		api.User{ID: null.IntFrom(2), Username: null.StringFrom("user2"), Password: null.StringFrom("pass")},
		api.User{ID: null.IntFrom(3), Username: null.StringFrom("user3"), Password: null.StringFrom("pass")},
	}
}
