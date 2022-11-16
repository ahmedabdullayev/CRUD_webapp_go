package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

func dsn() string {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}

func DbPool() (*sql.DB, error) {
	connect, err := sql.Open("postgres", dsn())

	if err != nil {
		return connect, err
	}

	err = connect.Ping()
	if err != nil {
		return connect, err
	}

	log.Println("Db connected!")
	return connect, nil
}
