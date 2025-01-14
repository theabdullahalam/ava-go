package tasks

import (
	"fmt"

	"github.com/theabdullahalam/ava-go/internal/nodes"
)

func RunInternalTask(name string, args []string) string {

	response := "Task not found"

	switch name {
	case "SayThis":
		response = SayThis(args[0])
	case "AddNode":
		response = AddNode(args[0])
	}


	return response
}

func SayThis(s string) string{
	fmt.Println(s)
	return "Said " + s
}

func AddNode(topic string) string {
	nodes.AddNode(nodes.NewNode("-", topic))
	return "Added node " + topic
}