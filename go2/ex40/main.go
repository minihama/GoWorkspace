package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	app.Get("/value/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
	})
	app.Get("/name/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") == "" {
			return fiber.NewError(400, "name is missing")
		}
		return c.SendString("Hello, " + c.Params("name"))
	})
	app.Get("/wildcard/*", func(c *fiber.Ctx) error {
		return c.SendString("Wildcard: " + c.Params("*"))
	})
	app.Static("/", "./public")
	if err := app.Listen(":1234"); err != nil {
		panic(err)
	}
}

/*
curl http://localhost:1234/hello
curl http://localhost:1234/value/good
curl http://localhost:1234/name/John
curl http://localhost:1234/name/
curl http://localhost:1234/wildcard/aa/bb/cc
curl http://localhost:1234/greeting.html
*/
