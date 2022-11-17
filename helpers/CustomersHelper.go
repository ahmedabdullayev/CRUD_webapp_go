package helpers

import (
	"CRUD_webapp_go/model"
	"database/sql"
	"time"
)

func CustomersToList(rows *sql.Rows) ([]model.Customer, error) {
	customer := model.Customer{}
	var customers []model.Customer

	for rows.Next() {
		var id int
		var first_name, last_name, gender, email, address string
		var birth_date time.Time
		err := rows.Scan(&id, &first_name, &last_name, &birth_date, &gender, &email, &address)

		if err != nil {
			return customers, err
		}

		customer.Id = id
		customer.FirstName = first_name
		customer.LastName = last_name
		customer.BirthDate = birth_date
		customer.Gender = gender
		customer.Email = email
		customer.Address = address

		customers = append(customers, customer)
	}

	return customers, nil
}
