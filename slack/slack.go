package slack

import (
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

type Configration struct {
	SlackToken string `json:"slack_token"`
	ChannelID  string `json:"channel_id"`
}

func readConf() (string, string) {
	getwd, err := os.Getwd()
	if err != nil {
		return "", ""
	}
	fmt.Println("getwd: ", getwd)
	file, _ := os.Open("autoGit/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration.SlackToken)
	fmt.Println(configuration.ChannelID)
	return configuration.SlackToken, configuration.ChannelID
}

func SendMessageToSlack(msg string) error {
	slackToken, channelID := readConf()

	api := slack.New(slackToken)

	attachment := slack.Attachment{
		Pretext: msg,
	}

	channelID, timestamp, err := api.PostMessage(
		channelID,
		slack.MsgOptionText("", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(false),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return nil
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return nil
}
