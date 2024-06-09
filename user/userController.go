package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	userService UserService
}

func NewController(userService UserService) *controller {
	return &controller{userService}
}

func (cn *controller) GetUsers(c *gin.Context) {
	users, err := cn.userService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var usersResponse []UserResponse

	for _, user := range users {
		userResponse := convertToUserResponse(user)

		usersResponse = append(usersResponse, userResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": usersResponse,
	})
}

func (cn *controller) GetUser(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Invalid user ID",
		})
		return
	}

	user, err := cn.userService.FindUserByID(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "User not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToUserResponse(user),
	})
}

func (cn *controller) CreateUser(c *gin.Context) {
	var userRequest UserCreateRequest

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on %s field, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	user, err := cn.userService.CreateUser(userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToUserResponse(user),
	})
}

func (cn *controller) UpdateUser(c *gin.Context) {
	var userRequest UserUpdateRequest

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on %s field, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Invalid user ID",
		})
		return
	}

	user, err := cn.userService.UpdateUser(id, userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToUserResponse(user),
	})
}

func convertToUserResponse(user User) UserResponse {
	return UserResponse{
		ID:                 user.ID,
		UserCode:           user.UserCode,
		Role:               user.Role,
		UserName:           user.UserName,
		Email:              user.Email,
		Password:           user.Password,
		UserWhatsappNumber: user.UserWhatsappNumber,
		UserGender:         user.UserGender,
	}
}
