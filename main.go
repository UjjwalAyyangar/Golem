package main

import (
	"GopherOps/make"
	"log"
)

func main() {
	botToken := "SLACK-BOT-TOKEN"
	sC, hC := make.CreateBot(botToken)
	botName := "gdgvit"
	make.AddReaction("Hello Bot", "What's up?")
	make.AddReaction("Give me GDG's offical link", "https://developers.google.com/groups/")
	make.AddReaction("Thanks", "You're welcome :) ")

	make.Run(botName, botToken, sC, hC)
	log.Println("Running")
}
