package controllers

import (
	"CRUD_webapp_go/contracts"
	"CRUD_webapp_go/model"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

type CustomerController struct {
	repository contracts.CustomerRepositoryInterface
}

func NewCustomerController(repository contracts.CustomerRepositoryInterface) CustomerController {
	return CustomerController{
		repository: repository,
	}
}

func (controller *CustomerController) Customers(w http.ResponseWriter, request *http.Request) {
	urlParams := request.URL.Query()

	firstName := urlParams.Get("firstName")
	lastName := urlParams.Get("lastName")
	offset, _ := strconv.Atoi(urlParams.Get("offset"))

	searchParams := model.SearchParams{
		FirstName: firstName,
		LastName:  lastName,
		Offset:    offset,
	}

	customers, err := controller.repository.GetAll(searchParams)

	if err != nil {
		panic(err.Error())
	}
	var tmpl = template.Must(template.ParseGlob("templates/*")) // our templates
	tmpl.ExecuteTemplate(w, "Customers.gohtml", customers)
}

func (controller *CustomerController) Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.Customers).Methods(http.MethodGet)

	return router
}
