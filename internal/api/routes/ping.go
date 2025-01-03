package routes

import (
	"gera-ai/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func PingRouter(app fiber.Router) {
	app.Get("/ping", handlers.Ping())
}
