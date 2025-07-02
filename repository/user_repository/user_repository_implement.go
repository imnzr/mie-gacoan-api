package userrepository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/imnzr/mie-gacoan-api/models"
)

type UserRepositoryImplementation struct{}

// FindByEmail implements UserRepositoryInterface.
func (u *UserRepositoryImplementation) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (models.User, error) {
	query := "SELECT id, username, email, password FROM user WHERE email = ?"
	rows, err := tx.QueryContext(ctx, query, email)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	user := models.User{}

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return models.User{}, fmt.Errorf("failed to scan row: %w", err)
		}
		return user, nil
	} else {
		return models.User{}, fmt.Errorf("user with email %s not found", email)
	}
}

// UpdateUsername implements UserRepositoryInterface.
func (u *UserRepositoryImplementation) UpdateUsername(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error) {
	query := "UPDATE user SET username = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, user.Username, user.Id)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to execute query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to fetch rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return models.User{}, fmt.Errorf("no rows updated, user with ID %d may not exist", user.Id)
	}

	return user, nil
}

// Create implements UserRepositoryInterface.
func (u *UserRepositoryImplementation) Create(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error) {
	query := "INSERT INTO user(username, email, password) VALUES(?,?,?)"
	result, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to execute query: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get last insert id: %w", err)
	}

	user.Id = int(id)

	return user, nil
}

// Delete implements UserRepositoryInterface.
func (u *UserRepositoryImplementation) Delete(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error) {
	query := "DELETE FROM user WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, user.Id)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to execute query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to fetch rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return models.User{}, fmt.Errorf("no rows deleted, user with ID %d may not exist", user.Id)
	}

	return user, nil
}

// FindByAll implements UserRepositoryInterface.
func (u *UserRepositoryImplementation) FindByAll(ctx context.Context, tx *sql.Tx) ([]models.User, error) {
	query := "SELECT id, username, email, password FROM user"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}

// FindById implements UserRepositoryInterface.
func (u *UserRepositoryImplementation) FindById(ctx context.Context, tx *sql.Tx, userId int) (models.User, error) {
	query := "SELECT id, username, email, password FROM user WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, userId)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to execute query: %w", err)
	}
	rows.Close()

	user := models.User{}

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return models.User{}, fmt.Errorf("failed to scan row: %w", err)
		}
		return user, nil
	} else {
		return models.User{}, fmt.Errorf("user with id %d not found", userId)
	}
}

// Login implements UserRepositoryInterface.
func (u *UserRepositoryImplementation) Login(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error) {
	query := "SELECT id, username, email, password WHERE email = ? AND password = ?"
	rows, err := tx.QueryContext(ctx, query, user.Email, user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return models.User{}, fmt.Errorf("failed to scan row: %w", err)
		}
		return user, nil
	} else {
		return models.User{}, fmt.Errorf("invalid email or password")
	}
}

// UpdateEmail implements UserRepositoryInterface.
func (u *UserRepositoryImplementation) UpdateEmail(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error) {
	query := "UPDATE user SET email = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, user.Email, user.Id)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to execute query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to fetch rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return models.User{}, fmt.Errorf("no rows updated, user with ID %d may not exist", user.Id)
	}

	return user, nil
}

// UpdatePassword implements UserRepositoryInterface.
func (u *UserRepositoryImplementation) UpdatePassword(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error) {
	query := "UPDATE user SET password = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, user.Password, user.Id)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to execute query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to fetch rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return models.User{}, fmt.Errorf("no rows updated, user with ID %d may not exist", user.Id)
	}

	return user, nil
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepositoryImplementation{}
}
