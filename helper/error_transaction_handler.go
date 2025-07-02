package helper

import (
	"fmt"
	"log"

	"github.com/imnzr/mie-gacoan-api/models"
)

func TransactionErrorHandler(err error) (models.User, error) {
	if err != nil {
		log.Printf("error starting transaction: %v", err)
		return models.User{}, fmt.Errorf("failed to start transaction: %w", err)
	}
	return models.User{}, nil
}
