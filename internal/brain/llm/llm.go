package llm

import (
	go_context "context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"github.com/theabdullahalam/ava-go/internal/brain/messages"
)

func getModel() (go_context.Context, *genai.GenerativeModel, *genai.Client) {
	ctx := go_context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	
	model := client.GenerativeModel("gemini-1.5-flash")
	return ctx, model, client
}

func GetResponse(message string) string {
	
	// ai stuff
	ctx, model, client := getModel()
	defer client.Close()

	// existing convo
	message_objs_convo := messages.GetConversation()
	cs := model.StartChat()

	for _, message_obj := range message_objs_convo {

		role := "user"

		if message_obj.Source == "ava" {
			role = "model"
		}

		cs.History = append(cs.History, &genai.Content{
			Parts: []genai.Part{
				genai.Text(message_obj.Message),
			},
			Role: role,
		})
	}

	resp, err := cs.SendMessage(ctx, genai.Text(message))
	if err != nil {
		log.Fatal(err)
	}


	// potntial response because the part thingy below is something I don't understand and 
	// might be incorrect or empty
	response := ""

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				response += fmt.Sprintf("%s", part)
			}
		}
	}

	return response

}