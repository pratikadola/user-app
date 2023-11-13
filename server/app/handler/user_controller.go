package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/user-assignment/app/models"
	"github.com/user-assignment/app/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{userService: service}
}

func (uc *UserController) CreateUser(c echo.Context) error {
	name := c.QueryParam("name")
	u := models.User{Name: name}
	resp, err := uc.userService.CreateUser(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, resp)
}

func (uc *UserController) UpdateUser(c echo.Context) error {
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	u := models.User{ID: uint(idVal), Name: name}
	resp, err := uc.userService.UpdateUser(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, resp)
}

func (uc *UserController) DeleteUser(c echo.Context) error {
	name := c.QueryParam("name")
	u := models.User{Name: name}
	err := uc.userService.DeleteUser(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "deleted")
}

func (uc *UserController) GelAllUsers(c echo.Context) error {
	users := uc.userService.GetAllUsers()
	return c.JSON(http.StatusOK, users)
}
