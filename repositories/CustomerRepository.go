package repositories

import (
	"CRUD_webapp_go/contracts"
	"CRUD_webapp_go/helpers"
	"CRUD_webapp_go/model"
	"database/sql"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const DefaultLimit = 5
const DefaultOffset = 5

type CustomerRepository struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) contracts.CustomerRepositoryInterface {
	return &CustomerRepository{DB: db}
}

func (rep CustomerRepository) GetOne(id int) (model.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (rep CustomerRepository) GetAll(searchParams model.SearchParams) (model.CustomersPage, error) {
	var customersPage model.CustomersPage
	var totalRows int
	rep.DB.QueryRow("SELECT COUNT(*) FROM customers WHERE "+
		"LOWER(first_name) LIKE '%' || $1 || '%' "+
		"AND LOWER(last_name) LIKE '%' || $2 || '%'",
		strings.ToLower(searchParams.FirstName), strings.ToLower(searchParams.LastName)).Scan(&totalRows)

	var queryString string
	if searchParams.Offset < DefaultOffset {
		queryString = "SELECT * FROM customers WHERE " +
			"LOWER(first_name) LIKE '%' || $1 || '%' " +
			"AND LOWER(last_name) LIKE '%' || $2 || '%' ORDER BY id DESC LIMIT $3"
	} else {
		queryString = "SELECT * FROM customers WHERE LOWER(first_name) LIKE '%' || $1 || '%' " +
			"AND LOWER(last_name) LIKE '%' || $2 || '%' ORDER BY id DESC LIMIT $3 OFFSET " +
			strconv.Itoa(searchParams.Offset)
	}

	customersQuery, err := rep.DB.Query(queryString,
		strings.ToLower(searchParams.FirstName),
		strings.ToLower(searchParams.LastName),
		DefaultLimit)
	if err != nil {
		return customersPage, err
	}
	customers, err := helpers.CustomersToList(customersQuery)

	totalPages := int(math.Ceil(float64(totalRows) / float64(DefaultLimit)))
	customersPage = model.CustomersPage{
		Customers: customers,
	}

	var constructedUrl string
	var offset int
	for i := 0; i < totalPages; i++ {
		offset = i * DefaultOffset
		constructedUrl = fmt.Sprintf("/?firstName=%s&lastName=%s&offset=%d",
			searchParams.FirstName, searchParams.LastName, offset)
		customersPage.Pages = append(customersPage.Pages, constructedUrl)
	}

	return customersPage, nil
}

func (rep CustomerRepository) Create(customer *model.Customer) error {
	stmt, err := rep.DB.Prepare("INSERT INTO customers (first_name, last_name, birth_date, gender, email, address) " +
		"VALUES($1, $2, $3, $4, $5, $6)")

	if err != nil {
		return err
	}

	stmt.Exec(customer.FirstName, customer.LastName, customer.BirthDate, customer.Gender, customer.Email, customer.Address)

	return nil
}

func (rep CustomerRepository) Update(customer *model.Customer) error {
	//TODO implement me
	panic("implement me")
}

func (rep CustomerRepository) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
