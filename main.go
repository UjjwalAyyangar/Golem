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
	make.AddReaction("who is ujjwal ?", "My sexy creator *_^")
	make.AddReaction("who is shubhodeep mukherjee?", "My sexy creator's best friend and someone who now thinks that he(creator) has lost it ..")
	make.AddReaction("harami", "Tu hai saale!")
	make.Run(botName, botToken, sC, hC)
	log.Println("Running")
}
