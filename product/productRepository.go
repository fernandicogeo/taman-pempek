package product

import (
	"errors"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]Product, error)
	FindProductByID(ID int) (Product, error)
	CreateProduct(product Product) (Product, error)
	UpdateProduct(product Product) (Product, error)
	DeleteProduct(product Product) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *repository) FindProductByID(ID int) (Product, error) {
	var product Product
	err := r.db.First(&product, ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Product{}, errors.New("Product not found")
	}
	return product, err
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateProduct(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *repository) UpdateProduct(product Product) (Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *repository) DeleteProduct(product Product) (Product, error) {
	err := r.db.Delete(&product).Error
	return product, err
}
