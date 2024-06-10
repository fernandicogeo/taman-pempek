package main

import (
	"log"
	"taman-pempek/bank"
	"taman-pempek/product"
	"taman-pempek/productcategory"
	"taman-pempek/transaction"
	"taman-pempek/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/taman-pempek?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection error")
	}

	router := gin.Default()
	v1 := router.Group("/v1")

	migration(db)

	routeUser(db, v1)
	routeProduct(db, v1)

	router.Run(":8888") // port
}

func migration(db *gorm.DB) {
	db.AutoMigrate(&bank.Bank{})
	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&productcategory.ProductCategory{})
	db.AutoMigrate(&transaction.Transaction{})
	db.AutoMigrate(&user.User{})
}

func routeUser(db *gorm.DB, v *gin.RouterGroup) {
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := user.NewController(userService)

	v.GET("/users", userController.GetUsers)
	v.GET("/user/:id", userController.GetUser)
	v.POST("/user/register", userController.CreateUser)
	v.PUT("/user/update/:id", userController.UpdateUser)
}

func routeProduct(db *gorm.DB, v *gin.RouterGroup) {
	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productController := product.NewController(productService)

	v.GET("/products", productController.GetProducts)
	v.GET("/product/:id", productController.GetProduct)
	v.POST("/product/create", productController.CreateProduct)
	v.PUT("/product/update/:id", productController.UpdateProduct)
}
