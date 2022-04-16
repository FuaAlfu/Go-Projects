package main

import(
	"fmt"
	"context"
	"os"
	"log"

	 "github.com/joho/godotenv"
	_"github.com/shomali11/slacker"
)

func main() {
	fmt.Println("loged.. !")
	err := godotenv.Load()
    if err != nil {
    log.Fatal("Error loading .env file")
  }

	os.Setenv("SLACK_BOT_TOKEN","")
	os.Setenv("SLACK_APP_TOKEN","")
}