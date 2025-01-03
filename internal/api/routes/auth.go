package routes

import (
	"gera-ai/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRouter(app fiber.Router, db *gorm.DB) {
	app.Post("/auth/register", handlers.Register(db))

	app.Post("/auth/login", handlers.Login(db))
}
