package messages

import (
	"encoding/json"
	"fmt"

	"github.com/theabdullahalam/ava-go/internal/utils"
)

type MessageObj struct {
	Sender    string
	Message   string
	Timestamp string
	Source    string
	Target    string
	Type      string
}



func NewMessageObj(message string, sender string, target string) MessageObj {
	return MessageObj{
		Sender:    sender,
		Message:   message,
		Timestamp: utils.GetTimeStampString(),
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