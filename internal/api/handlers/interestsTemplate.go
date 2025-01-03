package handlers

import (
	"encoding/json"
	dbmodels "gera-ai/internal/models/database"
	"gera-ai/internal/models/requests"
	"gera-ai/internal/models/responses"
	"gera-ai/internal/utils/jsonUtils"
	"gera-ai/internal/utils/jwtUtils"
	"gera-ai/internal/utils/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// CreateInterestsTemplate creates a new interests template
// @Summary Create Interests Template
// @Description Creates a new template based on the provided title and list of interests.
// @Tags Interests Template
// @Accept json
// @Produce json
// @Param input body requests.CreateInterestsTemplate true "Template data"
// @Success 200 {object} responses.CreateInterestsTemplateDTO "Successfully created template"
// @Failure 400 {object} responses.ErrorResponse "Invalid token or JSON"
// @Failure 422 {object} responses.ValidationErrorResponse "Validation error"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /interests/template [post]
func CreateInterestsTemplate(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		data := requests.CreateInterestsTemplate{}
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

		interestsJSON, err := json.Marshal(data.Interests)
		if err != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to serialize interests",
				Error:  err.Error(),
			})
		}

		taskTemplate := dbmodels.InterestsTemplate{
			AuthorID:  authorID,
			Title:     data.Title,
			Interests: interestsJSON,
			CreatedAt: time.Now(),
		}

		result := db.Create(&taskTemplate)
		if result.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to create task template",
				Error:  result.Error.Error(),
			})
		}

		interestsList, err := jsonUtils.ConvertInterestsToList(taskTemplate.Interests)
		if err != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to convert interests to list",
				Error:  err.Error(),
			})
		}

		return c.Status(200).JSON(responses.CreateInterestsTemplateDTO{
			Status: "task template created",
			TaskTemplate: responses.InterestsTemplateDTO{
				ID:        taskTemplate.ID,
				Title:     taskTemplate.Title,
				Interests: interestsList,
			},
		})
	}
}

// GetInterestsTemplate retrieves an interests template by ID
// @Summary Get Interests Template
// @Description Fetches a specific interests template by its ID, provided the user is the author.
// @Tags Interests Template
// @Produce json
// @Param id path int true "Template ID"
// @Success 200 {object} responses.GetInterestsTemplateDTO "Successfully retrieved template"
// @Failure 400 {object} responses.ErrorResponse "Invalid token or template ID"
// @Failure 403 {object} responses.ErrorResponse "Access forbidden"
// @Failure 404 {object} responses.ErrorResponse "Template not found"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /interests/template/{id} [get]
func GetInterestsTemplate(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		taskTemplateID, err := strconv.Atoi(c.Params("id"))
		if err != nil || taskTemplateID <= 0 {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid template ID",
				Error:  "template ID must be a positive integer",
			})
		}

		var taskTemplate dbmodels.InterestsTemplate
		result := db.First(&taskTemplate, taskTemplateID)
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(responses.ErrorResponse{
				Status: "task template not found",
			})
		} else if result.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "internal server error",
				Error:  result.Error.Error(),
			})
		}

		if taskTemplate.AuthorID != authorID {
			return c.Status(403).JSON(responses.ErrorResponse{
				Status: "forbidden",
				Error:  "you are not the author of this task template",
			})
		}

		interests, err := jsonUtils.ConvertInterestsToList(taskTemplate.Interests)
		if err != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to decode interests",
				Error:  err.Error(),
			})
		}

		return c.Status(200).JSON(responses.GetInterestsTemplateDTO{
			TaskTemplate: responses.InterestsTemplateDTO{
				ID:        taskTemplate.ID,
				Title:     taskTemplate.Title,
				Interests: interests,
			},
		})
	}
}

