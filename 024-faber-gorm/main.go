package main

import(
	"github.com/gofiber/fiber/v2"
	_"net/http"
	"log"
	"fmt"
)

func hi(c *fiber.Ctx) error{
	return c.SendString("hi..")
}

func run(){
	app := fiber.New()

	app.Get("/api", hi)
	log.Fatal(app.Listen(":3000"))
}

func main() {
	run()
	fmt.Println("testing..")
}