package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func main() {
	engine := django.New("./views", ".django")
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})
	app.Get("/layouts", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
			"user":  User{FirstName: "Bloom", LastName: "Filter"},
			"testMap": map[string]string{
				"testKey": "aaa",
			},
			"testSlice": []string{"This", "is", "a", "test", "string"},
		}, "layouts/main")
	})
	log.Fatal(app.Listen(":3000"))
}
