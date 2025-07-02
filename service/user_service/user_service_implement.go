package userservice

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/imnzr/mie-gacoan-api/helper"
	"github.com/imnzr/mie-gacoan-api/models"
	userrepository "github.com/imnzr/mie-gacoan-api/repository/user_repository"
	userwebrequest "github.com/imnzr/mie-gacoan-api/web/request/user_web_request"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImplementation struct {
	UserRepository userrepository.UserRepositoryInterface
	DB             *sql.DB
}

// Create implements UserServiceInterface.
func (service *UserServiceImplementation) Create(ctx context.Context, request userwebrequest.UserCreateRequest) (models.User, error) {
	tx, err := service.DB.Begin()
	helper.TransactionErrorHandler(err)
	defer helper.CommitOrRollback(tx)

	// Hashed Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to hash password: %w", err)
	}
	// Input hashed password to models user
	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
	}
	// Create user with hashed password
	savedUser, err := service.UserRepository.Create(ctx, tx, user)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	// Return the saved user
	return models.User{
		Id:       savedUser.Id,
		Username: savedUser.Username,
		Email:    savedUser.Email,
		Password: savedUser.Password,
	}, nil
}

// Delete implements UserServiceInterface.
func (service *UserServiceImplementation) Delete(ctx context.Context, userId int) (models.User, error) {
	tx, err := service.DB.Begin()
	helper.TransactionErrorHandler(err)
	defer helper.CommitOrRollback(tx)

	// Find user by ID
	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to find user: %w", err)
	}

	// Delete user if exists
	service.UserRepository.Delete(ctx, tx, user)

	return models.User{}, nil
}

// FindByAll implements UserServiceInterface.
func (service *UserServiceImplementation) FindByAll(ctx context.Context) ([]models.User, error) {
	tx, err := service.DB.Begin()
	helper.TransactionErrorHandler(err)
	defer helper.CommitOrRollback(tx)

	users, err := service.UserRepository.FindByAll(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to find users: %w", err)
	}

	var result []models.User

	for _, users := range users {
		result = append(result, models.User{
			Id:       users.Id,
			Username: users.Username,
			Email:    users.Email,
		})
	}

	return result, nil
}

// FindById implements UserServiceInterface.
func (service *UserServiceImplementation) FindById(ctx context.Context, userId int) (models.User, error) {
	tx, err := service.DB.Begin()
	helper.TransactionErrorHandler(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to find user by ID: %w", err)
	}

	return models.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

// Login implements UserServiceInterface.
func (service *UserServiceImplementation) Login(ctx context.Context, request userwebrequest.UserLoginRequest) (models.User, error) {
	tx, err := service.DB.Begin()
	helper.TransactionErrorHandler(err)
	defer helper.CommitOrRollback(tx)

	// find user by email
	result, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to find user by email: %w", err)
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(request.Password))
	if err != nil {
		return models.User{}, fmt.Errorf("password does not match: %w", err)
	}

	return models.User{}, nil
}

// UpdateEmail implements UserServiceInterface.
func (service *UserServiceImplementation) UpdateEmail(ctx context.Context, request userwebrequest.UserUpdateEmail) (models.User, error) {
	tx, err := service.DB.Begin()
	helper.TransactionErrorHandler(err)
	defer helper.CommitOrRollback(tx)

	// Find user by ID
	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to find user: %w", err)
	}

	// Update user email
	result, err := service.UserRepository.UpdateEmail(ctx, tx, user)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to update user email: %w", err)
	}

	return models.User{
		Email: result.Email,
	}, nil
}

// UpdatePassword implements UserServiceInterface.
func (service *UserServiceImplementation) UpdatePassword(ctx context.Context, request userwebrequest.UserUpdatePassword) (models.User, error) {
	tx, err := service.DB.Begin()
	helper.TransactionErrorHandler(err)
	defer helper.CommitOrRollback(tx)

	// Find user by ID
	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to find user: %w", err)
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to hash password: %w", err)
	}

	// Update user password
	user.Password = string(hashedPassword)

	_, err = service.UserRepository.UpdatePassword(ctx, tx, user)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to update user password: %w", err)
	}

	// Get the updated user
	_, err = service.UserRepository.FindById(ctx, tx, user.Id)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to find updated user: %w", err)
	}

	return models.User{}, nil
}

func NewUserService(userRepository userrepository.UserRepositoryInterface, db *sql.DB) UserServiceInterface {
	return &UserServiceImplementation{
		UserRepository: userRepository,
		DB:             db,
	}
}
