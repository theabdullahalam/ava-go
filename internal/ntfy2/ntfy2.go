package ntfy2

import (
	"encoding/json"
	"fmt"
)


func GetMessageFromEvent(event string) (string, bool) {
	var eventObject map[string]interface{}
	err := json.Unmarshal([]byte(event), &eventObject)
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	messageString, ok := eventObject["message"].(string)
	if !ok {
		return "", false
	}
	return messageString, true
}