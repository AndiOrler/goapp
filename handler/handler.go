package handler

import (
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

func Greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Hello my friend! :)")
}

func CreateUser(c echo.Context) error {
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

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	return c.JSON(http.StatusOK, users[id])
}

func UpdateUser(c echo.Context) error {
	u := new(user)

	if err := c.Bind(u); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))

	users[id].Name = u.Name

	return c.JSON(http.StatusOK, users[id])
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	delete(users, id)

	return c.NoContent(http.StatusNoContent)
}

func GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}