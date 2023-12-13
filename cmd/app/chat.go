package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ChatRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type ChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (app *application) getGptResponse(prompt string) (string, error) {

	if !app.config.Chat.Public {
		return "*Mamelon is currently sleeping and can't talk now. But your message will be passed on to the bot maintainers. No worries.*", nil
	}

	model := app.config.Chat.OpenAI.Model
	url := app.config.Chat.OpenAI.Url
	key := app.config.Chat.OpenAI.Key
	system := app.config.Chat.OpenAI.System

	requestPayload := ChatRequest{
		Model: model,
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "system",
				Content: system,
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	payloadBytes, err := json.Marshal(requestPayload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return "", err
	}
	defer resp.Body.Close()

	var cr ChatResponse
	err = json.NewDecoder(resp.Body).Decode(&cr)
	if err != nil {
		return "", err
	}
	return cr.Choices[0].Message.Content, nil
}
