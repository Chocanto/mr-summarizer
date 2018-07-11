package main

import (
	"encoding/json"
	"strings"
	"net/http"
	"time"
	"fmt"
	"log"
)

type Message struct {
	Username string `json:"username"`
	Channel  string `json:"channel"`
	Text     string `json:"text"`
}

func sendToMattermost(message string) {
	data := Message{
		Username: Config.Mattermost.Username,
		Channel:  Config.Mattermost.Channel,
		Text:     fmt.Sprintf("### %s - %s\n%s", Printer.Sprintf("MERGE REQUESTS WAITING FOR APPROVAL"), time.Now().Format("02/01/06"), message),
	}

	messageJSON, _ := json.Marshal(data)
	body := strings.NewReader(string(messageJSON))
	req, _ := http.NewRequest("POST", Config.Mattermost.Webhook, body)

	log.Println("Sending message to Mattermost...")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Message successfully sent !")

	defer resp.Body.Close()
}
