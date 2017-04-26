package user

import (
	"context"
	"reflect"
	"testing"

	api "github.com/asasmoyo/techtalk93-jdv/example-api"
	"github.com/asasmoyo/techtalk93-jdv/example-api/sql"
	null "gopkg.in/guregu/null.v3"
)

func TestAPIV1_GetAll(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.Request
	}

	type testCase struct {
		name    string
		a       *APIV1
		args    args
		wantRes *api.Response
	}

	var ctx1 = context.Background()
	ctx1 = context.WithValue(ctx1, api.CtxUserService, new(sql.UserService))

	var expectedRes1 = &api.Response{
		Status: 200,
		Payload: map[string]interface{}{
			"users": []api.User{
				api.User{ID: null.IntFrom(1), Username: null.StringFrom("user1"), Password: null.StringFrom("pass")},
				api.User{ID: null.IntFrom(2), Username: null.StringFrom("user2"), Password: null.StringFrom("pass")},
				api.User{ID: null.IntFrom(3), Username: null.StringFrom("user3"), Password: null.StringFrom("pass")},
			},
		},
	}

	tests := []testCase{
		testCase{name: "first test case", a: new(APIV1), args: args{ctx: ctx1, req: new(api.Request)}, wantRes: expectedRes1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := tt.a.GetAll(tt.args.ctx, tt.args.req); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("APIV1.GetAll() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