// EditInterestsTemplate updates an existing interests template
// @Summary Edit Interests Template
// @Description Updates the title or list of interests of an existing template.
// @Tags Interests Template
// @Accept json
// @Produce json
// @Param input body requests.EditInterestsTemplate true "Updated template data"
// @Success 200 {object} responses.EditInterestsTemplateDTO "Successfully updated template"
// @Failure 400 {object} responses.ErrorResponse "Invalid token or JSON"
// @Failure 403 {object} responses.ErrorResponse "Access forbidden"
// @Failure 404 {object} responses.ErrorResponse "Template not found"
// @Failure 422 {object} responses.ValidationErrorResponse "Validation error"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /interests/template [put]
func EditInterestsTemplate(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		data := requests.EditInterestsTemplate{}
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

		var taskTemplate dbmodels.InterestsTemplate
		result := db.First(&taskTemplate, data.ID)
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(responses.ErrorResponse{
				Status: "task template not found",
			})
		} else if result.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "internal server error",
				Error:  result.Error.Error(),
			})
		}

		if taskTemplate.AuthorID != authorID {
			return c.Status(403).JSON(responses.ErrorResponse{
				Status: "forbidden",
				Error:  "you are not the author of this task template",
			})
		}

		taskTemplate.Title = data.Title
		if data.Interests != nil {
			interestsJSON, err := json.Marshal(data.Interests)
			if err != nil {
				return c.Status(500).JSON(responses.ErrorResponse{
					Status: "failed to process interests",
					Error:  err.Error(),
				})
			}
			taskTemplate.Interests = interestsJSON
		}
		taskTemplate.UpdatedAt = time.Now()

		saveResult := db.Save(&taskTemplate)
		if saveResult.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to update task template",
				Error:  saveResult.Error.Error(),
			})
		}

		interests, err := jsonUtils.ConvertInterestsToList(taskTemplate.Interests)
		if err != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to decode interests",
				Error:  err.Error(),
			})
		}

		return c.Status(200).JSON(responses.EditInterestsTemplateDTO{
			Status: "task template updated",
			TaskTemplate: responses.InterestsTemplateDTO{
				ID:        taskTemplate.ID,
				Title:     taskTemplate.Title,
				Interests: interests,
			},
		})
	}
}

// DeleteInterestsTemplate deletes an interests template by ID
// @Summary Delete Interests Template
// @Description Deletes a specific interests template, provided the user is the author.
// @Tags Interests Template
// @Accept json
// @Produce json
// @Param input body requests.DeleteInterestsTemplate true "Template ID"
// @Success 200 {object} responses.DeleteInterestsTemplateDTO "Successfully deleted template"
// @Failure 400 {object} responses.ErrorResponse "Invalid token or JSON"
// @Failure 403 {object} responses.ErrorResponse "Access forbidden"
// @Failure 404 {object} responses.ErrorResponse "Template not found"
// @Failure 422 {object} responses.ValidationErrorResponse "Validation error"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /interests/template [delete]
func DeleteInterestsTemplate(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		data := requests.DeleteInterestsTemplate{}
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

		var taskTemplate dbmodels.InterestsTemplate
		result := db.First(&taskTemplate, data.ID)
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(responses.ErrorResponse{
				Status: "task template not found",
			})
		} else if result.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "internal server error",
				Error:  result.Error.Error(),
			})
		}

		if taskTemplate.AuthorID != authorID {
			return c.Status(403).JSON(responses.ErrorResponse{
				Status: "forbidden",
				Error:  "you are not the author of this task template",
			})
		}

		deleteResult := db.Delete(&taskTemplate)
		if deleteResult.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to delete task template",
				Error:  deleteResult.Error.Error(),
			})
		}

		return c.Status(200).JSON(responses.DeleteInterestsTemplateDTO{
			Status: "task template deleted",
		})
	}
}

// GetAllInterestsTemplates retrieves all templates for the current user
// @Summary Get All Interests Templates
// @Description Retrieves a paginated list of all interests templates created by the user.
// @Tags Interests Template
// @Produce json
// @Param offset query int false "Pagination offset (default: 0)"
// @Success 200 {object} responses.GetAllInterestsTemplatesDTO "Successfully retrieved templates"
// @Failure 400 {object} responses.ErrorResponse "Invalid token or offset"
// @Failure 404 {object} responses.ErrorResponse "No templates found"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /interests/templates [get]
func GetAllInterestsTemplates(db *gorm.DB) fiber.Handler {
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
		var templates []dbmodels.InterestsTemplate
		result := db.Where("author_id = ?", authorID).
			Order("id DESC").
			Offset(offset * limit).
			Limit(limit).
			Find(&templates)

		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				return c.Status(404).JSON(responses.ErrorResponse{
					Status: "no task templates found",
				})
			}
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "internal server error",
				Error:  result.Error.Error(),
			})
		}

		var taskTemplates []responses.InterestsTemplateDTO
		for _, template := range templates {
			interests, err := jsonUtils.ConvertInterestsToList(template.Interests)
			if err != nil {
				return c.Status(500).JSON(responses.ErrorResponse{
					Status: "failed to decode interests",
					Error:  err.Error(),
				})
			}
			taskTemplates = append(taskTemplates, responses.InterestsTemplateDTO{
				ID:        template.ID,
				Title:     template.Title,
				Interests: interests,
			})
		}

		return c.Status(200).JSON(responses.GetAllInterestsTemplatesDTO{
			TaskTemplates: taskTemplates,
		})
	}
}
