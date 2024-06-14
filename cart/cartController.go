package cart

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	cartService CartService
}

func NewController(cartService CartService) *controller {
	return &controller{cartService}
}

func (cn *controller) GetCarts(c *gin.Context) {
	carts, err := cn.cartService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err,
		})
		return
	}

	var cartsResponse []CartResponse

	for _, cart := range carts {
		cartResponse := convertToCartResponse(cart)

		cartsResponse = append(cartsResponse, cartResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  cartsResponse,
	})
}

func (cn *controller) GetCart(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid cart ID",
		})
		return
	}

	cart, err := cn.cartService.FindCartByID(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Cart not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  convertToCartResponse(cart),
	})
}

func (cn *controller) CreateCart(c *gin.Context) {
	var cartRequest CartCreateRequest

	err := c.ShouldBindJSON(&cartRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on %s field, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   errorMessages,
		})
		return
	}

	userID, errorID := c.Get("UserID")

	if !errorID {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   errorID,
		})
	}

	cartRequest.UserID = int(userID.(uint64))

	cart, err := cn.cartService.CreateCart(cartRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  convertToCartResponse(cart),
	})
}

func (cn *controller) UpdateCart(c *gin.Context) {
	var cartRequest CartUpdateRequest

	err := c.ShouldBindJSON(&cartRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on %s field, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid cart ID",
		})
		return
	}

	cart, err := cn.cartService.UpdateCart(id, cartRequest)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Cart not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  convertToCartResponse(cart),
	})
}

func (ch *controller) DeleteCart(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid cart ID",
		})
		return
	}

	cart, err := ch.cartService.DeleteCart(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Cart not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  convertToCartResponse(cart),
	})
}

func convertToCartResponse(cart Cart) CartResponse {
	return CartResponse{
		ID:         cart.ID,
		UserID:     cart.UserID,
		ProductID:  cart.ProductID,
		PaymentID:  cart.PaymentID,
		Quantity:   cart.Quantity,
		TotalPrice: cart.TotalPrice,
		IsActived:  cart.IsActived,
	}
}
