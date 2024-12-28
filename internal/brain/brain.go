package brain

import (
	"github.com/theabdullahalam/ava-go/internal/brain/messages"
	"github.com/theabdullahalam/ava-go/internal/utils"
	"github.com/theabdullahalam/ava-go/internal/brain/llm"
)

func GetResponse(messageObj messages.MessageObj) messages.MessageObj {
	if messageObj.Source == "ava" {
		return messageObj
	}

	return messages.MessageObj{
		Sender:    "Ava",
		Message:   llm.GetResponse(messageObj.Message),
		Timestamp: utils.GetTimeStampString(),
		Source:    "ava",
		Target:    messageObj.Sender,
		Type:      "message",
	}
}
