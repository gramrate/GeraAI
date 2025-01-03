package handlers

import (
	"gera-ai/internal/models/database"
	"gera-ai/internal/models/requests"
	"gera-ai/internal/models/responses"
	"gera-ai/internal/utils/jwtUtils"
	"gera-ai/internal/utils/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// CreateConditionTemplate godoc
// @Summary Create a new condition template
// @Description Creates a condition template for the current user
// @Tags ConditionTemplate
// @Accept json
// @Produce json
// @Param data body requests.CreateConditionTemplate true "Condition template data"
// @Success 200 {object} responses.CreateConditionTemplateDTO
// @Failure 400 {object} responses.ErrorResponse "Invalid token or request"
// @Failure 422 {object} responses.ValidationErrorResponse "Validation errors"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /api/condition-templates [post]
func CreateConditionTemplate(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		data := requests.CreateConditionTemplate{}
		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid json",
				Error:  err.Error(),
			})
		}

		validationErrors := validator.ValidateStruct(data)
		if validationErrors != nil {
			return c.Status(422).JSON(responses.ValidationErrorResponse{
				Status: "validation failed",
				Errors: validationErrors,
			})
		}

		conditionTemplate := database.ConditionTemplate{
			AuthorID:  authorID,
			Title:     data.Title,
			Condition: data.Condition,
			CreatedAt: time.Now(),
		}

		result := db.Create(&conditionTemplate)
		if result.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to create condition template",
				Error:  result.Error.Error(),
			})
		}

		return c.Status(200).JSON(responses.CreateConditionTemplateDTO{
			Status: "condition template created",
			TaskTemplate: responses.ConditionTemplateDTO{
				ID:        conditionTemplate.ID,
				Title:     conditionTemplate.Title,
				Condition: conditionTemplate.Condition,
			},
		})
	}
}

// GetConditionTemplate godoc
// @Summary Get a condition template by ID
// @Description Retrieves a specific condition template owned by the user
// @Tags ConditionTemplate
// @Param id path int true "Condition Template ID"
// @Success 200 {object} responses.GetConditionTemplateDTO
// @Failure 400 {object} responses.ErrorResponse "Invalid token or request"
// @Failure 403 {object} responses.ErrorResponse "Forbidden"
// @Failure 404 {object} responses.ErrorResponse "Not found"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /api/condition-templates/{id} [get]
func GetConditionTemplate(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		conditionTemplateID, err := strconv.Atoi(c.Params("id"))
		if err != nil || conditionTemplateID <= 0 {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid template ID",
				Error:  "template ID must be a positive integer",
			})
		}

		var conditionTemplate database.ConditionTemplate
		result := db.First(&conditionTemplate, conditionTemplateID)
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(responses.ErrorResponse{
				Status: "condition template not found",
			})
		} else if result.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "internal server error",
				Error:  result.Error.Error(),
			})
		}

		if conditionTemplate.AuthorID != authorID {
			return c.Status(403).JSON(responses.ErrorResponse{
				Status: "forbidden",
				Error:  "you are not the author of this condition template",
			})
		}

		return c.Status(200).JSON(responses.GetConditionTemplateDTO{
			TaskTemplate: responses.ConditionTemplateDTO{
				ID:        conditionTemplate.ID,
				Title:     conditionTemplate.Title,
				Condition: conditionTemplate.Condition,
			},
		})
	}
}

