package brain2

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/theabdullahalam/ava-go/internal/utils"
)

const PROMPT_FILE = "prompt.txt"

type Message struct {
	Content string
	Role string
	Timestamp string
}

// global variable is scary
// replace with actual conversation system
var conversation []Message

func init() {
	InitConversation()
}

func InitConversation() {
	ava_folder := utils.GetAvaFolder()
	prompt_file_path := filepath.Join(ava_folder, PROMPT_FILE)

	data, err := os.ReadFile(prompt_file_path)
	if err != nil {
		fmt.Println(err)
	}
	AddToConversation(string(data), "user")
}

func GetConversation() []Message {
	return conversation
}

func AddUserMessage(user_message string) Message {
	llmResponse := GetLLMResponse(user_message)

	AddToConversation(user_message, "user")
	return AddToConversation(llmResponse, "model")
}

func AddToConversation(message_content string, role string) Message {
	message := Message{Content: message_content, Role: role, Timestamp: utils.GetTimeStampString()}
	conversation = append(conversation, message)
	return message
}

func ClearConversation() {
	conversation = []Message{}
	InitConversation()
}