package payment

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	paymentService PaymentService
}

func NewController(paymentService PaymentService) *controller {
	return &controller{paymentService}
}

func (cn *controller) GetPayments(c *gin.Context) {
	payments, err := cn.paymentService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   err,
		})
		return
	}

	var paymentsResponse []PaymentResponse

	for _, payment := range payments {
		paymentResponse := convertToPaymentResponse(payment)

		paymentsResponse = append(paymentsResponse, paymentResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Success!",
		"data":  paymentsResponse,
	})
}

func (cn *controller) GetPayment(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid payment ID",
		})
		return
	}

	payment, err := cn.paymentService.FindPaymentByID(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Payment not found" {
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
		"data":  convertToPaymentResponse(payment),
	})
}

func (cn *controller) CreatePayment(c *gin.Context) {
	var paymentRequest PaymentCreateRequest

	err := c.ShouldBindJSON(&paymentRequest)

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

	paymentRequest.UserID = int(userID.(uint64))

	payment, err := cn.paymentService.CreatePayment(paymentRequest)

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
		"data":  convertToPaymentResponse(payment),
	})
}

func (cn *controller) UpdatePayment(c *gin.Context) {
	var paymentRequest PaymentUpdateRequest

	err := c.ShouldBindJSON(&paymentRequest)

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
			"msg":   "Invalid payment ID",
		})
		return
	}

	payment, err := cn.paymentService.UpdatePayment(id, paymentRequest)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Payment not found" {
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
		"data":  convertToPaymentResponse(payment),
	})
}

func (ch *controller) DeletePayment(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Invalid payment ID",
		})
		return
	}

	payment, err := ch.paymentService.DeletePayment(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Payment not found" {
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
		"data":  convertToPaymentResponse(payment),
	})
}

func convertToPaymentResponse(payment Payment) PaymentResponse {
	return PaymentResponse{
		ID:             payment.ID,
		UserID:         payment.UserID,
		BankID:         payment.BankID,
		DeliveryID:     payment.DeliveryID,
		TotalPrice:     payment.TotalPrice,
		PaymentStatus:  payment.PaymentStatus,
		DeliveryStatus: payment.DeliveryStatus,
	}
}
