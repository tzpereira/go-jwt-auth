package controller

import (
	"os"
	"go-jwt-auth/internal/service"
	"go-jwt-auth/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// POST /jwt-generate
func GenerateJWT(ctx *fiber.Ctx) error {
	var req struct {
		Password string `json:"password"`
		Sub      string `json:"sub"`
	}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
       envSub := os.Getenv("SUB")
       if envSub == "" || req.Sub != envSub || req.Password != os.Getenv("JWT_SECRET") {
	       return ctx.Status(401).JSON(fiber.Map{"error": "unauthorized"})
       }
	token, err := service.GenerateJWT(req.Sub)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "could not generate token"})
	}
	repository.AddToWhitelist(token)
	return ctx.JSON(fiber.Map{"token": token})
}

// GET /jwt-authorize
func AuthJWT(ctx *fiber.Ctx) error {
	token := ctx.Locals("jwt_token")
	jwtToken, ok := token.(*jwt.Token)
	if !ok || jwtToken == nil {
		return ctx.Status(401).JSON(fiber.Map{"error": "invalid token"})
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return ctx.Status(500).JSON(fiber.Map{"error": "cannot parse claims"})
	}
	return ctx.JSON(fiber.Map{"message": "JWT is valid!", "claims": claims})
}