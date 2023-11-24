package customer

import "github.com/S1nceU/CRMS/model"

type Service interface {
	ListCustomers() ([]model.Customer, error)                      // Get all Customers
	GetCustomersByCitizenship(in string) ([]model.Customer, error) // Get all Customers by citizenship
	GetCustomerByID(in string) (*model.Customer, error)            // Get Customer by ID
	GetCustomerByCustomerId(in int) (*model.Customer, error)       // Get Customer by customer_id
	CreateCustomer(in *model.Customer) (*model.Customer, error)    // Create a new Customer
	UpdateCustomer(in *model.Customer) (*model.Customer, error)    // Update customer data
	DeleteCustomer(in int) error                                   // Delete Customer by customer_id
}
