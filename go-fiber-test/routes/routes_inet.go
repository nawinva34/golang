package routes

import (
	"go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))

	//Grouping api
	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1

	v1.Get("/", controllers.HelloTest)

	v1.Post("/", controllers.BodyParser)

	v1.Get("/user/:name", controllers.Params)

	v1.Post("/inet", controllers.Query)

	v1.Post("/valid", controllers.Validation)
}
