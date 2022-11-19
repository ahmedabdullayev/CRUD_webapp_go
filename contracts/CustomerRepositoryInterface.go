package contracts

import "CRUD_webapp_go/model"

type CustomerRepositoryInterface interface {
	GetAll(searchParams model.SearchParams) (model.CustomersPage, error)
	GetOne(id int) (model.Customer, error)
	Create(customer *model.Customer) (int, error)
	Update(customer *model.Customer) (int, error)
	Delete(id int) error
}
