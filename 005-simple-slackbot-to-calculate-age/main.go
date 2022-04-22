package main

import(
	"fmt"
	"context"
	"os"
	"log"
	"strconv"

	 "github.com/joho/godotenv"
	"github.com/shomali11/slacker"
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

  //create/use your own tokens, these tokens not exist.
	os.Setenv("SLACK_BOT_TOKEN","xoxb-3401875288404-3399548909602-xYlxrd95sLIa2DFkQ0yK2enI")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A03C44F48U9-3399535889266-94f97b18b9eb03c153a1b58c7c2b380a364c2ff7afd9dac48578bcecf06bbfb5")

	bot := slacker.NewClient(os.Getenv("SLACK_APP_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
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
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}
}