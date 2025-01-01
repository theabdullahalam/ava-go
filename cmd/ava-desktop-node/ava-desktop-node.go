package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/google/uuid"

	"github.com/theabdullahalam/ava-go/internal/context"
)

func init () {
	// create topic if it does not exists
	_, ok := context.GetFromContext("ava-node.json", "topic")
	if !ok {
		context.SetContext("ava-node.json", "topic", uuid.New().String())
	}

	// create node name if it does not exist
	_, ok = context.GetFromContext("ava-node.json", "nodename")
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		context.SetContext("ava-node.json", "nodename", hostname)
	}

	// create actions array if it does not exist
	_, ok = context.GetFromContext("ava-node.json", "actions")
	if !ok {
		context.SetContext("ava-node.json", "actions", "[]")
	}
}

func listen() {
	topic, _ := context.GetFromContext("ava-node.json", "topic")
	
	topic_url := fmt.Sprintf("https://ntfy.sh/%s/json", topic)

	resp, err := http.Get(topic_url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}


func main() {
	args := os.Args[1:]

	// listen if no args provided
	if len(args) == 0 {
		listen()
	}

	if len(args) == 2 {
		// connect
		if args[0] == "connect"	{
			context.SetContext("ava-node.json", "ava", args[1])
			fmt.Println("Ava topic set to " + args[1])
		}
	}


}