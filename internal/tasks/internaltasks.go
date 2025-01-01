package tasks

import (
	"fmt"
)

func RunInternalTask(name string, args []string) string {

	response := "Task not found"

	switch name {
	case "SayThis":
		response = SayThis(args[0])
	}

	return response
}

func SayThis(s string) string{
	fmt.Println(s)
	return "Said " + s
}

func GetIpAddress() string {
	return "127.0.0.1"
}