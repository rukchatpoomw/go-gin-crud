package repositories

import (
	"context"
	"go-git-crud/config"
	"go-git-crud/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MflixRepository struct {
	collection *mongo.Collection
}

func NewMflixRepository(db *mongo.Database) *MflixRepository {
	return &MflixRepository{
		collection: db.Collection(config.GetEnv("MONGODB_COLLECTION_COMMENTS")),
	}
}

func (repo *MflixRepository) GetAll(skip int64, limit int64) ([]models.Mflix, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var mflixs []models.Mflix

	cursor, err := repo.collection.Find(ctx, bson.M{}, options.Find().SetSkip(skip).SetLimit(limit))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &mflixs); err != nil {
		return nil, err
	}

	return mflixs, nil
}

func (repo *MflixRepository) Create(mflix models.Mflix) (models.Mflix, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mflix.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	mflix.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	result, err := repo.collection.InsertOne(ctx, mflix)
	if err != nil {
		return models.Mflix{}, err
	}

	mflix.ID = result.InsertedID.(primitive.ObjectID)
	return mflix, nil
}

func (repo *MflixRepository) Update(mflix models.Mflix, id string) (models.Mflix, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Mflix{}, err
	}

	mflix.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	update := bson.M{
		"$set": bson.M{
			"title":       mflix.Title,
			"year":        mflix.Year,
			"director":    mflix.Director,
			"description": mflix.Description,
			"updated_at":  mflix.UpdatedAt,
		},
	}

	result := repo.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID},
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	var updatedMflix models.Mflix
	if err := result.Decode(&updatedMflix); err != nil {
		return models.Mflix{}, err
	}

	return updatedMflix, nil
}

func (repo *MflixRepository) Delete(id string) (models.Mflix, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Mflix{}, err
	}

	var deletedMflix models.Mflix
	err = repo.collection.FindOneAndDelete(
		ctx,
		bson.M{"_id": objectID},
	).Decode(&deletedMflix)

	if err != nil {
		return models.Mflix{}, err
	}

	return deletedMflix, nil
}
