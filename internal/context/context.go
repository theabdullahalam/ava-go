package context

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/theabdullahalam/ava-go/internal/utils"
)

func getContextFilePath(context_file string) string {
	ava_dir := utils.GetAvaFolder()
	return filepath.Join(ava_dir, context_file)
}


func createNewContext(context_file string) {

	context_file_path := getContextFilePath(context_file)

	// create the file
	context_file_obj, err := os.Create(context_file_path)
	if err != nil {
		fmt.Println(err)
	}
	defer context_file_obj.Close()

	// create empty context
	context_file_obj.WriteString("{}")
}

func createAvaTopic(context_file string) string {
	topic := uuid.New().String()
	SetContext(context_file, "topic", topic)
	return topic
}


func getContextObject(context_file string) map[string]interface{} {

	// open the file, creating it if it doesn't exist
	context_file_obj, err := os.Open(getContextFilePath(context_file))
	if err != nil {
		context_file_obj.Close()
		createNewContext(context_file)
		createAvaTopic(context_file)
		context_file_obj, _ = os.Open(getContextFilePath(context_file))
	}
	defer context_file_obj.Close()

	// read the file
	var context map[string]interface{}
	decoder := json.NewDecoder(context_file_obj)
	if err := decoder.Decode(&context); err != nil {
		fmt.Println(err)
		return context
	}

	return context
}

func GetFromContext(context_file string, key string) (string, bool) {

	context := getContextObject(context_file)

	if value, ok := context[key]; ok { // ok will be false if the key doesn't exist
		if strValue, ok := value.(string); ok { // ok will be false if the value is not a string
			return strValue, true
		}
	}

	return "", false

}

func SetContext(context_file string, key string, value string) {
	context := getContextObject(context_file)
	if context == nil { 
			return
	}

	// set the value
	context[key] = value

	// write the file
	context_file_path := getContextFilePath(context_file)
	context_file_obj, err := os.OpenFile(context_file_path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
			fmt.Println(err)
			return
	}
	defer context_file_obj.Close()

	// encode and write context object
	encoder := json.NewEncoder(context_file_obj)
	encoder.SetIndent("", "    ") // Add indentation
	if err := encoder.Encode(context); err != nil {
			fmt.Println(err)
			return
	}

}