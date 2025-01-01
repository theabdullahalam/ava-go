package messages

import (
	"encoding/json"
	"fmt"
	"strings"

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

type ActionObj struct {
	Task string
	target string
	args []string
}


// creates a new MessageObj
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

// converts MessageObj to a jsonstring
func (messageObj MessageObj) JsonString() (string, bool) {

	jsonString, err := json.Marshal(messageObj)
	if err != nil {
		fmt.Println(err)
		return "{}", false
	}

	return string(jsonString), true
}

func (messageObj MessageObj) GetActionObj() string {
	action_string :=strings.Split(strings.Split(messageObj.Message, "```json\n")[1], "```")[0]
	return action_string
}