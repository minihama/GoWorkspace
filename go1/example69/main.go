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
	var name, email string
	row := db.QueryRow("SELECT name, email FROM users where id = ?", 1)
	err = row.Scan(&name, &email)
	if err != nil {
		log.Fatal().Err(err).Msg("[db.Scan] failed")
	}
	fmt.Printf("name:%q, email:%q\n", name, email)
}
