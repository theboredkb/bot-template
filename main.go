package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/theboredkb/bot-template/internal/bot"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("env could not be loaded: %v", err)
	}
}

func main() {
	token := os.Getenv("BOT_TOKEN")
	bot, err := bot.New(token)
	if err != nil {
		log.Panicf("Error when creating the bot: %v", err)
	}

	bot.RegisterHandlers()

	s := bot.Session
	err = s.Open()
	if err != nil {
		log.Panicf("Cannot open the session: %v", err)
	}

	if err = bot.RegisterCommands(); err != nil {
		log.Panicf("Error when starting the bot: %v", err)
	}

	bot.PrintCommands()

	defer s.Close()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("Press Ctrl+C to exit")
	<-stop
}
