package messages

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/theabdullahalam/ava-go/internal/utils"
	"github.com/theabdullahalam/ava-go/internal/tasks"
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
	Name string
	Args []string
}


func (actionObj ActionObj) RunAction() string{
	task := tasks.GetTask(actionObj.Name)
	return task.Run(actionObj.Args)
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

func (messageObj MessageObj) HasAction() bool {
	return strings.Contains(messageObj.Message, "\n```json\n") 
}

func (messageObj MessageObj) GetActionObj() ActionObj {

	if !messageObj.HasAction() {
		return ActionObj{}
	}

	action_string :=strings.Split(strings.Split(messageObj.Message, "```json\n")[1], "```")[0]

	var actionObj ActionObj
	decoder := json.NewDecoder(strings.NewReader(action_string))
	if err := decoder.Decode(&actionObj); err != nil {
		fmt.Println(err)
	}

	return actionObj
}

func (messageObj MessageObj) GetMessageOnly() string {

	if !messageObj.HasAction() {
		return messageObj.Message
	}

	parts := strings.Split(messageObj.Message, "```json\n")
	if strings.Contains(parts[0], "```json\n"){
		return parts[1]
	} else {
		return parts[0]
	}
}