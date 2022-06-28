package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AndiOrler/goapp/models"

	"github.com/labstack/echo/v4"
)

var (
	// users = map[int]*models.User{}
	seq = 1
)

func Greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Hello my friend! :)")
}

func (h handler) CreateUser(c echo.Context) error {
	u := &models.User{
		ID: seq,
	}

	if err := c.Bind(u); err != nil {
		return err
	}

	if result := h.DB.Create(&u); result.Error != nil {
		fmt.Println(result.Error)
	}

	// users[u.ID] = u
	seq++

	return c.JSON(http.StatusCreated, u)
}

func (h handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	u := &models.User{
		ID: id,
	}

	if result := h.DB.First(&u); result.Error != nil {
		fmt.Println(result.Error)
	}

	return c.JSON(http.StatusOK, u)
}

func (h handler) UpdateUser(c echo.Context) error {
	u := new(models.User)

	if err := c.Bind(u); err != nil {
		return err
	}

	var user models.User

	if result := h.DB.First(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	user.Name = u.Name

	if result := h.DB.Save(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	// users[id].Name = u.Name

	return c.JSON(http.StatusOK, user)
}

func (h handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	u := &models.User{
		ID: id,
	}

	if result := h.DB.Delete(&u); result.Error != nil {
		fmt.Println(result.Error)
	}

	// delete(users, id)

	return c.NoContent(http.StatusNoContent)
}

func (h handler) GetAllUsers(c echo.Context) error {
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		fmt.Println(result.Error)
	}
	return c.JSON(http.StatusOK, users)
}
