package main

import (
	"fmt"

	"github.com/inasalifatus/bank-payment/config"
	"github.com/inasalifatus/bank-payment/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.InitDB()
	config.InitPort()
	routes.New(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.PORT)))
}
