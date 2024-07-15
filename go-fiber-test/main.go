package main

import (
	"go-fiber-test/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.InetRoutes(app)
	app.Listen(":3000")
}
