package main

import (
	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}

func main() {
	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		p := new(Person)
		if err := c.BodyParser(p); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.JSON(p)
	})
	if err := app.Listen(":1234"); err != nil {
		panic(err)
	}
}

/*
curl -X POST -H "Content-Type: application/json" --data '{"name":"john","pass":"doe"}' localhost:1234
curl -X POST -H "Content-Type: application/xml" --data "<login><name>john</name><pass>doe</pass></login>" localhost:1234
curl -X POST -H "Content-Type: application/x-www-form-urlencoded" --data "name=john&pass=doe" localhost:1234
curl -X POST -F name=john -F pass=doe http://localhost:1234
*/
