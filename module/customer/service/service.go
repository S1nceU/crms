package service

import (
	"go.mod/src/crms/model"
	"go.mod/src/crms/module/customer"
)

type CustomerService struct {
	repo customer.Repository
}

func NewCustomer(repo customer.Repository) customer.Service {
	return &CustomerService{
		repo: repo,
	}
}

func (u *CustomerService) GetCustomerList(map[string]interface{}) ([]*model.Customer, error) {
	return u.repo.GetCustomerList()
}
func (u *CustomerService) GetCustomer(in *model.Customer) (*model.Customer, error) {
	return u.repo.GetCustomer(in)
}
func (u *CustomerService) CreateCustomer(in *model.Customer) (*model.Customer, error) {
	return u.repo.CreateCustomer(in)
}
func (u *CustomerService) UpdateCustomer(in *model.Customer) (*model.Customer, error) {
	return u.repo.UpdateCustomer(in)
}
func (u *CustomerService) DeleteCustomer(in *model.Customer) error {
	return u.repo.DeleteCustomer(in)
}
