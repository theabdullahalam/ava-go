package messages

const CONVERSATION_FILE string = "conversation.json"

var conversation []MessageObj

func GetConversation() []MessageObj {
	return conversation
}

func AddToConversation(messageObj MessageObj) {
	conversation = append(conversation, messageObj)
}