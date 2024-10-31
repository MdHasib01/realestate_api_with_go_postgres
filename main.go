package main

import (
	"log"

	"github.com/MdHasib01/realestate_api_with_go_postgre/config"
	"github.com/MdHasib01/realestate_api_with_go_postgre/database"
	"github.com/MdHasib01/realestate_api_with_go_postgre/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Load configuration
	config.LoadEnv()

	// Connect to the database
	database.Connect()

	// Setup routes
	routes.SetupRoutes(app)

	port := config.GetPort()
	log.Fatal(app.Listen(":" + port))
}
