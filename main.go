package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/imnzr/mie-gacoan-api/database"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// ENV Call
	env := godotenv.Load()
	if env != nil {
		log.Println("warning: .env file is not loaded")
	}
	// DB Connection
	db, err := database.DatabaseConnection()
	if err != nil {
		log.Println("failed connect to database: %w", err)
	}
	defer db.Close()

	// Httprouter
	router := httprouter.New()
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	// App Listen
	err = server.ListenAndServe()
	if err != nil {
		log.Println("failed to start server: %w", err)
	}
}
