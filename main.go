package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type user struct {
	ID   int    `Json: ID`
	Name string `Json: Name`
}

var (
	users = map[int]*user{}
	seq   = 1
)

func main() {
	fmt.Println("programm start")

	e := echo.New()

	e.GET("/", greeting)
	e.POST("/user", createUser)
	e.GET("/users", getAllUsers)

	e.Logger.Fatal(e.Start(":3000"))
}

func greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Hello my friend! :)")
}

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}

	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++

	return c.JSON(http.StatusCreated, u)
}

func getAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}
