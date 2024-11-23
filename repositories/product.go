package repositories

import (
	"errors"
	"go-git-crud/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	// find all products and return all data without soft delete
	result := repo.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (repo *ProductRepository) GetProduct(id string) (models.Product, error) {
	var product models.Product

	result := repo.db.First(&product, id)
	if result.Error != nil {
		// return error with text "product not found"
		return models.Product{}, errors.New("product not found")
	}
	return product, nil
}

func (repo *ProductRepository) CreateProduct(product models.Product) (models.Product, error) {
	result := repo.db.Create(&product)

	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return product, nil
}

func (repo *ProductRepository) UpdateProduct(product models.Product, id string) (models.Product, error) {
	var result models.Product

	// Update product without returning the updated product
	// err := repo.db.Model(&result).Where("id = ?", id).Updates(&product)

	// Update product and return the updated product
	err := repo.db.Model(&result).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&product)

	// Handle error
	if err.Error != nil {
		return models.Product{}, err.Error
	}

	return result, nil
}

func (repo *ProductRepository) DeleteProduct(id string) (models.Product, error) {
	var result models.Product

	// Delete product without returning the deleted product
	err := repo.db.Delete(&result, id)

	// Delete product and return the deleted product
	// err := repo.db.Clauses(clause.Returning{}).Delete(&result, id)

	if err.Error != nil {
		return result, err.Error
	}

	return result, nil
}
