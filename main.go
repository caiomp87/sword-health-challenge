package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/caiomp87/sword-health-challenge/cache"
	"github.com/caiomp87/sword-health-challenge/db"
	"github.com/caiomp87/sword-health-challenge/middlewares"
	"github.com/caiomp87/sword-health-challenge/repository"
	"github.com/caiomp87/sword-health-challenge/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseService := db.NewDatabaseService("mysql", os.Getenv("DB_CONNECTION_STRING"))

	database, err := databaseService.Connect()
	if err != nil {
		log.Fatalf("could not connect on database: %v", err)
	}

	defer func() {
		if err := databaseService.Disconnect(database); err != nil {
			log.Fatalf("could not disconnect from database gracefully: %v", err)
		}
	}()

	repository.TaskRepository = repository.NewTaskRepository(database)
	repository.UserRepository = repository.NewUserRepository(database)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	cache.CacheService = cache.NewCache()
	if err := cache.CacheService.Ping(ctx); err != nil {
		log.Fatalf("could not initialize cache: %v", err)
	}

	env, ok := os.LookupEnv("ENV")
	if env == "production" || !ok {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()
	app.Use(middlewares.Authenticate())
	routes.AddRoutes(app)

	log.Println("API listening on port 3333")
	if err := app.Run(":3333"); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
