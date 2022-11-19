package contracts

import "CRUD_webapp_go/model"

type CustomerServiceInterface interface {
	GetAllByParams(searchParams model.SearchParams) (model.CustomersPage, error)
	GetOneById(id int) (model.Customer, error)
	Create(customer *model.Customer) (int, error)
	Update(customer *model.Customer) (int, error)
	Delete(id int) error
}
