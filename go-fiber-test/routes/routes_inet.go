package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// CRUD profiles (without Basic Auth)
	profile := v1.Group("/profile")
	profile.Get("", c.GetProfiles)
	profile.Get("/filter", c.GetProfile)
	profile.Get("/ages", c.GetProfileAnyAges)
	profile.Get("/search", c.SearchProfiles)

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))

	// CRUD profiles (with Basic Auth)
	profile.Post("/", c.AddProfile)
	profile.Put("/:id", c.UpdateProfile)
	profile.Delete("/:id", c.RemoveProfile)

	// Other APIs under v1
	v1.Get("/", c.HelloTest)
	v1.Post("/", c.BodyParser)
	v1.Get("/user/:name", c.Params)
	v1.Post("/inet", c.Query)
	v1.Post("/valid", c.Validation)
	v1.Get("/fact/:num", c.Factorial)

	// v3 APIs
	v3 := api.Group("/v3")
	v3.Get("/:nickname", c.ConvertTaxIDToASCII)

	// CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/deleted", c.GetDeletedDogs)
	dog.Get("/range", c.GetDogsRangeCountByDogId)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogsJsonSummary)
	dog.Post("/", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)

	// CRUD companies
	company := v1.Group("/company")
	company.Get("", c.GetCompanies)
	company.Get("/filter", c.GetCompany)
	company.Post("/", c.AddCompany)
	company.Put("/:id", c.UpdateCompany)
	company.Delete("/:id", c.RemoveCompany)
}
