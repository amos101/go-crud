package main

import (
	// "fmt"

	"github.com/amos101/go-crud/controllers"
	"github.com/amos101/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConectToDB()
}

func main() {
	r := gin.Default()

	// create
	r.POST("/posts", controllers.PostsCreate)

  // update
	r.PUT("/posts/:id", controllers.PostsUpdate)

	// read
	r.GET("/posts", controllers.Postsindex)
	r.GET("/posts/:id", controllers.PostsShow)

	// delete
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}