package middleware

import (
	"strings"

	"go-jwt-auth/internal/repository"
	"go-jwt-auth/internal/service"

	"github.com/gofiber/fiber/v2"
)

// JWTAuthMiddleware protects routes by requiring a valid, non-blacklisted, whitelisted JWT
func JWTAuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing authorization header"})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid authorization header format"})
	}

	if repository.IsBlacklisted(tokenString) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "token is blacklisted"})
	}

	// JWT validation
	token, err := service.ParseJWT(tokenString)
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
	}

	if !repository.IsWhitelisted(tokenString) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "token not whitelisted"})
	}

	// Store the validated token in context for downstream handlers
	c.Locals("jwt_token", token)
	c.Locals("jwt_token_string", tokenString)

	return c.Next()
}
