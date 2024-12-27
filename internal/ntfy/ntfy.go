package ntfy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	
	"ava/internal/context"
	"ava/internal/brain"
)



func PublishMessage(messageObj brain.MessageObj) {

	jsonString, ok := messageObj.JsonString()
	if !ok {
		return
	}

	ava_topic, _ := context.GetFromContext("ava_topic")
	topic_url := fmt.Sprintf("https://ntfy.sh/%s", ava_topic)
	http.Post(topic_url, "text/plain", strings.NewReader(jsonString))
}

func GetMessageFromEvent(event string) (brain.MessageObj, bool) {

	var messageString string;
	var ok bool;

	var eventObject map[string]interface{}
	decoder := json.NewDecoder(strings.NewReader(event))
	if err := decoder.Decode(&eventObject); err != nil {
		fmt.Println(err)
		return brain.MessageObj{}, false
	}

	if messageString, ok = eventObject["message"].(string); !ok {
		return brain.MessageObj{}, false
	}

	var messageObj brain.MessageObj
	decoder = json.NewDecoder(strings.NewReader(messageString))
	if err := decoder.Decode(&messageObj); err != nil {
		return brain.MessageObj{}, false
	}

	return messageObj, true
}