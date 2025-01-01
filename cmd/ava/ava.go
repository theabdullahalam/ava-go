package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"

	"github.com/theabdullahalam/ava-go/internal/brain"
	"github.com/theabdullahalam/ava-go/internal/brain/messages"
	"github.com/theabdullahalam/ava-go/internal/context"
	"github.com/theabdullahalam/ava-go/internal/ntfy"
)

func listen() {

	topic, _ := context.GetFromContext("ava.json","topic")
	topic_url := fmt.Sprintf("https://ntfy.sh/%s/json", topic)

	resp, err := http.Get(topic_url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		messageObj, ok := ntfy.GetMessageFromEvent(scanner.Text())

		if !ok {
			continue
		}

		if messageObj.Target == "ava" && messageObj.Type == "message" && messageObj.Source == "user" {
			fmt.Printf("Recieved message!\n")
			messages.AddToConversation(messageObj)

			responseObj := brain.GetResponse(messageObj)
			ntfy.PublishMessage(responseObj, topic)
			messages.AddToConversation(responseObj)

			fmt.Printf("Sent response!\n")

			continue
		}

	}
}

func main() {
	listen()
}
