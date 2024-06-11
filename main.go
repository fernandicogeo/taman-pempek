package main

import (
	"log"
	"taman-pempek/bank"
	"taman-pempek/category"
	"taman-pempek/delivery"
	"taman-pempek/product"
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
	routeBank(db, v1)
	routeCategory(db, v1)
	routeDelivery(db, v1)

	router.Run(":8888") // port
}

func migration(db *gorm.DB) {
	db.AutoMigrate(&bank.Bank{})
	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&category.Category{})
	db.AutoMigrate(&delivery.Delivery{})
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
	v.DELETE("/user/delete/:id", userController.DeleteUser)
}

func routeProduct(db *gorm.DB, v *gin.RouterGroup) {
	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productController := product.NewController(productService)

	v.GET("/products", productController.GetProducts)
	v.GET("/product/:id", productController.GetProduct)
	v.POST("/product/create", productController.CreateProduct)
	v.PUT("/product/update/:id", productController.UpdateProduct)
	v.DELETE("/product/delete/:id", productController.DeleteProduct)
}

func routeBank(db *gorm.DB, v *gin.RouterGroup) {
	bankRepository := bank.NewRepository(db)
	bankService := bank.NewService(bankRepository)
	bankController := bank.NewController(bankService)

	v.GET("/banks", bankController.GetBanks)
	v.GET("/bank/:id", bankController.GetBank)
	v.POST("/bank/create", bankController.CreateBank)
	v.PUT("/bank/update/:id", bankController.UpdateBank)
	v.DELETE("/bank/delete/:id", bankController.DeleteBank)
}

func routeCategory(db *gorm.DB, v *gin.RouterGroup) {
	categoryRepository := category.NewRepository(db)
	categoryService := category.NewService(categoryRepository)
	categoryController := category.NewController(categoryService)

	v.GET("/categories", categoryController.GetCategories)
	v.GET("/category/:id", categoryController.GetCategory)
	v.POST("/category/create", categoryController.CreateCategory)
	v.PUT("/category/update/:id", categoryController.UpdateCategory)
	v.DELETE("/category/delete/:id", categoryController.DeleteCategory)
}

func routeDelivery(db *gorm.DB, v *gin.RouterGroup) {
	deliveryRepository := delivery.NewRepository(db)
	deliveryService := delivery.NewService(deliveryRepository)
	deliveryController := delivery.NewController(deliveryService)

	v.GET("/deliveries", deliveryController.GetDeliveries)
	v.GET("/delivery/:id", deliveryController.GetDelivery)
	v.POST("/delivery/create", deliveryController.CreateDelivery)
	v.PUT("/delivery/update/:id", deliveryController.UpdateDelivery)
	v.DELETE("/delivery/delete/:id", deliveryController.DeleteDelivery)
}
