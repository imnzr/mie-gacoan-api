package userrepository

import (
	"context"
	"database/sql"

	"github.com/imnzr/mie-gacoan-api/models"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error)
	Delete(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (models.User, error)
	FindByAll(ctx context.Context, tx *sql.Tx) ([]models.User, error)

	UpdateEmail(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error)
	UpdatePassword(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error)
	UpdateUsername(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error)

	Login(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error)
	// forgot-password
}
