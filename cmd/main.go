package main

import (
	"go-jwt-auth/internal/view"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	app := fiber.New()
	view.InitializeRoutes(app)

	log.Fatal(app.Listen(":9001"))
}
