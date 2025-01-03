package handlers

import (
	"encoding/json"
	dbmodels "gera-ai/internal/models/database"
	"gera-ai/internal/models/requests"
	"gera-ai/internal/models/responses"
	jwtUtils "gera-ai/internal/utils/jwtUtils"
	"gera-ai/internal/utils/taskGenerator"
	"gera-ai/internal/utils/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

// GenerateTaskByInterest generates a task based on a list of interests
// @Summary Generate Task by Interests
// @Description Generates a task based on the provided list of interests and saves it in the history.
// @Tags Task Generation
// @Accept json
// @Produce json
// @Param input body requests.GenerateByInterests true "Data for task generation"
// @Success 200 {object} responses.GeneratedTaskResponse "Successfully generated task"
// @Failure 400 {object} responses.ErrorResponse "Invalid token or JSON"
// @Failure 422 {object} responses.ValidationErrorResponse "Validation error"
// @Failure 500 {object} responses.ErrorResponse "Error saving the generated task"
// @Router /generate/interest [post]
func GenerateTaskByInterest(db *gorm.DB, tg *taskGenerator.TaskGenerator) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Извлекаем информацию о пользователе из JWT
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		// Парсим входящие данные
		data := requests.GenerateByInterests{}
		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid json",
				Error:  err.Error(),
			})
		}

		// Валидация данных
		validationErrors := validator.ValidateStruct(data)
		if validationErrors != nil {
			return c.Status(422).JSON(responses.ValidationErrorResponse{
				Status: "validation failed",
				Errors: validationErrors,
			})
		}

		// Генерация задания
		taskText, err := tg.GenerateTaskWithInterests(data.Condition, data.Interests)
		if err != nil {
			return c.Status(422).JSON(responses.ErrorResponse{
				Status: "generation failed",
				Error:  err.Error(),
			})
		}

		// Преобразование Interests в JSON
		interestsJSON, err := json.Marshal(data.Interests)
		if err != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to process interests",
				Error:  err.Error(),
			})
		}

		// Сохранение в истории генераций
		generatedTask := dbmodels.GenerationByInterestsHistory{
			UserID:    authorID,
			Condition: data.Condition,
			Interests: interestsJSON,
			TaskText:  taskText,
			CreatedAt: time.Now(),
		}

		if err := db.Create(&generatedTask).Error; err != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to save generated task",
				Error:  err.Error(),
			})
		}

		// Ответ с успешным результатом
		return c.Status(200).JSON(responses.GeneratedTaskResponse{
			Status:        "generated successfully",
			GeneratedText: taskText,
		})
	}
}

// GenerateTaskByNoInterest generates a task based on reality without considering interests
// @Summary Generate Task Without Interests
// @Description Generates a task based solely on the provided condition and saves it in history.
// @Tags Task Generation
// @Accept json
// @Produce json
// @Param input body requests.GenerateByNoInterests true "Data for task generation"
// @Success 200 {object} responses.GeneratedTaskResponse "Successfully generated task"
// @Failure 400 {object} responses.ErrorResponse "Invalid token or JSON"
// @Failure 422 {object} responses.ValidationErrorResponse "Validation error"
// @Failure 500 {object} responses.ErrorResponse "Error saving the generated task"
// @Router /generate/no-interest [post]
func GenerateTaskByNoInterest(db *gorm.DB, tg *taskGenerator.TaskGenerator) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Извлекаем информацию о пользователе из JWT
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		// Парсим входящие данные
		data := requests.GenerateByNoInterests{}
		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid json",
				Error:  err.Error(),
			})
		}

		// Валидация данных
		validationErrors := validator.ValidateStruct(data)
		if validationErrors != nil {
			return c.Status(422).JSON(responses.ValidationErrorResponse{
				Status: "validation failed",
				Errors: validationErrors,
			})
		}

		// Генерация задания
		taskText, err := tg.GenerateTaskWithNoInterests(data.Condition)
		if err != nil {
			return c.Status(422).JSON(responses.ErrorResponse{
				Status: "generation failed",
				Error:  err.Error(),
			})
		}
		// Сохранение в истории генераций
		generatedTask := dbmodels.GenerationByNoInterestsHistory{
			UserID:    authorID,
			Condition: data.Condition,
			TaskText:  taskText,
			CreatedAt: time.Now(),
		}

		if err := db.Create(&generatedTask).Error; err != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to save generated task",
				Error:  err.Error(),
			})
		}

		// Ответ с успешным результатом
		return c.Status(200).JSON(responses.GeneratedTaskResponse{
			Status:        "generated successfully",
			GeneratedText: taskText,
		})
	}
}

// GenerateAnswerByCondition generates an answer based on a condition
// @Summary Generate Answer by Condition
// @Description Generates an answer based on the provided condition and saves it in history.
// @Tags Answer Generation
// @Accept json
// @Produce json
// @Param input body requests.GenerateAnswer true "Data for answer generation"
// @Success 200 {object} responses.GeneratedAnswerResponse "Successfully generated answer"
// @Failure 400 {object} responses.ErrorResponse "Invalid token or JSON"
// @Failure 422 {object} responses.ValidationErrorResponse "Validation error"
// @Failure 500 {object} responses.ErrorResponse "Error saving the generated answer"
// @Router /generate/answer [post]
func GenerateAnswer(db *gorm.DB, tg *taskGenerator.TaskGenerator) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Извлекаем информацию о пользователе из JWT
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		// Парсим входящие данные
		data := requests.GenerateAnswer{}
		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid json",
				Error:  err.Error(),
			})
		}

		// Валидация данных
		validationErrors := validator.ValidateStruct(data)
		if validationErrors != nil {
			return c.Status(422).JSON(responses.ValidationErrorResponse{
				Status: "validation failed",
				Errors: validationErrors,
			})
		}

		// Генерация задания
		answer, err := tg.GenerateAnswer(data.Condition)
		if err != nil {
			return c.Status(422).JSON(responses.ErrorResponse{
				Status: "generation failed",
				Error:  err.Error(),
			})
		}
		// Сохранение в истории генераций
		generatedAnswer := dbmodels.GenerationAnswersHistory{
			UserID:    authorID,
			Condition: data.Condition,
			Answer:    answer,
			CreatedAt: time.Now(),
		}

		if err := db.Create(&generatedAnswer).Error; err != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to save generated task",
				Error:  err.Error(),
			})
		}

		// Ответ с успешным результатом
		return c.Status(200).JSON(responses.GeneratedAnswerResponse{
			Status:        "generated successfully",
			GeneratedText: answer,
		})
	}
}
