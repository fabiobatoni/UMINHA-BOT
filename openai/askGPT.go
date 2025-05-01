package openai

import (
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
)

func AskGPT(question string) (string, error) {
	client := resty.New()

	url := "https://api.groq.com/openai/v1/chat/completions"

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+os.Getenv("GROQ_API_KEY")).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model": "llama-3.3-70b-versatile",
			"messages": []map[string]string{
				{"role": "system", "content": "Você é um especialista em Path of Exile 2. Responda como um jogador veterano, com foco em builds, farm e progressão."},
				{"role": "user", "content": question},
			},
		}).
		Post(url)

	if err != nil {
		return "", err
	}

	type Choice struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	}
	type Response struct {
		Choices []Choice `json:"choices"`
	}
	var res Response

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return "", err
	}

	if len(res.Choices) == 0 {
		return "Desculpe, não consegui entender sua pergunta ou a resposta não foi gerada corretamente.", nil
	}

	return res.Choices[0].Message.Content, nil
}
