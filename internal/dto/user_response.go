package dto

import "github.com/amanverma-765/basic-golang/internal/model"

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	//JWTToken string `json:"jwt_token"`
}

func FromDomain(user model.User) UserResponse {
	return UserResponse{
		Name:  user.Name,
		Email: user.Email,
		//JWTToken:
	}
}
