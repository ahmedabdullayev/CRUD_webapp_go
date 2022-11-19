package main

import (
	"CRUD_webapp_go/controllers"
	"CRUD_webapp_go/database"
	"CRUD_webapp_go/repositories"
	"CRUD_webapp_go/services"
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
	customerService := services.NewEmployeeService(customerRep)
	customerControllers := controllers.NewCustomerController(customerService)

	router := customerControllers.Router()
	http.ListenAndServe(":8084", router)

	return err
}
