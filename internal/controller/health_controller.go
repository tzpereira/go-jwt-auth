package controller

import "github.com/gofiber/fiber/v2"

// GET /health
func Health(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"status": "ok"})
}