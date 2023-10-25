package customer

import "crms/model"

type Repository interface {
	GetCustomerList() ([]*model.Customer, error)                                 // Get all Customer
	GetCustomerListForCitizenship(in *model.Customer) ([]*model.Customer, error) // Get all Customer by Citizenship
	GetCustomer(in *model.Customer) (*model.Customer, error)                     // Get Customer by ID
	GetCustomerForCID(in *model.Customer) (*model.Customer, error)               // Get Customer by CustomerId
	CreateCustomer(in *model.Customer) (*model.Customer, error)                  // Create a new customer
	UpdateCustomer(in *model.Customer) (*model.Customer, error)                  // Update Customer data
	DeleteCustomer(in *model.Customer) error                                     // Delete Customer by CustomerId
}
