package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	port := 1323
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	address := fmt.Sprintf("%s:%v", "0.0.0.0", port)
	e.Start(address)
}
