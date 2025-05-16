package main

import (
	"database/sql"
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
	result, err := db.Exec(
		"INSERT INTO users VALUES (?, ?, ?, ?)",
		3,
		"Jane Doe",
		"jain@email.address",
		"7618385caea382f562305468e586330381aae549829075b78dca7cb0d0aefac6",
	)
	if err != nil {
		log.Fatal().Err(err).Msg("[db.Exec] failed")
	}
	n, err := result.LastInsertId()
	if err != nil {
		log.Fatal().Err(err).Msg("[result.LastInsertId] failed")
	}
	log.Info().Int64("last_insert_id", n).Msg("[insert] succeed")
}
