package jwtUtils

import (
	"errors"
	"gera-ai/internal/utils/parser"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// ExtractUserID извлекает ID пользователя из JWT токена, если он валиден
func ExtractUserID(c *fiber.Ctx) (uint, error) {
	// Извлекаем JWT токен из контекста
	jwtToken, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return 0, errors.New("invalid token format")
	}

	// Извлекаем клеймы из токена
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// Получаем ID пользователя из клеймов
	jwtID, ok := claims["id"].(string)
	if !ok {
		return 0, errors.New("missing or invalid user ID in token")
	}

	// Преобразуем ID пользователя в uint
	authorID, err := parser.StringToUint(jwtID)
	if err != nil {
		return 0, err
	}

	return authorID, nil
}
