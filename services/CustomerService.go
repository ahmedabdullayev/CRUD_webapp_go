package services

import (
	"CRUD_webapp_go/contracts"
	"CRUD_webapp_go/model"
)

type CustomerService struct {
	repository contracts.CustomerRepositoryInterface
}

func NewEmployeeService(repository contracts.CustomerRepositoryInterface) contracts.CustomerServiceInterface {
	return &CustomerService{
		repository: repository,
	}
}

func (service CustomerService) GetAllByParams(searchParams model.SearchParams) (model.CustomersPage, error) {
	return service.repository.GetAll(searchParams)
}

func (service CustomerService) GetOneById(id int) (model.Customer, error) {
	return service.repository.GetOne(id)
}

func (service CustomerService) Create(customer *model.Customer) (int, error) {
	return service.repository.Create(customer)
}

func (service CustomerService) Update(customer *model.Customer) (int, error) {
	return service.repository.Update(customer)
}

func (service CustomerService) Delete(id int) error {
	return service.repository.Delete(id)
}
