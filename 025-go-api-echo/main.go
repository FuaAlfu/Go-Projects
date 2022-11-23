package main

import(
	_"github.com/FuaAlfu/Go-Projects/025-go-api-echo/cmd/api/service"
	"github.com/FuaAlfu/Go-Projects/025-go-api-echo/cmd/api/handlers"

	"github.com/labstack/echo"
)

func main(){
	e := echo.New()
	e.GET("/health-check", handlers.HealthCheackHandler)
	e.GET("/post/:id", handlers.PostSingleHandler)
	e.GET("/posts", handlers.PoastIndexHandler)

	e.Logger.Fatal(e.Start(":1323"))
}