package llm

import (
	go_context "context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// var model *genai.GenerativeModel
// var ctx go_context.Context

func init() {
	
}

func GetResponse(message string) string {

	ctx := go_context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(message))
	if err != nil {
		log.Fatal(err)
	}


	// potntial response because the part thingy below is something I don't understand and 
	// might be incorrect or empty
	var potential_response string

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				potential_response = fmt.Sprintf("%s", part)
			}
		}
	}

	return potential_response

}