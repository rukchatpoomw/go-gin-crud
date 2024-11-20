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
func CreateProduct(product models.Product) (models.Product, error) {
	product, err := repositories.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// Update product
func UpdateProduct(product models.Product, id string) ([]models.Product, error) {
	products, err := repositories.UpdateProduct(product, id)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Delete product
func DeleteProduct(id string) ([]models.Product, error) {
	products, err := repositories.DeleteProduct(id)
	if err != nil {
		return nil, err
	}
	return products, nil
}
