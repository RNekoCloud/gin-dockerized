package service

import (
	"database/sql"

	"github.com/RNekoCloud/gin-dockerized/db/sqlc"
)

var q *sqlc.Queries
var db *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresq://root:root@localhost:5432/movie?sslmode=disable"
)

func init() {
	var err error
	db, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		panic("Failed to connect database!")
	}

	q = sqlc.New(db)

}
