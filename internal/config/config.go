package config

import (
	"database/sql"
	"fmt"
	"os"
)

type Config struct {
	HTTPAddr string
}

func Read() Config {
	var c Config
	addr := os.Getenv("HTTP_ADDR")
	if addr != "" {
		c.HTTPAddr = addr
	}
	return c
}

func InitDB() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	return db, nil
}
