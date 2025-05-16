package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	User struct {
		Name  string `json:"name" validate:"required,min=3,max=32"`
		Email string `json:"email" validate:"required,email,min=6,max=32"`
		Job   Job    `json:"job" validate:"dive"`
	}
	Job struct {
		JobId int    `json:"job_id" validate:"required,number"`
		Type  string `json:"type" validate:"required,min=3,max=32"`
	}
	ValidationError struct {
		Field string      `json:"field"`
		Tag   string      `json:"tag"`
		Value interface{} `json:"value"`
	}
)

func main() {
	app := fiber.New()
	app.Post("/create-user", createUser)
	if err := app.Listen(":1234"); err != nil {
		panic(err)
	}
}

func validationError(err error) []ValidationError {
	errs := make([]ValidationError, 0)
	for _, terr := range err.(validator.ValidationErrors) {
		errs = append(errs, ValidationError{
			Field: terr.StructNamespace(),
			Tag:   terr.Tag(),
			Value: terr.Value(),
		})
	}
	return errs
}

func createUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationError(err))
	}
	return c.JSON(user)
}

// curl -X POST -H "Content-Type: application/json" --data '{"name":"john","email":"bloomfilter"}' http://localhost:1234/create-user
