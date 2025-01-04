package handlers

import (
	"strconv"
	"time"

	dbmodels "gera-ai/internal/models/database"
	"gera-ai/internal/models/requests"
	"gera-ai/internal/models/responses"
	"gera-ai/internal/utils/jwtUtils"
	"gera-ai/internal/utils/validator"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateTask creates a new task
// @Summary Create a new task
// @Description Creates a new task and saves it to the database
// @Tags Tasks
// @Accept json
// @Produce json
// @Param input body requests.CreateTask true "Task creation data"
// @Success 200 {object} responses.CreateTaskResponseDTO "Task successfully created"
// @Failure 400 {object} responses.ErrorResponse "Invalid token or JSON"
// @Failure 422 {object} responses.ValidationErrorResponse "Validation error"
// @Failure 500 {object} responses.ErrorResponse "Database error"
// @Router /api/task/new [post]
func CreateTask(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		data := requests.CreateTask{}
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

		// Создание задачи
		task := dbmodels.Task{
			AuthorID:  authorID,
			Title:     data.Title,
			Condition: data.Condition,
			Answer:    data.Answer,
			CreatedAt: time.Now(),
		}

		// Сохранение в базе данных
		result := db.Create(&task)
		if result.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to create task template",
				Error:  result.Error.Error(),
			})
		}

		// Ответ с успешным созданием
		return c.Status(200).JSON(responses.CreateTaskResponseDTO{
			Status: "task created",
			Task: responses.TaskDTO{
				ID:        task.ID,
				Title:     task.Title,
				Condition: task.Condition,
				Answer:    task.Answer,
			},
		})
	}
}

// GetTask retrieves a task by ID
// @Summary Retrieve a task by ID
// @Description Fetches a specific task by its ID
// @Tags Tasks
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} responses.GetTaskResponseDTO "Task successfully retrieved"
// @Failure 400 {object} responses.ErrorResponse "Invalid token or JSON"
// @Failure 403 {object} responses.ErrorResponse "Forbidden access"
// @Failure 404 {object} responses.ErrorResponse "Task not found"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /api/task/get/{id} [get]
func GetTask(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		paramsTaskID := c.Params("id")
		taskID, err := strconv.Atoi(paramsTaskID)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid json",
				Error:  err.Error(),
			})
		}

		// Валидация ID шаблона
		data := requests.GetTask{ID: uint(taskID)}
		validationErrors := validator.ValidateStruct(data)
		if validationErrors != nil {
			return c.Status(422).JSON(responses.ValidationErrorResponse{
				Status: "validation failed",
				Errors: validationErrors,
			})
		}

		// Получение шаблона задачи из базы данных
		var task dbmodels.Task
		result := db.First(&task, data.ID)
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(responses.ErrorResponse{
				Status: "task not found",
			})
		} else if result.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "internal server error",
				Error:  result.Error.Error(),
			})
		}

		// Проверка, что текущий пользователь является автором задания
		if task.AuthorID != authorID {
			return c.Status(403).JSON(responses.ErrorResponse{
				Status: "forbidden",
				Error:  "you are not the author of this task",
			})
		}

		// Возвращаем найденное задание
		return c.Status(200).JSON(responses.GetTaskResponseDTO{
			Status: "success",
			Task: responses.TaskDTO{
				ID:        task.ID,
				Title:     task.Title,
				Condition: task.Condition,
				Answer:    task.Answer,
			},
		})
	}
}

// EditTask edits an existing task
// @Summary Edit a task
// @Description Updates an existing task with new data
// @Tags Tasks
// @Accept json
// @Produce json
// @Param input body requests.EditTask true "Task editing data"
// @Success 200 {object} responses.EditTaskResponseDTO "Task successfully updated"
// @Failure 400 {object} responses.ErrorResponse "Invalid token or JSON"
// @Failure 403 {object} responses.ErrorResponse "Forbidden access"
// @Failure 404 {object} responses.ErrorResponse "Task not found"
// @Failure 422 {object} responses.ValidationErrorResponse "Validation error"
// @Failure 500 {object} responses.ErrorResponse "Database error"
// @Router /api/task/edit [put]
func EditTask(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		data := requests.EditTask{}
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

		// Поиск шаблона по ID
		var task dbmodels.Task
		result := db.First(&task, data.ID)
		if result.Error == gorm.ErrRecordNotFound {
			// Возвращаем ошибку, если шаблон не найден
			return c.Status(404).JSON(responses.ErrorResponse{
				Status: "task not found",
			})
		} else if result.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "internal server error",
				Error:  result.Error.Error(),
			})
		}

		// Проверка, что текущий пользователь является автором шаблона
		if task.AuthorID != authorID {
			return c.Status(403).JSON(responses.ErrorResponse{
				Status: "forbidden",
				Error:  "you are not the author of this task",
			})
		}

		// Обновление данных шаблона
		task.Title = data.Title
		task.Condition = data.Condition
		task.Answer = data.Answer
		task.UpdatedAt = time.Now()

		// Сохранение изменений
		saveResult := db.Save(&task)
		if saveResult.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to update task",
				Error:  saveResult.Error.Error(),
			})
		}

		// Возвращаем обновленный шаблон
		return c.Status(200).JSON(responses.EditTaskResponseDTO{
			Status: "task template updated",
			Task: responses.TaskDTO{
				ID:        task.ID,
				Title:     task.Title,
				Condition: task.Condition,
				Answer:    task.Answer,
			},
		})
	}
}

