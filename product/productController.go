package product

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type controller struct {
	productService ProductService
}

func NewController(productService ProductService) *controller {
	return &controller{productService}
}

func (cn *controller) GetProducts(c *gin.Context) {
	products, err := cn.productService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var productsResponse []ProductResponse

	for _, product := range products {
		productResponse := convertToProductResponse(product)

		productsResponse = append(productsResponse, productResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productsResponse,
	})
}

func (cn *controller) GetProduct(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Invalid product ID",
		})
		return
	}

	product, err := cn.productService.FindProductByID(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "product not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToProductResponse(product),
	})
}

func (cn *controller) CreateProduct(c *gin.Context) {
	var productRequest ProductCreateRequest

	err := c.ShouldBind(&productRequest)

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

	if err := os.MkdirAll("/public/product/image", 0755); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	userID, errorID := c.Get("UserID")

	if !errorID {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorID,
		})
	}

	productRequest.UserID = int(userID.(uint64))

	// rename file
	ext := filepath.Ext(productRequest.Image.Filename)
	newFileName := uuid.New().String() + ext

	// save file
	dst := filepath.Join("public/product/image", filepath.Base(newFileName))
	c.SaveUploadedFile(&productRequest.Image, dst)

	productRequest.Image.Filename = fmt.Sprintf("%s/public/product/image/%s", c.Request.Host, newFileName)
	product, err := cn.productService.CreateProduct(productRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToProductResponse(product),
	})
}

func (cn *controller) UpdateProduct(c *gin.Context) {
	var productRequest ProductUpdateRequest

	err := c.ShouldBind(&productRequest)

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

	if productRequest.Image != nil {
		if err := os.MkdirAll("/public/product/image", 0755); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}

		// rename file
		ext := filepath.Ext(productRequest.Image.Filename)
		newFileName := uuid.New().String() + ext

		// save file
		dst := filepath.Join("public/product/image", filepath.Base(newFileName))
		c.SaveUploadedFile(productRequest.Image, dst)

		productRequest.Image.Filename = fmt.Sprintf("%s/public/product/image/%s", c.Request.Host, newFileName)
	}

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Invalid product ID",
		})
		return
	}

	product, err := cn.productService.UpdateProduct(id, productRequest)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Product not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToProductResponse(product),
	})
}

func (ch *controller) DeleteProduct(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Invalid product ID",
		})
		return
	}

	product, err := ch.productService.DeleteProduct(id)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Product not found" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToProductResponse(product),
	})
}

func convertToProductResponse(product Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		UserID:      product.UserID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Image:       product.Image,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}
