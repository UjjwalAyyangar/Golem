# Golem
This is a chatops mico-framework made in golang. 
It aims in helping go-programmers write slack bots effortlessly.

<h3> Example :- </h3> <br/>

```
  	botToken := "SLACK-BOT-TOKEN"
	sC, hC := make.CreateBot(botToken)
	botName := "gdgvit"
	make.AddReaction("Hello Bot", "What's up?")
	make.AddReaction("Give me GDG's official link", "https://developers.google.com/groups/")
	make.AddReaction("Thanks", "You're welcome :) ")
	make.WelcomeMessage("Welcome to our team. May the force be with you my friend")
	make.Run(botName, botToken, sC, hC)

```

<h3>Output :- </h3> <br />

<h4> Sample conversation </h4> <br />

![alt text](https://github.com/UjjwalAyyangar/GopherOps/blob/master/sample.png)

<h4> Welcome message </h4> <br/>

![alt text](https://github.com/UjjwalAyyangar/GopherOps/blob/master/welcome.png)




Lots of features to be added, soon.
