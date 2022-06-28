package api

import (
	"github.com/AndiOrler/goapp/handler"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Api(e *echo.Echo, db *gorm.DB) *echo.Echo {

	h := handler.New(db)

	e.GET("/", handler.Greeting)
	e.POST("/user", h.CreateUser)
	e.GET("/user/:id", h.GetUser)
	e.PUT("/user/:id", h.UpdateUser)
	e.DELETE("/user/:id", h.DeleteUser)
	e.GET("/users", h.GetAllUsers)

	// e.Logger.Fatal(e.Start(":3000"))

	return e

}
