package main

import(
	"fmt"
	"os"

	"github.com/FuaAlfu/Go-Projects/020-fullstack-tracker/server/routes"
	"github.com/gin-contrip/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == ""{
		port = ":8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Default())

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries")
	router.GET("/entry/:id")
	router.GET("ingredient/:ingredient ")
	router.GET()
}