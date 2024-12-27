package main

import (
	"ava/internal/context"
	"ava/internal/ntfy"
	"ava/internal/brain"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
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
		fmt_string := "Ava: %s\nYou: "
		if ava_response, ok := ntfy.GetMessageFromEvent(scanner.Text()); ok {
			if ava_response.Target == "user" && ava_response.Type == "message" && ava_response.Source == "ava" {
				fmt.Printf(fmt_string, ava_response.Message)
			}
		} 
	}
}

func main() {

	listening := false
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nAva Chat\n----------\nYou: ")

	for {
		if !listening {
			listening = true
			go listen()
		}
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
		ntfy.PublishMessage(brain.NewMessageObj(user_message, "user", "ava"))
		// fmt.Printf("You: ")
	}

	fmt.Printf("\n")

}