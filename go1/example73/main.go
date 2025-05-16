package main

import (
	"os"
	"time"

	"example73/src/handler"
	"example73/src/mdware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"xorm.io/xorm"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})
	db, err := xorm.NewEngine("mysql", "go1:goldMa$k46@tcp(127.0.0.1:3306)/go1")
	if err != nil {
		log.Fatal().Err(err).Msg("[xorm.NewEngine] failed")
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello World!!")
	})
	g := e.Group("/user")
	g.Use(middleware.BasicAuth(mdware.AuthValidator()))
	g.Use(mdware.DbMiddleware(db))
	g.GET("/:id", handler.UserFindByIdHandler())
	g.GET("/", handler.UserListHandler())
	if err := e.Start(":1323"); err != nil {
		log.Fatal().Err(err).Msg("[e.Start] failed")
	}
}
