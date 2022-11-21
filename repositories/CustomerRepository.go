package repositories

import (
	"CRUD_webapp_go/contracts"
	"CRUD_webapp_go/customerHelpers"
	"CRUD_webapp_go/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"math"
	"strconv"
	"strings"
)

const DefaultLimit = 5
const DefaultOffset = 5
const DefaultOrderBy = "id"
const DefaultOrderType = "DESC"

type CustomerRepository struct {
	DB *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) contracts.CustomerRepositoryInterface {
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

	customers, _ := customerHelpers.CustomersToList(customersQuery)

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

func (rep CustomerRepository) Create(customer *model.Customer) (int, error) {
	stmt, err := rep.DB.PrepareNamed("INSERT INTO customers (first_name, last_name, birth_date, gender, email, address) " +
		"VALUES(:first_name, :last_name, :birth_date, :gender, :email, :address) RETURNING id")

	if err != nil {
		return 0, err
	}

	var insertedId int

	err = stmt.Get(&insertedId, customer)

	if err != nil {
		return 0, err
	}

	return insertedId, nil
}

func (rep CustomerRepository) Update(customer *model.Customer) (int, error) {
	query, err := rep.DB.PrepareNamed("UPDATE customers SET first_name = :first_name, last_name = :last_name, " +
		"birth_date = :birth_date, gender = :gender, email = :email, address = :address WHERE id= :id RETURNING id")

	if err != nil {
		return 0, err
	}

	var updatedId int

	err = query.Get(&updatedId, customer)

	if err != nil {
		return 0, err
	}

	return updatedId, err
}

func (rep CustomerRepository) Delete(id int) error {
	query, err := rep.DB.PrepareNamed("DELETE FROM customers WHERE id=$1")

	if err != nil {
		return err
	}

	_, err = query.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
