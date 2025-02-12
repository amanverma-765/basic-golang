package handlers

import (
	"github.com/amanverma-765/basic-golang/internal/dto"
	"github.com/amanverma-765/basic-golang/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	Repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (handler *UserHandler) GetUsers(app echo.Context) error {
	ctx := app.Request().Context()
	users, err := handler.Repo.GetUsers(ctx)
	if err != nil {
		return app.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var response []dto.UserResponse
	for _, user := range users {
		response = append(response, dto.FromDomain(user))
	}
	return app.JSON(http.StatusOK, response)
}

func (handler *UserHandler) CreateUser(app echo.Context) error {
	ctx := app.Request().Context()
	userReq := new(dto.UserRequest)
	if err := app.Bind(userReq); err != nil {
		return app.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user := userReq.ToDomain()

	if err := handler.Repo.CreateUser(ctx, &user); err != nil {
		return app.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return app.JSON(http.StatusCreated, userReq)
}
