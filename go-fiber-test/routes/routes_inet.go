package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))

	// Grouping API
	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1
	v3 := api.Group("/v3")   // /api/v3

	v1.Get("/", c.HelloTest)

	v1.Post("/", c.BodyParser)

	v1.Get("/user/:name", c.Params)

	v1.Post("/inet", c.Query)

	v1.Post("/valid", c.Validation)

	v1.Get("/fact/:num", c.Factorial)

	v3.Get("/:nickname", c.ConvertTaxIDToASCII)

	v1.Post("/register", c.Register)
}
