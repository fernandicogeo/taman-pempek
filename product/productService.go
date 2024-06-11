package product

import (
	"errors"

	"gorm.io/gorm"
)

type ProductService interface {
	FindAll() ([]Product, error)
	FindProductByID(ID int) (Product, error)
	CreateProduct(product ProductCreateRequest) (Product, error)
	UpdateProduct(ID int, product ProductUpdateRequest) (Product, error)
	DeleteProduct(ID int) (Product, error)
}

type service struct {
	productRepository ProductRepository
}

func NewService(productRepository ProductRepository) *service {
	return &service{productRepository}
}

func (s *service) FindAll() ([]Product, error) {
	return s.productRepository.FindAll()
}

func (s *service) FindProductByID(ID int) (Product, error) {
	return s.productRepository.FindProductByID(ID)
}

func (s *service) CreateProduct(productRequest ProductCreateRequest) (Product, error) {
	productData := Product{
		UserID:      productRequest.UserID,
		CategoryID:  productRequest.CategoryID,
		Name:        productRequest.Name,
		Image:       productRequest.Image,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		Stock:       productRequest.Stock,
	}

	product, err := s.productRepository.CreateProduct(productData)

	return product, err
}

func (s *service) UpdateProduct(ID int, productRequest ProductUpdateRequest) (Product, error) {
	product, err := s.productRepository.FindProductByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Product{}, errors.New("Product not found")
	}

	if err != nil {
		return Product{}, err
	}

	product.CategoryID = productRequest.CategoryID
	product.Name = productRequest.Name
	product.Image = productRequest.Image
	product.Description = productRequest.Description
	product.Price = productRequest.Price
	product.Stock = productRequest.Stock

	return s.productRepository.UpdateProduct(product)
}

func (s *service) DeleteProduct(ID int) (Product, error) {
	product, err := s.productRepository.FindProductByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Product{}, errors.New("Product not found")
	}

	if err != nil {
		return Product{}, err
	}

	return s.productRepository.DeleteProduct(product)
}
