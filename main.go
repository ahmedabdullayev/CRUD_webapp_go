package main

import (
	"CRUD_webapp_go/controllers"
	"CRUD_webapp_go/database"
	"CRUD_webapp_go/repositories"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	db, err := database.DbPool()
	if err != nil {
		return err
	}

	customerRep := repositories.NewCustomerRepository(db)
	controllers := controllers.NewCustomerController(customerRep)

	router := controllers.Router()
	http.ListenAndServe(":8084", router)

	return err
}
