package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	app_token := os.Getenv("SLACK_APP_TOKEN")
	bot_token := os.Getenv("SLACK_BOT_TOKEN")

	bot := slacker.NewClient(bot_token, app_token) // Creates the bot

	go printCommandEvents(bot.CommandEvents())

	/*
		A command is something a user types in Slack that bot can respond to.
		/hello
		/weather London
		/help
	*/

	bot.Command("age <year>", &slacker.CommandDefinition{
		Description: "birth calculation",
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			year := r.Param("year")
			birth, err := strconv.Atoi(year)
			if err != nil {
				w.Reply("Invalid year")
				return
			}
			age := 2025 - birth
			response := fmt.Sprintf("You are %d years old", age)
			w.Reply(response)
		},
	})

	// Create a new cancellable context from the background context.
	// This context can be used to signal cancellation to functions running with it.
	context, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the Slack bot's event listener, passing the context.
	// The listener will run until the context is cancelled or an error occurs.
	err = bot.Listen(context)
	if err != nil {
		log.Fatal(err)
	}
}

// <-chan circle means: a channel that can only read from
// chan<- circle means: a channel that can only write to.
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {

	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
