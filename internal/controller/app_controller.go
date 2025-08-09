package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// GET /app/protected
func AppProtected(ctx *fiber.Ctx) error {
	token := ctx.Locals("jwt_token")

	jwtToken, ok := token.(*jwt.Token)
	if !ok || jwtToken == nil {
		return ctx.Status(401).JSON(fiber.Map{"error": "invalid token"})
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return ctx.Status(500).JSON(fiber.Map{"error": "cannot parse claims"})
	}

	sub, _ := claims["sub"].(string)

	return ctx.JSON(fiber.Map{
		"message": "Access to application granted",
		"sub": sub,
		"app_data": fiber.Map{
			"info": "Protected app area",
		},
	})
}
