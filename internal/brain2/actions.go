package brain2

import (
	"os/exec"
)

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