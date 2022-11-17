package controllers

import (
	"CRUD_webapp_go/contracts"
	"CRUD_webapp_go/model"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
	"time"
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
func (controller *CustomerController) CreateView(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseGlob("templates/*"))
	tmpl.ExecuteTemplate(w, "CreateCustomer.gohtml", nil)
}
func (controller *CustomerController) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var errors []string
		var formatedBirthDate = time.Time{}

		validate := validator.New()
		layout := "2006-01-02T15:04"
		birthDate := r.FormValue("birthDate")

		if birthDate != "" {
			formatedBirthDate, _ = time.Parse(layout, birthDate)
			currentDate := time.Now()
			tooYoungBirthDate := formatedBirthDate.AddDate(18, 0, 0)
			tooOldBirthDate := formatedBirthDate.AddDate(60, 0, 0)
			isInRangeOf18to60Years := tooYoungBirthDate.Before(currentDate) && currentDate.Before(tooOldBirthDate)
			if !isInRangeOf18to60Years {
				errors = append(errors, "Customer should be in range of 18 to 60 years old")
			}
		} else {
			errors = append(errors, "Birth date cannot be empty")
		}
		customer := model.Customer{
			FirstName: r.FormValue("firstName"),
			LastName:  r.FormValue("lastName"),
			BirthDate: formatedBirthDate,
			Gender:    r.FormValue("gender"),
			Email:     r.FormValue("email"),
			Address:   r.FormValue("address"),
		}
		validations := validate.Struct(customer)

		if validations != nil {
			for _, err := range validations.(validator.ValidationErrors) {
				errors = append(errors, err.StructField())
			}
			fmt.Print(errors)
		}
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

func (controller *CustomerController) Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.Customers).Methods(http.MethodGet)
	router.HandleFunc("/create-customer", controller.CreateView).Methods(http.MethodGet)

	router.HandleFunc("/create-customer", controller.Post).Methods(http.MethodPost)

	return router
}
