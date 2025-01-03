package routes

import (
	"gera-ai/internal/api/handlers"
	"gera-ai/internal/api/middlewares"
	"gera-ai/internal/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ConditionTemplateRouter(app fiber.Router, db *gorm.DB) {
	jwt := middlewares.AuthMiddleware(config.Config.JWTSecret)
	app.Post("/template/condition/new", jwt, handlers.CreateConditionTemplate(db))
	app.Get("/template/condition/get/:id", jwt, handlers.GetConditionTemplate(db))
	app.Put("/template/condition/edit", jwt, handlers.EditConditionTemplate(db))
	app.Delete("/template/condition/delete", jwt, handlers.DeleteConditionTemplate(db))

	app.Get("/template/condition/all", jwt, handlers.GetAllConditionTemplates(db))
}
