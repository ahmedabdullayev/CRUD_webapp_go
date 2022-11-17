package contracts

import "CRUD_webapp_go/model"

type CustomerRepositoryInterface interface {
	GetAll(searchParams model.SearchParams) (model.CustomersPage, error)
	Create(customer *model.Customer) error
	UpdateOne(customer *model.Customer) error
	Delete(id int) error
}
