package routes

import (
	"blockchain-event-platform/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Events routes
	app.Get("/events", handlers.GetEvents)

	app.Get("/events/:txHash", handlers.GetEventByTx)

}
