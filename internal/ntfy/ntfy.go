package ntfy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/theabdullahalam/ava-go/internal/brain/messages"
)

func PublishMessage(messageObj messages.MessageObj, topic string) {

	jsonString, ok := messageObj.JsonString()
	if !ok {
		return
	}
	topic_url := fmt.Sprintf("https://ntfy.sh/%s", topic)
	http.Post(topic_url, "text/plain", strings.NewReader(jsonString))
}

func GetMessageFromEvent(event string) (messages.MessageObj, bool) {

	var messageString string
	var ok bool

	var eventObject map[string]interface{}
	decoder := json.NewDecoder(strings.NewReader(event))
	if err := decoder.Decode(&eventObject); err != nil {
		fmt.Println(err)
		return messages.MessageObj{}, false
	}

	if messageString, ok = eventObject["message"].(string); !ok {
		return messages.MessageObj{}, false
	}

	var messageObj messages.MessageObj
	decoder = json.NewDecoder(strings.NewReader(messageString))
	if err := decoder.Decode(&messageObj); err != nil {
		return messages.MessageObj{}, false
	}

	return messageObj, true
}
