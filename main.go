package main

import (
	"fmt"

	"github.com/AndiOrler/goapp/api"
	"github.com/AndiOrler/goapp/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("programm start")

	db := database.Init()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	api.Api(e, db)

	e.Logger.Fatal(e.Start(":3001"))

}
