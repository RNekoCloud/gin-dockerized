package controller

import (
	"database/sql"

	"github.com/RNekoCloud/gin-dockerized/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var q *sqlc.Queries
var db *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/movie?sslmode=disable"
)

func init() {
	var err error
	db, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		panic(err)
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
			"error":   err,
		})
		panic(err)
	}

	ctx.JSON(200, res)
}
