package make

import (
	"GopherOps/bot"
	//"fmt"
	//"github.com/gorilla/mux"
	"github.com/nlopes/slack"
	"golang.org/x/net/context"
	"log"
	"net"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func CreateBot(botToken string) (*slack.Client, *http.Client) {
	slackClient := slack.New(botToken)
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
	return slackClient, httpClient
}
func AddReaction(caller, response string) {
	botResponse := make(map[string][]string)
	caller = strings.ToLower(caller)
	if _, isPresent := botResponse[caller]; isPresent {
		log.Printf("Could not add caller %s, since it is already present \n", caller)
	} else {
		a := []string{}
		a = append(a, response)
		botResponse[caller] = a //append(botResponse[caller], response)
	}
	bot.SetResponse(botResponse, caller)
}
func WelcomeMessage(msg string){
	bot.SetWelcome(msg)
}

func Run(botName, botToken string, slackClient *slack.Client, httpClient *http.Client) {
	botRTM := slackClient.NewRTM()
	go botRTM.ManageConnection()
	runtime.Gosched()

	ctx := context.Background()
	b := bot.NewBot(ctx, slackClient, httpClient, botName, botToken, log.Printf)
	if err := b.Init(botRTM); err != nil {
		panic(err)
	}

	go func() {
		for msg := range botRTM.IncomingEvents {
			switch message := msg.Data.(type) {
			case *slack.MessageEvent:
				go b.HandleMessage(message)

			case *slack.TeamJoinEvent:
				go b.TeamJoined(message)
			}
		}
	}()

	/*go func() {
		r := mux.NewRouter()
		r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, info)
		}).
			Name("info")
		Methods("GET")

		s := http.Server{
			Addr:         ":8081",
			Handler:      r,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		log.Fatal(s.ListenAndServe())
	}()*/
	log.Println("%s is now running", botName)
	select {}
}
