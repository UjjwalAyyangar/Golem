# GopherOps
This is a chatops mico-framework made in golang. 
It aims in helping go-programmers write slack bot effortlessly.

<h3> Example :- </h3> <br/>

```
  	botToken := "SLACK-BOT-TOKEN"
	sC, hC := make.CreateBot(botToken)
	botName := "gdgvit"
	make.AddReaction("Hello Bot", "What's up?")
	make.AddReaction("Give me GDG's offical link", "https://developers.google.com/groups/")
	make.AddReaction("Thanks", "You're welcome :) ")

	make.Run(botName, botToken, sC, hC)
```

<h3>Output :- </h3> <br />

![alt text](https://github.com/UjjwalAyyangar/GopherOps/blob/master/screenshot.png)


Lot of features to be added, soon.
