package main

import (
	"CRUD_webapp_go/controllers"
	"CRUD_webapp_go/database"
	"CRUD_webapp_go/model"
	"CRUD_webapp_go/repositories"
	"CRUD_webapp_go/services"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

var customerOne = model.Customer{
	Id:        2,
	FirstName: "Liisa",
	LastName:  "Kertu",
	BirthDate: time.Date(1999, time.January, 19, 14, 26, 0, 0, time.FixedZone("", 0)),
	Gender:    "Female",
	Email:     "liisa@gmail.com",
	Address:   "Pärnu mnt.7",
}

var customerTwo = model.Customer{
	FirstName: "Peeter",
	LastName:  "Mere",
	BirthDate: time.Date(1988, time.January, 19, 14, 26, 0, 0, time.FixedZone("", 0)),
	Gender:    "Male",
	Email:     "liisa@gmail.com",
	Address:   "Pärnu mnt.7",
}

func TestServicePost(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	db, _ := database.DbPool()

	customerRepository := repositories.NewCustomerRepository(db)
	customerService := services.NewEmployeeService(customerRepository)

	customerId, err := customerService.Create(&customerTwo)
	if err != nil {
		return
	}

	assert.GreaterOrEqual(t, customerId, 1)
}

func TestServiceUpdate(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	db, _ := database.DbPool()

	customerRepository := repositories.NewCustomerRepository(db)
	customerService := services.NewEmployeeService(customerRepository)

	customerId, err := customerService.Update(&customerOne)
	if err != nil {
		return
	}

	assert.Equal(t, customerId, 2)
}

func TestServiceGetOneById(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	db, _ := database.DbPool()

	customerRepository := repositories.NewCustomerRepository(db)
	customerService := services.NewEmployeeService(customerRepository)

	customer, err := customerService.GetOneById(2)
	if err != nil {
		return
	}

	assert.Equal(t, customer, customerOne)
}

func TestUpdate(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	data := url.Values{}
	data.Add("id", "2")
	data.Add("firstName", "Liisa")
	data.Add("lastName", "Kertu")
	data.Add("birthDate", "1999-01-19T14:26")
	data.Add("gender", "Female")
	data.Add("email", "liisa@gmail.com")
	data.Add("address", "Pärnu mnt.7")

	db, _ := database.DbPool()

	customerRepository := repositories.NewCustomerRepository(db)
	customerService := services.NewEmployeeService(customerRepository)
	controller := controllers.NewCustomerController(customerService)

	req, _ := http.NewRequest("POST", "/edit-customer", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpRecorder := httptest.NewRecorder()

	router := controller.Router()
	router.HandleFunc("/edit-customer", controller.Post).Methods("POST")
	router.ServeHTTP(httpRecorder, req)

	assert.Equal(t, 303, httpRecorder.Code, "Redirect response is expected")
}

func TestGetOneCustomer(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	db, _ := database.DbPool()

	customerRepository := repositories.NewCustomerRepository(db)
	customerService := services.NewEmployeeService(customerRepository)
	controller := controllers.NewCustomerController(customerService)

	req, _ := http.NewRequest("GET", "/show-customer?id=2", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpRecorder := httptest.NewRecorder()

	router := controller.Router()
	router.HandleFunc("/show-customer?id=2", controller.Customers).Methods("GET")
	router.ServeHTTP(httpRecorder, req)

	assert.Equal(t, 200, httpRecorder.Code, "OK response is expected")
}

func TestGetAllCustomers(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	db, _ := database.DbPool()

	customerRepository := repositories.NewCustomerRepository(db)
	customerService := services.NewEmployeeService(customerRepository)
	controller := controllers.NewCustomerController(customerService)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpRecorder := httptest.NewRecorder()

	router := controller.Router()
	router.HandleFunc("/", controller.Customers).Methods("GET")
	router.ServeHTTP(httpRecorder, req)

	assert.Equal(t, 200, httpRecorder.Code, "OK response is expected")
}

func TestPost(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	//2000-01-19T14:26
	data := url.Values{}
	data.Add("firstName", "Karlo")
	data.Add("lastName", "Luberg")
	data.Add("birthDate", "2000-01-19T14:26")
	data.Add("gender", "Male")
	data.Add("email", "peeter@gmail.com")
	data.Add("address", "Pärnu mnt.5")

	db, _ := database.DbPool()

	customerRepository := repositories.NewCustomerRepository(db)
	customerService := services.NewEmployeeService(customerRepository)
	controller := controllers.NewCustomerController(customerService)

	req, _ := http.NewRequest("POST", "/create-customer", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpRecorder := httptest.NewRecorder()

	router := controller.Router()
	router.HandleFunc("/create-customer", controller.Post).Methods("POST")
	router.ServeHTTP(httpRecorder, req)

	assert.Equal(t, 303, httpRecorder.Code, "Redirect response is expected")
}

func TestPostFail(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	//2000-01-19T14:26
	data := url.Values{}
	data.Add("firstName", "")
	data.Add("lastName", "")
	data.Add("birthDate", "2000-01-19T14:26")
	data.Add("gender", "Male")
	data.Add("email", "peeter@gmail.com")
	data.Add("address", "Pärnu mnt.5")

	db, _ := database.DbPool()

	customerRepository := repositories.NewCustomerRepository(db)
	customerService := services.NewEmployeeService(customerRepository)
	controller := controllers.NewCustomerController(customerService)

	req, _ := http.NewRequest("POST", "/create-customer", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpRecorder := httptest.NewRecorder()

	router := controller.Router()
	router.HandleFunc("/create-customer", controller.Post).Methods("POST")
	router.ServeHTTP(httpRecorder, req)

	assert.Equal(t, 400, httpRecorder.Code, "Client error is expected")
}

func TestPostYoungBirthDateFail(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	//2000-01-19T14:26
	data := url.Values{}
	data.Add("firstName", "")
	data.Add("lastName", "")
	data.Add("birthDate", "2016-01-19T14:26")
	data.Add("gender", "Male")
	data.Add("email", "peeter@gmail.com")
	data.Add("address", "Pärnu mnt.5")

	db, _ := database.DbPool()

	customerRepository := repositories.NewCustomerRepository(db)
	customerService := services.NewEmployeeService(customerRepository)
	controller := controllers.NewCustomerController(customerService)

	req, _ := http.NewRequest("POST", "/create-customer", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpRecorder := httptest.NewRecorder()

	router := controller.Router()
	router.HandleFunc("/create-customer", controller.Post).Methods("POST")
	router.ServeHTTP(httpRecorder, req)

	assert.Equal(t, 400, httpRecorder.Code, "Client error is expected")
}
