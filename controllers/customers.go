package controllers

import (
	"net/http"
	"strconv"

	"github.com/inasalifatus/bank-payment/lib/database"
	"github.com/inasalifatus/bank-payment/middlewares"
	"github.com/inasalifatus/bank-payment/models"
	"github.com/labstack/echo"
)

func GetOneCustomersByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	customers, err := database.GetOneCustomerById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"users":   customers,
	})
}

func CreateCustomersController(c echo.Context) error {
	customer := models.Customers{}
	c.Bind(&customer)

	customerAdd, err := database.CreateCustomer(customer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    customerAdd,
	})
}

func DeleteCustomersByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	customers, err := database.GetOneCustomerById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	customersdeleted, err := database.DeleteCustomersById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":                "success delete customer selected",
		"customer before delete": customers,
		"customer after delete":  customersdeleted,
	})
}

func UpdateCustomersController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	customers := database.GetUpdateCustomers(id)
	c.Bind(&customers)
	customersUpdate, err1 := database.UpdateCustomers(customers)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot post data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success update customer",
		"update customer": customersUpdate,
	})
}

func UserAuthorized(customerId int, c echo.Context) error {
	customer, err := database.GetOneCustomerById(customerId)
	loggedInCustomerId := middlewares.ExtractTokenCustomerId(c)
	if loggedInCustomerId != int(customer.ID) || err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access")
	}
	return nil
}

func LoginCustomerController(c echo.Context) error {
	customer := models.Customers{}
	c.Bind(&customer)
	customers, err := database.LoginCustomers(customer.Email, customer.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Please check your Email and password",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login Success",
		"users":   customers,
	})
}

func LogoutCustomerController(c echo.Context) error {
	customerId, err := strconv.Atoi(c.Param("customer_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID",
		})
	}
	if err = UserAuthorized(customerId, c); err != nil {
		return err
	}
	logout, err := database.GetOneCustomerById(customerId)
	if logout.Token == "" {
		return c.JSON(http.StatusBadRequest, "Please login again")
	}
	logout.Token = ""

	c.Bind(&logout)
	customer, err := database.UpdateCustomers(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout, try again",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Thank you for using our service",
		"Customer ID": customer.ID,
		"Username":    customer.Username,
	})
}
