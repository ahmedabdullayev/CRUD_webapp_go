package repositories

import (
	"CRUD_webapp_go/contracts"
	"CRUD_webapp_go/model"
	"database/sql"
)

type CustomerRepository struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) contracts.CustomerRepositoryInterface {
	return &CustomerRepository{DB: db}
}

func (c CustomerRepository) GetAll() ([]model.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomerRepository) GetBySearch(name string) ([]model.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomerRepository) Create(customer *model.Customer) error {
	//TODO implement me
	panic("implement me")
}

func (c CustomerRepository) UpdateOne(customer *model.Customer) error {
	//TODO implement me
	panic("implement me")
}

func (c CustomerRepository) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
