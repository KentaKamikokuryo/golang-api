package main

import (
	"github.com/kentakamikokuryo/golang-api/initializers"
	"github.com/kentakamikokuryo/golang-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
