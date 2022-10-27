package main

import(
	"fmt"
	"os"
	"context"
	"log"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4295137471761-4276079093830-S6vP8Yh0Q63aqEFvFYVlv6XX")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A047W2WQDP1-4282573585939-375a7837108a74d1acff871e56d79933710dba69f40e5bcd86176a4a6744c972")


	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	
	go printCommandEvents(bot.CommandEvents())
	
	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	})

	bot.Command("Hii", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Hello! How're you today")
		},
	})

	bot.Command("I'm fine", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Well! great to have a person like you in conversation")
		},
	})

	bot.Command("How're you?", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("I'm great! my developers are quite interactive with me.")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}