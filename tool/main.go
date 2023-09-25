package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func test() error {
	return errors.New("this is an error")
}

func main() {
	test()
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Dockercon 2023 🐋!\n")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
