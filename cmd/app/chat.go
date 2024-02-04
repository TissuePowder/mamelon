package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
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

func (app *application) getGptTLDR(prompt string) (string, error) {

	model := app.config.Chat.OpenAI.Model
	url := app.config.Chat.OpenAI.Url
	key := app.config.Chat.OpenAI.Key
	// system := app.config.Chat.OpenAI.System

	requestPayload := ChatRequest{
		Model: model,
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "system",
				Content: "The prompt contains last 50 messages from a channel of a discord server centered to the manga/anime 'boku no kokoro no yabai yatsu'. Summarize it and give a short but precise TLDR of what have been discussed in the channel.",
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

func (app *application) getQAResponse(text string) (string, error) {

	condaEnvPath := app.config.Python.Conda
	pythonScriptPath := app.config.Python.Script

	cmd := exec.Command(condaEnvPath, pythonScriptPath, text)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil

}
