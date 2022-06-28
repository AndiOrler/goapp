package main

import (
	"fmt"

	"github.com/AndiOrler/goapp/api"
	"github.com/AndiOrler/goapp/database"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("programm start")

	db := database.Init()
	e := echo.New()

	api.Api(e, db)

	e.Logger.Fatal(e.Start(":3000"))

}
