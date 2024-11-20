package repositories

import (
	"errors"
	"go-git-crud/models"

	"gorm.io/gorm"
)

type ProductRepositoryType struct {
	db *gorm.DB
}

var products []models.Product

func ProductRepository(db *gorm.DB) *ProductRepositoryType {
	return &ProductRepositoryType{db: db}
}

func (repo *ProductRepositoryType) GetProducts() ([]models.Product, error) {
	var products []models.Product
	result := repo.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (repo *ProductRepositoryType) GetProduct(id string) (models.Product, error) {
	var product models.Product
	result := repo.db.First(&product, id)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	return product, nil
}

func CreateProduct(product models.Product) (models.Product, error) {

	for _, p := range products {
		if p.ID == product.ID {
			return models.Product{}, errors.New("product with the same ID already exists")
		}
	}

	products = append(products, product)
	return product, nil
}

func UpdateProduct(product models.Product, id any) ([]models.Product, error) {

	for index, value := range products {
		if value.ID == id {
			products[index] = product
			return products, nil
		}
	}
	return nil, errors.New("product not found")
}

func DeleteProduct(id any) ([]models.Product, error) {
	for index, value := range products {
		if value.ID == id {
			products = append(products[:index], products[index+1:]...)
			return products, nil
		}
	}
	return nil, errors.New("product not found")
}
