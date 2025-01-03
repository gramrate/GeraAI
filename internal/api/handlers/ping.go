package handlers

import (
	"gera-ai/internal/models/responses"
	"github.com/gofiber/fiber/v2"
)

// Ping checks the server status
// @Summary Server Health Check
// @Description Returns a simple status response to verify the server is running.
// @Tags Health
// @Produce json
// @Success 200 {object} responses.PingDTO "Server is running"
// @Router /ping [get]
func Ping() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON(responses.PingDTO{
			Status: "ok",
		})
	}
}
