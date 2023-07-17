package main

import (
	controller "github.com/RNekoCloud/gin-dockerized/controller/movie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Route GIN
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to my app!",
		})
	})

	// Add movie record
	r.POST("/movie", controller.AddMovie)

	// Fetching movies record
	r.GET("/movies", controller.GetMovies)
	// Fetching movie by its ID
	r.GET("/movies/:id", controller.GetMovie)

	// Initialize Gin App
	r.Run(":3003")
}
