package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Student struct{
	NIM string `json:"nim"`
	FirstName string `json:"namaAwal"`
	LastName string
	Class string
	Age int
}


func main() {
	e := echo.New()

	oki := Student{
		NIM: "123",
		FirstName: "Oki",
		LastName: "Syauqi",
		Class: "A",
		Age: 21,
	}

	e.GET("/awal/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, oki)
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Halaman Awal!")
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, there!")
	})

	e.GET("/bye", func(c echo.Context) error {
		return c.String(http.StatusOK, "Goodbye!")
	})

	e.Logger.Fatal(e.Start("localhost:1234"))
}