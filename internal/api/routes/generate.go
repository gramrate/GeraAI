package routes

import (
	"gera-ai/internal/api/handlers"
	"gera-ai/internal/api/middlewares"
	"gera-ai/internal/config"
	"gera-ai/internal/utils/taskGenerator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AIGeneratorRouter(app fiber.Router, db *gorm.DB, tg *taskGenerator.TaskGenerator) {
	jwt := middlewares.AuthMiddleware(config.Config.JWTSecret)

	app.Post("/generate/interests", jwt, handlers.GenerateTaskByInterest(db, tg))
	app.Post("/generate/nointerests", jwt, handlers.GenerateTaskByNoInterest(db, tg))
	app.Post("/generate/answer", jwt, handlers.GenerateAnswer(db, tg))
}
