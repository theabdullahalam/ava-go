package messages

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/theabdullahalam/ava-go/internal/utils"
)

const CONVERSATION_FILE string = "conversation.json"
const PROMPT_FILE = "prompt.txt"

var conversation []MessageObj

func init() {
	ava_folder := utils.GetAvaFolder()
	prompt_file_path := filepath.Join(ava_folder, PROMPT_FILE)

	data, err := os.ReadFile(prompt_file_path)
	if err != nil {
		fmt.Println(err)
	}
	AddToConversation(NewMessageObj(string(data), "user", "ava"))
	AddToConversation(NewMessageObj("Task List:\n Empty", "user", "ava"))
}

func GetConversation() []MessageObj {
	return conversation
}

func AddToConversation(messageObj MessageObj) {
	conversation = append(conversation, messageObj)
}