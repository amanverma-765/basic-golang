package repository

import (
	"context"
	"github.com/amanverma-765/basic-golang/internal/model"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
}
