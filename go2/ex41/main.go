package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Custom-Header", "FIBER")
		return c.Next()
	})
	api := app.Group("/api")
	api.Use(func(c *fiber.Ctx) error {
		c.Locals("xyz", "XYZ")
		return c.Next()
	})
	api.Get("/list", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"method":       c.Method(),
			"protocol":     c.Protocol(),
			"baseurl":      c.BaseURL(),
			"original-url": c.OriginalURL(),
			"path":         c.Path(),
			"query":        c.Query("q"),
			"content-type": c.Get("Content-Type"),
			"hostname":     c.Hostname(),
			"xyz":          c.Locals("xyz"),
		})
	})
	if err := app.Listen(":1234"); err != nil {
		panic(err)
	}
}

/*
curl -v -H "Content-Type:application/json" http://localhost:1234/api/list?q=abc
*/
