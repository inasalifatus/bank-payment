package controllers

import (
	"net/http"
	"strconv"

	"github.com/inasalifatus/bank-payment/lib/database"
	"github.com/inasalifatus/bank-payment/models"
	"github.com/labstack/echo"
)

func GetAllPaymentsController(c echo.Context) error {
	payments, err := database.GetPayments()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all payment",
		"data":   payments,
	})
}

func GetOnePaymentsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	payments, err := database.GetPaymentsById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//custom data for body response
	output := map[string]interface{}{
		"CreatedAt": payments.CreatedAt,
		"UpdatedAt": payments.UpdatedAt,
		"DeletedAt": payments.DeletedAt,
		"id":        payments.ID,
		"name":      payments.Name,
		"status":    payments.Status,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get payment data",
		"data":    output,
	})
}

func CreatePaymentsController(c echo.Context) error {
	// binding data
	payment := models.Payments{}
	c.Bind(&payment)
	payments, err := database.CreatePayments(payment)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//custom data for body response
	output := map[string]interface{}{
		"CreatedAt": payments.CreatedAt,
		"UpdatedAt": payments.UpdatedAt,
		"DeletedAt": payments.DeletedAt,
		"id":        payments.ID,
		"name":      payments.Name,
		"status":    payments.Status,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create payment",
		"data":     output,
	})
}

func DeletePaymentsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	payments, err := database.DeletePaymentsById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete payment",
		"data":    payments,
	})

}

func UpdatePaymentsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	payment := database.GetUpdatePayment(id)
	c.Bind(&payment)
	paymentUpdate, err := database.UpdatePayment(payment)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//custom data for body response
	output := map[string]interface{}{
		"CreatedAt": paymentUpdate.CreatedAt,
		"UpdatedAt": paymentUpdate.UpdatedAt,
		"DeletedAt": paymentUpdate.DeletedAt,
		"id":        paymentUpdate.ID,
		"name":      paymentUpdate.Name,
		"status":    paymentUpdate.Status,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating payment",
		"data":    output,
	})
}
