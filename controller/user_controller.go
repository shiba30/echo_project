package controller

import (
	"net/http"
	"strconv"
	"twitter_echo_backend/model"
	"twitter_echo_backend/usecase"

	"github.com/labstack/echo"
)

type UserController interface {
	CreateUser(c echo.Context) error
	GetAllUsers(c echo.Context) error
	GetUserById(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type userController struct {
	uu usecase.UserUsecase
}

// コンストラクタの作成、main.goで呼び出し
func NewUserController(uu usecase.UserUsecase) UserController {
	return &userController{uu}
}

func (uc *userController) CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) GetAllUsers(c echo.Context) error {
	usersRes, err := uc.uu.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, usersRes)
}

func (uc *userController) GetUserById(c echo.Context) error {
	id := c.Param("userId")
	userId, _ := strconv.Atoi(id)
	userRes, err := uc.uu.GetUserById(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userRes)
}

func (uc *userController) UpdateUser(c echo.Context) error {
	id := c.Param("userId")
	userId, _ := strconv.Atoi(id)

	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.UpdateUser(user, uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userRes)
}

func (uc *userController) DeleteUser(c echo.Context) error {
	id := c.Param("userId")
	userId, _ := strconv.Atoi(id)

	err := uc.uu.DeleteUser(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
