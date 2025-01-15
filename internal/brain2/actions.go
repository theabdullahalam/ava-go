package brain2

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type ActionMessage struct {
	Name string
	Args []string
}

func GetActionMessageObj(message string) (ActionMessage, bool) {
	var actionMessage ActionMessage
	decoder := json.NewDecoder(strings.NewReader(message))
	err := decoder.Decode(&actionMessage)
	if err != nil {
		fmt.Println(err)
		return ActionMessage{}, false
	}
	return actionMessage, true
}

func (node Node) Run(action string, args []string) string {
	for _, a := range node.Actions {
		if a.Name == action {
			return a.Run(args)
		}
	}
	return "Action not found"
}

func (action Action) Run(args []string) string {
	if action.Type == "script" {
		out, err := exec.Command(action.Name, args...).Output()
		if err != nil {
			return err.Error()
		}
		return string(out)
	}

	return "Could not run action"
}









// internal actions