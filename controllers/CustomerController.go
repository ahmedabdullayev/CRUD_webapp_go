package controllers

import (
	"CRUD_webapp_go/CustomerHelpers"
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
	orderBy := urlParams.Get("orderBy")
	orderType := urlParams.Get("orderType")
	offset, _ := strconv.Atoi(urlParams.Get("offset"))

	searchParams := model.SearchParams{
		FirstName: firstName,
		LastName:  lastName,
		Offset:    offset,
		OrderBy:   orderBy,
		OrderType: orderType,
	}

	customers, err := controller.repository.GetAll(searchParams)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var tmpl = template.Must(template.ParseGlob("templates/*")) // our templates
	err = tmpl.ExecuteTemplate(w, "Customers.gohtml", customers)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
func (controller *CustomerController) CreateView(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseGlob("templates/*"))
	err := tmpl.ExecuteTemplate(w, "CreateCustomer.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
func (controller *CustomerController) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		customer, errors := CustomerHelpers.ValidatePostPutActions(r)

		if len(errors) > 0 {
			var tmpl = template.Must(template.ParseGlob("templates/*")) // our templates
			tmpl.ExecuteTemplate(w, "CreateCustomer.gohtml", errors)
			return
		}
		create := controller.repository.Create(&customer)

		if create != nil {
			http.Error(w, create.Error(), 500)
			return
		}

	}
	http.Redirect(w, r, "/", 303)
}
func (controller *CustomerController) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		customer, errors := CustomerHelpers.ValidatePostPutActions(r)

		customerEdit := model.CustomerAction{
			Customer: customer,
			Errors:   errors,
		}

		if len(errors) > 0 {
			var tmpl = template.Must(template.ParseGlob("templates/*")) // our templates
			tmpl.ExecuteTemplate(w, "EditCustomer.gohtml", customerEdit)
			return
		}
		create := controller.repository.Update(&customer)

		if create != nil {
			http.Error(w, create.Error(), 500)
			return
		}

	}
	http.Redirect(w, r, "/", 303)
}

func (controller *CustomerController) EditView(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	customerId, err := strconv.Atoi(queryId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	customer, err := controller.repository.GetOne(customerId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	customerEdit := model.CustomerAction{
		Customer: customer,
		Errors:   nil,
	}
	var tmpl = template.Must(template.ParseGlob("templates/*"))
	tmpl.ExecuteTemplate(w, "EditCustomer.gohtml", customerEdit)
}

func (controller *CustomerController) Show(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	customerId, err := strconv.Atoi(queryId)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	customer, err := controller.repository.GetOne(customerId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var tmpl = template.Must(template.ParseGlob("templates/*"))
	tmpl.ExecuteTemplate(w, "ShowCustomer.gohtml", customer)
}

func (controller *CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	customerId, err := strconv.Atoi(queryId)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	deleteCustomer := controller.repository.Delete(customerId)

	if deleteCustomer != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", 301)
}

func (controller *CustomerController) Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.Customers).Methods(http.MethodGet)
	router.HandleFunc("/create-customer", controller.CreateView).Methods(http.MethodGet)
	router.HandleFunc("/edit-customer", controller.EditView).Methods(http.MethodGet)
	router.HandleFunc("/show-customer", controller.Show).Methods(http.MethodGet)
	router.HandleFunc("/delete-customer", controller.Delete).Methods(http.MethodGet)

	router.HandleFunc("/create-customer", controller.Post).Methods(http.MethodPost)
	router.HandleFunc("/edit-customer", controller.Update).Methods(http.MethodPost)

	return router
}
