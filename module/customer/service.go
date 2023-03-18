package customer

import "go.mod/src/crms/model"

type Service interface {
	GetCustomerList(string, string) ([]*model.Customer, error)
	GetCustomer(in *model.Customer) (*model.Customer, error)
	CreateCustomer(in *model.Customer) (*model.Customer, error)
	UpdateCustomer(in *model.Customer) (*model.Customer, error)
	DeleteCustomer(in *model.Customer) error
}
