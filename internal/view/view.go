package view

import (
	"go-jwt-auth/internal/controller"
	"go-jwt-auth/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
		// Public route for health check
		app.Get("/health", controller.Health)

		// Public route to generate JWT
		app.Post("/jwt-generate", controller.GenerateJWT)

		// Protected route to authenticate JWT
		app.Get("/jwt-authorize", middleware.JWTAuthMiddleware, controller.AuthJWT)

		// Protected route to test that auth works
		app.Get("/app/protected", middleware.JWTAuthMiddleware, controller.AppProtected)

	// Public route for fallback
	app.All("*", func(c *fiber.Ctx) error {
		return c.SendString("Fallback catch-all route: " + c.Method() + " " + c.Path())
	})
}
