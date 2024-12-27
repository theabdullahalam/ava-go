package brain

import (
	"encoding/json"
	"fmt"
	"time"
)

type MessageObj struct {
	Sender    string
	Message   string
	Timestamp string
	Source    string
	Target    string
	Type      string
}

func getTimeStampString() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

func NewMessageObj(message string, sender string, target string) MessageObj {
	return MessageObj{
		Sender:    sender,
		Message:   message,
		Timestamp: getTimeStampString(),
		Source:    "user",
		Target:    target,
		Type:      "message",
	}
}

func (messageObj MessageObj) JsonString() (string, bool) {

	jsonString, err := json.Marshal(messageObj)
	if err != nil {
		fmt.Println(err)
		return "{}", false
	}

	return string(jsonString), true
}

func GetResponse(messageObj MessageObj) MessageObj {
	if messageObj.Source == "ava" {
		return messageObj
	}

	return MessageObj{
		Sender:    "Ava",
		Message:   fmt.Sprintf("You said \"%s\"", messageObj.Message),
		Timestamp: getTimeStampString(),
		Source:    "ava",
		Target:    messageObj.Sender,
		Type:      "message",
	}
}
