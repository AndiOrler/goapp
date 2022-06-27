package main

import (
	"fmt"
	"net/http"
	"strconv"

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

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)

	if err := c.Bind(u); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))

	users[id].Name = u.Name

	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	delete(users, id)

	return c.NoContent(http.StatusNoContent)
}

func getAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func main() {
	fmt.Println("programm start")

	e := echo.New()

	e.GET("/", greeting)
	e.POST("/user", createUser)
	e.GET("/user/:id", getUser)
	e.PUT("/user/:id", updateUser)
	e.DELETE("/user/:id", deleteUser)
	e.GET("/users", getAllUsers)

	e.Logger.Fatal(e.Start(":3000"))
}
