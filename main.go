package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	fmt.Println("Hello, 世界")

	e := echo.New()
	e.GET("/hello", Greetings)
	e.Logger.Fatal(e.Start(":3000"))

}

func Greetings(c echo.Context) error {
	return c.JSON(http.StatusOK, Message{
		Message: "Hello World!",
	})
}
