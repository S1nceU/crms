package customer

import "go.mod/src/crms/model"

type Repository interface {
	GetCustomerList() ([]*model.Customer, error)                                 // Get all Customer
	GetCustomerListForCitizenship(in *model.Customer) ([]*model.Customer, error) // Get all Customer by citizenship
	GetCustomer(in *model.Customer) (*model.Customer, error)                     // Get Customer by ID and citizenship
	GetCustomerForID(customer_id int) (*model.Customer, error)                   // Get Customer by customer_id
	CreateCustomer(in *model.Customer) (*model.Customer, error)                  // Create a new customer
	UpdateCustomer(in *model.Customer) (*model.Customer, error)                  // Update customer data by a whole customer
	DeleteCustomer(in int) error                                                 // Delete customer by customer_id

}
