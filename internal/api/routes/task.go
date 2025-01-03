package routes

import (
	"gera-ai/internal/api/handlers"
	"gera-ai/internal/api/middlewares"
	"gera-ai/internal/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TaskRouter(app fiber.Router, db *gorm.DB) {
	jwt := middlewares.AuthMiddleware(config.Config.JWTSecret)
	app.Post("/task/new", jwt, handlers.CreateTask(db))
	app.Get("/task/get/:id", jwt, handlers.GetTask(db))
	app.Put("/task/edit", jwt, handlers.EditTask(db))
	app.Delete("/task/delete", jwt, handlers.DeleteTask(db))

	app.Get("/task/all", jwt, handlers.GetAllTasks(db))
}
