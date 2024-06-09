package main

import (
	"log"
	"taman-pempek/bank"
	"taman-pempek/product"
	"taman-pempek/productcategory"
	"taman-pempek/transaction"
	"taman-pempek/transactiondetail"
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

	migration(db)

	routeUser(db)
}

func migration(db *gorm.DB) {
	db.AutoMigrate(&bank.Bank{})
	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&productcategory.ProductCategory{})
	db.AutoMigrate(&transaction.Transaction{})
	db.AutoMigrate(&transactiondetail.TransactionDetail{})
	db.AutoMigrate(&user.User{})
}

func routeUser(db *gorm.DB) {
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := user.NewController(userService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/users", userController.GetUsers)
	v1.GET("/user/:id", userController.GetUser)
	v1.POST("/user/register", userController.CreateUser)
	v1.PUT("/user/update/:id", userController.UpdateUser)

	router.Run(":8888") // port
}
