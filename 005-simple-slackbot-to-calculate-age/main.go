package main

import(
	"fmt"
	"context"
	"os"
	"log"

	 "github.com/joho/godotenv"
	_"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for e := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(e.Timestamp)
		fmt.Println(e.Command)
		fmt.Println(e.Parameters)
		fmt.Println(e.Event)
		fmt.Println()
	}
}

func main() {
	fmt.Println("loged.. !")
	err := godotenv.Load()
    if err != nil {
    log.Fatal("Error loading .env file")
  }

	os.Setenv("SLACK_BOT_TOKEN","")
	os.Setenv("SLACK_APP_TOKEN","")

	bot := slacker.NewClient(os.Getenv("SLACK_APP_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
}