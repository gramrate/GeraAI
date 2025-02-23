package app

import (
	"gera-ai/internal/api/routes"
	"gera-ai/internal/config"
	dbmodels "gera-ai/internal/models/database"
	"gera-ai/internal/utils/database"
	"gera-ai/internal/utils/taskGenerator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
	"log"
)

type GeraApp struct {
	Fiber *fiber.App
	Db    *gorm.DB
}

func NewGeraApp() *GeraApp {
	// init config
	config.InitConfig()
	// init db connection
	db, dbErr := database.Connection()
	if dbErr != nil {
		log.Fatalf("failed to connect database: %v", dbErr.Error())
	}

	// migrate db models
	migrateErr := db.AutoMigrate(
		dbmodels.User{},
		dbmodels.Task{},
		dbmodels.InterestsTemplate{},
		dbmodels.ConditionTemplate{},
		dbmodels.GenerationByInterestsHistory{},
		dbmodels.GenerationByNoInterestsHistory{},
		dbmodels.GenerationAnswersHistory{},
	)
	if migrateErr != nil {
		log.Fatalf("failed to migrate database: %v", migrateErr.Error())
	}

	tg, err := taskGenerator.NewTaskGenerator(config.Config.ApiKey, config.Config.ProxyURL)
	if err != nil {
		log.Fatalf("Failed to create task generator: %v", err)
	}
	tg = tg
	// init new fiber app and use swagger
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	// map api routes and swagger
	api := app.Group("/api")
	app.Get("/swagger/*", swagger.HandlerDefault)
	routes.SwaggerRouter(api)
	routes.PingRouter(api)
	routes.AuthRouter(api, db)
	routes.ConditionTemplateRouter(api, db)
	routes.InterestsTemplateRouter(api, db)
	routes.TaskRouter(api, db)
	routes.AIGeneratorRouter(api, db, tg)
	return &GeraApp{
		Fiber: app,
		Db:    db,
	}
}

func Start(app *GeraApp) {
	if err := app.Fiber.Listen(":8080"); err != nil {
		panic("failed to listen: " + err.Error())
	}
}

//TODO реализовать авторизацию через яндекс ID
// TODO Реализовать шаблоны вариантов
// TODO проверить соответсвие валидации и бд
// TODO реализовать получение истории генераций
// TODO Убрать все коменты
// TODO почему-то бд спамит что-то про root
// TODO переписать эндпоинты в сваггере
// TODO протестить генерацию
