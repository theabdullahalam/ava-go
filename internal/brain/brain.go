package brain

import (
	"fmt"

	"github.com/theabdullahalam/ava-go/internal/brain/llm"
	"github.com/theabdullahalam/ava-go/internal/brain/messages"
	"github.com/theabdullahalam/ava-go/internal/utils"
)

func GetResponse(messageObj messages.MessageObj) messages.MessageObj {
	if messageObj.Source == "ava" {
		return messageObj
	}

	ava_response := llm.GetResponse(messageObj.Message) 
	
	ava_reponse_obj := messages.MessageObj{
		Sender:    "Ava",
		Message:   ava_response,
		Timestamp: utils.GetTimeStampString(),
		Source:    "ava",
		Target:    messageObj.Sender,
		Type:      "message",
	}


	// if there is an action, try to run it
	if ava_reponse_obj.HasAction() {
		actionObj := ava_reponse_obj.GetActionObj()
		result := actionObj.RunAction()
		fmt.Println(result)

	}

	return ava_reponse_obj
}
