package openai

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

func AskGPT(question string) (string, error) {
	client := resty.New()

	// Definir a URL da API Groq (substitua com o URL real se necessário)
	url := "https://api.groq.com/openai/v1/chat/completions"

	// Enviar a solicitação com a chave da API Groq
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+os.Getenv("GROQ_API_KEY")). // Usando a chave da Groq
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model": "mistral", // Use o modelo que a Groq oferece
			"messages": []map[string]string{
				{"role": "system", "content": "Você é um especialista em Path of Exile 2. Responda como um jogador veterano, com foco em builds, farm e progressão."},
				{"role": "user", "content": question},
			},
		}).
		Post(url)

	if err != nil {
		return "", err
	}

	// Log para ver o que está sendo retornado pela API
	fmt.Println("Resposta da API:", string(resp.Body()))

	// Estrutura da resposta da API Groq (ajuste conforme necessário)
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

	fmt.Println("Resposta decodificada:", res)

	if len(res.Choices) == 0 {
		return "Desculpe, não consegui entender sua pergunta ou a resposta não foi gerada corretamente.", nil
	}

	return res.Choices[0].Message.Content, nil
}
