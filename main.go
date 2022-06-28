package main

import (
	"fmt"

	"github.com/AndiOrler/goapp/api"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("programm start")

	e := echo.New()

	api.Api(e)

}
