package main

import (
	"notes/configs"
	"notes/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(":8000")
}
