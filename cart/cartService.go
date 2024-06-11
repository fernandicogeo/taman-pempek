package cart

import (
	"errors"

	"gorm.io/gorm"
)

type CartService interface {
	FindAll() ([]Cart, error)
	FindCartByID(ID int) (Cart, error)
	CreateCart(cart CartCreateRequest) (Cart, error)
	UpdateCart(ID int, cart CartUpdateRequest) (Cart, error)
	DeleteCart(ID int) (Cart, error)
}

type service struct {
	cartRepository CartRepository
}

func NewService(cartRepository CartRepository) *service {
	return &service{cartRepository}
}

func (s *service) FindAll() ([]Cart, error) {
	return s.cartRepository.FindAll()
}

func (s *service) FindCartByID(ID int) (Cart, error) {
	return s.cartRepository.FindCartByID(ID)
}

func (s *service) CreateCart(cartRequest CartCreateRequest) (Cart, error) {
	cartData := Cart{
		UserID:     cartRequest.UserID,
		ProductID:  cartRequest.ProductID,
		PaymentID:  cartRequest.PaymentID,
		Quantity:   cartRequest.Quantity,
		TotalPrice: cartRequest.TotalPrice,
		IsActived:  cartRequest.IsActived,
	}

	cart, err := s.cartRepository.CreateCart(cartData)

	return cart, err
}

func (s *service) UpdateCart(ID int, cartRequest CartUpdateRequest) (Cart, error) {
	cart, err := s.cartRepository.FindCartByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Cart{}, errors.New("Cart not found")
	}

	if err != nil {
		return Cart{}, err
	}

	cart.ProductID = cartRequest.ProductID
	cart.PaymentID = cartRequest.PaymentID
	cart.Quantity = cartRequest.Quantity
	cart.TotalPrice = cartRequest.TotalPrice
	cart.IsActived = cartRequest.IsActived

	return s.cartRepository.UpdateCart(cart)
}

func (s *service) DeleteCart(ID int) (Cart, error) {
	cart, err := s.cartRepository.FindCartByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Cart{}, errors.New("Cart not found")
	}

	if err != nil {
		return Cart{}, err
	}

	return s.cartRepository.DeleteCart(cart)
}
