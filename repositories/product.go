package repositories

import (
	"errors"
	"go-git-crud/models"
)

var products = []models.Product{
	{ID: "1", Name: "Product 1", Price: 10000},
	{ID: "2", Name: "Product 2", Price: 20000},
	{ID: "3", Name: "Product 3", Price: 30000},
}

func GetProducts() ([]models.Product, error) {
	return products, nil
}

func GetProduct(id string) (models.Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}

	// if result is nil, return error
	return models.Product{}, errors.New("product not found")
}

func CreateProduct(product models.Product) (models.Product, error) {
	// Simulating unique ID check
	for _, p := range products {
		if p.ID == product.ID {
			return models.Product{}, errors.New("product with the same ID already exists")
		}
	}

	products = append(products, product)
	return product, nil
}
