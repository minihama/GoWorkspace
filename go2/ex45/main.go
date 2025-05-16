package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

type User struct {
	FirstName  string
	LastName   string
	JuminBunho string `validate:"required,juminbunho"`
}

const juminBunho = `^(?:[0-9]{2}(?:0[1-9]|1[0-2])(?:0[1-9]|[1,2][0-9]|3[0,1]))[1-8][0-9]{6}$`

var regexJuminBunho = regexp.MustCompile(juminBunho)

func main() {
	user := User{
		FirstName:  "John",
		LastName:   "Smith",
		JuminBunho: "8803279123456",
	}
	validate := validator.New()
	validate.RegisterStructValidation(validateUser, User{})
	if err := validate.RegisterValidation("juminbunho", validateJuminBunho); err != nil {
		panic(err)
	}
	if errs := validate.Struct(user); errs != nil {
		panic(errs)
	}
	fmt.Println("OK")
}

func validateUser(sl validator.StructLevel) {
	user := sl.Current().Interface().(User)
	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "fname", "FirstName", "fnameorlname", "")
		sl.ReportError(user.LastName, "lname", "LastName", "fnameorlname", "")
	}
}

func validateJuminBunho(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return regexJuminBunho.MatchString(value)
}
