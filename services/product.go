package services

import (
	"go-git-crud/models"
	"go-git-crud/repositories"
)

func GetProducts() ([]models.Product, error) {
	return repositories.GetProducts()
}
