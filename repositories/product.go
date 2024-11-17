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
