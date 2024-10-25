package user

import (
	"context"
)

type UserRepositoryInterface interface {
	AddUsers(ctx context.Context, request *Users) (int, error)
	GetUsers(ctx context.Context, request *RequestGetUsers) ([]Users, int64, int, error)
}

type UserServiceinterface interface {
	ServiceAddUsers(ctx context.Context, request *RequestAddUsers) (any, int, error)
	ServiceGetUsers(ctx context.Context, request *RequestGetUsers) (any, int, error)
}
