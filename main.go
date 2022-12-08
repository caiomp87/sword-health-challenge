package main

import (
	"log"

	"github.com/caiomp87/sword-health-challenge/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	routes.AddRoutes(app)

	log.Println("API listening on port 3333")
	if err := app.Run(":3333"); err != nil {
		log.Fatalf("could not init server: %v", err)
	}
}
