package controllers

import (
	"CRUD_webapp_go/contracts"
	"CRUD_webapp_go/customerHelpers"
	"CRUD_webapp_go/models"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

type CustomerController struct {
	service contracts.CustomerServiceInterface
}

func NewCustomerController(service contracts.CustomerServiceInterface) CustomerController {
	return CustomerController{
		service: service,
	}
}

func (controller *CustomerController) Customers(w http.ResponseWriter, request *http.Request) {
	urlParams := request.URL.Query()

	firstName := urlParams.Get("firstName")
	lastName := urlParams.Get("lastName")
	orderBy := urlParams.Get("orderBy")
	orderType := urlParams.Get("orderType")
	offset, _ := strconv.Atoi(urlParams.Get("offset"))

	searchParams := models.SearchParams{
		FirstName: firstName,
		LastName:  lastName,
		Offset:    offset,
		OrderBy:   orderBy,
		OrderType: orderType,
	}

	customers, err := controller.service.GetAllByParams(searchParams)

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
		customer, errors := customerHelpers.ValidatePostPutActions(r)

		if len(errors) > 0 {
			var tmpl = template.Must(template.ParseGlob("templates/*")) // our templates
			w.WriteHeader(400)
			err := tmpl.ExecuteTemplate(w, "CreateCustomer.gohtml", errors)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			return
		}
		_, err := controller.service.Create(&customer)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

	}
	http.Redirect(w, r, "/", 303)
}
func (controller *CustomerController) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		customer, errors := customerHelpers.ValidatePostPutActions(r)

		customerEdit := models.CustomerAction{
			Customer: customer,
			Errors:   errors,
		}

		if len(errors) > 0 {
			var tmpl = template.Must(template.ParseGlob("templates/*")) // our templates
			err := tmpl.ExecuteTemplate(w, "EditCustomer.gohtml", customerEdit)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			return
		}
		_, err := controller.service.Update(&customer)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

	}
	http.Redirect(w, r, "/", 303)
}

func (controller *CustomerController) EditView(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	customerId, err := strconv.Atoi(queryId)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	customer, err := controller.service.GetOneById(customerId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	customerEdit := models.CustomerAction{
		Customer: customer,
		Errors:   nil,
	}
	var tmpl = template.Must(template.ParseGlob("templates/*"))
	err = tmpl.ExecuteTemplate(w, "EditCustomer.gohtml", customerEdit)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (controller *CustomerController) Show(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	customerId, err := strconv.Atoi(queryId)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	customer, err := controller.service.GetOneById(customerId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var tmpl = template.Must(template.ParseGlob("templates/*"))
	err = tmpl.ExecuteTemplate(w, "ShowCustomer.gohtml", customer)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (controller *CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	customerId, err := strconv.Atoi(queryId)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	deleteCustomer := controller.service.Delete(customerId)

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
