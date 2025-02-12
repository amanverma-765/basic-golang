package routes

import (
	"github.com/amanverma-765/basic-golang/internal/handlers"
	"github.com/amanverma-765/basic-golang/internal/repository"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(app *echo.Echo, userRepo repository.UserRepository) {

	userHandler := handlers.NewUserHandler(userRepo)

	app.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome to the Go Echo Project")
	})

	app.GET("/users", userHandler.GetUsers)
	app.POST("/users", userHandler.CreateUser)
}
