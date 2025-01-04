package handlers

import (
	"fmt"
	"gera-ai/internal/config"
	dbmodels "gera-ai/internal/models/database"
	"gera-ai/internal/models/requests"
	"gera-ai/internal/models/responses"
	"gera-ai/internal/utils/password"
	"gera-ai/internal/utils/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"time"
)

// Login handles user authentication
// @Summary      User Login
// @Description  Authenticate user and return JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      requests.Login  true  "Login details"
// @Success      200      {object}  responses.AuthDTO
// @Failure      400      {object}  responses.ErrorResponse
// @Failure      401      {object}  responses.ErrorResponse
// @Failure      404      {object}  responses.ErrorResponse
// @Failure      422      {object}  responses.ValidationErrorResponse
// @Failure      500      {object}  responses.ErrorResponse
// @Router       /api/auth/login [post]
func Login(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := requests.Login{}
		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "error",
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

		var user dbmodels.User
		result := db.First(&user, "login = ?", data.Login)
		if result.Error != nil {
			return c.Status(404).JSON(responses.ErrorResponse{
				Status: "error",
				Error:  "User not found",
			})
		}

		if !password.CheckPasswordHash(data.Password, user.PasswordHash) {
			return c.Status(401).JSON(responses.ErrorResponse{
				Status: "error",
				Error:  "Wrong password",
			})
		}

		claims := jwt.MapClaims{
			"id":  fmt.Sprint(user.ID),
			"exp": time.Now().Add(config.Config.JWTExpiration).Unix(),
			"iat": time.Now().Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, tokenErr := token.SignedString([]byte(config.Config.JWTSecret))
		if tokenErr != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "error",
				Error:  "Failed to sign token",
			})
		}

		return c.Status(200).JSON(responses.AuthDTO{
			Status: "ok",
			Token:  signedToken,
		})
	}
}

// Register handles user registration
// @Summary      User Registration
// @Description  Register a new user and return JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      requests.Register  true  "Registration details"
// @Success      200      {object}  responses.AuthDTO
// @Failure      400      {object}  responses.ErrorResponse
// @Failure      409      {object}  responses.ErrorResponse
// @Failure      422      {object}  responses.ValidationErrorResponse
// @Failure      500      {object}  responses.ErrorResponse
// @Router       /api/auth/register [post]
func Register(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := requests.Register{}

		// 1. Парсинг тела запроса
		if reqErr := c.BodyParser(&data); reqErr != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "error",
				Error:  "Invalid JSON format",
			})
		}

		// 2. Валидация данных
		validationErrors := validator.ValidateStruct(data)
		if validationErrors != nil {
			return c.Status(422).JSON(responses.ValidationErrorResponse{
				Status: "error",
				Errors: validationErrors,
			})
		}

		// 3. Проверка на существование пользователя с таким же логином
		var existingUser dbmodels.User
		checkResult := db.Where("login = ?", data.Login).First(&existingUser)
		if checkResult.Error == nil {
			return c.Status(409).JSON(responses.ErrorResponse{
				Status: "error",
				Error:  "User with this login already exists",
			})
		} else if checkResult.Error != gorm.ErrRecordNotFound {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "error",
				Error:  "Internal server error during user lookup",
			})
		}

		// 4. Хэширование пароля
		hashedPassword, hashErr := password.HashPassword(data.Password)
		if hashErr != nil {
			return c.Status(400).JSON(responses.ErrorResponse{
				Status: "error",
				Error:  "Failed to hash password",
			})
		}

		// 5. Создание пользователя
		user := &dbmodels.User{
			Login:        data.Login,
			Username:     data.Username,
			PasswordHash: hashedPassword,
			CreatedAt:    time.Now(),
		}
		result := db.Create(&user)
		if result.Error != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "error",
				Error:  "Failed to create user",
			})
		}

		// 6. Генерация токена
		claims := jwt.MapClaims{
			"id":  fmt.Sprint(user.ID),
			"exp": time.Now().Add(config.Config.JWTExpiration).Unix(),
			"iat": time.Now().Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, tokenErr := token.SignedString([]byte(config.Config.JWTSecret))
		if tokenErr != nil {
			return c.Status(500).JSON(responses.ErrorResponse{
				Status: "error",
				Error:  "Failed to sign token",
			})
		}

		// 7. Успешный ответ
		return c.Status(200).JSON(responses.AuthDTO{
			Status: "success",
			Token:  signedToken,
		})
	}
}
