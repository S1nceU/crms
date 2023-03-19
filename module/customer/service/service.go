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

func (u *CustomerService) GetCustomerList() ([]model.Customer, error) {
	var err error
	var point []*model.Customer
	var out []model.Customer
	point, err = u.repo.GetCustomerList()
	for i := 0; i < len(point); i++ {
		out = append(out, *point[i])
	}
	return out, err
}

func (u *CustomerService) GetCustomerListForCitizenship(in string) ([]model.Customer, error) {
	var err error
	var point []*model.Customer
	var out []model.Customer
	new_in := &model.Customer{
		Citizenship: in,
	}
	point, err = u.repo.GetCustomerListForCitizenship(new_in)
	for i := 0; i < len(point); i++ {
		out = append(out, *point[i])
	}
	return out, err
}
