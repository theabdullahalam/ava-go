package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	
	"github.com/theabdullahalam/ava-go/internal/context"
	"github.com/theabdullahalam/ava-go/internal/utils"
	"github.com/theabdullahalam/ava-go/internal/ntfy2"

	"github.com/theabdullahalam/ava-go/internal/brain2"
)

func handleResponse(message string) {
	fmt_string := "Ava: %s\nYou: "
	
	tags, err := utils.ExtractTags(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, tag := range tags {
		if tag.Name == "user" {
			fmt.Printf(fmt_string, tag.Content)
		}
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

		message, ok := ntfy2.GetMessageFromEvent(scanner.Text())

		if !ok {
			continue
		}

		handleResponse(message)
	}
}

func main() {

	ava, ok := brain2.GetAva()
	if !ok {
		fmt.Println("Ava not set")
		return
	}

	reader := bufio.NewReader(os.Stdin)
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
		tagged_user_message := brain2.GetTaggedString(user_message, "ava")

		// send it to ava
		ava.Publish(tagged_user_message)

	}

	fmt.Printf("\n")

}
