package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Square struct {
	Root   int `json:"root"`
	Square int `json:"square"`
}

func main() {
	e := echo.New()

	// Routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!")
	})
	e.GET("/square/:root", func(c echo.Context) error {
		root, _ := strconv.Atoi(c.Param("root"))
		return c.JSON(http.StatusOK, getSquare(root))
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func getSquare(root int) *Square {
	square := root * root
	s := &Square{
		Root:   root,
		Square: square,
	}
	return s
}
