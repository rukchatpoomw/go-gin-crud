package services

import (
	"go-git-crud/models"
	"go-git-crud/repositories"
)

type ProductServiceType struct {
	repo *repositories.ProductRepositoryType
}

func ProductService(repo *repositories.ProductRepositoryType) *ProductServiceType {
	return &ProductServiceType{repo: repo}
}

// Get all products
func (service *ProductServiceType) GetProducts() ([]models.Product, error) {
	products, err := service.repo.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Get product by id
func (service *ProductServiceType) GetProduct(id string) (models.Product, error) {
	product, err := service.repo.GetProduct(id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// Create product
func (service *ProductServiceType) CreateProduct(product models.Product) (models.Product, error) {
	product, err := service.repo.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// Update product
func (service *ProductServiceType) UpdateProduct(product models.Product, id string) (models.Product, error) {
	products, err := service.repo.UpdateProduct(product, id)
	if err != nil {
		return models.Product{}, err
	}
	return products, nil
}

// Delete product
func (service *ProductServiceType) DeleteProduct(id string) (models.Product, error) {
	products, err := service.repo.DeleteProduct(id)
	if err != nil {
		return models.Product{}, err
	}
	return products, nil
}
