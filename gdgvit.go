package main

import (
	"GDGVitSlack-Bot/bot"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nlopes/slack"
	"golang.org/x/net/context"
	"log"
	"net"
	"net/http"
	"runtime"
	"strings"
	"time"
)

var (
	botVersion = "HEAD"
	info       = `{ "version": "` + botVersion + `"}`
)

func main() {
	botName := "@gdgvit"
	slackBotToken := "xoxb-157875480528-KivR2yET9sCVYK7kcojcO608"
	slackBotAPI := slack.New(slackBotToken)
	httpClient := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   15 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   5 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	botName = strings.TrimPrefix(botName, "@")
	slackBotRTM := slackBotAPI.NewRTM()
	go slackBotRTM.ManageConnection()
	runtime.Gosched()

	ctx := context.Background()
	//projectID := "gdg-slack-bot"
	b := bot.NewBot(ctx, slackBotAPI, httpClient, botName, slackBotToken, log.Printf)

	if err := b.Init(slackBotRTM); err != nil {
		panic(err)
	}
	//_, err := b.GetLastSeenCl()
	//if err != nil {
	//	log.Printf("got error : %v\n", err)
	//	panic(err)
	//}
	go func() {
		for msg := range slackBotRTM.IncomingEvents {
			switch message := msg.Data.(type) {
			case *slack.MessageEvent:
				go b.HandleMessage(message)

			case *slack.TeamJoinEvent:
				go b.TeamJoined(message)

			}

		}
	}()

	go func() {
		r := mux.NewRouter()
		r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, info)
		}).
			Name("info").
			Methods("GET")

		s := http.Server{
			Addr:         ":8081",
			Handler:      r,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		log.Fatal(s.ListenAndServe())

	}()
	log.Println("GDGVIT bot is now running")
	select {}

}
