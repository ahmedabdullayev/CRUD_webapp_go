package main

import (
	"CRUD_webapp_go/database"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	database.DbPool()
}
