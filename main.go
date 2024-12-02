package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"delivery/config"
	"delivery/database/postgres"
	"delivery/routes"
)

func main() {
	r := gin.Default()
	cfg := config.LoadConfig()


	err := postgres.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer postgres.DB.Close()

	r.LoadHTMLGlob("templates/*.html")


	routes.InitializeRoutes(r)


	r.Run(":4040")
}
