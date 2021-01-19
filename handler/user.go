package handler

import (
	"bwastartup/helper"
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
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if err != nil {

		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userFormat := user.FormatUser(newUser, "tokenabcd")
	response := helper.APIResponse("Account has been registered", http.StatusCreated, "success", userFormat)

	c.JSON(http.StatusCreated, response)
}
