package services

import (
	"go-git-crud/models"
	"go-git-crud/repositories"
)

func GetProducts() ([]models.Product, error) {
	products, err := repositories.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetProduct(id string) (models.Product, error) {
	product, err := repositories.GetProduct(id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
