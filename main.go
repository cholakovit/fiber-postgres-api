package main

import (
	"fiber-postgres-api/initializer"
	"fiber-postgres-api/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {	
	app := fiber.New()
	routes.Routes(app)

	initializer.LoadEnv()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}