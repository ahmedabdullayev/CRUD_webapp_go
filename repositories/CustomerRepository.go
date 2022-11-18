package repositories

import (
	"CRUD_webapp_go/CustomerHelpers"
	"CRUD_webapp_go/contracts"
	"CRUD_webapp_go/model"
	"database/sql"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const DefaultLimit = 5
const DefaultOffset = 5
const DefaultOrderBy = "id"
const DefaultOrderType = "DESC"

type CustomerRepository struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) contracts.CustomerRepositoryInterface {
	return &CustomerRepository{DB: db}
}

func (rep CustomerRepository) GetOne(id int) (model.Customer, error) {
	customer := model.Customer{}
	row := rep.DB.QueryRow("SELECT * FROM customers WHERE id = $1", id)

	err := row.Scan(&customer.Id, &customer.FirstName,
		&customer.LastName, &customer.BirthDate,
		&customer.Gender, &customer.Email, &customer.Address)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (rep CustomerRepository) GetAll(searchParams model.SearchParams) (model.CustomersPage, error) {
	var customersPage model.CustomersPage
	var totalRows int
	err := rep.DB.QueryRow("SELECT COUNT(*) FROM customers WHERE "+
		"LOWER(first_name) LIKE '%' || $1 || '%' "+
		"AND LOWER(last_name) LIKE '%' || $2 || '%'",
		strings.ToLower(searchParams.FirstName), strings.ToLower(searchParams.LastName)).Scan(&totalRows)
	if err != nil {
		return model.CustomersPage{}, err
	}

	orderBy := DefaultOrderBy
	orderType := DefaultOrderType

	if searchParams.OrderBy != "" {
		orderBy = searchParams.OrderBy
	}
	if searchParams.OrderType != "" {
		orderType = searchParams.OrderType
	}

	var queryString string
	if searchParams.Offset < DefaultOffset {
		queryString = "SELECT * FROM customers WHERE " +
			"LOWER(first_name) LIKE '%' || $1 || '%' " +
			"AND LOWER(last_name) LIKE '%' || $2 || '%' ORDER BY " + orderBy + " " + orderType + " LIMIT $3"
	} else {
		queryString = "SELECT * FROM customers WHERE LOWER(first_name) LIKE '%' || $1 || '%' " +
			"AND LOWER(last_name) LIKE '%' || $2 || '%' ORDER BY " + orderBy + " " + orderType + " LIMIT $3 OFFSET " +
			strconv.Itoa(searchParams.Offset)
	}

	customersQuery, err := rep.DB.Query(queryString,
		strings.ToLower(searchParams.FirstName),
		strings.ToLower(searchParams.LastName),
		DefaultLimit)

	if err != nil {
		return customersPage, err
	}

	customers, _ := CustomerHelpers.CustomersToList(customersQuery)

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

	_, err = stmt.Exec(customer.FirstName, customer.LastName, customer.BirthDate, customer.Gender,
		customer.Email, customer.Address)
	if err != nil {
		return err
	}

	return nil
}

func (rep CustomerRepository) Update(customer *model.Customer) error {
	query, err := rep.DB.Prepare("UPDATE customers SET first_name=$1, last_name=$2, birth_date=$3, gender=$4, " +
		"email=$5, address=$6 WHERE id=$7")

	if err != nil {
		return err
	}

	_, err = query.Exec(customer.FirstName, customer.LastName, customer.BirthDate, customer.Gender, customer.Email,
		customer.Address, customer.Id)

	if err != nil {
		return err
	}

	return nil
}

func (rep CustomerRepository) Delete(id int) error {
	query, err := rep.DB.Prepare("DELETE FROM customers WHERE id=$1")
	if err != nil {
		return err
	}
	query.Exec(id)

	return nil
}
