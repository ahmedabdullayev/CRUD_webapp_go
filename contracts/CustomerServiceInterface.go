package contracts

import "CRUD_webapp_go/models"

type CustomerServiceInterface interface {
	GetAllByParams(searchParams models.SearchParams) (models.CustomersPage, error)
	GetOneById(id int) (models.Customer, error)
	Create(customer *models.Customer) (int, error)
	Update(customer *models.Customer) (int, error)
	Delete(id int) error
}
