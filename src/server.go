package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/images", "assets/images")
	e.Static("/favicon", "assets/favicon")
	e.Static("/css", "assets/css")
	e.File("/", "templates/index.html")

	e.Logger.Fatal(e.Start(":1323"))
}
