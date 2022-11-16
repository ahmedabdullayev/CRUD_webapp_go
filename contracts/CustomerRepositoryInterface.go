package contracts

import "CRUD_webapp_go/model"

type CustomerRepositoryInterface interface {
	GetAll() ([]model.Customer, error)
	GetBySearch(name string) ([]model.Customer, error)
	Create(customer *model.Customer) error
	UpdateOne(customer *model.Customer) error
	Delete(id int) error
}
