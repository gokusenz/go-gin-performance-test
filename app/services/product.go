package services

import (
	"fastwork/go-gin-performance-test/app/models"
	"fastwork/go-gin-performance-test/app/repos"
)

// ProductService model
type ProductService struct {
	ProductRepo repos.IProductRepo
}

// IProductService interface
type IProductService interface {
	GetbyID(string) (*models.Product, error)
}

// NewProductService function to create new product service
func NewProductService(
	productRepo repos.IProductRepo) *ProductService {
	return &ProductService{
		ProductRepo: productRepo,
	}
}

// GetbyID function of product service
func (ps *ProductService) GetbyID(userID string) (*models.Product, error) {
	productRepo := ps.ProductRepo

	// Get product
	product, err := productRepo.GetbyID(userID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
