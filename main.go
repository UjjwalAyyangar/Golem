package main

import (
	"GopherOps/make"
	"log"
)

func main() {
	botToken := "xoxb-157875480528-KivR2yET9sCVYK7kcojcO608"
	sC, hC := make.CreateBot(botToken)
	botName := "gdgvit"

	make.AddReaction("framework", "Are the most awesome things, don't you think?")
	make.Run(botName, botToken, sC, hC)
	log.Println("Running")
}
