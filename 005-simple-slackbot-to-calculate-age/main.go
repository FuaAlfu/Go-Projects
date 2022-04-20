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

	bot.Command("my yob is <year>", &slacker.CommandDefinition){
		Description: "yob calculator",
		Example: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil{
				println("error.. ")
			}
			age := 2022 - yob
			r := fmt.Sprintf("age is %d ", age)
			response.Reply(r)
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}
}