package database

import (
	"github.com/inasalifatus/bank-payment/config"
	"github.com/inasalifatus/bank-payment/middlewares"
	"github.com/inasalifatus/bank-payment/models"
)

func GetOneCustomerById(customerId int) (models.Customers, error) {
	var customer models.Customers
	if err := config.DB.Where("id=?", customerId).First(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func CreateCustomer(customers models.Customers) (interface{}, error) {
	if err := config.DB.Save(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

// login customer used to matching data in database
func LoginCustomers(email, password string) (interface{}, error) {
	var customer models.Customers
	var err error
	if err = config.DB.Where("email = ? AND password = ?", email, password).First(&customer).Error; err != nil {
		return nil, err
	}
	customer.Token, err = middlewares.CreateToken(int(customer.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(customer).Error; err != nil {
		return nil, err
	}
	return customer, err
}

func DeleteCustomersById(id int) (interface{}, error) {
	var customers models.Customers
	if err := config.DB.Where("id=?", id).Delete(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

//to update user information from database
func UpdateCustomers(customer models.Customers) (models.Customers, error) {
	if err := config.DB.Save(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func GetUpdateCustomers(id int) models.Customers {
	var customers models.Customers
	config.DB.Find(&customers, "id=?", id)
	return customers
}

func GetToken(id int) string {
	var customer models.Customers
	config.DB.Model(&customer).Select("token").Where("id=?", id)
	return customer.Token
}
