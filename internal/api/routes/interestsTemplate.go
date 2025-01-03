package routes

import (
	"gera-ai/internal/api/handlers"
	"gera-ai/internal/api/middlewares"
	"gera-ai/internal/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InterestsTemplateRouter(app fiber.Router, db *gorm.DB) {
	jwt := middlewares.AuthMiddleware(config.Config.JWTSecret)
	app.Post("/template/interests/new", jwt, handlers.CreateInterestsTemplate(db))
	app.Get("/template/interests/get/:id", jwt, handlers.GetInterestsTemplate(db))
	app.Put("/template/interests/edit", jwt, handlers.EditInterestsTemplate(db))
	app.Delete("/template/interests/delete", jwt, handlers.DeleteInterestsTemplate(db))

	app.Get("/template/interests/all", jwt, handlers.GetAllInterestsTemplates(db))
}
