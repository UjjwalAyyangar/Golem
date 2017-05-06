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
	make.AddReaction("Give me GDG's official link", "https://developers.google.com/groups/")
	make.AddReaction("Thanks", "You're welcome :) ")
	make.WelcomeMessage("Welcome to our team. May the force be with you my friend")
	make.Run(botName, botToken, sC, hC)
	log.Println("Running")
}
