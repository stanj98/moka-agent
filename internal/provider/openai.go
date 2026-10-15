package provider

//Moka uses a lightweight adapter layer to abstract away LLM endpoints. OpenAI/OpenRouter-compatible client adapter in Go

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Client struct {
	APIKey string
	Model  string
}

func NewClient() *Client {
	model := os.GetEnv("MOKA_MODEL")
	if model == "" {
		model = "gpt-4o-mini"
	}
	return &Client{
		APIKey: os.Getenv("OPENAI_API_KEY"),
		Model:  model,
	}
}

func (c *Client) Chat(messages []Message) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	payload := map[string]interface{}{
		"model":    c.Model,
		"messages": messages,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch completion from provider")
	}

	var result struct {
		Choices []struct {
			Message `json:"message"`
		} `json:"choices`
	}
	json.NewDecoder(resp.Body).Decode(&result)

	if len(result.Choices) > 0 {
		return result.Choices[0].Message.Content, nil
	}
	return "", errors.New("no response choices returned")
}
