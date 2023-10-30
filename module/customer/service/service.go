package service

import (
	"crms/model"
	"crms/module/customer"
	"encoding/json"
	"errors"
)

type CustomerService struct {
	repo customer.Repository
}

func NewCustomerService(repo customer.Repository) customer.Service {
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
	newCustomer := &model.Customer{
		Citizenship: in,
	}
	point, err = u.repo.GetCustomerListForCitizenship(newCustomer)
	for i := 0; i < len(point); i++ {
		out = append(out, *point[i])
	}
	return out, err
}

func (u *CustomerService) GetCustomer(in string) (*model.Customer, error) {
	var err error
	newCustomer := &model.Customer{
		ID: in,
	}
	if newCustomer, err = u.repo.GetCustomer(newCustomer); newCustomer.CustomerId == 0 {
		return nil, errors.New("error CRMS : There is no this customer")
	}
	return newCustomer, err
}

func (u *CustomerService) GetCustomerForCID(in int) (*model.Customer, error) {
	var err error
	newCustomer := &model.Customer{
		CustomerId: in,
	}
	if newCustomer, err = u.repo.GetCustomerForCID(newCustomer); newCustomer.Name == "" {
		return nil, errors.New("error CRMS : There is no this customer")
	}
	return newCustomer, err
}

func (u *CustomerService) CreateCustomer(in []byte) (*model.Customer, error) {
	var err error
	var newCustomer *model.Customer
	if err = json.Unmarshal(in, &newCustomer); err != nil {
		return nil, err
	}
	if newCustomer, err = u.repo.GetCustomer(newCustomer); newCustomer.CustomerId == 0 {
		newCustomer, err = u.repo.CreateCustomer(newCustomer)
		return newCustomer, err
	} else {
		return nil, errors.New("error CRMS : This customer is already existed")
	}
}

func (u *CustomerService) UpdateCustomer(in []byte) (*model.Customer, error) {
	var err error
	var newCustomer *model.Customer
	if err = json.Unmarshal(in, &newCustomer); err != nil {
		return nil, err
	}
	if _, err = u.GetCustomerForCID(newCustomer.CustomerId); err != nil {
		return nil, err
	}
	if newCustomer, err = u.repo.UpdateCustomer(newCustomer); err != nil {
		return nil, err
	}
	return newCustomer, err
}

func (u *CustomerService) DeleteCustomer(in int) error {
	var err error
	newCustomer := &model.Customer{
		CustomerId: in,
	}
	if _, err = u.GetCustomerForCID(newCustomer.CustomerId); err != nil {
		return err
	}
	if err = u.repo.DeleteCustomer(newCustomer); err != nil {
		return err
	}
	return nil
}
