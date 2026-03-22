package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"blockchain-event-platform/cache"
	"blockchain-event-platform/database"
	"blockchain-event-platform/routes"
)

func main() {

	database.ConnectDB()
	cache.ConnectRedis()

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "API running",
		})
	})

	routes.SetupRoutes(app)

	log.Println("Server running on port 3000")

	app.Listen(":3000")
}
