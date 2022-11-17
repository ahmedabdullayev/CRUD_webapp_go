package contracts

import "CRUD_webapp_go/model"

type CustomerRepositoryInterface interface {
	GetAll(searchParams model.SearchParams) (model.CustomersPage, error)
	GetOne(id int) (model.Customer, error)
	Create(customer *model.Customer) error
	Update(customer *model.Customer) error
	Delete(id int) error
}
