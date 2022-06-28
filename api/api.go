package api

import (
	"github.com/AndiOrler/goapp/handler"
	"github.com/labstack/echo/v4"
)

func Api(e *echo.Echo) *echo.Echo {

	e.GET("/", handler.Greeting)
	e.POST("/user", handler.CreateUser)
	e.GET("/user/:id", handler.GetUser)
	e.PUT("/user/:id", handler.UpdateUser)
	e.DELETE("/user/:id", handler.DeleteUser)
	e.GET("/users", handler.GetAllUsers)

	// e.Logger.Fatal(e.Start(":3000"))

	return e

}
