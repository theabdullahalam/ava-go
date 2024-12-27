package context

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/google/uuid"
)

const CONTEXT_FILE string = "context.json"

func createNewContext() {
	// create the file
	context_file, err := os.Create(CONTEXT_FILE)
	if err != nil {
		fmt.Println(err)
	}
	defer context_file.Close()

	// generate a new UUID for ava
	ava_topic := uuid.New().String()

	context_file.WriteString("{\"ava_topic\": \"" + ava_topic + "\"}")
}

func getContextObject() map[string]interface{} {
	// open the file, creating it if it doesn't exist
	context_file, err := os.Open(CONTEXT_FILE)
	if err != nil {
		context_file.Close()
		createNewContext()
		context_file, _ = os.Open(CONTEXT_FILE)
	}
	defer context_file.Close()

	// read the file
	var context map[string]interface{}
	decoder := json.NewDecoder(context_file)
	if err := decoder.Decode(&context); err != nil {
		fmt.Println(err)
		return context
	}

	return context
}

func GetFromContext(key string) (string, bool) {

	context := getContextObject()

	if value, ok := context[key]; ok { // ok will be false if the key doesn't exist
		if strValue, ok := value.(string); ok { // ok will be false if the value is not a string
			return strValue, true
		}
	}	

	return "", false

}