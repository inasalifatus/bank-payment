package main

import (
	"fmt"

	"github.com/inasalifatus/bank-payment/config"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.InitPort()

	Port := fmt.Sprintf(":%d", config.PORT)
	if err := e.Start(Port); err != nil {
		e.Logger.Fatal(err)
	}
	fmt.Println("Application started...")
}