// EditConditionTemplate godoc
// @Summary Edit an existing condition template
// @Description Updates a condition template owned by the user
// @Tags ConditionTemplate
// @Accept json
// @Produce json
// @Param data body requests.EditConditionTemplate true "Condition template data"
// @Success 200 {object} responses.EditConditionTemplateDTO
// @Failure 400 {object} responses.ErrorResponse "Invalid token or request"
// @Failure 403 {object} responses.ErrorResponse "Forbidden"
// @Failure 404 {object} responses.ErrorResponse "Not found"
// @Failure 422 {object} responses.ValidationErrorResponse "Validation errors"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /api/condition-templates [put]
func EditConditionTemplate(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		data := requests.EditConditionTemplate{}
		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid json",
				Error:  err.Error(),
			})
		}

		validationErrors := validator.ValidateStruct(data)
		if validationErrors != nil {
			return c.Status(422).JSON(responses.ValidationErrorResponse{
				Status: "validation failed",
				Errors: validationErrors,
			})
		}

		var conditionTemplate database.ConditionTemplate
		result := db.First(&conditionTemplate, data.ID)
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(responses.ErrorResponse{
				Status: "condition template not found",
			})
		} else if result.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "internal server error",
				Error:  result.Error.Error(),
			})
		}

		if conditionTemplate.AuthorID != authorID {
			return c.Status(403).JSON(responses.ErrorResponse{
				Status: "forbidden",
				Error:  "you are not the author of this condition template",
			})
		}

		conditionTemplate.Title = data.Title
		conditionTemplate.Condition = data.Condition
		conditionTemplate.UpdatedAt = time.Now()

		saveResult := db.Save(&conditionTemplate)
		if saveResult.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to update condition template",
				Error:  saveResult.Error.Error(),
			})
		}

		return c.Status(200).JSON(responses.EditConditionTemplateDTO{
			Status: "condition template updated",
			TaskTemplate: responses.ConditionTemplateDTO{
				ID:        conditionTemplate.ID,
				Title:     conditionTemplate.Title,
				Condition: conditionTemplate.Condition,
			},
		})
	}
}

// DeleteConditionTemplate godoc
// @Summary Delete a condition template by ID
// @Description Deletes a specific condition template owned by the user
// @Tags ConditionTemplate
// @Accept json
// @Produce json
// @Param data body requests.DeleteConditionTemplate true "Condition template ID"
// @Success 200 {object} responses.DeleteConditionTemplateDTO
// @Failure 400 {object} responses.ErrorResponse "Invalid token or request"
// @Failure 403 {object} responses.ErrorResponse "Forbidden"
// @Failure 404 {object} responses.ErrorResponse "Not found"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /api/condition-templates [delete]
func DeleteConditionTemplate(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		data := requests.DeleteConditionTemplate{}
		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid json",
				Error:  err.Error(),
			})
		}

		validationErrors := validator.ValidateStruct(data)
		if validationErrors != nil {
			return c.Status(422).JSON(responses.ValidationErrorResponse{
				Status: "validation failed",
				Errors: validationErrors,
			})
		}

		var conditionTemplate database.ConditionTemplate
		result := db.First(&conditionTemplate, data.ID)
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(responses.ErrorResponse{
				Status: "condition template not found",
			})
		} else if result.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "internal server error",
				Error:  result.Error.Error(),
			})
		}

		if conditionTemplate.AuthorID != authorID {
			return c.Status(403).JSON(responses.ErrorResponse{
				Status: "forbidden",
				Error:  "you are not the author of this condition template",
			})
		}

		deleteResult := db.Delete(&conditionTemplate)
		if deleteResult.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to delete condition template",
				Error:  deleteResult.Error.Error(),
			})
		}

		return c.Status(200).JSON(responses.DeleteConditionTemplateDTO{
			Status: "condition template deleted",
		})
	}
}

// GetAllConditionTemplates godoc
// @Summary Get all condition templates for the user
// @Description Retrieves all condition templates created by the current user with pagination
// @Tags ConditionTemplate
// @Param offset query int false "Pagination offset (default: 0)"
// @Success 200 {object} responses.GetAllConditionTemplatesDTO
// @Failure 400 {object} responses.ErrorResponse "Invalid token or request"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /api/condition-templates [get]
func GetAllConditionTemplates(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		offset, err := strconv.Atoi(c.Query("offset", "0"))
		if err != nil || offset < 0 {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid offset",
				Error:  "offset must be a non-negative integer",
			})
		}

		limit := 10
		var conditions []database.ConditionTemplate
		result := db.Where("author_id = ?", authorID).
			Order("id DESC").
			Offset(offset * limit).
			Limit(limit).
			Find(&conditions)

		if result.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "internal server error",
				Error:  result.Error.Error(),
			})
		}

		var responseConditions []responses.ConditionTemplateDTO
		for _, condition := range conditions {
			responseConditions = append(responseConditions, responses.ConditionTemplateDTO{
				ID:        condition.ID,
				Title:     condition.Title,
				Condition: condition.Condition,
			})
		}

		return c.Status(200).JSON(responses.GetAllConditionTemplatesDTO{
			TaskTemplates: responseConditions,
		})
	}
}
