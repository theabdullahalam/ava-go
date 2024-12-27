package main

import (
	"bufio"
	"log"
	"net/http"
	"fmt"

	"ava/internal/context"
	"ava/internal/ntfy"
	"ava/internal/brain"
)

func listen() {

	ava_topic, _ := context.GetFromContext("ava_topic")
	topic_url := fmt.Sprintf("https://ntfy.sh/%s/json", ava_topic)

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
			fmt.Printf("Recieved message: %s\n", messageObj.Message)
			responseObj := brain.GetResponse(messageObj)
			ntfy.PublishMessage(responseObj)
			fmt.Printf("Sent response: %s\n", responseObj.Message)
			continue
		}

	}
}

func main() {
	listen()
}