package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"

	"github.com/theabdullahalam/ava-go/internal/brain2"
	"github.com/theabdullahalam/ava-go/internal/context"
	"github.com/theabdullahalam/ava-go/internal/ntfy2"
	"github.com/theabdullahalam/ava-go/internal/utils"
)

func listen() {

	ava, ok := brain2.GetAva()
	if !ok {
		fmt.Println("Ava not set")
		return
	}
	topic, _ := context.GetFromContext("ava.json","topic")
	topic_url := fmt.Sprintf("https://ntfy.sh/%s/json", topic)

	resp, err := http.Get(topic_url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {

		// only respond if a message exists in the event
		message, ok := ntfy2.GetMessageFromEvent(scanner.Text())
		if !ok {
			continue
		}

		// only respond if the message is intended for ava or is a result
		user_message := ""
		tags, err := utils.ExtractTags(message)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, tag := range tags {
			if tag.Name == "ava" || tag.Name == "result" {
				user_message += fmt.Sprintf("%s\n\n", tag.Content)
			}
		}

		if user_message == "" {
			continue
		}


		// process user message
		response := brain2.AddUserMessage(user_message)
		_, ok = brain2.GetAva()
		if !ok {
			fmt.Println("Ava not set")
			continue
		}

		fmt.Println(response.Content)
		ava.Publish(response.Content)

	}
}

func main() {
	listen()
}
