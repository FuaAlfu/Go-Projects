package main

import (
   "fmt"
   "log"
   "os"
   _"database/sql"

   "github.com/gofiber/fiber/v2"

)

func homePage(c *fiber.Ctx) error {
	return c.SendString("home page")
 }
 func postHandler(c *fiber.Ctx) error {
	return c.SendString("post")
 }
 
 func putHandler(c *fiber.Ctx) error {
	return c.SendString("put")
 }
 func deleteHandler(c *fiber.Ctx) error {
	return c.SendString("delete")
 }

func connect(){
	app := fiber.New()
	
	app.Get("/", homePage)

	app.Post("/", postHandler)
 
	app.Post("/:name", putHandler) 
 
	app.Delete("/:name", deleteHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func main() {
  connect()
}