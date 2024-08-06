package LLM

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"projects/config"
	"projects/utils/math"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

const (
	apiEndpoint   = "https://api.openai.com/v1/chat/completions"
	generalPrompt = "This request is automated please respond carefully or it might break the system. Only generate output in the specified format. Give answers to the best of your ability. "
)

func UsernameSuggestion() string {
	apiKey := os.Getenv("OPENAI_API")
	client := resty.New()

	var (
		usernamePrompt = fmt.Sprintf("Task: Generate five creative 1 word usernames and also use these digits %d in random fashion for username generation! Output format : 1. name1 \n2. name2 \n3. name3\n4. name4\n5. name5\n", math.RandomInt())
	)

	finalPrompt := generalPrompt + usernamePrompt

	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":      config.GPT_MODEL,
			"messages":   []interface{}{map[string]interface{}{"role": "system", "content": finalPrompt}},
			"max_tokens": 50,
		}).
		Post(apiEndpoint)

	if err != nil {
		log.Fatalf("Error while sending the request: %v", err)
	}

	body := response.Body()

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error while decoding JSON response:", err)
		return ""
	}

	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	return content
}
