package contracts

import "CRUD_webapp_go/models"

type CustomerRepositoryInterface interface {
	GetAll(searchParams models.SearchParams) (models.CustomersPage, error)
	GetOne(id int) (models.Customer, error)
	Create(customer *models.Customer) (int, error)
	Update(customer *models.Customer) (int, error)
	Delete(id int) error
}
