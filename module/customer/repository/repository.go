package repository

import (
	"errors"
	"go.mod/src/crms/model"
	"go.mod/src/crms/module/customer"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	orm *gorm.DB
}

func NewCustomerRepository(orm *gorm.DB) customer.Repository {
	return &CustomerRepository{
		orm: orm,
	}
}

func (u *CustomerRepository) GetCustomerList() ([]*model.Customer, error) {
	var err error
	var in []*model.Customer
	err = u.orm.Find(&in).Error
	return in, err
}

func (u *CustomerRepository) GetCustomer(in *model.Customer) (*model.Customer, error) {
	var err error

	if err = u.orm.Where("ID = ?", in.ID).Find(&in).Error; in.Customer_id == 0 {
		return nil, errors.New("Error CRMS : There is no this customer.")
	}
	return in, err
}

func (u *CustomerRepository) GetCustomerForCID(in *model.Customer) (*model.Customer, error) {
	var err error

	if err = u.orm.Where("customer_id = ?", in.Customer_id).Find(&in).Error; in.Name == "" {
		return nil, errors.New("Error CRMS : There is no this customer.")
	}
	return in, err
}

func (u *CustomerRepository) CreateCustomer(in *model.Customer) (*model.Customer, error) {
	var err error
	err = u.orm.Create(in).Error
	return in, err
}

func (u *CustomerRepository) UpdateCustomer(in *model.Customer) (*model.Customer, error) {
	var err error
	if _, err = u.GetCustomerForCID(in); err != nil {
		return nil, err
	}
	err = u.orm.Save(&in).Error
	return in, err
}

func (u *CustomerRepository) DeleteCustomer(in *model.Customer) error {
	if _, err := u.GetCustomerForCID(in); err != nil {
		return err
	}
	err := u.orm.Where("customer_id = ?", in.Customer_id).Delete(&in).Error
	return err
}

func (u *CustomerRepository) GetCustomerListForCitizenship(in *model.Customer) ([]*model.Customer, error) {
	var err error
	var out []*model.Customer
	if err = u.orm.Where("Citizenship = ?", in.Citizenship).Find(&out).Error; err != nil {
		return nil, err
	}
	return out, err
}
