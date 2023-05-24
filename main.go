package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)

}

func main() {

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "PONG!")
	})

	e.GET("/user/:id", getUser)

	e.Logger.Fatal(e.Start(":9999"))
}
