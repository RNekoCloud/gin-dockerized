package controller

import (
	"database/sql"

	"github.com/RNekoCloud/gin-dockerized/db/sqlc"
	"github.com/gin-gonic/gin"
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

func AddMovie(ctx *gin.Context) {
	res, err := q.CreateMovie(ctx, sqlc.CreateMovieParams{
		Title: ctx.PostForm("title"),
		Genre: ctx.PostForm("genre"),
	})

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Failed to add new record",
		})
	}

	ctx.JSON(200, res)
}
