package main

import (
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"xorm.io/xorm"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})
	db, err := xorm.NewEngine("mysql", "go1:goldMa$k46@tcp(127.0.0.1:3306)/go1")
	if err != nil {
		log.Fatal().Err(err).Msg("[sql.Open] failed")
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal().Err(err).Msg("[db.Close] failed")
		}
	}()
	var users []User
	err = db.Table("users").Where("id < ?", 10).Find(&users)
	if err != nil {
		log.Fatal().Err(err).Msg("[db.Find] failed")
	}
	spew.Dump(users)

	newUser := User{
		Name:     "James Dean",
		Email:    "james@email.address",
		Password: "391a53bfacfad41e634f26cef691328c4added675a5ab698a57a3619ed666aeb",
	}
	n, err := db.Table("users").Insert(&newUser)
	if err != nil {
		log.Fatal().Err(err).Msg("[db.Insert] failed")
	}
	fmt.Printf("%d rows inserted", n)
}
