package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})
	db, err := sql.Open("mysql", "go1:goldMa$k46@tcp(127.0.0.1:3306)/go1")
	if err != nil {
		log.Fatal().Err(err).Msg("[sql.Open] failed")
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal().Err(err).Msg("[db.Close] failed")
		}
	}()
	var id int
	var name string
	rows, err := db.Query("SELECT id, name FROM users where id < ?", 10)
	if err != nil {
		log.Fatal().Err(err).Msg("[db.Query] failed")
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal().Err(err).Msg("[rows.Close] failed")
		}
	}()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Error().Err(err).Msg("[rows.Scan] failed")
		}
		fmt.Printf("%d: %s\n", id, name)
	}
}
