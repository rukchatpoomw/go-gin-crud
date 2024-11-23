package services

import (
	"go-git-crud/models"
	"go-git-crud/repositories"

	"gorm.io/gorm"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{repo: repositories.NewProductRepository(db)}
}

// Get all products
func (service *ProductService) GetProducts() ([]models.Product, error) {
	products, err := service.repo.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Get product by id
func (service *ProductService) GetProduct(id string) (models.Product, error) {
	product, err := service.repo.GetProduct(id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// Create product
func (service *ProductService) CreateProduct(product models.Product) (models.Product, error) {
	product, err := service.repo.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// Update product
func (service *ProductService) UpdateProduct(product models.Product, id string) (models.Product, error) {
	products, err := service.repo.UpdateProduct(product, id)
	if err != nil {
		return models.Product{}, err
	}
	return products, nil
}

// Delete product
func (service *ProductService) DeleteProduct(id string) (models.Product, error) {
	products, err := service.repo.DeleteProduct(id)
	if err != nil {
		return models.Product{}, err
	}
	return products, nil
}
