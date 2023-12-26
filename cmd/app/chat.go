package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"slices"
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

func (app *application) getGptResponse(authorID string, prompt string) (string, error) {

	isPrivilegedUser := false
	if slices.Contains(app.config.Bot.PrivilegedUsers, authorID) {
		isPrivilegedUser = true
	}

	if isPrivilegedUser && prompt == "sleep" {
		app.config.Chat.Public = false
		return "*Mamelon falls asleep*", nil
	}

	if isPrivilegedUser && prompt == "wake up" {
		app.config.Chat.Public = true
		return "Meow!", nil
	}

	if !app.config.Chat.Public {
		return "*Mamelon is sleeping and can't talk right now. Your message will be passed on to the bot maintainers.*", nil
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
