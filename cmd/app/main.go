package main

import (
	"context"
	"fmt"
	"github.com/amanverma-765/basic-golang/configs"
	"github.com/amanverma-765/basic-golang/internal/database"
	"github.com/amanverma-765/basic-golang/internal/repository"
	"github.com/amanverma-765/basic-golang/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {

	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	mongoClient, err := database.InitMongoDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()
	mongoDB := mongoClient.Database(cfg.Database)
	userRepo := repository.NewMongoUserRepository(mongoDB)

	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	routes.RegisterRoutes(app, userRepo)

	addr := fmt.Sprintf(":%s", cfg.Port)
	if err := app.Start(addr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
