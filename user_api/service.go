package user_api

import (
	"context"

	"efishery_test/user_api/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, ID int64) (*entity.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*entity.User, error)
}

// UserUsecase is a contract for usecases related to users
type UserUsecase interface {
	Register(ctx context.Context, user *entity.User) (*entity.UserForRegisterResponse, error)
	GetUser(ctx context.Context, ID int64) (*entity.UserPublic, error)
}
