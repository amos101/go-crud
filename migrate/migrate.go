package main

import (
	"github.com/amos101/go-crud/initializers"
	"github.com/amos101/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
