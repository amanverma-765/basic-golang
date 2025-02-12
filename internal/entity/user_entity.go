package entity

import (
	"github.com/amanverma-765/basic-golang/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserEntity struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
}

func (entity UserEntity) ToDomain() model.User {
	return model.User{
		ID:       entity.ID.Hex(),
		Name:     entity.Name,
		Email:    entity.Email,
		Password: entity.Password,
	}
}

func FromDomain(user model.User) UserEntity {
	return UserEntity{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
