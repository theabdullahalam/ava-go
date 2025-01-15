package ntfy2

import (
	"encoding/json"
	"fmt"
	"strings"
)



func GetMessageFromEvent(event string) (string, bool) {
	var eventObject map[string]interface{}
	decoder := json.NewDecoder(strings.NewReader(event))

	err := decoder.Decode(&eventObject)
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