package cart

import (
	"errors"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindAll() ([]Cart, error)
	FindCartByID(ID int) (Cart, error)
	CreateCart(cart Cart) (Cart, error)
	UpdateCart(cart Cart) (Cart, error)
	DeleteCart(cart Cart) (Cart, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) FindAll() ([]Cart, error) {
	var carts []Cart
	err := r.db.Find(&carts).Error
	return carts, err
}

func (r *repository) FindCartByID(ID int) (Cart, error) {
	var cart Cart
	err := r.db.First(&cart, ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Cart{}, errors.New("Cart not found")
	}
	return cart, err
}
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCart(cart Cart) (Cart, error) {
	err := r.db.Create(&cart).Error
	return cart, err
}

func (r *repository) UpdateCart(cart Cart) (Cart, error) {
	err := r.db.Save(&cart).Error
	return cart, err
}

func (r *repository) DeleteCart(cart Cart) (Cart, error) {
	err := r.db.Delete(&cart).Error
	return cart, err
}
