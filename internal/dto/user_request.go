package dto

import "github.com/amanverma-765/basic-golang/internal/model"

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request UserRequest) ToDomain() model.User {
	return model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}
