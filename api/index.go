package handler

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

var app = newApp()

func newApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service": "react-in-go",
			"status":  "ok",
		})
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	return app
}

// Handler is the Vercel Function entrypoint.
func Handler(w http.ResponseWriter, r *http.Request) {
	// Vercel forwards requests to this function under /api. Normalize that
	// prefix so Fiber routes work consistently for /api and /api/health.
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
	if r.URL.Path == "" {
		r.URL.Path = "/"
	}
	r.RequestURI = r.URL.RequestURI()

	adaptor.FiberApp(app)(w, r)
}