// DeleteTask deletes a task by ID
// @Summary Delete a task
// @Description Deletes a specific task by its ID
// @Tags Tasks
// @Accept json
// @Produce json
// @Param input body requests.DeleteTask true "Task deletion data"
// @Success 200 {object} responses.DeleteTaskResponseDTO "Task successfully deleted"
// @Failure 400 {object} responses.ErrorResponse "Invalid token or JSON"
// @Failure 403 {object} responses.ErrorResponse "Forbidden access"
// @Failure 404 {object} responses.ErrorResponse "Task not found"
// @Failure 500 {object} responses.ErrorResponse "Database error"
// @Router /api/task/delete [delete]
func DeleteTask(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		data := requests.DeleteTask{}
		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid json",
				Error:  err.Error(),
			})
		}

		// Валидация ID
		validationErrors := validator.ValidateStruct(data)
		if validationErrors != nil {
			return c.Status(422).JSON(responses.ValidationErrorResponse{
				Status: "validation failed",
				Errors: validationErrors,
			})
		}

		// Поиск шаблона для удаления
		var task dbmodels.Task
		result := db.First(&task, data.ID)
		if result.Error == gorm.ErrRecordNotFound {
			// Возвращаем ошибку, если шаблон не найден
			return c.Status(404).JSON(responses.ErrorResponse{
				Status: "task not found",
			})
		} else if result.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "internal server error",
				Error:  result.Error.Error(),
			})
		}

		// Проверка, что текущий пользователь является автором шаблона
		if task.AuthorID != authorID {
			return c.Status(403).JSON(responses.ErrorResponse{
				Status: "forbidden",
				Error:  "you are not the author of this task",
			})
		}

		deleteResult := db.Delete(&task)
		if deleteResult.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "failed to delete task",
				Error:  deleteResult.Error.Error(),
			})
		}

		// Возвращаем успешный ответ
		return c.Status(200).JSON(responses.DeleteTaskResponseDTO{
			Status: "task deleted",
		})
	}
}

// GetAllTasks retrieves all tasks for a user
// @Summary Retrieve all tasks
// @Description Fetches all tasks created by the authenticated user with pagination
// @Tags Tasks
// @Produce json
// @Param offset query int false "Pagination offset (default is 0)"
// @Success 200 {object} responses.GetAllTasksResponseDTO "Tasks successfully retrieved"
// @Failure 400 {object} responses.ErrorResponse "Invalid token or offset"
// @Failure 404 {object} responses.ErrorResponse "No tasks found"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /api/task/all [get]
func GetAllTasks(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorID, err := jwtUtils.ExtractUserID(c)
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid token",
				Error:  err.Error(),
			})
		}

		paramsOffset := c.Query("offset", "0")
		offset, err := strconv.Atoi(paramsOffset)
		if offset < 0 {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid offset",
				Error:  "offset must be a positive integer",
			})
		}
		if err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "invalid json",
				Error:  err.Error(),
			})
		}

		// Получение шаблона задачи из базы данных
		var tasks []dbmodels.Task
		limit := 10
		result := db.Where("author_id = ?", authorID).
			Order("id DESC").
			Offset(offset * limit).
			Limit(limit).
			Find(&tasks)
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

		tasksDTO := make([]responses.TaskDTO, len(tasks))
		for i, task := range tasks {
			tasksDTO[i] = responses.TaskDTO{
				ID:        task.ID,
				Title:     task.Title,
				Condition: task.Condition,
				Answer:    task.Answer,
			}
		}

		// Возвращаем найденные задачи
		return c.Status(200).JSON(responses.GetAllTasksResponseDTO{
			Tasks: tasksDTO,
		})
	}
}
