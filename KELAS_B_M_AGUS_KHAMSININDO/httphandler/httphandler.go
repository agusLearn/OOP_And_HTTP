package httphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"

	jwtware "github.com/gofiber/jwt/v2"
)

func TryHTTPHandler(app *fiber.App, secret string) {
	// Add logic in here to complete the tests
	basicauthHandler := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "rakamin",
		},
	})

	app.Use("/internal", basicauthHandler).Get("/internal", func(c *fiber.Ctx) (err error) {
		err = c.SendString("Welcome to the internal of Rakamin!")
		return
	})

	apiGroup := app.Group("/")

	var kunciRahasia = []byte(secret)

	apiGroup.Use("/admin", jwtware.New(jwtware.Config{
		SigningKey: kunciRahasia,
	}))

	apiGroup.Get("/admin", func(c *fiber.Ctx) error {
		return c.SendString("Welcome back, admin!")
	})
}
