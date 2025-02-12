package repository

import (
	"context"
	"github.com/amanverma-765/basic-golang/internal/entity"
	"github.com/amanverma-765/basic-golang/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
	"time"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) UserRepository {
	return &MongoUserRepository{
		collection: db.Collection("users"),
	}
}

func (repo *MongoUserRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	var results []entity.UserEntity
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Printf("error closing cursor: %v", err)
		}
	}(cursor, ctx)

	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	var users []model.User
	for _, userEntity := range results {
		users = append(users, userEntity.ToDomain())
	}
	return users, nil
}

func (repo *MongoUserRepository) CreateUser(ctx context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	userEntity := entity.FromDomain(*user)
	_, err := repo.collection.InsertOne(ctx, userEntity)
	return err
}
