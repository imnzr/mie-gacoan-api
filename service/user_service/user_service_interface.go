package userservice

import (
	"context"

	"github.com/imnzr/mie-gacoan-api/models"
	userwebrequest "github.com/imnzr/mie-gacoan-api/web/request/user_web_request"
)

type UserServiceInterface interface {
	Create(ctx context.Context, request userwebrequest.UserCreateRequest) (models.User, error)
	Delete(ctx context.Context, userId int) (models.User, error)
	FindById(ctx context.Context, userId int) (models.User, error)
	FindByAll(ctx context.Context) ([]models.User, error)

	UpdateEmail(ctx context.Context, request userwebrequest.UserUpdateEmail) (models.User, error)
	UpdatePassword(ctx context.Context, request userwebrequest.UserUpdatePassword) (models.User, error)

	Login(ctx context.Context, request userwebrequest.UserLoginRequest) (models.User, error)
}
