package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/theabdullahalam/ava-go/internal/brain2"
	"github.com/theabdullahalam/ava-go/internal/context"
	"github.com/theabdullahalam/ava-go/internal/ntfy2"
)


func processMessage(message string) {
	actionMessage, ok := brain2.GetActionMessageObj(message)
	if !ok {
		fmt.Println("User message: " + message)
		return
	}
	fmt.Println("Running action: " + actionMessage.Name)
	node, ok := brain2.GetThisNode()
	if !ok {
		fmt.Println("Could not find node.json")
		return
	}
	result := node.Run(actionMessage.Name, actionMessage.Args)
	taggedResult := brain2.GetTaggedString(result, "result")

	ava, ok := brain2.GetAva()
	if !ok {
		fmt.Println("Ava not set")
	}

	ava.Publish(taggedResult)
}

func set(key string, value string) {
	context.SetContext("node.json", key, value)
	fmt.Println("Set " + key + " to " + value)
}

func listen() {

	fmt.Println("Listening for messages...")

	node, ok := brain2.GetThisNode()
	if !ok {
		fmt.Println("Could not find node.json")
		return
	}
	topic := node.Topic
	topic_url := fmt.Sprintf("https://ntfy.sh/%s/json", topic)

	resp, err := http.Get(topic_url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		message, exists := ntfy2.GetMessageFromEvent(scanner.Text())
		if !exists {
			continue
		}
		processMessage(message)
	}
}

func main() {
	
	args := os.Args[1:]
	if len(args) == 3 {
		if args[0] == "set" {
			set(args[1], args[2])
			return
		}
	}
	
	fmt.Println("Booting up linux node...")
	listen()

}