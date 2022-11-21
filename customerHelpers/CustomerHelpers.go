package customerHelpers

import (
	"CRUD_webapp_go/models"
	"database/sql"
	"errors"
	"github.com/go-playground/validator"
	"net/http"
	"strconv"
	"time"
)

func CustomersToList(rows *sql.Rows) ([]models.Customer, error) {
	customer := models.Customer{}
	var customers []models.Customer

	for rows.Next() {
		var id int
		var firstName, lastName, gender, email, address string
		var birthDate time.Time
		err := rows.Scan(&id, &firstName, &lastName, &birthDate, &gender, &email, &address)

		if err != nil {
			return customers, err
		}

		customer.Id = id
		customer.FirstName = firstName
		customer.LastName = lastName
		customer.BirthDate = birthDate
		customer.Gender = gender
		customer.Email = email
		customer.Address = address

		customers = append(customers, customer)
	}

	return customers, nil
}

func ValidateBirthDate(birthDate string) (time.Time, error) {
	var formedBirthDate = time.Time{}
	layout := "2006-01-02T15:04"

	if birthDate != "" {
		formedBirthDate, _ = time.Parse(layout, birthDate)
		currentDate := time.Now()
		tooYoungBirthDate := formedBirthDate.AddDate(18, 0, 0)
		tooOldBirthDate := formedBirthDate.AddDate(60, 0, 0)
		isInRangeOf18to60Years := tooYoungBirthDate.Before(currentDate) && currentDate.Before(tooOldBirthDate)

		if !isInRangeOf18to60Years {
			return time.Time{}, errors.New("Customer should be in range of 18 to 60 years old")
		}
	} else {
		return time.Time{}, errors.New("Birth date cannot be empty")
	}

	return formedBirthDate, nil
}

func ValidatePostPutActions(r *http.Request) (models.Customer, []string) {
	var errors []string
	validate := validator.New()
	birthDate := r.FormValue("birthDate")

	formedBirthDate, err := ValidateBirthDate(birthDate)
	if err != nil {
		errors = append(errors, err.Error())
	}
	//
	id := r.FormValue("id")
	idInt, _ := strconv.Atoi(id)
	customer := models.Customer{
		Id:        idInt,
		FirstName: r.FormValue("firstName"),
		LastName:  r.FormValue("lastName"),
		BirthDate: formedBirthDate,
		Gender:    r.FormValue("gender"),
		Email:     r.FormValue("email"),
		Address:   r.FormValue("address"),
	}
	validations := validate.Struct(customer)

	if validations != nil {
		for _, err := range validations.(validator.ValidationErrors) {
			errors = append(errors, err.StructField())
		}
	}

	return customer, errors
}
