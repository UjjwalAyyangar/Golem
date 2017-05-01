package bot

import (
	"errors"
	"github.com/nlopes/slack"
	"golang.org/x/net/context"
	"net/http"
	//"regexp"
	"strings"
)

type (
	Channel struct {
		id          string
		description string
		welcome     bool
		special     bool
	}
	Logger  func(message string, args ...interface{})
	ChatBot struct {
		adminName      string
		id             string
		botCall        string
		name           string
		token          string
		users          map[string]string
		predefMessages string
		channels       map[string]Channel
		hclient        *http.Client
		client         *slack.Client
		logf           Logger
		ctx            context.Context
	}
)

var cmdLog Logger

func (bot *ChatBot) Init(rtm *slack.RTM) error {
	bot.logf("Determining the bot / user IDs\n")
	users, err := bot.client.GetUsers()
	if err != nil {
		return err
	}
	bot.users = map[string]string{}
	for _, user := range users {
		if user.IsBot {
			bot.id = user.ID

		} else if user.IsAdmin {
			bot.users[user.Name] = user.ID
			bot.adminName = user.Name
		}
	}
	if bot.id == "" {
		return errors.New("Unable to find bot in the list of users ")
	}

	//How the bot will be called?
	bot.botCall = strings.ToLower("<@" + bot.id + ">")
	users = nil
	bot.logf("Determining the channels ID\n")

	publicChannels, err := bot.client.GetChannels(true)
	//Set to true for excluding the archived channels

	for _, channel := range publicChannels {
		channelName := strings.ToLower(channel.Name)
		if chn, isPresent := bot.channels[channelName]; isPresent {
			chn.id = "#" + channel.ID
			bot.channels[channelName] = chn
		}
	}
	publicChannels = nil

	bot.logf("Determining groups ID \n")
	botGroups, err := bot.client.GetGroups(true)
	for _, group := range botGroups {
		groupName := strings.ToLower(group.Name)
		if chn, ok := bot.channels[groupName]; ok && bot.channels[groupName].id == "" {
			chn.id = group.ID
			bot.channels[groupName] = chn
		}
	}
	botGroups = nil

	bot.logf("Initialized %s with ID %s\n", bot.name, bot.id)

	msgParams := slack.PostMessageParameters{}
	_, _, err = bot.client.PostMessage(bot.users[bot.adminName], "Bot deployed", msgParams)
	if err != nil {
		bot.logf("Deployment failed", err)
	}
	return err
}

func (bot *ChatBot) TeamJoined(event *slack.TeamJoinEvent) {
	message := ` Hello ` + event.User.Name + `,
	You made it!
	Welcome to the GDG VIT Vellore Slack channel.
	The slack app is meant to be used as a medium for communication between the members of GDG VIT Vellore. These are the following channels you could join : `
	for idx, val := range bot.channels {
		if !val.welcome {
			continue
		} else {
			message += `<` + val.id + "|" + idx + `> -> ` + val.description + "\n"

			message += "Enjoy"
			msgParams := slack.PostMessageParameters{AsUser: true, LinkNames: 1}
			_, _, err := bot.client.PostMessage(event.User.ID, message, msgParams)
			if err != nil {
				bot.logf("%s\n", err)
				return
			}
		}
	}
}

func (b *ChatBot) isBotMessage(event *slack.MessageEvent, eventText string) bool {
	prefixes := []string{b.predefMessages, "gdgvit"}
	for _, p := range prefixes {
		if strings.HasPrefix(eventText, p) {
			return true
		}
	}
	return strings.HasPrefix(event.Channel, "D")
}

func (b *ChatBot) trimBot(msg string) string {
	msg = strings.Replace(msg, strings.ToLower(b.predefMessages), "", 1)
	msg = strings.TrimPrefix(msg, "gdgvit")
	msg = strings.Trim(msg, " :\n")
	return msg

}

var botResponse map[string][]string

func SetResponse(a map[string][]string) {
	botResponse = a

}
func addReaction(caller, response string) {
	if _, isPresent := botResponse[caller]; isPresent {
		cmdLog("Could not add caller %s, since it is already present", caller)
	} else {
		botResponse[caller] = append(botResponse[caller], response)
	}
}

func (bot *ChatBot) HandleMessage(event *slack.MessageEvent) {
	if event.BotID != "" || event.User == "" || event.SubType == "bot_message" {
		return
	}
	eventText := strings.Trim(strings.ToLower(event.Text), " \n\r")

	if !bot.isBotMessage(event, eventText) {
		return
	}
	eventText = bot.trimBot(eventText)

	for _, response := range botResponse[eventText] {
		respond(bot, event, response+"\n")
		return
	}
}
func respond(bot *ChatBot, event *slack.MessageEvent, response string) {
	msgParams := slack.PostMessageParameters{AsUser: true}
	_, _, err := bot.client.PostMessage(event.Channel, response, msgParams)
	if err != nil {
		bot.logf("%s\n", err)
	}
}

func NewBot(ctx context.Context, slackBotAPI *slack.Client, httpClient *http.Client, name, token string, log Logger) *ChatBot {
	return &ChatBot{
		ctx:     ctx,
		name:    name,
		token:   token,
		hclient: httpClient,
		client:  slackBotAPI,
		logf:    log,
		channels: map[string]Channel{
			"random":  {description: "For random stuff", welcome: true},
			"general": {description: "For general discussions", welcome: true},
		},
	}
}
