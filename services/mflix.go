package services

import (
	"go-git-crud/middleware"
	"go-git-crud/models"
	"go-git-crud/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type MflixService struct {
	repo *repositories.MflixRepository
}

func NewMflixService(db *mongo.Database) *MflixService {
	return &MflixService{repo: repositories.NewMflixRepository(db)}
}

func (service *MflixService) GetMflixs(pagination middleware.PaginationQuery) ([]models.Mflix, error) {
	return service.repo.GetAll(pagination.Skip, pagination.Limit)
}

func (service *MflixService) CreateMflix(mflix models.Mflix) (models.Mflix, error) {
	return service.repo.Create(mflix)
}

func (service *MflixService) UpdateMflix(mflix models.Mflix, id string) (models.Mflix, error) {
	return service.repo.Update(mflix, id)
}

func (service *MflixService) DeleteMflix(id string) (models.Mflix, error) {
	return service.repo.Delete(id)
}
