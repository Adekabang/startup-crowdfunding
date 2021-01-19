package handler

import (
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

//UserHandler interface handler
type UserHandler interface {
	RegisterUser(*gin.Context)
}

type userHandler struct {
	userService user.Service
}

//NewUserHandler create handler with service
func NewUserHandler(userService user.Service) UserHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusCreated, user)
}
