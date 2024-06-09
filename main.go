package main

import (
	"log"
	"os/user"
	"taman-pempek/bank"
	"taman-pempek/product"
	"taman-pempek/productcategory"
	"taman-pempek/transaction"
	"taman-pempek/transactiondetail"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/taman-pempek?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection error")
	}

	db.AutoMigrate(&bank.Bank{})
	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&productcategory.ProductCategory{})
	db.AutoMigrate(&transaction.Transaction{})
	db.AutoMigrate(&transactiondetail.TransactionDetail{})
	db.AutoMigrate(&user.User{})

}
