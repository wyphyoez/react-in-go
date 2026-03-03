package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// Load templates
	engine := html.New("./", ".html")

	// Create app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Configure static files
	app.Static("/", "./public")

	// Home page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	// API v2
	app.Get("/api/v2/message", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":      "react-in-go",
			"version":   "v2",
			"framework": "Fiber + React",
			"status":    "ok",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
			"features": []string{
				"Go backend with Fiber",
				"React 18 frontend",
				"TypeScript entrypoint",
				"esbuild bundling",
			},
		})
	})

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
