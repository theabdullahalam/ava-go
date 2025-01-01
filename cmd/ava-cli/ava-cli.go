package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/theabdullahalam/ava-go/internal/brain/messages"
	"github.com/theabdullahalam/ava-go/internal/context"
	"github.com/theabdullahalam/ava-go/internal/ntfy"
)

func handleResponse(ava_response messages.MessageObj) {
	fmt_string := "Ava: %sYou: "
	if ava_response.Target == "user" && ava_response.Type == "message" && ava_response.Source == "ava" {
		fmt.Printf(fmt_string, ava_response.GetMessageOnly())
	}
}

func listen() {

	topic, _ := context.GetFromContext("ava.json", "topic")
	topic_url := fmt.Sprintf("https://ntfy.sh/%s/json", topic)

	resp, err := http.Get(topic_url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if ava_response, ok := ntfy.GetMessageFromEvent(scanner.Text()); ok {
			handleResponse(ava_response)
		}
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	topic, _ := context.GetFromContext("ava.json", "topic")
	fmt.Printf("\nAva Chat\n----------\nYou: ")
	go listen()

	for {
		
		user_message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		// do nothing if the user just presses enter
		if user_message == "\n" {
			continue
		}

		// quit
		if user_message == "\\q\n" {
			break
		}

		// if it is a regular message,
		user_message = user_message[:len(user_message)-1]

		// send it to ava
		ntfy.PublishMessage(messages.NewMessageObj(user_message, "user", "ava"), topic)

	}

	fmt.Printf("\n")

}
