package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Mahanmmi/Go-Remote-Calculator/calculator"
	"github.com/labstack/echo"
)

func clientHandler(ao calculator.ArithmaticOperation) string {
	var ans float64
	switch ao.OperationType {
	case "sum":
		{
			ans = calculator.Sum(ao.A, ao.B)
		}
	case "subtract":
		{
			ans = calculator.Subtract(ao.A, ao.B)
		}
	case "multiply":
		{
			ans = calculator.Multiply(ao.A, ao.B)
		}
	case "divide":
		{
			ans = calculator.Divide(ao.A, ao.B)
		}
	}
	strAns := strconv.FormatFloat(ans, 'f', 2, 64)
	return strAns
}

func handler(c echo.Context) error {
	params := c.QueryParams()
	A, _ := strconv.ParseFloat(params["A"][0], 64)
	B, _ := strconv.ParseFloat(params["B"][0], 64)
	oper := calculator.ArithmaticOperation{
		params["type"][0],
		A,
		B,
	}

	fmt.Println(oper)

	return c.JSON(http.StatusOK, clientHandler(oper))
}

// InitializeServer TCP on port 18757
func InitializeServer() {
	e := echo.New()
	e.GET("/calculator", handler)
	fmt.Println("ere")
	e.Logger.Fatal(e.Start("127.0.0.1:18757"))
}
