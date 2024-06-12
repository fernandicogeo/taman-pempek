package main

import (
	"log"
	"taman-pempek/bank"
	"taman-pempek/cart"
	"taman-pempek/category"
	"taman-pempek/delivery"
	"taman-pempek/middleware"
	"taman-pempek/payment"
	"taman-pempek/product"
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

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userMiddleware := middleware.NewMiddleware(userService)

	requireAuth := userMiddleware.RequireAuth

	routeUser(db, v1, requireAuth)
	routeProduct(db, v1, requireAuth)
	routeBank(db, v1, requireAuth)
	routeCategory(db, v1, requireAuth)
	routeDelivery(db, v1, requireAuth)
	routeCart(db, v1, requireAuth)
	routePayment(db, v1, requireAuth)

	router.Run(":8888") // port
}

func migration(db *gorm.DB) {
	db.AutoMigrate(&bank.Bank{})
	db.AutoMigrate(&cart.Cart{})
	db.AutoMigrate(&category.Category{})
	db.AutoMigrate(&delivery.Delivery{})
	db.AutoMigrate(&payment.Payment{})
	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&user.User{})
}

func routeUser(db *gorm.DB, v *gin.RouterGroup, requireAuth func(c *gin.Context)) {
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := user.NewController(userService)

	v.GET("/users", requireAuth, userController.GetUsers)
	v.GET("/user/:id", requireAuth, userController.GetUser)
	v.POST("/user/register", requireAuth, userController.CreateUser)
	v.PUT("/user/update/:id", requireAuth, userController.UpdateUser)
	v.DELETE("/user/delete/:id", requireAuth, userController.DeleteUser)

	v.POST("/login", userController.Login)
}

func routeProduct(db *gorm.DB, v *gin.RouterGroup, requireAuth func(c *gin.Context)) {
	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productController := product.NewController(productService)

	v.GET("/products", requireAuth, productController.GetProducts)
	v.GET("/product/:id", requireAuth, productController.GetProduct)
	v.POST("/product/create", requireAuth, productController.CreateProduct)
	v.PUT("/product/update/:id", requireAuth, productController.UpdateProduct)
	v.DELETE("/product/delete/:id", requireAuth, productController.DeleteProduct)
}

func routeBank(db *gorm.DB, v *gin.RouterGroup, requireAuth func(c *gin.Context)) {
	bankRepository := bank.NewRepository(db)
	bankService := bank.NewService(bankRepository)
	bankController := bank.NewController(bankService)

	v.GET("/banks", requireAuth, bankController.GetBanks)
	v.GET("/bank/:id", requireAuth, bankController.GetBank)
	v.POST("/bank/create", requireAuth, bankController.CreateBank)
	v.PUT("/bank/update/:id", requireAuth, bankController.UpdateBank)
	v.DELETE("/bank/delete/:id", requireAuth, bankController.DeleteBank)
}

func routeCategory(db *gorm.DB, v *gin.RouterGroup, requireAuth func(c *gin.Context)) {
	categoryRepository := category.NewRepository(db)
	categoryService := category.NewService(categoryRepository)
	categoryController := category.NewController(categoryService)

	v.GET("/categories", requireAuth, categoryController.GetCategories)
	v.GET("/category/:id", requireAuth, categoryController.GetCategory)
	v.POST("/category/create", requireAuth, categoryController.CreateCategory)
	v.PUT("/category/update/:id", requireAuth, categoryController.UpdateCategory)
	v.DELETE("/category/delete/:id", requireAuth, categoryController.DeleteCategory)
}

func routeDelivery(db *gorm.DB, v *gin.RouterGroup, requireAuth func(c *gin.Context)) {
	deliveryRepository := delivery.NewRepository(db)
	deliveryService := delivery.NewService(deliveryRepository)
	deliveryController := delivery.NewController(deliveryService)

	v.GET("/deliveries", requireAuth, deliveryController.GetDeliveries)
	v.GET("/delivery/:id", requireAuth, deliveryController.GetDelivery)
	v.POST("/delivery/create", requireAuth, deliveryController.CreateDelivery)
	v.PUT("/delivery/update/:id", requireAuth, deliveryController.UpdateDelivery)
	v.DELETE("/delivery/delete/:id", requireAuth, deliveryController.DeleteDelivery)
}

func routeCart(db *gorm.DB, v *gin.RouterGroup, requireAuth func(c *gin.Context)) {
	cartRepository := cart.NewRepository(db)
	cartService := cart.NewService(cartRepository)
	cartController := cart.NewController(cartService)

	v.GET("/carts", requireAuth, cartController.GetCarts)
	v.GET("/cart/:id", requireAuth, cartController.GetCart)
	v.POST("/cart/create", requireAuth, cartController.CreateCart)
	v.PUT("/cart/update/:id", requireAuth, cartController.UpdateCart)
	v.DELETE("/cart/delete/:id", requireAuth, cartController.DeleteCart)
}

func routePayment(db *gorm.DB, v *gin.RouterGroup, requireAuth func(c *gin.Context)) {
	paymentRepository := payment.NewRepository(db)
	paymentService := payment.NewService(paymentRepository)
	paymentController := payment.NewController(paymentService)

	v.GET("/payments", requireAuth, paymentController.GetPayments)
	v.GET("/payment/:id", requireAuth, paymentController.GetPayment)
	v.POST("/payment/create", requireAuth, paymentController.CreatePayment)
	v.PUT("/payment/update/:id", requireAuth, paymentController.UpdatePayment)
	v.DELETE("/payment/delete/:id", requireAuth, paymentController.DeletePayment)
}
