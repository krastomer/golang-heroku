package main

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/labstack/echo"
)

func main() {
	server := echo.New()
	server.GET(path.Join("/"), Version)

	port := os.Getenv("PORT")

	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	server.Start(address)
}

func Version(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{"version": 1})
}
