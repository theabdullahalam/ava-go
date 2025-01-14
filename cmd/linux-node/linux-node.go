package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	// "path/filepath"

	"github.com/theabdullahalam/ava-go/internal/brain2"
	"github.com/theabdullahalam/ava-go/internal/context"
	"github.com/theabdullahalam/ava-go/internal/ntfy2"
)

type ActionMessage struct {
	Name string
	Args []string
}

func newActionMessage(message string) (ActionMessage, bool) {
	var actionMessage ActionMessage
	err := json.Unmarshal([]byte(message), &actionMessage)
	if err != nil {
		fmt.Println(err)
		return ActionMessage{}, false
	}
	return actionMessage, true
}

func processMessage(message string) {
	actionMessage, ok := newActionMessage(message)
	if !ok {
		fmt.Println("User message: " + message)
		return
	}
	fmt.Println("Action message: " + actionMessage.Name)
}

func set(key string, value string) {
	context.SetContext("node.json", key, value)
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
	fmt.Println("Booting up linux node...")
	
	args := os.Args[1:]
	if len(args) == 3 {
		if args[0] == "set" {
			set(args[1], args[2])
			return
		}
	}
	
	listen()

}