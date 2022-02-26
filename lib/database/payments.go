package database

import (
	"github.com/inasalifatus/bank-payment/config"
	"github.com/inasalifatus/bank-payment/models"
)

// creating new payment
func CreatePayments(createPayment models.Payments) (models.Payments, error) {
	if err := config.DB.Save(&createPayment).Error; err != nil {
		return createPayment, err
	}
	return createPayment, nil
}

//function get all payment  table
func GetPayments() (interface{}, error) {
	var payments []models.Payments
	if err := config.DB.Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

// get payment from id
func GetPaymentsById(id int) (models.Payments, error) {
	var payments models.Payments
	if err := config.DB.Find(&payments, "ID=?", id).Error; err != nil {
		return payments, err
	}
	return payments, nil
}

//check is payment exist on table payments
func CheckPayment(paymentId int, payment models.Payments) (interface{}, error) {
	if err := config.DB.Where("id=?", paymentId).First(&payment).Error; err != nil {
		return nil, err
	}
	return payment.ID, nil
}

// deleting payment by id
func DeletePaymentsById(id int) (interface{}, error) {
	var payments []models.Payments
	if err := config.DB.Find(&payments, "ID=?", id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&payments, "ID=?", id).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

//get 1 specified payment
func GetUpdatePayment(id int) models.Payments {
	var payment models.Payments
	config.DB.Find(&payment, "id=?", id)
	return payment
}

//update payment  from database
func UpdatePayment(payment models.Payments) (models.Payments, error) {
	if err := config.DB.Save(&payment).Error; err != nil {
		return payment, err
	}
	return payment, nil
}
